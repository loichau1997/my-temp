package ginext

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/jfcore/common/common"
)

func RequestIDMiddleware(c *gin.Context) {
	requestid := c.GetHeader(common.HeaderXRequestID)
	if requestid == "" {
		requestid = uuid.New().String()
		c.Request.Header.Set(common.HeaderXRequestID, requestid)
	}
	// set to context
	c.Set(common.HeaderXRequestID, requestid)

	// set to response header as well
	c.Header(common.HeaderXRequestID, requestid)

	c.Next()
}
