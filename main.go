package main

import (
	"fmt"
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/maiguangyang/graphql/events"

	// "github.com/rs/cors"
	"github.com/maiguangyang/graphql-gorm/gen"
	"github.com/maiguangyang/graphql-gorm/utils"
	"github.com/maiguangyang/graphql-gorm/middleware"
)

const (
	defaultPort = "80"
)

func main() {
	// mux := http.NewServeMux()
	mux := mux.NewRouter()
	mux.Use(middleware.AuthHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	urlString := os.Getenv("DATABASE_URL")
	if urlString == "" {
		panic(fmt.Errorf("missing DATABASE_URL environment variable"))
	}
	db := gen.NewDBWithString(urlString)
	defer db.Close()
	db.AutoMigrate()
	eventController, err := events.NewEventController()
	if err != nil {
		panic(err)
	}

	loaders := gen.GetLoaders(db)

	gqlHandler := handler.GraphQL(gen.NewExecutableSchema(gen.Config{Resolvers: NewResolver(db, &eventController)}),

		// 中间件进行登录Token校验
		handler.ResolverMiddleware(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
			// 检测是否需要登录
			path 		:= graphql.GetResolverContext(ctx).Path()
			isAuth 	:= utils.CheckRouterIsAuth(path)
			if isAuth == true {
				auth := ctx.Value("Authorization").(map[string]interface{})
				if len(auth) <= 0 {
					return nil, fmt.Errorf("Invalid Authorization")
				}
			}

			return next(ctx)
		}),
	)

	playgroundHandler := handler.Playground("GraphQL playground", "/graphql")
	mux.HandleFunc("/graphql", func(res http.ResponseWriter, req *http.Request) {
		// principalID := getPrincipalID(req)


		// ctx := context.WithValue(req.Context(), gen.KeyPrincipalID, principalID)
		ctx := context.WithValue(req.Context(), "loaders", loaders)
		req = req.WithContext(ctx)
		if req.Method == "GET" {
			playgroundHandler(res, req)
		} else {
			gqlHandler(res, req)
		}
	})
	mux.HandleFunc("/healthcheck", func(res http.ResponseWriter, req *http.Request) {
		if err := db.Ping(); err != nil {
			res.WriteHeader(400)
			res.Write([]byte("ERROR"))
			return
		}
		res.WriteHeader(200)
		res.Write([]byte("OK"))
	})
	handler := mux
	// use this line to allow cors for all origins/methods/headers (for development)
	// handler := cors.AllowAll().Handler(mux)

	log.Printf("connect to http://localhost:%s/graphql for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
func getPrincipalIDFromContext(ctx context.Context) *string {
	v, _ := ctx.Value(gen.KeyPrincipalID).(*string)
	return v
}
func getJWTClaimsFromContext(ctx context.Context) *JWTClaims {
	v, _ := ctx.Value(gen.KeyJWTClaims).(*JWTClaims)
	return v
}
func getPrincipalID(req *http.Request) *string {
	pID := req.Header.Get("principal-id")
	if pID != "" {
		return &pID
	}
	c, _ := getJWTClaims(req)
	if c == nil {
		return nil
	}
	return &c.Subject
}

type JWTClaims struct {
	jwtgo.StandardClaims
	Scope *string
}

func getJWTClaims(req *http.Request) (*JWTClaims, error) {
	var p *JWTClaims
	tokenStr := strings.Replace(req.Header.Get("authorization"), "Bearer ", "", 1)
	if tokenStr == "" {
		return p, nil
	}
	p = &JWTClaims{}
	jwtgo.ParseWithClaims(tokenStr, p, nil)
	return p, nil
}
