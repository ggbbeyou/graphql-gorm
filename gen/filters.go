package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

func (f *UserFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, "users", wheres, values, joins)
}
func (f *UserFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _values := f.WhereContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			err := or.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.Tasks != nil {
		_alias := alias + "_tasks"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("tasks")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("assigneeId")+" = "+dialect.Quote(alias)+".id")
		err := f.Tasks.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *UserFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}
	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}
	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}
	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}
	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}
	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}
	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.Email != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" = ?")
		values = append(values, f.Email)
	}
	if f.EmailNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" != ?")
		values = append(values, f.EmailNe)
	}
	if f.EmailGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" > ?")
		values = append(values, f.EmailGt)
	}
	if f.EmailLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" < ?")
		values = append(values, f.EmailLt)
	}
	if f.EmailGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" >= ?")
		values = append(values, f.EmailGte)
	}
	if f.EmailLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" <= ?")
		values = append(values, f.EmailLte)
	}
	if f.EmailIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" IN (?)")
		values = append(values, f.EmailIn)
	}
	if f.EmailLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EmailLike, "?", "_", -1), "*", "%", -1))
	}
	if f.EmailPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EmailPrefix))
	}
	if f.EmailSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EmailSuffix))
	}

	if f.FirstName != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" = ?")
		values = append(values, f.FirstName)
	}
	if f.FirstNameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" != ?")
		values = append(values, f.FirstNameNe)
	}
	if f.FirstNameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" > ?")
		values = append(values, f.FirstNameGt)
	}
	if f.FirstNameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" < ?")
		values = append(values, f.FirstNameLt)
	}
	if f.FirstNameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" >= ?")
		values = append(values, f.FirstNameGte)
	}
	if f.FirstNameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" <= ?")
		values = append(values, f.FirstNameLte)
	}
	if f.FirstNameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" IN (?)")
		values = append(values, f.FirstNameIn)
	}
	if f.FirstNameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.FirstNameLike, "?", "_", -1), "*", "%", -1))
	}
	if f.FirstNamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.FirstNamePrefix))
	}
	if f.FirstNameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.FirstNameSuffix))
	}

	if f.LastName != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" = ?")
		values = append(values, f.LastName)
	}
	if f.LastNameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" != ?")
		values = append(values, f.LastNameNe)
	}
	if f.LastNameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" > ?")
		values = append(values, f.LastNameGt)
	}
	if f.LastNameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" < ?")
		values = append(values, f.LastNameLt)
	}
	if f.LastNameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" >= ?")
		values = append(values, f.LastNameGte)
	}
	if f.LastNameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" <= ?")
		values = append(values, f.LastNameLte)
	}
	if f.LastNameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" IN (?)")
		values = append(values, f.LastNameIn)
	}
	if f.LastNameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LastNameLike, "?", "_", -1), "*", "%", -1))
	}
	if f.LastNamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LastNamePrefix))
	}
	if f.LastNameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LastNameSuffix))
	}

	if f.State != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" = ?")
		values = append(values, f.State)
	}
	if f.StateNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" != ?")
		values = append(values, f.StateNe)
	}
	if f.StateGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" > ?")
		values = append(values, f.StateGt)
	}
	if f.StateLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" < ?")
		values = append(values, f.StateLt)
	}
	if f.StateGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" >= ?")
		values = append(values, f.StateGte)
	}
	if f.StateLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" <= ?")
		values = append(values, f.StateLte)
	}
	if f.StateIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" IN (?)")
		values = append(values, f.StateIn)
	}

	if f.Del != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" = ?")
		values = append(values, f.Del)
	}
	if f.DelNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" != ?")
		values = append(values, f.DelNe)
	}
	if f.DelGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" > ?")
		values = append(values, f.DelGt)
	}
	if f.DelLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" < ?")
		values = append(values, f.DelLt)
	}
	if f.DelGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" >= ?")
		values = append(values, f.DelGte)
	}
	if f.DelLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" <= ?")
		values = append(values, f.DelLte)
	}
	if f.DelIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" IN (?)")
		values = append(values, f.DelIn)
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}
	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}
	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}
	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}
	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}
	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}
	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}
	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}
	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}
	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}
	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}
	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}
	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.DeletedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" = ?")
		values = append(values, f.DeletedBy)
	}
	if f.DeletedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" != ?")
		values = append(values, f.DeletedByNe)
	}
	if f.DeletedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" > ?")
		values = append(values, f.DeletedByGt)
	}
	if f.DeletedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" < ?")
		values = append(values, f.DeletedByLt)
	}
	if f.DeletedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" >= ?")
		values = append(values, f.DeletedByGte)
	}
	if f.DeletedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" <= ?")
		values = append(values, f.DeletedByLte)
	}
	if f.DeletedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" IN (?)")
		values = append(values, f.DeletedByIn)
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}
	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}
	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}
	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}
	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}
	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}
	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}
	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}
	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}
	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}
	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}
	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}
	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *UserFilterType) AndWith(f2 ...*UserFilterType) *UserFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &UserFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *UserFilterType) OrWith(f2 ...*UserFilterType) *UserFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &UserFilterType{
		Or: append(_f2, f),
	}
}

