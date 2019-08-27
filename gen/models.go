package gen

import (
	"fmt"
	"reflect"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/maiguangyang/graphql/resolvers"
	"github.com/mitchellh/mapstructure"
)

type UserResultType struct {
	resolvers.EntityResultType
}

type User struct {
	ID        string  `json:"id" gorm:"type:varchar(36) comment 'uuid';primary_key;NOT NULL;"`
	Email     *string `json:"email" gorm:"type:varchar(64) comment '用户邮箱地址';default:null;" validator:"required:true;type:email;"`
	FirstName *string `json:"firstName" gorm:"column:firstName;null;default:null"`
	LastName  *string `json:"lastName" gorm:"column:lastName;null;default:null"`
	State     *int64  `json:"state" gorm:"type:int(2) comment '状态：1/正常、2/禁用、3/下架';NOT NULL;default:1;" validator:"required:true;type:state;"`
	Del       *int64  `json:"del" gorm:"type:int(2) comment '状态：1/正常、2/删除';NOT NULL;default:1;" validator:"required:true;type:noOrYes;"`
	UpdatedAt *int64  `json:"updatedAt" gorm:"type:int(11) comment '更新时间';null;default:null"`
	CreatedAt *int64  `json:"createdAt" gorm:"type:int(11) comment '创建时间';null;default:null"`
	DeletedBy *string `json:"deletedBy" gorm:"column:deletedBy;null;default:null"`
	UpdatedBy *string `json:"updatedBy" gorm:"column:updatedBy;null;default:null"`
	CreatedBy *string `json:"createdBy" gorm:"column:createdBy;null;default:null"`

	Tasks []*Task `json:"tasks" gorm:"foreignkey:AssigneeID"`
}

func (m *User) Is_Entity() {}

type UserChanges struct {
	ID        string
	Email     *string
	FirstName *string
	LastName  *string
	State     *int64
	Del       *int64
	UpdatedAt *int64
	CreatedAt *int64
	DeletedBy *string
	UpdatedBy *string
	CreatedBy *string

	TasksIDs []*string
}

type TaskResultType struct {
	resolvers.EntityResultType
}

type Task struct {
	ID         string     `json:"id" gorm:"type:varchar(36) comment 'uuid';primary_key;NOT NULL;"`
	Title      *string    `json:"title" gorm:"column:title;null;default:null"`
	Completed  *bool      `json:"completed" gorm:"column:completed;null;default:null"`
	DueDate    *time.Time `json:"dueDate" gorm:"column:dueDate;null;default:null"`
	AssigneeID *string    `json:"assigneeId" gorm:"column:assigneeId;null;default:null"`
	State      *int64     `json:"state" gorm:"type:int(2) comment '状态：1/正常、2/禁用、3/下架';NOT NULL;default:1;" validator:"required:true;type:state;"`
	Del        *int64     `json:"del" gorm:"type:int(2) comment '状态：1/正常、2/删除';NOT NULL;default:1;" validator:"required:true;type:noOrYes;"`
	UpdatedAt  *int64     `json:"updatedAt" gorm:"type:int(11) comment '更新时间';null;default:null"`
	CreatedAt  *int64     `json:"createdAt" gorm:"type:int(11) comment '创建时间';null;default:null"`
	DeletedBy  *string    `json:"deletedBy" gorm:"column:deletedBy;null;default:null"`
	UpdatedBy  *string    `json:"updatedBy" gorm:"column:updatedBy;null;default:null"`
	CreatedBy  *string    `json:"createdBy" gorm:"column:createdBy;null;default:null"`

	Assignee *User `json:"assignee"`
}

func (m *Task) Is_Entity() {}

type TaskChanges struct {
	ID         string
	Title      *string
	Completed  *bool
	DueDate    *time.Time
	AssigneeID *string
	State      *int64
	Del        *int64
	UpdatedAt  *int64
	CreatedAt  *int64
	DeletedBy  *string
	UpdatedBy  *string
	CreatedBy  *string
}

type AdminResultType struct {
	resolvers.EntityResultType
}

