package services

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin/repositories"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
)

type ReleaseHistoryService interface {
	Find(appId, namespaceName, key string, pageSize, pageNum int) ([]*models.ReleaseHistory, error)
}

type releaseHistoryService struct {
	repository repositories.ReleaseHistoryRepository
	db         *gorm.DB
}

func NewReleaseHistoryService(
	db *gorm.DB,
	repository repositories.ReleaseHistoryRepository,
) ReleaseHistoryService {
	return &releaseHistoryService{
		db:         db,
		repository: repository,
	}
}

func (s releaseHistoryService) Find(appId, namespaceName, key string, pageSize, pageNum int) ([]*models.ReleaseHistory, error) {
	releaseHistorys, err := s.repository.Find(appId, namespaceName, key, pageSize, pageNum)
	if err != nil {
		return nil, errors.Wrap(err, "call ReleaseHistoryReposotory.Find() error")
	}
	return releaseHistorys, nil
}
