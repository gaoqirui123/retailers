package cron

import (
	"common/global"
	"common/model"
	"fmt"
	"github.com/robfig/cron/v3"
)

func PinkCron(unique int64) {
	c := cron.New(cron.WithSeconds())
	//24小时之后检测拼团的状态
	_, err := c.AddFunc("* * */24 * * *", func() {
		// 如果没有拼团成功，那就将参与拼团的单状态全部设置为取消
		pink := &model.Pink{}
		err := pink.GetPainUnique(unique)
		if err != nil {
			fmt.Printf("查询拼团失败: %v\n", err)
			return
		}
		if pink.Status == 0 {
			tx := global.DB.Begin()
			o := &model.Order{}
			err = o.UpdateOrderStatus(pink.OrderId, 9)
			if err != nil {
				tx.Rollback()
				fmt.Printf("更新订单状态失败: %v\n", err)
				return
			}
			// 拼团失败，返还库存
			com := &model.Combination{}
			err = com.ReverseCombinationStock(int64(pink.Pid), int64(pink.TotalNum))
			if err != nil {
				tx.Rollback()
				fmt.Printf("返还库存失败: %v\n", err)
				return
			}
			err = tx.Commit().Error
			if err != nil {
				return
			}
		}
	})
	if err != nil {
		fmt.Printf("添加定时任务失败: %v\n", err)
		return
	}
	c.Start()
	defer c.Stop()
}
