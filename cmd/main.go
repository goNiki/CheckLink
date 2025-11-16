package main

import (
	"context"
	"goNiki/CheckLink/internal/app"
	"goNiki/CheckLink/internal/domain"
	"goNiki/CheckLink/internal/http/middleware"
	"goNiki/CheckLink/internal/infrastructure/logger/sl"
	"goNiki/CheckLink/pkg/errorsAPP"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	c, err := app.NewConteiner("http://localhost:8081")
	if err != nil {
		panic(errorsAPP.ErrCreateDIContainer)
	}

	r := chi.NewRouter()

	r.Use(middleware.Cors)

	r.Use(middleware.GracefulShutdownMiddlleware(c.TaskService))

	r.Post("/links", c.LinksHandler.CheckLink)
	r.Post("/links/report", c.LinksHandler.GetReportLinks)

	srv := http.Server{
		Addr:    ":8081",
		Handler: r,
	}
	go func() {
		time.Sleep(10 * time.Second)
		if err := c.TaskService.ProcessPendingTasks(context.Background()); err != nil {
			c.Log.Error("Ошибка обработки отложенных задач", sl.Error(err))
		}
	}()

	go func() {
		c.Log.Info("Сервер запущен на 8081 порту")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			c.Log.Error("Ошибка запуска сервера", sl.Error(err))
			os.Exit(1)
		}
	}()

	ctxsheduler, shedulercancel := context.WithCancel(context.Background())

	go func() {
		if err := c.Scheduler.Start(ctxsheduler); err != nil && err != context.Canceled {
			c.Log.Error("Планировщик завершил работу из-за ошибки ", sl.Error(err))
		} else {
			c.Log.Info("Планировщик остановлен")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	domain.SetDraining(true)
	c.Log.Info("Режим draining активирован - новые запросы будут сохранены")

	time.Sleep(30 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	c.Log.Info("Завершение работы сервера")

	// Graceful shutdown - ждем завершения текущих запросов
	if err := srv.Shutdown(ctx); err != nil {
		c.Log.Error("Ошибка при завершении сервера", sl.Error(err))
	}

	shedulercancel()
	c.Log.Info("Сохранение данных перед завершением")
	if err := c.SaveService.SaveLinksToFile(ctx); err != nil {
		c.Log.Error("Ошибка сохранения данных: ", sl.Error(err))
	}
	if err := c.SaveService.SaveTasksToFile(ctx); err != nil {
		c.Log.Error("Ошибка сохранения данных: ", sl.Error(err))
	}

	c.Log.Info("Сервер остановлен")
}
