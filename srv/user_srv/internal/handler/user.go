package handler

import (
	"common/model"
	"common/model/user_level"
	"common/proto/user"
	"common/utlis"
	"errors"
	"time"
)

// UserLogin  TODO: 用户登录
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
	return &user.UserLoginResponse{UserId: users.Uid}, nil
}

// UserRegister TODO: 用户注册
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
	return &user.UserRegisterResponse{UserId: users.Uid}, nil
}

// UserDetail TODO: 个人资料显示
func UserDetail(in *user.UserDetailRequest) (*user.UserDetailResponse, error) {
	u := model.User{}
	detail, err := u.Detail(int(in.Uid))
	if err != nil {
		return nil, err
	}
	var list []*user.UserDetail
	for _, u := range detail {
		list = append(list, &user.UserDetail{
			Account:        u.Account,
			RealName:       u.RealName,
			Birthday:       u.Birthday,
			Nickname:       u.Nickname,
			Avatar:         u.Avatar,
			Phone:          u.Phone,
			NowMoney:       float32(u.NowMoney),
			BrokeragePrice: float32(u.BrokeragePrice),
			Integral:       float32(u.Integral),
		})
	}

	return &user.UserDetailResponse{Detail: list}, nil
}

// ImproveUser TODO： 完善用户信息
func ImproveUser(in *user.ImproveUserRequest) (*user.ImproveUserResponse, error) {
	u := model.User{
		RealName: in.RealName,        //真实姓名
		Birthday: int64(in.Birthday), //生日
		CardId:   in.CardId,          //身份证号码
		Mark:     in.Mark,            //用户备注
		Nickname: in.Nickname,        //用户昵称
		Avatar:   in.Avatar,          //用户头像
		Phone:    in.Phone,           //手机号码
		Address:  in.Address,         //地址
	}
	Id, err := u.FindId(int(in.Uid))
	if err != nil {
		return nil, err
	}
	if Id.Uid == 0 {
		return nil, errors.New("没有这个用户")
	}

	updated := u.Updated(int(Id.Uid), u)
	if !updated {
		return nil, errors.New("完善用户信息失败")
	}
	return &user.ImproveUserResponse{Success: true}, nil
}

// UpdatedPassword TODO： 修改密码
func UpdatedPassword(in *user.UpdatedPasswordRequest) (*user.UpdatedPasswordResponse, error) {
	u := model.User{}
	Id, err := u.FindId(int(in.Uid))
	if err != nil {
		return nil, err
	}
	if Id.Pwd == utlis.Encryption(in.NewPassword) {
		return nil, errors.New("旧密码和新密码一样，修改失败")
	}
	newPassword := u.UpdatedPassword(int(Id.Uid), utlis.Encryption(in.NewPassword))
	if !newPassword {
		return nil, errors.New("密码修改失败")
	}
	return &user.UpdatedPasswordResponse{Success: true}, nil
}

// UserLevelList TODO:会员页面展示
func UserLevelList(in *user.UserLevelListRequest) (*user.UserLevelListResponse, error) {
	ul := user_level.UserLevel{}
	level, err := ul.FindUsersLevel()
	if err != nil {
		return nil, errors.New("查询失败")
	}
	var list []*user.UserLevelList
	for _, i := range level {
		list = append(list, &user.UserLevelList{
			Name:         i.Name,
			MemberPoints: i.MemberPoints,
			IsShow:       int64(i.IsShow),
			Grade:        int64(i.Grade),
			Image:        i.Image,
			Icon:         i.Icon,
			Explain:      i.Explain,
		})
	}
	return &user.UserLevelListResponse{List: list}, nil
}

// UserLevelPowerList TODO:会员权益页面展示
func UserLevelPowerList(in *user.UserLevelPowerListRequest) (*user.UserLevelPowerListResponse, error) {
	ulp := user_level.UserLevelPower{}
	power, err := ulp.FindUserLevelPower()
	if err != nil {
		return nil, errors.New("查询失败")
	}
	var list []*user.UserLevelPowerList
	for _, i := range power {
		list = append(list, &user.UserLevelPowerList{
			Name:    i.Name,
			Grade:   int64(i.Grade),
			Explain: i.Explain,
		})
	}
	return &user.UserLevelPowerListResponse{List: list}, nil
}

// AddUsePower TODO:用户使用权益
func AddUsePower(in *user.AddUsePowerRequest) (*user.AddUsePowerResponse, error) {
	ulr := user_level.UserLevelRecord{}
	userRecords, err := ulr.FindRecords(int(in.Uid))
	if err != nil {
		return nil, err
	}
	ulup := user_level.UserLevelUsePower{
		Uid: uint32(userRecords.Uid),
		Qid: uint32(userRecords.Grade),
	}
	err = ulup.AddUserPower()
	if err != nil {
		return nil, errors.New("权益使用失败")
	}
	return &user.AddUsePowerResponse{Success: true}, nil
}

// UsePowerList TODO: 用户使用权益表展示
func UsePowerList(in *user.UsePowerListRequest) (*user.UsePowerListResponse, error) {
	ulup := user_level.UserLevelUsePower{}
	levelUsePowers, err := ulup.Finds()
	if err != nil {
		return nil, err
	}
	var list []*user.UsePowerList
	for _, i := range levelUsePowers {
		list = append(list, &user.UsePowerList{
			Uid:     int64(i.Uid),
			Qid:     int64(i.Qid),
			AddTime: i.AddTime.Format(time.DateTime),
		})
	}
	return &user.UsePowerListResponse{List: list}, nil
}

// AddText TODO:会员分添加记录
func AddText(in *user.AddTextRequest) (*user.AddTextResponse, error) {
	ulss := user_level.UserLevelScoreSource{}
	scoreSource, err := ulss.Find()
	if err != nil {
		return nil, err
	}
	switch scoreSource.Id {
	case 1: //消费20元+1积分

	case 2: //邀请一个人注册+20积分

	case 3: //用户签到+5积分

	case 4: //完成特定任务（观看短视频、阅读文章）+10积分

	default:

	}
	//会员分添加记录表
	ulat := user_level.UserLevelAddText{
		Uid:    uint32(in.Uid),
		Source: scoreSource.Source,
		Score:  uint32(scoreSource.Score),
	}
	err = ulat.Add()
	if err != nil {
		return nil, err
	}

	//用户表的剩余积分+++++
	u := model.User{}
	result, err := u.FindId(int(in.Uid))
	points := result.Integral + float64(scoreSource.Score)
	err = u.AddScore(points, in.Uid)
	if err != nil {
		return nil, err
	}

	return &user.AddTextResponse{Success: "会员分添加成功"}, nil
}

// AddUserAddress TODO:用户添加地址
func AddUserAddress(in *user.AddUserAddressRequest) (*user.AddUserAddressResponse, error) {
	u := model.User{}
	FindUser, err := u.FindId(int(in.Uid))
	if err != nil {
		return nil, err
	}
	ua := model.UserAddress{
		Uid:      FindUser.Uid,
		RealName: FindUser.RealName,
		Phone:    FindUser.Phone,
		Province: in.Province, //收货人所在省
		City:     in.City,     //收货人所在市
		District: in.District, //收货人所在区
		Detail:   in.Detail,   //收货人详细地址
	}
	err = ua.Created()
	if err != nil {
		return nil, errors.New("地址添加失败")
	}
	return &user.AddUserAddressResponse{Success: "地址添加成功"}, nil
}
