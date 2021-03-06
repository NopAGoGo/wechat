// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/sunday9th/wechat for the canonical source repository
// @license     https://github.com/sunday9th/wechat/blob/master/LICENSE
// @authors     sunday9th(sunday9th@gmail.com)

package mp

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"net/http"
	"strconv"

	"github.com/sunday9th/wechat/internal/util"
)

// 明文模式下回复消息给微信服务器.
//  要求 msg 是有效的消息数据结构(经过 encoding/xml marshal 后符合微信消息格式);
//  如果有必要可以修改 Request 里面的某些值, 比如 Timestamp, Nonce, Random.
func WriteRawResponse(w http.ResponseWriter, r *Request, msg interface{}) (err error) {
	if w == nil {
		return errors.New("nil http.ResponseWriter")
	}
	if msg == nil {
		return errors.New("nil message")
	}
	return xml.NewEncoder(w).Encode(msg)
}

// 安全模式下回复消息的 http body
type ResponseHttpBody struct {
	XMLName struct{} `xml:"xml" json:"-"`

	EncryptedMsg string `xml:"Encrypt"      json:"Encrypt"`
	MsgSignature string `xml:"MsgSignature" json:"MsgSignature"`
	Timestamp    int64  `xml:"TimeStamp"    json:"TimeStamp"`
	Nonce        string `xml:"Nonce"        json:"Nonce"`
}

// 安全模式下回复消息给微信服务器.
//  要求 msg 是有效的消息数据结构(经过 encoding/xml marshal 后符合微信消息格式);
//  如果有必要可以修改 Request 里面的某些值, 比如 Timestamp, Nonce, Random.
func WriteAESResponse(w http.ResponseWriter, r *Request, msg interface{}) (err error) {
	if w == nil {
		return errors.New("nil http.ResponseWriter")
	}
	if r == nil {
		return errors.New("nil Request")
	}
	if msg == nil {
		return errors.New("nil message")
	}

	rawMsgXML, err := xml.Marshal(msg)
	if err != nil {
		return
	}

	encryptedMsg := util.AESEncryptMsg(r.Random, rawMsgXML, r.AppId, r.AESKey)
	base64EncryptedMsg := base64.StdEncoding.EncodeToString(encryptedMsg)

	responseHttpBody := ResponseHttpBody{
		EncryptedMsg: base64EncryptedMsg,
		Timestamp:    r.Timestamp,
		Nonce:        r.Nonce,
	}

	TimestampStr := strconv.FormatInt(responseHttpBody.Timestamp, 10)
	responseHttpBody.MsgSignature = util.MsgSign(r.Token, TimestampStr, responseHttpBody.Nonce, responseHttpBody.EncryptedMsg)

	return xml.NewEncoder(w).Encode(&responseHttpBody)
}
