package model

type Sign struct {
	Id      int32  `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	Uid     int32  `gorm:"column:uid;type:int;comment:用户uid;not null;default:0;" json:"uid"`          // 用户uid
	Title   string `gorm:"column:title;type:varchar(255);comment:签到说明;not null;" json:"title"`        // 签到说明
	Number  int32  `gorm:"column:number;type:int;comment:获得积分;not null;default:0;" json:"number"`     // 获得积分
	Balance int32  `gorm:"column:balance;type:int;comment:剩余积分;not null;default:0;" json:"balance"`   // 剩余积分
	AddTime int32  `gorm:"column:add_time;type:int;comment:添加时间;not null;default:0;" json:"add_time"` // 添加时间
}
