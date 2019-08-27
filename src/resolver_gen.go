package src

import (
	"github.com/maiguangyang/graphql-gorm/gen"
	"github.com/maiguangyang/graphql/events"
)

func NewResolver(db *gen.DB, ec *events.EventController) *Resolver {
	handlers := gen.DefaultResolutionHandlers()
	return &Resolver{&gen.GeneratedResolver{Handlers: handlers, DB: db, EventController: ec}}
}

type Resolver struct {
	*gen.GeneratedResolver
}

type MutationResolver struct {
	*gen.GeneratedMutationResolver
}

type QueryResolver struct {
	*gen.GeneratedQueryResolver
}

func (r *Resolver) Mutation() gen.MutationResolver {
	return &MutationResolver{&gen.GeneratedMutationResolver{r.GeneratedResolver}}
}
func (r *Resolver) Query() gen.QueryResolver {
	return &QueryResolver{&gen.GeneratedQueryResolver{r.GeneratedResolver}}
}

type UserResultTypeResolver struct {
	*gen.GeneratedUserResultTypeResolver
}

func (r *Resolver) UserResultType() gen.UserResultTypeResolver {
	return &UserResultTypeResolver{&gen.GeneratedUserResultTypeResolver{r.GeneratedResolver}}
}

type UserResolver struct {
	*gen.GeneratedUserResolver
}

func (r *Resolver) User() gen.UserResolver {
	return &UserResolver{&gen.GeneratedUserResolver{r.GeneratedResolver}}
}

type TaskResultTypeResolver struct {
	*gen.GeneratedTaskResultTypeResolver
}

func (r *Resolver) TaskResultType() gen.TaskResultTypeResolver {
	return &TaskResultTypeResolver{&gen.GeneratedTaskResultTypeResolver{r.GeneratedResolver}}
}

type TaskResolver struct {
	*gen.GeneratedTaskResolver
}

func (r *Resolver) Task() gen.TaskResolver {
	return &TaskResolver{&gen.GeneratedTaskResolver{r.GeneratedResolver}}
}

type AdminResultTypeResolver struct {
	*gen.GeneratedAdminResultTypeResolver
}

func (r *Resolver) AdminResultType() gen.AdminResultTypeResolver {
	return &AdminResultTypeResolver{&gen.GeneratedAdminResultTypeResolver{r.GeneratedResolver}}
}

type AdminResolver struct {
	*gen.GeneratedAdminResolver
}

func (r *Resolver) Admin() gen.AdminResolver {
	return &AdminResolver{&gen.GeneratedAdminResolver{r.GeneratedResolver}}
}

type GroupResultTypeResolver struct {
	*gen.GeneratedGroupResultTypeResolver
}

func (r *Resolver) GroupResultType() gen.GroupResultTypeResolver {
	return &GroupResultTypeResolver{&gen.GeneratedGroupResultTypeResolver{r.GeneratedResolver}}
}

type GroupResolver struct {
	*gen.GeneratedGroupResolver
}

func (r *Resolver) Group() gen.GroupResolver {
	return &GroupResolver{&gen.GeneratedGroupResolver{r.GeneratedResolver}}
}

type RoleResultTypeResolver struct {
	*gen.GeneratedRoleResultTypeResolver
}

func (r *Resolver) RoleResultType() gen.RoleResultTypeResolver {
	return &RoleResultTypeResolver{&gen.GeneratedRoleResultTypeResolver{r.GeneratedResolver}}
}

type RoleResolver struct {
	*gen.GeneratedRoleResolver
}

func (r *Resolver) Role() gen.RoleResolver {
	return &RoleResolver{&gen.GeneratedRoleResolver{r.GeneratedResolver}}
}
