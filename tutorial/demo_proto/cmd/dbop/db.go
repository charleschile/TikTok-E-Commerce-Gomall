package main

import (
	"github.com/charleschile/TikTok-E-Commerce-Gomall/tutorial/demo_proto/biz/dal"
	"github.com/charleschile/TikTok-E-Commerce-Gomall/tutorial/demo_proto/biz/dal/mysql"
	"github.com/charleschile/TikTok-E-Commerce-Gomall/tutorial/demo_proto/biz/model"

	"github.com/joho/godotenv"

	"github.com/cloudwego/kitex/pkg/klog"
	// 注意教程中是kitex下的consul包而不是hertz下的
	// "github.com/hertz-contrib/registry/consul"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		klog.Fatal("Error loading .env file")
		panic(err)
	}
	dal.Init()

	// CRUD
	// 创建对象
	// mysql.DB.Create(&model.User{Email: "demo@example.com", Password: "12345"})
	// 修改数据
	// mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").Update("password", "2222222222")

	// // 读取数据：first读取一行，find是读取多行记录
	// var row model.User
	// mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").First(&row)
	// fmt.Printf("row: %+v\n", row)

	// // 删除数据
	// mysql.DB.Where("email = ?", "demo@example.com").Delete(&model.User{})

	// 强制删除
	mysql.DB.Unscoped().Where("email = ?", "demo@example.com").Delete(&model.User{})

}
