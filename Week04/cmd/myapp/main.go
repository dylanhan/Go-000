package main

import (
	"log"

	"../../internal/dao"
	"../../internal/service"
)

func main() {
	db, cleanup, err := dao.NewDB()
	defer cleanup()
	if err != nil {
		log.Printf("Init error:%v\n", err)
		return
	}
	d := dao.NewDao(db)
	serviceService := service.NewService(d)
	server := NewServer(serviceService)
	if err = srv.Run(); err != nil {
		log.Printf("Run error:%v\n", err)
		return
	}
}
