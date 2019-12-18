package tenant

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gomall/pkg/web"
	"strings"
	"sync"
)

var (
	DBMaps  map[string]*TenantInfo
	dbMutex sync.Mutex
)

type TenantInfo struct {
	Name    string
	Domain  string
	ConnStr string
	Driver  string
	db      *gorm.DB
}

type MultiTenantHandlerFunc func(*web.Context)

func DBProvider() []*TenantInfo {
	return make([]*TenantInfo, 2)
}

func ResolverOrigin(ctx *gin.Context) string {
	id := ctx.GetHeader("Origin")
	if strings.TrimSpace(id) == "" {
		id = "localhost"
	}
	return id
}

func Setup() {

}
