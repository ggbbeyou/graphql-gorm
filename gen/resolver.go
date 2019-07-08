package gen

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"reflect"
  "strings"
  "regexp"
  "strconv"

	"github.com/maiguangyang/graphql/events"
	"github.com/maiguangyang/graphql/resolvers"
	uuid "github.com/satori/go.uuid"
)

func getPrincipalID(ctx context.Context) *string {
	v, _ := ctx.Value(KeyPrincipalID).(*string)
	return v
}

var Rule = map[string]map[string]interface{}{
  "empty"    : { "rgx": "^\\S", "msg": "不能为空", "bool": true, },
  "int"      : { "rgx": "^[0-9]\\d*$", "msg": "必须是0-9的整数", },
  "justInt"  : { "rgx": "^[1-9]\\d*$", "msg": "必须是大于0的整数", },
  "code"     : { "rgx": "^([0-9]){6}$", "msg": "验证码必须是6位整数", },
  "url"      : { "rgx": "^https?:\\/\\/.+$", "msg": "网址格式不正确", },
  "password" : { "rgx": "^(\\S){6,20}$", "msg": "密码为6-20个字符", },
  "email"    : { "rgx": "^([a-z0-9\\+\\_\\-]+)(\\.[a-z0-9\\+\\_\\-]+)*@([a-z0-9\\-]+\\.)+[a-z]{2,6}$", "msg": "邮箱格式不正确", },
  "identity" : { "rgx": "^\\d{6}(18|19|20)?\\d{2}(0[1-9]|1[012])(0[1-9]|[12]\\d|3[01])\\d{3}(\\d|X|x)$", "msg": "身份证号码格式不正确",
  },
  "phone"    : { "rgx": "^(((13[0-9]{1})|(15[0-9]{1})|(18[0-9]{1})|(17[0-9]{1})|(14[0-9]{1}))+\\d{8})$", "msg": "必须是11位手机号码",
  },
}

// String转Int
func StrToInt(v string) int {
  s, _ := strconv.Atoi(v)
  return s
}

// String转Int64
func StrToInt64(v string) int64 {
  s, _ := strconv.ParseInt(v, 10, 64)
  return s
}


// Int转String
func IntToStr(v int) string {
  s := strconv.Itoa(v)
  return s
}

// Int64转String
func Int64ToStr(v int64) string {
  s := strconv.FormatInt(v, 10)
  return s
}

// ArrayIntToString
func ArrayInt64ToString(a interface{}) string {
  return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ",", -1), "[]")
}


// float64转String
func Float64ToStr(fv float64) string {
  return strconv.FormatFloat(fv, 'f', 0, 64)
}

// 查找数组并返回下标
func IndexOf(str []interface{}, data interface{}) int {
  for k, v := range str{
    if v == data {
      return k
    }
  }

  return - 1
}

// []StringTo[]interface
func ArrStrTointerface(data []string) []interface{} {
  newArr := make([]interface{}, len(data))
  for i, v := range data {
    newArr[i] = v
  }
  return newArr
}

