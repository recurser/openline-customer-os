package neo4j

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"time"
)

const username = "neo4j"
const password = "new-s3cr3t"

func startContainer(ctx context.Context, username, password string) (testcontainers.Container, error) {
	request := testcontainers.ContainerRequest{
		Image:        "neo4j:5-community",
		ExposedPorts: []string{"7687/tcp"},
		Env:          map[string]string{"NEO4J_AUTH": fmt.Sprintf("%s/%s", username, password)},
		WaitingFor:   wait.ForLog("Bolt enabled").WithStartupTimeout(300 * time.Second),
	}
	return testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: request,
		Started:          true,
	})
}

func InitTestNeo4jDB() (testcontainers.Container, *neo4j.DriverWithContext) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	var err error
	var neo4jContainer testcontainers.Container
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		neo4jContainer, err = startContainer(ctxWithTimeout, username, password)
		if err == nil {
			// success, break out of loop
			break
		}
		// error occurred
		if i == maxRetries-1 {
			// last attempt failed, panic
			log.Panic(err)
		}
		// wait for 1 second before retrying
		time.Sleep(1 * time.Second)
	}

	port, err := neo4jContainer.MappedPort(ctxWithTimeout, "7687")
	if err != nil {
		log.Panic(err)
	}
	address := fmt.Sprintf("bolt://localhost:%d", port.Int())
	driver, err := neo4j.NewDriverWithContext(address, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		log.Panic(err)
	}
	return neo4jContainer, &driver
}

func CloseDriver(driver neo4j.DriverWithContext) {
	err := driver.Close(context.Background())
	if err != nil {
		log.Panic("Neo4j driver should close")
	}
}

func Terminate(container testcontainers.Container, ctx context.Context) {
	err := container.Terminate(ctx)
	if err != nil {
		log.Fatal("Container should stop")
	}
}
