package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return err
	}
	return nil
}

func ParseQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return err
	}
	return nil
}

func ParseForm(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return err
	}
	return nil
}
func ParseUri(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindUri(obj); err != nil {
		return err
	}
	return nil
}
