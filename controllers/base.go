package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	successReturn = gin.H{"code": 200, "data": "ok", "msg": "ok"}
	err404        = gin.H{"code": 404, "data": "", "msg": "object not found"}
	err403        = gin.H{"code": 403, "data": "", "msg": "no permission"}
	errToken      = gin.H{"code": 401, "data": "", "msg": "wrong token"}
	badReq        = gin.H{"code": 400, "data": "", "msg": "bad request"}
	unauthorized  = gin.H{"code": 401, "data": "", "msg": "wrong name or password"}
	nameExist     = gin.H{"code": 4001, "data": "", "msg": "name has existed"}
)

func SuccessJson(c *gin.Context) {
	c.JSON(http.StatusOK, successReturn)
}

func GenJson(data interface{})(gin.H){
	return gin.H{"code": 200, "data": data, "msg": "ok"}
}
