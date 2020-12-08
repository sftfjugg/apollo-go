package services

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/repositories"
	models2 "go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.didapinche.com/time"
)

type RoleService interface {
	Create(role *models.Role) error
	Find(appId, userId string) (int, error)
	FindByAppId(appId string) (*models.Role, error)
}

type roleService struct {
	repository repositories.RoleRepository
	db         *gorm.DB
}

func NewRoleService(
	repository repositories.RoleRepository,
	db *gorm.DB,
) RoleService {
	return &roleService{
		db:         db,
		repository: repository,
	}
}

func (s roleService) Create(role *models.Role) error {
	roles := make([]*models2.Role, 0)
	for _, r := range role.Release {
		release := new(models2.Role)
		release.AppId = role.AppId
		release.UserID = r.UserId
		release.UserName = r.UserName
		release.DataChange_CreatedTime = time.Now()
		release.DataChange_CreatedBy = role.Operator
		release.Level = 2
		roles = append(roles, release)
	}
	for _, r := range role.Write {
		write := new(models2.Role)
		write.AppId = role.AppId
		write.UserID = r.UserId
		write.UserName = r.UserName
		write.DataChange_CreatedTime = time.Now()
		write.DataChange_CreatedBy = role.Operator
		write.Level = 2
		roles = append(roles, write)
	}
	db := s.db.Begin()
	if err := s.repository.Delete(db, role.AppId); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call RoleRepository.deleted failed")
	}
	if err := s.repository.Creates(db, roles); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call RoleRepository.create failed")
	}
	db.Commit()
	return nil
}

func (s roleService) Find(appId, userId string) (int, error) {

	roles, err := s.repository.Find(appId, userId)
	if err != nil {
		return 0, errors.Wrap(err, "call RoleSitory.Find failed")
	}
	i := 0
	for _, r := range roles {
		i = +r.Level
	}
	if i >= 4 {
		i = 4
	}
	return i, nil
}

func (s roleService) FindByAppId(appId string) (*models.Role, error) {
	roles, err := s.repository.FindByAppId(appId)
	if err != nil {
		return nil, errors.Wrap(err, "call RoleSitory.Find failed")
	}
	rs := new(models.Role)
	rs.AppId = appId
	write := make([]*models.User, 0)
	release := make([]*models.User, 0)
	for _, r := range roles {
		user := new(models.User)
		user.UserName = r.UserName
		user.UserId = r.UserID
		if r.Level == 1 {
			write = append(write, user)
		}
		if r.Level == 2 {
			release = append(release, user)
		}
	}
	rs.Write = write
	rs.Release = release
	return rs, nil

}
