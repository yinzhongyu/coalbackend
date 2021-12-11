package middleWare

import (
	"ginVue/common"
	"ginVue/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		//验证格式
		//如果header为空，或者前缀不为..
		//我日Bearer后面还有空格
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort() //将这次请求抛弃掉
			return
		}
		tokenString = tokenString[7:] // "Bearer "七个字符
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 400,
				"msg":  "权限不足",
			})
			ctx.Abort() //将这次请求抛弃掉
			return
		}
		//验证通过后获取token的userID
		userId := claims.UserId //可以看作是一个键，mysql表里面可以看到它是gorm.model里面设置的自增主键
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId) //查出一条数据,这个函数点击进去看就是select *from user where primary_key = userId
		//用户不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 400,
				"msg":  "权限不足",
			})
			ctx.Abort() //将这次请求抛弃掉
			return
		}
		//用户存在,将user信息写入上下文
		ctx.Set("user", user) //把user数据放入user键里面
		ctx.Set("key", user.PublicKey)
		ctx.Set("kind", user.Kind)
		ctx.Next() //接收下一条请求，个人感觉应该是指针往下一个

	}
}
