package utils

import(
  // "fmt"
  "strings"
  "regexp"
  // "reflect"
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