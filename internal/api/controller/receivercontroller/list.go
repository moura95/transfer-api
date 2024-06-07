package receivercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Receiver) list(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