func InitValidator(fieldInfo reflect.StructField, value interface{}) []map[string]interface{} {
  errMsgs := []map[string]interface{}{}
  tag := fieldInfo.Tag
  valid := tag.Get("validator")
  jsonKey := tag.Get("json")

  typ := reflect.TypeOf(value).String()

  if jsonKey == "" {
    jsonKey = fieldInfo.Name
  }

  // 为空的判断
  if value == "" {
    errMsgs = append(errMsgs, map[string]interface{}{
      "value": value,
      "msg":   jsonKey + " 不能为空",
    })
  } else if valid != "" {
    result := strings.Split(valid, ";")

    obj := make(map[string]interface{})

    if len(result) > 0 {
      for _, v := range result {
        result := strings.Split(v, ":")
        if len(result) == 2 {
          obj[result[0]] = result[1]
        }
      }
    }

    // 以下都是不为空的判断
    if obj["type"] != "" || obj["type"] != nil {
      rl := Rule[obj["type"].(string)]

      var bool bool
      if typ == "string" {
        bool = regexp.MustCompile(rl["rgx"].(string)).MatchString(fmt.Sprint(value.(string)))
      } else if typ == "int" {
        bool = regexp.MustCompile(rl["rgx"].(string)).MatchString(fmt.Sprint(value.(int)))
      } else if typ == "int64" {
        bool = regexp.MustCompile(rl["rgx"].(string)).MatchString(fmt.Sprint(value.(int64)))
      } else if typ == "float64" {
        bool = regexp.MustCompile(rl["rgx"].(string)).MatchString(fmt.Sprint(value.(float64)))
      }

      if bool != true {
        errMsgs = append(errMsgs, map[string]interface{}{
          "value": value,
          "msg":   jsonKey + " " + rl["msg"].(string),
        })
      }
    }
    if obj["min"] != nil {
      var bool bool
      if typ == "string" {
        bool = len(value.(string)) < StrToInt(obj["min"].(string))
      } else if typ == "int" {
        bool = len(IntToStr(value.(int))) < StrToInt(obj["min"].(string))
      } else if typ == "int64" {
        bool = len(Int64ToStr(value.(int64))) < StrToInt(obj["min"].(string))
      } else if typ == "float64" {
        bool = len(Float64ToStr(value.(float64))) < StrToInt(obj["min"].(string))
      }

      if bool == true {
        errMsgs = append(errMsgs, map[string]interface{}{
          "value": value,
          "msg":   jsonKey + " 不能小于" + obj["min"].(string) + "个字符",
        })
      }
    }

    if obj["max"] != nil {
      var bool bool
      if typ == "string" {
        bool = len(value.(string)) > StrToInt(obj["max"].(string))
      } else if typ == "int" {
        bool = len(IntToStr(value.(int))) > StrToInt(obj["max"].(string))
      } else if typ == "int64" {
        bool = len(Int64ToStr(value.(int64))) > StrToInt(obj["max"].(string))
      } else if typ == "float64" {
        bool = len(Float64ToStr(value.(float64))) > StrToInt(obj["max"].(string))
      }

      if bool == true {
        errMsgs = append(errMsgs, map[string]interface{}{
          "value": value,
          "msg":   jsonKey + " 最大不能超过" + obj["max"].(string) + "个字符",
        })

      }
    }

  }

  return errMsgs
}


func ResNodeEnvData(table interface{}) map[string]interface{} {
  v := reflect.ValueOf(table)
  v1 := v.Elem()
  k := v1.Type()

  errMsgs := []map[string]interface{}{}

  fliter := []interface{}{"int", "int64", "float64", "*int", "*int64", "*float64"}
  for i := 0; i < v1.NumField(); i++ {
    val := v1.Field(i)

    value := val.Interface()
    t1 := reflect.TypeOf(value)

    // 如果是结构体
    if t1.Kind() == reflect.Struct {
      v1 := reflect.ValueOf(value)

      for j := 0; j < t1.NumField(); j++ {

        tag := t1.Field(j).Tag
        valid := tag.Get("validator")
        json := tag.Get("json")
        if json == "" {
          json = k.Field(i).Name
        }

        // if v1.Field(j).Type().Kind() != reflect.Struct && valid != "" && IndexOf(ArrStrTointerface(required), json) != -1 || v1.Field(j).Type().Kind() != reflect.Struct && valid != "" {
        if v1.Field(j).Type().Kind() != reflect.Struct && valid != "" {
          // if v1.Field(j).Type().String() == "int" || v1.Field(j).Type().String() == "int64" || v1.Field(j).Type().String() == "float64" {
          if IndexOf(fliter, v1.Field(j).Type().String()) != -1 {
            var val int64 = 0
            if v1.Field(j).Type().String() == "*int64" {
              if v1.Field(j).Interface().(*int64) != nil {
                val = int64(*v1.Field(j).Interface().(*int64))
              }
            } else {
              val = v1.Field(j).Interface().(int64)
            }

            if val != 0 {
              res := InitValidator(t1.Field(j), val)
              for _, v := range res {
                errMsgs = append(errMsgs, v)
              }
            }
          } else {
            res := InitValidator(t1.Field(j), v1.Field(j).Interface())
            for _, v := range res {
              errMsgs = append(errMsgs, v)
            }
          }
        }
      }
    } else {
      tag := k.Field(i).Tag
      valid := tag.Get("validator")
      json := tag.Get("json")
      if json == "" {
        json = k.Field(i).Name
      }
      if valid != "" {
        // if t1.String() == "int" || t1.String() == "int64" || t1.String() == "float64" {
        if IndexOf(fliter, t1.String()) != -1{
          var val int64 = 0
          if t1.String() == "*int64" {
            if value.(*int64) != nil {
              val = int64(*value.(*int64))
            }
          } else {
            val = value.(int64)
          }

          if val != 0 {
            res := InitValidator(k.Field(i), val)
            for _, v := range res {
              errMsgs = append(errMsgs, v)
            }
          }
        } else {
          if t1.String() == "*string" {
            if value.(*string) != nil {
              value = string(*value.(*string))
            }
          }

          // if t1.String() == "*int64" {
            // val := *value.(*int64)
            // fmt.Println(int64(val))
          // }
          res := InitValidator(k.Field(i), value)
          for _, v := range res {
            errMsgs = append(errMsgs, v)
          }
        }

      }
    }
  }

  if len(errMsgs) > 0 {
    return map[string]interface{}{
      "status": 1,
      "data":   errMsgs,
    }
  }

  return map[string]interface{}{
    "status": 0,
    "data":   nil,
  }

}

