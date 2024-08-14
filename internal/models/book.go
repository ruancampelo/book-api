package models

type Book struct {
	Id       uint   `gorm:"column:id;primarykey;autoIncrement"`
	Name     string `gorm:"column:name`
	AuthorID uint   `gorm:"column:author_id`
	Author   Author `gorm:"foreignKey:author_id"` // Relacionamento de chave estrangeira
}
