package handler

import (
	"common/global"
	"common/model"
	"common/model/user_level"
	"common/proto/user"
	"common/utlis"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"strconv"
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
	detail, err := u.Detail(in.Uid)
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
	Id, err := u.FindId(int(in.Uid))
	if err != nil {
		return nil, err
	}
	if Id.Uid == 0 {
		return nil, errors.New("没有这个用户")
	}

	updated := u.Updated(Id.Uid, u)
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
	newPassword := u.UpdatedPassword(Id.Uid, utlis.Encryption(in.NewPassword))
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
		Uid:       FindUser.Uid,
		RealName:  FindUser.RealName,
		Phone:     FindUser.Phone,
		Province:  in.Province, //收货人所在省
		City:      in.City,     //收货人所在市
		District:  in.District, //收货人所在区
		Detail:    in.Detail,   //收货人详细地址
		IsDefault: in.IsDefault,
	}
	err = ua.Created()
	if err != nil {
		return nil, errors.New("地址添加失败")
	}
	return &user.AddUserAddressResponse{Success: "地址添加成功"}, nil
}

// UserSignIn TODO:用户签到
func UserSignIn(in *user.UserSignInRequest) (*user.UserSignInResponse, error) {
	// 如果传入了SignDate，支持自定义签到日期（方便测试）
	var signDate time.Time
	var err error
	if in.SignDate != "" {
		signDate, err = time.Parse("2006-01-02", in.SignDate) //测试的话格式(2025-03-28)
		if err != nil {
			return nil, fmt.Errorf("无效的签到日期格式")
		}
	} else {
		signDate = time.Now() // 默认使用当前日期
	}
	today := signDate.Format("2006-01-02")
	//1.检查今天是否已经签到
	todaykey := fmt.Sprintf("sign:user:%d:%s", in.UserId, today)
	offset := signDate.Day() // 位图的偏移量从0开始
	bit, err := global.Rdb.GetBit(global.Ctx, todaykey, int64(offset)).Result()
	if err != nil {
		return nil, err
	}
	if bit == 1 {
		return nil, fmt.Errorf("今天签到了")
	}
	//2.检查昨天是否已经签到，计算连续签到天数
	consecutiveDays := 1 // 默认连续1天
	yesterday := signDate.AddDate(0, 0, -1).Format("2006-01-02")
	yesterdayKey := fmt.Sprintf("sign:user:%d:%s", in.UserId, yesterday)
	//检查昨天的签到是否还存在
	exists, err := global.Rdb.Exists(global.Ctx, yesterdayKey).Result()
	if err != nil {
		return nil, fmt.Errorf("昨天的签到不存在")
	}
	// 如果昨天有签到，获取连续签到天数
	if exists > 0 {
		consecutiveKey := fmt.Sprintf("sign:consecutive:%d", in.UserId)
		days, err := global.Rdb.Get(global.Ctx, consecutiveKey).Int()
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
	userIntegral := &model.UserIntegral{}
	err = userIntegral.GetUserIntegral(in.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果没有记录，创建新记录
			userIntegral = &model.UserIntegral{
				UserId:        in.UserId,
				Integral:      int64(points),
				IntegralTotal: int64(points),
				CreateTime:    signDate,
				UpdateTime:    signDate,
			}
			err = userIntegral.AddUserIntegral()
			if err != nil {
				tx.Rollback()
				return nil, fmt.Errorf("创建积分记录失败")
			}
		} else {
			tx.Rollback()
			return nil, err
		}
	} else {
		// 更新现有记录
		err = userIntegral.UpdateUserIntegral(in.UserId, int64(points))
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("更新积分失败")
		}
	}
	// 6. 创建积分流水记录
	integralLog := &model.UserIntegralLog{
		UserId:        in.UserId,
		IntegralType:  model.IntegralTypeContinuous, // 连续签到类型
		Integral:      int64(points),
		Bak:           fmt.Sprintf("连续签到%d天", consecutiveDays),
		OperationTime: signDate,
		CreateTime:    signDate,
	}
	err = integralLog.AddUserIntegralLog()
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("创建积分流水失败")
	}
	if integralLog.Id == 0 {
		tx.Rollback()
		return nil, fmt.Errorf("创建积分流水失败")
	}

	// 7. 更新Redis中的签到状态和连续签到天数
	consecutiveKey := fmt.Sprintf("sign:consecutive:%d", in.UserId)
	pipe := global.Rdb.TxPipeline()
	pipe.SetBit(global.Ctx, todaykey, int64(offset), 1)
	pipe.Set(global.Ctx, consecutiveKey, consecutiveDays, 30*24*time.Hour) // 保留30天
	if _, err := pipe.Exec(global.Ctx); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("更新签到状态失败")
	}
	// 8. 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("提交事务失败")
	}
	return &user.UserSignInResponse{
		Message: fmt.Sprintf("签到成功，连续签到%d天", consecutiveDays),
		Points:  int64(points),
	}, nil
}

