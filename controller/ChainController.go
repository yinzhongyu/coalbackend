package controller

import (
	"ginVue/response"
	"ginVue/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type Req struct {
	method string	`json:"method" form:"method" binding:"required"`
	args []interface{} `json:"args" form:"args" binding:"required"`
}



//
// ChainInvoke
//  @Summary 调用链码
//  @Description: 区块链链码调用
//  @Tags 链码相关
//  @Accept application/json
//  @Produce application/json
//  @Param Authorization header string false "用户令牌"
//  @Param name query string true "方法名"
//  @param args query string true "参数json"
//  @Success 200 {object} dto.CommonDto
//  @failure 400 {object} dto.CommonDto
//  @Router /chain/invoke [post]
func ChainInvoke(ctx *gin.Context){
	json:=make(map[string]interface{})
	ctx.BindJSON(&json)
	log.Println(json)
	if err:=utils.Invoke(json["method"].(string),json["args"]);err!= nil{
		response.Fail(ctx,gin.H{"error":err.Error()},"fail")
		return
	}else{
		response.Success(ctx,gin.H{}, "success")
	}
}


//
// ChainQuery
//  @Summary 调用链码
//  @Description: 区块链链码调用
//  @Tags 链码相关
//  @Accept application/json
//  @Produce application/json
//  @Param Authorization header string false "用户令牌"
//  @Param name query string true "方法名"
//  @param args query []string true "参数列表"
//  @Success 200 {object} dto.CommonDto
//  @failure 400 {object} dto.CommonDto
//  @Router /chain/query [post]
func ChainQuery(ctx *gin.Context){
	json:=make(map[string]interface{})
	ctx.BindJSON(&json)
	log.Println("获取到如下json")
	log.Println(json)
	if err,result:=utils.Query(json["method"].(string),json["args"]);err!= nil{
		response.Fail(ctx,gin.H{"error":err.Error()},"fail")
		return
	}else{
		response.Success(ctx,gin.H{"result":result}, "success")
	}
}