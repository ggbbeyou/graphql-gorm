package gen

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/maiguangyang/graphql/events"
	"github.com/maiguangyang/graphql/resolvers"
	uuid "github.com/satori/go.uuid"
)

func getPrincipalID(ctx context.Context) *string {
	v, _ := ctx.Value(KeyPrincipalID).(*string)
	return v
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
		// data, _ := json.Marshal(event)
		// fmt.Println("?",string(data))
	}

	return
}
func (r *GeneratedMutationResolver) DeleteUser(ctx context.Context, id string) (item *User, err error) {
	item = &User{}
	err = resolvers.GetItem(ctx, r.DB.Query(), item, &id)
	if err != nil {
		return
	}

	err = r.DB.Query().Delete(item, "users.id = ?", id).Error

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
		// data, _ := json.Marshal(event)
		// fmt.Println("?",string(data))
	}

	return
}
func (r *GeneratedMutationResolver) DeleteTask(ctx context.Context, id string) (item *Task, err error) {
	item = &Task{}
	err = resolvers.GetItem(ctx, r.DB.Query(), item, &id)
	if err != nil {
		return
	}

	err = r.DB.Query().Delete(item, "tasks.id = ?", id).Error

	return
}

type GeneratedQueryResolver struct{ *GeneratedResolver }

func (r *GeneratedQueryResolver) User(ctx context.Context, id *string, q *string, filter *UserFilterType) (*User, error) {
	query := UserQueryFilter{q}
	offset := 0
	limit := 1
	current_page := 0
	per_page := 0
	rt := &UserResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset:      &offset,
			Limit:       &limit,
			CurrentPage: &current_page,
			PerPage:     &per_page,
			Query:       &query,
			Filter:      filter,
		},
	}
	qb := r.DB.Query()
	if id != nil {
		qb = qb.Where("users.id = ?", *id)
	}

	var items []*User
	err := rt.GetItems(ctx, qb, "users", &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, fmt.Errorf("record not found")
	}
	return items[0], err
}
func (r *GeneratedQueryResolver) Users(ctx context.Context, offset *int, limit *int, current_page *int, per_page *int, q *string, sort []UserSortType, filter *UserFilterType) (*UserResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range sort {
		_sort = append(_sort, s)
	}
	query := UserQueryFilter{q}
	return &UserResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset:      offset,
			Limit:       limit,
			CurrentPage: current_page,
			PerPage:     per_page,
			Query:       &query,
			Sort:        _sort,
			Filter:      filter,
		},
	}, nil
}

type GeneratedUserResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedUserResultTypeResolver) Data(ctx context.Context, obj *UserResultType) (items []*User, err error) {
	err = obj.GetItems(ctx, r.DB.db, "users", &items)
	return
}

func (r *GeneratedUserResultTypeResolver) Total(ctx context.Context, obj *UserResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &User{})
}

func (r *GeneratedUserResultTypeResolver) CurrentPage(ctx context.Context, obj *UserResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &User{})
}

func (r *GeneratedUserResultTypeResolver) PerPage(ctx context.Context, obj *UserResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &User{})
}

func (r *GeneratedUserResultTypeResolver) TotalPage(ctx context.Context, obj *UserResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &User{})
}

type GeneratedUserResolver struct{ *GeneratedResolver }

func (r *GeneratedUserResolver) Tasks(ctx context.Context, obj *User) (res []*Task, err error) {

	selects := resolvers.GetFieldsRequested(ctx, strings.ToLower("Tasks"))

	items := []*Task{}
	err = r.DB.Query().Select(selects).Model(obj).Related(&items, "Tasks").Error
	res = items

	return
}

func (r *GeneratedQueryResolver) Task(ctx context.Context, id *string, q *string, filter *TaskFilterType) (*Task, error) {
	query := TaskQueryFilter{q}
	offset := 0
	limit := 1
	current_page := 0
	per_page := 0
	rt := &TaskResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset:      &offset,
			Limit:       &limit,
			CurrentPage: &current_page,
			PerPage:     &per_page,
			Query:       &query,
			Filter:      filter,
		},
	}
	qb := r.DB.Query()
	if id != nil {
		qb = qb.Where("tasks.id = ?", *id)
	}

	var items []*Task
	err := rt.GetItems(ctx, qb, "tasks", &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, fmt.Errorf("record not found")
	}
	return items[0], err
}
func (r *GeneratedQueryResolver) Tasks(ctx context.Context, offset *int, limit *int, current_page *int, per_page *int, q *string, sort []TaskSortType, filter *TaskFilterType) (*TaskResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range sort {
		_sort = append(_sort, s)
	}
	query := TaskQueryFilter{q}
	return &TaskResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset:      offset,
			Limit:       limit,
			CurrentPage: current_page,
			PerPage:     per_page,
			Query:       &query,
			Sort:        _sort,
			Filter:      filter,
		},
	}, nil
}

type GeneratedTaskResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedTaskResultTypeResolver) Data(ctx context.Context, obj *TaskResultType) (items []*Task, err error) {
	err = obj.GetItems(ctx, r.DB.db, "tasks", &items)
	return
}

func (r *GeneratedTaskResultTypeResolver) Total(ctx context.Context, obj *TaskResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &Task{})
}

func (r *GeneratedTaskResultTypeResolver) CurrentPage(ctx context.Context, obj *TaskResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &Task{})
}

func (r *GeneratedTaskResultTypeResolver) PerPage(ctx context.Context, obj *TaskResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &Task{})
}

func (r *GeneratedTaskResultTypeResolver) TotalPage(ctx context.Context, obj *TaskResultType) (count int, err error) {
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
