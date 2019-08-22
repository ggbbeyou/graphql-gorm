package src

import (
	"context"
	"github.com/maiguangyang/graphql-gorm/gen"
	"github.com/maiguangyang/graphql/events"
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

func (r *MutationResolver) Token(ctx context.Context, input *string) (string, error) {
	return "token", nil
}