func (f *TaskFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, "tasks", wheres, values, joins)
}
func (f *TaskFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _values := f.WhereContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			err := or.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.Assignee != nil {
		_alias := alias + "_assignee"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("users")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("assigneeId"))
		err := f.Assignee.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *TaskFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}
	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}
	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}
	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}
	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}
	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}
	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.Title != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" = ?")
		values = append(values, f.Title)
	}
	if f.TitleNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" != ?")
		values = append(values, f.TitleNe)
	}
	if f.TitleGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" > ?")
		values = append(values, f.TitleGt)
	}
	if f.TitleLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" < ?")
		values = append(values, f.TitleLt)
	}
	if f.TitleGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" >= ?")
		values = append(values, f.TitleGte)
	}
	if f.TitleLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" <= ?")
		values = append(values, f.TitleLte)
	}
	if f.TitleIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" IN (?)")
		values = append(values, f.TitleIn)
	}
	if f.TitleLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.TitleLike, "?", "_", -1), "*", "%", -1))
	}
	if f.TitlePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TitlePrefix))
	}
	if f.TitleSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TitleSuffix))
	}

	if f.Completed != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" = ?")
		values = append(values, f.Completed)
	}
	if f.CompletedNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" != ?")
		values = append(values, f.CompletedNe)
	}
	if f.CompletedGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" > ?")
		values = append(values, f.CompletedGt)
	}
	if f.CompletedLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" < ?")
		values = append(values, f.CompletedLt)
	}
	if f.CompletedGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" >= ?")
		values = append(values, f.CompletedGte)
	}
	if f.CompletedLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" <= ?")
		values = append(values, f.CompletedLte)
	}
	if f.CompletedIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" IN (?)")
		values = append(values, f.CompletedIn)
	}

	if f.DueDate != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" = ?")
		values = append(values, f.DueDate)
	}
	if f.DueDateNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" != ?")
		values = append(values, f.DueDateNe)
	}
	if f.DueDateGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" > ?")
		values = append(values, f.DueDateGt)
	}
	if f.DueDateLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" < ?")
		values = append(values, f.DueDateLt)
	}
	if f.DueDateGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" >= ?")
		values = append(values, f.DueDateGte)
	}
	if f.DueDateLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" <= ?")
		values = append(values, f.DueDateLte)
	}
	if f.DueDateIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" IN (?)")
		values = append(values, f.DueDateIn)
	}

	if f.AssigneeID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" = ?")
		values = append(values, f.AssigneeID)
	}
	if f.AssigneeIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" != ?")
		values = append(values, f.AssigneeIDNe)
	}
	if f.AssigneeIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" > ?")
		values = append(values, f.AssigneeIDGt)
	}
	if f.AssigneeIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" < ?")
		values = append(values, f.AssigneeIDLt)
	}
	if f.AssigneeIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" >= ?")
		values = append(values, f.AssigneeIDGte)
	}
	if f.AssigneeIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" <= ?")
		values = append(values, f.AssigneeIDLte)
	}
	if f.AssigneeIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" IN (?)")
		values = append(values, f.AssigneeIDIn)
	}

	if f.State != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" = ?")
		values = append(values, f.State)
	}
	if f.StateNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" != ?")
		values = append(values, f.StateNe)
	}
	if f.StateGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" > ?")
		values = append(values, f.StateGt)
	}
	if f.StateLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" < ?")
		values = append(values, f.StateLt)
	}
	if f.StateGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" >= ?")
		values = append(values, f.StateGte)
	}
	if f.StateLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" <= ?")
		values = append(values, f.StateLte)
	}
	if f.StateIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" IN (?)")
		values = append(values, f.StateIn)
	}

	if f.Del != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" = ?")
		values = append(values, f.Del)
	}
	if f.DelNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" != ?")
		values = append(values, f.DelNe)
	}
	if f.DelGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" > ?")
		values = append(values, f.DelGt)
	}
	if f.DelLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" < ?")
		values = append(values, f.DelLt)
	}
	if f.DelGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" >= ?")
		values = append(values, f.DelGte)
	}
	if f.DelLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" <= ?")
		values = append(values, f.DelLte)
	}
	if f.DelIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" IN (?)")
		values = append(values, f.DelIn)
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}
	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}
	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}
	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}
	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}
	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}
	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}
	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}
	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}
	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}
	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}
	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}
	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.DeletedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" = ?")
		values = append(values, f.DeletedBy)
	}
	if f.DeletedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" != ?")
		values = append(values, f.DeletedByNe)
	}
	if f.DeletedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" > ?")
		values = append(values, f.DeletedByGt)
	}
	if f.DeletedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" < ?")
		values = append(values, f.DeletedByLt)
	}
	if f.DeletedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" >= ?")
		values = append(values, f.DeletedByGte)
	}
	if f.DeletedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" <= ?")
		values = append(values, f.DeletedByLte)
	}
	if f.DeletedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" IN (?)")
		values = append(values, f.DeletedByIn)
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}
	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}
	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}
	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}
	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}
	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}
	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}
	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}
	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}
	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}
	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}
	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}
	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *TaskFilterType) AndWith(f2 ...*TaskFilterType) *TaskFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &TaskFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *TaskFilterType) OrWith(f2 ...*TaskFilterType) *TaskFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &TaskFilterType{
		Or: append(_f2, f),
	}
}

