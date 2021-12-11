package common

import (
	"ginVue/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// var jwtKey = sha256.Sum256([]byte("this_is_a_secret_key"))
var jwtKey = []byte("dGhpc19pc19hX3NlY3JldF9rZXk=")

// 具体业务数据
type Claims struct {
	UserId    uint
	Kind      int
	PublicKey string
	jwt.StandardClaims
}

//登陆成功后发放token
func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(100 * 24 * time.Hour) //token有效期为100天
	claims := &Claims{
		UserId:    user.ID,
		PublicKey: user.PublicKey,
		Kind:      user.Kind,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),     //发放时间
			Issuer:    "The Government",      //发放者
			Subject:   "user token",          //主题
		},
	}
	// 使用SHA256进行签名,将头部和载荷组合  Header.Payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey) // 使用密钥生成token
	if err != nil {
		return "", err
	}
	return tokenString, nil
	/**
	会生成三部分 xxx.yyy.zzz
	header:使用的加密算法 (可用base64解密)
	payload:claims内容加密
	zzz:前面两部分hash加密的值
	*/
}

//从tokenString解析出claims
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	//解析tokenString，将内容放到toekn、claims里面，最主要的是获取claims的内容
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
