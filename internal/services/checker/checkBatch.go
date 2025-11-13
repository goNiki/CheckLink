package checker

import (
	"context"
	"errors"
	"fmt"
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

	links := make(map[string]domain.LinkStatus, len(urls))
	for res := range resultChan {
		links[res.URL] = res.Status
	}

	number := s.linkstorage.NextID()

	linkBatch := domain.LinkBatch{
		Links:  links,
		Number: number,
	}

	err := s.linkstorage.SaveDate(ctx, &linkBatch)
	if err != nil {
		return domain.LinkBatch{}, err
	}

	err = s.linkstorage.SaveInFile()
	if err != nil {
		return linkBatch, fmt.Errorf("%w: %v", errors.New("Error save"), err)
	}

	return linkBatch, nil

}
