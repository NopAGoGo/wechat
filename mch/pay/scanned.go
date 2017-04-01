// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/sunday9th/wechat for the canonical source repository
// @license     https://github.com/sunday9th/wechat/blob/master/LICENSE
// @authors     sunday9th(sunday9th@gmail.com)

package pay

import (
	"github.com/sunday9th/wechat/mch"
)

// 提交刷卡支付.
func MicroPay(pxy *mch.Proxy, req map[string]string) (resp map[string]string, err error) {
	return pxy.PostXML("https://api.mch.weixin.qq.com/pay/micropay", req)
}

// 撤销订单.
//  NOTE: 请求需要双向证书.
func Reverse(pxy *mch.Proxy, req map[string]string) (resp map[string]string, err error) {
	return pxy.PostXML("https://api.mch.weixin.qq.com/secapi/pay/reverse", req)
}
