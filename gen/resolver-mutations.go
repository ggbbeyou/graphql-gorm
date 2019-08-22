package gen

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/maiguangyang/graphql-gorm/utils"
	"github.com/maiguangyang/graphql/events"
	"github.com/maiguangyang/graphql/resolvers"
)

type GeneratedMutationResolver struct{ *GeneratedResolver }

func (r *GeneratedMutationResolver) CreateUser(ctx context.Context, input map[string]interface{}) (item *User, err error) {
	return r.Handlers.CreateUser(ctx, r.GeneratedResolver, input)
}
func CreateUserHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *User, err error) {
	principalID := getPrincipalIDFromContext(ctx)
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

	if input["state"] == nil {
		input["state"] = 1
	}

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

	if _, ok := input["state"]; ok && (item.State != changes.State) && (item.State == nil || changes.State == nil || *item.State != *changes.State) {
		item.State = changes.State
		event.AddNewValue("state", changes.State)
	}

	errText, resErr := utils.Validator(item)
	if resErr != nil {
		return item, &errText
	}

	if err = tx.Create(item).Error; err != nil {
		tx.Rollback()
		return
	}

	if ids, ok := input["tasksIds"].([]interface{}); ok {
		items := []Task{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Tasks")
		association.Replace(items)
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
	return r.Handlers.UpdateUser(ctx, r.GeneratedResolver, id, input)
}
func UpdateUserHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *User, err error) {
	principalID := getPrincipalIDFromContext(ctx)
	item = &User{}
	now := time.Now()
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "User",
		EntityID:    id,
		Date:        now.Unix(),
		PrincipalID: principalID,
	})

	if input["state"] == nil {
		input["state"] = 1
	}

	var changes UserChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

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

	if _, ok := input["state"]; ok && (item.State != changes.State) && (item.State == nil || changes.State == nil || *item.State != *changes.State) {
		event.AddOldValue("state", item.State)
		event.AddNewValue("state", changes.State)
		item.State = changes.State
	}

	errText, resErr := utils.Validator(item)
	if resErr != nil {
		return item, &errText
	}

	oldItem := &User{}
	err = resolvers.GetItem(ctx, tx, oldItem, &id)
	if err != nil {
		return oldItem, err
	}

	oldState := oldItem.State

	item.UpdatedBy = principalID
	item.ID = id

	if err = tx.Model(&item).Updates(item).Error; err != nil {
		tx.Rollback()
		return
	}

	items := []Task{}
	if ids, ok := input["tasksIds"].([]interface{}); ok {
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Tasks")
		association.Replace(items)
	}

	// 判断是不是改变状态
	if oldState != item.State {
		if err = tx.Model(&items).Where("assigneeId = ?", item.ID).Update("state", item.State).Error; err != nil {
			tx.Rollback()
			return
		}
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
	return r.Handlers.DeleteUser(ctx, r.GeneratedResolver, id)
}
func DeleteUserHandler(ctx context.Context, r *GeneratedResolver, id string) (item *User, err error) {
	principalID := getPrincipalIDFromContext(ctx)
	item = &User{}
	now := time.Now()
	tx := r.DB.db.Begin()

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	// 3为删除
	var state int64 = 3

	item.UpdatedBy = principalID
	item.State = &state

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "User",
		EntityID:    id,
		Date:        now.Unix(),
		PrincipalID: principalID,
	})

	// err = tx.Delete(item, "users.id = ?", id).Error

	if err = tx.Save(item).Error; err != nil {
		tx.Rollback()
		return
	}

	items := []Task{}
	if err = tx.Model(&items).Where("assigneeId = ?", id).Update("state", state).Error; err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = r.EventController.SendEvent(ctx, &event)

	return
}
func (r *GeneratedMutationResolver) DeleteAllUsers(ctx context.Context) (bool, error) {
	return r.Handlers.DeleteAllUsers(ctx, r.GeneratedResolver)
}
func DeleteAllUsersHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	err := r.DB.db.Delete(&User{}).Error
	return err == nil, err
}

func (r *GeneratedMutationResolver) CreateTask(ctx context.Context, input map[string]interface{}) (item *Task, err error) {
	return r.Handlers.CreateTask(ctx, r.GeneratedResolver, input)
}
func CreateTaskHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Task, err error) {
	principalID := getPrincipalIDFromContext(ctx)
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

	if input["state"] == nil {
		input["state"] = 1
	}

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

	if _, ok := input["state"]; ok && (item.State != changes.State) && (item.State == nil || changes.State == nil || *item.State != *changes.State) {
		item.State = changes.State
		event.AddNewValue("state", changes.State)
	}

	errText, resErr := utils.Validator(item)
	if resErr != nil {
		return item, &errText
	}

	if err = tx.Create(item).Error; err != nil {
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
	return r.Handlers.UpdateTask(ctx, r.GeneratedResolver, id, input)
}
func UpdateTaskHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Task, err error) {
	principalID := getPrincipalIDFromContext(ctx)
	item = &Task{}
	now := time.Now()
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Task",
		EntityID:    id,
		Date:        now.Unix(),
		PrincipalID: principalID,
	})

	if input["state"] == nil {
		input["state"] = 1
	}

	var changes TaskChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

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

	if _, ok := input["state"]; ok && (item.State != changes.State) && (item.State == nil || changes.State == nil || *item.State != *changes.State) {
		event.AddOldValue("state", item.State)
		event.AddNewValue("state", changes.State)
		item.State = changes.State
	}

	errText, resErr := utils.Validator(item)
	if resErr != nil {
		return item, &errText
	}

	oldItem := &Task{}
	err = resolvers.GetItem(ctx, tx, oldItem, &id)
	if err != nil {
		return oldItem, err
	}

	item.UpdatedBy = principalID
	item.ID = id

	if err = tx.Model(&item).Updates(item).Error; err != nil {
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
	return r.Handlers.DeleteTask(ctx, r.GeneratedResolver, id)
}
func DeleteTaskHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Task, err error) {
	principalID := getPrincipalIDFromContext(ctx)
	item = &Task{}
	now := time.Now()
	tx := r.DB.db.Begin()

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	// 3为删除
	var state int64 = 3

	item.UpdatedBy = principalID
	item.State = &state

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Task",
		EntityID:    id,
		Date:        now.Unix(),
		PrincipalID: principalID,
	})

	// err = tx.Delete(item, "tasks.id = ?", id).Error

	if err = tx.Save(item).Error; err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = r.EventController.SendEvent(ctx, &event)

	return
}
func (r *GeneratedMutationResolver) DeleteAllTasks(ctx context.Context) (bool, error) {
	return r.Handlers.DeleteAllTasks(ctx, r.GeneratedResolver)
}
func DeleteAllTasksHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	err := r.DB.db.Delete(&Task{}).Error
	return err == nil, err
}
