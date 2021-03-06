// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testMeetupsMeetups(t *testing.T) {
	t.Parallel()

	query := MeetupsMeetups()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMeetupsMeetupsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MeetupsMeetup{}
	if err = randomize.Struct(seed, o, meetupsMeetupDBTypes, true, meetupsMeetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MeetupsMeetups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMeetupsMeetupsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MeetupsMeetup{}
	if err = randomize.Struct(seed, o, meetupsMeetupDBTypes, true, meetupsMeetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := MeetupsMeetups().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MeetupsMeetups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMeetupsMeetupsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MeetupsMeetup{}
	if err = randomize.Struct(seed, o, meetupsMeetupDBTypes, true, meetupsMeetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MeetupsMeetupSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MeetupsMeetups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMeetupsMeetupsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MeetupsMeetup{}
	if err = randomize.Struct(seed, o, meetupsMeetupDBTypes, true, meetupsMeetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MeetupsMeetupExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if MeetupsMeetup exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MeetupsMeetupExists to return true, but got false.")
	}
}

func testMeetupsMeetupsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MeetupsMeetup{}
	if err = randomize.Struct(seed, o, meetupsMeetupDBTypes, true, meetupsMeetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	meetupsMeetupFound, err := FindMeetupsMeetup(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if meetupsMeetupFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMeetupsMeetupsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MeetupsMeetup{}
	if err = randomize.Struct(seed, o, meetupsMeetupDBTypes, true, meetupsMeetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = MeetupsMeetups().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMeetupsMeetupsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MeetupsMeetup{}
	if err = randomize.Struct(seed, o, meetupsMeetupDBTypes, true, meetupsMeetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := MeetupsMeetups().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMeetupsMeetupsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	meetupsMeetupOne := &MeetupsMeetup{}
	meetupsMeetupTwo := &MeetupsMeetup{}
	if err = randomize.Struct(seed, meetupsMeetupOne, meetupsMeetupDBTypes, false, meetupsMeetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup struct: %s", err)
	}
	if err = randomize.Struct(seed, meetupsMeetupTwo, meetupsMeetupDBTypes, false, meetupsMeetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = meetupsMeetupOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = meetupsMeetupTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MeetupsMeetups().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMeetupsMeetupsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	meetupsMeetupOne := &MeetupsMeetup{}
	meetupsMeetupTwo := &MeetupsMeetup{}
	if err = randomize.Struct(seed, meetupsMeetupOne, meetupsMeetupDBTypes, false, meetupsMeetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup struct: %s", err)
	}
	if err = randomize.Struct(seed, meetupsMeetupTwo, meetupsMeetupDBTypes, false, meetupsMeetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = meetupsMeetupOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = meetupsMeetupTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MeetupsMeetups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func meetupsMeetupBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *MeetupsMeetup) error {
	*o = MeetupsMeetup{}
	return nil
}

func meetupsMeetupAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *MeetupsMeetup) error {
	*o = MeetupsMeetup{}
	return nil
}

func meetupsMeetupAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *MeetupsMeetup) error {
	*o = MeetupsMeetup{}
	return nil
}

func meetupsMeetupBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MeetupsMeetup) error {
	*o = MeetupsMeetup{}
	return nil
}

func meetupsMeetupAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MeetupsMeetup) error {
	*o = MeetupsMeetup{}
	return nil
}

func meetupsMeetupBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MeetupsMeetup) error {
	*o = MeetupsMeetup{}
	return nil
}

func meetupsMeetupAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MeetupsMeetup) error {
	*o = MeetupsMeetup{}
	return nil
}

func meetupsMeetupBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MeetupsMeetup) error {
	*o = MeetupsMeetup{}
	return nil
}

func meetupsMeetupAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MeetupsMeetup) error {
	*o = MeetupsMeetup{}
	return nil
}

func testMeetupsMeetupsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &MeetupsMeetup{}
	o := &MeetupsMeetup{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, meetupsMeetupDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup object: %s", err)
	}

	AddMeetupsMeetupHook(boil.BeforeInsertHook, meetupsMeetupBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	meetupsMeetupBeforeInsertHooks = []MeetupsMeetupHook{}

	AddMeetupsMeetupHook(boil.AfterInsertHook, meetupsMeetupAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	meetupsMeetupAfterInsertHooks = []MeetupsMeetupHook{}

	AddMeetupsMeetupHook(boil.AfterSelectHook, meetupsMeetupAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	meetupsMeetupAfterSelectHooks = []MeetupsMeetupHook{}

	AddMeetupsMeetupHook(boil.BeforeUpdateHook, meetupsMeetupBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	meetupsMeetupBeforeUpdateHooks = []MeetupsMeetupHook{}

	AddMeetupsMeetupHook(boil.AfterUpdateHook, meetupsMeetupAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	meetupsMeetupAfterUpdateHooks = []MeetupsMeetupHook{}

	AddMeetupsMeetupHook(boil.BeforeDeleteHook, meetupsMeetupBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	meetupsMeetupBeforeDeleteHooks = []MeetupsMeetupHook{}

	AddMeetupsMeetupHook(boil.AfterDeleteHook, meetupsMeetupAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	meetupsMeetupAfterDeleteHooks = []MeetupsMeetupHook{}

	AddMeetupsMeetupHook(boil.BeforeUpsertHook, meetupsMeetupBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	meetupsMeetupBeforeUpsertHooks = []MeetupsMeetupHook{}

	AddMeetupsMeetupHook(boil.AfterUpsertHook, meetupsMeetupAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	meetupsMeetupAfterUpsertHooks = []MeetupsMeetupHook{}
}

func testMeetupsMeetupsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MeetupsMeetup{}
	if err = randomize.Struct(seed, o, meetupsMeetupDBTypes, true, meetupsMeetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MeetupsMeetups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMeetupsMeetupsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MeetupsMeetup{}
	if err = randomize.Struct(seed, o, meetupsMeetupDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(meetupsMeetupColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := MeetupsMeetups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMeetupsMeetupsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MeetupsMeetup{}
	if err = randomize.Struct(seed, o, meetupsMeetupDBTypes, true, meetupsMeetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMeetupsMeetupsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MeetupsMeetup{}
	if err = randomize.Struct(seed, o, meetupsMeetupDBTypes, true, meetupsMeetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MeetupsMeetupSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMeetupsMeetupsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MeetupsMeetup{}
	if err = randomize.Struct(seed, o, meetupsMeetupDBTypes, true, meetupsMeetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MeetupsMeetups().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	meetupsMeetupDBTypes = map[string]string{`ID`: `bigint`, `Title`: `character varying`, `Address`: `character varying`, `Image`: `character varying`, `Description`: `text`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`}
	_                    = bytes.MinRead
)

func testMeetupsMeetupsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(meetupsMeetupPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(meetupsMeetupAllColumns) == len(meetupsMeetupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MeetupsMeetup{}
	if err = randomize.Struct(seed, o, meetupsMeetupDBTypes, true, meetupsMeetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MeetupsMeetups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, meetupsMeetupDBTypes, true, meetupsMeetupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMeetupsMeetupsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(meetupsMeetupAllColumns) == len(meetupsMeetupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MeetupsMeetup{}
	if err = randomize.Struct(seed, o, meetupsMeetupDBTypes, true, meetupsMeetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MeetupsMeetups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, meetupsMeetupDBTypes, true, meetupsMeetupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(meetupsMeetupAllColumns, meetupsMeetupPrimaryKeyColumns) {
		fields = meetupsMeetupAllColumns
	} else {
		fields = strmangle.SetComplement(
			meetupsMeetupAllColumns,
			meetupsMeetupPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := MeetupsMeetupSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMeetupsMeetupsUpsert(t *testing.T) {
	t.Parallel()

	if len(meetupsMeetupAllColumns) == len(meetupsMeetupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := MeetupsMeetup{}
	if err = randomize.Struct(seed, &o, meetupsMeetupDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MeetupsMeetup: %s", err)
	}

	count, err := MeetupsMeetups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, meetupsMeetupDBTypes, false, meetupsMeetupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MeetupsMeetup struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MeetupsMeetup: %s", err)
	}

	count, err = MeetupsMeetups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
