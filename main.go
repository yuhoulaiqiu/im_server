package main

import (
	"flag"
	"fmt"
	"yim_server/core"
	"yim_server/yim_chat/chat_models"
	"yim_server/yim_group/group_models"
	"yim_server/yim_user/user_models"
)

type Option struct {
	DB bool
}

func main() {
	var opt Option
	flag.BoolVar(&opt.DB, "db", false, "初始化数据库")
	flag.Parse()
	mysqlDataSource := "root:zxc3240858086@tcp(127.0.0.1:3306)/yim_server_db?charset=utf8mb4&parseTime=True&loc=Local"
	if opt.DB {
		db := core.InitMysql(mysqlDataSource)
		err := db.AutoMigrate(&user_models.UserModel{},
			&user_models.FriendModel{},
			&user_models.FriendVerifyModel{},
			&user_models.FriendModel{},
			&user_models.UserConfModel{},

			&chat_models.ChatModel{},

			&group_models.GroupMemberModel{},
			&group_models.GroupModel{},
			&group_models.GroupVerifyModel{},
			&group_models.GroupMsgModel{},
		)
		if err != nil {
			fmt.Println("数据库初始化失败", err)
			return
		}
		fmt.Println("数据库初始化成功")
	}
}
