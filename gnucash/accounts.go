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

// Account is an object representing the database table.
type Account struct {
	GUID          string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	Name          string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	AccountType   string      `boil:"account_type" json:"account_type" toml:"account_type" yaml:"account_type"`
	CommodityGUID null.String `boil:"commodity_guid" json:"commodity_guid,omitempty" toml:"commodity_guid" yaml:"commodity_guid,omitempty"`
	CommodityScu  int64       `boil:"commodity_scu" json:"commodity_scu" toml:"commodity_scu" yaml:"commodity_scu"`
	NonSTDScu     int64       `boil:"non_std_scu" json:"non_std_scu" toml:"non_std_scu" yaml:"non_std_scu"`
	ParentGUID    null.String `boil:"parent_guid" json:"parent_guid,omitempty" toml:"parent_guid" yaml:"parent_guid,omitempty"`
	Code          null.String `boil:"code" json:"code,omitempty" toml:"code" yaml:"code,omitempty"`
	Description   null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	Hidden        null.Int64  `boil:"hidden" json:"hidden,omitempty" toml:"hidden" yaml:"hidden,omitempty"`
	Placeholder   null.Int64  `boil:"placeholder" json:"placeholder,omitempty" toml:"placeholder" yaml:"placeholder,omitempty"`

	R *accountR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L accountL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AccountColumns = struct {
	GUID          string
	Name          string
	AccountType   string
	CommodityGUID string
	CommodityScu  string
	NonSTDScu     string
	ParentGUID    string
	Code          string
	Description   string
	Hidden        string
	Placeholder   string
}{
	GUID:          "guid",
	Name:          "name",
	AccountType:   "account_type",
	CommodityGUID: "commodity_guid",
	CommodityScu:  "commodity_scu",
	NonSTDScu:     "non_std_scu",
	ParentGUID:    "parent_guid",
	Code:          "code",
	Description:   "description",
	Hidden:        "hidden",
	Placeholder:   "placeholder",
}

var AccountTableColumns = struct {
	GUID          string
	Name          string
	AccountType   string
	CommodityGUID string
	CommodityScu  string
	NonSTDScu     string
	ParentGUID    string
	Code          string
	Description   string
	Hidden        string
	Placeholder   string
}{
	GUID:          "accounts.guid",
	Name:          "accounts.name",
	AccountType:   "accounts.account_type",
	CommodityGUID: "accounts.commodity_guid",
	CommodityScu:  "accounts.commodity_scu",
	NonSTDScu:     "accounts.non_std_scu",
	ParentGUID:    "accounts.parent_guid",
	Code:          "accounts.code",
	Description:   "accounts.description",
	Hidden:        "accounts.hidden",
	Placeholder:   "accounts.placeholder",
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
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_String) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_String) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
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
func (w whereHelpernull_Int64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Int64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var AccountWhere = struct {
	GUID          whereHelperstring
	Name          whereHelperstring
	AccountType   whereHelperstring
	CommodityGUID whereHelpernull_String
	CommodityScu  whereHelperint64
	NonSTDScu     whereHelperint64
	ParentGUID    whereHelpernull_String
	Code          whereHelpernull_String
	Description   whereHelpernull_String
	Hidden        whereHelpernull_Int64
	Placeholder   whereHelpernull_Int64
}{
	GUID:          whereHelperstring{field: "\"accounts\".\"guid\""},
	Name:          whereHelperstring{field: "\"accounts\".\"name\""},
	AccountType:   whereHelperstring{field: "\"accounts\".\"account_type\""},
	CommodityGUID: whereHelpernull_String{field: "\"accounts\".\"commodity_guid\""},
	CommodityScu:  whereHelperint64{field: "\"accounts\".\"commodity_scu\""},
	NonSTDScu:     whereHelperint64{field: "\"accounts\".\"non_std_scu\""},
	ParentGUID:    whereHelpernull_String{field: "\"accounts\".\"parent_guid\""},
	Code:          whereHelpernull_String{field: "\"accounts\".\"code\""},
	Description:   whereHelpernull_String{field: "\"accounts\".\"description\""},
	Hidden:        whereHelpernull_Int64{field: "\"accounts\".\"hidden\""},
	Placeholder:   whereHelpernull_Int64{field: "\"accounts\".\"placeholder\""},
}

// AccountRels is where relationship names are stored.
var AccountRels = struct {
}{}

// accountR is where relationships are stored.
type accountR struct {
}

// NewStruct creates a new relationship struct
func (*accountR) NewStruct() *accountR {
	return &accountR{}
}

// accountL is where Load methods for each relationship are stored.
type accountL struct{}

var (
	accountAllColumns            = []string{"guid", "name", "account_type", "commodity_guid", "commodity_scu", "non_std_scu", "parent_guid", "code", "description", "hidden", "placeholder"}
	accountColumnsWithoutDefault = []string{"guid", "name", "account_type", "commodity_scu", "non_std_scu"}
	accountColumnsWithDefault    = []string{"commodity_guid", "parent_guid", "code", "description", "hidden", "placeholder"}
	accountPrimaryKeyColumns     = []string{"guid"}
	accountGeneratedColumns      = []string{}
)

type (
	// AccountSlice is an alias for a slice of pointers to Account.
	// This should almost always be used instead of []Account.
	AccountSlice []*Account
	// AccountHook is the signature for custom Account hook methods
	AccountHook func(context.Context, boil.ContextExecutor, *Account) error

	accountQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	accountType                 = reflect.TypeOf(&Account{})
	accountMapping              = queries.MakeStructMapping(accountType)
	accountPrimaryKeyMapping, _ = queries.BindMapping(accountType, accountMapping, accountPrimaryKeyColumns)
	accountInsertCacheMut       sync.RWMutex
	accountInsertCache          = make(map[string]insertCache)
	accountUpdateCacheMut       sync.RWMutex
	accountUpdateCache          = make(map[string]updateCache)
	accountUpsertCacheMut       sync.RWMutex
	accountUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var accountAfterSelectHooks []AccountHook

var accountBeforeInsertHooks []AccountHook
var accountAfterInsertHooks []AccountHook

var accountBeforeUpdateHooks []AccountHook
var accountAfterUpdateHooks []AccountHook

var accountBeforeDeleteHooks []AccountHook
var accountAfterDeleteHooks []AccountHook

var accountBeforeUpsertHooks []AccountHook
var accountAfterUpsertHooks []AccountHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Account) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Account) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Account) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Account) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Account) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Account) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Account) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Account) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Account) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAccountHook registers your hook function for all future operations.
func AddAccountHook(hookPoint boil.HookPoint, accountHook AccountHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		accountAfterSelectHooks = append(accountAfterSelectHooks, accountHook)
	case boil.BeforeInsertHook:
		accountBeforeInsertHooks = append(accountBeforeInsertHooks, accountHook)
	case boil.AfterInsertHook:
		accountAfterInsertHooks = append(accountAfterInsertHooks, accountHook)
	case boil.BeforeUpdateHook:
		accountBeforeUpdateHooks = append(accountBeforeUpdateHooks, accountHook)
	case boil.AfterUpdateHook:
		accountAfterUpdateHooks = append(accountAfterUpdateHooks, accountHook)
	case boil.BeforeDeleteHook:
		accountBeforeDeleteHooks = append(accountBeforeDeleteHooks, accountHook)
	case boil.AfterDeleteHook:
		accountAfterDeleteHooks = append(accountAfterDeleteHooks, accountHook)
	case boil.BeforeUpsertHook:
		accountBeforeUpsertHooks = append(accountBeforeUpsertHooks, accountHook)
	case boil.AfterUpsertHook:
		accountAfterUpsertHooks = append(accountAfterUpsertHooks, accountHook)
	}
}

