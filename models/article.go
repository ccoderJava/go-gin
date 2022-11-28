package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Article TagID 声明为gorm:index 表示该字段为索引
// 下列中Tag字段为结构体，表示当前TagID字段与之模型关联 能够达到Article Tag相关联
// time.Now().Unix() 表示返回当前时间
type Article struct {
	Model
	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func ExistArticleById(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)

	return article.Id > 0
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preloads("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

// GetArticle Article有一个TagId 成员变量 理解为外键。gorm会通过类名+ID寻找对应两个类之间关联关系
// Article有一个结构体成员变量，即为上面嵌套在Article中的Tag结构体。使得我们可以通过Related进行关联查询
func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}

// EditArticle 编辑文章
func EditArticle(id int, maps interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(maps)
	return true
}

// AddArticle  添加文章
func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})
	return true
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedOn", time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	err := scope.SetColumn("ModifiedOn", time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}