func (f *AdminFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, "admins", wheres, values, joins)
}
func (f *AdminFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _values := f.WhereContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			err := or.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.Groups != nil {
		_alias := alias + "_groups"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("admin_groups")+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("admin_id")+" LEFT JOIN "+dialect.Quote("groups")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("group_id")+" = "+dialect.Quote(_alias)+".id")
		err := f.Groups.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Roles != nil {
		_alias := alias + "_roles"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("admin_roles")+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("admin_id")+" LEFT JOIN "+dialect.Quote("roles")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("role_id")+" = "+dialect.Quote(_alias)+".id")
		err := f.Roles.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *AdminFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}
	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}
	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}
	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}
	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}
	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}
	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.Phone != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Phone")+" = ?")
		values = append(values, f.Phone)
	}
	if f.PhoneNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Phone")+" != ?")
		values = append(values, f.PhoneNe)
	}
	if f.PhoneGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Phone")+" > ?")
		values = append(values, f.PhoneGt)
	}
	if f.PhoneLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Phone")+" < ?")
		values = append(values, f.PhoneLt)
	}
	if f.PhoneGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Phone")+" >= ?")
		values = append(values, f.PhoneGte)
	}
	if f.PhoneLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Phone")+" <= ?")
		values = append(values, f.PhoneLte)
	}
	if f.PhoneIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Phone")+" IN (?)")
		values = append(values, f.PhoneIn)
	}
	if f.PhoneLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Phone")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.PhoneLike, "?", "_", -1), "*", "%", -1))
	}
	if f.PhonePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Phone")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PhonePrefix))
	}
	if f.PhoneSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Phone")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PhoneSuffix))
	}

	if f.Password != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Password")+" = ?")
		values = append(values, f.Password)
	}
	if f.PasswordNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Password")+" != ?")
		values = append(values, f.PasswordNe)
	}
	if f.PasswordGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Password")+" > ?")
		values = append(values, f.PasswordGt)
	}
	if f.PasswordLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Password")+" < ?")
		values = append(values, f.PasswordLt)
	}
	if f.PasswordGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Password")+" >= ?")
		values = append(values, f.PasswordGte)
	}
	if f.PasswordLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Password")+" <= ?")
		values = append(values, f.PasswordLte)
	}
	if f.PasswordIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Password")+" IN (?)")
		values = append(values, f.PasswordIn)
	}
	if f.PasswordLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Password")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.PasswordLike, "?", "_", -1), "*", "%", -1))
	}
	if f.PasswordPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Password")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PasswordPrefix))
	}
	if f.PasswordSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Password")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PasswordSuffix))
	}

	if f.Username != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Username")+" = ?")
		values = append(values, f.Username)
	}
	if f.UsernameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Username")+" != ?")
		values = append(values, f.UsernameNe)
	}
	if f.UsernameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Username")+" > ?")
		values = append(values, f.UsernameGt)
	}
	if f.UsernameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Username")+" < ?")
		values = append(values, f.UsernameLt)
	}
	if f.UsernameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Username")+" >= ?")
		values = append(values, f.UsernameGte)
	}
	if f.UsernameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Username")+" <= ?")
		values = append(values, f.UsernameLte)
	}
	if f.UsernameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Username")+" IN (?)")
		values = append(values, f.UsernameIn)
	}
	if f.UsernameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Username")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.UsernameLike, "?", "_", -1), "*", "%", -1))
	}
	if f.UsernamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Username")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.UsernamePrefix))
	}
	if f.UsernameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Username")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.UsernameSuffix))
	}

	if f.Money != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Money")+" = ?")
		values = append(values, f.Money)
	}
	if f.MoneyNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Money")+" != ?")
		values = append(values, f.MoneyNe)
	}
	if f.MoneyGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Money")+" > ?")
		values = append(values, f.MoneyGt)
	}
	if f.MoneyLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Money")+" < ?")
		values = append(values, f.MoneyLt)
	}
	if f.MoneyGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Money")+" >= ?")
		values = append(values, f.MoneyGte)
	}
	if f.MoneyLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Money")+" <= ?")
		values = append(values, f.MoneyLte)
	}
	if f.MoneyIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Money")+" IN (?)")
		values = append(values, f.MoneyIn)
	}

	if f.Sex != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Sex")+" = ?")
		values = append(values, f.Sex)
	}
	if f.SexNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Sex")+" != ?")
		values = append(values, f.SexNe)
	}
	if f.SexGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Sex")+" > ?")
		values = append(values, f.SexGt)
	}
	if f.SexLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Sex")+" < ?")
		values = append(values, f.SexLt)
	}
	if f.SexGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Sex")+" >= ?")
		values = append(values, f.SexGte)
	}
	if f.SexLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Sex")+" <= ?")
		values = append(values, f.SexLte)
	}
	if f.SexIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Sex")+" IN (?)")
		values = append(values, f.SexIn)
	}

	if f.Super != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Super")+" = ?")
		values = append(values, f.Super)
	}
	if f.SuperNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Super")+" != ?")
		values = append(values, f.SuperNe)
	}
	if f.SuperGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Super")+" > ?")
		values = append(values, f.SuperGt)
	}
	if f.SuperLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Super")+" < ?")
		values = append(values, f.SuperLt)
	}
	if f.SuperGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Super")+" >= ?")
		values = append(values, f.SuperGte)
	}
	if f.SuperLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Super")+" <= ?")
		values = append(values, f.SuperLte)
	}
	if f.SuperIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Super")+" IN (?)")
		values = append(values, f.SuperIn)
	}

	if f.LoginCount != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LoginCount")+" = ?")
		values = append(values, f.LoginCount)
	}
	if f.LoginCountNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LoginCount")+" != ?")
		values = append(values, f.LoginCountNe)
	}
	if f.LoginCountGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LoginCount")+" > ?")
		values = append(values, f.LoginCountGt)
	}
	if f.LoginCountLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LoginCount")+" < ?")
		values = append(values, f.LoginCountLt)
	}
	if f.LoginCountGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LoginCount")+" >= ?")
		values = append(values, f.LoginCountGte)
	}
	if f.LoginCountLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LoginCount")+" <= ?")
		values = append(values, f.LoginCountLte)
	}
	if f.LoginCountIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LoginCount")+" IN (?)")
		values = append(values, f.LoginCountIn)
	}

	if f.LoginIP != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LoginIp")+" = ?")
		values = append(values, f.LoginIP)
	}
	if f.LoginIPNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LoginIp")+" != ?")
		values = append(values, f.LoginIPNe)
	}
	if f.LoginIPGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LoginIp")+" > ?")
		values = append(values, f.LoginIPGt)
	}
	if f.LoginIPLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LoginIp")+" < ?")
		values = append(values, f.LoginIPLt)
	}
	if f.LoginIPGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LoginIp")+" >= ?")
		values = append(values, f.LoginIPGte)
	}
	if f.LoginIPLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LoginIp")+" <= ?")
		values = append(values, f.LoginIPLte)
	}
	if f.LoginIPIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LoginIp")+" IN (?)")
		values = append(values, f.LoginIPIn)
	}
	if f.LoginIPLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LoginIp")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LoginIPLike, "?", "_", -1), "*", "%", -1))
	}
	if f.LoginIPPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LoginIp")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LoginIPPrefix))
	}
	if f.LoginIPSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LoginIp")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LoginIPSuffix))
	}

	if f.LastIP != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LastIp")+" = ?")
		values = append(values, f.LastIP)
	}
	if f.LastIPNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LastIp")+" != ?")
		values = append(values, f.LastIPNe)
	}
	if f.LastIPGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LastIp")+" > ?")
		values = append(values, f.LastIPGt)
	}
	if f.LastIPLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LastIp")+" < ?")
		values = append(values, f.LastIPLt)
	}
	if f.LastIPGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LastIp")+" >= ?")
		values = append(values, f.LastIPGte)
	}
	if f.LastIPLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LastIp")+" <= ?")
		values = append(values, f.LastIPLte)
	}
	if f.LastIPIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LastIp")+" IN (?)")
		values = append(values, f.LastIPIn)
	}
	if f.LastIPLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LastIp")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LastIPLike, "?", "_", -1), "*", "%", -1))
	}
	if f.LastIPPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LastIp")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LastIPPrefix))
	}
	if f.LastIPSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("LastIp")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LastIPSuffix))
	}

	if f.State != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" = ?")
		values = append(values, f.State)
	}
	if f.StateNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" != ?")
		values = append(values, f.StateNe)
	}
	if f.StateGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" > ?")
		values = append(values, f.StateGt)
	}
	if f.StateLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" < ?")
		values = append(values, f.StateLt)
	}
	if f.StateGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" >= ?")
		values = append(values, f.StateGte)
	}
	if f.StateLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" <= ?")
		values = append(values, f.StateLte)
	}
	if f.StateIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" IN (?)")
		values = append(values, f.StateIn)
	}

	if f.Del != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" = ?")
		values = append(values, f.Del)
	}
	if f.DelNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" != ?")
		values = append(values, f.DelNe)
	}
	if f.DelGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" > ?")
		values = append(values, f.DelGt)
	}
	if f.DelLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" < ?")
		values = append(values, f.DelLt)
	}
	if f.DelGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" >= ?")
		values = append(values, f.DelGte)
	}
	if f.DelLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" <= ?")
		values = append(values, f.DelLte)
	}
	if f.DelIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" IN (?)")
		values = append(values, f.DelIn)
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}
	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}
	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}
	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}
	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}
	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}
	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}
	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}
	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}
	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}
	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}
	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}
	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.DeletedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" = ?")
		values = append(values, f.DeletedBy)
	}
	if f.DeletedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" != ?")
		values = append(values, f.DeletedByNe)
	}
	if f.DeletedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" > ?")
		values = append(values, f.DeletedByGt)
	}
	if f.DeletedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" < ?")
		values = append(values, f.DeletedByLt)
	}
	if f.DeletedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" >= ?")
		values = append(values, f.DeletedByGte)
	}
	if f.DeletedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" <= ?")
		values = append(values, f.DeletedByLte)
	}
	if f.DeletedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" IN (?)")
		values = append(values, f.DeletedByIn)
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}
	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}
	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}
	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}
	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}
	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}
	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}
	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}
	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}
	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}
	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}
	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}
	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *AdminFilterType) AndWith(f2 ...*AdminFilterType) *AdminFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &AdminFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *AdminFilterType) OrWith(f2 ...*AdminFilterType) *AdminFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &AdminFilterType{
		Or: append(_f2, f),
	}
}

