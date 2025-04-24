package handler

import (
	"common/global"
	"common/model"
	"common/proto/user"
	"common/utlis"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	return &user.UserLoginResponse{UserId: uint64(users.Uid)}, nil
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
	return &user.UserRegisterResponse{UserId: uint64(users.Uid)}, nil
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
	ul := model.EbSystemUserLevel{}
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

// UserLevelPowerList TODO:会员权益页面展示
func UserLevelPowerList(in *user.UserLevelPowerListRequest) (*user.UserLevelPowerListResponse, error) {
	ulp := model.EbSystemUserPower{}
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

// GroupBuying GroupBuying TODO:用户发起拼团
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
	return &user.GroupBuyingResponse{Success: true}, nil
}

// UserSignIn TODO:用户签到
func UserSignIn(in *user.UserSignInRequest) (*user.UserSignInResponse, error) {
	// 如果传入了SignDate，支持自定义签到日期（方便测试）
	var signDate time.Time
	if in.SignDate != "" {
		signDate, err := time.Parse("2006-01-02", in.SignDate) //测试的话格式(2025-03-28)
		if err != nil {
			return nil, fmt.Errorf("无效的签到日期格式")
		}
	} else {
		signDate = time.Now() // 默认使用当前日期
	}
	today := signDate.Format("2006-01-02")

	//1.检查今天是否已经签到
	todaykey := fmt.Sprintf("sign:user:%d:%s", in.UserId, today)
	offset := signDate.Day() - 1 // 位图的偏移量从0开始
	bit, err := global.Rdb.GetBit(global.CTX, todaykey, int64(offset)).Result()
	if err != nil {
		return nil, err
	}
	if bit == 1 {
		return nil, fmt.Errorf("今天签到了")
	}
	//2.检查昨天是否已经签到，计算连续签到天数
	consecutiveDays := 1 // 默认连续1天
	yesterday := signDate.AddDate(0, 0, -1).Format("2006-01-02")
	yesterdayKey := fmt.Sprintf("sign:user:%d:%s", req.UserId, yesterday)
	//检查昨天的签到是否还存在
	exists, err := global.Rdb.Exists(global.CTX, yesterdayKey).Result()
	if err != nil {
		return nil, fmt.Errorf("昨天的签到不存在")
	}
	// 如果昨天有签到，获取连续签到天数
	if exists > 0 {
		consecutiveKey := fmt.Sprintf("sign:consecutive:%d", req.UserId)
		days, err := global.Rdb.Get(global.CTX, consecutiveKey).Int()
		if err != nil && err != redis.Nil {
			return nil, err
		}
		if days > 0 {
			consecutiveDays = days + 1
		}
	}
	// 3. 计算本次签到应得积分
	points := consecutiveDays // 第N天连续签到得N分
	// 4. 开启事务处理
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 5. 更新用户总积分
	var userIntegral model.UserIntegral
	if err := tx.Where("user_id = ?", req.UserId).First(&userIntegral).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果没有记录，创建新记录
			userIntegral = model.UserIntegral{
				ID:            uuid.New().String(),
				UserID:        int(req.UserId),
				Integral:      points,
				IntegralTotal: points,
				CreateTime:    signDate,
				UpdateTime:    signDate,
			}
			if err := tx.Create(&userIntegral).Error; err != nil {
				tx.Rollback()
				return nil, fmt.Errorf("创建积分记录失败")
			}
		} else {
			tx.Rollback()
			return nil, err
		}
	} else {
		// 更新现有记录
		if err := tx.Model(&userIntegral).Updates(map[string]interface{}{
			"integral":       gorm.Expr("integral + ?", points),
			"integral_total": gorm.Expr("integral_total + ?", points),
			"update_time":    signDate,
		}).Error; err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("更新积分失败")
		}
	}
	// 6. 创建积分流水记录
	integralLog := model.UserIntegralLog{
		ID:            uuid.New().String(),
		UserID:        int(req.UserId),
		IntegralType:  model.IntegralTypeContinuous, // 连续签到类型
		Integral:      points,
		Bak:           fmt.Sprintf("连续签到%d天", consecutiveDays),
		OperationTime: signDate,
		CreateTime:    signDate,
	}
	if err := tx.Create(&integralLog).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("创建积分流水失败")
	}

	// 7. 更新Redis中的签到状态和连续签到天数
	consecutiveKey := fmt.Sprintf("sign:consecutive:%d", req.UserId)
	pipe := global.Rdb.TxPipeline()
	pipe.SetBit(global.CTX, todaykey, int64(offset), 1)
	pipe.Set(global.CTX, consecutiveKey, consecutiveDays, 30*24*time.Hour) // 保留30天
	if _, err := pipe.Exec(global.CTX); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("更新签到状态失败")
	}
	// 8. 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("提交事务失败")
	}
	return &user.UserSignInResponse{
		Success: true,
		Message: fmt.Sprintf("签到成功，连续签到%d天", consecutiveDays),
		Points:  int32(points),
	}, nil

}

