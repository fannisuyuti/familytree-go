package model

import (
	"gorm.io/gorm"
	"suyuti.com/famtrees/common"
)

type Person struct {
	gorm.Model
	Name   string `gorm:"column:name;type:varchar(100);not null"`
	Father uint   `gorm:"column:father;type:bigint;null"`
	Mother uint   `gorm:"column:mother;type:bigint;null"`
	Birth  string `gorm:"column:birth;type:varchar(10);null"`
	Death  string `gorm:"column:death;type:varchar(10);null"`
	Gender string `gorm:"column:gender;type:varchar(10);null"`
}

type PersonResponse struct {
	Id          uint   `json:"id"`
	Parents     []uint `json:"parents"`
	Title       string `json:"title"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Image       string `json:"image"`

	//    { id: 3, parents: [1, 2], title: "Sibling 1", label: "Sibling 1", description: "Sibling 1", image: "/api/images/photos/s.png" },
}

type Option struct {
	ID   uint
	Name string
}

func GetOne(id uint) (Person, error) {
	db := common.GetDB()
	var person Person
	err := db.First(&person, id).Error
	return person, err
}

func GetAll() ([]Person, error) {
	db := common.GetDB()
	var persons []Person
	err := db.Find(&persons).Error
	return persons, err
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func (model *Person) Update(data interface{}) error {
	db := common.GetDB()
	err := db.Model(model).Updates(data).Error
	return err
}