type Admin struct {
	ID         string  `json:"id" gorm:"type:varchar(36) comment 'uuid';primary_key;NOT NULL;"`
	Phone      *string `json:"Phone" gorm:"type:varchar(32) comment '手机号码';NOT NULL;" validator:"required:true;type:phone;"`
	Password   *string `json:"Password" gorm:"type:varchar(64) comment '登录密码';NOT NULL;" validator:"required:true;type:password;"`
	Username   *string `json:"Username" gorm:"type:varchar(64) comment '姓名';NOT NULL;"`
	Money      *int64  `json:"Money" gorm:"type:int(11) comment '账户余额';default:0;" validator:"required:true;type:justInt;"`
	Sex        *int64  `json:"Sex" gorm:"type:int(2) comment '性别：0/未知、1/男、2/女';default:0;" validator:"required:true;type:sex;"`
	Super      *int64  `json:"Super" gorm:"type:int(2) comment '超级账户：1/否，2/是';default:1;" validator:"required:true;type:noOrYes;"`
	LoginCount *int64  `json:"LoginCount" gorm:"type:int(11) comment '登陆次数';default:0;" validator:"required:true;type:justInt;"`
	LoginIP    *string `json:"LoginIp" gorm:"type:varchar(64) comment '登录Ip';default:null;"`
	LastIP     *string `json:"LastIp" gorm:"type:varchar(64) comment '上次登陆Ip';default:null;"`
	State      *int64  `json:"state" gorm:"type:int(2) comment '状态：1/正常、2/禁用、3/下架';NOT NULL;default:1;" validator:"required:true;type:state;"`
	Del        *int64  `json:"del" gorm:"type:int(2) comment '状态：1/正常、2/删除';NOT NULL;default:1;" validator:"required:true;type:noOrYes;"`
	UpdatedAt  *int64  `json:"updatedAt" gorm:"type:int(11) comment '更新时间';null;default:null"`
	CreatedAt  *int64  `json:"createdAt" gorm:"type:int(11) comment '创建时间';null;default:null"`
	DeletedBy  *string `json:"deletedBy" gorm:"column:deletedBy;null;default:null"`
	UpdatedBy  *string `json:"updatedBy" gorm:"column:updatedBy;null;default:null"`
	CreatedBy  *string `json:"createdBy" gorm:"column:createdBy;null;default:null"`

	Groups []*Group `json:"groups" gorm:"many2many:admin_groups;jointable_foreignkey:admin_id;association_jointable_foreignkey:group_id"`

	Roles []*Role `json:"roles" gorm:"many2many:admin_roles;jointable_foreignkey:admin_id;association_jointable_foreignkey:role_id"`
}

func (m *Admin) Is_Entity() {}

type AdminChanges struct {
	ID         string
	Phone      *string
	Password   *string
	Username   *string
	Money      *int64
	Sex        *int64
	Super      *int64
	LoginCount *int64
	LoginIP    *string
	LastIP     *string
	State      *int64
	Del        *int64
	UpdatedAt  *int64
	CreatedAt  *int64
	DeletedBy  *string
	UpdatedBy  *string
	CreatedBy  *string

	GroupsIDs []*string
	RolesIDs  []*string
}

type GroupResultType struct {
	resolvers.EntityResultType
}

type Group struct {
	ID        string  `json:"id" gorm:"type:varchar(36) comment 'uuid';primary_key;NOT NULL;"`
	Name      *string `json:"Name" gorm:"type:varchar(255) comment '部门名称';NOT NULL;"`
	State     *int64  `json:"state" gorm:"type:int(2) comment '状态：1/正常、2/禁用、3/下架';NOT NULL;default:1;" validator:"required:true;type:state;"`
	Del       *int64  `json:"del" gorm:"type:int(2) comment '状态：1/正常、2/删除';NOT NULL;default:1;" validator:"required:true;type:noOrYes;"`
	UpdatedAt *int64  `json:"updatedAt" gorm:"type:int(11) comment '更新时间';null;default:null"`
	CreatedAt *int64  `json:"createdAt" gorm:"type:int(11) comment '创建时间';null;default:null"`
	DeletedBy *string `json:"deletedBy" gorm:"column:deletedBy;null;default:null"`
	UpdatedBy *string `json:"updatedBy" gorm:"column:updatedBy;null;default:null"`
	CreatedBy *string `json:"createdBy" gorm:"column:createdBy;null;default:null"`

	Admin []*Admin `json:"admin" gorm:"many2many:admin_groups;jointable_foreignkey:group_id;association_jointable_foreignkey:admin_id"`

	Roles []*Role `json:"roles" gorm:"many2many:group_roles;jointable_foreignkey:group_id;association_jointable_foreignkey:role_id"`
}

