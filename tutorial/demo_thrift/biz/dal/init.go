package dal

import (
	"github.com/charleschile/TikTok-E-Commerce-Gomall/tutorial/demo_thrift/biz/dal/mysql"
	"github.com/charleschile/TikTok-E-Commerce-Gomall/tutorial/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
