package utlis

import (
	"common/global"
	"fmt"
	"strconv"
	"time"
)

const (
	KEY      = "start_stock:start_id_"
	LOCK_KEY = "start_lock"
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

func ProductCreateRedis(num, startId int) {
	// 生成唯一的锁值
	lockValue := fmt.Sprintf("%d_%d", time.Now().UnixNano(), time.Now().Nanosecond())

	// redis加锁
	if !CreateLock(LOCK_KEY, lockValue, 5*time.Minute) {
		return
	}

	// 释放锁
	defer DeleteLock(LOCK_KEY, lockValue)

	id := strconv.Itoa(startId)
	for i := 0; i < num; i++ {
		err := global.Rdb.LPush(global.Ctx, KEY+id, startId).Err()
		if err != nil {
			return
		}
	}
}

//判断redis库存是否添加成功

func GetProductRedis(startId int) int64 {
	id := strconv.Itoa(startId)
	return global.Rdb.LLen(global.Ctx, KEY+id).Val()
}

//扣减秒杀商品redis列表库存

func UpdateProductRedis(startId, num int64) bool {
	id := strconv.Itoa(int(startId))
	err := global.Rdb.LRem(global.Ctx, KEY+id, num, startId)
	if err != nil {
		return false
	}
	return true
}

//清除redis列表库存

func DelProductRedis(startId int) {
	id := strconv.Itoa(startId)
	global.Rdb.Del(global.Ctx, KEY+id)
}
