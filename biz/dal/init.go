package dal

import (
	"fmt"
	"github.com/U1traVeno/tiktok-shop/biz/dal/model"
	query "github.com/U1traVeno/tiktok-shop/biz/dal/query/user"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func Init() {

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
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai search_path=xd_schema",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	// 初始化 GORM 数据库对象
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// auto migrate
	err = db.AutoMigrate(model.NewAllModels().Models...)
	if err != nil {
		panic(err)
	}
	// 初始化 query
	query.SetDefault(db)

}
