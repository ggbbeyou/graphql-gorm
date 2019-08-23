package middleware

import(
	// "fmt"
	// "log"
	"time"
  "errors"
	"strings"
	"context"
	"net/http"
	jwt "github.com/dgrijalva/jwt-go"

	// "github.com/maiguangyang/graphql-gorm/gen"
  "github.com/maiguangyang/graphql-gorm/utils"
)

// var user gen.User

func AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, req *http.Request) {
		// header := req.Header.Get("Authorization")
		// header := strings.Replace(req.Header.Get("Authorization"), "Bearer ", "", 1)

		// if header == "" {
		// 	next.ServeHTTP(response, req)
		// } else {

			// token := jwt.New(jwt.SigningMethodHS256)
			// token.Claims = jwt.MapClaims{
			// 	"userid" :   user.ID,
			// 	"exp"    :   time.Now().Add(time.Hour * 24).Unix(),
			// }
			// tokenstring, err := token.SignedString(secretkey)
			// if err != nil {
			// 	log.Fatal("Error while generating token ", err)
			// }
			// ctxt := context.WithValue(req.Context(), "Authorization", tokenstring)
			// next.ServeHTTP(response, req.WithContext(ctxt))
		// }

    // 返回前端的Token

    ip := utils.RemoteIp(req)
    // token := SetToken(map[string]interface{}{
    //   "id": "998fc3fb-59c5-4af9-86b4-987cb14363f1",
    // }, utils.EncryptMd5(ip + SecretKey["admin"].(string)), "admin")

    // fmt.Println(token)

    res, _ := HandleUserJWTToken(req, "admin")
    ctxt := context.WithValue(req.Context(), "Authorization", res)
    ctxt = context.WithValue(req.Context(), "RemoteIp", ip)
    next.ServeHTTP(response, req.WithContext(ctxt))
	})
}

// 旧Token，未完成
// 1、生成token -> 记录token
// 2、验证长度 -> 检查token记录 -> 校验token -> 检查ip地址

// role group
var SecretKey = map[string]interface{}{
  "user": "AdPllFsFCVlNIFyorcY0K3o1OQldYPe5",
  "admin": "btafOY5CSD3prfJM1lUSxHIJipTfe26K",
}

// 检查user的Token
func CheckAuthUser(req *http.Request) error {
  return Verify(req, "user")
}

// 检查admin的Token
func CheckAuthAdmin(req *http.Request) error {
  return Verify(req, "admin")
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
func Verify(req *http.Request, role string) error {
  token := strings.Replace(req.Header.Get("Authorization"), "Bearer ", "", 1)

  if token == "" {
    return errors.New("未登录")
  }

  // 校验token
  tokenData, err := ParseToken(token, []byte(SecretKey[role].(string)))
  if len(tokenData) > 0 {
    hash := utils.EncryptMd5(utils.RemoteIp(req) + SecretKey[role].(string))

    if tokenData["hash"] != hash {
      return errors.New("账号已在其他设备登陆")
    }
  }

  if err != nil {
    return errors.New("登录授权已失效")
  }

  return nil
}


// 获取用户JWT信息
func HandleUserJWTToken(req *http.Request, role string) (map[string]interface{}, error) {
  err := CheckAuthAdmin(req)

  if err != nil {
    return nil, err
  }

  // 获取服务端用户信息
  author  := req.Header.Get("Authorization")
  userinfo, err := DecryptToken(author, role)

  if userinfo == nil || err != nil {
    return nil, errors.New("user data is empty")
  }

  reqData := userinfo.(map[string]interface{})

  return reqData, nil
}

