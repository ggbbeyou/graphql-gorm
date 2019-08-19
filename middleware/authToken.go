package middleware

import(
	"fmt"
	"log"
	"time"
	"strings"
	"context"
	"net/http"
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/maiguangyang/graphql-gorm/gen"
)

var contxt context.Context
var secretkey = []byte("secret_key")
var user gen.User

func AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		// header := request.Header.Get("Authorization")
		header := strings.Replace(request.Header.Get("Authorization"), "Bearer ", "", 1)
		fmt.Println(header)

		if header == "" {
			next.ServeHTTP(response, request)
		} else {
			token := jwt.New(jwt.SigningMethodHS256)
			token.Claims = jwt.MapClaims{
				"userid" :   user.ID,
				"exp"    :   time.Now().Add(time.Hour * 24).Unix(),
			}
			tokenstring, err := token.SignedString(secretkey)
			if err != nil {
				log.Fatal("Error while generating token ", err)
			}
			ctxt := context.WithValue(request.Context(), "Authorization", tokenstring)
			next.ServeHTTP(response, request.WithContext(ctxt))
		}
	})
}

// 旧Token，未完成
// 1、生成token -> 记录token
// 2、验证长度 -> 检查token记录 -> 校验token -> 检查ip地址


import (
  // "fmt"
  "time"
  "errors"
  jwt "github.com/dgrijalva/jwt-go"

  Public "erp-go/common/public"
  Utils "erp-go/utils"
)

// role group
var SecretKey = context.Map {
  "user": "AdPllFsFCVlNIFyorcY0K3o1OQldYPe5",
  "admin": "btafOY5CSD3prfJM1lUSxHIJipTfe26K",
}

// 检查user的Token
func CheckAuthUser(ctx context.Context) {
  Verify(ctx.GetHeader("Authorization"), "user", ctx)
}

// 检查admin的Token
func CheckAuthAdmin(ctx context.Context) {
  Verify(ctx.GetHeader("Authorization"), "admin", ctx)
}

/**
 * 统一使用
 * 生成json web token
 */
func SetToken(str, hash interface{}, role string) string {
  timeNow := time.Now().Unix()
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "hash"      : hash,
    "content"   : str,
    "nbf"       : int64(timeNow),
    "exp"       : int64(timeNow + 60 * 60 * 24),
    "timestamp" : int64(timeNow),
  })

  ss, _ := token.SignedString([]byte(SecretKey[role].(string)))
  return ss
}

/**
 * 校验token是否有效
 */
func ParseToken(tokenStr string, key []byte) (jwt.MapClaims, error) {
  token, err := jwt.Parse(tokenStr, func (token *jwt.Token) (interface{}, error) {
    return key, nil
  })

  if err != nil {
    return nil, err
  }

  claims := token.Claims.(jwt.MapClaims)

  return claims, nil
}

/**
 * token解密
 */
func DecryptToken(tokenStr, role string) (interface{}, error) {
  claims, err := GetTokenContent(tokenStr, role)
  return claims, err
}


/**
 * 获取token内容
 */
func GetTokenContent(authHeader string, role string) (interface{}, error) {
  claims, err := ParseToken(authHeader[7:], []byte(SecretKey[role].(string)))
  return claims["content"], err
}

/**
 * 授权验证
 * 验证长度 -> 检查token记录 -> 校验token -> 检查ip地址
 */
func Verify(authHeader, role string, ctx context.Context) {
  if authHeader == "" || len(authHeader) <= 7 {
    ctx.JSON(Utils.NewResData(401, "未登录", ctx))
    return
  }

  token := authHeader[7:]

  // 校验token
  tokenData, err := ParseToken(token, []byte(SecretKey[role].(string)))
  if len(tokenData) > 0 {
    hash := Public.EncryptMd5(ctx.RemoteAddr() + SecretKey[role].(string))

    if tokenData["hash"] != hash {
      ctx.JSON(Utils.NewResData(405, "账号已在其他设备登陆", ctx))
      return
    }
  }

  if err != nil {
    ctx.JSON(Utils.NewResData(401, "登录授权已失效", ctx))
    return
  }

  ctx.Next()
}


// 获取用户JWT信息
func HandleUserJWTToken(ctx context.Context, tye string) (map[string]interface{}, error) {
  // 获取服务端用户信息
  author      := ctx.GetHeader("Authorization")
  userinfo, _ := DecryptToken(author, tye)
  reqData     := userinfo.(map[string]interface{})

  if len(reqData) <= 0 {
    return nil, errors.New("user data is empty")
  }

  if tye == "admin" && reqData["gid"] == "1" {
    reqData["gid"] = ""
  }
  return reqData, nil
}

