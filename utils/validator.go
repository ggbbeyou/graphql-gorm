package utils

import(
	"fmt"
	"strings"
	"regexp"
	"reflect"

	"github.com/vektah/gqlparser/gqlerror"
)

func Validator(table interface{}) (gqlerror.Error, error) {

  data := reflect.ValueOf(table)
  elem := data.Elem()

  key := elem.Type()
  var errText gqlerror.Error
  var err error = nil

  for i := 0; i < elem.NumField(); i++ {
  	val := elem.Field(i)
    value := val.Interface()

    tag := key.Field(i).Tag
    valid := tag.Get("validator")
    json := tag.Get("json")
    if json == "" {
      json = key.Field(i).Name
    }

    // 判定tag和自定义验证标签是否为空
    if !IsEmpty(tag) && !IsEmpty(valid) {
    	arr := StrToArr(valid, ";")
  		resData := make(map[string]interface{})

    	if len(arr) > 0 {
	    	for _, v := range arr {
	    		field := strings.Split(v, ":")
					if len(field) == 2 {
						resData[field[0]] = field[1]
					}
	    	}
    	}

    	// 字段验证
    	if res, resErr := checkField(resData, value, json); resErr != nil {
    		errText.Path = append(errText.Path, res)
    	}
    }
  }

	if len(errText.Path) > 0 {
		errText.Message = "请检查以下字段是否正确"
		err = fmt.Errorf("error")
	}
  return errText, err
}

/**
 * checkField 字段校验函数
 * @param  {[type]} resData map[string]interface{} 自定义校验指令
 * @param  {[type]} value   interface{}            值
 * @param  {[type]} json    string)                  (gqlerror.Error, error [description]
 * @return {[type]}         [description]
 */
func checkField(resData map[string]interface{}, value interface{}, json string) (string, error) {
  // var errText gqlerror.Error
  var errText string

	// 反射类型，进行判定处理
	// tye := reflect.TypeOf(value).String()
	tye := reflect.TypeOf(value).String()

	// 类型判定，取值
	var newValue interface{}
  switch tye {
	  case "string":
	    newValue = string(value.(string))
	  case "*string":
	  	if value.(*string) != nil {
	    	newValue = string(*value.(*string))
	  	}
	  case "int64":
	    newValue = int64(value.(int64))
	  case "*int64":
	  	if value.(*int64) != nil {
	    	newValue = int64(*value.(*int64))
	  	}
	  case "float64":
	    newValue = float64(value.(float64))
	  case "*float64":
	  	if value.(*float64) != nil {
	    	newValue = float64(*value.(*float64))
	  	}
  }


  // 正则格式校验
	if resData["required"] == "true" && IsEmpty(newValue) {
		errText = json + "不能为空"
	} else if resData["type"] != "" {
		var bool bool
		rl := Rule[resData["type"].(string)]
		bool = regexp.MustCompile(rl["rgx"].(string)).MatchString(fmt.Sprint(newValue))
		// if resData["type"] == "password" {
		// 	EncryptPassword(newValue.(string))
		// }

		msgText := rl["msg"].(string)
		if msgText == "" {
			msgText = "格式不正确"
		}

		if bool != true {
			errText = json + " " + msgText
		}
	}

	if errText != "" {
		// errText.Message = "请检查以下字段是否正确"
		return errText, fmt.Errorf("error")
	}

	return errText, nil
}