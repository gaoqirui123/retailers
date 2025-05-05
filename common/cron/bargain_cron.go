package cron

//func BargainCron(orderSn string) {
//	c := cron.New(cron.WithSeconds())
//	//24小时之后检测砍价的状态
//	_, err := c.AddFunc("* * */24 * * *", func() {
//		// 如果没有砍成功，那就将此订单状态设置为取消
//		bar := &model.User{}
//		bar.GetBargainStatus()
//		if o.Status != 1 {
//			o := &model.Order{}
//			err := o.UpdateOrderStatus(orderSn, 9)
//			if err != nil {
//				fmt.Printf("更新订单状态失败: %v\n", err)
//				return
//			}
//			// 砍价失败，返还库存
//			bar.UpdateBargainStock()
//		}
//	})
//	if err != nil {
//		fmt.Printf("添加定时任务失败: %v\n", err)
//		return
//	}
//	c.Start()
//	defer c.Stop()
//}
