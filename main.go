package main

import (
	"c2n/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":9999") // 启动服务
}
