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

// Employee is an object representing the database table.
type Employee struct {
	GUID         string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	Username     string      `boil:"username" json:"username" toml:"username" yaml:"username"`
	ID           string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Language     string      `boil:"language" json:"language" toml:"language" yaml:"language"`
	ACL          string      `boil:"acl" json:"acl" toml:"acl" yaml:"acl"`
	Active       int64       `boil:"active" json:"active" toml:"active" yaml:"active"`
	Currency     string      `boil:"currency" json:"currency" toml:"currency" yaml:"currency"`
	CcardGUID    null.String `boil:"ccard_guid" json:"ccard_guid,omitempty" toml:"ccard_guid" yaml:"ccard_guid,omitempty"`
	WorkdayNum   int64       `boil:"workday_num" json:"workday_num" toml:"workday_num" yaml:"workday_num"`
	WorkdayDenom int64       `boil:"workday_denom" json:"workday_denom" toml:"workday_denom" yaml:"workday_denom"`
	RateNum      int64       `boil:"rate_num" json:"rate_num" toml:"rate_num" yaml:"rate_num"`
	RateDenom    int64       `boil:"rate_denom" json:"rate_denom" toml:"rate_denom" yaml:"rate_denom"`
	AddrName     null.String `boil:"addr_name" json:"addr_name,omitempty" toml:"addr_name" yaml:"addr_name,omitempty"`
	AddrAddr1    null.String `boil:"addr_addr1" json:"addr_addr1,omitempty" toml:"addr_addr1" yaml:"addr_addr1,omitempty"`
	AddrAddr2    null.String `boil:"addr_addr2" json:"addr_addr2,omitempty" toml:"addr_addr2" yaml:"addr_addr2,omitempty"`
	AddrAddr3    null.String `boil:"addr_addr3" json:"addr_addr3,omitempty" toml:"addr_addr3" yaml:"addr_addr3,omitempty"`
	AddrAddr4    null.String `boil:"addr_addr4" json:"addr_addr4,omitempty" toml:"addr_addr4" yaml:"addr_addr4,omitempty"`
	AddrPhone    null.String `boil:"addr_phone" json:"addr_phone,omitempty" toml:"addr_phone" yaml:"addr_phone,omitempty"`
	AddrFax      null.String `boil:"addr_fax" json:"addr_fax,omitempty" toml:"addr_fax" yaml:"addr_fax,omitempty"`
	AddrEmail    null.String `boil:"addr_email" json:"addr_email,omitempty" toml:"addr_email" yaml:"addr_email,omitempty"`

	R *employeeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L employeeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EmployeeColumns = struct {
	GUID         string
	Username     string
	ID           string
	Language     string
	ACL          string
	Active       string
	Currency     string
	CcardGUID    string
	WorkdayNum   string
	WorkdayDenom string
	RateNum      string
	RateDenom    string
	AddrName     string
	AddrAddr1    string
	AddrAddr2    string
	AddrAddr3    string
	AddrAddr4    string
	AddrPhone    string
	AddrFax      string
	AddrEmail    string
}{
	GUID:         "guid",
	Username:     "username",
	ID:           "id",
	Language:     "language",
	ACL:          "acl",
	Active:       "active",
	Currency:     "currency",
	CcardGUID:    "ccard_guid",
	WorkdayNum:   "workday_num",
	WorkdayDenom: "workday_denom",
	RateNum:      "rate_num",
	RateDenom:    "rate_denom",
	AddrName:     "addr_name",
	AddrAddr1:    "addr_addr1",
	AddrAddr2:    "addr_addr2",
	AddrAddr3:    "addr_addr3",
	AddrAddr4:    "addr_addr4",
	AddrPhone:    "addr_phone",
	AddrFax:      "addr_fax",
	AddrEmail:    "addr_email",
}

var EmployeeTableColumns = struct {
	GUID         string
	Username     string
	ID           string
	Language     string
	ACL          string
	Active       string
	Currency     string
	CcardGUID    string
	WorkdayNum   string
	WorkdayDenom string
	RateNum      string
	RateDenom    string
	AddrName     string
	AddrAddr1    string
	AddrAddr2    string
	AddrAddr3    string
	AddrAddr4    string
	AddrPhone    string
	AddrFax      string
	AddrEmail    string
}{
	GUID:         "employees.guid",
	Username:     "employees.username",
	ID:           "employees.id",
	Language:     "employees.language",
	ACL:          "employees.acl",
	Active:       "employees.active",
	Currency:     "employees.currency",
	CcardGUID:    "employees.ccard_guid",
	WorkdayNum:   "employees.workday_num",
	WorkdayDenom: "employees.workday_denom",
	RateNum:      "employees.rate_num",
	RateDenom:    "employees.rate_denom",
	AddrName:     "employees.addr_name",
	AddrAddr1:    "employees.addr_addr1",
	AddrAddr2:    "employees.addr_addr2",
	AddrAddr3:    "employees.addr_addr3",
	AddrAddr4:    "employees.addr_addr4",
	AddrPhone:    "employees.addr_phone",
	AddrFax:      "employees.addr_fax",
	AddrEmail:    "employees.addr_email",
}

// Generated where

var EmployeeWhere = struct {
	GUID         whereHelperstring
	Username     whereHelperstring
	ID           whereHelperstring
	Language     whereHelperstring
	ACL          whereHelperstring
	Active       whereHelperint64
	Currency     whereHelperstring
	CcardGUID    whereHelpernull_String
	WorkdayNum   whereHelperint64
	WorkdayDenom whereHelperint64
	RateNum      whereHelperint64
	RateDenom    whereHelperint64
	AddrName     whereHelpernull_String
	AddrAddr1    whereHelpernull_String
	AddrAddr2    whereHelpernull_String
	AddrAddr3    whereHelpernull_String
	AddrAddr4    whereHelpernull_String
	AddrPhone    whereHelpernull_String
	AddrFax      whereHelpernull_String
	AddrEmail    whereHelpernull_String
}{
	GUID:         whereHelperstring{field: "\"employees\".\"guid\""},
	Username:     whereHelperstring{field: "\"employees\".\"username\""},
	ID:           whereHelperstring{field: "\"employees\".\"id\""},
	Language:     whereHelperstring{field: "\"employees\".\"language\""},
	ACL:          whereHelperstring{field: "\"employees\".\"acl\""},
	Active:       whereHelperint64{field: "\"employees\".\"active\""},
	Currency:     whereHelperstring{field: "\"employees\".\"currency\""},
	CcardGUID:    whereHelpernull_String{field: "\"employees\".\"ccard_guid\""},
	WorkdayNum:   whereHelperint64{field: "\"employees\".\"workday_num\""},
	WorkdayDenom: whereHelperint64{field: "\"employees\".\"workday_denom\""},
	RateNum:      whereHelperint64{field: "\"employees\".\"rate_num\""},
	RateDenom:    whereHelperint64{field: "\"employees\".\"rate_denom\""},
	AddrName:     whereHelpernull_String{field: "\"employees\".\"addr_name\""},
	AddrAddr1:    whereHelpernull_String{field: "\"employees\".\"addr_addr1\""},
	AddrAddr2:    whereHelpernull_String{field: "\"employees\".\"addr_addr2\""},
	AddrAddr3:    whereHelpernull_String{field: "\"employees\".\"addr_addr3\""},
	AddrAddr4:    whereHelpernull_String{field: "\"employees\".\"addr_addr4\""},
	AddrPhone:    whereHelpernull_String{field: "\"employees\".\"addr_phone\""},
	AddrFax:      whereHelpernull_String{field: "\"employees\".\"addr_fax\""},
	AddrEmail:    whereHelpernull_String{field: "\"employees\".\"addr_email\""},
}

// EmployeeRels is where relationship names are stored.
var EmployeeRels = struct {
}{}

// employeeR is where relationships are stored.
type employeeR struct {
}

// NewStruct creates a new relationship struct
func (*employeeR) NewStruct() *employeeR {
	return &employeeR{}
}

// employeeL is where Load methods for each relationship are stored.
type employeeL struct{}

var (
	employeeAllColumns            = []string{"guid", "username", "id", "language", "acl", "active", "currency", "ccard_guid", "workday_num", "workday_denom", "rate_num", "rate_denom", "addr_name", "addr_addr1", "addr_addr2", "addr_addr3", "addr_addr4", "addr_phone", "addr_fax", "addr_email"}
	employeeColumnsWithoutDefault = []string{"guid", "username", "id", "language", "acl", "active", "currency", "workday_num", "workday_denom", "rate_num", "rate_denom"}
	employeeColumnsWithDefault    = []string{"ccard_guid", "addr_name", "addr_addr1", "addr_addr2", "addr_addr3", "addr_addr4", "addr_phone", "addr_fax", "addr_email"}
	employeePrimaryKeyColumns     = []string{"guid"}
	employeeGeneratedColumns      = []string{}
)

type (
	// EmployeeSlice is an alias for a slice of pointers to Employee.
	// This should almost always be used instead of []Employee.
	EmployeeSlice []*Employee
	// EmployeeHook is the signature for custom Employee hook methods
	EmployeeHook func(context.Context, boil.ContextExecutor, *Employee) error

	employeeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	employeeType                 = reflect.TypeOf(&Employee{})
	employeeMapping              = queries.MakeStructMapping(employeeType)
	employeePrimaryKeyMapping, _ = queries.BindMapping(employeeType, employeeMapping, employeePrimaryKeyColumns)
	employeeInsertCacheMut       sync.RWMutex
	employeeInsertCache          = make(map[string]insertCache)
	employeeUpdateCacheMut       sync.RWMutex
	employeeUpdateCache          = make(map[string]updateCache)
	employeeUpsertCacheMut       sync.RWMutex
	employeeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var employeeAfterSelectHooks []EmployeeHook

var employeeBeforeInsertHooks []EmployeeHook
var employeeAfterInsertHooks []EmployeeHook

var employeeBeforeUpdateHooks []EmployeeHook
var employeeAfterUpdateHooks []EmployeeHook

var employeeBeforeDeleteHooks []EmployeeHook
var employeeAfterDeleteHooks []EmployeeHook

var employeeBeforeUpsertHooks []EmployeeHook
var employeeAfterUpsertHooks []EmployeeHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Employee) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range employeeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Employee) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range employeeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Employee) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range employeeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Employee) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range employeeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Employee) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range employeeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Employee) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range employeeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Employee) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range employeeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Employee) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range employeeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Employee) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range employeeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEmployeeHook registers your hook function for all future operations.
func AddEmployeeHook(hookPoint boil.HookPoint, employeeHook EmployeeHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		employeeAfterSelectHooks = append(employeeAfterSelectHooks, employeeHook)
	case boil.BeforeInsertHook:
		employeeBeforeInsertHooks = append(employeeBeforeInsertHooks, employeeHook)
	case boil.AfterInsertHook:
		employeeAfterInsertHooks = append(employeeAfterInsertHooks, employeeHook)
	case boil.BeforeUpdateHook:
		employeeBeforeUpdateHooks = append(employeeBeforeUpdateHooks, employeeHook)
	case boil.AfterUpdateHook:
		employeeAfterUpdateHooks = append(employeeAfterUpdateHooks, employeeHook)
	case boil.BeforeDeleteHook:
		employeeBeforeDeleteHooks = append(employeeBeforeDeleteHooks, employeeHook)
	case boil.AfterDeleteHook:
		employeeAfterDeleteHooks = append(employeeAfterDeleteHooks, employeeHook)
	case boil.BeforeUpsertHook:
		employeeBeforeUpsertHooks = append(employeeBeforeUpsertHooks, employeeHook)
	case boil.AfterUpsertHook:
		employeeAfterUpsertHooks = append(employeeAfterUpsertHooks, employeeHook)
	}
}

