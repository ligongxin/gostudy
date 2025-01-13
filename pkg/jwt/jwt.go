package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 定义过期时间
const TokenExpireDuration = time.Hour * 24 * 365

var CustomSecret = []byte("我爱学习")

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	UserID int64
	jwt.StandardClaims
}

// GenToken 生成JWT
//func GenToken(username string, userId int64) (string, error) {
//	// 创建一个我们自己的声明
//	fmt.Println(username, userId)
//	claims := CustomClaims{
//		username,
//		userId,
//		jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
//			Issuer:    "go", //签发人
//		},
//	}
//	// 使用指定的签名方法创建签名对象
//	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
//	// 使用指定的secret签名并获得完整的编码后的字符串token
//	return token.SignedString(CustomSecret)
//}

func GenToken(userId int64) (string, error) {
	// 创建一个我们自己的声明
	claims := CustomClaims{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "go", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(CustomSecret)
}

// ParseToken 解析JWT
func ParseToken(tokenStr string) (*CustomClaims, error) {
	//解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func GenTokenV1(userId int64) (aToken, rToken string, err error) {
	// 创建一个我们自己的声明
	claims := CustomClaims{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "go", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(CustomSecret)
	//refresh token 不需要任何定义
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodES256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		Issuer:    "go", // 签发人
	}).SignedString(CustomSecret)
	fmt.Println(rToken, err)
	return
}
