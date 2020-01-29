package message

import "encoding/xml"

type Message struct {
	XMLName      xml.Name `xml:"xml"` // 指定最外层的标签为config
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int      `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	MsgId        uint64   `xml:"MsgId"`

	Content      string `xml:"Content"`
	PicUrl       uint64 `xml:"PicUrl"`
	MediaId      uint64 `xml:"MediaId"`
	Format       uint64 `xml:"Format"`
	Recognition  uint64 `xml:"Recognition"`
	ThumbMediaId uint64 `xml:"ThumbMediaId"`
	Location_X   uint64 `xml:"Location_X"`
	Location_Y   uint64 `xml:"Location_Y"`
	Scale        uint64 `xml:"Scale"`
	Label        uint64 `xml:"Label"`
	Title        uint64 `xml:"Title"`
	Description  uint64 `xml:"Description"`
	Url          uint64 `xml:"Url"`
}

const (
	Text     = "text"
	image    = "image"
	voice    = "voice"
	video    = "video"
	location = "location"
	link     = "link"
)