// One returns a single employee record from the query.
func (q employeeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Employee, error) {
	o := &Employee{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "gnucash: failed to execute a one query for employees")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Employee records from the query.
func (q employeeQuery) All(ctx context.Context, exec boil.ContextExecutor) (EmployeeSlice, error) {
	var o []*Employee

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "gnucash: failed to assign all query results to Employee slice")
	}

	if len(employeeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Employee records in the query.
func (q employeeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to count employees rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q employeeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "gnucash: failed to check if employees exists")
	}

	return count > 0, nil
}

// Employees retrieves all the records using an executor.
func Employees(mods ...qm.QueryMod) employeeQuery {
	mods = append(mods, qm.From("\"employees\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"employees\".*"})
	}

	return employeeQuery{q}
}

// FindEmployee retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEmployee(ctx context.Context, exec boil.ContextExecutor, gUID string, selectCols ...string) (*Employee, error) {
	employeeObj := &Employee{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"employees\" where \"guid\"=?", sel,
	)

	q := queries.Raw(query, gUID)

	err := q.Bind(ctx, exec, employeeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "gnucash: unable to select from employees")
	}

	if err = employeeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return employeeObj, err
	}

	return employeeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Employee) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("gnucash: no employees provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(employeeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	employeeInsertCacheMut.RLock()
	cache, cached := employeeInsertCache[key]
	employeeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			employeeAllColumns,
			employeeColumnsWithDefault,
			employeeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(employeeType, employeeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(employeeType, employeeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"employees\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"employees\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "gnucash: unable to insert into employees")
	}

	if !cached {
		employeeInsertCacheMut.Lock()
		employeeInsertCache[key] = cache
		employeeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Employee.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Employee) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	employeeUpdateCacheMut.RLock()
	cache, cached := employeeUpdateCache[key]
	employeeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			employeeAllColumns,
			employeePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("gnucash: unable to update employees, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"employees\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, employeePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(employeeType, employeeMapping, append(wl, employeePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "gnucash: unable to update employees row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by update for employees")
	}

	if !cached {
		employeeUpdateCacheMut.Lock()
		employeeUpdateCache[key] = cache
		employeeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q employeeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to update all for employees")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to retrieve rows affected for employees")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EmployeeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), employeePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"employees\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, employeePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to update all in employee slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to retrieve rows affected all in update all employee")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Employee) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("gnucash: no employees provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(employeeColumnsWithDefault, o)

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

	employeeUpsertCacheMut.RLock()
	cache, cached := employeeUpsertCache[key]
	employeeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			employeeAllColumns,
			employeeColumnsWithDefault,
			employeeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			employeeAllColumns,
			employeePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("gnucash: unable to upsert employees, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(employeePrimaryKeyColumns))
			copy(conflict, employeePrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"employees\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(employeeType, employeeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(employeeType, employeeMapping, ret)
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
		return errors.Wrap(err, "gnucash: unable to upsert employees")
	}

	if !cached {
		employeeUpsertCacheMut.Lock()
		employeeUpsertCache[key] = cache
		employeeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Employee record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Employee) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("gnucash: no Employee provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), employeePrimaryKeyMapping)
	sql := "DELETE FROM \"employees\" WHERE \"guid\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to delete from employees")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by delete for employees")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q employeeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("gnucash: no employeeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to delete all from employees")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by deleteall for employees")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EmployeeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(employeeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), employeePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"employees\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, employeePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to delete all from employee slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by deleteall for employees")
	}

	if len(employeeAfterDeleteHooks) != 0 {
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
func (o *Employee) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEmployee(ctx, exec, o.GUID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EmployeeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EmployeeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), employeePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"employees\".* FROM \"employees\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, employeePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "gnucash: unable to reload all in EmployeeSlice")
	}

	*o = slice

	return nil
}

// EmployeeExists checks if the Employee row exists.
func EmployeeExists(ctx context.Context, exec boil.ContextExecutor, gUID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"employees\" where \"guid\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, gUID)
	}
	row := exec.QueryRowContext(ctx, sql, gUID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "gnucash: unable to check if employees exists")
	}

	return exists, nil
}

// Exists checks if the Employee row exists.
func (o *Employee) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return EmployeeExists(ctx, exec, o.GUID)
}
