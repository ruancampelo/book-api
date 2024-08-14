package models

type Author struct {
	Id    uint   `gorm:"column:id;primarykey;autoIncrement"`
	Name  string `gorm:"column:name;"`
	Books []Book
}
