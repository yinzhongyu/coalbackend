package controller

import (
	"github.com/gin-gonic/gin"
)

//
// RegisterCer
//  @Summary 注册证书
//  @Description: 煤炭场对证书进行申请
//  @Tags 证书管理
//  @Accept application/json
//  @Produce application/json
//  @Param Authorization header string false "用户令牌"
//  @Success 200 {object} dto.CommonDto
//  @failure 400 {object} dto.CommonDto
//  @Router /cer/registerCer [post]
func RegisterCer(ctx *gin.Context) {
	//获得公钥
	//publickey := ctx.Get("key")
	//查询此公钥是否已经有证书了
	//utils.Query()
	//若没有 则调用注册合约
	//返回结果信息
}

//
// ApprovalCer
//  @Summary 审批证书
//  @Description: 政府部门对证书进行审批
//  @Tags 证书管理
//  @Accept application/json
//  @Produce application/json
//  @Param Authorization header string false "用户令牌"
//  @Param coalPublicKey query string true "煤炭场公钥"
//  @Param coalName query string true "秘钥"
//  @Success 200 {object} dto.CommonDto
//  @failure 400 {object} dto.CommonDto
//  @Router /cer/approvalCer [post]
func ApprovalCer(ctx *gin.Context) {
	//获取参数

	//验证权限

	//通过煤炭场公钥获取证书

	//使用秘钥对证书签名

	//返回结果
}

//
// RevokeCer
//  @Summary 撤销证书
//  @Description: 指定一个煤炭场的名字，撤销此证书
//  @Tags 接口类别
//  @Accept application/json
//  @Produce application/json
//  @Param Authorization header string false "用户令牌"
//  @Param coalPublicKey query string true "煤炭场公钥"
//  @Success 200 {object} dto.CommonDto
//  @failure 400 {object} dto.CommonDto
//  @Router /cer/revokeCer [post]
func RevokeCer(ctx *gin.Context) {
	//获取参数
	//验证权限
	//通过煤炭场公钥撤销证书
	//返回结果
}

//
// ViewOneCer
//  @Summary 获取一个证书
//  @Description: 获取一个证书
//  @Tags 接口类别
//  @Accept application/json
//  @Produce application/json
//  @Param Authorization header string false "用户令牌"
//  @Param coalPublicKey query string true "煤炭场公钥"
//  @Success 200 {object} dto.CommonDto
//  @failure 400 {object} dto.CommonDto
//  @Router /cer/GetOneCer [post]
func ViewOneCer(ctx *gin.Context) {
	//获取参数
	//验证权限
	//通过煤炭场公钥获取证书信息

	//返回证书信息
}

//
// ViewAllCer
//  @Summary 获取所有证书
//  @Description: 获取所有煤炭场证书，包括已经审批的和未审批的
//  @Tags 接口类别
//  @Accept application/json
//  @Produce application/json
//  @Param Authorization header string false "用户令牌"
//  @Success 200 {object} dto.CommonDto
//  @failure 400 {object} dto.CommonDto
//  @Router /cer/GetAllCer [post]
func ViewAllCer(ctx *gin.Context) {
	//获取参数
	//验证权限
	//获取所有证书

}
