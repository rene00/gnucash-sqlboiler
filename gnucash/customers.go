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

// Customer is an object representing the database table.
type Customer struct {
	GUID          string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	Name          string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	ID            string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Notes         string      `boil:"notes" json:"notes" toml:"notes" yaml:"notes"`
	Active        int64       `boil:"active" json:"active" toml:"active" yaml:"active"`
	DiscountNum   int64       `boil:"discount_num" json:"discount_num" toml:"discount_num" yaml:"discount_num"`
	DiscountDenom int64       `boil:"discount_denom" json:"discount_denom" toml:"discount_denom" yaml:"discount_denom"`
	CreditNum     int64       `boil:"credit_num" json:"credit_num" toml:"credit_num" yaml:"credit_num"`
	CreditDenom   int64       `boil:"credit_denom" json:"credit_denom" toml:"credit_denom" yaml:"credit_denom"`
	Currency      string      `boil:"currency" json:"currency" toml:"currency" yaml:"currency"`
	TaxOverride   int64       `boil:"tax_override" json:"tax_override" toml:"tax_override" yaml:"tax_override"`
	AddrName      null.String `boil:"addr_name" json:"addr_name,omitempty" toml:"addr_name" yaml:"addr_name,omitempty"`
	AddrAddr1     null.String `boil:"addr_addr1" json:"addr_addr1,omitempty" toml:"addr_addr1" yaml:"addr_addr1,omitempty"`
	AddrAddr2     null.String `boil:"addr_addr2" json:"addr_addr2,omitempty" toml:"addr_addr2" yaml:"addr_addr2,omitempty"`
	AddrAddr3     null.String `boil:"addr_addr3" json:"addr_addr3,omitempty" toml:"addr_addr3" yaml:"addr_addr3,omitempty"`
	AddrAddr4     null.String `boil:"addr_addr4" json:"addr_addr4,omitempty" toml:"addr_addr4" yaml:"addr_addr4,omitempty"`
	AddrPhone     null.String `boil:"addr_phone" json:"addr_phone,omitempty" toml:"addr_phone" yaml:"addr_phone,omitempty"`
	AddrFax       null.String `boil:"addr_fax" json:"addr_fax,omitempty" toml:"addr_fax" yaml:"addr_fax,omitempty"`
	AddrEmail     null.String `boil:"addr_email" json:"addr_email,omitempty" toml:"addr_email" yaml:"addr_email,omitempty"`
	ShipaddrName  null.String `boil:"shipaddr_name" json:"shipaddr_name,omitempty" toml:"shipaddr_name" yaml:"shipaddr_name,omitempty"`
	ShipaddrAddr1 null.String `boil:"shipaddr_addr1" json:"shipaddr_addr1,omitempty" toml:"shipaddr_addr1" yaml:"shipaddr_addr1,omitempty"`
	ShipaddrAddr2 null.String `boil:"shipaddr_addr2" json:"shipaddr_addr2,omitempty" toml:"shipaddr_addr2" yaml:"shipaddr_addr2,omitempty"`
	ShipaddrAddr3 null.String `boil:"shipaddr_addr3" json:"shipaddr_addr3,omitempty" toml:"shipaddr_addr3" yaml:"shipaddr_addr3,omitempty"`
	ShipaddrAddr4 null.String `boil:"shipaddr_addr4" json:"shipaddr_addr4,omitempty" toml:"shipaddr_addr4" yaml:"shipaddr_addr4,omitempty"`
	ShipaddrPhone null.String `boil:"shipaddr_phone" json:"shipaddr_phone,omitempty" toml:"shipaddr_phone" yaml:"shipaddr_phone,omitempty"`
	ShipaddrFax   null.String `boil:"shipaddr_fax" json:"shipaddr_fax,omitempty" toml:"shipaddr_fax" yaml:"shipaddr_fax,omitempty"`
	ShipaddrEmail null.String `boil:"shipaddr_email" json:"shipaddr_email,omitempty" toml:"shipaddr_email" yaml:"shipaddr_email,omitempty"`
	Terms         null.String `boil:"terms" json:"terms,omitempty" toml:"terms" yaml:"terms,omitempty"`
	TaxIncluded   null.Int64  `boil:"tax_included" json:"tax_included,omitempty" toml:"tax_included" yaml:"tax_included,omitempty"`
	Taxtable      null.String `boil:"taxtable" json:"taxtable,omitempty" toml:"taxtable" yaml:"taxtable,omitempty"`

	R *customerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L customerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CustomerColumns = struct {
	GUID          string
	Name          string
	ID            string
	Notes         string
	Active        string
	DiscountNum   string
	DiscountDenom string
	CreditNum     string
	CreditDenom   string
	Currency      string
	TaxOverride   string
	AddrName      string
	AddrAddr1     string
	AddrAddr2     string
	AddrAddr3     string
	AddrAddr4     string
	AddrPhone     string
	AddrFax       string
	AddrEmail     string
	ShipaddrName  string
	ShipaddrAddr1 string
	ShipaddrAddr2 string
	ShipaddrAddr3 string
	ShipaddrAddr4 string
	ShipaddrPhone string
	ShipaddrFax   string
	ShipaddrEmail string
	Terms         string
	TaxIncluded   string
	Taxtable      string
}{
	GUID:          "guid",
	Name:          "name",
	ID:            "id",
	Notes:         "notes",
	Active:        "active",
	DiscountNum:   "discount_num",
	DiscountDenom: "discount_denom",
	CreditNum:     "credit_num",
	CreditDenom:   "credit_denom",
	Currency:      "currency",
	TaxOverride:   "tax_override",
	AddrName:      "addr_name",
	AddrAddr1:     "addr_addr1",
	AddrAddr2:     "addr_addr2",
	AddrAddr3:     "addr_addr3",
	AddrAddr4:     "addr_addr4",
	AddrPhone:     "addr_phone",
	AddrFax:       "addr_fax",
	AddrEmail:     "addr_email",
	ShipaddrName:  "shipaddr_name",
	ShipaddrAddr1: "shipaddr_addr1",
	ShipaddrAddr2: "shipaddr_addr2",
	ShipaddrAddr3: "shipaddr_addr3",
	ShipaddrAddr4: "shipaddr_addr4",
	ShipaddrPhone: "shipaddr_phone",
	ShipaddrFax:   "shipaddr_fax",
	ShipaddrEmail: "shipaddr_email",
	Terms:         "terms",
	TaxIncluded:   "tax_included",
	Taxtable:      "taxtable",
}

var CustomerTableColumns = struct {
	GUID          string
	Name          string
	ID            string
	Notes         string
	Active        string
	DiscountNum   string
	DiscountDenom string
	CreditNum     string
	CreditDenom   string
	Currency      string
	TaxOverride   string
	AddrName      string
	AddrAddr1     string
	AddrAddr2     string
	AddrAddr3     string
	AddrAddr4     string
	AddrPhone     string
	AddrFax       string
	AddrEmail     string
	ShipaddrName  string
	ShipaddrAddr1 string
	ShipaddrAddr2 string
	ShipaddrAddr3 string
	ShipaddrAddr4 string
	ShipaddrPhone string
	ShipaddrFax   string
	ShipaddrEmail string
	Terms         string
	TaxIncluded   string
	Taxtable      string
}{
	GUID:          "customers.guid",
	Name:          "customers.name",
	ID:            "customers.id",
	Notes:         "customers.notes",
	Active:        "customers.active",
	DiscountNum:   "customers.discount_num",
	DiscountDenom: "customers.discount_denom",
	CreditNum:     "customers.credit_num",
	CreditDenom:   "customers.credit_denom",
	Currency:      "customers.currency",
	TaxOverride:   "customers.tax_override",
	AddrName:      "customers.addr_name",
	AddrAddr1:     "customers.addr_addr1",
	AddrAddr2:     "customers.addr_addr2",
	AddrAddr3:     "customers.addr_addr3",
	AddrAddr4:     "customers.addr_addr4",
	AddrPhone:     "customers.addr_phone",
	AddrFax:       "customers.addr_fax",
	AddrEmail:     "customers.addr_email",
	ShipaddrName:  "customers.shipaddr_name",
	ShipaddrAddr1: "customers.shipaddr_addr1",
	ShipaddrAddr2: "customers.shipaddr_addr2",
	ShipaddrAddr3: "customers.shipaddr_addr3",
	ShipaddrAddr4: "customers.shipaddr_addr4",
	ShipaddrPhone: "customers.shipaddr_phone",
	ShipaddrFax:   "customers.shipaddr_fax",
	ShipaddrEmail: "customers.shipaddr_email",
	Terms:         "customers.terms",
	TaxIncluded:   "customers.tax_included",
	Taxtable:      "customers.taxtable",
}

// Generated where

var CustomerWhere = struct {
	GUID          whereHelperstring
	Name          whereHelperstring
	ID            whereHelperstring
	Notes         whereHelperstring
	Active        whereHelperint64
	DiscountNum   whereHelperint64
	DiscountDenom whereHelperint64
	CreditNum     whereHelperint64
	CreditDenom   whereHelperint64
	Currency      whereHelperstring
	TaxOverride   whereHelperint64
	AddrName      whereHelpernull_String
	AddrAddr1     whereHelpernull_String
	AddrAddr2     whereHelpernull_String
	AddrAddr3     whereHelpernull_String
	AddrAddr4     whereHelpernull_String
	AddrPhone     whereHelpernull_String
	AddrFax       whereHelpernull_String
	AddrEmail     whereHelpernull_String
	ShipaddrName  whereHelpernull_String
	ShipaddrAddr1 whereHelpernull_String
	ShipaddrAddr2 whereHelpernull_String
	ShipaddrAddr3 whereHelpernull_String
	ShipaddrAddr4 whereHelpernull_String
	ShipaddrPhone whereHelpernull_String
	ShipaddrFax   whereHelpernull_String
	ShipaddrEmail whereHelpernull_String
	Terms         whereHelpernull_String
	TaxIncluded   whereHelpernull_Int64
	Taxtable      whereHelpernull_String
}{
	GUID:          whereHelperstring{field: "\"customers\".\"guid\""},
	Name:          whereHelperstring{field: "\"customers\".\"name\""},
	ID:            whereHelperstring{field: "\"customers\".\"id\""},
	Notes:         whereHelperstring{field: "\"customers\".\"notes\""},
	Active:        whereHelperint64{field: "\"customers\".\"active\""},
	DiscountNum:   whereHelperint64{field: "\"customers\".\"discount_num\""},
	DiscountDenom: whereHelperint64{field: "\"customers\".\"discount_denom\""},
	CreditNum:     whereHelperint64{field: "\"customers\".\"credit_num\""},
	CreditDenom:   whereHelperint64{field: "\"customers\".\"credit_denom\""},
	Currency:      whereHelperstring{field: "\"customers\".\"currency\""},
	TaxOverride:   whereHelperint64{field: "\"customers\".\"tax_override\""},
	AddrName:      whereHelpernull_String{field: "\"customers\".\"addr_name\""},
	AddrAddr1:     whereHelpernull_String{field: "\"customers\".\"addr_addr1\""},
	AddrAddr2:     whereHelpernull_String{field: "\"customers\".\"addr_addr2\""},
	AddrAddr3:     whereHelpernull_String{field: "\"customers\".\"addr_addr3\""},
	AddrAddr4:     whereHelpernull_String{field: "\"customers\".\"addr_addr4\""},
	AddrPhone:     whereHelpernull_String{field: "\"customers\".\"addr_phone\""},
	AddrFax:       whereHelpernull_String{field: "\"customers\".\"addr_fax\""},
	AddrEmail:     whereHelpernull_String{field: "\"customers\".\"addr_email\""},
	ShipaddrName:  whereHelpernull_String{field: "\"customers\".\"shipaddr_name\""},
	ShipaddrAddr1: whereHelpernull_String{field: "\"customers\".\"shipaddr_addr1\""},
	ShipaddrAddr2: whereHelpernull_String{field: "\"customers\".\"shipaddr_addr2\""},
	ShipaddrAddr3: whereHelpernull_String{field: "\"customers\".\"shipaddr_addr3\""},
	ShipaddrAddr4: whereHelpernull_String{field: "\"customers\".\"shipaddr_addr4\""},
	ShipaddrPhone: whereHelpernull_String{field: "\"customers\".\"shipaddr_phone\""},
	ShipaddrFax:   whereHelpernull_String{field: "\"customers\".\"shipaddr_fax\""},
	ShipaddrEmail: whereHelpernull_String{field: "\"customers\".\"shipaddr_email\""},
	Terms:         whereHelpernull_String{field: "\"customers\".\"terms\""},
	TaxIncluded:   whereHelpernull_Int64{field: "\"customers\".\"tax_included\""},
	Taxtable:      whereHelpernull_String{field: "\"customers\".\"taxtable\""},
}

// CustomerRels is where relationship names are stored.
var CustomerRels = struct {
}{}

// customerR is where relationships are stored.
type customerR struct {
}

// NewStruct creates a new relationship struct
func (*customerR) NewStruct() *customerR {
	return &customerR{}
}

// customerL is where Load methods for each relationship are stored.
type customerL struct{}

var (
	customerAllColumns            = []string{"guid", "name", "id", "notes", "active", "discount_num", "discount_denom", "credit_num", "credit_denom", "currency", "tax_override", "addr_name", "addr_addr1", "addr_addr2", "addr_addr3", "addr_addr4", "addr_phone", "addr_fax", "addr_email", "shipaddr_name", "shipaddr_addr1", "shipaddr_addr2", "shipaddr_addr3", "shipaddr_addr4", "shipaddr_phone", "shipaddr_fax", "shipaddr_email", "terms", "tax_included", "taxtable"}
	customerColumnsWithoutDefault = []string{"guid", "name", "id", "notes", "active", "discount_num", "discount_denom", "credit_num", "credit_denom", "currency", "tax_override"}
	customerColumnsWithDefault    = []string{"addr_name", "addr_addr1", "addr_addr2", "addr_addr3", "addr_addr4", "addr_phone", "addr_fax", "addr_email", "shipaddr_name", "shipaddr_addr1", "shipaddr_addr2", "shipaddr_addr3", "shipaddr_addr4", "shipaddr_phone", "shipaddr_fax", "shipaddr_email", "terms", "tax_included", "taxtable"}
	customerPrimaryKeyColumns     = []string{"guid"}
	customerGeneratedColumns      = []string{}
)

type (
	// CustomerSlice is an alias for a slice of pointers to Customer.
	// This should almost always be used instead of []Customer.
	CustomerSlice []*Customer
	// CustomerHook is the signature for custom Customer hook methods
	CustomerHook func(context.Context, boil.ContextExecutor, *Customer) error

	customerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	customerType                 = reflect.TypeOf(&Customer{})
	customerMapping              = queries.MakeStructMapping(customerType)
	customerPrimaryKeyMapping, _ = queries.BindMapping(customerType, customerMapping, customerPrimaryKeyColumns)
	customerInsertCacheMut       sync.RWMutex
	customerInsertCache          = make(map[string]insertCache)
	customerUpdateCacheMut       sync.RWMutex
	customerUpdateCache          = make(map[string]updateCache)
	customerUpsertCacheMut       sync.RWMutex
	customerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var customerAfterSelectHooks []CustomerHook

var customerBeforeInsertHooks []CustomerHook
var customerAfterInsertHooks []CustomerHook

var customerBeforeUpdateHooks []CustomerHook
var customerAfterUpdateHooks []CustomerHook

var customerBeforeDeleteHooks []CustomerHook
var customerAfterDeleteHooks []CustomerHook

var customerBeforeUpsertHooks []CustomerHook
var customerAfterUpsertHooks []CustomerHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Customer) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Customer) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Customer) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Customer) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Customer) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Customer) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Customer) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Customer) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Customer) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCustomerHook registers your hook function for all future operations.
func AddCustomerHook(hookPoint boil.HookPoint, customerHook CustomerHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		customerAfterSelectHooks = append(customerAfterSelectHooks, customerHook)
	case boil.BeforeInsertHook:
		customerBeforeInsertHooks = append(customerBeforeInsertHooks, customerHook)
	case boil.AfterInsertHook:
		customerAfterInsertHooks = append(customerAfterInsertHooks, customerHook)
	case boil.BeforeUpdateHook:
		customerBeforeUpdateHooks = append(customerBeforeUpdateHooks, customerHook)
	case boil.AfterUpdateHook:
		customerAfterUpdateHooks = append(customerAfterUpdateHooks, customerHook)
	case boil.BeforeDeleteHook:
		customerBeforeDeleteHooks = append(customerBeforeDeleteHooks, customerHook)
	case boil.AfterDeleteHook:
		customerAfterDeleteHooks = append(customerAfterDeleteHooks, customerHook)
	case boil.BeforeUpsertHook:
		customerBeforeUpsertHooks = append(customerBeforeUpsertHooks, customerHook)
	case boil.AfterUpsertHook:
		customerAfterUpsertHooks = append(customerAfterUpsertHooks, customerHook)
	}
}

