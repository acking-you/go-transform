package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	kGoodCode = 0
	kBadCode  = 1
)

type StatusResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type ListResponse struct {
	StatusCode int      `json:"status_code"`
	ListMsg    []string `json:"list_msg"`
}

func GoodResponse(c *gin.Context, msg string) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, &StatusResponse{
		StatusCode: kGoodCode,
		StatusMsg:  msg,
	})
}

func BadResponse(c *gin.Context, msg string) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, &StatusResponse{
		StatusCode: kBadCode,
		StatusMsg:  msg,
	})
}

func GoodListResponse(c *gin.Context, listMsg []string) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, &ListResponse{
		StatusCode: kGoodCode,
		ListMsg:    listMsg,
	})
}

func BadListResponse(c *gin.Context, msg string) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, &ListResponse{
		StatusCode: kBadCode,
		ListMsg:    []string{msg},
	})
}
