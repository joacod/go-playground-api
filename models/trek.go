package models

import (
	"playground/database"

	_ "github.com/go-sql-driver/mysql" // MySQL driver, needed to work properly
)

// CrewMember crew member
type CrewMember struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Starship string `json:"starship"`
}

// TableName name of the table
func (b *CrewMember) TableName() string {
	return "crew"
}

// GetAllMembers get all
func GetAllMembers(member *[]CrewMember) (err error) {
	if err = database.DB.Find(member).Error; err != nil {
		return err
	}
	return nil
}

// CreateMember create
func CreateMember(member *CrewMember) (err error) {
	if err = database.DB.Create(member).Error; err != nil {
		return err
	}
	return nil
}

// GetMember get
func GetMember(member *CrewMember, id string) (err error) {
	if err := database.DB.Where("id = ?", id).First(member).Error; err != nil {
		return err
	}
	return nil
}

// UpdateMember update
func UpdateMember(member *CrewMember, id string) (err error) {
	database.DB.Save(member)
	return nil
}

// DeleteMember delete
func DeleteMember(member *CrewMember, id string) (err error) {
	database.DB.Where("id = ?", id).Delete(member)
	return nil
}
