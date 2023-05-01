package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/config"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/service"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"sync"
	"time"
)

type TaskQueue struct {
	name          string
	taskFunctions []func()
	mutex         sync.Mutex
	waitGroup     sync.WaitGroup
}

func (t *TaskQueue) AddTask(function func()) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.taskFunctions = append(t.taskFunctions, function)
}

func (t *TaskQueue) RunTasks() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if len(t.taskFunctions) == 0 {
		logrus.Warn("No task found, exiting")
		return
	}
	for _, task := range t.taskFunctions {
		t.waitGroup.Add(1)
		go func(fn func()) {
			defer t.waitGroup.Done()
			fn()
		}(task)
	}
	t.taskFunctions = nil
	t.waitGroup.Wait()
}

func main() {
	cfg := loadConfiguration()
	config.InitLogger(cfg)

	// init openline postgres db client
	sqlDb, gormDb, errPostgres := config.NewPostgresClient(cfg)
	if errPostgres != nil {
		logrus.Fatalf("failed opening connection to postgres: %v", errPostgres.Error())
	}
	defer sqlDb.Close()

	ctx := context.Background()
	// init openline neo4j db client
	neo4jDriver, errNeo4j := config.NewDriver(cfg)
	if errNeo4j != nil {
		logrus.Fatalf("failed opening connection to neo4j: %v", errNeo4j.Error())
	}
	defer (*neo4jDriver).Close(ctx)

	// init airbyte postgres db pool
	airbyteStoreDb := config.InitPoolManager(cfg)

	services := service.InitServices(neo4jDriver, gormDb, airbyteStoreDb)

	services.InitService.Init()

	var taskQueueSyncCustomerOsData = &TaskQueue{name: "Sync Customer OS Data"}
	var taskQueueSyncToEventStore = &TaskQueue{name: "Sync Neo4j Data to EventStore"}

	go runTaskQueue(taskQueueSyncCustomerOsData, cfg.SyncCustomerOsData.TimeoutAfterTaskRun, []func(){
		func() {
			runId, _ := uuid.NewRandom()
			logrus.Infof("run id: %s syncing data into customer-os at %v", runId.String(), time.Now().UTC())
			services.SyncCustomerOsDataService.Sync(ctx, runId.String())
			logrus.Infof("run id: %s sync completed at %v", runId.String(), time.Now().UTC())
		},
	})
	if cfg.SyncToEventStore.Enabled {
		go runTaskQueue(taskQueueSyncToEventStore, cfg.SyncToEventStore.TimeoutAfterTaskRun, []func(){
			func() {
				services.SyncToEventStoreService.SyncEmails(ctx)
			},
		})
	}

	select {}
}

func runTaskQueue(taskQueue *TaskQueue, timeoutAfterTaskRun int, taskFuncs []func()) {
	for {
		for _, v := range taskFuncs {
			taskQueue.AddTask(v)
		}

		taskQueue.RunTasks()

		// Wait for a fixed amount of time before running the tasks again
		timeout := time.Second * time.Duration(timeoutAfterTaskRun)
		logrus.Infof("waiting %v seconds before next run for %s", timeout.Seconds(), taskQueue.name)
		time.Sleep(timeout)
	}
}

func loadConfiguration() *config.Config {
	if err := godotenv.Load(); err != nil {
		logrus.Warn("Failed loading .env file")
	}

	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		logrus.Errorf("%+v", err)
	}

	return &cfg
}
