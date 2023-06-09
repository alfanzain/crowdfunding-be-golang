package main

import (
	"crowdfunding/handler"
	"crowdfunding/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/local_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	// userInput := user.RegisterInput{}
	// userInput.Name = "Test simpan dari service"
	// userInput.Email = "contoh@gmail.com"
	// userInput.Occupation = "thief"
	// userInput.Password = "password"

	// userService.Register(userInput)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("api/v1")
	api.POST("/users", userHandler.RegisterUser)

	router.Run()

	// fmt.Println("Connection to database is good")

	// var users []user.M_users
	// db.Find(&users)

	// for _, user := range users {
	// 	fmt.Println(user.Name)
	// 	fmt.Println(user.Email)
	// 	fmt.Println("=======================")
	// }

	// router := gin.Default()
	// router.GET("/h", handler)
	// router.Run()
}

// func handler(c *gin.Context) {
// 	// dsn := "root:@tcp(127.0.0.1:3306)/local_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
// 	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	// if err != nil {
// 	// 	log.Fatal(err.Error())
// 	// }

// 	// var users []user.User
// 	// db.Find(&users)

// 	// c.JSON(http.StatusOK, users)
// }
