package app

import (
	"goNiki/CheckLink/internal/client"
	httpclient "goNiki/CheckLink/internal/client/http"
	"goNiki/CheckLink/internal/http/handler/links"
	"goNiki/CheckLink/internal/infrastructure/logger"
	"goNiki/CheckLink/internal/infrastructure/logger/sl"
	"goNiki/CheckLink/internal/infrastructure/scheduler"
	"goNiki/CheckLink/internal/services"
	"goNiki/CheckLink/internal/services/checker"
	"goNiki/CheckLink/internal/services/linkreport"
	"goNiki/CheckLink/internal/services/save"
	"goNiki/CheckLink/internal/services/task"
	"goNiki/CheckLink/internal/storage"
	"goNiki/CheckLink/internal/storage/filestorage"
	linksstorage "goNiki/CheckLink/internal/storage/links"
	"goNiki/CheckLink/internal/storage/tasks"
	"log/slog"
	"time"
)

type Container struct {
	Log        *slog.Logger
	HTTPClient client.HTTPClient

	FileStorage  storage.FileStorage
	LinksStorage storage.LinksStorage
	TasksStorage storage.TasksStorage

	LinksCheckerService services.LinksChecker
	ReportService       services.ReportService
	SaveService         services.SaveService
	TaskService         services.TaskService

	Scheduler scheduler.Scheduler

	LinksHandler links.LinksHandler
}

func NewConteiner(baseURL string) (*Container, error) {
	c := &Container{}

	c.Log = logger.NewLogger()

	c.HTTPClient = httpclient.NewHttpClient(baseURL)

	c.FileStorage = filestorage.NewFileStorage()

	var err error
	c.LinksStorage, err = linksstorage.NewLinksStorage(c.FileStorage)
	if err != nil {
		c.Log.Error("LinksStorage Error: ", sl.Error(err))
	}

	c.TasksStorage, err = tasks.NewTaskStorage(c.FileStorage)
	if err != nil {
		c.Log.Error("TasksStorage Error: ", sl.Error(err))
	}

	c.LinksCheckerService = checker.NewLinksChecker(c.HTTPClient, c.LinksStorage, c.Log)
	c.ReportService = linkreport.NewReportService(c.LinksStorage)
	c.SaveService = save.NewSaveService(c.LinksStorage, c.TasksStorage, c.FileStorage)
	c.TaskService = task.NewTasksService(c.TasksStorage, c.HTTPClient)

	c.Scheduler = *scheduler.NewScheduler(30*time.Second, c.SaveService, c.Log)

	c.LinksHandler = *links.NewLinksHandler(c.Log, c.LinksCheckerService, c.ReportService)

	return c, nil
}
