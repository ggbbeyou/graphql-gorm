package gen

import (
	"context"

	"github.com/maiguangyang/graphql/events"
)

type ResolutionHandlers struct {
	CreateUser     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *User, err error)
	UpdateUser     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *User, err error)
	DeleteUser     func(ctx context.Context, r *GeneratedResolver, id string) (item *User, err error)
	DeleteAllUsers func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryUser      func(ctx context.Context, r *GeneratedResolver, opts QueryUserHandlerOptions) (*User, error)
	QueryUsers     func(ctx context.Context, r *GeneratedResolver, opts QueryUsersHandlerOptions) (*UserResultType, error)

	UserTasks func(ctx context.Context, r *GeneratedUserResolver, obj *User) (res []*Task, err error)

	CreateTask     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Task, err error)
	UpdateTask     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Task, err error)
	DeleteTask     func(ctx context.Context, r *GeneratedResolver, id string) (item *Task, err error)
	DeleteAllTasks func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryTask      func(ctx context.Context, r *GeneratedResolver, opts QueryTaskHandlerOptions) (*Task, error)
	QueryTasks     func(ctx context.Context, r *GeneratedResolver, opts QueryTasksHandlerOptions) (*TaskResultType, error)

	TaskAssignee func(ctx context.Context, r *GeneratedTaskResolver, obj *Task) (res *User, err error)
}

func DefaultResolutionHandlers() ResolutionHandlers {
	handlers := ResolutionHandlers{

		CreateUser:     CreateUserHandler,
		UpdateUser:     UpdateUserHandler,
		DeleteUser:     DeleteUserHandler,
		DeleteAllUsers: DeleteAllUsersHandler,
		QueryUser:      QueryUserHandler,
		QueryUsers:     QueryUsersHandler,

		UserTasks: UserTasksHandler,

		CreateTask:     CreateTaskHandler,
		UpdateTask:     UpdateTaskHandler,
		DeleteTask:     DeleteTaskHandler,
		DeleteAllTasks: DeleteAllTasksHandler,
		QueryTask:      QueryTaskHandler,
		QueryTasks:     QueryTasksHandler,

		TaskAssignee: TaskAssigneeHandler,
	}
	return handlers
}

type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *events.EventController
}
