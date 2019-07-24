package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jinzhu/gorm"
	"github.com/vektah/gqlparser/ast"
)

type UserQueryFilter struct {
	Query *string
}

func (qf *UserQueryFilter) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}
	fields := []*ast.Field{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		fields = append(fields, f.Field)
	}

	ors := []string{}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		if err := qf.applyQueryWithFields(dialect, fields, part, "users", &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *UserQueryFilter) applyQueryWithFields(dialect gorm.Dialect, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = f
	}

	if _, ok := fieldsMap["email"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]s"+dialect.Quote("email")+" LIKE ? OR %[1]s"+dialect.Quote("email")+" LIKE ?", dialect.Quote(alias)+"."))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	if _, ok := fieldsMap["firstName"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]s"+dialect.Quote("firstName")+" LIKE ? OR %[1]s"+dialect.Quote("firstName")+" LIKE ?", dialect.Quote(alias)+"."))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	if _, ok := fieldsMap["lastName"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]s"+dialect.Quote("lastName")+" LIKE ? OR %[1]s"+dialect.Quote("lastName")+" LIKE ?", dialect.Quote(alias)+"."))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	if f, ok := fieldsMap["tasks"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_tasks"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("tasks")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("assigneeId")+" = "+dialect.Quote(alias)+".id")

		for _, s := range f.SelectionSet {
			if f, ok := s.(*ast.Field); ok {
				_fields = append(_fields, f)
			}
		}
		q := TaskQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

type TaskQueryFilter struct {
	Query *string
}

func (qf *TaskQueryFilter) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}
	fields := []*ast.Field{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		fields = append(fields, f.Field)
	}

	ors := []string{}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		if err := qf.applyQueryWithFields(dialect, fields, part, "tasks", &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *TaskQueryFilter) applyQueryWithFields(dialect gorm.Dialect, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = f
	}

	if _, ok := fieldsMap["title"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]s"+dialect.Quote("title")+" LIKE ? OR %[1]s"+dialect.Quote("title")+" LIKE ?", dialect.Quote(alias)+"."))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	if f, ok := fieldsMap["assignee"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_assignee"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("users")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("assigneeId"))

		for _, s := range f.SelectionSet {
			if f, ok := s.(*ast.Field); ok {
				_fields = append(_fields, f)
			}
		}
		q := UserQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}
