package services

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	models2 "go.didapinche.com/foundation/apollo-plus/internal/app/admin/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/admin/repositories"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/models"
	"go.didapinche.com/time"
	"sort"
	"strconv"
	"strings"
)

type ItemService interface {
	Create(item *models.Item) error
	CreateByText(items *models2.ItemText) error
	Creates(item []*models.Item) error
	Update(item *models.Item) error
	DeleteById(id, operator string) error
	DeleteByNamespaceId(namespaceId string) error
	FindItemByAppIdAndKey(appId, cluster, key, format, comment string) ([]*models2.AppNamespace, error)
	FindItemByNamespaceId(namespaceID, comment string) ([]*models.Item, error)
	FindItemByNamespaceIdOnRelease(namespaceID string) ([]*models.Item, error)
	FindItemByKeyForPage(cluster, key, format string, pageSize, pageNum int) (*models2.ItemPage, error)
	FindAppItemByKeyForPage(cluster, key, format string, pageSize, pageNum int) (*models2.AppNamespacePage, error)
	FindItemByNamespaceIdAndKey(namespaceId, key string) ([]*models.Item, error)
	FindOneItemByNamespaceIdAndKey(namespaceId uint64, key string) (*models.Item, error)
	FindAllComment(appId string) ([]string, error)
}

type itemService struct {
	repository repositories.ItemRepisitory
	db         *gorm.DB
}

func NewItemService(
	repository repositories.ItemRepisitory,
	db *gorm.DB,
) ItemService {
	return &itemService{
		db:         db,
		repository: repository,
	}
}

//通过文本编辑修改值，先获得该配置下所有items，然后通过map对比改动，原来不存在的值则新增，原来存在，现在存在的值则删除，同时存在则修改,因为删除操作为逻辑删除，所以可以使用一个sql语句进行
func (s itemService) CreateByText(itemText *models2.ItemText) error {
	items, err := s.repository.FindItemByNamespaceId(strconv.FormatUint(itemText.NamespaceId, 10), "") //目前所有主版本配置
	if err != nil {
		return errors.Wrap(err, "call itemService.FindItemByNamespaceId() error")
	}
	m := make(map[string]int)
	for i := range items {
		m[items[i].Key] = i
	}
	//解析text
	texts := strings.Split(itemText.Text, "\n")
	if len(texts) == 0 {
		return errors.New("格式错误或者内容为空")
	}
	itemsSave := make([]*models.Item, 0) //需要新增，修改，删除的
	mrepeat := make(map[string]int)
	for _, t := range texts { //第一次循环将新增修改的添加itemsSave并剔除于m，使m中只留下需要删除的
		if strings.Count(t, "=") >= 1 {

			k := strings.Index(t, "=")
			key := strings.Trim(t[0:k], " ")
			value := strings.Trim(t[k+1:len(t)], " ")
			if _, ok := mrepeat[key]; ok {
				return errors.New("存在重复的key")
			}
			mrepeat[key] = 0
			if key != "" {
				if _, ok := m[key]; ok {
					i := m[key]
					if value != items[i].Value {
						items[i].Value = strings.ReplaceAll(value, "\\", "\\\\")
						items[i].Status = 2
						items[i].DataChange_LastModifiedBy = itemText.Operator
						itemsSave = append(itemsSave, items[i])
						delete(m, key)
					} else {
						items[i].Value = strings.ReplaceAll(value, "\\", "\\\\")
						itemsSave = append(itemsSave, items[i])
						delete(m, key)
					}
				} else {
					item := new(models.Item)
					item.Key = strings.ReplaceAll(key, "\\", "\\\\")
					item.Value = strings.ReplaceAll(value, "\\", "\\\\")
					item.Status = 0
					item.NamespaceId = itemText.NamespaceId
					item.DataChange_CreatedBy = itemText.Operator
					item.DataChange_LastModifiedBy = itemText.Operator
					item.DataChange_CreatedTime = time.Now()
					itemsSave = append(itemsSave, item)
				}
			}
		}
	}
	//整理所有需要删除的
	for _, i := range m {
		items[i].Status = 3
		items[i].DataChange_LastModifiedBy = itemText.Operator
		itemsSave = append(itemsSave, items[i])
	}
	db := s.db.Begin()
	if err := s.repository.Saves(db, itemsSave); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.Saves() error")
	}
	db.Commit()
	return nil
}

