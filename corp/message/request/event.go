// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/sunday9th/wechat for the canonical source repository
// @license     https://github.com/sunday9th/wechat/blob/master/LICENSE
// @authors     sunday9th(sunday9th@gmail.com)

package request

import (
	"github.com/sunday9th/wechat/corp"
)

const (
	// 微信服务器推送过来的事件类型
	EventTypeSubscribe   = "subscribe"          // 订阅, 包括点击订阅和扫描二维码
	EventTypeUnsubscribe = "unsubscribe"        // 取消订阅
	EventTypeLocation    = "LOCATION"           // 上报地理位置事件
	//上报菜单事件
	EventTypeClick       = "click"              // 点击菜单拉取消息的事件推送
	EventTypeView        = "view"               // 点击菜单跳转链接的事件推送
	EventTypeScanCodeP   = "scancode_push"      // 扫码推事件的事件推送
	EventTypeScanCodeW   = "scancode_waitmsg"   // 扫码推事件且弹出“消息接收中”提示框的事件推送
	EventTypePicSysPhoto = "pic_sysphoto"       // 弹出系统拍照发图的事件推送
	EventTypePicPoA      = "pic_photo_or_album" // 弹出拍照或者相册发图的事件推送
	EventTypePicWeChat   = "pic_weixin"         // 弹出微信相册发图器的事件推送
	EventTypeLocSelect   = "location_select"    // 弹出地理位置选择器的事件推送
	EventTypeEnterAgent  = "enter_agent"        // 成员进入应用的事件推送
	EventTypeBJR         = "batch_job_result"   // 异步任务完成事件推送

)

// 关注事件
//  特别的, 默认企业小助手可以用于获取整个企业号的关注状况.
type SubscribeEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	corp.MessageHeader

	Event string `xml:"Event" json:"Event"` // 事件类型, subscribe(订阅)
}

func GetSubscribeEvent(msg *corp.MixedMessage) *SubscribeEvent {
	return &SubscribeEvent{
		MessageHeader: msg.MessageHeader,
		Event:         msg.Event,
	}
}

// 取消关注
type UnsubscribeEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	corp.MessageHeader

	Event string `xml:"Event" json:"Event"` // 事件类型, unsubscribe(取消订阅)
}

func GetUnsubscribeEvent(msg *corp.MixedMessage) *UnsubscribeEvent {
	return &UnsubscribeEvent{
		MessageHeader: msg.MessageHeader,
		Event:         msg.Event,
	}
}

// 上报地理位置事件
type LocationEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	corp.MessageHeader

	Event     string  `xml:"Event"     json:"Event"`     // 事件类型, 此时固定为: LOCATION
	Latitude  float64 `xml:"Latitude"  json:"Latitude"`  // 地理位置纬度
	Longitude float64 `xml:"Longitude" json:"Longitude"` // 地理位置经度
	Precision float64 `xml:"Precision" json:"Precision"` // 地理位置精度
}

func GetLocationEvent(msg *corp.MixedMessage) *LocationEvent {
	return &LocationEvent{
		MessageHeader: msg.MessageHeader,
		Event:         msg.Event,
		Latitude:      msg.Latitude,
		Longitude:     msg.Longitude,
		Precision:     msg.Precision,
	}
}

