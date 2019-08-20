package utils

// 需要登录验证的路由
var router = []string{
  "Task",
}

// 替换方法名称
var replaceStr = []string {
	"Create",
	"Update",
	"DeleteAll",
	"Delete",
}

// 检测路由是否需要登录
func CheckRouterIsAuth(path []interface{}) bool {
	var colName string = ""
	if len(path) > 0 {
		colName = path[0].(string)
	}

	colName = StrFirstToUpper(colName)

	for _, v := range replaceStr {
		colName = ReplaceAll(colName, v, "")
	}

	index := IndexOf(ArrStrTointerface(router), colName)

	return index != -1
}