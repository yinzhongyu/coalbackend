package controller

import "github.com/gin-gonic/gin"

//
// ViewOneCoalBalance
//  @Summary 获取一个煤炭场的余量
//  @Description: 通过输入一个煤炭场的名字，来获取煤炭场的余量
//  @Tags 煤炭场数据管理
//  @Accept application/json
//  @Produce application/json
//  @Param Authorization header string false "用户令牌"
//  @Param coalName query string true "煤炭场名字"
//  @Success 200 {object} dto.CommonDto
//  @failure 400 {object} dto.CommonDto
//  @Router /coal/getOneBalance [post]
func ViewOneCoalBalance(ctx *gin.Context) {

}

//
// ViewAllCoalBalance
//  @Summary 获取所有煤炭场的煤炭余量
//  @Description: 传入token,验证权限，然后获取煤炭场的所有煤炭余量
//  @Tags 煤炭场数据管理
//  @Accept application/json
//  @Produce application/json
//  @Param Authorization header string false "用户令牌"
//  @Success 200 {object} dto.CommonDto
//  @failure 400 {object} dto.CommonDto
//  @Router /coal/getAllBalance [post]
func ViewAllCoalBalance(ctx *gin.Context) {

}

//
// ViewAllUse
//  @Summary 查看所有煤炭使用数据
//  @Description: 传入token，验证身份，查看所有煤炭使用的数据
//  @Tags 煤炭使用数据管理
//  @Accept application/json
//  @Produce application/json
//  @Param Authorization header string false "用户令牌"
//  @Success 200 {object} dto.CommonDto
//  @failure 400 {object} dto.CommonDto
//  @Router /coal/getAllUse [post]
func ViewAllUse(ctx *gin.Context) {

}

//
// ViewGenElectricityUse
//  @Summary 发电用煤数据
//  @Description: 传入token，验证身份，查看所有发电煤炭使用数据
//  @Tags 煤炭使用数据管理
//  @Accept application/json
//  @Produce application/json
//  @Param Authorization header string false "用户令牌"
//  @Success 200 {object} dto.CommonDto
//  @failure 400 {object} dto.CommonDto
//  @Router /coal/GetGenElectricityUse [post]
func ViewGenElectricityUse(ctx *gin.Context) {

}

//
// ViewOneTx
//  @Summary 查看一个煤炭场的交易记录
//  @Description: 输入token，验证身份，查看一个煤炭场的交易记录
//  @Tags 煤炭场使用数据管理
//  @Accept application/json
//  @Produce application/json
//  @Param Authorization header string false "用户令牌"
//  @Param coalName query string true "煤炭场名字"
//  @Success 200 {object} dto.CommonDto
//  @failure 400 {object} dto.CommonDto
//  @Router /coal/GetOneTx [post]
func ViewOneTx(ctx *gin.Context) {

}
