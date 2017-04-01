// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/sunday9th/wechat for the canonical source repository
// @license     https://github.com/sunday9th/wechat/blob/master/LICENSE
// @authors     sunday9th(sunday9th@gmail.com)

package mch

import (
	"log"
)

var LogInfoln = log.Println

// 沒有加锁, 请确保在初始化阶段调用!
func SetLogInfoln(fn func(v ...interface{})) {
	if fn == nil {
		return
	}
	LogInfoln = fn
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}
