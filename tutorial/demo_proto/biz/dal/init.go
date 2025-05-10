package dal

import (
	"github.com/charleschile/TikTok-E-Commerce-Gomall/tutorial/demo_proto/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
