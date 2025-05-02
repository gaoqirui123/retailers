package cron

import (
	"common/model"
	"common/utlis"
	"github.com/robfig/cron/v3"
)

func OrderCron(orderSn string) {
	c := cron.New(cron.WithSeconds())

	//30分钟之后检测订单的支付状态
	c.AddFunc("*/30 * * * * ?", func() {
		o := model.Order{}
		err := o.GetOrderStatus(orderSn)
		if err != nil {
			return
		}
		if o.Status != 1 {
			err = o.UpdateOrderStatus(orderSn, 3)
			if err != nil {
				return
			}
		}
		// 订单未支付，反还redis，mysql库存
		if o.Status != 1 {
			utlis.ProductCreateRedis(int(o.TotalNum), int(o.SeckillId))
			s := &model.Seckill{}
			err = s.StartsStockReverse(o.SeckillId, o.TotalNum)
			if err != nil {
				return
			}
		}
		c.Stop()
	})
	c.Start()
}
