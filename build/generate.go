/*
Usage:
在 项目根目录 下创建.env文件，内容如下：

DB_USER=youruser
DB_PASSWORD=yourpassword
DB_HOST=113.xx.xx.xxx
DB_PORT=xxxxx
DB_NAME=yourdbname

之后编辑71行的OutPath，将 your_model_name 替换为你的模型名，
编辑 77 行的 g.ApplyBasic(model.User{})，将 User 替换为你的模型名
然后运行 go run ./build/generate.go

脚本会在 biz/dal/query/your_model_name/ 目录下生成对应的查询代码
*/
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

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
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai search_path=xd_schema",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	// 初始化 GORM 数据库对象
	gormdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	// 初始化 GORM Generator
	// 这里的OutPath改成自己模块的路径
	// 示例: OutPath: "./biz/dal/query/user/", 会生成到 biz/dal/query/user/ 目录下

	g := gen.NewGenerator(gen.Config{

		// Replace your_model_name with your model name
		OutPath: "./biz/dal/query/your_model_name/", // 输出路径

		Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// 使用数据库对象
	g.UseDB(gormdb)

	// Add your own model here
	//g.ApplyBasic(model.User{})

	g.Execute()
}
