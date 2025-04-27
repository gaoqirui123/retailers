package user_level

import "common/global"

// 会员分数提升来源表
type UserLevelScoreSource struct {
	Id     uint32 `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	Score  int32  `gorm:"column:score;type:int;comment:提升的会员分;default:NULL;" json:"score"`         // 提升的会员分
	Source string `gorm:"column:source;type:varchar(30);comment:提升来源;default:NULL;" json:"source"` // 提升来源
}

func (s *UserLevelScoreSource) Find() (result UserLevelScoreSource, err error) {
	err = global.DB.Debug().Table("user_level_score_source").Find(&result).Error
	if err != nil {
		return UserLevelScoreSource{}, err
	}
	return result, nil
}