func UserMakeupSignIn(in *user.UserMakeupSignInRequest) (*user.UserMakeupSignInResponse, error) {
	// 1. 解析补签日期
	makeupDate, err := time.Parse("2006-01-02", in.SignDate)
	if err != nil {
		return nil, fmt.Errorf("无效的补签日期格式")
	}

	// 2. 检验补签日期 是不是在一周之内的
	if time.Since(makeupDate) > 7*24*time.Hour {
		return nil, fmt.Errorf("只能补签过去7天内的签到")
	}

	// 3. 检查是否已签到
	dateKey := fmt.Sprintf("sign:user:%d:%s", in.UserId, makeupDate.Format("2006-01-02"))
	offset := makeupDate.Day() - 1
	bit, err := global.Rdb.GetBit(ctx, dateKey, int64(offset)).Result()
	if err != nil {
		return nil, err
	}
	if bit == 1 {
		return nil, fmt.Errorf("该日期已签到，无需补签")
	}

	// 4. 检查用户是否有补签卡
	var makeupCard model.UserMakeupCard
	if err := global.DB.Where("user_id = ?", in.UserId).First(&makeupCard).Error; err != nil {
		return nil, fmt.Errorf("没有可用的补签卡")
	}
	if makeupCard.CardCount <= 0 {
		return nil, fmt.Errorf("没有可用的补签卡")
	}

	// 5. 计算积分（补签固定得1分）
	points := 1

	// 6. 开启事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 7. 扣除补签卡
	if err := tx.Model(&model.UserMakeupCard{}).
		Where("user_id = ? AND cardCount > 0", in.UserId).
		Update("cardCount", gorm.Expr("cardCount - 1")).
		Update("update_time", time.Now()).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("扣除补签卡失败")
	}

	// 8. 更新用户积分
	if err := tx.Model(&model.UserIntegral{}).
		Where("user_id = ?", in.UserId).
		Updates(map[string]interface{}{
			"integral":       gorm.Expr("integral + ?", points),
			"integral_total": gorm.Expr("integral_total + ?", points),
			"update_time":    time.Now(),
		}).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("更新积分失败")
	}

	// 9. 创建积分流水记录
	integralLog := model.Sign{
		ID:            uuid.New().String(),
		UserID:        int(in.UserId),
		IntegralType:  model.IntegralTypeReplenish, // 补签类型
		Integral:      points,
		Bak:           "补签获得",
		OperationTime: makeupDate,
		CreateTime:    time.Now(),
	}
	if err := tx.Create(&integralLog).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("创建积分流水失败")
	}

	// 10. 更新Redis签到状态（但不更新连续签到）
	_, err = global.Rdb.SetBit(global.Ctx, dateKey, int64(offset), 1).Result()
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("更新签到状态失败")
	}

	// 11. 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("提交事务失败")
	}

	return &user.UserMakeupSignInResponse{
		Success: true,
		Message: "补签成功",
		Points:  int32(points),
	}, nil
}
