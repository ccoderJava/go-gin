package models

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
