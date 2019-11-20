// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testHubs(t *testing.T) {
	t.Parallel()

	query := Hubs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testHubsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hub{}
	if err = randomize.Struct(seed, o, hubDBTypes, true, hubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
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

	count, err := Hubs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHubsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hub{}
	if err = randomize.Struct(seed, o, hubDBTypes, true, hubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Hubs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Hubs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHubsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hub{}
	if err = randomize.Struct(seed, o, hubDBTypes, true, hubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := HubSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Hubs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHubsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hub{}
	if err = randomize.Struct(seed, o, hubDBTypes, true, hubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := HubExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Hub exists: %s", err)
	}
	if !e {
		t.Errorf("Expected HubExists to return true, but got false.")
	}
}

func testHubsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hub{}
	if err = randomize.Struct(seed, o, hubDBTypes, true, hubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	hubFound, err := FindHub(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if hubFound == nil {
		t.Error("want a record, got nil")
	}
}

func testHubsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hub{}
	if err = randomize.Struct(seed, o, hubDBTypes, true, hubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Hubs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testHubsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hub{}
	if err = randomize.Struct(seed, o, hubDBTypes, true, hubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Hubs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testHubsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	hubOne := &Hub{}
	hubTwo := &Hub{}
	if err = randomize.Struct(seed, hubOne, hubDBTypes, false, hubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}
	if err = randomize.Struct(seed, hubTwo, hubDBTypes, false, hubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = hubOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = hubTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Hubs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testHubsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	hubOne := &Hub{}
	hubTwo := &Hub{}
	if err = randomize.Struct(seed, hubOne, hubDBTypes, false, hubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}
	if err = randomize.Struct(seed, hubTwo, hubDBTypes, false, hubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = hubOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = hubTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Hubs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func hubBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Hub) error {
	*o = Hub{}
	return nil
}

func hubAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Hub) error {
	*o = Hub{}
	return nil
}

func hubAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Hub) error {
	*o = Hub{}
	return nil
}

func hubBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Hub) error {
	*o = Hub{}
	return nil
}

func hubAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Hub) error {
	*o = Hub{}
	return nil
}

func hubBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Hub) error {
	*o = Hub{}
	return nil
}

func hubAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Hub) error {
	*o = Hub{}
	return nil
}

func hubBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Hub) error {
	*o = Hub{}
	return nil
}

func hubAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Hub) error {
	*o = Hub{}
	return nil
}

func testHubsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Hub{}
	o := &Hub{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, hubDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Hub object: %s", err)
	}

	AddHubHook(boil.BeforeInsertHook, hubBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	hubBeforeInsertHooks = []HubHook{}

	AddHubHook(boil.AfterInsertHook, hubAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	hubAfterInsertHooks = []HubHook{}

	AddHubHook(boil.AfterSelectHook, hubAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	hubAfterSelectHooks = []HubHook{}

	AddHubHook(boil.BeforeUpdateHook, hubBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	hubBeforeUpdateHooks = []HubHook{}

	AddHubHook(boil.AfterUpdateHook, hubAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	hubAfterUpdateHooks = []HubHook{}

	AddHubHook(boil.BeforeDeleteHook, hubBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	hubBeforeDeleteHooks = []HubHook{}

	AddHubHook(boil.AfterDeleteHook, hubAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	hubAfterDeleteHooks = []HubHook{}

	AddHubHook(boil.BeforeUpsertHook, hubBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	hubBeforeUpsertHooks = []HubHook{}

	AddHubHook(boil.AfterUpsertHook, hubAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	hubAfterUpsertHooks = []HubHook{}
}

func testHubsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hub{}
	if err = randomize.Struct(seed, o, hubDBTypes, true, hubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Hubs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHubsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hub{}
	if err = randomize.Struct(seed, o, hubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(hubColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Hubs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHubToManyChatMSGS(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Hub
	var b, c ChatMSG

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, hubDBTypes, true, hubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, chatMSGDBTypes, false, chatMSGColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, chatMSGDBTypes, false, chatMSGColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.HubID = a.ID
	c.HubID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ChatMSGS().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.HubID == b.HubID {
			bFound = true
		}
		if v.HubID == c.HubID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := HubSlice{&a}
	if err = a.L.LoadChatMSGS(ctx, tx, false, (*[]*Hub)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ChatMSGS); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ChatMSGS = nil
	if err = a.L.LoadChatMSGS(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ChatMSGS); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testHubToManyHubUsers(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Hub
	var b, c HubUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, hubDBTypes, true, hubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, hubUserDBTypes, false, hubUserColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, hubUserDBTypes, false, hubUserColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.HubID = a.ID
	c.HubID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.HubUsers().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.HubID == b.HubID {
			bFound = true
		}
		if v.HubID == c.HubID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := HubSlice{&a}
	if err = a.L.LoadHubUsers(ctx, tx, false, (*[]*Hub)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.HubUsers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.HubUsers = nil
	if err = a.L.LoadHubUsers(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.HubUsers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testHubToManyAddOpChatMSGS(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Hub
	var b, c, d, e ChatMSG

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, hubDBTypes, false, strmangle.SetComplement(hubPrimaryKeyColumns, hubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ChatMSG{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, chatMSGDBTypes, false, strmangle.SetComplement(chatMSGPrimaryKeyColumns, chatMSGColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*ChatMSG{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddChatMSGS(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.HubID {
			t.Error("foreign key was wrong value", a.ID, first.HubID)
		}
		if a.ID != second.HubID {
			t.Error("foreign key was wrong value", a.ID, second.HubID)
		}

		if first.R.Hub != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Hub != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ChatMSGS[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ChatMSGS[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ChatMSGS().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testHubToManyAddOpHubUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Hub
	var b, c, d, e HubUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, hubDBTypes, false, strmangle.SetComplement(hubPrimaryKeyColumns, hubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*HubUser{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, hubUserDBTypes, false, strmangle.SetComplement(hubUserPrimaryKeyColumns, hubUserColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*HubUser{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddHubUsers(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.HubID {
			t.Error("foreign key was wrong value", a.ID, first.HubID)
		}
		if a.ID != second.HubID {
			t.Error("foreign key was wrong value", a.ID, second.HubID)
		}

		if first.R.Hub != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Hub != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.HubUsers[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.HubUsers[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.HubUsers().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testHubToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Hub
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, hubDBTypes, false, hubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
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

	slice := HubSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*Hub)(&slice), nil); err != nil {
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

func testHubToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Hub
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, hubDBTypes, false, strmangle.SetComplement(hubPrimaryKeyColumns, hubColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Hubs[0] != &a {
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

func testHubsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hub{}
	if err = randomize.Struct(seed, o, hubDBTypes, true, hubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
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

func testHubsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hub{}
	if err = randomize.Struct(seed, o, hubDBTypes, true, hubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := HubSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testHubsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Hub{}
	if err = randomize.Struct(seed, o, hubDBTypes, true, hubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Hubs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	hubDBTypes = map[string]string{`ID`: `text`, `Logo`: `bytea`, `UserID`: `text`, `IsPrivate`: `boolean`, `IsClose`: `boolean`}
	_          = bytes.MinRead
)

func testHubsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(hubPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(hubAllColumns) == len(hubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Hub{}
	if err = randomize.Struct(seed, o, hubDBTypes, true, hubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Hubs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, hubDBTypes, true, hubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testHubsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(hubAllColumns) == len(hubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Hub{}
	if err = randomize.Struct(seed, o, hubDBTypes, true, hubColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Hubs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, hubDBTypes, true, hubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(hubAllColumns, hubPrimaryKeyColumns) {
		fields = hubAllColumns
	} else {
		fields = strmangle.SetComplement(
			hubAllColumns,
			hubPrimaryKeyColumns,
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

	slice := HubSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testHubsUpsert(t *testing.T) {
	t.Parallel()

	if len(hubAllColumns) == len(hubPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Hub{}
	if err = randomize.Struct(seed, &o, hubDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Hub: %s", err)
	}

	count, err := Hubs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, hubDBTypes, false, hubPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Hub struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Hub: %s", err)
	}

	count, err = Hubs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}