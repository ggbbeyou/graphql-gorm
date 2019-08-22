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

	items := []*Task{}
	err = r.DB.Query().Model(obj).Related(&items, "Tasks").Error
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

	loaders := ctx.Value("loaders").(map[string]*dataloader.Loader)
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
