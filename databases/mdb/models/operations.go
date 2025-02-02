// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Operation is an object representing the database table.
type Operation struct {
	ID         int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	UID        string      `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	TypeID     int64       `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	CreatedAt  time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	Station    null.String `boil:"station" json:"station,omitempty" toml:"station" yaml:"station,omitempty"`
	UserID     null.Int64  `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id,omitempty"`
	Details    null.String `boil:"details" json:"details,omitempty" toml:"details" yaml:"details,omitempty"`
	Properties null.JSON   `boil:"properties" json:"properties,omitempty" toml:"properties" yaml:"properties,omitempty"`

	R *operationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L operationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OperationColumns = struct {
	ID         string
	UID        string
	TypeID     string
	CreatedAt  string
	Station    string
	UserID     string
	Details    string
	Properties string
}{
	ID:         "id",
	UID:        "uid",
	TypeID:     "type_id",
	CreatedAt:  "created_at",
	Station:    "station",
	UserID:     "user_id",
	Details:    "details",
	Properties: "properties",
}

// Generated where

var OperationWhere = struct {
	ID         whereHelperint64
	UID        whereHelperstring
	TypeID     whereHelperint64
	CreatedAt  whereHelpertime_Time
	Station    whereHelpernull_String
	UserID     whereHelpernull_Int64
	Details    whereHelpernull_String
	Properties whereHelpernull_JSON
}{
	ID:         whereHelperint64{field: "\"operations\".\"id\""},
	UID:        whereHelperstring{field: "\"operations\".\"uid\""},
	TypeID:     whereHelperint64{field: "\"operations\".\"type_id\""},
	CreatedAt:  whereHelpertime_Time{field: "\"operations\".\"created_at\""},
	Station:    whereHelpernull_String{field: "\"operations\".\"station\""},
	UserID:     whereHelpernull_Int64{field: "\"operations\".\"user_id\""},
	Details:    whereHelpernull_String{field: "\"operations\".\"details\""},
	Properties: whereHelpernull_JSON{field: "\"operations\".\"properties\""},
}

// OperationRels is where relationship names are stored.
var OperationRels = struct {
	Type  string
	User  string
	Files string
}{
	Type:  "Type",
	User:  "User",
	Files: "Files",
}

// operationR is where relationships are stored.
type operationR struct {
	Type  *OperationType
	User  *User
	Files FileSlice
}

// NewStruct creates a new relationship struct
func (*operationR) NewStruct() *operationR {
	return &operationR{}
}

// operationL is where Load methods for each relationship are stored.
type operationL struct{}

var (
	operationAllColumns            = []string{"id", "uid", "type_id", "created_at", "station", "user_id", "details", "properties"}
	operationColumnsWithoutDefault = []string{"uid", "type_id", "station", "user_id", "details", "properties"}
	operationColumnsWithDefault    = []string{"id", "created_at"}
	operationPrimaryKeyColumns     = []string{"id"}
)

type (
	// OperationSlice is an alias for a slice of pointers to Operation.
	// This should generally be used opposed to []Operation.
	OperationSlice []*Operation

	operationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	operationType                 = reflect.TypeOf(&Operation{})
	operationMapping              = queries.MakeStructMapping(operationType)
	operationPrimaryKeyMapping, _ = queries.BindMapping(operationType, operationMapping, operationPrimaryKeyColumns)
	operationInsertCacheMut       sync.RWMutex
	operationInsertCache          = make(map[string]insertCache)
	operationUpdateCacheMut       sync.RWMutex
	operationUpdateCache          = make(map[string]updateCache)
	operationUpsertCacheMut       sync.RWMutex
	operationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single operation record from the query.
func (q operationQuery) One(exec boil.Executor) (*Operation, error) {
	o := &Operation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for operations")
	}

	return o, nil
}

// All returns all Operation records from the query.
func (q operationQuery) All(exec boil.Executor) (OperationSlice, error) {
	var o []*Operation

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Operation slice")
	}

	return o, nil
}

// Count returns the count of all Operation records in the query.
func (q operationQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count operations rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q operationQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if operations exists")
	}

	return count > 0, nil
}

// Type pointed to by the foreign key.
func (o *Operation) Type(mods ...qm.QueryMod) operationTypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := OperationTypes(queryMods...)
	queries.SetFrom(query.Query, "\"operation_types\"")

	return query
}

