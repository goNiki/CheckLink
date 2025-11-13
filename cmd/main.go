package main

import (
	linkshandler "goNiki/CheckLink/internal/http/handler/links"
	"goNiki/CheckLink/internal/http/handler/middleware"
	"goNiki/CheckLink/internal/infrastructure/logger"
	"goNiki/CheckLink/internal/infrastructure/logger/sl"
	"goNiki/CheckLink/internal/services/checker"
	"goNiki/CheckLink/internal/services/linkreport"
	"goNiki/CheckLink/internal/storage/filestorage"
	linksstorage "goNiki/CheckLink/internal/storage/links"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	log := logger.NewLogger()
	log.Info("logger initialized")

	filestorage := filestorage.NewFileStorage()

	linksstorage := linksstorage.NewStorage(filestorage)

	linkchecker := checker.NewChecker(&http.Client{}, linksstorage)
	reportservice := linkreport.NewLinkReportService(linksstorage)

	handler := linkshandler.NewLinksHandler(log, linkchecker, reportservice)

	r := chi.NewRouter()

	// CORS middleware
	r.Use(middleware.Cors)

	r.Post("/links", handler.CheckLink)
	r.Post("/links/report", handler.GetReportLinks)

	srv := http.Server{
		Addr:    ":8081",
		Handler: r,
	}
	log.Info("Поднимаю сервер на 8081")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Error("Ошибка запуска сервера: ", sl.Error(err))
	}

	log.Info("Сервер остановлен")

}
