package main

import (
	"fmt"
	router "gsam/router"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var balance = 1000

func main() {
	r := gin.Default()

	router.Router(r)

	r.Run(":8080")
}

func HelloWorld(context *gin.Context) {
	var msg = "您的帳戶內有:" + strconv.Itoa(balance) + "元"
	context.JSON(http.StatusOK, gin.H{
		"amount":  balance,
		"status":  "ok",
		"message": msg,
	})

}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	sumnum := append(nums1, nums2...)

	fmt.Println("####", sumnum)
	return 0
}
