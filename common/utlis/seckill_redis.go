package utlis

import (
	"common/global"
	"fmt"
	"strconv"
	"time"
)

const (
	KEY      = "seckill_stock:seckill_id_"
	LOCK_KEY = "seckill_stock"
)

// 获取锁

func CreateLock(lockKey string, lockValue string, expiration time.Duration) bool {
	set, err := global.Rdb.SetNX(global.Ctx, lockKey, lockValue, expiration).Result()
	if err != nil {
		fmt.Println("获取锁时发生错误:", err)
		return false
	}
	return set
}

// 释放锁

func DeleteLock(lockKey string, lockValue string) {
	err := global.Rdb.Del(global.Ctx, lockKey, lockValue).Err()
	if err != nil {
		fmt.Println("释放锁时发生错误:", err)
	}
}

//将秒杀商品添加redis的list中

func SeckillCreateRedis(num, seckillId int) {
	// 生成唯一的锁值
	lockValue := fmt.Sprintf("%d_%d", time.Now().UnixNano(), time.Now().Nanosecond())

	// redis加锁
	if !CreateLock(LOCK_KEY, lockValue, 5*time.Minute) {
		return
	}

	// 释放锁
	defer DeleteLock(LOCK_KEY, lockValue)
	id := strconv.Itoa(seckillId)
	for i := 0; i < num; i++ {
		err := global.Rdb.LPush(global.Ctx, KEY+id, seckillId).Err()
		if err != nil {
			return
		}
	}
}

//判断redis库存是否添加成功

func GetSeckillRedis(seckillId int) int64 {
	id := strconv.Itoa(seckillId)
	return global.Rdb.LLen(global.Ctx, KEY+id).Val()
}

//扣减秒杀商品redis列表库存

func UpdateSeckillRedis(seckillId, num int64) bool {
	id := strconv.Itoa(int(seckillId))
	err := global.Rdb.LRem(global.Ctx, KEY+id, num, seckillId).Err()
	if err != nil {
		return false
	}
	return true
}

//清除redis列表库存

func DelSeckillRedis(seckillId int) {
	id := strconv.Itoa(seckillId)
	global.Rdb.Del(global.Ctx, KEY+id)
}
