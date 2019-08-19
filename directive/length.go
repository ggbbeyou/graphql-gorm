package directive

import(
	"fmt"
	"context"
	"github.com/99designs/gqlgen/graphql"
)

func Length(ctx context.Context, obj interface{}, next graphql.Resolver, min int, max *int, message *string) (interface{}, error) {
	// rc := graphql.GetResolverContext(ctx)
	// fmt.Println("field %s has value %#v", rc.Field.Name, obj)
	// return nil, fmt.Errorf("too long")

	e := func(msg string) error {
		if message == nil {
			return fmt.Errorf(msg)
		}
		return fmt.Errorf(*message)
	}

	res, err := next(ctx)
	if err != nil {
		return nil, err
	}

	s := string(*res.(*string))
	if len(s) < min {
		return nil, e("too short")
	}
	if max != nil && len(s) > *max {
		return nil, e("too long")
	}
	return res, err
}