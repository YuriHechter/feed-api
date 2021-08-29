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

// Publisher is an object representing the database table.
type Publisher struct {
	ID      int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	UID     string      `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	Pattern null.String `boil:"pattern" json:"pattern,omitempty" toml:"pattern" yaml:"pattern,omitempty"`

	R *publisherR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L publisherL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PublisherColumns = struct {
	ID      string
	UID     string
	Pattern string
}{
	ID:      "id",
	UID:     "uid",
	Pattern: "pattern",
}

// Generated where

var PublisherWhere = struct {
	ID      whereHelperint64
	UID     whereHelperstring
	Pattern whereHelpernull_String
}{
	ID:      whereHelperint64{field: "\"publishers\".\"id\""},
	UID:     whereHelperstring{field: "\"publishers\".\"uid\""},
	Pattern: whereHelpernull_String{field: "\"publishers\".\"pattern\""},
}

// PublisherRels is where relationship names are stored.
var PublisherRels = struct {
	ContentUnits   string
	PublisherI18ns string
}{
	ContentUnits:   "ContentUnits",
	PublisherI18ns: "PublisherI18ns",
}

// publisherR is where relationships are stored.
type publisherR struct {
	ContentUnits   ContentUnitSlice
	PublisherI18ns PublisherI18nSlice
}

// NewStruct creates a new relationship struct
func (*publisherR) NewStruct() *publisherR {
	return &publisherR{}
}

// publisherL is where Load methods for each relationship are stored.
type publisherL struct{}

var (
	publisherAllColumns            = []string{"id", "uid", "pattern"}
	publisherColumnsWithoutDefault = []string{"uid", "pattern"}
	publisherColumnsWithDefault    = []string{"id"}
	publisherPrimaryKeyColumns     = []string{"id"}
)

type (
	// PublisherSlice is an alias for a slice of pointers to Publisher.
	// This should generally be used opposed to []Publisher.
	PublisherSlice []*Publisher

	publisherQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	publisherType                 = reflect.TypeOf(&Publisher{})
	publisherMapping              = queries.MakeStructMapping(publisherType)
	publisherPrimaryKeyMapping, _ = queries.BindMapping(publisherType, publisherMapping, publisherPrimaryKeyColumns)
	publisherInsertCacheMut       sync.RWMutex
	publisherInsertCache          = make(map[string]insertCache)
	publisherUpdateCacheMut       sync.RWMutex
	publisherUpdateCache          = make(map[string]updateCache)
	publisherUpsertCacheMut       sync.RWMutex
	publisherUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single publisher record from the query.
func (q publisherQuery) One(exec boil.Executor) (*Publisher, error) {
	o := &Publisher{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for publishers")
	}

	return o, nil
}

// All returns all Publisher records from the query.
func (q publisherQuery) All(exec boil.Executor) (PublisherSlice, error) {
	var o []*Publisher

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Publisher slice")
	}

	return o, nil
}

// Count returns the count of all Publisher records in the query.
func (q publisherQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count publishers rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q publisherQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if publishers exists")
	}

	return count > 0, nil
}

// ContentUnits retrieves all the content_unit's ContentUnits with an executor.
func (o *Publisher) ContentUnits(mods ...qm.QueryMod) contentUnitQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("\"content_units_publishers\" on \"content_units\".\"id\" = \"content_units_publishers\".\"content_unit_id\""),
		qm.Where("\"content_units_publishers\".\"publisher_id\"=?", o.ID),
	)

	query := ContentUnits(queryMods...)
	queries.SetFrom(query.Query, "\"content_units\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"content_units\".*"})
	}

	return query
}

// PublisherI18ns retrieves all the publisher_i18n's PublisherI18ns with an executor.
func (o *Publisher) PublisherI18ns(mods ...qm.QueryMod) publisherI18nQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"publisher_i18n\".\"publisher_id\"=?", o.ID),
	)

	query := PublisherI18ns(queryMods...)
	queries.SetFrom(query.Query, "\"publisher_i18n\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"publisher_i18n\".*"})
	}

	return query
}

// LoadContentUnits allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (publisherL) LoadContentUnits(e boil.Executor, singular bool, maybePublisher interface{}, mods queries.Applicator) error {
	var slice []*Publisher
	var object *Publisher

	if singular {
		object = maybePublisher.(*Publisher)
	} else {
		slice = *maybePublisher.(*[]*Publisher)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &publisherR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &publisherR{}
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
		qm.Select("\"content_units\".*, \"a\".\"publisher_id\""),
		qm.From("\"content_units\""),
		qm.InnerJoin("\"content_units_publishers\" as \"a\" on \"content_units\".\"id\" = \"a\".\"content_unit_id\""),
		qm.WhereIn("\"a\".\"publisher_id\" in ?", args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load content_units")
	}

	var resultSlice []*ContentUnit

	var localJoinCols []int64
	for results.Next() {
		one := new(ContentUnit)
		var localJoinCol int64

		err = results.Scan(&one.ID, &one.UID, &one.TypeID, &one.CreatedAt, &one.Properties, &one.Secure, &one.Published, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for content_units")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice content_units")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on content_units")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for content_units")
	}

	if singular {
		object.R.ContentUnits = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &contentUnitR{}
			}
			foreign.R.Publishers = append(foreign.R.Publishers, object)
		}
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.ID == localJoinCol {
				local.R.ContentUnits = append(local.R.ContentUnits, foreign)
				if foreign.R == nil {
					foreign.R = &contentUnitR{}
				}
				foreign.R.Publishers = append(foreign.R.Publishers, local)
				break
			}
		}
	}

	return nil
}

// LoadPublisherI18ns allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (publisherL) LoadPublisherI18ns(e boil.Executor, singular bool, maybePublisher interface{}, mods queries.Applicator) error {
	var slice []*Publisher
	var object *Publisher

	if singular {
		object = maybePublisher.(*Publisher)
	} else {
		slice = *maybePublisher.(*[]*Publisher)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &publisherR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &publisherR{}
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

	query := NewQuery(qm.From(`publisher_i18n`), qm.WhereIn(`publisher_i18n.publisher_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load publisher_i18n")
	}

	var resultSlice []*PublisherI18n
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice publisher_i18n")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on publisher_i18n")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for publisher_i18n")
	}

	if singular {
		object.R.PublisherI18ns = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &publisherI18nR{}
			}
			foreign.R.Publisher = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.PublisherID {
				local.R.PublisherI18ns = append(local.R.PublisherI18ns, foreign)
				if foreign.R == nil {
					foreign.R = &publisherI18nR{}
				}
				foreign.R.Publisher = local
				break
			}
		}
	}

	return nil
}

