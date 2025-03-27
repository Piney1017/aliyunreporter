package uredis

import (
	"context"
	"strconv"
	"time"

	redis "github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	logger "github.com/sirupsen/logrus"
)

var rdb *redis.Client

const (
	MilliSecondsInDay = 60 * 60 * 24 * 1000
)

func Open(url string) {
	opt, err := redis.ParseURL(url)
	if err != nil {
		logger.WithContext(context.Background()).Errorf("redis.ParseURL url: %v Failed: %+v\n", url, err)
		panic("redis.ParseURL Failed")
	}

	rdb = redis.NewClient(opt)

	for {
		e := rdb.Ping(context.Background()).Err()
		if e == nil {
			break
		}
		logger.WithContext(context.Background()).Errorf("Connect To Redis Failed: %+v\n", e)
		time.Sleep(time.Second)
	}

	logger.WithContext(context.Background()).Infoln("Connected to Redis Success")
}

func Close() {
	rdb.Close()
	rdb = nil
}

func Set(key string, value interface{}, exp time.Duration) {
	rdb.Set(context.Background(), key, value, exp)
}

func Del(key string) {

	rdb.Del(context.Background(), key)
}

func Get(key string) interface{} {
	result, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		return nil
	}
	return result
}

func Exists(key string) bool {
	cmd := rdb.Exists(context.Background(), key)
	return cmd.Val() > 0
}

func GetInt64(key string) int64 {
	result, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		return 0
	}
	r, err := strconv.ParseInt(result, 10, 64)
	if err != nil {
		return 0
	}
	return r
}

func HIncr(key string, field string, diff int64) (int64, error) {
	r := rdb.HIncrBy(context.Background(), key, field, diff)
	return r.Result()
}

func HGetAll(key string) (map[string]int64, error) {

	r := rdb.HGetAll(context.Background(), key)

	val, err := r.Result()
	if err != nil {
		return nil, err
	}

	rtv := lo.MapValues(val, func(v, k string) int64 {
		i, _ := strconv.ParseInt(v, 10, 64)
		return i
	})

	return rtv, nil
}

func Expire(key string, exp time.Duration) {
	rdb.Expire(context.Background(), key, exp)
}

func AddSet(key string, value string) (int64, error) {
	r := rdb.SAdd(context.Background(), key, value)
	return r.Result()
}

func RemoveSet(key string, value string) {
	rdb.SRem(context.Background(), key, value)
}

func MembersSet(key string) []string {
	result, _ := rdb.SMembers(context.Background(), key).Result()
	return result
}

func MatchKeys(key string) []string {
	result, _ := rdb.Keys(context.Background(), key).Result()
	return result
}

func ZAdd(key string, member redis.Z) {
	rdb.ZAdd(context.Background(), key, member)
}

func ZRem(key string, member redis.Z) {
	rdb.ZRem(context.Background(), key, member)
}

func ZRemRangeByRank(key string, start, end int64) {
	rdb.ZRemRangeByRank(context.Background(), key, start, end)
}

func ZRang(key string, start, end int64) *redis.StringSliceCmd {
	list := rdb.ZRange(context.Background(), key, start, end)
	return list
}

func ZRevRange(key string, start, end int64) *redis.StringSliceCmd {
	list := rdb.ZRevRange(context.Background(), key, start, end)
	return list
}

func ZRangeByScore(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	return rdb.ZRangeByScore(context.Background(), key, opt)
}

func ZCard(key string) *redis.IntCmd {
	return rdb.ZCard(context.Background(), key)
}
