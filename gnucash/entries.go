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

// Entry is an object representing the database table.
type Entry struct {
	GUID           string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	Date           string      `boil:"date" json:"date" toml:"date" yaml:"date"`
	DateEntered    null.String `boil:"date_entered" json:"date_entered,omitempty" toml:"date_entered" yaml:"date_entered,omitempty"`
	Description    null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	Action         null.String `boil:"action" json:"action,omitempty" toml:"action" yaml:"action,omitempty"`
	Notes          null.String `boil:"notes" json:"notes,omitempty" toml:"notes" yaml:"notes,omitempty"`
	QuantityNum    null.Int64  `boil:"quantity_num" json:"quantity_num,omitempty" toml:"quantity_num" yaml:"quantity_num,omitempty"`
	QuantityDenom  null.Int64  `boil:"quantity_denom" json:"quantity_denom,omitempty" toml:"quantity_denom" yaml:"quantity_denom,omitempty"`
	IAcct          null.String `boil:"i_acct" json:"i_acct,omitempty" toml:"i_acct" yaml:"i_acct,omitempty"`
	IPriceNum      null.Int64  `boil:"i_price_num" json:"i_price_num,omitempty" toml:"i_price_num" yaml:"i_price_num,omitempty"`
	IPriceDenom    null.Int64  `boil:"i_price_denom" json:"i_price_denom,omitempty" toml:"i_price_denom" yaml:"i_price_denom,omitempty"`
	IDiscountNum   null.Int64  `boil:"i_discount_num" json:"i_discount_num,omitempty" toml:"i_discount_num" yaml:"i_discount_num,omitempty"`
	IDiscountDenom null.Int64  `boil:"i_discount_denom" json:"i_discount_denom,omitempty" toml:"i_discount_denom" yaml:"i_discount_denom,omitempty"`
	Invoice        null.String `boil:"invoice" json:"invoice,omitempty" toml:"invoice" yaml:"invoice,omitempty"`
	IDiscType      null.String `boil:"i_disc_type" json:"i_disc_type,omitempty" toml:"i_disc_type" yaml:"i_disc_type,omitempty"`
	IDiscHow       null.String `boil:"i_disc_how" json:"i_disc_how,omitempty" toml:"i_disc_how" yaml:"i_disc_how,omitempty"`
	ITaxable       null.Int64  `boil:"i_taxable" json:"i_taxable,omitempty" toml:"i_taxable" yaml:"i_taxable,omitempty"`
	ITaxincluded   null.Int64  `boil:"i_taxincluded" json:"i_taxincluded,omitempty" toml:"i_taxincluded" yaml:"i_taxincluded,omitempty"`
	ITaxtable      null.String `boil:"i_taxtable" json:"i_taxtable,omitempty" toml:"i_taxtable" yaml:"i_taxtable,omitempty"`
	BAcct          null.String `boil:"b_acct" json:"b_acct,omitempty" toml:"b_acct" yaml:"b_acct,omitempty"`
	BPriceNum      null.Int64  `boil:"b_price_num" json:"b_price_num,omitempty" toml:"b_price_num" yaml:"b_price_num,omitempty"`
	BPriceDenom    null.Int64  `boil:"b_price_denom" json:"b_price_denom,omitempty" toml:"b_price_denom" yaml:"b_price_denom,omitempty"`
	Bill           null.String `boil:"bill" json:"bill,omitempty" toml:"bill" yaml:"bill,omitempty"`
	BTaxable       null.Int64  `boil:"b_taxable" json:"b_taxable,omitempty" toml:"b_taxable" yaml:"b_taxable,omitempty"`
	BTaxincluded   null.Int64  `boil:"b_taxincluded" json:"b_taxincluded,omitempty" toml:"b_taxincluded" yaml:"b_taxincluded,omitempty"`
	BTaxtable      null.String `boil:"b_taxtable" json:"b_taxtable,omitempty" toml:"b_taxtable" yaml:"b_taxtable,omitempty"`
	BPaytype       null.Int64  `boil:"b_paytype" json:"b_paytype,omitempty" toml:"b_paytype" yaml:"b_paytype,omitempty"`
	Billable       null.Int64  `boil:"billable" json:"billable,omitempty" toml:"billable" yaml:"billable,omitempty"`
	BilltoType     null.Int64  `boil:"billto_type" json:"billto_type,omitempty" toml:"billto_type" yaml:"billto_type,omitempty"`
	BilltoGUID     null.String `boil:"billto_guid" json:"billto_guid,omitempty" toml:"billto_guid" yaml:"billto_guid,omitempty"`
	OrderGUID      null.String `boil:"order_guid" json:"order_guid,omitempty" toml:"order_guid" yaml:"order_guid,omitempty"`

	R *entryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L entryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EntryColumns = struct {
	GUID           string
	Date           string
	DateEntered    string
	Description    string
	Action         string
	Notes          string
	QuantityNum    string
	QuantityDenom  string
	IAcct          string
	IPriceNum      string
	IPriceDenom    string
	IDiscountNum   string
	IDiscountDenom string
	Invoice        string
	IDiscType      string
	IDiscHow       string
	ITaxable       string
	ITaxincluded   string
	ITaxtable      string
	BAcct          string
	BPriceNum      string
	BPriceDenom    string
	Bill           string
	BTaxable       string
	BTaxincluded   string
	BTaxtable      string
	BPaytype       string
	Billable       string
	BilltoType     string
	BilltoGUID     string
	OrderGUID      string
}{
	GUID:           "guid",
	Date:           "date",
	DateEntered:    "date_entered",
	Description:    "description",
	Action:         "action",
	Notes:          "notes",
	QuantityNum:    "quantity_num",
	QuantityDenom:  "quantity_denom",
	IAcct:          "i_acct",
	IPriceNum:      "i_price_num",
	IPriceDenom:    "i_price_denom",
	IDiscountNum:   "i_discount_num",
	IDiscountDenom: "i_discount_denom",
	Invoice:        "invoice",
	IDiscType:      "i_disc_type",
	IDiscHow:       "i_disc_how",
	ITaxable:       "i_taxable",
	ITaxincluded:   "i_taxincluded",
	ITaxtable:      "i_taxtable",
	BAcct:          "b_acct",
	BPriceNum:      "b_price_num",
	BPriceDenom:    "b_price_denom",
	Bill:           "bill",
	BTaxable:       "b_taxable",
	BTaxincluded:   "b_taxincluded",
	BTaxtable:      "b_taxtable",
	BPaytype:       "b_paytype",
	Billable:       "billable",
	BilltoType:     "billto_type",
	BilltoGUID:     "billto_guid",
	OrderGUID:      "order_guid",
}

var EntryTableColumns = struct {
	GUID           string
	Date           string
	DateEntered    string
	Description    string
	Action         string
	Notes          string
	QuantityNum    string
	QuantityDenom  string
	IAcct          string
	IPriceNum      string
	IPriceDenom    string
	IDiscountNum   string
	IDiscountDenom string
	Invoice        string
	IDiscType      string
	IDiscHow       string
	ITaxable       string
	ITaxincluded   string
	ITaxtable      string
	BAcct          string
	BPriceNum      string
	BPriceDenom    string
	Bill           string
	BTaxable       string
	BTaxincluded   string
	BTaxtable      string
	BPaytype       string
	Billable       string
	BilltoType     string
	BilltoGUID     string
	OrderGUID      string
}{
	GUID:           "entries.guid",
	Date:           "entries.date",
	DateEntered:    "entries.date_entered",
	Description:    "entries.description",
	Action:         "entries.action",
	Notes:          "entries.notes",
	QuantityNum:    "entries.quantity_num",
	QuantityDenom:  "entries.quantity_denom",
	IAcct:          "entries.i_acct",
	IPriceNum:      "entries.i_price_num",
	IPriceDenom:    "entries.i_price_denom",
	IDiscountNum:   "entries.i_discount_num",
	IDiscountDenom: "entries.i_discount_denom",
	Invoice:        "entries.invoice",
	IDiscType:      "entries.i_disc_type",
	IDiscHow:       "entries.i_disc_how",
	ITaxable:       "entries.i_taxable",
	ITaxincluded:   "entries.i_taxincluded",
	ITaxtable:      "entries.i_taxtable",
	BAcct:          "entries.b_acct",
	BPriceNum:      "entries.b_price_num",
	BPriceDenom:    "entries.b_price_denom",
	Bill:           "entries.bill",
	BTaxable:       "entries.b_taxable",
	BTaxincluded:   "entries.b_taxincluded",
	BTaxtable:      "entries.b_taxtable",
	BPaytype:       "entries.b_paytype",
	Billable:       "entries.billable",
	BilltoType:     "entries.billto_type",
	BilltoGUID:     "entries.billto_guid",
	OrderGUID:      "entries.order_guid",
}

// Generated where

var EntryWhere = struct {
	GUID           whereHelperstring
	Date           whereHelperstring
	DateEntered    whereHelpernull_String
	Description    whereHelpernull_String
	Action         whereHelpernull_String
	Notes          whereHelpernull_String
	QuantityNum    whereHelpernull_Int64
	QuantityDenom  whereHelpernull_Int64
	IAcct          whereHelpernull_String
	IPriceNum      whereHelpernull_Int64
	IPriceDenom    whereHelpernull_Int64
	IDiscountNum   whereHelpernull_Int64
	IDiscountDenom whereHelpernull_Int64
	Invoice        whereHelpernull_String
	IDiscType      whereHelpernull_String
	IDiscHow       whereHelpernull_String
	ITaxable       whereHelpernull_Int64
	ITaxincluded   whereHelpernull_Int64
	ITaxtable      whereHelpernull_String
	BAcct          whereHelpernull_String
	BPriceNum      whereHelpernull_Int64
	BPriceDenom    whereHelpernull_Int64
	Bill           whereHelpernull_String
	BTaxable       whereHelpernull_Int64
	BTaxincluded   whereHelpernull_Int64
	BTaxtable      whereHelpernull_String
	BPaytype       whereHelpernull_Int64
	Billable       whereHelpernull_Int64
	BilltoType     whereHelpernull_Int64
	BilltoGUID     whereHelpernull_String
	OrderGUID      whereHelpernull_String
}{
	GUID:           whereHelperstring{field: "\"entries\".\"guid\""},
	Date:           whereHelperstring{field: "\"entries\".\"date\""},
	DateEntered:    whereHelpernull_String{field: "\"entries\".\"date_entered\""},
	Description:    whereHelpernull_String{field: "\"entries\".\"description\""},
	Action:         whereHelpernull_String{field: "\"entries\".\"action\""},
	Notes:          whereHelpernull_String{field: "\"entries\".\"notes\""},
	QuantityNum:    whereHelpernull_Int64{field: "\"entries\".\"quantity_num\""},
	QuantityDenom:  whereHelpernull_Int64{field: "\"entries\".\"quantity_denom\""},
	IAcct:          whereHelpernull_String{field: "\"entries\".\"i_acct\""},
	IPriceNum:      whereHelpernull_Int64{field: "\"entries\".\"i_price_num\""},
	IPriceDenom:    whereHelpernull_Int64{field: "\"entries\".\"i_price_denom\""},
	IDiscountNum:   whereHelpernull_Int64{field: "\"entries\".\"i_discount_num\""},
	IDiscountDenom: whereHelpernull_Int64{field: "\"entries\".\"i_discount_denom\""},
	Invoice:        whereHelpernull_String{field: "\"entries\".\"invoice\""},
	IDiscType:      whereHelpernull_String{field: "\"entries\".\"i_disc_type\""},
	IDiscHow:       whereHelpernull_String{field: "\"entries\".\"i_disc_how\""},
	ITaxable:       whereHelpernull_Int64{field: "\"entries\".\"i_taxable\""},
	ITaxincluded:   whereHelpernull_Int64{field: "\"entries\".\"i_taxincluded\""},
	ITaxtable:      whereHelpernull_String{field: "\"entries\".\"i_taxtable\""},
	BAcct:          whereHelpernull_String{field: "\"entries\".\"b_acct\""},
	BPriceNum:      whereHelpernull_Int64{field: "\"entries\".\"b_price_num\""},
	BPriceDenom:    whereHelpernull_Int64{field: "\"entries\".\"b_price_denom\""},
	Bill:           whereHelpernull_String{field: "\"entries\".\"bill\""},
	BTaxable:       whereHelpernull_Int64{field: "\"entries\".\"b_taxable\""},
	BTaxincluded:   whereHelpernull_Int64{field: "\"entries\".\"b_taxincluded\""},
	BTaxtable:      whereHelpernull_String{field: "\"entries\".\"b_taxtable\""},
	BPaytype:       whereHelpernull_Int64{field: "\"entries\".\"b_paytype\""},
	Billable:       whereHelpernull_Int64{field: "\"entries\".\"billable\""},
	BilltoType:     whereHelpernull_Int64{field: "\"entries\".\"billto_type\""},
	BilltoGUID:     whereHelpernull_String{field: "\"entries\".\"billto_guid\""},
	OrderGUID:      whereHelpernull_String{field: "\"entries\".\"order_guid\""},
}

// EntryRels is where relationship names are stored.
var EntryRels = struct {
}{}

// entryR is where relationships are stored.
type entryR struct {
}

// NewStruct creates a new relationship struct
func (*entryR) NewStruct() *entryR {
	return &entryR{}
}

// entryL is where Load methods for each relationship are stored.
type entryL struct{}

var (
	entryAllColumns            = []string{"guid", "date", "date_entered", "description", "action", "notes", "quantity_num", "quantity_denom", "i_acct", "i_price_num", "i_price_denom", "i_discount_num", "i_discount_denom", "invoice", "i_disc_type", "i_disc_how", "i_taxable", "i_taxincluded", "i_taxtable", "b_acct", "b_price_num", "b_price_denom", "bill", "b_taxable", "b_taxincluded", "b_taxtable", "b_paytype", "billable", "billto_type", "billto_guid", "order_guid"}
	entryColumnsWithoutDefault = []string{"guid", "date"}
	entryColumnsWithDefault    = []string{"date_entered", "description", "action", "notes", "quantity_num", "quantity_denom", "i_acct", "i_price_num", "i_price_denom", "i_discount_num", "i_discount_denom", "invoice", "i_disc_type", "i_disc_how", "i_taxable", "i_taxincluded", "i_taxtable", "b_acct", "b_price_num", "b_price_denom", "bill", "b_taxable", "b_taxincluded", "b_taxtable", "b_paytype", "billable", "billto_type", "billto_guid", "order_guid"}
	entryPrimaryKeyColumns     = []string{"guid"}
	entryGeneratedColumns      = []string{}
)

type (
	// EntrySlice is an alias for a slice of pointers to Entry.
	// This should almost always be used instead of []Entry.
	EntrySlice []*Entry
	// EntryHook is the signature for custom Entry hook methods
	EntryHook func(context.Context, boil.ContextExecutor, *Entry) error

	entryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	entryType                 = reflect.TypeOf(&Entry{})
	entryMapping              = queries.MakeStructMapping(entryType)
	entryPrimaryKeyMapping, _ = queries.BindMapping(entryType, entryMapping, entryPrimaryKeyColumns)
	entryInsertCacheMut       sync.RWMutex
	entryInsertCache          = make(map[string]insertCache)
	entryUpdateCacheMut       sync.RWMutex
	entryUpdateCache          = make(map[string]updateCache)
	entryUpsertCacheMut       sync.RWMutex
	entryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var entryAfterSelectHooks []EntryHook

var entryBeforeInsertHooks []EntryHook
var entryAfterInsertHooks []EntryHook

var entryBeforeUpdateHooks []EntryHook
var entryAfterUpdateHooks []EntryHook

var entryBeforeDeleteHooks []EntryHook
var entryAfterDeleteHooks []EntryHook

var entryBeforeUpsertHooks []EntryHook
var entryAfterUpsertHooks []EntryHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Entry) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Entry) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Entry) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Entry) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Entry) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Entry) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Entry) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Entry) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Entry) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEntryHook registers your hook function for all future operations.
func AddEntryHook(hookPoint boil.HookPoint, entryHook EntryHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		entryAfterSelectHooks = append(entryAfterSelectHooks, entryHook)
	case boil.BeforeInsertHook:
		entryBeforeInsertHooks = append(entryBeforeInsertHooks, entryHook)
	case boil.AfterInsertHook:
		entryAfterInsertHooks = append(entryAfterInsertHooks, entryHook)
	case boil.BeforeUpdateHook:
		entryBeforeUpdateHooks = append(entryBeforeUpdateHooks, entryHook)
	case boil.AfterUpdateHook:
		entryAfterUpdateHooks = append(entryAfterUpdateHooks, entryHook)
	case boil.BeforeDeleteHook:
		entryBeforeDeleteHooks = append(entryBeforeDeleteHooks, entryHook)
	case boil.AfterDeleteHook:
		entryAfterDeleteHooks = append(entryAfterDeleteHooks, entryHook)
	case boil.BeforeUpsertHook:
		entryBeforeUpsertHooks = append(entryBeforeUpsertHooks, entryHook)
	case boil.AfterUpsertHook:
		entryAfterUpsertHooks = append(entryAfterUpsertHooks, entryHook)
	}
}

