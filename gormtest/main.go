package main

import (
	"log"
	"time"
)
import "gorm.io/gorm"
import "gorm.io/driver/mysql"

type UserBasic struct {
	Uid         int        `json:"uid,omitempty"`
	Identity    string     `json:"identity"`
	Name        string     `json:"name"`
	Password    string     `json:"password"`
	Email       string     `json:"email"`
	CreatedTime *time.Time `json:"created_time"`
	//CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" `
}

func (table UserBasic) TableName() string {
	return "user_basic"
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/go-cloud-disk?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	db.Debug(). /*.Select("identity", "name", "password")*/ Create(&UserBasic{
		Identity: "ada",
		Name:     "ada",
		Password: "ada",
		//Email:    "ada",
	})
}
