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
	State     *int64  `json:"state" gorm:"type:int(2) comment '状态：1/正常、2/禁用、3/删除';NOT NULL;default:1;"`
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
	State      *int64     `json:"state" gorm:"type:int(2) comment '状态：1/正常、2/禁用、3/删除';NOT NULL;default:1;"`
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
	UpdatedAt  *int64
	CreatedAt  *int64
	DeletedBy  *string
	UpdatedBy  *string
	CreatedBy  *string
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