func (f *GroupFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, "groups", wheres, values, joins)
}
func (f *GroupFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _values := f.WhereContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			err := or.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.Admin != nil {
		_alias := alias + "_admin"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("admin_groups")+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("group_id")+" LEFT JOIN "+dialect.Quote("admins")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("admin_id")+" = "+dialect.Quote(_alias)+".id")
		err := f.Admin.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Roles != nil {
		_alias := alias + "_roles"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("group_roles")+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("group_id")+" LEFT JOIN "+dialect.Quote("roles")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("role_id")+" = "+dialect.Quote(_alias)+".id")
		err := f.Roles.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *GroupFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}
	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}
	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}
	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}
	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}
	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}
	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Name")+" = ?")
		values = append(values, f.Name)
	}
	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Name")+" != ?")
		values = append(values, f.NameNe)
	}
	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Name")+" > ?")
		values = append(values, f.NameGt)
	}
	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Name")+" < ?")
		values = append(values, f.NameLt)
	}
	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Name")+" >= ?")
		values = append(values, f.NameGte)
	}
	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Name")+" <= ?")
		values = append(values, f.NameLte)
	}
	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Name")+" IN (?)")
		values = append(values, f.NameIn)
	}
	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Name")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1))
	}
	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}
	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.State != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" = ?")
		values = append(values, f.State)
	}
	if f.StateNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" != ?")
		values = append(values, f.StateNe)
	}
	if f.StateGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" > ?")
		values = append(values, f.StateGt)
	}
	if f.StateLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" < ?")
		values = append(values, f.StateLt)
	}
	if f.StateGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" >= ?")
		values = append(values, f.StateGte)
	}
	if f.StateLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" <= ?")
		values = append(values, f.StateLte)
	}
	if f.StateIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" IN (?)")
		values = append(values, f.StateIn)
	}

	if f.Del != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" = ?")
		values = append(values, f.Del)
	}
	if f.DelNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" != ?")
		values = append(values, f.DelNe)
	}
	if f.DelGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" > ?")
		values = append(values, f.DelGt)
	}
	if f.DelLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" < ?")
		values = append(values, f.DelLt)
	}
	if f.DelGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" >= ?")
		values = append(values, f.DelGte)
	}
	if f.DelLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" <= ?")
		values = append(values, f.DelLte)
	}
	if f.DelIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" IN (?)")
		values = append(values, f.DelIn)
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}
	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}
	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}
	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}
	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}
	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}
	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}
	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}
	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}
	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}
	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}
	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}
	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.DeletedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" = ?")
		values = append(values, f.DeletedBy)
	}
	if f.DeletedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" != ?")
		values = append(values, f.DeletedByNe)
	}
	if f.DeletedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" > ?")
		values = append(values, f.DeletedByGt)
	}
	if f.DeletedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" < ?")
		values = append(values, f.DeletedByLt)
	}
	if f.DeletedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" >= ?")
		values = append(values, f.DeletedByGte)
	}
	if f.DeletedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" <= ?")
		values = append(values, f.DeletedByLte)
	}
	if f.DeletedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" IN (?)")
		values = append(values, f.DeletedByIn)
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}
	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}
	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}
	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}
	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}
	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}
	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}
	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}
	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}
	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}
	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}
	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}
	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *GroupFilterType) AndWith(f2 ...*GroupFilterType) *GroupFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &GroupFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *GroupFilterType) OrWith(f2 ...*GroupFilterType) *GroupFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &GroupFilterType{
		Or: append(_f2, f),
	}
}