func (s itemService) Create(item *models.Item) error {
	item2, err := s.FindOneItemByNamespaceIdAndKey(item.NamespaceId, item.Key)
	if err != nil {
		return errors.Wrap(err, "call itemService.FindItemByNamespaceIdAndKey() error")
	}
	if item2.Key != "" {
		return errors.New("item already exists")
	}
	db := s.db.Begin()
	if err := s.repository.Create(db, item); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.Create() error")
	}
	db.Commit()
	return nil
}

func (s itemService) Creates(items []*models.Item) error {
	db := s.db.Begin()
	if err := s.repository.Creates(db, items); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.Creates() error")
	}
	db.Commit()
	return nil
}

func (s itemService) DeleteById(id, operator string) error {
	db := s.db.Begin()
	if err := s.repository.DeleteById(db, id, operator); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.DeleteByNamespaceIdAndKey() error")
	}
	db.Commit()
	return nil
}

func (s itemService) DeleteByNamespaceId(namespaceId string) error {
	db := s.db.Begin()
	if err := s.repository.DeleteByNamespaceId(db, namespaceId); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.DeleteByNamespaceIdAndKey() error")
	}
	db.Commit()
	return nil
}

func (s itemService) Update(item *models.Item) error {
	item2, err := s.FindOneItemByNamespaceIdAndKey(item.NamespaceId, item.Key)
	if err != nil {
		return errors.Wrap(err, "call itemService.FindItemByNamespaceIdAndKey() error")
	}
	if item2.Key != "" && item2.Key != item.Key {
		return errors.New("item already exists")
	}
	db := s.db.Begin()
	if err := s.repository.Update(db, item); err != nil {
		db.Rollback()
		return errors.Wrap(err, "call ItemRepository.Update() error")
	}
	db.Commit()
	return nil
}

func (s itemService) FindItemByNamespaceId(namespaceID, comment string) ([]*models.Item, error) {
	items, err := s.repository.FindItemByNamespaceId(namespaceID, comment)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindItemByNamespaceId() error")
	}
	return items, nil
}

func (s itemService) FindItemByNamespaceIdOnRelease(namespaceID string) ([]*models.Item, error) {
	items, err := s.repository.FindItemByNamespaceIdOnRelease(namespaceID)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindItemByNamespaceIdOnRelease() error")
	}
	return items, nil
}

func (s itemService) FindItemByKeyForPage(cluster, key, format string, pageSize, pageNum int) (*models2.ItemPage, error) {
	items, err := s.repository.FindItemByKeyForPage(cluster, key, format, pageSize, pageNum)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindItemByKeyForPage() error")
	}
	count, err := s.repository.FindItemCountByKey(key)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindItemCountByKey() error")
	}
	itemPage := new(models2.ItemPage)
	itemPage.Items = items
	itemPage.Total = count
	return itemPage, nil
}

func (s itemService) FindAppItemByKeyForPage(cluster, key, format string, pageSize, pageNum int) (*models2.AppNamespacePage, error) {
	items, err := s.repository.FindItemByKeyForPage(cluster, key, format, pageSize, pageNum)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindItemByKeyForPage() error")
	}
	appNamespaces := s.ItemChangeAppNamespace(items)
	count, err := s.repository.FindItemCountByKey(key)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindItemCountByKey() error")
	}
	appNamespacePage := new(models2.AppNamespacePage)
	appNamespacePage.Total = count
	appNamespacePage.AppNamespaces = appNamespaces
	return appNamespacePage, nil
}

