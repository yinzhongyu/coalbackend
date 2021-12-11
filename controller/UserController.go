package controller

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"ginVue/common"
	"ginVue/dto"
	"ginVue/model"
	"ginVue/response"
	"ginVue/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

//
// Register
//  @Summary 用户注册
//  @Description: 用户注册接口
//  @Tags 用户管理
//  @Accept application/json
//  @Produce application/json
//  @Param name query string true "用户名"
//  @Param telephone query string true "电话号码"
//  @Param password query string true "密码"
//  @Param kind query int true "权限级别 2.煤炭场 3.普通用户"
//  @Success 200 {object} dto.CommonDto
//  @failure 400 {object} dto.CommonDto
//  @Router /auth/register [post]
func Register(ctx *gin.Context) {
	DB := common.GetDB()
	// 单纯的下面三行获取不到html传参
	//name := ctx.PostForm("name")
	//telephone := ctx.PostForm("telephone")
	//password := ctx.PostForm("password")

	//方法1
	//var requsetMap = make(map[string]string)
	//json.NewDecoder(ctx.Request.Body).Decode(requsetMap) //将传参的body解析到map

	//方法2
	//var requestUser = model.User{}
	//json.NewDecoder(ctx.Request.Body).Decode(requestUser)

	//方法3 gin提供的bind
	var requestUser = model.User{}
	ctx.ShouldBind(&requestUser)
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	fmt.Println(password)
	fmt.Println()
	fmt.Println()
	if len(telephone) != 11 {
		//ctx.JSON(http.StatusUnprocessableEntity,gin.H{
		//	"code":422,
		//	"msg":"手机号长度必须为11位",
		//})
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号长度必须为11位")
		return
	}
	if len(password) < 3 {
		//ctx.JSON(http.StatusUnprocessableEntity,gin.H{
		//	"code":422,
		//	"msg":"密码长度不小于3位",
		//})
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码长度不小于3位")
		return
	}
	if len(name) == 0 {
		name = utils.RandomString(13)
	}
	log.Println(name, telephone, password, requestUser.Kind)
	if utils.TelExists(DB, telephone) {
		//ctx.JSON(http.StatusUnprocessableEntity,gin.H{
		//	"code":422,
		//	"msg":"用户已经存在",
		//})
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}
	//创建用户，密码加密
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		//ctx.JSON(http.StatusUnprocessableEntity,gin.H{
		//	"code":500,
		//	"msg":"加密错误",
		//})
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密错误")
		return
	}

	//获取公钥
	_,publickey:=genRsaKey()


	user := model.User{
		Telephone: telephone,
		Name:      name,
		Password:  string(hashPwd),
		PublicKey: string(publickey),
	}
	DB.Create(&user)
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		log.Printf("token generates error :%v", err)
		return
	}
	response.Success(ctx, gin.H{"token": token,"publickey":user.PublicKey}, "注册成功")
}

//
// Login
//  @Summary 用户登录
//  @Description: 用户登录，获取token
//  @Tags 用户管理
//  @Accept application/json
//  @Produce application/json
//  @Param telephone query string true "手机号"
//  @Param password query string true "密码"
//  @Success 200 {object} dto.CommonDto
//  @failure 400 {object} dto.CommonDto
//  @Router /auth/login [post]
func Login(ctx *gin.Context) {
	DB := common.GetDB()
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")


	fmt.Println(telephone)
	fmt.Println(password)
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号长度必须为11位",
		})
		return
	}
	if len(password) < 3 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码长度不小于3位",
		})
		return
	}
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户不存在",
		})
		return
	}
	//判断密码(用户密码不能明文保存)
	//arg1:查询的密码，arg2:传参的密码
	fmt.Println()
	fmt.Println()
	fmt.Println(user.Password)
	fmt.Println(password)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 400,
			"msg":  "密码错误",
		})
		return
	}

	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		log.Printf("token generates error :%v", err)
		return
	}
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

//
// Info
//  @Summary 获取个人信息
//  @Description: 根据token获取用户个人信息
//  @Tags 用户管理
//  @Accept application/json
//  @Produce application/json
//  @Param Authorization header string false "用户令牌"
//  @Success 200 {object} dto.CommonDto
//  @failure 400 {object} dto.CommonDto
//  @Router /auth/info [post]
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user") //直接从上下文获取用户信息
	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		"code": 200,
		"user": dto.ToUserDto(user.(model.User)), //类型断言r
	})
}


//RSA公钥私钥产生
func genRsaKey() (prvkey, pubkey []byte) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)

	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	prvkey = pem.EncodeToMemory(block)
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	//pubkey,_=hex.DecodeString(string(derPkix))


	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Bytes: derPkix,
	}
	pubkey = pem.EncodeToMemory(block)

	return
}