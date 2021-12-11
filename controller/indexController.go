package controller

import "github.com/gin-gonic/gin"

//
// TransferIndex
//  @Summary 转移指标
//  @Description: 传入起点和终点，以及份额，进行指标的转移
//  @Tags 指标管理
//  @Accept application/json
//  @Produce application/json
//  @Param Authorization header string false "用户令牌"
//  @Param from query string true "起点名字"
//  @Param to query string true "终点名字"
//  @Param num query int true "指标数量"
//  @Success 200 {object} dto.CommonDto
//  @failure 400 {object} dto.CommonDto
//  @Router /index/transferIndex [post]
func TransferIndex(ctx *gin.Context) {
}

//
// DestroyIndex
//  @Summary 删除某个煤炭场的指标
//  @Description: 删除某个煤炭场的现有指标
//  @Tags 指标管理
//  @Accept application/json
//  @Produce application/json
//  @Param Authorization header string false "用户令牌"
//  @Param coalName query string true "煤炭场名字"
//  @Param num query int true "指标数量"
//  @Success 200 {object} dto.CommonDto
//  @failure 400 {object} dto.CommonDto
//  @Router /index/DestroyIndex [post]
func DestroyIndex(ctx *gin.Context) {
}

//
// UseIndex
//  @Summary 使用指标
//  @Description: 使用现有指标
//  @Tags 指标管理
//  @Accept application/json
//  @Produce application/json
//  @Param Authorization header string false "用户令牌"
//  @Param coalName query string true "使用人名字"
//  @Param coalName query string true "使用方法"
//  @Param num query int true "指标数量"
//  @Success 200 {object} dto.CommonDto
//  @failure 400 {object} dto.CommonDto
//  @Router /index/UseIndex [post]
func UseIndex(ctx *gin.Context) {
}
