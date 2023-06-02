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

// Budget is an object representing the database table.
type Budget struct {
	GUID        string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	Name        string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Description null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	NumPeriods  int64       `boil:"num_periods" json:"num_periods" toml:"num_periods" yaml:"num_periods"`

	R *budgetR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L budgetL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BudgetColumns = struct {
	GUID        string
	Name        string
	Description string
	NumPeriods  string
}{
	GUID:        "guid",
	Name:        "name",
	Description: "description",
	NumPeriods:  "num_periods",
}

var BudgetTableColumns = struct {
	GUID        string
	Name        string
	Description string
	NumPeriods  string
}{
	GUID:        "budgets.guid",
	Name:        "budgets.name",
	Description: "budgets.description",
	NumPeriods:  "budgets.num_periods",
}

// Generated where

var BudgetWhere = struct {
	GUID        whereHelperstring
	Name        whereHelperstring
	Description whereHelpernull_String
	NumPeriods  whereHelperint64
}{
	GUID:        whereHelperstring{field: "\"budgets\".\"guid\""},
	Name:        whereHelperstring{field: "\"budgets\".\"name\""},
	Description: whereHelpernull_String{field: "\"budgets\".\"description\""},
	NumPeriods:  whereHelperint64{field: "\"budgets\".\"num_periods\""},
}

// BudgetRels is where relationship names are stored.
var BudgetRels = struct {
}{}

// budgetR is where relationships are stored.
type budgetR struct {
}

// NewStruct creates a new relationship struct
func (*budgetR) NewStruct() *budgetR {
	return &budgetR{}
}

// budgetL is where Load methods for each relationship are stored.
type budgetL struct{}

var (
	budgetAllColumns            = []string{"guid", "name", "description", "num_periods"}
	budgetColumnsWithoutDefault = []string{"guid", "name", "num_periods"}
	budgetColumnsWithDefault    = []string{"description"}
	budgetPrimaryKeyColumns     = []string{"guid"}
	budgetGeneratedColumns      = []string{}
)

type (
	// BudgetSlice is an alias for a slice of pointers to Budget.
	// This should almost always be used instead of []Budget.
	BudgetSlice []*Budget
	// BudgetHook is the signature for custom Budget hook methods
	BudgetHook func(context.Context, boil.ContextExecutor, *Budget) error

	budgetQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	budgetType                 = reflect.TypeOf(&Budget{})
	budgetMapping              = queries.MakeStructMapping(budgetType)
	budgetPrimaryKeyMapping, _ = queries.BindMapping(budgetType, budgetMapping, budgetPrimaryKeyColumns)
	budgetInsertCacheMut       sync.RWMutex
	budgetInsertCache          = make(map[string]insertCache)
	budgetUpdateCacheMut       sync.RWMutex
	budgetUpdateCache          = make(map[string]updateCache)
	budgetUpsertCacheMut       sync.RWMutex
	budgetUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var budgetAfterSelectHooks []BudgetHook

var budgetBeforeInsertHooks []BudgetHook
var budgetAfterInsertHooks []BudgetHook

var budgetBeforeUpdateHooks []BudgetHook
var budgetAfterUpdateHooks []BudgetHook

var budgetBeforeDeleteHooks []BudgetHook
var budgetAfterDeleteHooks []BudgetHook

var budgetBeforeUpsertHooks []BudgetHook
var budgetAfterUpsertHooks []BudgetHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Budget) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range budgetAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Budget) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range budgetBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Budget) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range budgetAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Budget) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range budgetBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Budget) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range budgetAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Budget) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range budgetBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Budget) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range budgetAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Budget) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range budgetBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Budget) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range budgetAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBudgetHook registers your hook function for all future operations.
func AddBudgetHook(hookPoint boil.HookPoint, budgetHook BudgetHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		budgetAfterSelectHooks = append(budgetAfterSelectHooks, budgetHook)
	case boil.BeforeInsertHook:
		budgetBeforeInsertHooks = append(budgetBeforeInsertHooks, budgetHook)
	case boil.AfterInsertHook:
		budgetAfterInsertHooks = append(budgetAfterInsertHooks, budgetHook)
	case boil.BeforeUpdateHook:
		budgetBeforeUpdateHooks = append(budgetBeforeUpdateHooks, budgetHook)
	case boil.AfterUpdateHook:
		budgetAfterUpdateHooks = append(budgetAfterUpdateHooks, budgetHook)
	case boil.BeforeDeleteHook:
		budgetBeforeDeleteHooks = append(budgetBeforeDeleteHooks, budgetHook)
	case boil.AfterDeleteHook:
		budgetAfterDeleteHooks = append(budgetAfterDeleteHooks, budgetHook)
	case boil.BeforeUpsertHook:
		budgetBeforeUpsertHooks = append(budgetBeforeUpsertHooks, budgetHook)
	case boil.AfterUpsertHook:
		budgetAfterUpsertHooks = append(budgetAfterUpsertHooks, budgetHook)
	}
}

