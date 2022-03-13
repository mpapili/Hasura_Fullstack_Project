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

func testDjangoAdminLogs(t *testing.T) {
	t.Parallel()

	query := DjangoAdminLogs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDjangoAdminLogsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoAdminLog{}
	if err = randomize.Struct(seed, o, djangoAdminLogDBTypes, true, djangoAdminLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
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

	count, err := DjangoAdminLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDjangoAdminLogsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoAdminLog{}
	if err = randomize.Struct(seed, o, djangoAdminLogDBTypes, true, djangoAdminLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DjangoAdminLogs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DjangoAdminLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDjangoAdminLogsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoAdminLog{}
	if err = randomize.Struct(seed, o, djangoAdminLogDBTypes, true, djangoAdminLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DjangoAdminLogSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DjangoAdminLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDjangoAdminLogsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoAdminLog{}
	if err = randomize.Struct(seed, o, djangoAdminLogDBTypes, true, djangoAdminLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DjangoAdminLogExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DjangoAdminLog exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DjangoAdminLogExists to return true, but got false.")
	}
}

func testDjangoAdminLogsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoAdminLog{}
	if err = randomize.Struct(seed, o, djangoAdminLogDBTypes, true, djangoAdminLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	djangoAdminLogFound, err := FindDjangoAdminLog(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if djangoAdminLogFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDjangoAdminLogsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoAdminLog{}
	if err = randomize.Struct(seed, o, djangoAdminLogDBTypes, true, djangoAdminLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DjangoAdminLogs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDjangoAdminLogsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoAdminLog{}
	if err = randomize.Struct(seed, o, djangoAdminLogDBTypes, true, djangoAdminLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DjangoAdminLogs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDjangoAdminLogsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	djangoAdminLogOne := &DjangoAdminLog{}
	djangoAdminLogTwo := &DjangoAdminLog{}
	if err = randomize.Struct(seed, djangoAdminLogOne, djangoAdminLogDBTypes, false, djangoAdminLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
	}
	if err = randomize.Struct(seed, djangoAdminLogTwo, djangoAdminLogDBTypes, false, djangoAdminLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = djangoAdminLogOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = djangoAdminLogTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DjangoAdminLogs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDjangoAdminLogsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	djangoAdminLogOne := &DjangoAdminLog{}
	djangoAdminLogTwo := &DjangoAdminLog{}
	if err = randomize.Struct(seed, djangoAdminLogOne, djangoAdminLogDBTypes, false, djangoAdminLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
	}
	if err = randomize.Struct(seed, djangoAdminLogTwo, djangoAdminLogDBTypes, false, djangoAdminLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = djangoAdminLogOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = djangoAdminLogTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DjangoAdminLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func djangoAdminLogBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DjangoAdminLog) error {
	*o = DjangoAdminLog{}
	return nil
}

func djangoAdminLogAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DjangoAdminLog) error {
	*o = DjangoAdminLog{}
	return nil
}

func djangoAdminLogAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DjangoAdminLog) error {
	*o = DjangoAdminLog{}
	return nil
}

func djangoAdminLogBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DjangoAdminLog) error {
	*o = DjangoAdminLog{}
	return nil
}

func djangoAdminLogAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DjangoAdminLog) error {
	*o = DjangoAdminLog{}
	return nil
}

func djangoAdminLogBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DjangoAdminLog) error {
	*o = DjangoAdminLog{}
	return nil
}

func djangoAdminLogAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DjangoAdminLog) error {
	*o = DjangoAdminLog{}
	return nil
}

func djangoAdminLogBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DjangoAdminLog) error {
	*o = DjangoAdminLog{}
	return nil
}

func djangoAdminLogAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DjangoAdminLog) error {
	*o = DjangoAdminLog{}
	return nil
}

func testDjangoAdminLogsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DjangoAdminLog{}
	o := &DjangoAdminLog{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, djangoAdminLogDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog object: %s", err)
	}

	AddDjangoAdminLogHook(boil.BeforeInsertHook, djangoAdminLogBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	djangoAdminLogBeforeInsertHooks = []DjangoAdminLogHook{}

	AddDjangoAdminLogHook(boil.AfterInsertHook, djangoAdminLogAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	djangoAdminLogAfterInsertHooks = []DjangoAdminLogHook{}

	AddDjangoAdminLogHook(boil.AfterSelectHook, djangoAdminLogAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	djangoAdminLogAfterSelectHooks = []DjangoAdminLogHook{}

	AddDjangoAdminLogHook(boil.BeforeUpdateHook, djangoAdminLogBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	djangoAdminLogBeforeUpdateHooks = []DjangoAdminLogHook{}

	AddDjangoAdminLogHook(boil.AfterUpdateHook, djangoAdminLogAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	djangoAdminLogAfterUpdateHooks = []DjangoAdminLogHook{}

	AddDjangoAdminLogHook(boil.BeforeDeleteHook, djangoAdminLogBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	djangoAdminLogBeforeDeleteHooks = []DjangoAdminLogHook{}

	AddDjangoAdminLogHook(boil.AfterDeleteHook, djangoAdminLogAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	djangoAdminLogAfterDeleteHooks = []DjangoAdminLogHook{}

	AddDjangoAdminLogHook(boil.BeforeUpsertHook, djangoAdminLogBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	djangoAdminLogBeforeUpsertHooks = []DjangoAdminLogHook{}

	AddDjangoAdminLogHook(boil.AfterUpsertHook, djangoAdminLogAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	djangoAdminLogAfterUpsertHooks = []DjangoAdminLogHook{}
}

func testDjangoAdminLogsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoAdminLog{}
	if err = randomize.Struct(seed, o, djangoAdminLogDBTypes, true, djangoAdminLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DjangoAdminLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDjangoAdminLogsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoAdminLog{}
	if err = randomize.Struct(seed, o, djangoAdminLogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(djangoAdminLogColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DjangoAdminLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDjangoAdminLogToOneDjangoContentTypeUsingContentType(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DjangoAdminLog
	var foreign DjangoContentType

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, djangoAdminLogDBTypes, true, djangoAdminLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, djangoContentTypeDBTypes, false, djangoContentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.ContentTypeID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ContentType().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DjangoAdminLogSlice{&local}
	if err = local.L.LoadContentType(ctx, tx, false, (*[]*DjangoAdminLog)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ContentType == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ContentType = nil
	if err = local.L.LoadContentType(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ContentType == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDjangoAdminLogToOneAuthUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DjangoAdminLog
	var foreign AuthUser

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, djangoAdminLogDBTypes, false, djangoAdminLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, authUserDBTypes, false, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DjangoAdminLogSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*DjangoAdminLog)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDjangoAdminLogToOneSetOpDjangoContentTypeUsingContentType(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DjangoAdminLog
	var b, c DjangoContentType

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, djangoAdminLogDBTypes, false, strmangle.SetComplement(djangoAdminLogPrimaryKeyColumns, djangoAdminLogColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, djangoContentTypeDBTypes, false, strmangle.SetComplement(djangoContentTypePrimaryKeyColumns, djangoContentTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, djangoContentTypeDBTypes, false, strmangle.SetComplement(djangoContentTypePrimaryKeyColumns, djangoContentTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*DjangoContentType{&b, &c} {
		err = a.SetContentType(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ContentType != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ContentTypeDjangoAdminLogs[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.ContentTypeID, x.ID) {
			t.Error("foreign key was wrong value", a.ContentTypeID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ContentTypeID))
		reflect.Indirect(reflect.ValueOf(&a.ContentTypeID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.ContentTypeID, x.ID) {
			t.Error("foreign key was wrong value", a.ContentTypeID, x.ID)
		}
	}
}

func testDjangoAdminLogToOneRemoveOpDjangoContentTypeUsingContentType(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DjangoAdminLog
	var b DjangoContentType

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, djangoAdminLogDBTypes, false, strmangle.SetComplement(djangoAdminLogPrimaryKeyColumns, djangoAdminLogColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, djangoContentTypeDBTypes, false, strmangle.SetComplement(djangoContentTypePrimaryKeyColumns, djangoContentTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetContentType(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveContentType(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.ContentType().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.ContentType != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.ContentTypeID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.ContentTypeDjangoAdminLogs) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testDjangoAdminLogToOneSetOpAuthUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DjangoAdminLog
	var b, c AuthUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, djangoAdminLogDBTypes, false, strmangle.SetComplement(djangoAdminLogPrimaryKeyColumns, djangoAdminLogColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, authUserDBTypes, false, strmangle.SetComplement(authUserPrimaryKeyColumns, authUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, authUserDBTypes, false, strmangle.SetComplement(authUserPrimaryKeyColumns, authUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*AuthUser{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserDjangoAdminLogs[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testDjangoAdminLogsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoAdminLog{}
	if err = randomize.Struct(seed, o, djangoAdminLogDBTypes, true, djangoAdminLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
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

func testDjangoAdminLogsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoAdminLog{}
	if err = randomize.Struct(seed, o, djangoAdminLogDBTypes, true, djangoAdminLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DjangoAdminLogSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDjangoAdminLogsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoAdminLog{}
	if err = randomize.Struct(seed, o, djangoAdminLogDBTypes, true, djangoAdminLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DjangoAdminLogs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	djangoAdminLogDBTypes = map[string]string{`ID`: `integer`, `ActionTime`: `timestamp with time zone`, `ObjectID`: `text`, `ObjectRepr`: `character varying`, `ActionFlag`: `smallint`, `ChangeMessage`: `text`, `ContentTypeID`: `integer`, `UserID`: `integer`}
	_                     = bytes.MinRead
)

func testDjangoAdminLogsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(djangoAdminLogPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(djangoAdminLogAllColumns) == len(djangoAdminLogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DjangoAdminLog{}
	if err = randomize.Struct(seed, o, djangoAdminLogDBTypes, true, djangoAdminLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DjangoAdminLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, djangoAdminLogDBTypes, true, djangoAdminLogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDjangoAdminLogsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(djangoAdminLogAllColumns) == len(djangoAdminLogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DjangoAdminLog{}
	if err = randomize.Struct(seed, o, djangoAdminLogDBTypes, true, djangoAdminLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DjangoAdminLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, djangoAdminLogDBTypes, true, djangoAdminLogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(djangoAdminLogAllColumns, djangoAdminLogPrimaryKeyColumns) {
		fields = djangoAdminLogAllColumns
	} else {
		fields = strmangle.SetComplement(
			djangoAdminLogAllColumns,
			djangoAdminLogPrimaryKeyColumns,
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

	slice := DjangoAdminLogSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDjangoAdminLogsUpsert(t *testing.T) {
	t.Parallel()

	if len(djangoAdminLogAllColumns) == len(djangoAdminLogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DjangoAdminLog{}
	if err = randomize.Struct(seed, &o, djangoAdminLogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DjangoAdminLog: %s", err)
	}

	count, err := DjangoAdminLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, djangoAdminLogDBTypes, false, djangoAdminLogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DjangoAdminLog struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DjangoAdminLog: %s", err)
	}

	count, err = DjangoAdminLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}