// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package gnucash

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// TaxtableEntry is an object representing the database table.
type TaxtableEntry struct {
	ID          int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	Taxtable    string `boil:"taxtable" json:"taxtable" toml:"taxtable" yaml:"taxtable"`
	Account     string `boil:"account" json:"account" toml:"account" yaml:"account"`
	AmountNum   int64  `boil:"amount_num" json:"amount_num" toml:"amount_num" yaml:"amount_num"`
	AmountDenom int64  `boil:"amount_denom" json:"amount_denom" toml:"amount_denom" yaml:"amount_denom"`
	Type        int64  `boil:"type" json:"type" toml:"type" yaml:"type"`

	R *taxtableEntryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L taxtableEntryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TaxtableEntryColumns = struct {
	ID          string
	Taxtable    string
	Account     string
	AmountNum   string
	AmountDenom string
	Type        string
}{
	ID:          "id",
	Taxtable:    "taxtable",
	Account:     "account",
	AmountNum:   "amount_num",
	AmountDenom: "amount_denom",
	Type:        "type",
}

var TaxtableEntryTableColumns = struct {
	ID          string
	Taxtable    string
	Account     string
	AmountNum   string
	AmountDenom string
	Type        string
}{
	ID:          "taxtable_entries.id",
	Taxtable:    "taxtable_entries.taxtable",
	Account:     "taxtable_entries.account",
	AmountNum:   "taxtable_entries.amount_num",
	AmountDenom: "taxtable_entries.amount_denom",
	Type:        "taxtable_entries.type",
}

// Generated where

var TaxtableEntryWhere = struct {
	ID          whereHelperint64
	Taxtable    whereHelperstring
	Account     whereHelperstring
	AmountNum   whereHelperint64
	AmountDenom whereHelperint64
	Type        whereHelperint64
}{
	ID:          whereHelperint64{field: "\"taxtable_entries\".\"id\""},
	Taxtable:    whereHelperstring{field: "\"taxtable_entries\".\"taxtable\""},
	Account:     whereHelperstring{field: "\"taxtable_entries\".\"account\""},
	AmountNum:   whereHelperint64{field: "\"taxtable_entries\".\"amount_num\""},
	AmountDenom: whereHelperint64{field: "\"taxtable_entries\".\"amount_denom\""},
	Type:        whereHelperint64{field: "\"taxtable_entries\".\"type\""},
}

// TaxtableEntryRels is where relationship names are stored.
var TaxtableEntryRels = struct {
}{}

// taxtableEntryR is where relationships are stored.
type taxtableEntryR struct {
}

// NewStruct creates a new relationship struct
func (*taxtableEntryR) NewStruct() *taxtableEntryR {
	return &taxtableEntryR{}
}

// taxtableEntryL is where Load methods for each relationship are stored.
type taxtableEntryL struct{}

var (
	taxtableEntryAllColumns            = []string{"id", "taxtable", "account", "amount_num", "amount_denom", "type"}
	taxtableEntryColumnsWithoutDefault = []string{"taxtable", "account", "amount_num", "amount_denom", "type"}
	taxtableEntryColumnsWithDefault    = []string{"id"}
	taxtableEntryPrimaryKeyColumns     = []string{"id"}
	taxtableEntryGeneratedColumns      = []string{"id"}
)

type (
	// TaxtableEntrySlice is an alias for a slice of pointers to TaxtableEntry.
	// This should almost always be used instead of []TaxtableEntry.
	TaxtableEntrySlice []*TaxtableEntry
	// TaxtableEntryHook is the signature for custom TaxtableEntry hook methods
	TaxtableEntryHook func(context.Context, boil.ContextExecutor, *TaxtableEntry) error

	taxtableEntryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	taxtableEntryType                 = reflect.TypeOf(&TaxtableEntry{})
	taxtableEntryMapping              = queries.MakeStructMapping(taxtableEntryType)
	taxtableEntryPrimaryKeyMapping, _ = queries.BindMapping(taxtableEntryType, taxtableEntryMapping, taxtableEntryPrimaryKeyColumns)
	taxtableEntryInsertCacheMut       sync.RWMutex
	taxtableEntryInsertCache          = make(map[string]insertCache)
	taxtableEntryUpdateCacheMut       sync.RWMutex
	taxtableEntryUpdateCache          = make(map[string]updateCache)
	taxtableEntryUpsertCacheMut       sync.RWMutex
	taxtableEntryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var taxtableEntryAfterSelectHooks []TaxtableEntryHook

var taxtableEntryBeforeInsertHooks []TaxtableEntryHook
var taxtableEntryAfterInsertHooks []TaxtableEntryHook

var taxtableEntryBeforeUpdateHooks []TaxtableEntryHook
var taxtableEntryAfterUpdateHooks []TaxtableEntryHook

var taxtableEntryBeforeDeleteHooks []TaxtableEntryHook
var taxtableEntryAfterDeleteHooks []TaxtableEntryHook

var taxtableEntryBeforeUpsertHooks []TaxtableEntryHook
var taxtableEntryAfterUpsertHooks []TaxtableEntryHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TaxtableEntry) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxtableEntryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TaxtableEntry) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxtableEntryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TaxtableEntry) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxtableEntryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TaxtableEntry) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxtableEntryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TaxtableEntry) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxtableEntryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TaxtableEntry) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxtableEntryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TaxtableEntry) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxtableEntryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TaxtableEntry) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxtableEntryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TaxtableEntry) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taxtableEntryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTaxtableEntryHook registers your hook function for all future operations.
func AddTaxtableEntryHook(hookPoint boil.HookPoint, taxtableEntryHook TaxtableEntryHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		taxtableEntryAfterSelectHooks = append(taxtableEntryAfterSelectHooks, taxtableEntryHook)
	case boil.BeforeInsertHook:
		taxtableEntryBeforeInsertHooks = append(taxtableEntryBeforeInsertHooks, taxtableEntryHook)
	case boil.AfterInsertHook:
		taxtableEntryAfterInsertHooks = append(taxtableEntryAfterInsertHooks, taxtableEntryHook)
	case boil.BeforeUpdateHook:
		taxtableEntryBeforeUpdateHooks = append(taxtableEntryBeforeUpdateHooks, taxtableEntryHook)
	case boil.AfterUpdateHook:
		taxtableEntryAfterUpdateHooks = append(taxtableEntryAfterUpdateHooks, taxtableEntryHook)
	case boil.BeforeDeleteHook:
		taxtableEntryBeforeDeleteHooks = append(taxtableEntryBeforeDeleteHooks, taxtableEntryHook)
	case boil.AfterDeleteHook:
		taxtableEntryAfterDeleteHooks = append(taxtableEntryAfterDeleteHooks, taxtableEntryHook)
	case boil.BeforeUpsertHook:
		taxtableEntryBeforeUpsertHooks = append(taxtableEntryBeforeUpsertHooks, taxtableEntryHook)
	case boil.AfterUpsertHook:
		taxtableEntryAfterUpsertHooks = append(taxtableEntryAfterUpsertHooks, taxtableEntryHook)
	}
}

