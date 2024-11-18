package Twitter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TwitterUsersRouter struct {}

// InitTwitterUsersRouter 初始化 twitterUsers表 路由信息
func (s *TwitterUsersRouter) InitTwitterUsersRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	twitterUsersRouter := Router.Group("twitterUsers").Use(middleware.OperationRecord())
	twitterUsersRouterWithoutRecord := Router.Group("twitterUsers")
	twitterUsersRouterWithoutAuth := PublicRouter.Group("twitterUsers")
	{
		twitterUsersRouter.POST("createTwitterUsers", twitterUsersApi.CreateTwitterUsers)   // 新建twitterUsers表
		twitterUsersRouter.DELETE("deleteTwitterUsers", twitterUsersApi.DeleteTwitterUsers) // 删除twitterUsers表
		twitterUsersRouter.DELETE("deleteTwitterUsersByIds", twitterUsersApi.DeleteTwitterUsersByIds) // 批量删除twitterUsers表
		twitterUsersRouter.PUT("updateTwitterUsers", twitterUsersApi.UpdateTwitterUsers)    // 更新twitterUsers表
	}
	{
		twitterUsersRouterWithoutRecord.GET("findTwitterUsers", twitterUsersApi.FindTwitterUsers)        // 根据ID获取twitterUsers表
		twitterUsersRouterWithoutRecord.GET("getTwitterUsersList", twitterUsersApi.GetTwitterUsersList)  // 获取twitterUsers表列表
	}
	{
	    twitterUsersRouterWithoutAuth.GET("getTwitterUsersPublic", twitterUsersApi.GetTwitterUsersPublic)  // twitterUsers表开放接口
	}
}
