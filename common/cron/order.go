package cron

import (
	"common/model"
	"common/utlis"
	"fmt"
	"github.com/robfig/cron/v3"
)

func OrderCron(orderSn string) {
	c := cron.New(cron.WithSeconds())

	//30分钟之后检测订单的支付状态
	_, err := c.AddFunc("0 */30 * * * *", func() {
		o := &model.Order{}
		err := o.GetOrderStatus(orderSn)
		if err != nil {
			fmt.Printf("获取订单状态失败: %v\n", err)
			return
		}
		if o.Status != 1 {
			err = o.UpdateOrderStatus(orderSn, 3)
			if err != nil {
				fmt.Printf("更新订单状态失败: %v\n", err)
				return
			}
			// 订单未支付，返还 redis，mysql 库存
			utlis.SeckillCreateRedis(int(o.TotalNum), int(o.SeckillId))
			s := &model.Seckill{}
			err = s.StartsStockReverse(o.SeckillId, o.TotalNum)
			if err != nil {
				fmt.Printf("返还库存失败: %v\n", err)
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
