package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	tokenExpireDuration = time.Hour * 24
)

var (
	JwtSecret = []byte("勋")
)

type UserClaims struct {
	// 可根据需要自行添加字段
	Username   string `json:"username"`
	Permission string `json:"permission"`
}

type Claims struct {
	UserClaims
	jwt.RegisteredClaims // 内嵌标准的声明
}

// GenerateToken 生成token
func GenerateToken(username string, permission string) (string, error) {
	// UserClaims是自定义存储在token中的信息
	user := UserClaims{
		Username:   username,
		Permission: permission,
	}
	// 创建一个我们自己的声明
	claims := Claims{
		user, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpireDuration)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                          // 签发时间
			Issuer:    "blog",                                                  // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(JwtSecret)
}

// ParseToken 解析JWT  主要干两件事：1.将请求体中的token进行解析； 2.解析完成后将其断言成Claims类型，据此拿到token中自定义携带的个人信息
func ParseToken(tokenString string) (*Claims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*Claims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
