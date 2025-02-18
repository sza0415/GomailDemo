package dal

import (
	"github.com/sza0415/GomailDemo/demo/demo_proto/biz/dal/mysql"
	"github.com/sza0415/GomailDemo/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