// One returns a single taxtableEntry record from the query.
func (q taxtableEntryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TaxtableEntry, error) {
	o := &TaxtableEntry{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "gnucash: failed to execute a one query for taxtable_entries")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TaxtableEntry records from the query.
func (q taxtableEntryQuery) All(ctx context.Context, exec boil.ContextExecutor) (TaxtableEntrySlice, error) {
	var o []*TaxtableEntry

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "gnucash: failed to assign all query results to TaxtableEntry slice")
	}

	if len(taxtableEntryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TaxtableEntry records in the query.
func (q taxtableEntryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to count taxtable_entries rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q taxtableEntryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "gnucash: failed to check if taxtable_entries exists")
	}

	return count > 0, nil
}

// TaxtableEntries retrieves all the records using an executor.
func TaxtableEntries(mods ...qm.QueryMod) taxtableEntryQuery {
	mods = append(mods, qm.From("\"taxtable_entries\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"taxtable_entries\".*"})
	}

	return taxtableEntryQuery{q}
}

// FindTaxtableEntry retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTaxtableEntry(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*TaxtableEntry, error) {
	taxtableEntryObj := &TaxtableEntry{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"taxtable_entries\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, taxtableEntryObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "gnucash: unable to select from taxtable_entries")
	}

	if err = taxtableEntryObj.doAfterSelectHooks(ctx, exec); err != nil {
		return taxtableEntryObj, err
	}

	return taxtableEntryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TaxtableEntry) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("gnucash: no taxtable_entries provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(taxtableEntryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	taxtableEntryInsertCacheMut.RLock()
	cache, cached := taxtableEntryInsertCache[key]
	taxtableEntryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			taxtableEntryAllColumns,
			taxtableEntryColumnsWithDefault,
			taxtableEntryColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, taxtableEntryGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(taxtableEntryType, taxtableEntryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(taxtableEntryType, taxtableEntryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"taxtable_entries\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"taxtable_entries\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "gnucash: unable to insert into taxtable_entries")
	}

	if !cached {
		taxtableEntryInsertCacheMut.Lock()
		taxtableEntryInsertCache[key] = cache
		taxtableEntryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TaxtableEntry.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TaxtableEntry) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	taxtableEntryUpdateCacheMut.RLock()
	cache, cached := taxtableEntryUpdateCache[key]
	taxtableEntryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			taxtableEntryAllColumns,
			taxtableEntryPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, taxtableEntryGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("gnucash: unable to update taxtable_entries, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"taxtable_entries\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, taxtableEntryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(taxtableEntryType, taxtableEntryMapping, append(wl, taxtableEntryPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to update taxtable_entries row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by update for taxtable_entries")
	}

	if !cached {
		taxtableEntryUpdateCacheMut.Lock()
		taxtableEntryUpdateCache[key] = cache
		taxtableEntryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q taxtableEntryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to update all for taxtable_entries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to retrieve rows affected for taxtable_entries")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TaxtableEntrySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("gnucash: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taxtableEntryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"taxtable_entries\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, taxtableEntryPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to update all in taxtableEntry slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to retrieve rows affected all in update all taxtableEntry")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TaxtableEntry) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("gnucash: no taxtable_entries provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(taxtableEntryColumnsWithDefault, o)

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

	taxtableEntryUpsertCacheMut.RLock()
	cache, cached := taxtableEntryUpsertCache[key]
	taxtableEntryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			taxtableEntryAllColumns,
			taxtableEntryColumnsWithDefault,
			taxtableEntryColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			taxtableEntryAllColumns,
			taxtableEntryPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("gnucash: unable to upsert taxtable_entries, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(taxtableEntryPrimaryKeyColumns))
			copy(conflict, taxtableEntryPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"taxtable_entries\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(taxtableEntryType, taxtableEntryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(taxtableEntryType, taxtableEntryMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "gnucash: unable to upsert taxtable_entries")
	}

	if !cached {
		taxtableEntryUpsertCacheMut.Lock()
		taxtableEntryUpsertCache[key] = cache
		taxtableEntryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TaxtableEntry record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TaxtableEntry) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("gnucash: no TaxtableEntry provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), taxtableEntryPrimaryKeyMapping)
	sql := "DELETE FROM \"taxtable_entries\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to delete from taxtable_entries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by delete for taxtable_entries")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q taxtableEntryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("gnucash: no taxtableEntryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to delete all from taxtable_entries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by deleteall for taxtable_entries")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TaxtableEntrySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(taxtableEntryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taxtableEntryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"taxtable_entries\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, taxtableEntryPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to delete all from taxtableEntry slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by deleteall for taxtable_entries")
	}

	if len(taxtableEntryAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TaxtableEntry) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTaxtableEntry(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TaxtableEntrySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TaxtableEntrySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taxtableEntryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"taxtable_entries\".* FROM \"taxtable_entries\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, taxtableEntryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "gnucash: unable to reload all in TaxtableEntrySlice")
	}

	*o = slice

	return nil
}

// TaxtableEntryExists checks if the TaxtableEntry row exists.
func TaxtableEntryExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"taxtable_entries\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "gnucash: unable to check if taxtable_entries exists")
	}

	return exists, nil
}

// Exists checks if the TaxtableEntry row exists.
func (o *TaxtableEntry) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TaxtableEntryExists(ctx, exec, o.ID)
}
