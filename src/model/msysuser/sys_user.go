/*
create by subingjie
date:2015-09-23 10:14:00
discription:系统用户模块
*/

package msysuser

import "gorm.io/gorm"

type SysUser struct {
	gorm.Model
	Username     string `gorm:"type:varchar(20);not null;uniqueIndex;comment:用户名"`
	Password     string `gorm:"type:varchar(20);default:'';comment:密码"`
	Email        string `gorm:"type:varchar(50);default:'';comment:邮箱"`
	Phone        string `gorm:"type:varchar(20);not null;comment:手机号"`
	Avatar       string `gorm:"type:varchar(100);default:'';comment:头像"`
	WechatOpenid string `gorm:"type:varchar(50);default:'';comment:微信openid"`
}

type SysUserModel struct {
}

func (s *SysUserModel) Create(user *SysUser) (uint, error) {

	return 0, nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func longestPalindrome(s string) int {
	mp := map[byte]int{}

	for i := 0; i < len(s); i++ {
		mp[s[i]]++
	}

	res := 0
	for _, v := range mp {
		if v&1 == 1 {
			res += v - 1
		} else {
			res += v
		}
	}

	if res < len(s) {
		res++
	}

	return res
}
