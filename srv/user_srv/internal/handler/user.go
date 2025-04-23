package handler

import (
	"common/model"
	"common/model/user_level"
	"common/proto/user"
	"common/utlis"
	"errors"
	"github.com/google/uuid"
	"time"
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

// TODO: 个人资料显示
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

// TODO： 完善用户信息
func ImproveUser(in *user.ImproveUserRequest) (*user.ImproveUserResponse, error) {
	u := model.User{
		RealName: in.RealName, //真实姓名
		Birthday: in.Birthday, //生日
		CardId:   in.CardId,   //身份证号码
		Mark:     in.Mark,     //用户备注
		Nickname: in.Nickname, //用户昵称
		Avatar:   in.Avatar,   //用户头像
		Phone:    in.Phone,    //手机号码
		Address:  in.Address,  //地址
	}
	Id, err := u.FindId(int(in.Id))
	if err != nil {
		return nil, err
	}
	if Id.Uid == 0 {
		return &user.ImproveUserResponse{Success: "没有这个用户"}, nil
	}

	updated := u.Updated(int(Id.Uid), u)
	if !updated {
		return &user.ImproveUserResponse{Success: "完善用户信息失败"}, nil
	}

	return &user.ImproveUserResponse{Success: "完善用户信息成功"}, nil
}

// TODO： 修改密码
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
	return &user.UpdatedPasswordResponse{Success: "密码修改成功"}, nil
}

// TODO:会员页面展示
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
			IsShow:       int32(i.IsShow),
			Grade:        i.Grade,
			Image:        i.Image,
			Icon:         i.Icon,
			Explain:      i.Explain,
		})
	}
	return &user.UserLevelListResponse{List: list}, nil
}

// TODO:会员权益页面展示
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
			Grade:   int32(i.Grade),
			Explain: i.Explain,
		})
	}
	return &user.UserLevelPowerListResponse{List: list}, nil
}

// GroupBuying TODO:用户发起拼团
func GroupBuying(in *user.GroupBuyingRequest) (*user.GroupBuyingResponse, error) {
	// 假设拼团时长为 1 小时，计算结束时间
	addtime := time.Now().Add(0).Format("2006-01-02 15:04:05")
	stopTime := time.Now().Add(time.Hour).Format("2006-01-02 15:04:05")
	c := model.Combination{}
	combination, err := c.GetCombinationById(in.Pid)
	if err != nil {
		return nil, err
	}
	orderId := uuid.New().String()
	totalPrice := float64(in.Num) * combination.Price
	p := model.Pink{
		Uid:        int(in.Uid),
		OrderId:    orderId,
		OrderIdKey: 0,
		TotalNum:   int(in.Num),
		TotalPrice: totalPrice,
		Cid:        int(in.Pid),
		Pid:        combination.ProductId,
		People:     combination.People,
		Price:      combination.Price,
		AddTime:    addtime,
		StopTime:   stopTime,
	}
	err = p.Create()
	if err != nil {
		return nil, err
	}
	return &user.GroupBuyingResponse{Success: "发起拼团成功"}, nil
}

// TODO:用户使用权益
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
	return &user.AddUsePowerResponse{Success: "用户使用权益成功"}, nil
}

// TODO: 用户使用权益表展示
func UsePowerList(in *user.UsePowerListRequest) (*user.UsePowerListResponse, error) {
	ulup := user_level.UserLevelUsePower{}
	levelUsePowers, err := ulup.Finds()
	if err != nil {
		return nil, err
	}
	var list []*user.UsePowerList
	for _, i := range levelUsePowers {
		list = append(list, &user.UsePowerList{
			Uid:     int32(i.Uid),
			Qid:     int32(i.Qid),
			AddTime: i.AddTime.Format(time.DateTime),
		})
	}
	return &user.UsePowerListResponse{List: list}, nil
}
