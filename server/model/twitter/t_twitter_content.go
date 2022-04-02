package twitter

import (
	"smart-admin/server/global"
	"time"
)

// TTwitterContent 结构体
// 如果含有time.Time 请自行import time包
type TTwitterContent struct {
	global.GVA_MODEL
	Content          string     `json:"content" form:"content" gorm:"column:content;comment:内容;size:2000;"`
	ContentId        string     `json:"contentId" form:"contentId" gorm:"column:content_id;comment:推文id;size:30;"`
	ContentTypeId    string     `json:"contentTypeId" form:"contentTypeId" gorm:"column:content_type_id;comment:content_type;size:50;"`
	HashTags         string     `json:"hashTags" form:"hashTags" gorm:"column:hash_tags;comment:标签;size:20;"`
	Html             string     `json:"html" form:"html" gorm:"column:html;comment:html内容;size:2000;"`
	IsPin            *bool      `json:"isPin" form:"isPin" gorm:"column:is_pin;comment:是否置顶;"`
	IsQuoted         *bool      `json:"isQuoted" form:"isQuoted" gorm:"column:is_quoted;comment:是否引用;"`
	IsReply          *bool      `json:"isReply" form:"isReply" gorm:"column:is_reply;comment:是否答复;"`
	IsRetweet        *bool      `json:"isRetweet" form:"isRetweet" gorm:"column:is_retweet;comment:是否转载;"`
	Likes            *int       `json:"likes" form:"likes" gorm:"column:likes;comment:喜欢数;size:10;"`
	PermanentUrl     string     `json:"permanentUrl" form:"permanentUrl" gorm:"column:permanent_url;comment:永久地址;size:100;"`
	Photos           string     `json:"photos" form:"photos" gorm:"column:photos;comment:图片地址;size:200;"`
	Replies          *int       `json:"replies" form:"replies" gorm:"column:replies;comment:回复数;size:10;"`
	Retweets         *int       `json:"retweets" form:"retweets" gorm:"column:retweets;comment:转载数;size:10;"`
	SensitiveContent *bool      `json:"sensitiveContent" form:"sensitiveContent" gorm:"column:sensitive_content;comment:内容;"`
	SubTypeId        string     `json:"subTypeId" form:"subTypeId" gorm:"column:sub_type_id;comment:sub_type;size:50;"`
	TimeParsed       *time.Time `json:"timeParsed" form:"timeParsed" gorm:"column:time_parsed;comment:时间;"`
	Timestamp        *int       `json:"timestamp" form:"timestamp" gorm:"column:timestamp;comment:时间戳;size:10;"`
	Urls             string     `json:"urls" form:"urls" gorm:"column:urls;comment:地址;size:200;"`
	UserId           string     `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:20;"`
	UserName         string     `json:"userName" form:"userName" gorm:"column:user_name;comment:用户名;size:50;"`
	VideosUrl        string     `json:"videosUrl" form:"videosUrl" gorm:"column:videos_url;comment:视频地址;size:200;"`
}

// TableName TTwitterContent 表名
func (TTwitterContent) TableName() string {
	return "t_twitter_content"
}
