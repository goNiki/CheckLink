package save

import "goNiki/CheckLink/internal/storage"

type service struct {
	linkstorage storage.LinksStorage
	taskstorage storage.TasksStorage
	filestorage storage.FileStorage
}

func NewSaveService(linkstorage storage.LinksStorage, taskstorage storage.TasksStorage, filestorage storage.FileStorage) *service {
	return &service{
		linkstorage: linkstorage,
		taskstorage: taskstorage,
		filestorage: filestorage,
	}
}
