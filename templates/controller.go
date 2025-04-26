package templates

const ControllerTmplStr = `package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"{{ .ProjectName }}/services"
)


// GetList 查询{{ .Comment }}
func Get{{ .StructName }}List(c *gin.Context) {
	var req services.GetList{{ .StructName }}Params
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var {{ .StructName }}Service = services.New{{ .StructName }}Service()
	list, err := {{ .VarName }}Service.GetList(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, list)
}

// Create 创建{{ .Comment }}
func Create{{ .StructName }}(c *gin.Context) {
	var req services.Create{{ .StructName }}Params
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var {{ .StructName }}Service = services.New{{ .StructName }}Service()
	if err := {{ .VarName }}Service.Create(c, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "创建成功"})
}

// Update 更新{{ .Comment }}
func Update{{ .StructName }}(c *gin.Context) {
	var req services.Update{{ .StructName }}Params
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var {{ .StructName }}Service = services.New{{ .StructName }}Service()
	if err := {{ .VarName }}Service.Update(c, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

// Delete 删除{{ .Comment }}
func Delete{{ .StructName }}(c *gin.Context) {
	var req services.Delete{{ .StructName }}Params
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var {{ .StructName }}Service = services.New{{ .StructName }}Service()
	if err := {{ .VarName }}Service.Delete(c, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
`
