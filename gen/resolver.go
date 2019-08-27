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

	CreateAdmin     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Admin, err error)
	UpdateAdmin     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Admin, err error)
	DeleteAdmin     func(ctx context.Context, r *GeneratedResolver, id string) (item *Admin, err error)
	DeleteAllAdmins func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryAdmin      func(ctx context.Context, r *GeneratedResolver, opts QueryAdminHandlerOptions) (*Admin, error)
	QueryAdmins     func(ctx context.Context, r *GeneratedResolver, opts QueryAdminsHandlerOptions) (*AdminResultType, error)

	AdminGroups func(ctx context.Context, r *GeneratedAdminResolver, obj *Admin) (res []*Group, err error)

	AdminRoles func(ctx context.Context, r *GeneratedAdminResolver, obj *Admin) (res []*Role, err error)

	CreateGroup     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Group, err error)
	UpdateGroup     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Group, err error)
	DeleteGroup     func(ctx context.Context, r *GeneratedResolver, id string) (item *Group, err error)
	DeleteAllGroups func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryGroup      func(ctx context.Context, r *GeneratedResolver, opts QueryGroupHandlerOptions) (*Group, error)
	QueryGroups     func(ctx context.Context, r *GeneratedResolver, opts QueryGroupsHandlerOptions) (*GroupResultType, error)

	GroupAdmin func(ctx context.Context, r *GeneratedGroupResolver, obj *Group) (res []*Admin, err error)

	GroupRoles func(ctx context.Context, r *GeneratedGroupResolver, obj *Group) (res []*Role, err error)

	CreateRole     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Role, err error)
	UpdateRole     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Role, err error)
	DeleteRole     func(ctx context.Context, r *GeneratedResolver, id string) (item *Role, err error)
	DeleteAllRoles func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryRole      func(ctx context.Context, r *GeneratedResolver, opts QueryRoleHandlerOptions) (*Role, error)
	QueryRoles     func(ctx context.Context, r *GeneratedResolver, opts QueryRolesHandlerOptions) (*RoleResultType, error)

	RoleAdmin func(ctx context.Context, r *GeneratedRoleResolver, obj *Role) (res []*Admin, err error)

	RoleGroup func(ctx context.Context, r *GeneratedRoleResolver, obj *Role) (res []*Admin, err error)
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

		CreateAdmin:     CreateAdminHandler,
		UpdateAdmin:     UpdateAdminHandler,
		DeleteAdmin:     DeleteAdminHandler,
		DeleteAllAdmins: DeleteAllAdminsHandler,
		QueryAdmin:      QueryAdminHandler,
		QueryAdmins:     QueryAdminsHandler,

		AdminGroups: AdminGroupsHandler,

		AdminRoles: AdminRolesHandler,

		CreateGroup:     CreateGroupHandler,
		UpdateGroup:     UpdateGroupHandler,
		DeleteGroup:     DeleteGroupHandler,
		DeleteAllGroups: DeleteAllGroupsHandler,
		QueryGroup:      QueryGroupHandler,
		QueryGroups:     QueryGroupsHandler,

		GroupAdmin: GroupAdminHandler,

		GroupRoles: GroupRolesHandler,

		CreateRole:     CreateRoleHandler,
		UpdateRole:     UpdateRoleHandler,
		DeleteRole:     DeleteRoleHandler,
		DeleteAllRoles: DeleteAllRolesHandler,
		QueryRole:      QueryRoleHandler,
		QueryRoles:     QueryRolesHandler,

		RoleAdmin: RoleAdminHandler,

		RoleGroup: RoleGroupHandler,
	}
	return handlers
}

type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *events.EventController
}
