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

	if input["del"] == nil {
		input["del"] = 1
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

	if _, ok := input["del"]; ok && (item.Del != changes.Del) && (item.Del == nil || changes.Del == nil || *item.Del != *changes.Del) {
		item.Del = changes.Del
		event.AddNewValue("del", changes.Del)
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

		for k, _ := range items {
			items[k].State = item.State
			items[k].Del = item.Del
		}

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

	if input["del"] == nil {
		input["del"] = 1
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

	if _, ok := input["del"]; ok && (item.Del != changes.Del) && (item.Del == nil || changes.Del == nil || *item.Del != *changes.Del) {
		event.AddOldValue("del", item.Del)
		event.AddNewValue("del", changes.Del)
		item.Del = changes.Del
	}

	errText, resErr := utils.Validator(item)
	if resErr != nil {
		return item, &errText
	}

	item.UpdatedBy = principalID
	item.ID = id

	if err = tx.Model(&item).Updates(item).Error; err != nil {
		tx.Rollback()
		return
	}

	if ids, ok := input["tasksIds"].([]interface{}); ok {
		items := []Task{}
		tx.Find(&items, "id IN (?)", ids)

		for k, _ := range items {
			items[k].State = item.State
			items[k].Del = item.Del
		}

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

	// 2为删除
	var del int64 = 2

	item.UpdatedBy = principalID
	item.Del = &del

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

	tasks := []Task{}
	if err = tx.Model(&tasks).Where("assigneeId = ?", id).Update("del", del).Error; err != nil {
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

	if input["del"] == nil {
		input["del"] = 1
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

	if _, ok := input["del"]; ok && (item.Del != changes.Del) && (item.Del == nil || changes.Del == nil || *item.Del != *changes.Del) {
		item.Del = changes.Del
		event.AddNewValue("del", changes.Del)
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

	if input["del"] == nil {
		input["del"] = 1
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

	if _, ok := input["del"]; ok && (item.Del != changes.Del) && (item.Del == nil || changes.Del == nil || *item.Del != *changes.Del) {
		event.AddOldValue("del", item.Del)
		event.AddNewValue("del", changes.Del)
		item.Del = changes.Del
	}

	errText, resErr := utils.Validator(item)
	if resErr != nil {
		return item, &errText
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

	// 2为删除
	var del int64 = 2

	item.UpdatedBy = principalID
	item.Del = &del

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

func (r *GeneratedMutationResolver) CreateAdmin(ctx context.Context, input map[string]interface{}) (item *Admin, err error) {
	return r.Handlers.CreateAdmin(ctx, r.GeneratedResolver, input)
}
func CreateAdminHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Admin, err error) {
	principalID := getPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Admin{ID: uuid.Must(uuid.NewV4()).String(), CreatedBy: principalID}
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Admin",
		EntityID:    item.ID,
		Date:        now.Unix(),
		PrincipalID: principalID,
	})

	if input["state"] == nil {
		input["state"] = 1
	}

	if input["del"] == nil {
		input["del"] = 1
	}

	var changes AdminChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["Phone"]; ok && (item.Phone != changes.Phone) && (item.Phone == nil || changes.Phone == nil || *item.Phone != *changes.Phone) {
		item.Phone = changes.Phone
		event.AddNewValue("Phone", changes.Phone)
	}

	if _, ok := input["Password"]; ok && (item.Password != changes.Password) && (item.Password == nil || changes.Password == nil || *item.Password != *changes.Password) {
		item.Password = changes.Password
		event.AddNewValue("Password", changes.Password)
	}

	if _, ok := input["Username"]; ok && (item.Username != changes.Username) && (item.Username == nil || changes.Username == nil || *item.Username != *changes.Username) {
		item.Username = changes.Username
		event.AddNewValue("Username", changes.Username)
	}

	if _, ok := input["Money"]; ok && (item.Money != changes.Money) && (item.Money == nil || changes.Money == nil || *item.Money != *changes.Money) {
		item.Money = changes.Money
		event.AddNewValue("Money", changes.Money)
	}

	if _, ok := input["Sex"]; ok && (item.Sex != changes.Sex) && (item.Sex == nil || changes.Sex == nil || *item.Sex != *changes.Sex) {
		item.Sex = changes.Sex
		event.AddNewValue("Sex", changes.Sex)
	}

	if _, ok := input["Super"]; ok && (item.Super != changes.Super) && (item.Super == nil || changes.Super == nil || *item.Super != *changes.Super) {
		item.Super = changes.Super
		event.AddNewValue("Super", changes.Super)
	}

	if _, ok := input["LoginCount"]; ok && (item.LoginCount != changes.LoginCount) && (item.LoginCount == nil || changes.LoginCount == nil || *item.LoginCount != *changes.LoginCount) {
		item.LoginCount = changes.LoginCount
		event.AddNewValue("LoginCount", changes.LoginCount)
	}

	if _, ok := input["LoginIp"]; ok && (item.LoginIP != changes.LoginIP) && (item.LoginIP == nil || changes.LoginIP == nil || *item.LoginIP != *changes.LoginIP) {
		item.LoginIP = changes.LoginIP
		event.AddNewValue("LoginIp", changes.LoginIP)
	}

	if _, ok := input["LastIp"]; ok && (item.LastIP != changes.LastIP) && (item.LastIP == nil || changes.LastIP == nil || *item.LastIP != *changes.LastIP) {
		item.LastIP = changes.LastIP
		event.AddNewValue("LastIp", changes.LastIP)
	}

	if _, ok := input["state"]; ok && (item.State != changes.State) && (item.State == nil || changes.State == nil || *item.State != *changes.State) {
		item.State = changes.State
		event.AddNewValue("state", changes.State)
	}

	if _, ok := input["del"]; ok && (item.Del != changes.Del) && (item.Del == nil || changes.Del == nil || *item.Del != *changes.Del) {
		item.Del = changes.Del
		event.AddNewValue("del", changes.Del)
	}

	errText, resErr := utils.Validator(item)
	if resErr != nil {
		return item, &errText
	}

	if err = tx.Create(item).Error; err != nil {
		tx.Rollback()
		return
	}

	if ids, ok := input["groupsIds"].([]interface{}); ok {
		items := []Group{}
		tx.Find(&items, "id IN (?)", ids)

		for k, _ := range items {
			items[k].State = item.State
			items[k].Del = item.Del
		}

		association := tx.Model(&item).Association("Groups")
		association.Replace(items)
	}

	if ids, ok := input["rolesIds"].([]interface{}); ok {
		items := []Role{}
		tx.Find(&items, "id IN (?)", ids)

		for k, _ := range items {
			items[k].State = item.State
			items[k].Del = item.Del
		}

		association := tx.Model(&item).Association("Roles")
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
func (r *GeneratedMutationResolver) UpdateAdmin(ctx context.Context, id string, input map[string]interface{}) (item *Admin, err error) {
	return r.Handlers.UpdateAdmin(ctx, r.GeneratedResolver, id, input)
}
func UpdateAdminHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Admin, err error) {
	principalID := getPrincipalIDFromContext(ctx)
	item = &Admin{}
	now := time.Now()
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Admin",
		EntityID:    id,
		Date:        now.Unix(),
		PrincipalID: principalID,
	})

	if input["state"] == nil {
		input["state"] = 1
	}

	if input["del"] == nil {
		input["del"] = 1
	}

	var changes AdminChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

	if _, ok := input["Phone"]; ok && (item.Phone != changes.Phone) && (item.Phone == nil || changes.Phone == nil || *item.Phone != *changes.Phone) {
		event.AddOldValue("Phone", item.Phone)
		event.AddNewValue("Phone", changes.Phone)
		item.Phone = changes.Phone
	}

	if _, ok := input["Password"]; ok && (item.Password != changes.Password) && (item.Password == nil || changes.Password == nil || *item.Password != *changes.Password) {
		event.AddOldValue("Password", item.Password)
		event.AddNewValue("Password", changes.Password)
		item.Password = changes.Password
	}

	if _, ok := input["Username"]; ok && (item.Username != changes.Username) && (item.Username == nil || changes.Username == nil || *item.Username != *changes.Username) {
		event.AddOldValue("Username", item.Username)
		event.AddNewValue("Username", changes.Username)
		item.Username = changes.Username
	}

	if _, ok := input["Money"]; ok && (item.Money != changes.Money) && (item.Money == nil || changes.Money == nil || *item.Money != *changes.Money) {
		event.AddOldValue("Money", item.Money)
		event.AddNewValue("Money", changes.Money)
		item.Money = changes.Money
	}

	if _, ok := input["Sex"]; ok && (item.Sex != changes.Sex) && (item.Sex == nil || changes.Sex == nil || *item.Sex != *changes.Sex) {
		event.AddOldValue("Sex", item.Sex)
		event.AddNewValue("Sex", changes.Sex)
		item.Sex = changes.Sex
	}

	if _, ok := input["Super"]; ok && (item.Super != changes.Super) && (item.Super == nil || changes.Super == nil || *item.Super != *changes.Super) {
		event.AddOldValue("Super", item.Super)
		event.AddNewValue("Super", changes.Super)
		item.Super = changes.Super
	}

	if _, ok := input["LoginCount"]; ok && (item.LoginCount != changes.LoginCount) && (item.LoginCount == nil || changes.LoginCount == nil || *item.LoginCount != *changes.LoginCount) {
		event.AddOldValue("LoginCount", item.LoginCount)
		event.AddNewValue("LoginCount", changes.LoginCount)
		item.LoginCount = changes.LoginCount
	}

	if _, ok := input["LoginIp"]; ok && (item.LoginIP != changes.LoginIP) && (item.LoginIP == nil || changes.LoginIP == nil || *item.LoginIP != *changes.LoginIP) {
		event.AddOldValue("LoginIp", item.LoginIP)
		event.AddNewValue("LoginIp", changes.LoginIP)
		item.LoginIP = changes.LoginIP
	}

	if _, ok := input["LastIp"]; ok && (item.LastIP != changes.LastIP) && (item.LastIP == nil || changes.LastIP == nil || *item.LastIP != *changes.LastIP) {
		event.AddOldValue("LastIp", item.LastIP)
		event.AddNewValue("LastIp", changes.LastIP)
		item.LastIP = changes.LastIP
	}

	if _, ok := input["state"]; ok && (item.State != changes.State) && (item.State == nil || changes.State == nil || *item.State != *changes.State) {
		event.AddOldValue("state", item.State)
		event.AddNewValue("state", changes.State)
		item.State = changes.State
	}

	if _, ok := input["del"]; ok && (item.Del != changes.Del) && (item.Del == nil || changes.Del == nil || *item.Del != *changes.Del) {
		event.AddOldValue("del", item.Del)
		event.AddNewValue("del", changes.Del)
		item.Del = changes.Del
	}

	errText, resErr := utils.Validator(item)
	if resErr != nil {
		return item, &errText
	}

	item.UpdatedBy = principalID
	item.ID = id

	if err = tx.Model(&item).Updates(item).Error; err != nil {
		tx.Rollback()
		return
	}

	if ids, ok := input["groupsIds"].([]interface{}); ok {
		items := []Group{}
		tx.Find(&items, "id IN (?)", ids)

		for k, _ := range items {
			items[k].State = item.State
			items[k].Del = item.Del
		}

		association := tx.Model(&item).Association("Groups")
		association.Replace(items)
	}

	if ids, ok := input["rolesIds"].([]interface{}); ok {
		items := []Role{}
		tx.Find(&items, "id IN (?)", ids)

		for k, _ := range items {
			items[k].State = item.State
			items[k].Del = item.Del
		}

		association := tx.Model(&item).Association("Roles")
		association.Replace(items)
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
func (r *GeneratedMutationResolver) DeleteAdmin(ctx context.Context, id string) (item *Admin, err error) {
	return r.Handlers.DeleteAdmin(ctx, r.GeneratedResolver, id)
}
func DeleteAdminHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Admin, err error) {
	principalID := getPrincipalIDFromContext(ctx)
	item = &Admin{}
	now := time.Now()
	tx := r.DB.db.Begin()

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	// 2为删除
	var del int64 = 2

	item.UpdatedBy = principalID
	item.Del = &del

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Admin",
		EntityID:    id,
		Date:        now.Unix(),
		PrincipalID: principalID,
	})

	// err = tx.Delete(item, "admins.id = ?", id).Error

	if err = tx.Save(item).Error; err != nil {
		tx.Rollback()
		return
	}

	groups := []Group{}
	if err = tx.Model(&groups).Where("adminId = ?", id).Update("del", del).Error; err != nil {
		tx.Rollback()
		return
	}

	roles := []Role{}
	if err = tx.Model(&roles).Where("adminId = ?", id).Update("del", del).Error; err != nil {
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
func (r *GeneratedMutationResolver) DeleteAllAdmins(ctx context.Context) (bool, error) {
	return r.Handlers.DeleteAllAdmins(ctx, r.GeneratedResolver)
}
func DeleteAllAdminsHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	err := r.DB.db.Delete(&Admin{}).Error
	return err == nil, err
}

func (r *GeneratedMutationResolver) CreateGroup(ctx context.Context, input map[string]interface{}) (item *Group, err error) {
	return r.Handlers.CreateGroup(ctx, r.GeneratedResolver, input)
}
func CreateGroupHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Group, err error) {
	principalID := getPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Group{ID: uuid.Must(uuid.NewV4()).String(), CreatedBy: principalID}
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Group",
		EntityID:    item.ID,
		Date:        now.Unix(),
		PrincipalID: principalID,
	})

	if input["state"] == nil {
		input["state"] = 1
	}

	if input["del"] == nil {
		input["del"] = 1
	}

	var changes GroupChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["Name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		item.Name = changes.Name
		event.AddNewValue("Name", changes.Name)
	}

	if _, ok := input["state"]; ok && (item.State != changes.State) && (item.State == nil || changes.State == nil || *item.State != *changes.State) {
		item.State = changes.State
		event.AddNewValue("state", changes.State)
	}

	if _, ok := input["del"]; ok && (item.Del != changes.Del) && (item.Del == nil || changes.Del == nil || *item.Del != *changes.Del) {
		item.Del = changes.Del
		event.AddNewValue("del", changes.Del)
	}

	errText, resErr := utils.Validator(item)
	if resErr != nil {
		return item, &errText
	}

	if err = tx.Create(item).Error; err != nil {
		tx.Rollback()
		return
	}

	if ids, ok := input["adminIds"].([]interface{}); ok {
		items := []Admin{}
		tx.Find(&items, "id IN (?)", ids)

		for k, _ := range items {
			items[k].State = item.State
			items[k].Del = item.Del
		}

		association := tx.Model(&item).Association("Admin")
		association.Replace(items)
	}

	if ids, ok := input["rolesIds"].([]interface{}); ok {
		items := []Role{}
		tx.Find(&items, "id IN (?)", ids)

		for k, _ := range items {
			items[k].State = item.State
			items[k].Del = item.Del
		}

		association := tx.Model(&item).Association("Roles")
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
func (r *GeneratedMutationResolver) UpdateGroup(ctx context.Context, id string, input map[string]interface{}) (item *Group, err error) {
	return r.Handlers.UpdateGroup(ctx, r.GeneratedResolver, id, input)
}
func UpdateGroupHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Group, err error) {
	principalID := getPrincipalIDFromContext(ctx)
	item = &Group{}
	now := time.Now()
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Group",
		EntityID:    id,
		Date:        now.Unix(),
		PrincipalID: principalID,
	})

	if input["state"] == nil {
		input["state"] = 1
	}

	if input["del"] == nil {
		input["del"] = 1
	}

	var changes GroupChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

	if _, ok := input["Name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		event.AddOldValue("Name", item.Name)
		event.AddNewValue("Name", changes.Name)
		item.Name = changes.Name
	}

	if _, ok := input["state"]; ok && (item.State != changes.State) && (item.State == nil || changes.State == nil || *item.State != *changes.State) {
		event.AddOldValue("state", item.State)
		event.AddNewValue("state", changes.State)
		item.State = changes.State
	}

	if _, ok := input["del"]; ok && (item.Del != changes.Del) && (item.Del == nil || changes.Del == nil || *item.Del != *changes.Del) {
		event.AddOldValue("del", item.Del)
		event.AddNewValue("del", changes.Del)
		item.Del = changes.Del
	}

	errText, resErr := utils.Validator(item)
	if resErr != nil {
		return item, &errText
	}

	item.UpdatedBy = principalID
	item.ID = id

	if err = tx.Model(&item).Updates(item).Error; err != nil {
		tx.Rollback()
		return
	}

	if ids, ok := input["adminIds"].([]interface{}); ok {
		items := []Admin{}
		tx.Find(&items, "id IN (?)", ids)

		for k, _ := range items {
			items[k].State = item.State
			items[k].Del = item.Del
		}

		association := tx.Model(&item).Association("Admin")
		association.Replace(items)
	}

	if ids, ok := input["rolesIds"].([]interface{}); ok {
		items := []Role{}
		tx.Find(&items, "id IN (?)", ids)

		for k, _ := range items {
			items[k].State = item.State
			items[k].Del = item.Del
		}

		association := tx.Model(&item).Association("Roles")
		association.Replace(items)
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
func (r *GeneratedMutationResolver) DeleteGroup(ctx context.Context, id string) (item *Group, err error) {
	return r.Handlers.DeleteGroup(ctx, r.GeneratedResolver, id)
}
func DeleteGroupHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Group, err error) {
	principalID := getPrincipalIDFromContext(ctx)
	item = &Group{}
	now := time.Now()
	tx := r.DB.db.Begin()

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	// 2为删除
	var del int64 = 2

	item.UpdatedBy = principalID
	item.Del = &del

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Group",
		EntityID:    id,
		Date:        now.Unix(),
		PrincipalID: principalID,
	})

	// err = tx.Delete(item, "groups.id = ?", id).Error

	if err = tx.Save(item).Error; err != nil {
		tx.Rollback()
		return
	}

	admin := []Admin{}
	if err = tx.Model(&admin).Where("groupsId = ?", id).Update("del", del).Error; err != nil {
		tx.Rollback()
		return
	}

	roles := []Role{}
	if err = tx.Model(&roles).Where("groupId = ?", id).Update("del", del).Error; err != nil {
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
func (r *GeneratedMutationResolver) DeleteAllGroups(ctx context.Context) (bool, error) {
	return r.Handlers.DeleteAllGroups(ctx, r.GeneratedResolver)
}
func DeleteAllGroupsHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	err := r.DB.db.Delete(&Group{}).Error
	return err == nil, err
}

func (r *GeneratedMutationResolver) CreateRole(ctx context.Context, input map[string]interface{}) (item *Role, err error) {
	return r.Handlers.CreateRole(ctx, r.GeneratedResolver, input)
}
func CreateRoleHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Role, err error) {
	principalID := getPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Role{ID: uuid.Must(uuid.NewV4()).String(), CreatedBy: principalID}
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Role",
		EntityID:    item.ID,
		Date:        now.Unix(),
		PrincipalID: principalID,
	})

	if input["state"] == nil {
		input["state"] = 1
	}

	if input["del"] == nil {
		input["del"] = 1
	}

	var changes RoleChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["Name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		item.Name = changes.Name
		event.AddNewValue("Name", changes.Name)
	}

	if _, ok := input["Pid"]; ok && (item.Pid != changes.Pid) && (item.Pid == nil || changes.Pid == nil || *item.Pid != *changes.Pid) {
		item.Pid = changes.Pid
		event.AddNewValue("Pid", changes.Pid)
	}

	if _, ok := input["state"]; ok && (item.State != changes.State) && (item.State == nil || changes.State == nil || *item.State != *changes.State) {
		item.State = changes.State
		event.AddNewValue("state", changes.State)
	}

	if _, ok := input["del"]; ok && (item.Del != changes.Del) && (item.Del == nil || changes.Del == nil || *item.Del != *changes.Del) {
		item.Del = changes.Del
		event.AddNewValue("del", changes.Del)
	}

	errText, resErr := utils.Validator(item)
	if resErr != nil {
		return item, &errText
	}

	if err = tx.Create(item).Error; err != nil {
		tx.Rollback()
		return
	}

	if ids, ok := input["adminIds"].([]interface{}); ok {
		items := []Admin{}
		tx.Find(&items, "id IN (?)", ids)

		for k, _ := range items {
			items[k].State = item.State
			items[k].Del = item.Del
		}

		association := tx.Model(&item).Association("Admin")
		association.Replace(items)
	}

	if ids, ok := input["groupIds"].([]interface{}); ok {
		items := []Admin{}
		tx.Find(&items, "id IN (?)", ids)

		for k, _ := range items {
			items[k].State = item.State
			items[k].Del = item.Del
		}

		association := tx.Model(&item).Association("Group")
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
func (r *GeneratedMutationResolver) UpdateRole(ctx context.Context, id string, input map[string]interface{}) (item *Role, err error) {
	return r.Handlers.UpdateRole(ctx, r.GeneratedResolver, id, input)
}
func UpdateRoleHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Role, err error) {
	principalID := getPrincipalIDFromContext(ctx)
	item = &Role{}
	now := time.Now()
	tx := r.DB.db.Begin()

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Role",
		EntityID:    id,
		Date:        now.Unix(),
		PrincipalID: principalID,
	})

	if input["state"] == nil {
		input["state"] = 1
	}

	if input["del"] == nil {
		input["del"] = 1
	}

	var changes RoleChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		return
	}

	if _, ok := input["Name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		event.AddOldValue("Name", item.Name)
		event.AddNewValue("Name", changes.Name)
		item.Name = changes.Name
	}

	if _, ok := input["Pid"]; ok && (item.Pid != changes.Pid) && (item.Pid == nil || changes.Pid == nil || *item.Pid != *changes.Pid) {
		event.AddOldValue("Pid", item.Pid)
		event.AddNewValue("Pid", changes.Pid)
		item.Pid = changes.Pid
	}

	if _, ok := input["state"]; ok && (item.State != changes.State) && (item.State == nil || changes.State == nil || *item.State != *changes.State) {
		event.AddOldValue("state", item.State)
		event.AddNewValue("state", changes.State)
		item.State = changes.State
	}

	if _, ok := input["del"]; ok && (item.Del != changes.Del) && (item.Del == nil || changes.Del == nil || *item.Del != *changes.Del) {
		event.AddOldValue("del", item.Del)
		event.AddNewValue("del", changes.Del)
		item.Del = changes.Del
	}

	errText, resErr := utils.Validator(item)
	if resErr != nil {
		return item, &errText
	}

	item.UpdatedBy = principalID
	item.ID = id

	if err = tx.Model(&item).Updates(item).Error; err != nil {
		tx.Rollback()
		return
	}

	if ids, ok := input["adminIds"].([]interface{}); ok {
		items := []Admin{}
		tx.Find(&items, "id IN (?)", ids)

		for k, _ := range items {
			items[k].State = item.State
			items[k].Del = item.Del
		}

		association := tx.Model(&item).Association("Admin")
		association.Replace(items)
	}

	if ids, ok := input["groupIds"].([]interface{}); ok {
		items := []Admin{}
		tx.Find(&items, "id IN (?)", ids)

		for k, _ := range items {
			items[k].State = item.State
			items[k].Del = item.Del
		}

		association := tx.Model(&item).Association("Group")
		association.Replace(items)
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
func (r *GeneratedMutationResolver) DeleteRole(ctx context.Context, id string) (item *Role, err error) {
	return r.Handlers.DeleteRole(ctx, r.GeneratedResolver, id)
}
func DeleteRoleHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Role, err error) {
	principalID := getPrincipalIDFromContext(ctx)
	item = &Role{}
	now := time.Now()
	tx := r.DB.db.Begin()

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	// 2为删除
	var del int64 = 2

	item.UpdatedBy = principalID
	item.Del = &del

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Role",
		EntityID:    id,
		Date:        now.Unix(),
		PrincipalID: principalID,
	})

	// err = tx.Delete(item, "roles.id = ?", id).Error

	if err = tx.Save(item).Error; err != nil {
		tx.Rollback()
		return
	}

	admin := []Admin{}
	if err = tx.Model(&admin).Where("rolesId = ?", id).Update("del", del).Error; err != nil {
		tx.Rollback()
		return
	}

	group := []Admin{}
	if err = tx.Model(&group).Where("rolesId = ?", id).Update("del", del).Error; err != nil {
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
func (r *GeneratedMutationResolver) DeleteAllRoles(ctx context.Context) (bool, error) {
	return r.Handlers.DeleteAllRoles(ctx, r.GeneratedResolver)
}
func DeleteAllRolesHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	err := r.DB.db.Delete(&Role{}).Error
	return err == nil, err
}
