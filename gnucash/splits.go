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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Split is an object representing the database table.
type Split struct {
	GUID           string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	TXGUID         string      `boil:"tx_guid" json:"tx_guid" toml:"tx_guid" yaml:"tx_guid"`
	AccountGUID    string      `boil:"account_guid" json:"account_guid" toml:"account_guid" yaml:"account_guid"`
	Memo           string      `boil:"memo" json:"memo" toml:"memo" yaml:"memo"`
	Action         string      `boil:"action" json:"action" toml:"action" yaml:"action"`
	ReconcileState string      `boil:"reconcile_state" json:"reconcile_state" toml:"reconcile_state" yaml:"reconcile_state"`
	ReconcileDate  null.String `boil:"reconcile_date" json:"reconcile_date,omitempty" toml:"reconcile_date" yaml:"reconcile_date,omitempty"`
	ValueNum       int64       `boil:"value_num" json:"value_num" toml:"value_num" yaml:"value_num"`
	ValueDenom     int64       `boil:"value_denom" json:"value_denom" toml:"value_denom" yaml:"value_denom"`
	QuantityNum    int64       `boil:"quantity_num" json:"quantity_num" toml:"quantity_num" yaml:"quantity_num"`
	QuantityDenom  int64       `boil:"quantity_denom" json:"quantity_denom" toml:"quantity_denom" yaml:"quantity_denom"`
	LotGUID        null.String `boil:"lot_guid" json:"lot_guid,omitempty" toml:"lot_guid" yaml:"lot_guid,omitempty"`

	R *splitR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L splitL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SplitColumns = struct {
	GUID           string
	TXGUID         string
	AccountGUID    string
	Memo           string
	Action         string
	ReconcileState string
	ReconcileDate  string
	ValueNum       string
	ValueDenom     string
	QuantityNum    string
	QuantityDenom  string
	LotGUID        string
}{
	GUID:           "guid",
	TXGUID:         "tx_guid",
	AccountGUID:    "account_guid",
	Memo:           "memo",
	Action:         "action",
	ReconcileState: "reconcile_state",
	ReconcileDate:  "reconcile_date",
	ValueNum:       "value_num",
	ValueDenom:     "value_denom",
	QuantityNum:    "quantity_num",
	QuantityDenom:  "quantity_denom",
	LotGUID:        "lot_guid",
}

var SplitTableColumns = struct {
	GUID           string
	TXGUID         string
	AccountGUID    string
	Memo           string
	Action         string
	ReconcileState string
	ReconcileDate  string
	ValueNum       string
	ValueDenom     string
	QuantityNum    string
	QuantityDenom  string
	LotGUID        string
}{
	GUID:           "splits.guid",
	TXGUID:         "splits.tx_guid",
	AccountGUID:    "splits.account_guid",
	Memo:           "splits.memo",
	Action:         "splits.action",
	ReconcileState: "splits.reconcile_state",
	ReconcileDate:  "splits.reconcile_date",
	ValueNum:       "splits.value_num",
	ValueDenom:     "splits.value_denom",
	QuantityNum:    "splits.quantity_num",
	QuantityDenom:  "splits.quantity_denom",
	LotGUID:        "splits.lot_guid",
}

// Generated where

var SplitWhere = struct {
	GUID           whereHelperstring
	TXGUID         whereHelperstring
	AccountGUID    whereHelperstring
	Memo           whereHelperstring
	Action         whereHelperstring
	ReconcileState whereHelperstring
	ReconcileDate  whereHelpernull_String
	ValueNum       whereHelperint64
	ValueDenom     whereHelperint64
	QuantityNum    whereHelperint64
	QuantityDenom  whereHelperint64
	LotGUID        whereHelpernull_String
}{
	GUID:           whereHelperstring{field: "\"splits\".\"guid\""},
	TXGUID:         whereHelperstring{field: "\"splits\".\"tx_guid\""},
	AccountGUID:    whereHelperstring{field: "\"splits\".\"account_guid\""},
	Memo:           whereHelperstring{field: "\"splits\".\"memo\""},
	Action:         whereHelperstring{field: "\"splits\".\"action\""},
	ReconcileState: whereHelperstring{field: "\"splits\".\"reconcile_state\""},
	ReconcileDate:  whereHelpernull_String{field: "\"splits\".\"reconcile_date\""},
	ValueNum:       whereHelperint64{field: "\"splits\".\"value_num\""},
	ValueDenom:     whereHelperint64{field: "\"splits\".\"value_denom\""},
	QuantityNum:    whereHelperint64{field: "\"splits\".\"quantity_num\""},
	QuantityDenom:  whereHelperint64{field: "\"splits\".\"quantity_denom\""},
	LotGUID:        whereHelpernull_String{field: "\"splits\".\"lot_guid\""},
}

// SplitRels is where relationship names are stored.
var SplitRels = struct {
}{}

// splitR is where relationships are stored.
type splitR struct {
}

// NewStruct creates a new relationship struct
func (*splitR) NewStruct() *splitR {
	return &splitR{}
}

// splitL is where Load methods for each relationship are stored.
type splitL struct{}

var (
	splitAllColumns            = []string{"guid", "tx_guid", "account_guid", "memo", "action", "reconcile_state", "reconcile_date", "value_num", "value_denom", "quantity_num", "quantity_denom", "lot_guid"}
	splitColumnsWithoutDefault = []string{"guid", "tx_guid", "account_guid", "memo", "action", "reconcile_state", "value_num", "value_denom", "quantity_num", "quantity_denom"}
	splitColumnsWithDefault    = []string{"reconcile_date", "lot_guid"}
	splitPrimaryKeyColumns     = []string{"guid"}
	splitGeneratedColumns      = []string{}
)

type (
	// SplitSlice is an alias for a slice of pointers to Split.
	// This should almost always be used instead of []Split.
	SplitSlice []*Split
	// SplitHook is the signature for custom Split hook methods
	SplitHook func(context.Context, boil.ContextExecutor, *Split) error

	splitQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	splitType                 = reflect.TypeOf(&Split{})
	splitMapping              = queries.MakeStructMapping(splitType)
	splitPrimaryKeyMapping, _ = queries.BindMapping(splitType, splitMapping, splitPrimaryKeyColumns)
	splitInsertCacheMut       sync.RWMutex
	splitInsertCache          = make(map[string]insertCache)
	splitUpdateCacheMut       sync.RWMutex
	splitUpdateCache          = make(map[string]updateCache)
	splitUpsertCacheMut       sync.RWMutex
	splitUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var splitAfterSelectHooks []SplitHook

var splitBeforeInsertHooks []SplitHook
var splitAfterInsertHooks []SplitHook

var splitBeforeUpdateHooks []SplitHook
var splitAfterUpdateHooks []SplitHook

var splitBeforeDeleteHooks []SplitHook
var splitAfterDeleteHooks []SplitHook

var splitBeforeUpsertHooks []SplitHook
var splitAfterUpsertHooks []SplitHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Split) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range splitAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Split) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range splitBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Split) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range splitAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Split) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range splitBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Split) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range splitAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Split) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range splitBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Split) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range splitAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Split) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range splitBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Split) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range splitAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSplitHook registers your hook function for all future operations.
func AddSplitHook(hookPoint boil.HookPoint, splitHook SplitHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		splitAfterSelectHooks = append(splitAfterSelectHooks, splitHook)
	case boil.BeforeInsertHook:
		splitBeforeInsertHooks = append(splitBeforeInsertHooks, splitHook)
	case boil.AfterInsertHook:
		splitAfterInsertHooks = append(splitAfterInsertHooks, splitHook)
	case boil.BeforeUpdateHook:
		splitBeforeUpdateHooks = append(splitBeforeUpdateHooks, splitHook)
	case boil.AfterUpdateHook:
		splitAfterUpdateHooks = append(splitAfterUpdateHooks, splitHook)
	case boil.BeforeDeleteHook:
		splitBeforeDeleteHooks = append(splitBeforeDeleteHooks, splitHook)
	case boil.AfterDeleteHook:
		splitAfterDeleteHooks = append(splitAfterDeleteHooks, splitHook)
	case boil.BeforeUpsertHook:
		splitBeforeUpsertHooks = append(splitBeforeUpsertHooks, splitHook)
	case boil.AfterUpsertHook:
		splitAfterUpsertHooks = append(splitAfterUpsertHooks, splitHook)
	}
}

