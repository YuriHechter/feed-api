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

// PublisherI18n is an object representing the database table.
type PublisherI18n struct {
	PublisherID      int64       `boil:"publisher_id" json:"publisher_id" toml:"publisher_id" yaml:"publisher_id"`
	Language         string      `boil:"language" json:"language" toml:"language" yaml:"language"`
	OriginalLanguage null.String `boil:"original_language" json:"original_language,omitempty" toml:"original_language" yaml:"original_language,omitempty"`
	Name             null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Description      null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	UserID           null.Int64  `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id,omitempty"`
	CreatedAt        time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *publisherI18nR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L publisherI18nL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PublisherI18nColumns = struct {
	PublisherID      string
	Language         string
	OriginalLanguage string
	Name             string
	Description      string
	UserID           string
	CreatedAt        string
}{
	PublisherID:      "publisher_id",
	Language:         "language",
	OriginalLanguage: "original_language",
	Name:             "name",
	Description:      "description",
	UserID:           "user_id",
	CreatedAt:        "created_at",
}

// Generated where

var PublisherI18nWhere = struct {
	PublisherID      whereHelperint64
	Language         whereHelperstring
	OriginalLanguage whereHelpernull_String
	Name             whereHelpernull_String
	Description      whereHelpernull_String
	UserID           whereHelpernull_Int64
	CreatedAt        whereHelpertime_Time
}{
	PublisherID:      whereHelperint64{field: "\"publisher_i18n\".\"publisher_id\""},
	Language:         whereHelperstring{field: "\"publisher_i18n\".\"language\""},
	OriginalLanguage: whereHelpernull_String{field: "\"publisher_i18n\".\"original_language\""},
	Name:             whereHelpernull_String{field: "\"publisher_i18n\".\"name\""},
	Description:      whereHelpernull_String{field: "\"publisher_i18n\".\"description\""},
	UserID:           whereHelpernull_Int64{field: "\"publisher_i18n\".\"user_id\""},
	CreatedAt:        whereHelpertime_Time{field: "\"publisher_i18n\".\"created_at\""},
}

// PublisherI18nRels is where relationship names are stored.
var PublisherI18nRels = struct {
	Publisher string
	User      string
}{
	Publisher: "Publisher",
	User:      "User",
}

// publisherI18nR is where relationships are stored.
type publisherI18nR struct {
	Publisher *Publisher
	User      *User
}

// NewStruct creates a new relationship struct
func (*publisherI18nR) NewStruct() *publisherI18nR {
	return &publisherI18nR{}
}

// publisherI18nL is where Load methods for each relationship are stored.
type publisherI18nL struct{}

var (
	publisherI18nAllColumns            = []string{"publisher_id", "language", "original_language", "name", "description", "user_id", "created_at"}
	publisherI18nColumnsWithoutDefault = []string{"publisher_id", "language", "original_language", "name", "description", "user_id"}
	publisherI18nColumnsWithDefault    = []string{"created_at"}
	publisherI18nPrimaryKeyColumns     = []string{"publisher_id", "language"}
)

type (
	// PublisherI18nSlice is an alias for a slice of pointers to PublisherI18n.
	// This should generally be used opposed to []PublisherI18n.
	PublisherI18nSlice []*PublisherI18n

	publisherI18nQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	publisherI18nType                 = reflect.TypeOf(&PublisherI18n{})
	publisherI18nMapping              = queries.MakeStructMapping(publisherI18nType)
	publisherI18nPrimaryKeyMapping, _ = queries.BindMapping(publisherI18nType, publisherI18nMapping, publisherI18nPrimaryKeyColumns)
	publisherI18nInsertCacheMut       sync.RWMutex
	publisherI18nInsertCache          = make(map[string]insertCache)
	publisherI18nUpdateCacheMut       sync.RWMutex
	publisherI18nUpdateCache          = make(map[string]updateCache)
	publisherI18nUpsertCacheMut       sync.RWMutex
	publisherI18nUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single publisherI18n record from the query.
func (q publisherI18nQuery) One(exec boil.Executor) (*PublisherI18n, error) {
	o := &PublisherI18n{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for publisher_i18n")
	}

	return o, nil
}

// All returns all PublisherI18n records from the query.
func (q publisherI18nQuery) All(exec boil.Executor) (PublisherI18nSlice, error) {
	var o []*PublisherI18n

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PublisherI18n slice")
	}

	return o, nil
}

// Count returns the count of all PublisherI18n records in the query.
func (q publisherI18nQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count publisher_i18n rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q publisherI18nQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if publisher_i18n exists")
	}

	return count > 0, nil
}