// One returns a single entry record from the query.
func (q entryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Entry, error) {
	o := &Entry{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "gnucash: failed to execute a one query for entries")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Entry records from the query.
func (q entryQuery) All(ctx context.Context, exec boil.ContextExecutor) (EntrySlice, error) {
	var o []*Entry

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "gnucash: failed to assign all query results to Entry slice")
	}

	if len(entryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Entry records in the query.
func (q entryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to count entries rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q entryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "gnucash: failed to check if entries exists")
	}

	return count > 0, nil
}

// Entries retrieves all the records using an executor.
func Entries(mods ...qm.QueryMod) entryQuery {
	mods = append(mods, qm.From("\"entries\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"entries\".*"})
	}

	return entryQuery{q}
}

// FindEntry retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEntry(ctx context.Context, exec boil.ContextExecutor, gUID string, selectCols ...string) (*Entry, error) {
	entryObj := &Entry{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"entries\" where \"guid\"=?", sel,
	)

	q := queries.Raw(query, gUID)

	err := q.Bind(ctx, exec, entryObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "gnucash: unable to select from entries")
	}

	if err = entryObj.doAfterSelectHooks(ctx, exec); err != nil {
		return entryObj, err
	}

	return entryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Entry) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("gnucash: no entries provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(entryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	entryInsertCacheMut.RLock()
	cache, cached := entryInsertCache[key]
	entryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			entryAllColumns,
			entryColumnsWithDefault,
			entryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(entryType, entryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(entryType, entryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"entries\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"entries\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "gnucash: unable to insert into entries")
	}

	if !cached {
		entryInsertCacheMut.Lock()
		entryInsertCache[key] = cache
		entryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Entry.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Entry) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	entryUpdateCacheMut.RLock()
	cache, cached := entryUpdateCache[key]
	entryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			entryAllColumns,
			entryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("gnucash: unable to update entries, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"entries\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, entryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(entryType, entryMapping, append(wl, entryPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "gnucash: unable to update entries row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by update for entries")
	}

	if !cached {
		entryUpdateCacheMut.Lock()
		entryUpdateCache[key] = cache
		entryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q entryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to update all for entries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to retrieve rows affected for entries")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EntrySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), entryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"entries\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, entryPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to update all in entry slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to retrieve rows affected all in update all entry")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Entry) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("gnucash: no entries provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(entryColumnsWithDefault, o)

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

	entryUpsertCacheMut.RLock()
	cache, cached := entryUpsertCache[key]
	entryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			entryAllColumns,
			entryColumnsWithDefault,
			entryColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			entryAllColumns,
			entryPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("gnucash: unable to upsert entries, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(entryPrimaryKeyColumns))
			copy(conflict, entryPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"entries\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(entryType, entryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(entryType, entryMapping, ret)
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
		return errors.Wrap(err, "gnucash: unable to upsert entries")
	}

	if !cached {
		entryUpsertCacheMut.Lock()
		entryUpsertCache[key] = cache
		entryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Entry record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Entry) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("gnucash: no Entry provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), entryPrimaryKeyMapping)
	sql := "DELETE FROM \"entries\" WHERE \"guid\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to delete from entries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by delete for entries")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q entryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("gnucash: no entryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to delete all from entries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by deleteall for entries")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EntrySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(entryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), entryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"entries\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, entryPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: unable to delete all from entry slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "gnucash: failed to get rows affected by deleteall for entries")
	}

	if len(entryAfterDeleteHooks) != 0 {
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
func (o *Entry) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEntry(ctx, exec, o.GUID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EntrySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EntrySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), entryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"entries\".* FROM \"entries\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, entryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "gnucash: unable to reload all in EntrySlice")
	}

	*o = slice

	return nil
}

// EntryExists checks if the Entry row exists.
func EntryExists(ctx context.Context, exec boil.ContextExecutor, gUID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"entries\" where \"guid\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, gUID)
	}
	row := exec.QueryRowContext(ctx, sql, gUID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "gnucash: unable to check if entries exists")
	}

	return exists, nil
}

// Exists checks if the Entry row exists.
func (o *Entry) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return EntryExists(ctx, exec, o.GUID)
}