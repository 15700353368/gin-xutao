package middleware

import (
	"crypto/md5"
	"fmt"
	"gin-xutao/global"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math"
	"strconv"
	"time"
)

func RestGate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取参数加签
		sign 		:= c.Request.Header.Get("Sign")
		timeStamp 	:= c.Request.Header.Get("Time-Stamp")
		method 		:= c.Request.Method

		//调试不用验证
		if sign == "debug"{
			c.Next()
			return
		}

		//客户端时间戳转换
		clientTime,_ := strconv.ParseInt(timeStamp,10,64)

		//这两个参数不能为空 前后端时间戳相差不超过120s
		ucsTimeStamp := global.GVA_CONFIG.System.UcsTimeStamp

		if sign == "" || timeStamp == "" || (math.Abs(float64(clientTime - time.Now().Unix()))   > float64(ucsTimeStamp) ){
			c.JSON(500,gin.H{"mesg":"签名错误,或者时间戳过期"})
			c.Abort()
			return
		}

		//开始验证签名
		//请求字符串
		var paramsStr string
		if method == "POST"{
			paramsBody,_ 	:= ioutil.ReadAll(c.Request.Body)
			paramsStr = string(paramsBody)
		}

		if method == "GET"{
			paramsStr = c.Request.URL.RawQuery
		}


		//md5加密
		postSighStr 	:= paramsStr  + timeStamp + global.GVA_CONFIG.System.UcsKey
		data := []byte(string(postSighStr))
		has := md5.Sum(data)
		md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
		fmt.Println(md5str1)
		if md5str1 != sign{
			c.JSON(500,gin.H{"mesg":"签名错误,或者时间戳过期"})
			c.Abort()
			return
		}

		c.Next()
	}

}

