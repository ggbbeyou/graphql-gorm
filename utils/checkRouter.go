package utils

import(
	"fmt"
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
)

// 需要登录验证的路由
var router = []string{
  "User",
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

	if StrFirstToUpper(colName[len(colName) - 1:]) == "S" {
		colName = colName[:len(colName) - 1]
	}

	index := IndexOf(ArrStrTointerface(router), colName)

	return index != -1
}

// 导出路由验证中间件
var	RouterIsAuthMiddleware = handler.ResolverMiddleware(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
	// 检测是否需要登录
	path 		:= graphql.GetResolverContext(ctx).Path()
	isAuth 	:= CheckRouterIsAuth(path)
	if isAuth == true {
		auth := ctx.Value("Authorization").(map[string]interface{})
		if len(auth) <= 0 {
			return nil, fmt.Errorf("Invalid Authorization")
		}
	}

	return next(ctx)
})