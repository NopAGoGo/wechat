// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/sunday9th/wechat for the canonical source repository
// @license     https://github.com/sunday9th/wechat/blob/master/LICENSE
// @authors     sunday9th(sunday9th@gmail.com)

package bizwifi

import (
	"github.com/sunday9th/wechat/mp"
)

const (
	// 推送到公众号URL上的事件类型
	EventTypeWifiConnected = "WifiConnected" // Wi-Fi连网成功事件
)

type WifiConnectedEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.MessageHeader

	Event string `xml:"Event" json:"Event"` // 事件类型，WifiConnected (Wi-Fi连网成功)

	ConnectTime int64  `xml:"ConnectTime" json:"ConnectTime"` // 连网时间（整型）
	ExpireTime  int64  `xml:"ExpireTime"  json:"ExpireTime"`  // 系统保留字段，固定值
	VendorId    string `xml:"VendorId"    json:"VendorId"`    // 系统保留字段，固定值
	PlaceId     int64  `xml:"PlaceId"     json:"PlaceId"`     // 连网的门店id
	DeviceNo    string `xml:"DeviceNo"    json:"DeviceNo"`    // 连网的设备无线mac地址，对应bssid
}

func GetWifiConnectedEvent(msg *mp.MixedMessage) *WifiConnectedEvent {
	return &WifiConnectedEvent{
		MessageHeader: msg.MessageHeader,
		Event:         msg.Event,
		ConnectTime:   msg.ConnectTime,
		ExpireTime:    msg.ExpireTime,
		VendorId:      msg.VendorId,
		PlaceId:       msg.PlaceId,
		DeviceNo:      msg.DeviceNo,
	}
}