// User pointed to by the foreign key.
func (o *Operation) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// Files retrieves all the file's Files with an executor.
func (o *Operation) Files(mods ...qm.QueryMod) fileQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("\"files_operations\" on \"files\".\"id\" = \"files_operations\".\"file_id\""),
		qm.Where("\"files_operations\".\"operation_id\"=?", o.ID),
	)

	query := Files(queryMods...)
	queries.SetFrom(query.Query, "\"files\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"files\".*"})
	}

	return query
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (operationL) LoadType(e boil.Executor, singular bool, maybeOperation interface{}, mods queries.Applicator) error {
	var slice []*Operation
	var object *Operation

	if singular {
		object = maybeOperation.(*Operation)
	} else {
		slice = *maybeOperation.(*[]*Operation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &operationR{}
		}
		args = append(args, object.TypeID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &operationR{}
			}

			for _, a := range args {
				if a == obj.TypeID {
					continue Outer
				}
			}

			args = append(args, obj.TypeID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`operation_types`), qm.WhereIn(`operation_types.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load OperationType")
	}

	var resultSlice []*OperationType
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice OperationType")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for operation_types")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for operation_types")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Type = foreign
		if foreign.R == nil {
			foreign.R = &operationTypeR{}
		}
		foreign.R.TypeOperations = append(foreign.R.TypeOperations, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TypeID == foreign.ID {
				local.R.Type = foreign
				if foreign.R == nil {
					foreign.R = &operationTypeR{}
				}
				foreign.R.TypeOperations = append(foreign.R.TypeOperations, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (operationL) LoadUser(e boil.Executor, singular bool, maybeOperation interface{}, mods queries.Applicator) error {
	var slice []*Operation
	var object *Operation

	if singular {
		object = maybeOperation.(*Operation)
	} else {
		slice = *maybeOperation.(*[]*Operation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &operationR{}
		}
		if !queries.IsNil(object.UserID) {
			args = append(args, object.UserID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &operationR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.UserID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.UserID) {
				args = append(args, obj.UserID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`users`), qm.WhereIn(`users.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.Operations = append(foreign.R.Operations, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.UserID, foreign.ID) {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Operations = append(foreign.R.Operations, local)
				break
			}
		}
	}

	return nil
}

// LoadFiles allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (operationL) LoadFiles(e boil.Executor, singular bool, maybeOperation interface{}, mods queries.Applicator) error {
	var slice []*Operation
	var object *Operation

	if singular {
		object = maybeOperation.(*Operation)
	} else {
		slice = *maybeOperation.(*[]*Operation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &operationR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &operationR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.Select("\"files\".*, \"a\".\"operation_id\""),
		qm.From("\"files\""),
		qm.InnerJoin("\"files_operations\" as \"a\" on \"files\".\"id\" = \"a\".\"file_id\""),
		qm.WhereIn("\"a\".\"operation_id\" in ?", args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load files")
	}

	var resultSlice []*File

	var localJoinCols []int64
	for results.Next() {
		one := new(File)
		var localJoinCol int64

		err = results.Scan(&one.ID, &one.UID, &one.Name, &one.Size, &one.Type, &one.SubType, &one.MimeType, &one.Sha1, &one.ContentUnitID, &one.CreatedAt, &one.Language, &one.BackupCount, &one.FirstBackupTime, &one.Properties, &one.ParentID, &one.FileCreatedAt, &one.Secure, &one.Published, &one.RemovedAt, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for files")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice files")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on files")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for files")
	}

	if singular {
		object.R.Files = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &fileR{}
			}
			foreign.R.Operations = append(foreign.R.Operations, object)
		}
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.ID == localJoinCol {
				local.R.Files = append(local.R.Files, foreign)
				if foreign.R == nil {
					foreign.R = &fileR{}
				}
				foreign.R.Operations = append(foreign.R.Operations, local)
				break
			}
		}
	}

	return nil
}

// SetType of the operation to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeOperations.
func (o *Operation) SetType(exec boil.Executor, insert bool, related *OperationType) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"operations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, operationPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.ID
	if o.R == nil {
		o.R = &operationR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &operationTypeR{
			TypeOperations: OperationSlice{o},
		}
	} else {
		related.R.TypeOperations = append(related.R.TypeOperations, o)
	}

	return nil
}

// SetUser of the operation to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Operations.
func (o *Operation) SetUser(exec boil.Executor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"operations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, operationPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.UserID, related.ID)
	if o.R == nil {
		o.R = &operationR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			Operations: OperationSlice{o},
		}
	} else {
		related.R.Operations = append(related.R.Operations, o)
	}

	return nil
}

// RemoveUser relationship.
// Sets o.R.User to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Operation) RemoveUser(exec boil.Executor, related *User) error {
	var err error

	queries.SetScanner(&o.UserID, nil)
	if _, err = o.Update(exec, boil.Whitelist("user_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.User = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Operations {
		if queries.Equal(o.UserID, ri.UserID) {
			continue
		}

		ln := len(related.R.Operations)
		if ln > 1 && i < ln-1 {
			related.R.Operations[i] = related.R.Operations[ln-1]
		}
		related.R.Operations = related.R.Operations[:ln-1]
		break
	}
	return nil
}

