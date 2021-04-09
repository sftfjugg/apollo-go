package services

import (
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/repositories"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
)

type DingdingService interface {
	Create(dingding *models.Dingding) error
	FindAll(pageNum, pageSize int) ([]*models.Dingding, int, error)
	Update(dingding *models.Dingding) error
	Delete(id int) error
	Find(Type, deptName, env string, level int) (*models.Dingding, error)
}

type dingdingService struct {
	repository repositories.DingdingRepository
}

func NewDingdingService(repository repositories.DingdingRepository) DingdingService {
	return &dingdingService{repository: repository}
}

func (d dingdingService) Create(dingding *models.Dingding) error {
	isRepeat := d.repository.FindByName(dingding.Name)
	if !isRepeat {
		return errors.New("dingding exists now")
	}
	dingding2, err := d.Find(dingding.Type, dingding.DeptName, dingding.Env, dingding.Level)
	if err != nil {
		return errors.Wrap(err, "call dingdingService.Find error")
	}
	if dingding2.Token != "" {
		return errors.New("dingding exists now")
	}
	if err := d.repository.Create(dingding); err != nil {
		return errors.Wrap(err, "call dingdingService.Create error")
	}
	return nil
}

func (d dingdingService) FindAll(pageNum, pageSize int) ([]*models.Dingding, int, error) {
	dingdings, err := d.repository.FindAll(pageNum, pageSize)
	if err != nil {
		return nil, 0, errors.Wrap(err, "call dingdingService.FindAll error")
	}
	count, err := d.repository.FindCount()
	if err != nil {
		return nil, 0, errors.Wrap(err, "call dingdingService.FindAll error")
	}
	return dingdings, count, nil
}

func (d dingdingService) Update(dingding *models.Dingding) error {
	dingding2, err := d.Find(dingding.Type, dingding.DeptName, dingding.Env, dingding.Level)
	if err != nil {
		return errors.Wrap(err, "call dingdingService.Find error")
	}
	if dingding2.Token != "" {
		return errors.New("dingding exists now")
	}
	if err := d.repository.Update(dingding); err != nil {
		return errors.Wrap(err, "call dingdingService.Update error")
	}
	return nil
}

func (d dingdingService) Delete(id int) error {
	if err := d.repository.Delete(id); err != nil {
		return errors.Wrap(err, "call dingdingService.Delete error")
	}
	return nil
}

func (d dingdingService) Find(Type, deptName, env string, level int) (*models.Dingding, error) {
	dingding, err := d.repository.Find(Type, deptName, env, level)
	if err != nil {
		return nil, errors.Wrap(err, "call dingdingService.Find error")
	}
	return dingding, nil
}
