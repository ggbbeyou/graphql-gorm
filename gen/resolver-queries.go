package gen

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/graph-gophers/dataloader"
	"github.com/maiguangyang/graphql/resolvers"
	"github.com/vektah/gqlparser/ast"
)

type GeneratedQueryResolver struct{ *GeneratedResolver }

type QueryUserHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *UserFilterType
}

func (r *GeneratedQueryResolver) User(ctx context.Context, id *string, q *string, filter *UserFilterType) (*User, error) {
	opts := QueryUserHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryUser(ctx, r.GeneratedResolver, opts)
}
func QueryUserHandler(ctx context.Context, r *GeneratedResolver, opts QueryUserHandlerOptions) (*User, error) {
	query := UserQueryFilter{opts.Q}
	current_page := 0
	per_page := 0
	rt := &UserResultType{
		EntityResultType: resolvers.EntityResultType{
			CurrentPage: &current_page,
			PerPage:     &per_page,
			Query:       &query,
			Filter:      opts.Filter,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where("users.id = ?", *opts.ID)
	}

	var items []*User
	err := rt.GetData(ctx, qb, "users", &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, fmt.Errorf("User not found")
	}
	return items[0], err
}

type QueryUsersHandlerOptions struct {
	CurrentPage *int
	PerPage     *int
	Q           *string
	Sort        []UserSortType
	Filter      *UserFilterType
}

func (r *GeneratedQueryResolver) Users(ctx context.Context, current_page *int, per_page *int, q *string, sort []UserSortType, filter *UserFilterType) (*UserResultType, error) {
	opts := QueryUsersHandlerOptions{
		CurrentPage: current_page,
		PerPage:     per_page,
		Q:           q,
		Sort:        sort,
		Filter:      filter,
	}
	return r.Handlers.QueryUsers(ctx, r.GeneratedResolver, opts)
}
func QueryUsersHandler(ctx context.Context, r *GeneratedResolver, opts QueryUsersHandlerOptions) (*UserResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range opts.Sort {
		_sort = append(_sort, s)
	}
	query := UserQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	return &UserResultType{
		EntityResultType: resolvers.EntityResultType{
			CurrentPage:  opts.CurrentPage,
			PerPage:      opts.PerPage,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedUserResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedUserResultTypeResolver) Data(ctx context.Context, obj *UserResultType) (items []*User, err error) {
	err = obj.GetData(ctx, r.DB.db, "users", &items)
	return
}

func (r *GeneratedUserResultTypeResolver) Total(ctx context.Context, obj *UserResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, &User{})
}

func (r *GeneratedUserResultTypeResolver) CurrentPage(ctx context.Context, obj *UserResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedUserResultTypeResolver) PerPage(ctx context.Context, obj *UserResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}

func (r *GeneratedUserResultTypeResolver) TotalPage(ctx context.Context, obj *UserResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))

	return totalPage, nil
}

type GeneratedUserResolver struct{ *GeneratedResolver }

func (r *GeneratedUserResolver) Tasks(ctx context.Context, obj *User) (res []*Task, err error) {

	selects := resolvers.GetFieldsRequested(ctx, strings.ToLower("Tasks"))

	items := []*Task{}
	err = r.DB.Query().Where("state = ?", 1).Select(selects).Model(obj).Related(&items, "Tasks").Error
	res = items

	return
}
func UserTasksHandler(ctx context.Context, r *GeneratedUserResolver, obj *User) (res []*Task, err error) {

	selects := resolvers.GetFieldsRequested(ctx, strings.ToLower("Tasks"))

	items := []*Task{}
	err = r.DB.Query().Where("state = ?", 1).Select(selects).Model(obj).Related(&items, "Tasks").Error
	res = items

	return
}

func (r *GeneratedUserResolver) TasksIds(ctx context.Context, obj *User) (ids []string, err error) {
	ids = []string{}

	items := []*Task{}
	err = r.DB.Query().Model(obj).Select("tasks.id").Related(&items, "Tasks").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

type QueryTaskHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *TaskFilterType
}

