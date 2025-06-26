package srvhandler

import (
	"context"
	"fmt"
	"math/rand"
	"time"
	"zg6zk1/apiway/basic/global"
	"zg6zk1/apiway/basic/models"
	__ "zg6zk1/service/protobuf"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedUserServer
}

func (s *Server) SendSms(_ context.Context, req *__.SendSmsRequest) (*__.SendSmsResponse, error) {

	// 查询数据库是否注册
	var login models.Logins
	err := global.DB.Where("mobile LIKE ?", "%"+req.Mobile+"%").Find(&login).Error
	//未注册直接创建
	if err != nil {
		return &__.SendSmsResponse{Msg: "fail"}, nil
	}

	Intn := rand.Intn(9000) + 1000
	global.Client.Set(context.Background(), req.Mobile, Intn, time.Hour*10)
	fmt.Println(Intn)
	return &__.SendSmsResponse{Msg: "success"}, nil

}
func (s *Server) Register(_ context.Context, req *__.RegisterRequest) (*__.RegisterResponse, error) {

	// 查询数据库是否注册
	var login models.Logins
	err := global.DB.Where("mobile LIKE ?", "%"+req.Mobile+"%").Find(&login).Error
	//未注册直接创建

	if err != nil {
		var create = models.Users{
			Username:    req.Mobile,
			HeaderImage: req.HeaderImage,
			Mobile:      req.Mobile,
		}
		global.DB.Create(&create)
	}

	return &__.RegisterResponse{Msg: "success"}, nil

}

func (s *Server) Login(_ context.Context, req *__.LoginRequest) (*__.LoginResponse, error) {

	// 获取验证码
	get := global.Client.Get(context.Background(), req.Mobile)
	if get.Val() != req.SendSms {
		return &__.LoginResponse{Msg: "fail"}, nil
	}

	// 查询数据库，获取用户信息
	var login models.Logins
	err := global.DB.Where("mobile LIKE ?", "%"+req.Mobile+"%").Find(&login).Error
	if err != nil {
		return &__.LoginResponse{Msg: "fail"}, nil
	}

	return &__.LoginResponse{Msg: "success"}, nil

}

// List函数用于根据手机号查询用户列表
func (s *Server) List(_ context.Context, req *__.ListRequest) (*__.ListResponse, error) {

	// 定义一个用户切片
	var user []models.Users
	// 在数据库中查询手机号包含req.Mobile的用户
	global.DB.Where("mobile LIKE ?", "%"+req.Mobile+"%").Find(&user)

	// 定义一个ListInfo切片
	var slice []*__.ListInfo

	// 遍历用户切片
	for _, V := range user {
		// 将用户信息添加到ListInfo切片中
		slice = append(slice, &__.ListInfo{
			Username:    V.Username,
			HeaderImage: V.HeaderImage,
		})
	}

	// 返回查询结果
	return &__.ListResponse{
		List: slice,
		Msg:  "success"}, nil
}
