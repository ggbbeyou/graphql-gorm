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
	ID        string     `json:"id" gorm:"column:id;primary_key"`
	Email     *string    `json:"email" gorm:"column:email"`
	FirstName *string    `json:"firstName" gorm:"column:firstName"`
	LastName  *string    `json:"lastName" gorm:"column:lastName"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy *string    `json:"createdBy" gorm:"column:createdBy"`

	Tasks []*Task `json:"tasks" gorm:"foreignkey:AssigneeID"`
}

type UserChanges struct {
	ID        string
	Email     *string
	FirstName *string
	LastName  *string
	UpdatedAt *time.Time
	CreatedAt time.Time
	UpdatedBy *string
	CreatedBy *string
}

type TaskResultType struct {
	resolvers.EntityResultType
}

type Task struct {
	ID         string     `json:"id" gorm:"column:id;primary_key"`
	Title      *string    `json:"title" gorm:"column:title"`
	Completed  *bool      `json:"completed" gorm:"column:completed"`
	DueDate    *time.Time `json:"dueDate" gorm:"column:dueDate"`
	AssigneeID *string    `json:"assigneeId" gorm:"column:assigneeId"`
	UpdatedAt  *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt  time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy  *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy  *string    `json:"createdBy" gorm:"column:createdBy"`

	Assignee *User `json:"assignee"`
}

type TaskChanges struct {
	ID         string
	Title      *string
	Completed  *bool
	DueDate    *time.Time
	AssigneeID *string
	UpdatedAt  *time.Time
	CreatedAt  time.Time
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
