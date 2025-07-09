package handler

import (
	"gsam/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
	testController controller.TestController
}

func NewTestHandler(testController controller.TestController) *TestHandler {
	return &TestHandler{
		testController: testController,
	}
}

func (ctrl *TestHandler) TestCase22(c *gin.Context) {
	if err := ctrl.testController.TestCase22(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}
