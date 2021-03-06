// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/sunday9th/wechat for the canonical source repository
// @license     https://github.com/sunday9th/wechat/blob/master/LICENSE
// @authors     sunday9th(sunday9th@gmail.com)

package mmpaymkttransfers

import (
	"github.com/sunday9th/wechat/mch"
)

// 红包查询接口.
//  NOTE: 请求需要双向证书
func GetRedPackInfo(pxy *mch.Proxy, req map[string]string) (resp map[string]string, err error) {
	return pxy.PostXML("https://api.mch.weixin.qq.com/mmpaymkttransfers/gethbinfo", req)
}
