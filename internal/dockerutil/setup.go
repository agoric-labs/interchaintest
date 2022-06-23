package dockerutil

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

// labelKey is a key for docker labels. Used when cleaning up docker resources.
const labelKey = "ibctest"

// DockerSetup sets up a new dockertest.Pool (which is a client connection
// to a Docker engine) and configures a network associated with t.
//
// If any part of the setup fails, t.Fatal is called.
func DockerSetup(t *testing.T) (*dockertest.Pool, string) {
	t.Helper()

	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatalf("failed to create dockertest pool: %v", err)
	}

	// Clean up docker resources at end of test.
	t.Cleanup(dockerCleanup(t.Name(), pool))

	// Also eagerly clean up any leftover resources from a previous test run,
	// e.g. if the test was interrupted.
	dockerCleanup(t.Name(), pool)()

	network, err := pool.Client.CreateNetwork(docker.CreateNetworkOptions{
		Name:           fmt.Sprintf("ibctest-%s", RandLowerCaseLetterString(8)),
		Options:        map[string]interface{}{},
		Labels:         map[string]string{labelKey: t.Name()},
		CheckDuplicate: true,
		Internal:       false,
		EnableIPv6:     false,
		Context:        context.Background(),
	})
	if err != nil {
		t.Fatalf("failed to create docker network: %v", err)
	}

	return pool, network.ID
}

// dockerCleanup will clean up Docker containers, networks, and the other various config files generated in testing
func dockerCleanup(testName string, pool *dockertest.Pool) func() {
	return func() {
		showContainerLogs := os.Getenv("SHOW_CONTAINER_LOGS")
		cont, _ := pool.Client.ListContainers(docker.ListContainersOptions{All: true})
		for _, c := range cont {
			for k, v := range c.Labels {
				if k == labelKey && v == testName {
					_ = pool.Client.StopContainer(c.ID, 10)
					ctxWait, cancelWait := context.WithTimeout(context.Background(), time.Second*5)
					defer cancelWait()
					_, _ = pool.Client.WaitContainerWithContext(c.ID, ctxWait)
					if showContainerLogs != "" {
						stdout := new(bytes.Buffer)
						stderr := new(bytes.Buffer)
						ctxLogs, cancelLogs := context.WithTimeout(context.Background(), time.Second*5)
						defer cancelLogs()
						_ = pool.Client.Logs(docker.LogsOptions{Context: ctxLogs, Container: c.ID, OutputStream: stdout, ErrorStream: stderr, Stdout: true, Stderr: true, Tail: "50", Follow: false, Timestamps: false})
					}
					_ = pool.Client.RemoveContainer(docker.RemoveContainerOptions{ID: c.ID, Force: true})
					break
				}
			}
		}
		nets, _ := pool.Client.ListNetworks()
		for _, n := range nets {
			for k, v := range n.Labels {
				if k == labelKey && v == testName {
					_ = pool.Client.RemoveNetwork(n.ID)
					break
				}
			}
		}
	}
}
