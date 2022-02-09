package models

import (

	//"gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=mikes_db password=mikes_db dbname=mikes_db port=5432 sslmode=disable TimeZone=US/Eastern"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("could not open db")
	}
	db.AutoMigrate(&Meetup{})
	DB = db

	// insert some dummy data as we go
	/*
		fmt.Println("NOTE - generating an extra dummy record")
		db.Create(&Product{Code: "D42", Price: 100})
	*/
}
