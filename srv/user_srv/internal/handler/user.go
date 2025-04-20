package handler

import (
	"common/model"
	"common/proto/user"
	"common/utlis"
	"errors"
)

func UserLogin(in *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	users := &model.User{}
	err := users.UserLogin(in.Account)
	if err != nil {
		return nil, err
	}
	if users.Uid == 0 {
		return nil, errors.New("账号未注册，请先去注册")
	}
	if utlis.Encryption(in.PassWord) != users.Pwd {
		return nil, errors.New("密码错误，请重新输入")
	}
	return &user.UserLoginResponse{UserId: uint64(users.Uid)}, nil
}

func UserRegister(in *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	if in.PassWord != in.Pass {
		return nil, errors.New("两次密码不一致")
	}
	users := &model.User{}
	err := users.UserLogin(in.Account)
	if err != nil {
		return nil, err
	}
	if users.Uid != 0 {
		return nil, errors.New("账号已注册，请去登录")
	}
	//val := redis.GetSendSms(in.Mobile, in.Source)
	//if val != in.SendSms {
	//	return nil, errors.New("验证码错误，请重新输入")
	//}
	//if in.PassWord != in.Pass {
	//	return nil, errors.New("两次密码不一致")
	//}
	users = &model.User{
		Account: in.Account,
		Pwd:     utlis.Encryption(in.PassWord),
	}
	err = users.UserRegister()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	if users.Uid == 0 {
		return nil, errors.New("注册失败")
	}
	//err = redis.DelSendSms(in.Mobile, in.Source)
	//if err != nil {
	//	return nil, errors.New("验证码销毁失败")
	//}
	return &user.UserRegisterResponse{UserId: uint64(users.Uid)}, nil
}
