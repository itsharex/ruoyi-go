package session

import (
	"github.com/gin-gonic/gin"
	userModel "robvi/app/system/model/system"
)

func IsAdminUser(user *userModel.SysUser) bool {
	if user.UserId == 1 {
		return true
	} else {
		return false
	}
}

// 获得用户信息详情
func GetTenantId(c *gin.Context) int64 {
	return 0
}