func (s itemService) FindItemByAppIdAndKey(appId, cluster, key, format, comment string) ([]*models2.AppNamespace, error) {
	items, err := s.repository.FindItemByAppIdAndKey(appId, cluster, key, format, comment)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindItemByAppIdAndKey() error")
	}
	appNamespaces := s.ItemChangeAppNamespace(items)
	return appNamespaces, nil
}

func (s itemService) FindItemByNamespaceIdAndKey(namespaceId, key string) ([]*models.Item, error) {
	items, err := s.repository.FindItemByNamespaceIdAndKey(namespaceId, key)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindItemByNamespaceId() error")
	}
	return items, nil
}

func (s itemService) FindOneItemByNamespaceIdAndKey(namespaceId uint64, key string) (*models.Item, error) {
	item, err := s.repository.FindOneItemByNamespaceIdAndKey(namespaceId, key)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindOneItemByNamespaceIdAndKey() error")
	}
	return item, nil
}

func (s itemService) FindAllComment(appId string) ([]string, error) {
	items, err := s.repository.FindAllComment(appId)
	if err != nil {
		return nil, errors.Wrap(err, "call ItemRepository.FindAllComment() error")
	}
	comments := make([]string, 0)
	for _, i := range items {
		if i.Comment != "" {
			comments = append(comments, i.Comment)
		}
	}
	return comments, nil
}

//作用是将Item格式转化为前端展示格式
func (s itemService) ItemChangeAppNamespace(items []*models2.Item) []*models2.AppNamespace {
	names := make(map[string][]int)
	for i, n := range items {
		name := n.AppId + n.ClusterName + n.Name
		if _, ok := names[name]; !ok {
			names[name] = make([]int, 0)
		}
		names[name] = append(names[name], i)
	}
	appNamespaces := make([]*models2.AppNamespace, 0)
	for _, v := range names {
		appNamespace := new(models2.AppNamespace)
		lanes := make(map[string][]*models2.Item)
		for _, j := range v {
			if items[j].LaneName == "default" || items[j].LaneName == "主版本" {
				appNamespace.Format = items[j].Format
				appNamespace.Name = items[j].Name
				appNamespace.AppId = items[j].AppId
				appNamespace.ClusterName = items[j].ClusterName
				appNamespace.IsPublic = items[j].IsPublic
				appNamespace.DeptName = items[j].DeptName
				appNamespace.IsDisplay = items[j].IsDisplay
				appNamespace.Comment = items[j].NamespaceComment
			}
			lanes[items[j].LaneName] = append(lanes[items[j].LaneName], items[j])

		}
		for _, c := range lanes {
			namespace := new(models2.Namespace)
			its := make([]*models.Item, 0)
			for _, s := range c {
				itemModel := new(models.Item)
				itemModel.Id = s.Id
				itemModel.Value = s.Value
				itemModel.Key = s.Key
				itemModel.NamespaceId = s.NamespaceId
				itemModel.DataChange_CreatedTime = s.DataChange_CreatedTime
				itemModel.DataChange_LastTime = s.DataChange_LastTime
				itemModel.DataChange_LastModifiedBy = s.DataChange_LastModifiedBy
				itemModel.DataChange_CreatedBy = s.DataChange_CreatedBy
				itemModel.Describe = s.Describe
				itemModel.Comment = s.Comment
				itemModel.ReleaseValue = s.ReleaseValue
				itemModel.Status = s.Status
				its = append(its, itemModel)
				namespace.LaneName = s.LaneName
				namespace.Id = s.NamespaceId
			}
			namespace.Items = its
			appNamespace.Namespaces = append(appNamespace.Namespaces, namespace)
		}
		appNamespaces = append(appNamespaces, appNamespace)
	}
	sort.Sort(models2.AppNamespaceSlice(appNamespaces))
	return appNamespaces
}
