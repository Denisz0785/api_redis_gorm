package domain

import (
	"redis_gorm_fiber/model"
)

// NovelRepo is an interface that defines the methods for creating and getting novels.
type NovelRepo interface {
	// CreateNovel creates a new novel in the repository.
	CreateNovel(createNovel model.Novel) error

	// GetNovelById retrieves a novel by its ID from the repository.
	GetNovelById(id int) (model.Novel, error)

	// DeleteNovel deletes a novel from the repository.
	DeleteNovel(id int) error

	// UpdateNovel updates a novel in the repository.
	UpdateNovel(id int, updateNovel model.Novel) error
}

// NovelUseCase is an interface that defines the methods for creating and getting novels.
type NovelUseCase interface {
	// CreateNovel creates a new novel in the use case.
	// It calls the CreateNovel method of the repository.
	CreateNovel(createNovel model.Novel) error

	// GetNovelById retrieves a novel by its ID from the use case.
	// It calls the GetNovelById method of the repository.
	GetNovelById(id int) (model.Novel, error)

	// DeleteNovel deletes a novel from the use case.
	// It calls the DeleteNovel method of the repository.
	DeleteNovel(id int) error

	// UpdateNovel updates a novel in the use case.
	// It calls the UpdateNovel method of the repository.
	UpdateNovel(id int, updateNovel model.Novel) error
}