// One returns a single account record from the query.
func (q accountQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Account, error) {
	o := &Account{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "gnucash: failed to execute a one query for accounts")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Account records from the query.
func (q accountQuery) All(ctx context.Context, exec boil.ContextExecutor) (AccountSlice, error) {
	var o []*Account

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "gnucash: failed to assign all query results to Account slice")
	}

	if len(accountAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Account records in the query.
func (q accountQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to count accounts rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q accountQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "gnucash: failed to check if accounts exists")
	}

	return count > 0, nil
}

// Accounts retrieves all the records using an executor.
func Accounts(mods ...qm.QueryMod) accountQuery {
	mods = append(mods, qm.From("\"accounts\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"accounts\".*"})
	}

	return accountQuery{q}
}

// FindAccount retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAccount(ctx context.Context, exec boil.ContextExecutor, gUID string, selectCols ...string) (*Account, error) {
	accountObj := &Account{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"accounts\" where \"guid\"=?", sel,
	)

	q := queries.Raw(query, gUID)

	err := q.Bind(ctx, exec, accountObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "gnucash: unable to select from accounts")
	}

	if err = accountObj.doAfterSelectHooks(ctx, exec); err != nil {
		return accountObj, err
	}

	return accountObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Account) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("gnucash: no accounts provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(accountColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	accountInsertCacheMut.RLock()
	cache, cached := accountInsertCache[key]
	accountInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			accountAllColumns,
			accountColumnsWithDefault,
			accountColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(accountType, accountMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(accountType, accountMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"accounts\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"accounts\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "gnucash: unable to insert into accounts")
	}

	if !cached {
		accountInsertCacheMut.Lock()
		accountInsertCache[key] = cache
		accountInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Account.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Account) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	accountUpdateCacheMut.RLock()
	cache, cached := accountUpdateCache[key]
	accountUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			accountAllColumns,
			accountPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("gnucash: unable to update accounts, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"accounts\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, accountPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(accountType, accountMapping, append(wl, accountPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "gnucash: unable to update accounts row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by update for accounts")
	}

	if !cached {
		accountUpdateCacheMut.Lock()
		accountUpdateCache[key] = cache
		accountUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q accountQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to update all for accounts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to retrieve rows affected for accounts")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AccountSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"accounts\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, accountPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to update all in account slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to retrieve rows affected all in update all account")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Account) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("gnucash: no accounts provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(accountColumnsWithDefault, o)

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

	accountUpsertCacheMut.RLock()
	cache, cached := accountUpsertCache[key]
	accountUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			accountAllColumns,
			accountColumnsWithDefault,
			accountColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			accountAllColumns,
			accountPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("gnucash: unable to upsert accounts, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(accountPrimaryKeyColumns))
			copy(conflict, accountPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"accounts\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(accountType, accountMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(accountType, accountMapping, ret)
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
		return errors.Wrap(err, "gnucash: unable to upsert accounts")
	}

	if !cached {
		accountUpsertCacheMut.Lock()
		accountUpsertCache[key] = cache
		accountUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Account record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Account) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("gnucash: no Account provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), accountPrimaryKeyMapping)
	sql := "DELETE FROM \"accounts\" WHERE \"guid\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to delete from accounts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by delete for accounts")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q accountQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("gnucash: no accountQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to delete all from accounts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by deleteall for accounts")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AccountSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(accountBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"accounts\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, accountPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to delete all from account slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by deleteall for accounts")
	}

	if len(accountAfterDeleteHooks) != 0 {
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
func (o *Account) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAccount(ctx, exec, o.GUID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AccountSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AccountSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"accounts\".* FROM \"accounts\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, accountPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "gnucash: unable to reload all in AccountSlice")
	}

	*o = slice

	return nil
}

// AccountExists checks if the Account row exists.
func AccountExists(ctx context.Context, exec boil.ContextExecutor, gUID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"accounts\" where \"guid\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, gUID)
	}
	row := exec.QueryRowContext(ctx, sql, gUID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "gnucash: unable to check if accounts exists")
	}

	return exists, nil
}

// Exists checks if the Account row exists.
func (o *Account) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return AccountExists(ctx, exec, o.GUID)
}
