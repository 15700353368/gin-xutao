package utils

import (
	"context"
	"fmt"
	"gin-xutao/global"
	//"gin-xutao/params/rdl"
	"go.uber.org/zap"
	"reflect"
)


//binding type interface 要修改的结构体
//value type interace 有数据的结构体
func StructAssign(binding interface{}, value interface{}) {
	bVal := reflect.ValueOf(binding).Elem() //获取reflect.Type类型
	vVal := reflect.ValueOf(value).Elem()   //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := bVal.FieldByName(name).IsValid(); ok {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
}


//set redis
func RedisSetString(key,value string) error {
	ctx := context.Background()
	//记录redis
	err := global.GVA_REDIS.Set(ctx,key, value, 0).Err()
	fmt.Println(err)
	return err
}


//del redis
func RedisDelString(key string) error {
	ctx := context.Background()
	//记录redis
	err := global.GVA_REDIS.Del(ctx,key).Err()
	return err
}

//get redis
func RedisGetString(key string) (err error,val string) {
	ctx := context.Background()
	val, err = global.GVA_REDIS.Get(ctx,key).Result()

	if err != nil {
	global.GVA_LOG.Error("RedisStoreGetError!", zap.Error(err))
	return err,""
	}

	return err,val
}