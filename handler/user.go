package handler

import (
	"fmt"
	"gsam/controller"
	"gsam/domain/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserController controller.UserController
}

func NewUserHandler(userController controller.UserController) *UserHandler {
	return &UserHandler{
		UserController: userController,
	}
}

func (ctrl *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("#### id:", id)
	msg, err := ctrl.UserController.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, msg)
}

func (ctrl *UserHandler) UserPost(c *gin.Context) {
	var input entity.UserPost
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func (ctrl *UserHandler) TestMap(c *gin.Context) {
	var input entity.TestMap
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.UserController.TestMap(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func (ctrl *UserHandler) TestRoutine(c *gin.Context) {
	if err := ctrl.UserController.TestRoutine(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entity.TestMap{})
}

func (ctrl *UserHandler) InitSqlLite(c *gin.Context) {
	ctrl.UserController.InitSqlLite()

	c.Status(http.StatusNoContent)
}

func (ctrl *UserHandler) GetVideo(c *gin.Context) {
	result, err := ctrl.UserController.GetVideo()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
