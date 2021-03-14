package main

import (
	"errors"
	"fmt"
	"github.com/TheEgid/articles-site-go/domain/entity/food_entity"
	"github.com/TheEgid/articles-site-go/handlers/food_handler"
	"github.com/TheEgid/articles-site-go/infrastructure/dtb"
	"github.com/TheEgid/articles-site-go/middleware"
	"github.com/TheEgid/articles-site-go/routes"
	"github.com/TheEgid/articles-site-go/utils"
	"github.com/gin-gonic/gin"
	"runtime"
)

func GetFoodObject(repos *dtb.Repositories, id uint64) *food_entity.Food {
	FoodObject, err := repos.Food.GetFood(id)
	if err != nil {
		fmt.Println(errors.Unwrap(err))
	}
	return FoodObject
}

func GetFoodObjects(repos *dtb.Repositories) *[]food_entity.Food {
	FoodObjects, err := repos.Food.GetAllFood()
	if err != nil {
		fmt.Println(errors.Unwrap(err))
	}
	return &FoodObjects
}

func main() {
	var DbHost string
	if runtime.GOOS != "windows" {
		DbHost = *utils.GoDotEnvVariable("DB_HOST")
	} else {
		DbHost = "localhost"
	}

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.Use(gin.Logger())

	DbName := utils.GoDotEnvVariable("DB_NAME")
	DbPort := utils.GoDotEnvVariable("DB_PORT")
	dbPassword := utils.GoDotEnvVariable("DB_PASSWORD")
	dbUser := utils.GoDotEnvVariable("DB_USER")

	repos, _ := dtb.NewRepositories(*dbUser, *dbPassword, *DbPort, DbHost, *DbName)

	defer repos.Close()

	f := GetFoodObject(repos, 1)
	if f == nil {
		repos.Automigrate()
	}

	//backupDb, _ := utils.Setup()
	//if err := backupDb.CreateDump(); err != nil { fmt.Println(err) }

	foods := food_handler.NewFood(repos.Food)
	//fmt.Printf("%#v\n", GetFoodObject(repos, 1))
	//fmt.Printf("%#v\n", GetFoodObjects(repos))

	//router.Static("/assets", "./assets")
	router.StaticFile("favicon.ico", "assets/favicon.ico")
	router.LoadHTMLGlob("templates/*")
	router.Use(middleware.CORSMiddleware()) //For CORS

	routes.InitializeRoutes(router)
	router.GET("/food", foods.GetAllFood)
	_ = router.Run(":80")
}
