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
	CreateBackDoor(userId string) error
	DeleteByUserId(userId string) error
	Find(appId, userId, cluster, env string) (*models.Auth, error)
	Finds(userId, cluster, env string) (map[string][]*models.NamespaceRole, error)
	FindByAppId(appId, cluster, env, name string) (*models.Role, error)
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

func (s roleService) CreateBackDoor(userId string) error {
	if userId == "" {
		return errors.New("don't exist userId")
	}
	role := new(models2.Role)
	role.AppId = "root"
	role.Level = 4
	role.UserID = userId
	role.DataChange_CreatedTime = time.Now()
	db := s.db.Begin()
	if err := s.repository.Create(db, role); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call RoleRepository.create failed")
	}
	db.Commit()
	return nil
}

func (s roleService) Create(role *models.Role) error {
	roles := make([]*models2.Role, 0)
	for _, r := range role.Release {
		release := new(models2.Role)
		release.AppId = role.AppId
		release.Namespace = role.Namespace
		release.Env = role.Env
		release.Cluster = role.Cluster
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
		write.Cluster = role.Cluster
		write.Namespace = role.Namespace
		write.Env = role.Env
		write.UserID = r.UserId
		write.UserName = r.UserName
		write.DataChange_CreatedTime = time.Now()
		write.DataChange_CreatedBy = role.Operator
		write.Level = 1
		roles = append(roles, write)
	}
	db := s.db.Begin()
	if err := s.repository.Delete(db, role.AppId, role.Cluster, role.Env, role.Namespace); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call RoleRepository.deleted failed")
	}
	if len(roles) > 0 {
		if err := s.repository.Creates(db, roles); err != nil {
			db.Rollback()
			return errors.Wrap(err, "call RoleRepository.creates failed")
		}
	}
	db.Commit()
	return nil
}

//前端全局搜索用
func (s roleService) Finds(userId, cluster, env string) (map[string][]*models.NamespaceRole, error) {
	m := make(map[string][]*models.NamespaceRole)
	roles, err := s.repository.Finds(userId, cluster, env)
	if err != nil {
		return nil, errors.Wrap(err, "call RoleSitory.Find failed")
	}
	for _, role := range roles {
		if n, ok := m[role.AppId]; ok {
			for i, _ := range n {
				if n[i].Name == role.Namespace {
					n[i].Level = n[i].Level + role.Level
					break
				}
				if i == len(n)-1 {
					namespaceRole := new(models.NamespaceRole)
					namespaceRole.Name = role.Namespace
					namespaceRole.Level = role.Level
					n = append(n, namespaceRole)
					break
				}
			}
		} else {
			namespaceRole := new(models.NamespaceRole)
			namespaceRole.Name = role.Namespace
			namespaceRole.Level = role.Level
			m[role.AppId] = append(m[role.AppId], namespaceRole)
		}
	}
	return m, nil
}

//前端控制按钮用
func (s roleService) Find(appId, userId, cluster, env string) (*models.Auth, error) {
	auths := new(models.Auth)
	namespaceRole := make([]*models.NamespaceRole, 0)
	m := make(map[string]int)
	roles, err := s.repository.Find(appId, userId, cluster, env)
	if err != nil {
		return nil, errors.Wrap(err, "call RoleSitory.Find failed")
	}
	j := 0
	for _, r := range roles {
		if _, ok := m[r.Namespace]; !ok {
			namespace := new(models.NamespaceRole)
			namespace.Name = r.Namespace
			namespace.Level = r.Level
			if r.Level >= 4 {
				auths.IsOwner = true
				r := make([]*models.NamespaceRole, 0)
				auths.Role = r
				return auths, nil
			}
			namespaceRole = append(namespaceRole, namespace)
			m[r.Namespace] = j
			j++
		} else {
			namespaceRole[m[r.Namespace]].Level += r.Level
		}
	}
	auths.Role = namespaceRole
	auths.IsOwner = false
	return auths, nil
}

func (s roleService) FindByAppId(appId, cluster, env, name string) (*models.Role, error) {
	roles, err := s.repository.FindByAppId(appId, cluster, env, name)
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

func (s roleService) DeleteByUserId(userId string) error {
	db := s.db.Begin()
	if err := s.repository.DeleteByUserId(db, userId); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call RoleRepository.DeleteByUserId failed")
	}
	db.Commit()
	return nil
}
