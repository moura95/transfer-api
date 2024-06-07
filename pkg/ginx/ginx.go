package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/moura95/transferapi/pkg/errors"
)

// Parse body json data to struct
func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return errors.BadRequest("", "Failed to parse json: %s", err.Error())
	}
	return nil
}

// Parse query parameter to struct
func ParseQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return errors.BadRequest("", "Failed to parse query: %s", err.Error())
	}
	return nil
}

// Parse body form data to struct
func ParseForm(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return errors.BadRequest("", "Failed to parse form: %s", err.Error())
	}
	return nil
}
func ParseUri(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindUri(obj); err != nil {
		return errors.BadRequest("", "Failed to parse uri: %s", err.Error())
	}
	return nil
}
