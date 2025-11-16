package checker

import (
	"context"
	"goNiki/CheckLink/internal/domain"
	"goNiki/CheckLink/internal/infrastructure/logger/sl"
	"sync"
	"time"
)

// TODO сделать канал для вывода всех ошибок связанных с проверкой.
func (s *service) CheckBatch(ctx context.Context, urls []string) (domain.LinkBatch, error) {

	resultChan := make(chan domain.Link, len(urls))
	errorChan := make(chan error, len(urls))

	var wg sync.WaitGroup

	semaphore := make(chan struct{}, 10)

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()

			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			time.Sleep(50 * time.Millisecond)

			link, err := s.CheckLink(ctx, u)
			if err != nil {
				link.Status = domain.StatusNotAvailable
				errorChan <- err
			}
			resultChan <- link
		}(url)
	}

	wg.Wait()
	close(resultChan)
	close(errorChan)

	errorResult := make([]error, 0, len(urls))
	links := make(map[string]domain.LinkStatus, len(urls))
	for res := range resultChan {
		links[res.URL] = res.Status
	}

	for err := range errorChan {
		errorResult = append(errorResult, err)
	}

	for _, err := range errorResult {
		s.log.Error("error", sl.Error(err))
	}

	number := s.linkstorage.NextID()

	linkBatch := domain.LinkBatch{
		Links:  links,
		Number: number,
	}

	err := s.linkstorage.SaveLinks(ctx, &linkBatch)
	if err != nil {
		return domain.LinkBatch{}, err
	}

	return linkBatch, nil

}