// One returns a single budget record from the query.
func (q budgetQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Budget, error) {
	o := &Budget{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "gnucash: failed to execute a one query for budgets")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Budget records from the query.
func (q budgetQuery) All(ctx context.Context, exec boil.ContextExecutor) (BudgetSlice, error) {
	var o []*Budget

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "gnucash: failed to assign all query results to Budget slice")
	}

	if len(budgetAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Budget records in the query.
func (q budgetQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to count budgets rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q budgetQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "gnucash: failed to check if budgets exists")
	}

	return count > 0, nil
}

// Budgets retrieves all the records using an executor.
func Budgets(mods ...qm.QueryMod) budgetQuery {
	mods = append(mods, qm.From("\"budgets\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"budgets\".*"})
	}

	return budgetQuery{q}
}

// FindBudget retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBudget(ctx context.Context, exec boil.ContextExecutor, gUID string, selectCols ...string) (*Budget, error) {
	budgetObj := &Budget{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"budgets\" where \"guid\"=?", sel,
	)

	q := queries.Raw(query, gUID)

	err := q.Bind(ctx, exec, budgetObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "gnucash: unable to select from budgets")
	}

	if err = budgetObj.doAfterSelectHooks(ctx, exec); err != nil {
		return budgetObj, err
	}

	return budgetObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Budget) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("gnucash: no budgets provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(budgetColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	budgetInsertCacheMut.RLock()
	cache, cached := budgetInsertCache[key]
	budgetInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			budgetAllColumns,
			budgetColumnsWithDefault,
			budgetColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(budgetType, budgetMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(budgetType, budgetMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"budgets\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"budgets\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "gnucash: unable to insert into budgets")
	}

	if !cached {
		budgetInsertCacheMut.Lock()
		budgetInsertCache[key] = cache
		budgetInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Budget.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Budget) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	budgetUpdateCacheMut.RLock()
	cache, cached := budgetUpdateCache[key]
	budgetUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			budgetAllColumns,
			budgetPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("gnucash: unable to update budgets, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"budgets\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, budgetPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(budgetType, budgetMapping, append(wl, budgetPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "gnucash: unable to update budgets row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by update for budgets")
	}

	if !cached {
		budgetUpdateCacheMut.Lock()
		budgetUpdateCache[key] = cache
		budgetUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q budgetQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to update all for budgets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to retrieve rows affected for budgets")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BudgetSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), budgetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"budgets\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, budgetPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to update all in budget slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to retrieve rows affected all in update all budget")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Budget) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("gnucash: no budgets provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(budgetColumnsWithDefault, o)

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

	budgetUpsertCacheMut.RLock()
	cache, cached := budgetUpsertCache[key]
	budgetUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			budgetAllColumns,
			budgetColumnsWithDefault,
			budgetColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			budgetAllColumns,
			budgetPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("gnucash: unable to upsert budgets, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(budgetPrimaryKeyColumns))
			copy(conflict, budgetPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"budgets\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(budgetType, budgetMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(budgetType, budgetMapping, ret)
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
		return errors.Wrap(err, "gnucash: unable to upsert budgets")
	}

	if !cached {
		budgetUpsertCacheMut.Lock()
		budgetUpsertCache[key] = cache
		budgetUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Budget record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Budget) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("gnucash: no Budget provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), budgetPrimaryKeyMapping)
	sql := "DELETE FROM \"budgets\" WHERE \"guid\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to delete from budgets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by delete for budgets")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q budgetQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("gnucash: no budgetQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to delete all from budgets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by deleteall for budgets")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BudgetSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(budgetBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), budgetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"budgets\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, budgetPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to delete all from budget slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by deleteall for budgets")
	}

	if len(budgetAfterDeleteHooks) != 0 {
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
func (o *Budget) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBudget(ctx, exec, o.GUID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BudgetSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BudgetSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), budgetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"budgets\".* FROM \"budgets\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, budgetPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "gnucash: unable to reload all in BudgetSlice")
	}

	*o = slice

	return nil
}

// BudgetExists checks if the Budget row exists.
func BudgetExists(ctx context.Context, exec boil.ContextExecutor, gUID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"budgets\" where \"guid\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, gUID)
	}
	row := exec.QueryRowContext(ctx, sql, gUID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "gnucash: unable to check if budgets exists")
	}

	return exists, nil
}

// Exists checks if the Budget row exists.
func (o *Budget) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return BudgetExists(ctx, exec, o.GUID)
}