// Publisher pointed to by the foreign key.
func (o *PublisherI18n) Publisher(mods ...qm.QueryMod) publisherQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.PublisherID),
	}

	queryMods = append(queryMods, mods...)

	query := Publishers(queryMods...)
	queries.SetFrom(query.Query, "\"publishers\"")

	return query
}

// User pointed to by the foreign key.
func (o *PublisherI18n) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// LoadPublisher allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (publisherI18nL) LoadPublisher(e boil.Executor, singular bool, maybePublisherI18n interface{}, mods queries.Applicator) error {
	var slice []*PublisherI18n
	var object *PublisherI18n

	if singular {
		object = maybePublisherI18n.(*PublisherI18n)
	} else {
		slice = *maybePublisherI18n.(*[]*PublisherI18n)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &publisherI18nR{}
		}
		args = append(args, object.PublisherID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &publisherI18nR{}
			}

			for _, a := range args {
				if a == obj.PublisherID {
					continue Outer
				}
			}

			args = append(args, obj.PublisherID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`publishers`), qm.WhereIn(`publishers.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Publisher")
	}

	var resultSlice []*Publisher
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Publisher")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for publishers")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for publishers")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Publisher = foreign
		if foreign.R == nil {
			foreign.R = &publisherR{}
		}
		foreign.R.PublisherI18ns = append(foreign.R.PublisherI18ns, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PublisherID == foreign.ID {
				local.R.Publisher = foreign
				if foreign.R == nil {
					foreign.R = &publisherR{}
				}
				foreign.R.PublisherI18ns = append(foreign.R.PublisherI18ns, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (publisherI18nL) LoadUser(e boil.Executor, singular bool, maybePublisherI18n interface{}, mods queries.Applicator) error {
	var slice []*PublisherI18n
	var object *PublisherI18n

	if singular {
		object = maybePublisherI18n.(*PublisherI18n)
	} else {
		slice = *maybePublisherI18n.(*[]*PublisherI18n)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &publisherI18nR{}
		}
		if !queries.IsNil(object.UserID) {
			args = append(args, object.UserID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &publisherI18nR{}
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
		foreign.R.PublisherI18ns = append(foreign.R.PublisherI18ns, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.UserID, foreign.ID) {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.PublisherI18ns = append(foreign.R.PublisherI18ns, local)
				break
			}
		}
	}

	return nil
}

// SetPublisher of the publisherI18n to the related item.
// Sets o.R.Publisher to related.
// Adds o to related.R.PublisherI18ns.
func (o *PublisherI18n) SetPublisher(exec boil.Executor, insert bool, related *Publisher) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"publisher_i18n\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"publisher_id"}),
		strmangle.WhereClause("\"", "\"", 2, publisherI18nPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.PublisherID, o.Language}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PublisherID = related.ID
	if o.R == nil {
		o.R = &publisherI18nR{
			Publisher: related,
		}
	} else {
		o.R.Publisher = related
	}

	if related.R == nil {
		related.R = &publisherR{
			PublisherI18ns: PublisherI18nSlice{o},
		}
	} else {
		related.R.PublisherI18ns = append(related.R.PublisherI18ns, o)
	}

	return nil
}

// SetUser of the publisherI18n to the related item.
// Sets o.R.User to related.
// Adds o to related.R.PublisherI18ns.
func (o *PublisherI18n) SetUser(exec boil.Executor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"publisher_i18n\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, publisherI18nPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.PublisherID, o.Language}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.UserID, related.ID)
	if o.R == nil {
		o.R = &publisherI18nR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			PublisherI18ns: PublisherI18nSlice{o},
		}
	} else {
		related.R.PublisherI18ns = append(related.R.PublisherI18ns, o)
	}

	return nil
}

// RemoveUser relationship.
// Sets o.R.User to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *PublisherI18n) RemoveUser(exec boil.Executor, related *User) error {
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

	for i, ri := range related.R.PublisherI18ns {
		if queries.Equal(o.UserID, ri.UserID) {
			continue
		}

		ln := len(related.R.PublisherI18ns)
		if ln > 1 && i < ln-1 {
			related.R.PublisherI18ns[i] = related.R.PublisherI18ns[ln-1]
		}
		related.R.PublisherI18ns = related.R.PublisherI18ns[:ln-1]
		break
	}
	return nil
}

// PublisherI18ns retrieves all the records using an executor.
func PublisherI18ns(mods ...qm.QueryMod) publisherI18nQuery {
	mods = append(mods, qm.From("\"publisher_i18n\""))
	return publisherI18nQuery{NewQuery(mods...)}
}

// FindPublisherI18n retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPublisherI18n(exec boil.Executor, publisherID int64, language string, selectCols ...string) (*PublisherI18n, error) {
	publisherI18nObj := &PublisherI18n{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"publisher_i18n\" where \"publisher_id\"=$1 AND \"language\"=$2", sel,
	)

	q := queries.Raw(query, publisherID, language)

	err := q.Bind(nil, exec, publisherI18nObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from publisher_i18n")
	}

	return publisherI18nObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PublisherI18n) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no publisher_i18n provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(publisherI18nColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	publisherI18nInsertCacheMut.RLock()
	cache, cached := publisherI18nInsertCache[key]
	publisherI18nInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			publisherI18nAllColumns,
			publisherI18nColumnsWithDefault,
			publisherI18nColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(publisherI18nType, publisherI18nMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(publisherI18nType, publisherI18nMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"publisher_i18n\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"publisher_i18n\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into publisher_i18n")
	}

	if !cached {
		publisherI18nInsertCacheMut.Lock()
		publisherI18nInsertCache[key] = cache
		publisherI18nInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the PublisherI18n.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PublisherI18n) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	publisherI18nUpdateCacheMut.RLock()
	cache, cached := publisherI18nUpdateCache[key]
	publisherI18nUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			publisherI18nAllColumns,
			publisherI18nPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update publisher_i18n, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"publisher_i18n\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, publisherI18nPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(publisherI18nType, publisherI18nMapping, append(wl, publisherI18nPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update publisher_i18n row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for publisher_i18n")
	}

	if !cached {
		publisherI18nUpdateCacheMut.Lock()
		publisherI18nUpdateCache[key] = cache
		publisherI18nUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q publisherI18nQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for publisher_i18n")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for publisher_i18n")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PublisherI18nSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), publisherI18nPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"publisher_i18n\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, publisherI18nPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in publisherI18n slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all publisherI18n")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PublisherI18n) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no publisher_i18n provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(publisherI18nColumnsWithDefault, o)

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

	publisherI18nUpsertCacheMut.RLock()
	cache, cached := publisherI18nUpsertCache[key]
	publisherI18nUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			publisherI18nAllColumns,
			publisherI18nColumnsWithDefault,
			publisherI18nColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			publisherI18nAllColumns,
			publisherI18nPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert publisher_i18n, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(publisherI18nPrimaryKeyColumns))
			copy(conflict, publisherI18nPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"publisher_i18n\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(publisherI18nType, publisherI18nMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(publisherI18nType, publisherI18nMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert publisher_i18n")
	}

	if !cached {
		publisherI18nUpsertCacheMut.Lock()
		publisherI18nUpsertCache[key] = cache
		publisherI18nUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single PublisherI18n record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PublisherI18n) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PublisherI18n provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), publisherI18nPrimaryKeyMapping)
	sql := "DELETE FROM \"publisher_i18n\" WHERE \"publisher_id\"=$1 AND \"language\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from publisher_i18n")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for publisher_i18n")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q publisherI18nQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no publisherI18nQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from publisher_i18n")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for publisher_i18n")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PublisherI18nSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), publisherI18nPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"publisher_i18n\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, publisherI18nPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from publisherI18n slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for publisher_i18n")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PublisherI18n) Reload(exec boil.Executor) error {
	ret, err := FindPublisherI18n(exec, o.PublisherID, o.Language)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PublisherI18nSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PublisherI18nSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), publisherI18nPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"publisher_i18n\".* FROM \"publisher_i18n\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, publisherI18nPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PublisherI18nSlice")
	}

	*o = slice

	return nil
}

// PublisherI18nExists checks if the PublisherI18n row exists.
func PublisherI18nExists(exec boil.Executor, publisherID int64, language string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"publisher_i18n\" where \"publisher_id\"=$1 AND \"language\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, publisherID, language)
	}
	row := exec.QueryRow(sql, publisherID, language)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if publisher_i18n exists")
	}

	return exists, nil
}
