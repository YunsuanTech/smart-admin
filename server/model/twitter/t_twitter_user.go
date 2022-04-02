package twitter

import (
	"smart-admin/server/global"
	"time"
)

// TTwitterUser 结构体
// 如果含有time.Time 请自行import time包
type TTwitterUser struct {
	global.GVA_MODEL
	AvatarUrl      string     `json:"avatarUrl" form:"avatarUrl" gorm:"column:avatar_url;comment:用户头像url;size:100;"`
	BannerUrl      string     `json:"bannerUrl" form:"bannerUrl" gorm:"column:banner_url;comment:背景url;size:100;"`
	Biography      string     `json:"biography" form:"biography" gorm:"column:biography;comment:个人介绍;size:255;"`
	Birthday       string     `json:"birthday" form:"birthday" gorm:"column:birthday;comment:出生日期;size:20;"`
	FollowersCount *int       `json:"followersCount" form:"followersCount" gorm:"column:followers_count;comment:粉丝数;size:10;"`
	FollowingCount *int       `json:"followingCount" form:"followingCount" gorm:"column:following_count;comment:追随数;size:10;"`
	FriendsCount   *int       `json:"friendsCount" form:"friendsCount" gorm:"column:friends_count;comment:好友数;size:10;"`
	IsPrivate      *bool      `json:"isPrivate" form:"isPrivate" gorm:"column:is_private;comment:是否私有;"`
	IsVerified     *bool      `json:"isVerified" form:"isVerified" gorm:"column:is_verified;comment:是否校验;"`
	JoinedTime     *time.Time `json:"joinedTime" form:"joinedTime" gorm:"column:joined_time;comment:加入日期;"`
	LikesCount     *int       `json:"likesCount" form:"likesCount" gorm:"column:likes_count;comment:喜欢数;size:10;"`
	ListedCount    *int       `json:"listedCount" form:"listedCount" gorm:"column:listed_count;comment:;size:10;"`
	Location       string     `json:"location" form:"location" gorm:"column:location;comment:;size:100;"`
	TweetsCount    *int       `json:"tweetsCount" form:"tweetsCount" gorm:"column:tweets_count;comment:推文数;size:10;"`
	UserId         string     `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:20;"`
	UserName       string     `json:"userName" form:"userName" gorm:"column:user_name;comment:用户名称;size:50;"`
	UserTypeId     string     `json:"userTypeId" form:"userTypeId" gorm:"column:user_type_id;comment:外键id;size:50;"`
	WebsiteUrl     string     `json:"websiteUrl" form:"websiteUrl" gorm:"column:website_url;comment:个人推特地址;size:100;"`
}

// TableName TTwitterUser 表名
func (TTwitterUser) TableName() string {
	return "t_twitter_user"
}
