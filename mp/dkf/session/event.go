// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/sunday9th/wechat for the canonical source repository
// @license     https://github.com/sunday9th/wechat/blob/master/LICENSE
// @authors     sunday9th(sunday9th@gmail.com)

package session

import (
	"github.com/sunday9th/wechat/mp"
)

const (
	EventTypeKfCreateSession = "kf_create_session" // 接入会话
	EventTypeKfCloseSession  = "kf_close_session"  // 关闭会话
	EventTypeKfSwitchSession = "kf_switch_session" // 转接会话
)

type KfCreateSessionEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.MessageHeader

	Event     string `xml:"Event"     json:"Event"`
	KfAccount string `xml:"KfAccount" json:"KfAccount"`
}

func GetKfCreateSessionEvent(msg *mp.MixedMessage) *KfCreateSessionEvent {
	return &KfCreateSessionEvent{
		MessageHeader: msg.MessageHeader,
		Event:         msg.Event,
		KfAccount:     msg.KfAccount,
	}
}

type KfCloseSessionEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.MessageHeader

	Event     string `xml:"Event"     json:"Event"`
	KfAccount string `xml:"KfAccount" json:"KfAccount"`
}

func GetKfCloseSessionEvent(msg *mp.MixedMessage) *KfCloseSessionEvent {
	return &KfCloseSessionEvent{
		MessageHeader: msg.MessageHeader,
		Event:         msg.Event,
		KfAccount:     msg.KfAccount,
	}
}

type KfSwitchSessionEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.MessageHeader

	Event         string `xml:"Event"         json:"Event"`
	FromKfAccount string `xml:"FromKfAccount" json:"FromKfAccount"`
	ToKfAccount   string `xml:"ToKfAccount"   json:"ToKfAccount"`
}

func GetKfSwitchSessionEvent(msg *mp.MixedMessage) *KfSwitchSessionEvent {
	return &KfSwitchSessionEvent{
		MessageHeader: msg.MessageHeader,
		Event:         msg.Event,
		FromKfAccount: msg.FromKfAccount,
		ToKfAccount:   msg.ToKfAccount,
	}
}
