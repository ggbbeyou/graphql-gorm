package src

import (
	// "fmt"
	"context"
	"encoding/json"
	"github.com/maiguangyang/graphql-gorm/gen"
	"github.com/maiguangyang/graphql/events"
	"github.com/maiguangyang/graphql-gorm/utils"
	"github.com/maiguangyang/graphql-gorm/middleware"
	"github.com/maiguangyang/graphql-gorm/cache"
)

func New(db *gen.DB, ec *events.EventController) *Resolver {
	resolver := NewResolver(db, ec)

	// resolver.Handlers.CreateUser = func(ctx context.Context, r *gen.GeneratedMutationResolver, input map[string]interface{}) (item *gen.Company, err error) {
	// 	return gen.CreateUserHandler(ctx, r, input)
	// }

	return resolver
}

// You can extend QueryResolver for adding custom fields in schema
// func (r *QueryResolver) Hello(ctx context.Context) (string, error) {
// 	return "world", nil
// }

// 自定义登录方法
func (r *MutationResolver) Login(ctx context.Context, email string) (*interface{}, error) {
	var resData interface{}
	var token string
	var user *gen.User
	var err error

  ip := ctx.Value("RemoteIp")

  // rUser, has := gen.RidesCache.HGetAll("user")

  // if has == false {
		// 根据条件查询用户
		var opts gen.QueryUserHandlerOptions
		opts.Filter = &gen.UserFilterType{
			Email: &email,
		}

		user, err = gen.QueryUserHandler(ctx, r.GeneratedResolver, opts)

		if err != nil {
			resData = "登录密码错误"
			return &resData, nil
		}

		// Struct To Map
		userInfo := make(map[string]interface{})
    j, _ := json.Marshal(user)
    json.Unmarshal(j, &userInfo)

		gen.RidesCache.HMSet(cache.RidesKeys["userInfo"] + user.ID, userInfo)
  // }

	// 生成JWT Token
  token = middleware.SetToken(map[string]interface{}{
    "id": user.ID,
  }, utils.EncryptMd5(ip.(string) + middleware.SecretKey["admin"].(string)), "admin")

	// 组装返回数据
	resData = map[string]interface{}{
		"user": map[string]interface{}{
			"id"    : user.ID,
			"email" : user.Email,
			"state" : user.State,
		},
		"token": token,
	}

	return &resData, nil
}