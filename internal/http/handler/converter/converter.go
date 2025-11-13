package converter

import (
	"goNiki/CheckLink/internal/domain"
	"goNiki/CheckLink/internal/dto"
)

func LinkBatchToResponce(l *domain.LinkBatch) *dto.ResponseCheckLink {
	return &dto.ResponseCheckLink{
		Links:    *LinksToDTOLinks(&l.Links),
		LinksNum: l.Number,
	}
}

func LinksToDTOLinks(m *map[string]domain.LinkStatus) *map[string]string {
	res := make(map[string]string, len(*m))

	for u, s := range *m {
		res[u] = s.String()
	}

	return &res

}