// AddContentUnits adds the given related objects to the existing relationships
// of the publisher, optionally inserting them as new records.
// Appends related to o.R.ContentUnits.
// Sets related.R.Publishers appropriately.
func (o *Publisher) AddContentUnits(exec boil.Executor, insert bool, related ...*ContentUnit) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into \"content_units_publishers\" (\"publisher_id\", \"content_unit_id\") values ($1, $2)"
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
		o.R = &publisherR{
			ContentUnits: related,
		}
	} else {
		o.R.ContentUnits = append(o.R.ContentUnits, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &contentUnitR{
				Publishers: PublisherSlice{o},
			}
		} else {
			rel.R.Publishers = append(rel.R.Publishers, o)
		}
	}
	return nil
}

// SetContentUnits removes all previously related items of the
// publisher replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Publishers's ContentUnits accordingly.
// Replaces o.R.ContentUnits with related.
// Sets related.R.Publishers's ContentUnits accordingly.
func (o *Publisher) SetContentUnits(exec boil.Executor, insert bool, related ...*ContentUnit) error {
	query := "delete from \"content_units_publishers\" where \"publisher_id\" = $1"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	removeContentUnitsFromPublishersSlice(o, related)
	if o.R != nil {
		o.R.ContentUnits = nil
	}
	return o.AddContentUnits(exec, insert, related...)
}

// RemoveContentUnits relationships from objects passed in.
// Removes related items from R.ContentUnits (uses pointer comparison, removal does not keep order)
// Sets related.R.Publishers.
func (o *Publisher) RemoveContentUnits(exec boil.Executor, related ...*ContentUnit) error {
	var err error
	query := fmt.Sprintf(
		"delete from \"content_units_publishers\" where \"publisher_id\" = $1 and \"content_unit_id\" in (%s)",
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
	removeContentUnitsFromPublishersSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.ContentUnits {
			if rel != ri {
				continue
			}

			ln := len(o.R.ContentUnits)
			if ln > 1 && i < ln-1 {
				o.R.ContentUnits[i] = o.R.ContentUnits[ln-1]
			}
			o.R.ContentUnits = o.R.ContentUnits[:ln-1]
			break
		}
	}

	return nil
}

