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

func testAuthPermissions(t *testing.T) {
	t.Parallel()

	query := AuthPermissions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAuthPermissionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
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

	count, err := AuthPermissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthPermissionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := AuthPermissions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AuthPermissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthPermissionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AuthPermissionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AuthPermissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthPermissionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AuthPermissionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if AuthPermission exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AuthPermissionExists to return true, but got false.")
	}
}

func testAuthPermissionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	authPermissionFound, err := FindAuthPermission(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if authPermissionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAuthPermissionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = AuthPermissions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAuthPermissionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := AuthPermissions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAuthPermissionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authPermissionOne := &AuthPermission{}
	authPermissionTwo := &AuthPermission{}
	if err = randomize.Struct(seed, authPermissionOne, authPermissionDBTypes, false, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}
	if err = randomize.Struct(seed, authPermissionTwo, authPermissionDBTypes, false, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = authPermissionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = authPermissionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AuthPermissions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAuthPermissionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	authPermissionOne := &AuthPermission{}
	authPermissionTwo := &AuthPermission{}
	if err = randomize.Struct(seed, authPermissionOne, authPermissionDBTypes, false, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}
	if err = randomize.Struct(seed, authPermissionTwo, authPermissionDBTypes, false, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = authPermissionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = authPermissionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AuthPermissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func authPermissionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *AuthPermission) error {
	*o = AuthPermission{}
	return nil
}

func authPermissionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *AuthPermission) error {
	*o = AuthPermission{}
	return nil
}

func authPermissionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *AuthPermission) error {
	*o = AuthPermission{}
	return nil
}

func authPermissionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AuthPermission) error {
	*o = AuthPermission{}
	return nil
}

func authPermissionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AuthPermission) error {
	*o = AuthPermission{}
	return nil
}

func authPermissionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AuthPermission) error {
	*o = AuthPermission{}
	return nil
}

func authPermissionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AuthPermission) error {
	*o = AuthPermission{}
	return nil
}

func authPermissionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AuthPermission) error {
	*o = AuthPermission{}
	return nil
}

func authPermissionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AuthPermission) error {
	*o = AuthPermission{}
	return nil
}

func testAuthPermissionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &AuthPermission{}
	o := &AuthPermission{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, authPermissionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AuthPermission object: %s", err)
	}

	AddAuthPermissionHook(boil.BeforeInsertHook, authPermissionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	authPermissionBeforeInsertHooks = []AuthPermissionHook{}

	AddAuthPermissionHook(boil.AfterInsertHook, authPermissionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	authPermissionAfterInsertHooks = []AuthPermissionHook{}

	AddAuthPermissionHook(boil.AfterSelectHook, authPermissionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	authPermissionAfterSelectHooks = []AuthPermissionHook{}

	AddAuthPermissionHook(boil.BeforeUpdateHook, authPermissionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	authPermissionBeforeUpdateHooks = []AuthPermissionHook{}

	AddAuthPermissionHook(boil.AfterUpdateHook, authPermissionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	authPermissionAfterUpdateHooks = []AuthPermissionHook{}

	AddAuthPermissionHook(boil.BeforeDeleteHook, authPermissionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	authPermissionBeforeDeleteHooks = []AuthPermissionHook{}

	AddAuthPermissionHook(boil.AfterDeleteHook, authPermissionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	authPermissionAfterDeleteHooks = []AuthPermissionHook{}

	AddAuthPermissionHook(boil.BeforeUpsertHook, authPermissionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	authPermissionBeforeUpsertHooks = []AuthPermissionHook{}

	AddAuthPermissionHook(boil.AfterUpsertHook, authPermissionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	authPermissionAfterUpsertHooks = []AuthPermissionHook{}
}

func testAuthPermissionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AuthPermissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthPermissionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(authPermissionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := AuthPermissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthPermissionToManyPermissionAuthGroupPermissions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AuthPermission
	var b, c AuthGroupPermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, authGroupPermissionDBTypes, false, authGroupPermissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, authGroupPermissionDBTypes, false, authGroupPermissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.PermissionID = a.ID
	c.PermissionID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.PermissionAuthGroupPermissions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.PermissionID == b.PermissionID {
			bFound = true
		}
		if v.PermissionID == c.PermissionID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AuthPermissionSlice{&a}
	if err = a.L.LoadPermissionAuthGroupPermissions(ctx, tx, false, (*[]*AuthPermission)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PermissionAuthGroupPermissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.PermissionAuthGroupPermissions = nil
	if err = a.L.LoadPermissionAuthGroupPermissions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PermissionAuthGroupPermissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testAuthPermissionToManyPermissionAuthUserUserPermissions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AuthPermission
	var b, c AuthUserUserPermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, authUserUserPermissionDBTypes, false, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, authUserUserPermissionDBTypes, false, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.PermissionID = a.ID
	c.PermissionID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.PermissionAuthUserUserPermissions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.PermissionID == b.PermissionID {
			bFound = true
		}
		if v.PermissionID == c.PermissionID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AuthPermissionSlice{&a}
	if err = a.L.LoadPermissionAuthUserUserPermissions(ctx, tx, false, (*[]*AuthPermission)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PermissionAuthUserUserPermissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.PermissionAuthUserUserPermissions = nil
	if err = a.L.LoadPermissionAuthUserUserPermissions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PermissionAuthUserUserPermissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testAuthPermissionToManyAddOpPermissionAuthGroupPermissions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AuthPermission
	var b, c, d, e AuthGroupPermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authPermissionDBTypes, false, strmangle.SetComplement(authPermissionPrimaryKeyColumns, authPermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AuthGroupPermission{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, authGroupPermissionDBTypes, false, strmangle.SetComplement(authGroupPermissionPrimaryKeyColumns, authGroupPermissionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*AuthGroupPermission{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPermissionAuthGroupPermissions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.PermissionID {
			t.Error("foreign key was wrong value", a.ID, first.PermissionID)
		}
		if a.ID != second.PermissionID {
			t.Error("foreign key was wrong value", a.ID, second.PermissionID)
		}

		if first.R.Permission != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Permission != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.PermissionAuthGroupPermissions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.PermissionAuthGroupPermissions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.PermissionAuthGroupPermissions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testAuthPermissionToManyAddOpPermissionAuthUserUserPermissions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AuthPermission
	var b, c, d, e AuthUserUserPermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authPermissionDBTypes, false, strmangle.SetComplement(authPermissionPrimaryKeyColumns, authPermissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AuthUserUserPermission{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, authUserUserPermissionDBTypes, false, strmangle.SetComplement(authUserUserPermissionPrimaryKeyColumns, authUserUserPermissionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*AuthUserUserPermission{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPermissionAuthUserUserPermissions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.PermissionID {
			t.Error("foreign key was wrong value", a.ID, first.PermissionID)
		}
		if a.ID != second.PermissionID {
			t.Error("foreign key was wrong value", a.ID, second.PermissionID)
		}

		if first.R.Permission != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Permission != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.PermissionAuthUserUserPermissions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.PermissionAuthUserUserPermissions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.PermissionAuthUserUserPermissions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testAuthPermissionToOneDjangoContentTypeUsingContentType(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local AuthPermission
	var foreign DjangoContentType

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, authPermissionDBTypes, false, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, djangoContentTypeDBTypes, false, djangoContentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ContentTypeID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ContentType().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := AuthPermissionSlice{&local}
	if err = local.L.LoadContentType(ctx, tx, false, (*[]*AuthPermission)(&slice), nil); err != nil {
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

func testAuthPermissionToOneSetOpDjangoContentTypeUsingContentType(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AuthPermission
	var b, c DjangoContentType

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authPermissionDBTypes, false, strmangle.SetComplement(authPermissionPrimaryKeyColumns, authPermissionColumnsWithoutDefault)...); err != nil {
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

		if x.R.ContentTypeAuthPermissions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ContentTypeID != x.ID {
			t.Error("foreign key was wrong value", a.ContentTypeID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ContentTypeID))
		reflect.Indirect(reflect.ValueOf(&a.ContentTypeID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ContentTypeID != x.ID {
			t.Error("foreign key was wrong value", a.ContentTypeID, x.ID)
		}
	}
}

func testAuthPermissionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
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

func testAuthPermissionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AuthPermissionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAuthPermissionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AuthPermissions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	authPermissionDBTypes = map[string]string{`ID`: `integer`, `Name`: `character varying`, `ContentTypeID`: `integer`, `Codename`: `character varying`}
	_                     = bytes.MinRead
)

func testAuthPermissionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(authPermissionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(authPermissionAllColumns) == len(authPermissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AuthPermissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAuthPermissionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(authPermissionAllColumns) == len(authPermissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AuthPermission{}
	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AuthPermissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, authPermissionDBTypes, true, authPermissionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(authPermissionAllColumns, authPermissionPrimaryKeyColumns) {
		fields = authPermissionAllColumns
	} else {
		fields = strmangle.SetComplement(
			authPermissionAllColumns,
			authPermissionPrimaryKeyColumns,
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

	slice := AuthPermissionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAuthPermissionsUpsert(t *testing.T) {
	t.Parallel()

	if len(authPermissionAllColumns) == len(authPermissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := AuthPermission{}
	if err = randomize.Struct(seed, &o, authPermissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AuthPermission: %s", err)
	}

	count, err := AuthPermissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, authPermissionDBTypes, false, authPermissionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthPermission struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AuthPermission: %s", err)
	}

	count, err = AuthPermissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
