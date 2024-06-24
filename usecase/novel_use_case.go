package usecase

import (
	"errors"
	"redis_gorm_fiber/domain"
	"redis_gorm_fiber/model"
)

type novelUseCase struct {
	novelRepo domain.NovelRepo
}

// GetNovelById implements domain.NovelUseCase.
func (n *novelUseCase) GetNovelById(id int) (model.Novel, error) {

	res, err := n.novelRepo.GetNovelById(id)
	if err != nil {
		return model.Novel{}, errors.New("failed to get novel:" + err.Error())
	}

	return res, nil
}

// CreateNovel implements domain.NovelUseCase.
func (n *novelUseCase) CreateNovel(createNovel model.Novel) error {

	err := n.novelRepo.CreateNovel(createNovel)
	if err != nil {

		return errors.New("failed to create novel:" + err.Error())
	}
	return nil
}

func (n *novelUseCase) DeleteNovel(id int) error {
	err := n.novelRepo.DeleteNovel(id)
	if err != nil {
		return errors.New("failed to delete novel:" + err.Error())
	}
	return nil
}

func (n *novelUseCase) UpdateNovel(id int, updateNovel model.Novel) error {
	err := n.novelRepo.UpdateNovel(id, updateNovel)
	if err != nil {
		return errors.New("failed to update novel:" + err.Error())
	}
	return nil
}

func NewNovelUseCase(novelRepo domain.NovelRepo) domain.NovelUseCase {
	return &novelUseCase{novelRepo: novelRepo}
}
