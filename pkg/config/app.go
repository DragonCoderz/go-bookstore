package config

import (
	"github.com/jinzhu/gorm" //helps you talk to mySQL
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() { //Point of this function is to create a variable that allows us to interface with our database :) 
	d, err := gorm.Open("mysql", "akhil:Axlesharma@12@/simplerest?charset=utf8&parseTime=True&loc=Local") //Opens up connection w/ database
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}