// One returns a single customer record from the query.
func (q customerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Customer, error) {
	o := &Customer{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "gnucash: failed to execute a one query for customers")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Customer records from the query.
func (q customerQuery) All(ctx context.Context, exec boil.ContextExecutor) (CustomerSlice, error) {
	var o []*Customer

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "gnucash: failed to assign all query results to Customer slice")
	}

	if len(customerAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Customer records in the query.
func (q customerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to count customers rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q customerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "gnucash: failed to check if customers exists")
	}

	return count > 0, nil
}

// Customers retrieves all the records using an executor.
func Customers(mods ...qm.QueryMod) customerQuery {
	mods = append(mods, qm.From("\"customers\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"customers\".*"})
	}

	return customerQuery{q}
}

// FindCustomer retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCustomer(ctx context.Context, exec boil.ContextExecutor, gUID string, selectCols ...string) (*Customer, error) {
	customerObj := &Customer{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"customers\" where \"guid\"=?", sel,
	)

	q := queries.Raw(query, gUID)

	err := q.Bind(ctx, exec, customerObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "gnucash: unable to select from customers")
	}

	if err = customerObj.doAfterSelectHooks(ctx, exec); err != nil {
		return customerObj, err
	}

	return customerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Customer) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("gnucash: no customers provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(customerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	customerInsertCacheMut.RLock()
	cache, cached := customerInsertCache[key]
	customerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			customerAllColumns,
			customerColumnsWithDefault,
			customerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(customerType, customerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(customerType, customerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"customers\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"customers\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "gnucash: unable to insert into customers")
	}

	if !cached {
		customerInsertCacheMut.Lock()
		customerInsertCache[key] = cache
		customerInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Customer.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Customer) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	customerUpdateCacheMut.RLock()
	cache, cached := customerUpdateCache[key]
	customerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			customerAllColumns,
			customerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("gnucash: unable to update customers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"customers\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, customerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(customerType, customerMapping, append(wl, customerPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "gnucash: unable to update customers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by update for customers")
	}

	if !cached {
		customerUpdateCacheMut.Lock()
		customerUpdateCache[key] = cache
		customerUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q customerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to update all for customers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to retrieve rows affected for customers")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CustomerSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), customerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"customers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, customerPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to update all in customer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to retrieve rows affected all in update all customer")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Customer) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("gnucash: no customers provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(customerColumnsWithDefault, o)

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

	customerUpsertCacheMut.RLock()
	cache, cached := customerUpsertCache[key]
	customerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			customerAllColumns,
			customerColumnsWithDefault,
			customerColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			customerAllColumns,
			customerPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("gnucash: unable to upsert customers, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(customerPrimaryKeyColumns))
			copy(conflict, customerPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"customers\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(customerType, customerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(customerType, customerMapping, ret)
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
		return errors.Wrap(err, "gnucash: unable to upsert customers")
	}

	if !cached {
		customerUpsertCacheMut.Lock()
		customerUpsertCache[key] = cache
		customerUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Customer record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Customer) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("gnucash: no Customer provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), customerPrimaryKeyMapping)
	sql := "DELETE FROM \"customers\" WHERE \"guid\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to delete from customers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by delete for customers")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q customerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("gnucash: no customerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to delete all from customers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by deleteall for customers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CustomerSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(customerBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), customerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"customers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, customerPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to delete all from customer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by deleteall for customers")
	}

	if len(customerAfterDeleteHooks) != 0 {
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
func (o *Customer) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCustomer(ctx, exec, o.GUID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CustomerSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CustomerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), customerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"customers\".* FROM \"customers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, customerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "gnucash: unable to reload all in CustomerSlice")
	}

	*o = slice

	return nil
}

// CustomerExists checks if the Customer row exists.
func CustomerExists(ctx context.Context, exec boil.ContextExecutor, gUID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"customers\" where \"guid\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, gUID)
	}
	row := exec.QueryRowContext(ctx, sql, gUID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "gnucash: unable to check if customers exists")
	}

	return exists, nil
}

// Exists checks if the Customer row exists.
func (o *Customer) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return CustomerExists(ctx, exec, o.GUID)
}
