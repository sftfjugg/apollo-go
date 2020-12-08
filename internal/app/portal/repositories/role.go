package repositories

import (
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.didapinche.com/time"
)

type RoleRepository interface {
	Creates(db *gorm.DB, role []*models.Role) error
	//Update(db *gorm.DB, role *models.Role) error
	Delete(db *gorm.DB, appId string) error
	Find(appId, userId string) ([]*models.Role, error) //查找用户在对应项目下权限
	FindByAppId(appId string) ([]*models.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleReposotory(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r roleRepository) Creates(db *gorm.DB, role []*models.Role) error {
	s := "insert into Role(`AppId`,`UserID`,`UserName`,`Level`,`DataChange_CreatedBy`,`DataChange_CreatedTime`) values"
	var buffer bytes.Buffer
	if _, err := buffer.WriteString(s); err != nil {
		return errors.Wrap(err, "creates releaseMessage error")
	}
	for i, r := range role {
		if i == len(role)-1 {
			buffer.WriteString(fmt.Sprintf("('%s','%s','%s','%v','%s','%s');", r.AppId, r.UserID, r.UserName, r.Level, r.DataChange_CreatedBy, time.Now()))
		} else {
			buffer.WriteString(fmt.Sprintf("('%s','%s','%s','%v','%s','%s'),", r.AppId, r.UserID, r.UserName, r.Level, r.DataChange_CreatedBy, time.Now()))
		}
	}
	if err := db.Exec(buffer.String()).Error; err != nil {
		return errors.Wrap(err, "creates Role error")
	}
	return nil
}

func (r roleRepository) Delete(db *gorm.DB, appId string) error {
	if err := db.Table(models.RoleTableName).Where("AppId= ? and (Level=1 or Level=2)", appId).Update("IsDeleted", 1).Error; err != nil {
		return errors.Wrap(err, "roleRepository.Delete failed")
	}
	return nil
}

func (r roleRepository) Find(appId, userId string) ([]*models.Role, error) {
	role := make([]*models.Role, 0)
	if err := r.db.Table(models.RoleTableName).Find(&role, "appId=? and userId=?", appId, userId).Error; err != nil {
		return nil, errors.Wrap(err, "roleRepository.FindByAppId failed")
	}
	return role, nil
}

func (r roleRepository) FindByAppId(appId string) ([]*models.Role, error) {
	role := make([]*models.Role, 0)
	if err := r.db.Table(models.RoleTableName).Find(&role, "appId=?", appId).Error; err != nil {
		return nil, errors.Wrap(err, "roleRepository.FindByAppId failed")
	}
	return role, nil
}
