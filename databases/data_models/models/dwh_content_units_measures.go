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

// DWHContentUnitsMeasure is an object representing the database table.
type DWHContentUnitsMeasure struct {
	EventUnitUID                string     `boil:"event_unit_uid" json:"event_unit_uid" toml:"event_unit_uid" yaml:"event_unit_uid"`
	AllEventsCount              null.Int64 `boil:"all_events_count" json:"all_events_count,omitempty" toml:"all_events_count" yaml:"all_events_count,omitempty"`
	UniqueUsersCount            null.Int64 `boil:"unique_users_count" json:"unique_users_count,omitempty" toml:"unique_users_count" yaml:"unique_users_count,omitempty"`
	DWHUpdateDatetime           null.Time  `boil:"dwh_update_datetime" json:"dwh_update_datetime,omitempty" toml:"dwh_update_datetime" yaml:"dwh_update_datetime,omitempty"`
	UniqueUsersWatchingNowCount null.Int64 `boil:"unique_users_watching_now_count" json:"unique_users_watching_now_count,omitempty" toml:"unique_users_watching_now_count" yaml:"unique_users_watching_now_count,omitempty"`

	R *dwhContentUnitsMeasureR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dwhContentUnitsMeasureL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DWHContentUnitsMeasureColumns = struct {
	EventUnitUID                string
	AllEventsCount              string
	UniqueUsersCount            string
	DWHUpdateDatetime           string
	UniqueUsersWatchingNowCount string
}{
	EventUnitUID:                "event_unit_uid",
	AllEventsCount:              "all_events_count",
	UniqueUsersCount:            "unique_users_count",
	DWHUpdateDatetime:           "dwh_update_datetime",
	UniqueUsersWatchingNowCount: "unique_users_watching_now_count",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var DWHContentUnitsMeasureWhere = struct {
	EventUnitUID                whereHelperstring
	AllEventsCount              whereHelpernull_Int64
	UniqueUsersCount            whereHelpernull_Int64
	DWHUpdateDatetime           whereHelpernull_Time
	UniqueUsersWatchingNowCount whereHelpernull_Int64
}{
	EventUnitUID:                whereHelperstring{field: "\"dwh_content_units_measures\".\"event_unit_uid\""},
	AllEventsCount:              whereHelpernull_Int64{field: "\"dwh_content_units_measures\".\"all_events_count\""},
	UniqueUsersCount:            whereHelpernull_Int64{field: "\"dwh_content_units_measures\".\"unique_users_count\""},
	DWHUpdateDatetime:           whereHelpernull_Time{field: "\"dwh_content_units_measures\".\"dwh_update_datetime\""},
	UniqueUsersWatchingNowCount: whereHelpernull_Int64{field: "\"dwh_content_units_measures\".\"unique_users_watching_now_count\""},
}

// DWHContentUnitsMeasureRels is where relationship names are stored.
var DWHContentUnitsMeasureRels = struct {
}{}

// dwhContentUnitsMeasureR is where relationships are stored.
type dwhContentUnitsMeasureR struct {
}

// NewStruct creates a new relationship struct
func (*dwhContentUnitsMeasureR) NewStruct() *dwhContentUnitsMeasureR {
	return &dwhContentUnitsMeasureR{}
}

// dwhContentUnitsMeasureL is where Load methods for each relationship are stored.
type dwhContentUnitsMeasureL struct{}

var (
	dwhContentUnitsMeasureAllColumns            = []string{"event_unit_uid", "all_events_count", "unique_users_count", "dwh_update_datetime", "unique_users_watching_now_count"}
	dwhContentUnitsMeasureColumnsWithoutDefault = []string{"event_unit_uid", "all_events_count", "unique_users_count", "dwh_update_datetime", "unique_users_watching_now_count"}
	dwhContentUnitsMeasureColumnsWithDefault    = []string{}
	dwhContentUnitsMeasurePrimaryKeyColumns     = []string{"event_unit_uid"}
)

type (
	// DWHContentUnitsMeasureSlice is an alias for a slice of pointers to DWHContentUnitsMeasure.
	// This should generally be used opposed to []DWHContentUnitsMeasure.
	DWHContentUnitsMeasureSlice []*DWHContentUnitsMeasure

	dwhContentUnitsMeasureQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dwhContentUnitsMeasureType                 = reflect.TypeOf(&DWHContentUnitsMeasure{})
	dwhContentUnitsMeasureMapping              = queries.MakeStructMapping(dwhContentUnitsMeasureType)
	dwhContentUnitsMeasurePrimaryKeyMapping, _ = queries.BindMapping(dwhContentUnitsMeasureType, dwhContentUnitsMeasureMapping, dwhContentUnitsMeasurePrimaryKeyColumns)
	dwhContentUnitsMeasureInsertCacheMut       sync.RWMutex
	dwhContentUnitsMeasureInsertCache          = make(map[string]insertCache)
	dwhContentUnitsMeasureUpdateCacheMut       sync.RWMutex
	dwhContentUnitsMeasureUpdateCache          = make(map[string]updateCache)
	dwhContentUnitsMeasureUpsertCacheMut       sync.RWMutex
	dwhContentUnitsMeasureUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single dwhContentUnitsMeasure record from the query.
func (q dwhContentUnitsMeasureQuery) One(exec boil.Executor) (*DWHContentUnitsMeasure, error) {
	o := &DWHContentUnitsMeasure{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for dwh_content_units_measures")
	}

	return o, nil
}

// All returns all DWHContentUnitsMeasure records from the query.
func (q dwhContentUnitsMeasureQuery) All(exec boil.Executor) (DWHContentUnitsMeasureSlice, error) {
	var o []*DWHContentUnitsMeasure

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DWHContentUnitsMeasure slice")
	}

	return o, nil
}

// Count returns the count of all DWHContentUnitsMeasure records in the query.
func (q dwhContentUnitsMeasureQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count dwh_content_units_measures rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dwhContentUnitsMeasureQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if dwh_content_units_measures exists")
	}

	return count > 0, nil
}

// DWHContentUnitsMeasures retrieves all the records using an executor.
func DWHContentUnitsMeasures(mods ...qm.QueryMod) dwhContentUnitsMeasureQuery {
	mods = append(mods, qm.From("\"dwh_content_units_measures\""))
	return dwhContentUnitsMeasureQuery{NewQuery(mods...)}
}

// FindDWHContentUnitsMeasure retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDWHContentUnitsMeasure(exec boil.Executor, eventUnitUID string, selectCols ...string) (*DWHContentUnitsMeasure, error) {
	dwhContentUnitsMeasureObj := &DWHContentUnitsMeasure{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"dwh_content_units_measures\" where \"event_unit_uid\"=$1", sel,
	)

	q := queries.Raw(query, eventUnitUID)

	err := q.Bind(nil, exec, dwhContentUnitsMeasureObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from dwh_content_units_measures")
	}

	return dwhContentUnitsMeasureObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DWHContentUnitsMeasure) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no dwh_content_units_measures provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(dwhContentUnitsMeasureColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dwhContentUnitsMeasureInsertCacheMut.RLock()
	cache, cached := dwhContentUnitsMeasureInsertCache[key]
	dwhContentUnitsMeasureInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dwhContentUnitsMeasureAllColumns,
			dwhContentUnitsMeasureColumnsWithDefault,
			dwhContentUnitsMeasureColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dwhContentUnitsMeasureType, dwhContentUnitsMeasureMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dwhContentUnitsMeasureType, dwhContentUnitsMeasureMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"dwh_content_units_measures\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"dwh_content_units_measures\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into dwh_content_units_measures")
	}

	if !cached {
		dwhContentUnitsMeasureInsertCacheMut.Lock()
		dwhContentUnitsMeasureInsertCache[key] = cache
		dwhContentUnitsMeasureInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the DWHContentUnitsMeasure.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DWHContentUnitsMeasure) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	dwhContentUnitsMeasureUpdateCacheMut.RLock()
	cache, cached := dwhContentUnitsMeasureUpdateCache[key]
	dwhContentUnitsMeasureUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dwhContentUnitsMeasureAllColumns,
			dwhContentUnitsMeasurePrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update dwh_content_units_measures, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"dwh_content_units_measures\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, dwhContentUnitsMeasurePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dwhContentUnitsMeasureType, dwhContentUnitsMeasureMapping, append(wl, dwhContentUnitsMeasurePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update dwh_content_units_measures row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for dwh_content_units_measures")
	}

	if !cached {
		dwhContentUnitsMeasureUpdateCacheMut.Lock()
		dwhContentUnitsMeasureUpdateCache[key] = cache
		dwhContentUnitsMeasureUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q dwhContentUnitsMeasureQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for dwh_content_units_measures")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for dwh_content_units_measures")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DWHContentUnitsMeasureSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dwhContentUnitsMeasurePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"dwh_content_units_measures\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, dwhContentUnitsMeasurePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in dwhContentUnitsMeasure slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all dwhContentUnitsMeasure")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DWHContentUnitsMeasure) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no dwh_content_units_measures provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(dwhContentUnitsMeasureColumnsWithDefault, o)

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

	dwhContentUnitsMeasureUpsertCacheMut.RLock()
	cache, cached := dwhContentUnitsMeasureUpsertCache[key]
	dwhContentUnitsMeasureUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dwhContentUnitsMeasureAllColumns,
			dwhContentUnitsMeasureColumnsWithDefault,
			dwhContentUnitsMeasureColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			dwhContentUnitsMeasureAllColumns,
			dwhContentUnitsMeasurePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert dwh_content_units_measures, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(dwhContentUnitsMeasurePrimaryKeyColumns))
			copy(conflict, dwhContentUnitsMeasurePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"dwh_content_units_measures\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(dwhContentUnitsMeasureType, dwhContentUnitsMeasureMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dwhContentUnitsMeasureType, dwhContentUnitsMeasureMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert dwh_content_units_measures")
	}

	if !cached {
		dwhContentUnitsMeasureUpsertCacheMut.Lock()
		dwhContentUnitsMeasureUpsertCache[key] = cache
		dwhContentUnitsMeasureUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single DWHContentUnitsMeasure record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DWHContentUnitsMeasure) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DWHContentUnitsMeasure provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dwhContentUnitsMeasurePrimaryKeyMapping)
	sql := "DELETE FROM \"dwh_content_units_measures\" WHERE \"event_unit_uid\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from dwh_content_units_measures")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for dwh_content_units_measures")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q dwhContentUnitsMeasureQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no dwhContentUnitsMeasureQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from dwh_content_units_measures")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for dwh_content_units_measures")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DWHContentUnitsMeasureSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dwhContentUnitsMeasurePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"dwh_content_units_measures\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, dwhContentUnitsMeasurePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from dwhContentUnitsMeasure slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for dwh_content_units_measures")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DWHContentUnitsMeasure) Reload(exec boil.Executor) error {
	ret, err := FindDWHContentUnitsMeasure(exec, o.EventUnitUID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DWHContentUnitsMeasureSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DWHContentUnitsMeasureSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dwhContentUnitsMeasurePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"dwh_content_units_measures\".* FROM \"dwh_content_units_measures\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, dwhContentUnitsMeasurePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DWHContentUnitsMeasureSlice")
	}

	*o = slice

	return nil
}

// DWHContentUnitsMeasureExists checks if the DWHContentUnitsMeasure row exists.
func DWHContentUnitsMeasureExists(exec boil.Executor, eventUnitUID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"dwh_content_units_measures\" where \"event_unit_uid\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, eventUnitUID)
	}
	row := exec.QueryRow(sql, eventUnitUID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if dwh_content_units_measures exists")
	}

	return exists, nil
}
