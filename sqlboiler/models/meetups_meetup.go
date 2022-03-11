// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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

// MeetupsMeetup is an object representing the database table.
type MeetupsMeetup struct {
	ID          int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title       string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	Address     string    `boil:"address" json:"address" toml:"address" yaml:"address"`
	Image       string    `boil:"image" json:"image" toml:"image" yaml:"image"`
	Description string    `boil:"description" json:"description" toml:"description" yaml:"description"`
	CreatedAt   null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt   null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt   null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *meetupsMeetupR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L meetupsMeetupL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MeetupsMeetupColumns = struct {
	ID          string
	Title       string
	Address     string
	Image       string
	Description string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}{
	ID:          "id",
	Title:       "title",
	Address:     "address",
	Image:       "image",
	Description: "description",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

var MeetupsMeetupTableColumns = struct {
	ID          string
	Title       string
	Address     string
	Image       string
	Description string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}{
	ID:          "meetups_meetup.id",
	Title:       "meetups_meetup.title",
	Address:     "meetups_meetup.address",
	Image:       "meetups_meetup.image",
	Description: "meetups_meetup.description",
	CreatedAt:   "meetups_meetup.created_at",
	UpdatedAt:   "meetups_meetup.updated_at",
	DeletedAt:   "meetups_meetup.deleted_at",
}

// Generated where

var MeetupsMeetupWhere = struct {
	ID          whereHelperint64
	Title       whereHelperstring
	Address     whereHelperstring
	Image       whereHelperstring
	Description whereHelperstring
	CreatedAt   whereHelpernull_Time
	UpdatedAt   whereHelpernull_Time
	DeletedAt   whereHelpernull_Time
}{
	ID:          whereHelperint64{field: "\"meetups_meetup\".\"id\""},
	Title:       whereHelperstring{field: "\"meetups_meetup\".\"title\""},
	Address:     whereHelperstring{field: "\"meetups_meetup\".\"address\""},
	Image:       whereHelperstring{field: "\"meetups_meetup\".\"image\""},
	Description: whereHelperstring{field: "\"meetups_meetup\".\"description\""},
	CreatedAt:   whereHelpernull_Time{field: "\"meetups_meetup\".\"created_at\""},
	UpdatedAt:   whereHelpernull_Time{field: "\"meetups_meetup\".\"updated_at\""},
	DeletedAt:   whereHelpernull_Time{field: "\"meetups_meetup\".\"deleted_at\""},
}

// MeetupsMeetupRels is where relationship names are stored.
var MeetupsMeetupRels = struct {
}{}

// meetupsMeetupR is where relationships are stored.
type meetupsMeetupR struct {
}

// NewStruct creates a new relationship struct
func (*meetupsMeetupR) NewStruct() *meetupsMeetupR {
	return &meetupsMeetupR{}
}

// meetupsMeetupL is where Load methods for each relationship are stored.
type meetupsMeetupL struct{}

var (
	meetupsMeetupAllColumns            = []string{"id", "title", "address", "image", "description", "created_at", "updated_at", "deleted_at"}
	meetupsMeetupColumnsWithoutDefault = []string{"title", "address", "image", "description"}
	meetupsMeetupColumnsWithDefault    = []string{"id", "created_at", "updated_at", "deleted_at"}
	meetupsMeetupPrimaryKeyColumns     = []string{"id"}
	meetupsMeetupGeneratedColumns      = []string{}
)

type (
	// MeetupsMeetupSlice is an alias for a slice of pointers to MeetupsMeetup.
	// This should almost always be used instead of []MeetupsMeetup.
	MeetupsMeetupSlice []*MeetupsMeetup
	// MeetupsMeetupHook is the signature for custom MeetupsMeetup hook methods
	MeetupsMeetupHook func(context.Context, boil.ContextExecutor, *MeetupsMeetup) error

	meetupsMeetupQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	meetupsMeetupType                 = reflect.TypeOf(&MeetupsMeetup{})
	meetupsMeetupMapping              = queries.MakeStructMapping(meetupsMeetupType)
	meetupsMeetupPrimaryKeyMapping, _ = queries.BindMapping(meetupsMeetupType, meetupsMeetupMapping, meetupsMeetupPrimaryKeyColumns)
	meetupsMeetupInsertCacheMut       sync.RWMutex
	meetupsMeetupInsertCache          = make(map[string]insertCache)
	meetupsMeetupUpdateCacheMut       sync.RWMutex
	meetupsMeetupUpdateCache          = make(map[string]updateCache)
	meetupsMeetupUpsertCacheMut       sync.RWMutex
	meetupsMeetupUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var meetupsMeetupAfterSelectHooks []MeetupsMeetupHook

var meetupsMeetupBeforeInsertHooks []MeetupsMeetupHook
var meetupsMeetupAfterInsertHooks []MeetupsMeetupHook

var meetupsMeetupBeforeUpdateHooks []MeetupsMeetupHook
var meetupsMeetupAfterUpdateHooks []MeetupsMeetupHook

var meetupsMeetupBeforeDeleteHooks []MeetupsMeetupHook
var meetupsMeetupAfterDeleteHooks []MeetupsMeetupHook

var meetupsMeetupBeforeUpsertHooks []MeetupsMeetupHook
var meetupsMeetupAfterUpsertHooks []MeetupsMeetupHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MeetupsMeetup) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range meetupsMeetupAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MeetupsMeetup) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range meetupsMeetupBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MeetupsMeetup) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range meetupsMeetupAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MeetupsMeetup) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range meetupsMeetupBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MeetupsMeetup) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range meetupsMeetupAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MeetupsMeetup) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range meetupsMeetupBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MeetupsMeetup) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range meetupsMeetupAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MeetupsMeetup) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range meetupsMeetupBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MeetupsMeetup) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range meetupsMeetupAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMeetupsMeetupHook registers your hook function for all future operations.
func AddMeetupsMeetupHook(hookPoint boil.HookPoint, meetupsMeetupHook MeetupsMeetupHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		meetupsMeetupAfterSelectHooks = append(meetupsMeetupAfterSelectHooks, meetupsMeetupHook)
	case boil.BeforeInsertHook:
		meetupsMeetupBeforeInsertHooks = append(meetupsMeetupBeforeInsertHooks, meetupsMeetupHook)
	case boil.AfterInsertHook:
		meetupsMeetupAfterInsertHooks = append(meetupsMeetupAfterInsertHooks, meetupsMeetupHook)
	case boil.BeforeUpdateHook:
		meetupsMeetupBeforeUpdateHooks = append(meetupsMeetupBeforeUpdateHooks, meetupsMeetupHook)
	case boil.AfterUpdateHook:
		meetupsMeetupAfterUpdateHooks = append(meetupsMeetupAfterUpdateHooks, meetupsMeetupHook)
	case boil.BeforeDeleteHook:
		meetupsMeetupBeforeDeleteHooks = append(meetupsMeetupBeforeDeleteHooks, meetupsMeetupHook)
	case boil.AfterDeleteHook:
		meetupsMeetupAfterDeleteHooks = append(meetupsMeetupAfterDeleteHooks, meetupsMeetupHook)
	case boil.BeforeUpsertHook:
		meetupsMeetupBeforeUpsertHooks = append(meetupsMeetupBeforeUpsertHooks, meetupsMeetupHook)
	case boil.AfterUpsertHook:
		meetupsMeetupAfterUpsertHooks = append(meetupsMeetupAfterUpsertHooks, meetupsMeetupHook)
	}
}

