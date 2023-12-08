package routes

import (
	"github.com/gin-gonic/gin"
	"interastral-peace.com/alnitak/internal/api/v1"
)

func CollectAuthRoutes(r *gin.RouterGroup) {
	authApi := api.UserAuthApi{}

	// 用户注册
	r.POST("auth/register", authApi.Register)
	// 用户登录(密码)
	r.POST("auth/login", authApi.Login)
	// 更新token
	r.POST("auth/updateToken", authApi.UpdateToken)
}
