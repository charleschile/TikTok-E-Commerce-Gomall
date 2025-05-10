package mysql

import (
	"fmt"
	"os"

	"github.com/charleschile/TikTok-E-Commerce-Gomall/tutorial/demo_proto/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	// 根据格式化字符串和提供的参数创建新的字符串
	// 由于conf中的dsn是“"%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local"”
	// sprintf会填充

	// Data Source Name
	// 在终端使用export MYSQL_USER=root，在项目启动的时候就能读取到
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	)
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	type Version struct {
		Version string
	}
	var v Version
	// “as version”将查询结果设置成了字段version
	// gorm会自动将字段名与结构体中的字段名匹配，并且不区分大小写
	// 所以SQL查询结果的version会被影射到v中的Version字段
	err = DB.Raw("select version() as version").Scan(&v).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(v.Version)
}
