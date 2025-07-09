package router

import (
	"gsam/controller"
	repo "gsam/domain/repository"
	userSvc "gsam/domain/service/user"
	"gsam/handler"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Router(r *gin.Engine) {
	db, err := sqlx.Open("sqlite", "./videos.db")
	if err != nil {
		log.Fatal(err)
	}
	videoRepo := repo.NewVideoRepo(db)

	userService := userSvc.NewUserService() // 初始化 service
	userController := controller.NewUserController(userService, videoRepo)
	UserHandler := handler.NewUserHandler(userController)
	testController := controller.NewTestController()
	TestHandler := handler.NewTestHandler(testController)

	r.GET("/hello/", HelloWorld)
	userGroup := r.Group("/user")
	{
		userGroup.GET("/db/init", UserHandler.InitSqlLite)
		userGroup.GET("/db/video", UserHandler.GetVideo)

		userGroup.GET("/:id", UserHandler.GetUser)
		userGroup.POST("/", UserHandler.UserPost)
		userGroup.POST("/test/map", UserHandler.TestMap)
		userGroup.GET("/test/routine", UserHandler.TestRoutine)
	}
	testGroup := r.Group("/test")
	{
		testGroup.GET("/22", TestHandler.TestCase22)
	}

}

var balance = 1000

func HelloWorld(context *gin.Context) {
	var msg = "您的帳戶內有:" + strconv.Itoa(balance) + "元"
	context.JSON(http.StatusOK, gin.H{
		"amount":  balance,
		"status":  "ok",
		"message": msg,
	})

}
