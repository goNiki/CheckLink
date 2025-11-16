package filestorage

import "goNiki/CheckLink/internal/storage/models"

type storage struct {
	filepathLinks string
	filepathTasks string
}

func NewFileStorage() *storage {
	return &storage{
		filepathLinks: models.DateDir + "/" + models.StorageLinksFile,
		filepathTasks: models.DateDir + "/" + models.StorageTasksFile,
	}
}
