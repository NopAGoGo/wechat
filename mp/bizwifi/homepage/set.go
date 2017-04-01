// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/sunday9th/wechat for the canonical source repository
// @license     https://github.com/sunday9th/wechat/blob/master/LICENSE
// @authors     sunday9th(sunday9th@gmail.com)

package homepage

import (
	"github.com/sunday9th/wechat/mp"
)

// 默认模板
func NewSetParameters1(shopId int64) interface{} {
	return &struct {
		ShopId     int64 `json:"shop_id"`
		TemplateId int64 `json:"template_id"`
	}{
		ShopId:     shopId,
		TemplateId: 0,
	}
}

// 自定义url
func NewSetParameters2(shopId int64, url string) interface{} {
	para := struct {
		ShopId     int64 `json:"shop_id"`
		TemplateId int64 `json:"template_id"`
		Struct     struct {
			URL string `json:"url"`
		} `json:"struct"`
	}{
		ShopId:     shopId,
		TemplateId: 1,
	}

	para.Struct.URL = url
	return &para
}

// 设置商家主页
//  要求 para 经过 encoding/json 后满足指定的格式要求
func Set(clt *mp.Client, para interface{}) (err error) {
	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/bizwifi/homepage/set?access_token="
	if err = clt.PostJSON(incompleteURL, para, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}