// One returns a single meetupsMeetup record from the query.
func (q meetupsMeetupQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MeetupsMeetup, error) {
	o := &MeetupsMeetup{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for meetups_meetup")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MeetupsMeetup records from the query.
func (q meetupsMeetupQuery) All(ctx context.Context, exec boil.ContextExecutor) (MeetupsMeetupSlice, error) {
	var o []*MeetupsMeetup

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MeetupsMeetup slice")
	}

	if len(meetupsMeetupAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MeetupsMeetup records in the query.
func (q meetupsMeetupQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count meetups_meetup rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q meetupsMeetupQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if meetups_meetup exists")
	}

	return count > 0, nil
}

// MeetupsMeetups retrieves all the records using an executor.
func MeetupsMeetups(mods ...qm.QueryMod) meetupsMeetupQuery {
	mods = append(mods, qm.From("\"meetups_meetup\""))
	return meetupsMeetupQuery{NewQuery(mods...)}
}

// FindMeetupsMeetup retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMeetupsMeetup(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*MeetupsMeetup, error) {
	meetupsMeetupObj := &MeetupsMeetup{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"meetups_meetup\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, meetupsMeetupObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from meetups_meetup")
	}

	if err = meetupsMeetupObj.doAfterSelectHooks(ctx, exec); err != nil {
		return meetupsMeetupObj, err
	}

	return meetupsMeetupObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MeetupsMeetup) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no meetups_meetup provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(meetupsMeetupColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	meetupsMeetupInsertCacheMut.RLock()
	cache, cached := meetupsMeetupInsertCache[key]
	meetupsMeetupInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			meetupsMeetupAllColumns,
			meetupsMeetupColumnsWithDefault,
			meetupsMeetupColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(meetupsMeetupType, meetupsMeetupMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(meetupsMeetupType, meetupsMeetupMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"meetups_meetup\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"meetups_meetup\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into meetups_meetup")
	}

	if !cached {
		meetupsMeetupInsertCacheMut.Lock()
		meetupsMeetupInsertCache[key] = cache
		meetupsMeetupInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the MeetupsMeetup.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MeetupsMeetup) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	meetupsMeetupUpdateCacheMut.RLock()
	cache, cached := meetupsMeetupUpdateCache[key]
	meetupsMeetupUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			meetupsMeetupAllColumns,
			meetupsMeetupPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update meetups_meetup, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"meetups_meetup\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, meetupsMeetupPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(meetupsMeetupType, meetupsMeetupMapping, append(wl, meetupsMeetupPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update meetups_meetup row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for meetups_meetup")
	}

	if !cached {
		meetupsMeetupUpdateCacheMut.Lock()
		meetupsMeetupUpdateCache[key] = cache
		meetupsMeetupUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q meetupsMeetupQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for meetups_meetup")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for meetups_meetup")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MeetupsMeetupSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), meetupsMeetupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"meetups_meetup\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, meetupsMeetupPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in meetupsMeetup slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all meetupsMeetup")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MeetupsMeetup) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no meetups_meetup provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(meetupsMeetupColumnsWithDefault, o)

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

	meetupsMeetupUpsertCacheMut.RLock()
	cache, cached := meetupsMeetupUpsertCache[key]
	meetupsMeetupUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			meetupsMeetupAllColumns,
			meetupsMeetupColumnsWithDefault,
			meetupsMeetupColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			meetupsMeetupAllColumns,
			meetupsMeetupPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert meetups_meetup, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(meetupsMeetupPrimaryKeyColumns))
			copy(conflict, meetupsMeetupPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"meetups_meetup\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(meetupsMeetupType, meetupsMeetupMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(meetupsMeetupType, meetupsMeetupMapping, ret)
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
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert meetups_meetup")
	}

	if !cached {
		meetupsMeetupUpsertCacheMut.Lock()
		meetupsMeetupUpsertCache[key] = cache
		meetupsMeetupUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single MeetupsMeetup record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MeetupsMeetup) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MeetupsMeetup provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), meetupsMeetupPrimaryKeyMapping)
	sql := "DELETE FROM \"meetups_meetup\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from meetups_meetup")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for meetups_meetup")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q meetupsMeetupQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no meetupsMeetupQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from meetups_meetup")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for meetups_meetup")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MeetupsMeetupSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(meetupsMeetupBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), meetupsMeetupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"meetups_meetup\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, meetupsMeetupPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from meetupsMeetup slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for meetups_meetup")
	}

	if len(meetupsMeetupAfterDeleteHooks) != 0 {
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
func (o *MeetupsMeetup) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMeetupsMeetup(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MeetupsMeetupSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MeetupsMeetupSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), meetupsMeetupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"meetups_meetup\".* FROM \"meetups_meetup\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, meetupsMeetupPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MeetupsMeetupSlice")
	}

	*o = slice

	return nil
}

// MeetupsMeetupExists checks if the MeetupsMeetup row exists.
func MeetupsMeetupExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"meetups_meetup\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if meetups_meetup exists")
	}

	return exists, nil
}
