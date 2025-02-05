package controller

import (
	"codebase-golang/internal/app/utils"
	"codebase-golang/internal/app/utils/database/mongodb"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	data := mongodb.GetAll(c)
	if data == nil {
		utils.Error(c, http.StatusInternalServerError)
		return
	}
	utils.Response(c, data, "Get all data has been successfully")
}

func Get(c *gin.Context) {
	data := mongodb.Get(c)
	if data == nil {
		utils.Error(c, http.StatusInternalServerError)
		return
	}
	utils.Response(c, data, "Get a data has been successfully")
}

func Add(c *gin.Context) {
	data := mongodb.Add(c)
	if data == nil {
		utils.Error(c, http.StatusInternalServerError)
		return
	}
	utils.Response(c, data, "Add data has been successfully")
}

func Update(c *gin.Context) {
	data := mongodb.Update(c)
	if data == nil {
		utils.Error(c, http.StatusInternalServerError)
		return
	}
	utils.Response(c, data, "Update data has been successfully")
}

func Delete(c *gin.Context) {
	data := mongodb.Delete(c)
	if data == 0 {
		utils.Error(c, http.StatusInternalServerError)
		return
	}
	utils.Response(c, data, "Delete all data has been successfully")
}
