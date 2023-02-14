package ginext

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"gitlab.com/jfcore/common/logger"
)

func TestMain(m *testing.M) {
	logger.Init("ginext.test")
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