// AddFiles adds the given related objects to the existing relationships
// of the operation, optionally inserting them as new records.
// Appends related to o.R.Files.
// Sets related.R.Operations appropriately.
func (o *Operation) AddFiles(exec boil.Executor, insert bool, related ...*File) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into \"files_operations\" (\"operation_id\", \"file_id\") values ($1, $2)"
		values := []interface{}{o.ID, rel.ID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, query)
			fmt.Fprintln(boil.DebugWriter, values)
		}
		_, err = exec.Exec(query, values...)
		if err != nil {
			return errors.Wrap(err, "failed to insert into join table")
		}
	}
	if o.R == nil {
		o.R = &operationR{
			Files: related,
		}
	} else {
		o.R.Files = append(o.R.Files, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &fileR{
				Operations: OperationSlice{o},
			}
		} else {
			rel.R.Operations = append(rel.R.Operations, o)
		}
	}
	return nil
}

// SetFiles removes all previously related items of the
// operation replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Operations's Files accordingly.
// Replaces o.R.Files with related.
// Sets related.R.Operations's Files accordingly.
func (o *Operation) SetFiles(exec boil.Executor, insert bool, related ...*File) error {
	query := "delete from \"files_operations\" where \"operation_id\" = $1"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	removeFilesFromOperationsSlice(o, related)
	if o.R != nil {
		o.R.Files = nil
	}
	return o.AddFiles(exec, insert, related...)
}

// RemoveFiles relationships from objects passed in.
// Removes related items from R.Files (uses pointer comparison, removal does not keep order)
// Sets related.R.Operations.
func (o *Operation) RemoveFiles(exec boil.Executor, related ...*File) error {
	var err error
	query := fmt.Sprintf(
		"delete from \"files_operations\" where \"operation_id\" = $1 and \"file_id\" in (%s)",
		strmangle.Placeholders(dialect.UseIndexPlaceholders, len(related), 2, 1),
	)
	values := []interface{}{o.ID}
	for _, rel := range related {
		values = append(values, rel.ID)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	_, err = exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}
	removeFilesFromOperationsSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Files {
			if rel != ri {
				continue
			}

			ln := len(o.R.Files)
			if ln > 1 && i < ln-1 {
				o.R.Files[i] = o.R.Files[ln-1]
			}
			o.R.Files = o.R.Files[:ln-1]
			break
		}
	}

	return nil
}

func removeFilesFromOperationsSlice(o *Operation, related []*File) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.Operations {
			if o.ID != ri.ID {
				continue
			}

			ln := len(rel.R.Operations)
			if ln > 1 && i < ln-1 {
				rel.R.Operations[i] = rel.R.Operations[ln-1]
			}
			rel.R.Operations = rel.R.Operations[:ln-1]
			break
		}
	}
}

// Operations retrieves all the records using an executor.
func Operations(mods ...qm.QueryMod) operationQuery {
	mods = append(mods, qm.From("\"operations\""))
	return operationQuery{NewQuery(mods...)}
}

// FindOperation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOperation(exec boil.Executor, iD int64, selectCols ...string) (*Operation, error) {
	operationObj := &Operation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"operations\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, operationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from operations")
	}

	return operationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Operation) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no operations provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(operationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	operationInsertCacheMut.RLock()
	cache, cached := operationInsertCache[key]
	operationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			operationAllColumns,
			operationColumnsWithDefault,
			operationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(operationType, operationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(operationType, operationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"operations\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"operations\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into operations")
	}

	if !cached {
		operationInsertCacheMut.Lock()
		operationInsertCache[key] = cache
		operationInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Operation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Operation) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	operationUpdateCacheMut.RLock()
	cache, cached := operationUpdateCache[key]
	operationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			operationAllColumns,
			operationPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update operations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"operations\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, operationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(operationType, operationMapping, append(wl, operationPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update operations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for operations")
	}

	if !cached {
		operationUpdateCacheMut.Lock()
		operationUpdateCache[key] = cache
		operationUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q operationQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for operations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for operations")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OperationSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), operationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"operations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, operationPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in operation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all operation")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Operation) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no operations provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(operationColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	operationUpsertCacheMut.RLock()
	cache, cached := operationUpsertCache[key]
	operationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			operationAllColumns,
			operationColumnsWithDefault,
			operationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			operationAllColumns,
			operationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert operations, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(operationPrimaryKeyColumns))
			copy(conflict, operationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"operations\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(operationType, operationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(operationType, operationMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert operations")
	}

	if !cached {
		operationUpsertCacheMut.Lock()
		operationUpsertCache[key] = cache
		operationUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Operation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Operation) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Operation provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), operationPrimaryKeyMapping)
	sql := "DELETE FROM \"operations\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from operations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for operations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q operationQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no operationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from operations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for operations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OperationSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), operationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"operations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, operationPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from operation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for operations")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Operation) Reload(exec boil.Executor) error {
	ret, err := FindOperation(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OperationSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OperationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), operationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"operations\".* FROM \"operations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, operationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OperationSlice")
	}

	*o = slice

	return nil
}

// OperationExists checks if the Operation row exists.
func OperationExists(exec boil.Executor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"operations\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if operations exists")
	}

	return exists, nil
}