// UserMakeupSignIn TODO:用户补签
func UserMakeupSignIn(in *user.UserMakeupSignInRequest) (*user.UserMakeupSignInResponse, error) {
	// 1. 解析补签日期
	makeupDate, err := time.Parse("2006-01-02", in.SignDate)
	if err != nil {
		return nil, errors.New("无效的补签日期格式")
	}
	// 2. 检验补签日期 是不是在一周之内的
	if time.Since(makeupDate) > 7*24*time.Hour {
		return nil, errors.New("只能补签过去7天内的签到")
	}

	// 3. 检查是否已签到
	dateKey := fmt.Sprintf("sign:user:%d:%s", in.UserId, makeupDate.Format("2006-01-02"))
	offset := makeupDate.Day() - 1
	bit, err := global.Rdb.GetBit(global.Ctx, dateKey, int64(offset)).Result()
	if err != nil {
		return nil, err
	}
	if bit == 1 {
		return nil, errors.New("该日期已签到，无需补签")
	}

	// 4. 检查用户是否有补签卡
	makeupCard := &model.UserMakeup{}
	err = makeupCard.GetUserMakeupCard(in.UserId)
	if err != nil {
		return nil, errors.New("没有可用的补签卡")
	}
	if makeupCard.Cardcount <= 0 {
		return nil, errors.New("没有可用的补签卡")
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
	err = makeupCard.UpdateUserMakeupCard(in.UserId)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("扣除补签卡失败")
	}

	// 8. 更新用户积分
	ui := &model.UserIntegral{}
	err = ui.UpdateUserIntegral(in.UserId, int64(points))
	if err != nil {
		tx.Rollback()
		return nil, errors.New("更新积分失败")
	}

	// 9. 创建积分流水记录
	integralLog := &model.UserIntegralLog{
		UserId:        in.UserId,
		IntegralType:  model.IntegralTypeReplenish, // 补签类型
		Integral:      int64(points),
		Bak:           "补签获得",
		OperationTime: makeupDate,
		CreateTime:    time.Now(),
	}
	err = integralLog.AddUserIntegralLog()
	if err != nil {
		tx.Rollback()
		return nil, errors.New("创建积分流水失败")
	}

	// 10. 更新Redis签到状态（但不更新连续签到）
	_, err = global.Rdb.SetBit(global.Ctx, dateKey, int64(offset), 1).Result()
	if err != nil {
		tx.Rollback()
		return nil, errors.New("更新签到状态失败")
	}

	// 11. 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, errors.New("提交事务失败")
	}

	return &user.UserMakeupSignInResponse{
		Success: true,
		Message: "补签成功",
		Points:  int64(points),
	}, nil
}

// UserApplication TODO:用户申请发票
func UserApplication(in *user.UserApplicationRequest) (*user.UserApplicationResponse, error) {
	u := model.User{}
	FindUser, err := u.FindId(int(in.UserId))
	if err != nil {
		return nil, errors.New("用户查询失败")
	}
	o := model.Order{}
	FindOrder, err := o.FindId(in.OrderId)
	if err != nil {
		return nil, errors.New("订单查找失败")
	}
	ua := model.UserAddress{}
	FindUserAddress, err := ua.FindId(in.UserId)
	if err != nil {
		return nil, errors.New("用户地址查询失败")
	}

	ia := model.InvoiceApplication{
		UserId:                       in.UserId,
		OrderId:                      in.OrderId,
		InvoiceType:                  in.InvoiceType,                  //发票类型：普通发票、增值税专用发票
		InvoiceTitle:                 in.InvoiceTitle,                 //发票抬头
		TaxpayerIdentificationNumber: in.TaxpayerIdentificationNumber, //纳税人识别号
		InvoiceAmount:                FindOrder.PayPrice,              //发票金额
		Email:                        FindUser.Email,
		Address:                      FindUserAddress.Detail,
		Phone:                        FindUser.Phone,
		Type:                         in.Type, //发票材质：纸质、电子
		MerId:                        FindOrder.MerId,
	}
	err = ia.UserApplication()
	if err != nil {
		return nil, errors.New("用户发票申请失败")
	}
	return &user.UserApplicationResponse{Success: "用户成功申请发票"}, nil
}

