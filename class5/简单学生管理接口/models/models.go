package models

//package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)
type Student struct{
	Sid string `gorm:"primary_key"`
	Name string
	Age int64
	Sex string
}

func RegisterDB(){
	db, err :=gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal("failed to connect database")
	}
	db.AutoMigrate(&Student{})
	defer db.Close()
}
func QueryRow(sid string) *Student{
	db, err :=gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("failed to connect database")
	}
	defer db.Close()
	s := &Student{}
	s.Sid = sid
	if db.First(s).RecordNotFound(){
		//fmt.Println("not ")
		return nil
	}
	return s
}
func QueryAll() []Student{
	db, err :=gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("failed to connect database")
	}
	defer db.Close()
	s := make([]Student, 0)
	db.Find(&s)
	return s
}
func Insert(stu []Student) []Student {
	db, err :=gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("failed to connect database")
	}
	defer db.Close()
	sIn := make([]Student,0)
	for _,s := range stu{
		if QueryRow(s.Sid) != nil {
			sIn = append(sIn,s)
			fmt.Println("学号：" + s.Sid + " already exist")
		}else {
			if err := db.Create(s).Error; err != nil {
				log.Fatal("插入失败", err)
			}
		}


	}
	return sIn
}
func DeleteFew(stu []string) error{
	db, err :=gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("failed to connect database")
	}
	defer db.Close()
	for _,s := range stu{
		if err := db.Where("Sid = ?", s).Delete(Student{}).Error; err != nil{
			log.Fatal(err)
		}
	}
	return nil
}
func DeleteAll() error{
	db, err :=gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("failed to connect database")
	}
	defer db.Close()
	db.Delete(Student{})
	return nil
}
func Modify(stu []Student) error{
	db, err :=gorm.Open("mysql", "root:666234@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("failed to connect database")
	}
	defer db.Close()
	for _,s := range stu{
		db.Save(&s)
	}
	return nil
}
/*
//func main(){
//	//RegisterDB()
//	s := make([]Student, 0)
//	for i := 0; i < 10; i++{
//	//
//		s = append(s, Student{strconv.Itoa(i),strconv.Itoa(i),int64(i),strconv.Itoa(i)})
//	}
//	//s := make([]string,0)
//	Insert(s)
//	//s = append(s, "1")
//	//DeleteFew(s)
//	//DeleteAll()
//	//s = append(s, Student{strconv.Itoa(1),strconv.Itoa(4),int64(3),strconv.Itoa(4)})
//	//Modify(s)
//	fmt.Println(QueryAll())
//}
*/