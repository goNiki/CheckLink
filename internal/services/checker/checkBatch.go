package checker

import (
	"context"
	"goNiki/CheckLink/internal/domain"
	"sync"
)

func (s *service) CheckBatch(ctx context.Context, urls []string) (domain.LinkBatch, error) {

	resultChan := make(chan domain.Link, len(urls))

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			link, err := s.CheckLink(ctx, u)
			if err != nil {
				link.Status = domain.StatusNotAvailable
			}
			resultChan <- link
		}(url)
	}

	wg.Wait()
	close(resultChan)

	links := make(map[string]domain.LinkStatus)
	for res := range resultChan {
		links[res.URL] = res.Status
	}

	number := s.linkstorage.NextID()

	linkBatch := domain.LinkBatch{
		Links:  links,
		Number: number,
	}

	err := s.linkstorage.SaveDate(&linkBatch)
	if err != nil {
		return domain.LinkBatch{}, err
	}

	return linkBatch, nil

}
