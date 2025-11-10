package main

import (
	linkshandler "goNiki/CheckLink/internal/http/handler/links"
	"goNiki/CheckLink/internal/infrastructure/logger"
	"goNiki/CheckLink/internal/infrastructure/logger/sl"
	linkchecker "goNiki/CheckLink/internal/services/LinkChecker"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	log := logger.NewLogger()
	log.Info("logger initialized")

	linkchecker := linkchecker.NewLinkChecker(&http.Client{})

	handler := linkshandler.NewLinksHandler(log, linkchecker)

	r := chi.NewRouter()
	r.Post("/links", handler.CheckLink)

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
