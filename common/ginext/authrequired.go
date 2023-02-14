package ginext

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/jfcore/common/common"
	"net/http"
	"strconv"
)

// AuthRequiredMiddleware is required the request has to have x-user-id in header
// (it's usually set by API Gateway)
func AuthRequiredMiddleware(c *gin.Context) {
	headers := struct {
		UserID    string `header:"x-user-id" validate:"required,min=1"`
		UserMeta  string `header:"x-user-meta"`
		TenantID  uint64 `header:"x-tenant-id"`
		ClientKey string `header:"x-client-key"`
	}{}

	if c.ShouldBindHeader(&headers) != nil {
		_ = c.Error(NewError(http.StatusUnauthorized, "unauthorized"))
		c.Status(common.CODE_UNAUTHORIZED) // in case of we don't use this middleware with ErrorHandler
		c.Abort()
		return
	}

	if c.Param(common.GinParamObject) != "" && !common.SliceContains(c.Param(common.GinParamObject), common.GinObject) {
		_ = c.Error(NewError(http.StatusUnauthorized, "invalid route"))
		c.Status(common.CODE_BAD_REQUEST) // in case of we don't use this middleware with ErrorHandler
		c.Abort()
		return
	}

	c.Set(common.HeaderUserID, headers.UserID)
	c.Set(common.HeaderUserMeta, headers.UserMeta)
	c.Set(common.HeaderTenantID, headers.TenantID)
	c.Set(common.HeaderClientKey, headers.ClientKey)

	c.Next()
}

type GetStringer interface {
	GetString(key string) string
	Param(object string) string
}

// GetUserID returns the user ID embedded in Gin context
func GetUserID(c GetStringer) string {
	return c.GetString(common.HeaderUserID)
}

// GetOrder
func GetObject(c GetStringer) string {
	return c.Param(common.GinParamObject)
}

// GetClientKey returns the client Key embedded in Gin context
func GetClientKey(c GetStringer) string {
	return c.GetString(common.HeaderClientKey)
}

func Uint64HeaderValue(c *gin.Context, headerName string) uint64 {
	sValue := c.GetHeader(headerName)
	if sValue == "" {
		return 0
	}
	v, err := strconv.ParseUint(sValue, 10, 64)
	if err != nil {
		return 0
	}

	return v
}

func Uint64UserID(c *gin.Context) uint64 {
	return Uint64HeaderValue(c, common.HeaderUserID)
}

func Uint64TenantID(c *gin.Context) uint64 {
	return Uint64HeaderValue(c, common.HeaderTenantID)
}
