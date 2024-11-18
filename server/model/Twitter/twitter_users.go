// 自动生成模板TwitterUsers
package Twitter

import (
	"time"
)

// twitterUsers表 结构体  TwitterUsers
type TwitterUsers struct {
	Id              *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;"`                               //ID
	UserId          *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:;"`                             //用户唯一ID
	Username        string     `json:"username" form:"username" gorm:"column:username;comment:;"`                        //用户名
	DisplayName     *string    `json:"displayName" form:"displayName" gorm:"column:display_name;comment:;"`              //显示名称
	Bio             *string    `json:"bio" form:"bio" gorm:"column:bio;comment:;"`                                       //用户简介
	Location        *string    `json:"location" form:"location" gorm:"column:location;comment:;"`                        //位置信息
	Website         *string    `json:"website" form:"website" gorm:"column:website;comment:;"`                           //网站链接
	JoinDate        *string    `json:"joinDate" form:"joinDate" gorm:"column:join_date;comment:;"`                       //注册日期
	FollowersCount  *int       `json:"followersCount" form:"followersCount" gorm:"column:followers_count;comment:;"`     //关注者数量
	FollowingCount  *int       `json:"followingCount" form:"followingCount" gorm:"column:following_count;comment:;"`     //关注数量
	TweetCount      *int       `json:"tweetCount" form:"tweetCount" gorm:"column:tweet_count;comment:;"`                 //总推文数
	LikesCount      *int       `json:"likesCount" form:"likesCount" gorm:"column:likes_count;comment:;"`                 //点赞数量
	ProfileImageUrl string     `json:"profileImageUrl" form:"profileImageUrl" gorm:"column:profile_image_url;comment:;"` //头像URL
	CoverImageUrl   string     `json:"coverImageUrl" form:"coverImageUrl" gorm:"column:cover_image_url;comment:;"`       //封面图片URL
	LastScraped     *time.Time `json:"lastScraped" form:"lastScraped" gorm:"column:last_scraped;comment:;"`              //最后一次爬取时间
	Weight          *float64   `json:"weight" form:"weight" gorm:"column:weight;comment:;"`                              //权重
}

// TableName twitterUsers表 TwitterUsers自定义表名 twitter_users
func (TwitterUsers) TableName() string {
	return "twitter_users"
}
