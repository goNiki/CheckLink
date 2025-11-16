package scheduler

import (
	"context"
	"goNiki/CheckLink/internal/infrastructure/logger/sl"
	"goNiki/CheckLink/internal/services"
	"log/slog"
	"sync"
	"time"
)

type Scheduler struct {
	Interval time.Duration
	Worker   services.SaveService
	log      *slog.Logger
}

func NewScheduler(interval time.Duration, worker services.SaveService, log *slog.Logger) *Scheduler {
	return &Scheduler{
		Interval: interval,
		Worker:   worker,
		log:      log,
	}
}

func (s *Scheduler) Start(ctx context.Context) error {
	ticket := time.NewTicker(s.Interval)

	defer ticket.Stop()

	var wg sync.WaitGroup

	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			return ctx.Err()
		case <-ticket.C:
			wg.Add(1)
			go func() {
				defer wg.Done()
				defer func() {
					if r := recover(); r != nil {
						s.log.Error("Паника в работе планировщика: ", r)
					}
				}()

				jobCtx, cancel := context.WithTimeout(ctx, 4*time.Minute)
				defer cancel()

				if err := s.Worker.SaveLinksToFile(jobCtx); err != nil {
					s.log.Error("worker Error: ", sl.Error(err))
				}
			}()
		}
	}
}
