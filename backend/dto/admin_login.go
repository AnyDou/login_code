package dto

import (
	"github.com/e421083458/go_gateway_demo/public"
	"github.com/gin-gonic/gin"
	"time"
)

type AdminSessionInfo struct {
	ID 			int `json:"id"`
	UserName 	string `json:"username"`
	LoginTime 	time.Time `json:"login_time"`
}

//定义结构体
//input为输入,output输出
type AdminLoginInput struct{
	//json是输出时运用，form是输入时运用
	//example是swagger文档的默认值
	//comment是错误输出时使用设置值
	//validate检验是否为required,is_valid_username定义验证器
	UserName string `json:"username" form:"username" comment:"用户名" example:"admin" validate:"required,is_valid_username"`
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"`
}

//*AdminLoginInput即AdminLoginInput指针
//param *AdminLoginInput即AdminLoginInput指针定义为param
//BindValidParam绑定结构体并检验参数
//error绑定不成功抛出error
func (param *AdminLoginInput) BindValidParam (c *gin.Context) error{
	//检验
	return public.DefaultGetValidParams(c, param)
}

type AdminLoginOutput struct{
	//给前端传token
	Token string `json:"token" form:"token" comment:"token" example:"token" validate:""`
}