func (r *GeneratedQueryResolver) Task(ctx context.Context, id *string, q *string, filter *TaskFilterType) (*Task, error) {
	opts := QueryTaskHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryTask(ctx, r.GeneratedResolver, opts)
}
func QueryTaskHandler(ctx context.Context, r *GeneratedResolver, opts QueryTaskHandlerOptions) (*Task, error) {
	query := TaskQueryFilter{opts.Q}
	current_page := 0
	per_page := 0
	rt := &TaskResultType{
		EntityResultType: resolvers.EntityResultType{
			CurrentPage: &current_page,
			PerPage:     &per_page,
			Query:       &query,
			Filter:      opts.Filter,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where("tasks.id = ?", *opts.ID)
	}

	var items []*Task
	err := rt.GetData(ctx, qb, "tasks", &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, fmt.Errorf("Task not found")
	}
	return items[0], err
}

type QueryTasksHandlerOptions struct {
	CurrentPage *int
	PerPage     *int
	Q           *string
	Sort        []TaskSortType
	Filter      *TaskFilterType
}

func (r *GeneratedQueryResolver) Tasks(ctx context.Context, current_page *int, per_page *int, q *string, sort []TaskSortType, filter *TaskFilterType) (*TaskResultType, error) {
	opts := QueryTasksHandlerOptions{
		CurrentPage: current_page,
		PerPage:     per_page,
		Q:           q,
		Sort:        sort,
		Filter:      filter,
	}
	return r.Handlers.QueryTasks(ctx, r.GeneratedResolver, opts)
}
func QueryTasksHandler(ctx context.Context, r *GeneratedResolver, opts QueryTasksHandlerOptions) (*TaskResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range opts.Sort {
		_sort = append(_sort, s)
	}
	query := TaskQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	return &TaskResultType{
		EntityResultType: resolvers.EntityResultType{
			CurrentPage:  opts.CurrentPage,
			PerPage:      opts.PerPage,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedTaskResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedTaskResultTypeResolver) Data(ctx context.Context, obj *TaskResultType) (items []*Task, err error) {
	err = obj.GetData(ctx, r.DB.db, "tasks", &items)
	return
}

func (r *GeneratedTaskResultTypeResolver) Total(ctx context.Context, obj *TaskResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, &Task{})
}

func (r *GeneratedTaskResultTypeResolver) CurrentPage(ctx context.Context, obj *TaskResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedTaskResultTypeResolver) PerPage(ctx context.Context, obj *TaskResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}

func (r *GeneratedTaskResultTypeResolver) TotalPage(ctx context.Context, obj *TaskResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))

	return totalPage, nil
}

type GeneratedTaskResolver struct{ *GeneratedResolver }

func (r *GeneratedTaskResolver) Assignee(ctx context.Context, obj *Task) (res *User, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.AssigneeID != nil {
		item, _err := loaders["User"].Load(ctx, dataloader.StringKey(*obj.AssigneeID))()
		res, _ = item.(*User)

		err = _err
	}

	return
}
func TaskAssigneeHandler(ctx context.Context, r *GeneratedTaskResolver, obj *Task) (res *User, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.AssigneeID != nil {
		item, _err := loaders["User"].Load(ctx, dataloader.StringKey(*obj.AssigneeID))()
		res, _ = item.(*User)

		err = _err
	}

	return
}

type QueryAdminHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *AdminFilterType
}

func (r *GeneratedQueryResolver) Admin(ctx context.Context, id *string, q *string, filter *AdminFilterType) (*Admin, error) {
	opts := QueryAdminHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryAdmin(ctx, r.GeneratedResolver, opts)
}
func QueryAdminHandler(ctx context.Context, r *GeneratedResolver, opts QueryAdminHandlerOptions) (*Admin, error) {
	query := AdminQueryFilter{opts.Q}
	current_page := 0
	per_page := 0
	rt := &AdminResultType{
		EntityResultType: resolvers.EntityResultType{
			CurrentPage: &current_page,
			PerPage:     &per_page,
			Query:       &query,
			Filter:      opts.Filter,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where("admins.id = ?", *opts.ID)
	}

	var items []*Admin
	err := rt.GetData(ctx, qb, "admins", &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, fmt.Errorf("Admin not found")
	}
	return items[0], err
}

type QueryAdminsHandlerOptions struct {
	CurrentPage *int
	PerPage     *int
	Q           *string
	Sort        []AdminSortType
	Filter      *AdminFilterType
}

func (r *GeneratedQueryResolver) Admins(ctx context.Context, current_page *int, per_page *int, q *string, sort []AdminSortType, filter *AdminFilterType) (*AdminResultType, error) {
	opts := QueryAdminsHandlerOptions{
		CurrentPage: current_page,
		PerPage:     per_page,
		Q:           q,
		Sort:        sort,
		Filter:      filter,
	}
	return r.Handlers.QueryAdmins(ctx, r.GeneratedResolver, opts)
}
func QueryAdminsHandler(ctx context.Context, r *GeneratedResolver, opts QueryAdminsHandlerOptions) (*AdminResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range opts.Sort {
		_sort = append(_sort, s)
	}
	query := AdminQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	return &AdminResultType{
		EntityResultType: resolvers.EntityResultType{
			CurrentPage:  opts.CurrentPage,
			PerPage:      opts.PerPage,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedAdminResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedAdminResultTypeResolver) Data(ctx context.Context, obj *AdminResultType) (items []*Admin, err error) {
	err = obj.GetData(ctx, r.DB.db, "admins", &items)
	return
}

func (r *GeneratedAdminResultTypeResolver) Total(ctx context.Context, obj *AdminResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, &Admin{})
}

func (r *GeneratedAdminResultTypeResolver) CurrentPage(ctx context.Context, obj *AdminResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedAdminResultTypeResolver) PerPage(ctx context.Context, obj *AdminResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}

func (r *GeneratedAdminResultTypeResolver) TotalPage(ctx context.Context, obj *AdminResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))

	return totalPage, nil
}

