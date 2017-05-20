package integrationtests

import (
	"fmt"
	"os"
	"testing"

	"github.com/RichardKnop/machinery/v1/config"
)

func TestRedisMemcache(t *testing.T) {
	redisURL := os.Getenv("REDIS_URL")
	memcacheURL := os.Getenv("MEMCACHE_URL")
	if redisURL == "" || memcacheURL == "" {
		return
	}

	// Redis broker, Redis result backend
	server := setup(&config.Config{
		Broker:        fmt.Sprintf("redis://%v", redisURL),
		DefaultQueue:  "test_queue",
		ResultBackend: fmt.Sprintf("memcache://%v", memcacheURL),
	})
	worker := server.NewWorker("test_worker")
	go worker.Launch()
	testSendTask(server, t)
	testSendGroup(server, t)
	testSendChord(server, t)
	testSendChain(server, t)
	testReturnJustError(server, t)
	testReturnMultipleValues(server, t)
	testPanic(server, t)
	worker.Quit()
}