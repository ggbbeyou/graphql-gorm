package utils

import(
  "strings"
  "regexp"
  "net"
  "net/http"
  "reflect"
)

func IsEmpty(v interface{}) bool {
  if v == nil {
    return true
  }

  switch v.(type) {
  case string:
    if v == "" {
      return true
    }
  default:
    return false
  }
  return false
}

// 查找数组并返回下标
func IndexOf(str []interface{}, data interface{}) int {
  for k, v := range str{
    if v == data {
      return k
    }
  }

  return - 1
}

// []StringTo[]interface
func ArrStrTointerface(data []string) []interface{} {
  newArr := make([]interface{}, len(data))
  for i, v := range data {
    newArr[i] = v
  }
  return newArr
}

// stringToArray
func StrToArr(data, split string) []string {
  if IsEmpty(data) {
    return nil
  }
  return strings.Split(data, split)
}

// 正则替换
func ReplaceAll(data, reg, target string) string {
  req, _ := regexp.Compile(reg);
  rep := req.ReplaceAllString(data, target);
  return rep
}

func ToLower(str string) string {
  return strings.ToLower(str)
}

const (
  XForwardedFor = "X-Forwarded-For"
  XRealIP       = "X-Real-IP"
)

// RemoteIp 返回远程客户端的 IP，如 192.168.1.1
func RemoteIp(req *http.Request) string {
  remoteAddr := req.RemoteAddr
  if ip := req.Header.Get(XRealIP); ip != "" {
    remoteAddr = ip
  } else if ip = req.Header.Get(XForwardedFor); ip != "" {
    remoteAddr = ip
  } else {
    remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
  }

  if remoteAddr == "::1" {
    remoteAddr = "127.0.0.1"
  }

  return remoteAddr
}

// 首字母大写
func StrFirstToUpper(str string) string {
  if len(str) < 1 {
    return ""
  }

  strArry := []rune(str)
  if strArry[0] >= 97 && strArry[0] <= 122  {
    strArry[0] -=  32
  }
  return string(strArry)
}

// 转驼峰命名
func CamelString(s string) string {
  data := make([]byte, 0, len(s))
  j := false
  k := false
  num := len(s) - 1

  for i := 0; i <= num; i++ {
    d := s[i]
    if k == false && d >= 'A' && d <= 'Z' {
      k = true
    }
    if d >= 'a' && d <= 'z' && (j || k == false) {
      d = d - 32
      j = false
      k = true
    }

    if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
      j = true
      continue
    }

    data = append(data, d)
  }
  return string(data[:])
}

// Struct To Map
func StructToMap(obj interface{}) map[string]interface{} {
  t := reflect.ValueOf(obj)
  elem := t.Elem()
  key := elem.Type()

  var data = make(map[string]interface{})
  for i := 0; i < elem.NumField(); i++ {
    val := elem.Field(i)
    value := val.Interface()

    tag := key.Field(i).Tag
    json := tag.Get("json")
    if json == "" {
      json = key.Field(i).Name
    }

    data[json] = value
  }

  return data
}
