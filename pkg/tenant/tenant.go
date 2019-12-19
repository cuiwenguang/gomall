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
	DB      *gorm.DB
}

type MultiTenantHandlerFunc func(*web.Context)

func DBProvider() []*TenantInfo {
	// 配置文件或者redis中获取所有租户连接信息
	return make([]*TenantInfo, 2)
}

func ResolverOrigin(ctx *gin.Context) string {
	id := ctx.GetHeader("Origin")
	if strings.TrimSpace(id) == "" {
		id = "default"
	}
	return id
}

func Setup() {

}
