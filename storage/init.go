package storage

import "strings"

func Setup() {
	// 数据库
	initDB()
	// redis
	initRedis()
}

func GetDomain(id string) string {
	if strings.TrimSpace(id) == "" ||
		strings.Contains(id, "localhost") ||
		strings.Contains(id, "127.0.0.1") {
		id = "default"
	}
	return id
}
