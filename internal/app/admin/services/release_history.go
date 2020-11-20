package services

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	models2 "go.didapinche.com/foundation/apollo-plus/internal/app/admin/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin/repositories"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
)

type ReleaseHistoryService interface {
	Find(appId, namespaceName, cluster, key string, pageSize, pageNum int) (*models2.ReleaseHistoryPage, error)
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

func (s releaseHistoryService) Find(appId, namespaceName, cluster, key string, pageSize, pageNum int) (*models2.ReleaseHistoryPage, error) {
	releaseHistorys, err := s.repository.Find(appId, namespaceName, cluster, key, pageSize, pageNum)
	if err != nil {
		return nil, errors.Wrap(err, "call ReleaseHistoryReposotory.Find() error")
	}
	total, err := s.repository.FindCount(appId, namespaceName, cluster, key)
	if err != nil {
		return nil, errors.Wrap(err, "call ReleaseHistoryReposotory.FindCount() error")
	}
	releaseHistoryPage := new(models2.ReleaseHistoryPage)
	releaseHistoryPage.Total = total
	releaseHistory2 := make([]*models2.ReleaseHistory, 0)
	for _, r := range releaseHistorys {
		releaseHistory := new(models2.ReleaseHistory)
		releaseHistory.AppId = r.AppId
		releaseHistory.Operation = r.Operation
		releaseHistory.ClusterName = r.ClusterName
		releaseHistory.BranchName = r.BranchName
		releaseHistory.NamespaceName = r.NamespaceName
		releaseHistory.DataChange_CreatedBy = r.DataChange_CreatedBy
		releaseHistory.DataChange_CreatedTime = r.DataChange_CreatedTime
		releaseHistory.Id = r.Id
		items1 := make([]*models.Item, 0)
		items2 := make([]*models.Item, 0)
		if err := json.Unmarshal([]byte(r.OperationContext), &items1); err != nil {
			return nil, errors.Wrap(err, "call json.Unmarshal([]byte(r.OperationContext) error")

		}
		releaseHistory.OperationContext = items1
		if err := json.Unmarshal([]byte(r.ReleaseContext), &items2); err != nil {
			return nil, errors.Wrap(err, "call json.Unmarshal([]byte(r.ReleaseContext) error")

		}
		releaseHistory.ReleaseContext = items2
		releaseHistory2 = append(releaseHistory2, releaseHistory)
	}
	releaseHistoryPage.ReleaseHistory = releaseHistory2

	return releaseHistoryPage, nil
}