func (m *Group) Is_Entity() {}

type GroupChanges struct {
	ID        string
	Name      *string
	State     *int64
	Del       *int64
	UpdatedAt *int64
	CreatedAt *int64
	DeletedBy *string
	UpdatedBy *string
	CreatedBy *string

	AdminIDs []*string
	RolesIDs []*string
}

type RoleResultType struct {
	resolvers.EntityResultType
}

type Role struct {
	ID        string  `json:"id" gorm:"type:varchar(36) comment 'uuid';primary_key;NOT NULL;"`
	Name      *string `json:"Name" gorm:"type:varchar(255) comment '部门名称';NOT NULL;"`
	Pid       *string `json:"Pid" gorm:"type:varchar(36) comment '所属上级：0/一级，其他对应该表的id字段';NOT NULL;"`
	State     *int64  `json:"state" gorm:"type:int(2) comment '状态：1/正常、2/禁用、3/下架';NOT NULL;default:1;" validator:"required:true;type:state;"`
	Del       *int64  `json:"del" gorm:"type:int(2) comment '状态：1/正常、2/删除';NOT NULL;default:1;" validator:"required:true;type:noOrYes;"`
	UpdatedAt *int64  `json:"updatedAt" gorm:"type:int(11) comment '更新时间';null;default:null"`
	CreatedAt *int64  `json:"createdAt" gorm:"type:int(11) comment '创建时间';null;default:null"`
	DeletedBy *string `json:"deletedBy" gorm:"column:deletedBy;null;default:null"`
	UpdatedBy *string `json:"updatedBy" gorm:"column:updatedBy;null;default:null"`
	CreatedBy *string `json:"createdBy" gorm:"column:createdBy;null;default:null"`

	Admin []*Admin `json:"admin" gorm:"many2many:admin_roles;jointable_foreignkey:role_id;association_jointable_foreignkey:admin_id"`

	Group []*Admin `json:"group" gorm:"many2many:admin_roles;jointable_foreignkey:role_id;association_jointable_foreignkey:group_id"`
}

func (m *Role) Is_Entity() {}

type RoleChanges struct {
	ID        string
	Name      *string
	Pid       *string
	State     *int64
	Del       *int64
	UpdatedAt *int64
	CreatedAt *int64
	DeletedBy *string
	UpdatedBy *string
	CreatedBy *string

	AdminIDs []*string
	GroupIDs []*string
}

// used to convert map[string]interface{} to EntityChanges struct
func ApplyChanges(changes map[string]interface{}, to interface{}) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: true,
		TagName:     "json",
		Result:      to,
		ZeroFields:  true,
		// This is needed to get mapstructure to call the gqlgen unmarshaler func for custom scalars (eg Date)
		DecodeHook: func(a reflect.Type, b reflect.Type, v interface{}) (interface{}, error) {

			if b == reflect.TypeOf(time.Time{}) {
				switch a.Kind() {
				case reflect.String:
					return time.Parse(time.RFC3339, v.(string))
				case reflect.Float64:
					return time.Unix(0, int64(v.(float64))*int64(time.Millisecond)), nil
				case reflect.Int64:
					return time.Unix(0, v.(int64)*int64(time.Millisecond)), nil
				default:
					return v, fmt.Errorf("Unable to parse date from %v", v)
				}
			}

			if reflect.PtrTo(b).Implements(reflect.TypeOf((*graphql.Unmarshaler)(nil)).Elem()) {
				resultType := reflect.New(b)
				result := resultType.MethodByName("UnmarshalGQL").Call([]reflect.Value{reflect.ValueOf(v)})
				err, _ := result[0].Interface().(error)
				return resultType.Elem().Interface(), err
			}

			return v, nil
		},
	})

	if err != nil {
		return err
	}

	return dec.Decode(changes)
}
