// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/sunday9th/wechat for the canonical source repository
// @license     https://github.com/sunday9th/wechat/blob/master/LICENSE
// @authors     sunday9th(sunday9th@gmail.com)

package component

import (
	"github.com/sunday9th/wechat/mp"
)

type PreAuthCode struct {
	Value     string `json:"pre_auth_code"`
	ExpiresIn int64  `json:"expires_in"`
}

// 获取预授权码.
func (clt *Client) CreatePreAuthCode() (code *PreAuthCode, err error) {
	request := struct {
		ComponentAppId string `json:"component_appid"`
	}{
		ComponentAppId: clt.AppId,
	}

	var result struct {
		mp.Error
		PreAuthCode
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	code = &result.PreAuthCode
	return
}