func removeContentUnitsFromPublishersSlice(o *Publisher, related []*ContentUnit) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.Publishers {
			if o.ID != ri.ID {
				continue
			}

			ln := len(rel.R.Publishers)
			if ln > 1 && i < ln-1 {
				rel.R.Publishers[i] = rel.R.Publishers[ln-1]
			}
			rel.R.Publishers = rel.R.Publishers[:ln-1]
			break
		}
	}
}

// AddPublisherI18ns adds the given related objects to the existing relationships
// of the publisher, optionally inserting them as new records.
// Appends related to o.R.PublisherI18ns.
// Sets related.R.Publisher appropriately.
func (o *Publisher) AddPublisherI18ns(exec boil.Executor, insert bool, related ...*PublisherI18n) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.PublisherID = o.ID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"publisher_i18n\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"publisher_id"}),
				strmangle.WhereClause("\"", "\"", 2, publisherI18nPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.PublisherID, rel.Language}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}
			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.PublisherID = o.ID
		}
	}

	if o.R == nil {
		o.R = &publisherR{
			PublisherI18ns: related,
		}
	} else {
		o.R.PublisherI18ns = append(o.R.PublisherI18ns, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &publisherI18nR{
				Publisher: o,
			}
		} else {
			rel.R.Publisher = o
		}
	}
	return nil
}

// Publishers retrieves all the records using an executor.
func Publishers(mods ...qm.QueryMod) publisherQuery {
	mods = append(mods, qm.From("\"publishers\""))
	return publisherQuery{NewQuery(mods...)}
}

// FindPublisher retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPublisher(exec boil.Executor, iD int64, selectCols ...string) (*Publisher, error) {
	publisherObj := &Publisher{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"publishers\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, publisherObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from publishers")
	}

	return publisherObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Publisher) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no publishers provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(publisherColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	publisherInsertCacheMut.RLock()
	cache, cached := publisherInsertCache[key]
	publisherInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			publisherAllColumns,
			publisherColumnsWithDefault,
			publisherColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(publisherType, publisherMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(publisherType, publisherMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"publishers\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"publishers\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into publishers")
	}

	if !cached {
		publisherInsertCacheMut.Lock()
		publisherInsertCache[key] = cache
		publisherInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Publisher.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Publisher) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	publisherUpdateCacheMut.RLock()
	cache, cached := publisherUpdateCache[key]
	publisherUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			publisherAllColumns,
			publisherPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update publishers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"publishers\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, publisherPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(publisherType, publisherMapping, append(wl, publisherPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update publishers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for publishers")
	}

	if !cached {
		publisherUpdateCacheMut.Lock()
		publisherUpdateCache[key] = cache
		publisherUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q publisherQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for publishers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for publishers")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PublisherSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), publisherPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"publishers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, publisherPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in publisher slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all publisher")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Publisher) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no publishers provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(publisherColumnsWithDefault, o)

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

	publisherUpsertCacheMut.RLock()
	cache, cached := publisherUpsertCache[key]
	publisherUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			publisherAllColumns,
			publisherColumnsWithDefault,
			publisherColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			publisherAllColumns,
			publisherPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert publishers, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(publisherPrimaryKeyColumns))
			copy(conflict, publisherPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"publishers\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(publisherType, publisherMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(publisherType, publisherMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert publishers")
	}

	if !cached {
		publisherUpsertCacheMut.Lock()
		publisherUpsertCache[key] = cache
		publisherUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Publisher record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Publisher) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Publisher provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), publisherPrimaryKeyMapping)
	sql := "DELETE FROM \"publishers\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from publishers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for publishers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q publisherQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no publisherQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from publishers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for publishers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PublisherSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), publisherPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"publishers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, publisherPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from publisher slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for publishers")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Publisher) Reload(exec boil.Executor) error {
	ret, err := FindPublisher(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PublisherSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PublisherSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), publisherPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"publishers\".* FROM \"publishers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, publisherPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PublisherSlice")
	}

	*o = slice

	return nil
}

// PublisherExists checks if the Publisher row exists.
func PublisherExists(exec boil.Executor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"publishers\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if publishers exists")
	}

	return exists, nil
}
