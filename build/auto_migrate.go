/*
Description:
自动迁移数据库表结构

Usage:
1. 在 项目根目录 下创建.env文件，内容如下：
DB_USER=myuser
DB_PASSWORD=mypassword
DB_HOST=1xx.xx.xx.xxx
DB_PORT=xxxxx
DB_NAME=mydb

2. 运行 go run ./build/auto_migrate.go

脚本会自动迁移所有模型到数据库中

在实际运行之前, 建议可以把 52 行的search_path=xd_schema 替换为你的数据库 xd_test 的 schema 名称
xd_test 里面看起来没有问题之后, 再替换回来
*/
package main

import (
	"fmt"
	"github.com/U1traVeno/tiktok-shop/biz/dal/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	// 加载 .env 文件中的环境变量
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// 从环境变量读取数据库配置
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		panic("database configuration is not fully set in .env file")
	}

	// 构建 DSN
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai search_path=xd_test",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// 初始化所有模型
	allModels := model.NewAllModels()

	// 自动迁移
	err = db.AutoMigrate(allModels.Models...)
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}
}