type GeneratedResolver struct {
	DB              *DB
	EventController *events.EventController
}

func (r *GeneratedResolver) Mutation() MutationResolver {
	return &GeneratedMutationResolver{r}
}
func (r *GeneratedResolver) Query() QueryResolver {
	return &GeneratedQueryResolver{r}
}

func (r *GeneratedResolver) UserResultType() UserResultTypeResolver {
	return &GeneratedUserResultTypeResolver{r}
}

func (r *GeneratedResolver) User() UserResolver {
	return &GeneratedUserResolver{r}
}

func (r *GeneratedResolver) TaskResultType() TaskResultTypeResolver {
	return &GeneratedTaskResultTypeResolver{r}
}

func (r *GeneratedResolver) Task() TaskResolver {
	return &GeneratedTaskResolver{r}
}

type GeneratedMutationResolver struct{ *GeneratedResolver }

func (r *GeneratedMutationResolver) CreateUser(ctx context.Context, input map[string]interface{}) (item *User, err error) {
	principalID := getPrincipalID(ctx)
	now := time.Now()
	item = &User{ID: uuid.Must(uuid.NewV4()).String(), CreatedBy: principalID}

	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "User",
		EntityID:    item.ID,
		Date:        now.Unix(),
		PrincipalID: principalID,
	})

	var changes UserChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["email"]; ok && (item.Email != changes.Email) && (item.Email == nil || changes.Email == nil || *item.Email != *changes.Email) {
		item.Email = changes.Email
		event.AddNewValue("email", changes.Email)
	}

	if _, ok := input["age"]; ok && (item.Age != changes.Age) && (item.Age == nil || changes.Age == nil || *item.Age != *changes.Age) {
		item.Age = changes.Age
		event.AddNewValue("age", changes.Age)
	}

	if _, ok := input["firstName"]; ok && (item.FirstName != changes.FirstName) && (item.FirstName == nil || changes.FirstName == nil || *item.FirstName != *changes.FirstName) {
		item.FirstName = changes.FirstName
		event.AddNewValue("firstName", changes.FirstName)
	}

	if _, ok := input["lastName"]; ok && (item.LastName != changes.LastName) && (item.LastName == nil || changes.LastName == nil || *item.LastName != *changes.LastName) {
		item.LastName = changes.LastName
		event.AddNewValue("lastName", changes.LastName)
	}

	if ids, ok := input["tasksIds"].([]interface{}); ok {
		items := []Task{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Tasks")
		association.Replace(items)
	}

  envData := ResNodeEnvData(item)
  fmt.Println(envData)
  return

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		err = r.EventController.SendEvent(ctx, &event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateUser(ctx context.Context, id string, input map[string]interface{}) (item *User, err error) {
	principalID := getPrincipalID(ctx)
	item = &User{}
	now := time.Now()
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "User",
		EntityID:    item.ID,
		Date:        now.Unix(),
		PrincipalID: principalID,
	})

	var changes UserChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["email"]; ok && (item.Email != changes.Email) && (item.Email == nil || changes.Email == nil || *item.Email != *changes.Email) {
		event.AddOldValue("email", item.Email)
		event.AddNewValue("email", changes.Email)
		item.Email = changes.Email
	}

	if _, ok := input["age"]; ok && (item.Age != changes.Age) && (item.Age == nil || changes.Age == nil || *item.Age != *changes.Age) {
		event.AddOldValue("age", item.Age)
		event.AddNewValue("age", changes.Age)
		item.Age = changes.Age
	}

	if _, ok := input["firstName"]; ok && (item.FirstName != changes.FirstName) && (item.FirstName == nil || changes.FirstName == nil || *item.FirstName != *changes.FirstName) {
		event.AddOldValue("firstName", item.FirstName)
		event.AddNewValue("firstName", changes.FirstName)
		item.FirstName = changes.FirstName
	}

	if _, ok := input["lastName"]; ok && (item.LastName != changes.LastName) && (item.LastName == nil || changes.LastName == nil || *item.LastName != *changes.LastName) {
		event.AddOldValue("lastName", item.LastName)
		event.AddNewValue("lastName", changes.LastName)
		item.LastName = changes.LastName
	}

	if ids, ok := input["tasksIds"].([]interface{}); ok {
		items := []Task{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Tasks")
		association.Replace(items)
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		err = r.EventController.SendEvent(ctx, &event)
		data, _ := json.Marshal(event)
		fmt.Println("??", string(data))
	}

	return
}
func (r *GeneratedMutationResolver) DeleteUser(ctx context.Context, id string) (item *User, err error) {
	principalID := getPrincipalID(ctx)
	item = &User{}
	tx := r.DB.db.Begin()

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	// err = r.DB.Query().Delete(item, "id = ?", id).Error

	item.DeletedBy = principalID

	err = tx.Delete(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	return
}

func (r *GeneratedMutationResolver) CreateTask(ctx context.Context, input map[string]interface{}) (item *Task, err error) {
	principalID := getPrincipalID(ctx)
	now := time.Now()
	item = &Task{ID: uuid.Must(uuid.NewV4()).String(), CreatedBy: principalID}
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Task",
		EntityID:    item.ID,
		Date:        now.Unix(),
		PrincipalID: principalID,
	})

	var changes TaskChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["title"]; ok && (item.Title != changes.Title) && (item.Title == nil || changes.Title == nil || *item.Title != *changes.Title) {
		item.Title = changes.Title
		event.AddNewValue("title", changes.Title)
	}

	if _, ok := input["completed"]; ok && (item.Completed != changes.Completed) && (item.Completed == nil || changes.Completed == nil || *item.Completed != *changes.Completed) {
		item.Completed = changes.Completed
		event.AddNewValue("completed", changes.Completed)
	}

	if _, ok := input["dueDate"]; ok && (item.DueDate != changes.DueDate) && (item.DueDate == nil || changes.DueDate == nil || *item.DueDate != *changes.DueDate) {
		item.DueDate = changes.DueDate
		event.AddNewValue("dueDate", changes.DueDate)
	}

	if _, ok := input["assigneeId"]; ok && (item.AssigneeID != changes.AssigneeID) && (item.AssigneeID == nil || changes.AssigneeID == nil || *item.AssigneeID != *changes.AssigneeID) {
		item.AssigneeID = changes.AssigneeID
		event.AddNewValue("assigneeId", changes.AssigneeID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		err = r.EventController.SendEvent(ctx, &event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateTask(ctx context.Context, id string, input map[string]interface{}) (item *Task, err error) {
	principalID := getPrincipalID(ctx)
	item = &Task{}
	now := time.Now()
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Task",
		EntityID:    item.ID,
		Date:        now.Unix(),
		PrincipalID: principalID,
	})

	var changes TaskChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["title"]; ok && (item.Title != changes.Title) && (item.Title == nil || changes.Title == nil || *item.Title != *changes.Title) {
		event.AddOldValue("title", item.Title)
		event.AddNewValue("title", changes.Title)
		item.Title = changes.Title
	}

	if _, ok := input["completed"]; ok && (item.Completed != changes.Completed) && (item.Completed == nil || changes.Completed == nil || *item.Completed != *changes.Completed) {
		event.AddOldValue("completed", item.Completed)
		event.AddNewValue("completed", changes.Completed)
		item.Completed = changes.Completed
	}

	if _, ok := input["dueDate"]; ok && (item.DueDate != changes.DueDate) && (item.DueDate == nil || changes.DueDate == nil || *item.DueDate != *changes.DueDate) {
		event.AddOldValue("dueDate", item.DueDate)
		event.AddNewValue("dueDate", changes.DueDate)
		item.DueDate = changes.DueDate
	}

	if _, ok := input["assigneeId"]; ok && (item.AssigneeID != changes.AssigneeID) && (item.AssigneeID == nil || changes.AssigneeID == nil || *item.AssigneeID != *changes.AssigneeID) {
		event.AddOldValue("assigneeId", item.AssigneeID)
		event.AddNewValue("assigneeId", changes.AssigneeID)
		item.AssigneeID = changes.AssigneeID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		err = r.EventController.SendEvent(ctx, &event)
		data, _ := json.Marshal(event)
		fmt.Println("??", string(data))
	}

	return
}
func (r *GeneratedMutationResolver) DeleteTask(ctx context.Context, id string) (item *Task, err error) {
	principalID := getPrincipalID(ctx)
	item = &Task{}
	tx := r.DB.db.Begin()

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	// err = r.DB.Query().Delete(item, "id = ?", id).Error

	item.DeletedBy = principalID

	err = tx.Delete(item).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	return
}

type GeneratedQueryResolver struct{ *GeneratedResolver }

func (r *GeneratedQueryResolver) User(ctx context.Context, id *string, q *string) (*User, error) {
	t := User{}
	err := resolvers.GetItem(ctx, r.DB.Query(), &t, id)
	return &t, err
}
func (r *GeneratedQueryResolver) Users(ctx context.Context, offset *int, limit *int, q *string, sort []UserSortType, filter *UserFilterType) (*UserResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range sort {
		_sort = append(_sort, s)
	}
	query := UserQueryFilter{q}
	return &UserResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset: offset,
			Limit:  limit,
			Query:  &query,
			Sort:   _sort,
			Filter: filter,
		},
	}, nil
}

type GeneratedUserResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedUserResultTypeResolver) Items(ctx context.Context, obj *UserResultType) (items []*User, err error) {
	err = obj.GetItems(ctx, r.DB.db, "users", &items)
	return
}

func (r *GeneratedUserResultTypeResolver) Count(ctx context.Context, obj *UserResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &User{})
}

type GeneratedUserResolver struct{ *GeneratedResolver }

func (r *GeneratedUserResolver) Tasks(ctx context.Context, obj *User) (res []*Task, err error) {

	items := []*Task{}
	err = r.DB.Query().Model(obj).Related(&items, "Tasks").Error
	res = items

	return
}

func (r *GeneratedQueryResolver) Task(ctx context.Context, id *string, q *string) (*Task, error) {
	t := Task{}
	err := resolvers.GetItem(ctx, r.DB.Query(), &t, id)
	return &t, err
}
func (r *GeneratedQueryResolver) Tasks(ctx context.Context, offset *int, limit *int, q *string, sort []TaskSortType, filter *TaskFilterType) (*TaskResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range sort {
		_sort = append(_sort, s)
	}
	query := TaskQueryFilter{q}
	return &TaskResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset: offset,
			Limit:  limit,
			Query:  &query,
			Sort:   _sort,
			Filter: filter,
		},
	}, nil
}

type GeneratedTaskResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedTaskResultTypeResolver) Items(ctx context.Context, obj *TaskResultType) (items []*Task, err error) {
	err = obj.GetItems(ctx, r.DB.db, "tasks", &items)
	return
}

func (r *GeneratedTaskResultTypeResolver) Count(ctx context.Context, obj *TaskResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &Task{})
}

type GeneratedTaskResolver struct{ *GeneratedResolver }

func (r *GeneratedTaskResolver) Assignee(ctx context.Context, obj *Task) (res *User, err error) {

	item := User{}
	_res := r.DB.Query().Model(obj).Related(&item, "Assignee")
	if _res.RecordNotFound() {
		return
	} else {
		err = _res.Error
	}
	res = &item

	return
}
