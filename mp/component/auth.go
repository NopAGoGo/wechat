// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/sunday9th/wechat for the canonical source repository
// @license     https://github.com/sunday9th/wechat/blob/master/LICENSE
// @authors     sunday9th(sunday9th@gmail.com)

package component

import (
	"net/url"
)

// 微信公众号登录授权入口地址.
//  授权后跳转到 redirectURI?auth_code=xxx&expires_in=600
func AuthCodeURL(componentAppId, preAuthCode, redirectURI, state string) string {
	return "https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=" + url.QueryEscape(componentAppId) +
		"&pre_auth_code=" + url.QueryEscape(preAuthCode) +
		"&redirect_uri=" + url.QueryEscape(redirectURI) +
		"&state=" + url.QueryEscape(state)
}
