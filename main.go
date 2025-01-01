package main

import (
	"basic-project/internal/repository"
	"basic-project/internal/repository/dao"
	"basic-project/internal/service"
	user "basic-project/internal/web"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	dsn := "root:root@tcp(127.0.0.1:3306)/webook?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	dao.InitTables(db)
	ud := dao.NewUserDao(db)
	ur := repository.NewUserRepository(ud)
	us := service.NewUserService(ur)

	uhld := user.NewUserHandler(us)
	//c := &user.UserHandler{}
	uhld.RegisterRouters(r)
	r.Run(":8081")
}
