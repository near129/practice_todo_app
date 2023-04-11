package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/near129/practice_todo_app/Config"
	"github.com/near129/practice_todo_app/Models"
	"github.com/near129/practice_todo_app/Routes"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("status: ", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Todo{})

	r := Routes.SetupRouter()
	r.Run()
}
