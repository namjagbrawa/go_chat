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
	PicUrl       string `xml:"PicUrl"`
	MediaId      string `xml:"MediaId"`
	Format       string `xml:"Format"`
	Recognition  string `xml:"Recognition"`
	ThumbMediaId string `xml:"ThumbMediaId"`
	Location_X   string `xml:"Location_X"`
	Location_Y   string `xml:"Location_Y"`
	Scale        string `xml:"Scale"`
	Label        string `xml:"Label"`
	Title        string `xml:"Title"`
	Description  string `xml:"Description"`
	Url          string `xml:"Url"`
}

const (
	Text     = "text"
	image    = "image"
	voice    = "voice"
	video    = "video"
	location = "location"
	link     = "link"
)