// UpdatedAddress TODO:用户修改地址
func UpdatedAddress(in *user.UpdatedAddressRequest) (*user.UpdatedAddressResponse, error) {
	//查找用户
	u := model.User{}
	FindUser, err := u.FindId(int(in.Uid))
	if err != nil {
		return nil, err
	}

	//查找用户的所有地址，是否是默认的
	ua := model.UserAddress{}
	findDefault, err := ua.FindDefault(FindUser.Uid, in.UserAddressId)
	if err != nil {
		return nil, err
	}

	//  findDefault   该地址不是默认地址是，修改成默认地址，并把原先的修改成不默认
	if findDefault == true {
		err = ua.UpdatedAddressDefault(FindUser.Uid)
		if err != nil {
			return nil, errors.New("修改失败")
		}

		ua = model.UserAddress{
			Uid:       FindUser.Uid,
			RealName:  in.RealName,
			Phone:     in.Phone,
			Province:  in.Province, //收货人所在省
			City:      in.City,     //收货人所在市
			District:  in.District, //收货人所在区
			Detail:    in.Detail,   //收货人详细地址
			IsDefault: 1,
		}
		err = ua.UpdatedAddress(in.UserAddressId)
		if err != nil {
			return nil, errors.New("用户地址修改失败")
		}
	}
	ua = model.UserAddress{
		Uid:      FindUser.Uid,
		RealName: in.RealName,
		Phone:    in.Phone,
		Province: in.Province, //收货人所在省
		City:     in.City,     //收货人所在市
		District: in.District, //收货人所在区
		Detail:   in.Detail,   //收货人详细地址
	}
	err = ua.UpdatedAddress(in.UserAddressId)
	if err != nil {
		return nil, errors.New("用户地址修改失败")
	}

	return &user.UpdatedAddressResponse{Success: "用户地址修改成功"}, nil
}

// UserReceiveCoupon TODO:用户领取优惠券
func UserReceiveCoupon(in *user.UserReceiveCouponRequest) (*user.UserReceiveCouponResponse, error) {
	cou := &model.Coupon{}
	if err := cou.GetCouponIdBy(in.CouponId); err != nil {
		return nil, err
	}
	uc := &model.CouponUser{}
	err := uc.GetUserCouponIdBy(in.CouponId, in.UserId)
	if err != nil {
		return nil, err
	}
	if uc.Id != 0 {
		return nil, errors.New("不可重复领取")
	}
	addTime, _ := strconv.Atoi(time.Now().AddDate(0, 0, 0).Format("20060102"))
	endTime, _ := strconv.Atoi(time.Now().AddDate(0, 0, int(cou.CouponTime)).Format("20060102"))
	uc = &model.CouponUser{
		Cid:         in.CouponId,
		Uid:         in.UserId,
		CouponTitle: cou.Title,
		CouponPrice: cou.CouponPrice,
		UseMinPrice: cou.UseMinPrice,
		AddTime:     int64(addTime),
		EndTime:     int64(endTime),
		Status:      0,
	}
	if err := uc.AddCouponUser(); err != nil {
		return nil, err
	}
	if uc.Id == 0 {
		return nil, errors.New("领取优惠券失败")
	}
	return &user.UserReceiveCouponResponse{Success: true}, nil
}

// UserWithdraw TODO: 用户提现
func UserWithdraw(in *user.UserWithdrawRequest) (*user.UserWithdrawResponse, error) {
	//检查用户是否存在
	u := model.User{}
	userInfo, err := u.FindId(int(in.UserId))
	if err != nil {
		return nil, err
	}
	if userInfo.Uid == 0 {
		return nil, errors.New("用户不存在")
	}

	//检查用户余额是否足够
	if userInfo.NowMoney < float64(in.Amount) {
		return nil, errors.New("余额不足")
	}

	//开启事务处理
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 扣除用户余额
	newBalance := userInfo.NowMoney - float64(in.Amount)
	err = u.UpdateBalance(in.UserId, newBalance)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	//记录提现记录
	ue := &model.UserExtract{
		Uid:          in.UserId,
		RealName:     userInfo.RealName,
		ExtractType:  in.WithdrawMethod, //支付类型
		BankCode:     in.AccountInfo,
		AlipayCode:   in.AccountInfo, //支付宝账号
		ExtractPrice: float64(in.Amount),
		AddTime:      time.Time{},
		Status:       0,
		Wechat:       in.AccountInfo, //微信号
	}
	if !ue.CreateUserExtract() {
		tx.Rollback()
		return nil, err
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, err
	}

	return &user.UserWithdrawResponse{Success: true}, nil
}

// UserAddressList TODO: 用户地址列表
func UserAddressList(in *user.UserAddressListRequest) (*user.UserAddressListResponse, error) {
	ua := model.UserAddress{}
	userAddress, err := ua.FindIds(in.Uid)
	if err != nil {
		return nil, err
	}
	var userAddressList []*user.UserAddressList
	for _, i := range userAddress {
		userAddressList = append(userAddressList, &user.UserAddressList{
			UserAddressId: i.Id,
			RealName:      i.RealName,
			Phone:         i.Phone,
			Province:      i.Province,
			City:          i.City,
			District:      i.District,
			Detail:        i.Detail,
			IsDefault:     i.IsDefault,
		})
	}
	return &user.UserAddressListResponse{List: userAddressList}, nil
}
