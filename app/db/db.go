package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	//mysql import
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/revel/revel"
)

/*
// Base Model's definition
type Model struct {
  ID        uint `gorm:"primary_key"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt *time.Time
}
*/
// DB connection
var DB *gorm.DB

const sECRET = "donotputthesecretinyourcode"

//Survey is the base DBN object upon which we build our surveys
type Survey struct {
	gorm.Model
	Name      string
	Questions []Question `gorm:"many2many:survey_to_question;"`
}

//Question is the table that will hold our survey Questions
type Question struct {
	gorm.Model
	Question string
	Answers  []Answers
	Surveys  []Survey `gorm:"many2many:survey_to_question;"`
}

//QuestionType of question, I.E. multiple choice, single response, video, audio, etc
type QuestionType struct {
	gorm.Model
	name string
}
type Answers struct {
	gorm.Model
	answer     string
	QuestionID uint
}
type Response struct {
}

func InitDB() {
	fmt.Println("eyyyy")
	dbCon, err := gorm.Open("mysql", "root:root@/survey?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		revel.ERROR.Printf("%s", err)
	}
	DB = dbCon

	DB.AutoMigrate(&User{})
	DB.DB().SetMaxOpenConns(100)

}
