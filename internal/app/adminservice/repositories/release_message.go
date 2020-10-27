package repositories

import (
	"apollo-adminserivce/internal/pkg/models"
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.didapinche.com/time"
)

type ReleaseMessageRepository interface {
	Create(db *gorm.DB, releaseMessage *models.ReleaseMessage) error
	Creates(db *gorm.DB, releaseMessages []*models.ReleaseMessage) error
	DeleteByMessage(db *gorm.DB, message string) error
	DeleteByMessages(db *gorm.DB, messages []string) error
}

type releaseMessageRepository struct {
}

func NewReleaseMessageRepository() ReleaseMessageRepository {
	return &releaseMessageRepository{}
}

func (r releaseMessageRepository) Create(db *gorm.DB, releaseMessage *models.ReleaseMessage) error {
	releaseMessage.DataChange_LastTime = time.Now()
	if err := db.Create(releaseMessage).Error; err != nil {
		return errors.Wrap(err, "create releaseMessage error")
	}
	return nil
}

func (r releaseMessageRepository) Creates(db *gorm.DB, releaseMessages []*models.ReleaseMessage) error {
	s := "insert into ReleaseMessage(Message,DataChange_LastTime) values"
	var buffer bytes.Buffer
	if _, err := buffer.WriteString(s); err != nil {
		return errors.Wrap(err, "creates releaseMessage error")
	}
	for i, r := range releaseMessages {
		if i == len(releaseMessages)-1 {
			r.DataChange_LastTime = time.Now()
			buffer.WriteString(fmt.Sprintf("('%s','%s');", r.Message, r.DataChange_LastTime))
		} else {
			buffer.WriteString(fmt.Sprintf("('%s','%s'),", r.Message, r.DataChange_LastTime))
		}
	}
	if err := db.Exec(buffer.String()).Error; err != nil {
		return errors.Wrap(err, "creates releaseMessage error")
	}
	return nil
}

func (r releaseMessageRepository) DeleteByMessage(db *gorm.DB, message string) error {
	if err := db.Table(models.ReleaseMessageTableName).Delete(&models.ReleaseMessage{}, "Message= ?", message).Error; err != nil {
		return errors.Wrap(err, "delete previous releaseMessage error")
	}
	return nil
}

func (r releaseMessageRepository) DeleteByMessages(db *gorm.DB, messages []string) error {
	if err := db.Delete(&models.ReleaseMessage{}, "Message IN  (?)", messages).Error; err != nil {
		return errors.Wrap(err, "delete previous releaseMessage error")
	}
	return nil
}
