package linkreport

import (
	"context"
	"fmt"
	"goNiki/CheckLink/pkg/errorsAPP"
	"goNiki/CheckLink/pkg/pdf"
	"io"
)

func (s *service) CreateReport(ctx context.Context, numbers []int64) (io.Reader, error) {

	linkBatch, err := s.linksstorage.GetByIDs(ctx, numbers)
	if err != nil {
		return nil, fmt.Errorf("%w, %v", errorsAPP.ErrInternalDB, err)
	}

	r := pdf.New()

	r.CreateRepors("Report", linkBatch)

	readers, err := r.OutputToReader()
	if err != nil {
		return nil, err
	}

	return readers, nil

}
