// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/sunday9th/wechat for the canonical source repository
// @license     https://github.com/sunday9th/wechat/blob/master/LICENSE
// @authors     sunday9th(sunday9th@gmail.com)

package code

import (
	"github.com/sunday9th/wechat/mp"
)

// Code解码接口
func Decrypt(clt *mp.Client, encryptCode string) (code string, err error) {
	request := struct {
		EncryptCode string `json:"encrypt_code"`
	}{
		EncryptCode: encryptCode,
	}

	var result struct {
		mp.Error
		Code string `json:"code"`
	}

	incompleteURL := "https://api.weixin.qq.com/card/code/decrypt?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	code = result.Code
	return
}