// One returns a single split record from the query.
func (q splitQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Split, error) {
	o := &Split{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "gnucash: failed to execute a one query for splits")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Split records from the query.
func (q splitQuery) All(ctx context.Context, exec boil.ContextExecutor) (SplitSlice, error) {
	var o []*Split

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "gnucash: failed to assign all query results to Split slice")
	}

	if len(splitAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Split records in the query.
func (q splitQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to count splits rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q splitQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "gnucash: failed to check if splits exists")
	}

	return count > 0, nil
}

// Splits retrieves all the records using an executor.
func Splits(mods ...qm.QueryMod) splitQuery {
	mods = append(mods, qm.From("\"splits\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"splits\".*"})
	}

	return splitQuery{q}
}

// FindSplit retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSplit(ctx context.Context, exec boil.ContextExecutor, gUID string, selectCols ...string) (*Split, error) {
	splitObj := &Split{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"splits\" where \"guid\"=?", sel,
	)

	q := queries.Raw(query, gUID)

	err := q.Bind(ctx, exec, splitObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "gnucash: unable to select from splits")
	}

	if err = splitObj.doAfterSelectHooks(ctx, exec); err != nil {
		return splitObj, err
	}

	return splitObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Split) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("gnucash: no splits provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(splitColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	splitInsertCacheMut.RLock()
	cache, cached := splitInsertCache[key]
	splitInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			splitAllColumns,
			splitColumnsWithDefault,
			splitColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(splitType, splitMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(splitType, splitMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"splits\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"splits\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "gnucash: unable to insert into splits")
	}

	if !cached {
		splitInsertCacheMut.Lock()
		splitInsertCache[key] = cache
		splitInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Split.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Split) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	splitUpdateCacheMut.RLock()
	cache, cached := splitUpdateCache[key]
	splitUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			splitAllColumns,
			splitPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("gnucash: unable to update splits, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"splits\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, splitPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(splitType, splitMapping, append(wl, splitPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "gnucash: unable to update splits row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by update for splits")
	}

	if !cached {
		splitUpdateCacheMut.Lock()
		splitUpdateCache[key] = cache
		splitUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q splitQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to update all for splits")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to retrieve rows affected for splits")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SplitSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), splitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"splits\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, splitPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to update all in split slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to retrieve rows affected all in update all split")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Split) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("gnucash: no splits provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(splitColumnsWithDefault, o)

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

	splitUpsertCacheMut.RLock()
	cache, cached := splitUpsertCache[key]
	splitUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			splitAllColumns,
			splitColumnsWithDefault,
			splitColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			splitAllColumns,
			splitPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("gnucash: unable to upsert splits, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(splitPrimaryKeyColumns))
			copy(conflict, splitPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"splits\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(splitType, splitMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(splitType, splitMapping, ret)
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
		return errors.Wrap(err, "gnucash: unable to upsert splits")
	}

	if !cached {
		splitUpsertCacheMut.Lock()
		splitUpsertCache[key] = cache
		splitUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Split record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Split) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("gnucash: no Split provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), splitPrimaryKeyMapping)
	sql := "DELETE FROM \"splits\" WHERE \"guid\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to delete from splits")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by delete for splits")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q splitQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("gnucash: no splitQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to delete all from splits")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by deleteall for splits")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SplitSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(splitBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), splitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"splits\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, splitPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to delete all from split slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by deleteall for splits")
	}

	if len(splitAfterDeleteHooks) != 0 {
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
func (o *Split) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSplit(ctx, exec, o.GUID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SplitSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SplitSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), splitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"splits\".* FROM \"splits\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, splitPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "gnucash: unable to reload all in SplitSlice")
	}

	*o = slice

	return nil
}

// SplitExists checks if the Split row exists.
func SplitExists(ctx context.Context, exec boil.ContextExecutor, gUID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"splits\" where \"guid\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, gUID)
	}
	row := exec.QueryRowContext(ctx, sql, gUID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "gnucash: unable to check if splits exists")
	}

	return exists, nil
}

// Exists checks if the Split row exists.
func (o *Split) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return SplitExists(ctx, exec, o.GUID)
}