type GeneratedAdminResolver struct{ *GeneratedResolver }

func (r *GeneratedAdminResolver) Groups(ctx context.Context, obj *Admin) (res []*Group, err error) {

	selects := resolvers.GetFieldsRequested(ctx, strings.ToLower("Groups"))

	items := []*Group{}
	err = r.DB.Query().Where("state = ?", 1).Select(selects).Model(obj).Related(&items, "Groups").Error
	res = items

	return
}
func AdminGroupsHandler(ctx context.Context, r *GeneratedAdminResolver, obj *Admin) (res []*Group, err error) {

	selects := resolvers.GetFieldsRequested(ctx, strings.ToLower("Groups"))

	items := []*Group{}
	err = r.DB.Query().Where("state = ?", 1).Select(selects).Model(obj).Related(&items, "Groups").Error
	res = items

	return
}

func (r *GeneratedAdminResolver) GroupsIds(ctx context.Context, obj *Admin) (ids []string, err error) {
	ids = []string{}

	items := []*Group{}
	err = r.DB.Query().Model(obj).Select("groups.id").Related(&items, "Groups").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

func (r *GeneratedAdminResolver) Roles(ctx context.Context, obj *Admin) (res []*Role, err error) {

	selects := resolvers.GetFieldsRequested(ctx, strings.ToLower("Roles"))

	items := []*Role{}
	err = r.DB.Query().Where("state = ?", 1).Select(selects).Model(obj).Related(&items, "Roles").Error
	res = items

	return
}
func AdminRolesHandler(ctx context.Context, r *GeneratedAdminResolver, obj *Admin) (res []*Role, err error) {

	selects := resolvers.GetFieldsRequested(ctx, strings.ToLower("Roles"))

	items := []*Role{}
	err = r.DB.Query().Where("state = ?", 1).Select(selects).Model(obj).Related(&items, "Roles").Error
	res = items

	return
}

func (r *GeneratedAdminResolver) RolesIds(ctx context.Context, obj *Admin) (ids []string, err error) {
	ids = []string{}

	items := []*Role{}
	err = r.DB.Query().Model(obj).Select("roles.id").Related(&items, "Roles").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

type QueryGroupHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *GroupFilterType
}

func (r *GeneratedQueryResolver) Group(ctx context.Context, id *string, q *string, filter *GroupFilterType) (*Group, error) {
	opts := QueryGroupHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryGroup(ctx, r.GeneratedResolver, opts)
}
func QueryGroupHandler(ctx context.Context, r *GeneratedResolver, opts QueryGroupHandlerOptions) (*Group, error) {
	query := GroupQueryFilter{opts.Q}
	current_page := 0
	per_page := 0
	rt := &GroupResultType{
		EntityResultType: resolvers.EntityResultType{
			CurrentPage: &current_page,
			PerPage:     &per_page,
			Query:       &query,
			Filter:      opts.Filter,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where("groups.id = ?", *opts.ID)
	}

	var items []*Group
	err := rt.GetData(ctx, qb, "groups", &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, fmt.Errorf("Group not found")
	}
	return items[0], err
}

type QueryGroupsHandlerOptions struct {
	CurrentPage *int
	PerPage     *int
	Q           *string
	Sort        []GroupSortType
	Filter      *GroupFilterType
}

func (r *GeneratedQueryResolver) Groups(ctx context.Context, current_page *int, per_page *int, q *string, sort []GroupSortType, filter *GroupFilterType) (*GroupResultType, error) {
	opts := QueryGroupsHandlerOptions{
		CurrentPage: current_page,
		PerPage:     per_page,
		Q:           q,
		Sort:        sort,
		Filter:      filter,
	}
	return r.Handlers.QueryGroups(ctx, r.GeneratedResolver, opts)
}
func QueryGroupsHandler(ctx context.Context, r *GeneratedResolver, opts QueryGroupsHandlerOptions) (*GroupResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range opts.Sort {
		_sort = append(_sort, s)
	}
	query := GroupQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	return &GroupResultType{
		EntityResultType: resolvers.EntityResultType{
			CurrentPage:  opts.CurrentPage,
			PerPage:      opts.PerPage,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedGroupResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedGroupResultTypeResolver) Data(ctx context.Context, obj *GroupResultType) (items []*Group, err error) {
	err = obj.GetData(ctx, r.DB.db, "groups", &items)
	return
}

func (r *GeneratedGroupResultTypeResolver) Total(ctx context.Context, obj *GroupResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, &Group{})
}

func (r *GeneratedGroupResultTypeResolver) CurrentPage(ctx context.Context, obj *GroupResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedGroupResultTypeResolver) PerPage(ctx context.Context, obj *GroupResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}

func (r *GeneratedGroupResultTypeResolver) TotalPage(ctx context.Context, obj *GroupResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))

	return totalPage, nil
}

type GeneratedGroupResolver struct{ *GeneratedResolver }

func (r *GeneratedGroupResolver) Admin(ctx context.Context, obj *Group) (res []*Admin, err error) {

	selects := resolvers.GetFieldsRequested(ctx, strings.ToLower("Admin"))

	items := []*Admin{}
	err = r.DB.Query().Where("state = ?", 1).Select(selects).Model(obj).Related(&items, "Admin").Error
	res = items

	return
}
func GroupAdminHandler(ctx context.Context, r *GeneratedGroupResolver, obj *Group) (res []*Admin, err error) {

	selects := resolvers.GetFieldsRequested(ctx, strings.ToLower("Admin"))

	items := []*Admin{}
	err = r.DB.Query().Where("state = ?", 1).Select(selects).Model(obj).Related(&items, "Admin").Error
	res = items

	return
}

func (r *GeneratedGroupResolver) AdminIds(ctx context.Context, obj *Group) (ids []string, err error) {
	ids = []string{}

	items := []*Admin{}
	err = r.DB.Query().Model(obj).Select("admins.id").Related(&items, "Admin").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

func (r *GeneratedGroupResolver) Roles(ctx context.Context, obj *Group) (res []*Role, err error) {

	selects := resolvers.GetFieldsRequested(ctx, strings.ToLower("Roles"))

	items := []*Role{}
	err = r.DB.Query().Where("state = ?", 1).Select(selects).Model(obj).Related(&items, "Roles").Error
	res = items

	return
}
func GroupRolesHandler(ctx context.Context, r *GeneratedGroupResolver, obj *Group) (res []*Role, err error) {

	selects := resolvers.GetFieldsRequested(ctx, strings.ToLower("Roles"))

	items := []*Role{}
	err = r.DB.Query().Where("state = ?", 1).Select(selects).Model(obj).Related(&items, "Roles").Error
	res = items

	return
}

func (r *GeneratedGroupResolver) RolesIds(ctx context.Context, obj *Group) (ids []string, err error) {
	ids = []string{}

	items := []*Role{}
	err = r.DB.Query().Model(obj).Select("roles.id").Related(&items, "Roles").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

type QueryRoleHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *RoleFilterType
}

func (r *GeneratedQueryResolver) Role(ctx context.Context, id *string, q *string, filter *RoleFilterType) (*Role, error) {
	opts := QueryRoleHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryRole(ctx, r.GeneratedResolver, opts)
}
func QueryRoleHandler(ctx context.Context, r *GeneratedResolver, opts QueryRoleHandlerOptions) (*Role, error) {
	query := RoleQueryFilter{opts.Q}
	current_page := 0
	per_page := 0
	rt := &RoleResultType{
		EntityResultType: resolvers.EntityResultType{
			CurrentPage: &current_page,
			PerPage:     &per_page,
			Query:       &query,
			Filter:      opts.Filter,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where("roles.id = ?", *opts.ID)
	}

	var items []*Role
	err := rt.GetData(ctx, qb, "roles", &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, fmt.Errorf("Role not found")
	}
	return items[0], err
}

type QueryRolesHandlerOptions struct {
	CurrentPage *int
	PerPage     *int
	Q           *string
	Sort        []RoleSortType
	Filter      *RoleFilterType
}

func (r *GeneratedQueryResolver) Roles(ctx context.Context, current_page *int, per_page *int, q *string, sort []RoleSortType, filter *RoleFilterType) (*RoleResultType, error) {
	opts := QueryRolesHandlerOptions{
		CurrentPage: current_page,
		PerPage:     per_page,
		Q:           q,
		Sort:        sort,
		Filter:      filter,
	}
	return r.Handlers.QueryRoles(ctx, r.GeneratedResolver, opts)
}
func QueryRolesHandler(ctx context.Context, r *GeneratedResolver, opts QueryRolesHandlerOptions) (*RoleResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range opts.Sort {
		_sort = append(_sort, s)
	}
	query := RoleQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	return &RoleResultType{
		EntityResultType: resolvers.EntityResultType{
			CurrentPage:  opts.CurrentPage,
			PerPage:      opts.PerPage,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedRoleResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedRoleResultTypeResolver) Data(ctx context.Context, obj *RoleResultType) (items []*Role, err error) {
	err = obj.GetData(ctx, r.DB.db, "roles", &items)
	return
}

func (r *GeneratedRoleResultTypeResolver) Total(ctx context.Context, obj *RoleResultType) (count int, err error) {
	return obj.GetTotal(ctx, r.DB.db, &Role{})
}

func (r *GeneratedRoleResultTypeResolver) CurrentPage(ctx context.Context, obj *RoleResultType) (count int, err error) {
	return int(*obj.EntityResultType.CurrentPage), nil
}

func (r *GeneratedRoleResultTypeResolver) PerPage(ctx context.Context, obj *RoleResultType) (count int, err error) {
	return int(*obj.EntityResultType.PerPage), nil
}

func (r *GeneratedRoleResultTypeResolver) TotalPage(ctx context.Context, obj *RoleResultType) (count int, err error) {
	total, _ := r.Total(ctx, obj)
	perPage, _ := r.PerPage(ctx, obj)
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))

	return totalPage, nil
}

type GeneratedRoleResolver struct{ *GeneratedResolver }

func (r *GeneratedRoleResolver) Admin(ctx context.Context, obj *Role) (res []*Admin, err error) {

	selects := resolvers.GetFieldsRequested(ctx, strings.ToLower("Admin"))

	items := []*Admin{}
	err = r.DB.Query().Where("state = ?", 1).Select(selects).Model(obj).Related(&items, "Admin").Error
	res = items

	return
}
func RoleAdminHandler(ctx context.Context, r *GeneratedRoleResolver, obj *Role) (res []*Admin, err error) {

	selects := resolvers.GetFieldsRequested(ctx, strings.ToLower("Admin"))

	items := []*Admin{}
	err = r.DB.Query().Where("state = ?", 1).Select(selects).Model(obj).Related(&items, "Admin").Error
	res = items

	return
}

func (r *GeneratedRoleResolver) AdminIds(ctx context.Context, obj *Role) (ids []string, err error) {
	ids = []string{}

	items := []*Admin{}
	err = r.DB.Query().Model(obj).Select("admins.id").Related(&items, "Admin").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

func (r *GeneratedRoleResolver) Group(ctx context.Context, obj *Role) (res []*Admin, err error) {

	selects := resolvers.GetFieldsRequested(ctx, strings.ToLower("Group"))

	items := []*Admin{}
	err = r.DB.Query().Where("state = ?", 1).Select(selects).Model(obj).Related(&items, "Group").Error
	res = items

	return
}
func RoleGroupHandler(ctx context.Context, r *GeneratedRoleResolver, obj *Role) (res []*Admin, err error) {

	selects := resolvers.GetFieldsRequested(ctx, strings.ToLower("Group"))

	items := []*Admin{}
	err = r.DB.Query().Where("state = ?", 1).Select(selects).Model(obj).Related(&items, "Group").Error
	res = items

	return
}

func (r *GeneratedRoleResolver) GroupIds(ctx context.Context, obj *Role) (ids []string, err error) {
	ids = []string{}

	items := []*Admin{}
	err = r.DB.Query().Model(obj).Select("admins.id").Related(&items, "Group").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}
