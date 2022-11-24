package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Tag 创建Tag结构体，编译Gorm使用,JSON标签便于后续gin中进行格式转化
type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// GetTags 获取所有标签。
// 函数return中没有变量，函数后已经声明了变量，函数体中可以直接使用，并且使用指针&，因此无需指定返回值
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

// GetTagTotal db *gorm.DB为同一个models package下变量。可直接使用。
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

// AddTag 创建tag
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}

// ExistTagByName 通过name判断tag是否存在
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name=?", name).First(&tag)
	if tag.Id > 0 {
		return true
	}
	return false
}

// BeforeCreate gorm的回调方法。将回调方法定义为模型结构指针，在创建、更新、删除、查询时被调用，
// 如果在回掉时出现异常，gorm将停止操作进行回滚更改
// 创建：BeforeSave AfterSave BeforeCreate AfterCreate
// 更新: BeforeSave AfterSave BeforeUpdate AfterUpdate
// 删除 BeforeDelete AfterDelete
// 查询 AfterFind
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedOn", time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	err := scope.SetColumn("ModifiedOn", time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}

func ExistTagById(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	return tag.Id > 0
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}