func (f *RoleFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, "roles", wheres, values, joins)
}
func (f *RoleFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _values := f.WhereContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			err := or.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.Admin != nil {
		_alias := alias + "_admin"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("admin_roles")+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("role_id")+" LEFT JOIN "+dialect.Quote("admins")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("admin_id")+" = "+dialect.Quote(_alias)+".id")
		err := f.Admin.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Group != nil {
		_alias := alias + "_group"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("admin_roles")+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("role_id")+" LEFT JOIN "+dialect.Quote("admins")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("group_id")+" = "+dialect.Quote(_alias)+".id")
		err := f.Group.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *RoleFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}
	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}
	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}
	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}
	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}
	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}
	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Name")+" = ?")
		values = append(values, f.Name)
	}
	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Name")+" != ?")
		values = append(values, f.NameNe)
	}
	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Name")+" > ?")
		values = append(values, f.NameGt)
	}
	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Name")+" < ?")
		values = append(values, f.NameLt)
	}
	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Name")+" >= ?")
		values = append(values, f.NameGte)
	}
	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Name")+" <= ?")
		values = append(values, f.NameLte)
	}
	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Name")+" IN (?)")
		values = append(values, f.NameIn)
	}
	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Name")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1))
	}
	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}
	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.Pid != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Pid")+" = ?")
		values = append(values, f.Pid)
	}
	if f.PidNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Pid")+" != ?")
		values = append(values, f.PidNe)
	}
	if f.PidGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Pid")+" > ?")
		values = append(values, f.PidGt)
	}
	if f.PidLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Pid")+" < ?")
		values = append(values, f.PidLt)
	}
	if f.PidGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Pid")+" >= ?")
		values = append(values, f.PidGte)
	}
	if f.PidLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Pid")+" <= ?")
		values = append(values, f.PidLte)
	}
	if f.PidIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Pid")+" IN (?)")
		values = append(values, f.PidIn)
	}
	if f.PidLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Pid")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.PidLike, "?", "_", -1), "*", "%", -1))
	}
	if f.PidPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Pid")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PidPrefix))
	}
	if f.PidSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("Pid")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PidSuffix))
	}

	if f.State != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" = ?")
		values = append(values, f.State)
	}
	if f.StateNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" != ?")
		values = append(values, f.StateNe)
	}
	if f.StateGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" > ?")
		values = append(values, f.StateGt)
	}
	if f.StateLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" < ?")
		values = append(values, f.StateLt)
	}
	if f.StateGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" >= ?")
		values = append(values, f.StateGte)
	}
	if f.StateLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" <= ?")
		values = append(values, f.StateLte)
	}
	if f.StateIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" IN (?)")
		values = append(values, f.StateIn)
	}

	if f.Del != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" = ?")
		values = append(values, f.Del)
	}
	if f.DelNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" != ?")
		values = append(values, f.DelNe)
	}
	if f.DelGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" > ?")
		values = append(values, f.DelGt)
	}
	if f.DelLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" < ?")
		values = append(values, f.DelLt)
	}
	if f.DelGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" >= ?")
		values = append(values, f.DelGte)
	}
	if f.DelLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" <= ?")
		values = append(values, f.DelLte)
	}
	if f.DelIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("del")+" IN (?)")
		values = append(values, f.DelIn)
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}
	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}
	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}
	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}
	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}
	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}
	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}
	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}
	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}
	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}
	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}
	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}
	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.DeletedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" = ?")
		values = append(values, f.DeletedBy)
	}
	if f.DeletedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" != ?")
		values = append(values, f.DeletedByNe)
	}
	if f.DeletedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" > ?")
		values = append(values, f.DeletedByGt)
	}
	if f.DeletedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" < ?")
		values = append(values, f.DeletedByLt)
	}
	if f.DeletedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" >= ?")
		values = append(values, f.DeletedByGte)
	}
	if f.DeletedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" <= ?")
		values = append(values, f.DeletedByLte)
	}
	if f.DeletedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("deletedBy")+" IN (?)")
		values = append(values, f.DeletedByIn)
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}
	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}
	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}
	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}
	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}
	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}
	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}
	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}
	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}
	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}
	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}
	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}
	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *RoleFilterType) AndWith(f2 ...*RoleFilterType) *RoleFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &RoleFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *RoleFilterType) OrWith(f2 ...*RoleFilterType) *RoleFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &RoleFilterType{
		Or: append(_f2, f),
	}
}
