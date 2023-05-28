package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/thoas/go-funk"
	"net/http"
	"simple-mall/global"
	"simple-mall/models"
	"simple-mall/models/user"
	"strings"
	"time"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.JWTConfig.SigningKey), //可以设置过期时间
	}
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims models.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*models.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now

		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(1 * time.Hour))
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

// JWTAuth token校验
func JWTAuth(whitelist []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		curRoutePath := c.Request.URL.Path
		// 白名单路由内不验证token
		isWhitelistRouting := funk.Contains(whitelist, curRoutePath)
		// swagger 路由不验证token
		isSwaggerRoute := strings.Contains(curRoutePath, "/swagger/")

		if isWhitelistRouting || isSwaggerRoute {
			c.Next()
			return
		}

		// 我们这里jwt鉴权取头部信息 authorization 登录时回返回token信息
		token := c.Request.Header.Get("authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token不存在！",
			})
			c.Abort()
			return
		}
		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				if err == TokenExpired {
					c.JSON(http.StatusUnauthorized, gin.H{
						"code": 401,
						"msg":  "授权已过期！",
					})
					c.Abort()
					return
				}
			}

			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "未登陆！",
			})
			c.Abort()
			return
		}
		c.Set("userId", claims.ID)
		c.Set("authorityId", claims.AuthorityId)
		c.Next()
	}
}

// GenerateLoginToken 生成登陆token
func GenerateLoginToken(user user.User) (token string, err error) {
	j := NewJWT()
	claims := models.CustomClaims{
		ID:          user.ID,
		NickName:    user.NickName,
		AuthorityId: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 签名的生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24小时
			Issuer:    "1057",
		},
	}
	return j.CreateToken(claims)
}
