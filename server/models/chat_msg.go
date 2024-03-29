// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// ChatMSG is an object representing the database table.
type ChatMSG struct {
	ID       string     `boil:"id" json:"id" toml:"id" yaml:"id"`
	HubID    string     `boil:"hub_id" json:"hub_id" toml:"hub_id" yaml:"hub_id"`
	GifData  null.Bytes `boil:"gif_data" json:"gif_data,omitempty" toml:"gif_data" yaml:"gif_data,omitempty"`
	CreateAt null.Time  `boil:"create_at" json:"create_at,omitempty" toml:"create_at" yaml:"create_at,omitempty"`
	UserID   string     `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`

	R *chatMSGR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L chatMSGL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ChatMSGColumns = struct {
	ID       string
	HubID    string
	GifData  string
	CreateAt string
	UserID   string
}{
	ID:       "id",
	HubID:    "hub_id",
	GifData:  "gif_data",
	CreateAt: "create_at",
	UserID:   "user_id",
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

type whereHelpernull_Bytes struct{ field string }

func (w whereHelpernull_Bytes) EQ(x null.Bytes) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Bytes) NEQ(x null.Bytes) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Bytes) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Bytes) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Bytes) LT(x null.Bytes) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Bytes) LTE(x null.Bytes) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Bytes) GT(x null.Bytes) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Bytes) GTE(x null.Bytes) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var ChatMSGWhere = struct {
	ID       whereHelperstring
	HubID    whereHelperstring
	GifData  whereHelpernull_Bytes
	CreateAt whereHelpernull_Time
	UserID   whereHelperstring
}{
	ID:       whereHelperstring{field: "\"chat_msg\".\"id\""},
	HubID:    whereHelperstring{field: "\"chat_msg\".\"hub_id\""},
	GifData:  whereHelpernull_Bytes{field: "\"chat_msg\".\"gif_data\""},
	CreateAt: whereHelpernull_Time{field: "\"chat_msg\".\"create_at\""},
	UserID:   whereHelperstring{field: "\"chat_msg\".\"user_id\""},
}

// ChatMSGRels is where relationship names are stored.
var ChatMSGRels = struct {
	Hub  string
	User string
}{
	Hub:  "Hub",
	User: "User",
}

// chatMSGR is where relationships are stored.
type chatMSGR struct {
	Hub  *Hub
	User *User
}

// NewStruct creates a new relationship struct
func (*chatMSGR) NewStruct() *chatMSGR {
	return &chatMSGR{}
}

// chatMSGL is where Load methods for each relationship are stored.
type chatMSGL struct{}

var (
	chatMSGAllColumns            = []string{"id", "hub_id", "gif_data", "create_at", "user_id"}
	chatMSGColumnsWithoutDefault = []string{"id", "hub_id", "gif_data", "create_at", "user_id"}
	chatMSGColumnsWithDefault    = []string{}
	chatMSGPrimaryKeyColumns     = []string{"id"}
)

type (
	// ChatMSGSlice is an alias for a slice of pointers to ChatMSG.
	// This should generally be used opposed to []ChatMSG.
	ChatMSGSlice []*ChatMSG
	// ChatMSGHook is the signature for custom ChatMSG hook methods
	ChatMSGHook func(context.Context, boil.ContextExecutor, *ChatMSG) error

	chatMSGQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	chatMSGType                 = reflect.TypeOf(&ChatMSG{})
	chatMSGMapping              = queries.MakeStructMapping(chatMSGType)
	chatMSGPrimaryKeyMapping, _ = queries.BindMapping(chatMSGType, chatMSGMapping, chatMSGPrimaryKeyColumns)
	chatMSGInsertCacheMut       sync.RWMutex
	chatMSGInsertCache          = make(map[string]insertCache)
	chatMSGUpdateCacheMut       sync.RWMutex
	chatMSGUpdateCache          = make(map[string]updateCache)
	chatMSGUpsertCacheMut       sync.RWMutex
	chatMSGUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var chatMSGBeforeInsertHooks []ChatMSGHook
var chatMSGBeforeUpdateHooks []ChatMSGHook
var chatMSGBeforeDeleteHooks []ChatMSGHook
var chatMSGBeforeUpsertHooks []ChatMSGHook

var chatMSGAfterInsertHooks []ChatMSGHook
var chatMSGAfterSelectHooks []ChatMSGHook
var chatMSGAfterUpdateHooks []ChatMSGHook
var chatMSGAfterDeleteHooks []ChatMSGHook
var chatMSGAfterUpsertHooks []ChatMSGHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ChatMSG) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatMSGBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ChatMSG) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatMSGBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ChatMSG) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatMSGBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ChatMSG) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatMSGBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ChatMSG) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatMSGAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ChatMSG) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatMSGAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ChatMSG) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatMSGAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ChatMSG) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatMSGAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ChatMSG) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chatMSGAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddChatMSGHook registers your hook function for all future operations.
func AddChatMSGHook(hookPoint boil.HookPoint, chatMSGHook ChatMSGHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		chatMSGBeforeInsertHooks = append(chatMSGBeforeInsertHooks, chatMSGHook)
	case boil.BeforeUpdateHook:
		chatMSGBeforeUpdateHooks = append(chatMSGBeforeUpdateHooks, chatMSGHook)
	case boil.BeforeDeleteHook:
		chatMSGBeforeDeleteHooks = append(chatMSGBeforeDeleteHooks, chatMSGHook)
	case boil.BeforeUpsertHook:
		chatMSGBeforeUpsertHooks = append(chatMSGBeforeUpsertHooks, chatMSGHook)
	case boil.AfterInsertHook:
		chatMSGAfterInsertHooks = append(chatMSGAfterInsertHooks, chatMSGHook)
	case boil.AfterSelectHook:
		chatMSGAfterSelectHooks = append(chatMSGAfterSelectHooks, chatMSGHook)
	case boil.AfterUpdateHook:
		chatMSGAfterUpdateHooks = append(chatMSGAfterUpdateHooks, chatMSGHook)
	case boil.AfterDeleteHook:
		chatMSGAfterDeleteHooks = append(chatMSGAfterDeleteHooks, chatMSGHook)
	case boil.AfterUpsertHook:
		chatMSGAfterUpsertHooks = append(chatMSGAfterUpsertHooks, chatMSGHook)
	}
}

// One returns a single chatMSG record from the query.
func (q chatMSGQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ChatMSG, error) {
	o := &ChatMSG{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for chat_msg")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ChatMSG records from the query.
func (q chatMSGQuery) All(ctx context.Context, exec boil.ContextExecutor) (ChatMSGSlice, error) {
	var o []*ChatMSG

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ChatMSG slice")
	}

	if len(chatMSGAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ChatMSG records in the query.
func (q chatMSGQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count chat_msg rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q chatMSGQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if chat_msg exists")
	}

	return count > 0, nil
}

// Hub pointed to by the foreign key.
func (o *ChatMSG) Hub(mods ...qm.QueryMod) hubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.HubID),
	}

	queryMods = append(queryMods, mods...)

	query := Hubs(queryMods...)
	queries.SetFrom(query.Query, "\"hubs\"")

	return query
}

// User pointed to by the foreign key.
func (o *ChatMSG) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// LoadHub allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (chatMSGL) LoadHub(ctx context.Context, e boil.ContextExecutor, singular bool, maybeChatMSG interface{}, mods queries.Applicator) error {
	var slice []*ChatMSG
	var object *ChatMSG

	if singular {
		object = maybeChatMSG.(*ChatMSG)
	} else {
		slice = *maybeChatMSG.(*[]*ChatMSG)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &chatMSGR{}
		}
		args = append(args, object.HubID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &chatMSGR{}
			}

			for _, a := range args {
				if a == obj.HubID {
					continue Outer
				}
			}

			args = append(args, obj.HubID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`hubs`), qm.WhereIn(`hubs.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Hub")
	}

	var resultSlice []*Hub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Hub")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for hubs")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for hubs")
	}

	if len(chatMSGAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Hub = foreign
		if foreign.R == nil {
			foreign.R = &hubR{}
		}
		foreign.R.ChatMSGS = append(foreign.R.ChatMSGS, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.HubID == foreign.ID {
				local.R.Hub = foreign
				if foreign.R == nil {
					foreign.R = &hubR{}
				}
				foreign.R.ChatMSGS = append(foreign.R.ChatMSGS, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (chatMSGL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeChatMSG interface{}, mods queries.Applicator) error {
	var slice []*ChatMSG
	var object *ChatMSG

	if singular {
		object = maybeChatMSG.(*ChatMSG)
	} else {
		slice = *maybeChatMSG.(*[]*ChatMSG)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &chatMSGR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &chatMSGR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`users`), qm.WhereIn(`users.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(chatMSGAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.ChatMSGS = append(foreign.R.ChatMSGS, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.ChatMSGS = append(foreign.R.ChatMSGS, local)
				break
			}
		}
	}

	return nil
}

// SetHub of the chatMSG to the related item.
// Sets o.R.Hub to related.
// Adds o to related.R.ChatMSGS.
func (o *ChatMSG) SetHub(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Hub) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"chat_msg\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"hub_id"}),
		strmangle.WhereClause("\"", "\"", 2, chatMSGPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.HubID = related.ID
	if o.R == nil {
		o.R = &chatMSGR{
			Hub: related,
		}
	} else {
		o.R.Hub = related
	}

	if related.R == nil {
		related.R = &hubR{
			ChatMSGS: ChatMSGSlice{o},
		}
	} else {
		related.R.ChatMSGS = append(related.R.ChatMSGS, o)
	}

	return nil
}

// SetUser of the chatMSG to the related item.
// Sets o.R.User to related.
// Adds o to related.R.ChatMSGS.
func (o *ChatMSG) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"chat_msg\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, chatMSGPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &chatMSGR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			ChatMSGS: ChatMSGSlice{o},
		}
	} else {
		related.R.ChatMSGS = append(related.R.ChatMSGS, o)
	}

	return nil
}

// ChatMSGS retrieves all the records using an executor.
func ChatMSGS(mods ...qm.QueryMod) chatMSGQuery {
	mods = append(mods, qm.From("\"chat_msg\""))
	return chatMSGQuery{NewQuery(mods...)}
}

// FindChatMSG retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindChatMSG(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*ChatMSG, error) {
	chatMSGObj := &ChatMSG{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"chat_msg\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, chatMSGObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from chat_msg")
	}

	return chatMSGObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ChatMSG) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no chat_msg provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(chatMSGColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	chatMSGInsertCacheMut.RLock()
	cache, cached := chatMSGInsertCache[key]
	chatMSGInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			chatMSGAllColumns,
			chatMSGColumnsWithDefault,
			chatMSGColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(chatMSGType, chatMSGMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(chatMSGType, chatMSGMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"chat_msg\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"chat_msg\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into chat_msg")
	}

	if !cached {
		chatMSGInsertCacheMut.Lock()
		chatMSGInsertCache[key] = cache
		chatMSGInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ChatMSG.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ChatMSG) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	chatMSGUpdateCacheMut.RLock()
	cache, cached := chatMSGUpdateCache[key]
	chatMSGUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			chatMSGAllColumns,
			chatMSGPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update chat_msg, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"chat_msg\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, chatMSGPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(chatMSGType, chatMSGMapping, append(wl, chatMSGPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update chat_msg row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for chat_msg")
	}

	if !cached {
		chatMSGUpdateCacheMut.Lock()
		chatMSGUpdateCache[key] = cache
		chatMSGUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q chatMSGQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for chat_msg")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for chat_msg")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ChatMSGSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), chatMSGPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"chat_msg\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, chatMSGPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in chatMSG slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all chatMSG")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ChatMSG) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no chat_msg provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(chatMSGColumnsWithDefault, o)

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

	chatMSGUpsertCacheMut.RLock()
	cache, cached := chatMSGUpsertCache[key]
	chatMSGUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			chatMSGAllColumns,
			chatMSGColumnsWithDefault,
			chatMSGColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			chatMSGAllColumns,
			chatMSGPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert chat_msg, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(chatMSGPrimaryKeyColumns))
			copy(conflict, chatMSGPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"chat_msg\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(chatMSGType, chatMSGMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(chatMSGType, chatMSGMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert chat_msg")
	}

	if !cached {
		chatMSGUpsertCacheMut.Lock()
		chatMSGUpsertCache[key] = cache
		chatMSGUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ChatMSG record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ChatMSG) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ChatMSG provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), chatMSGPrimaryKeyMapping)
	sql := "DELETE FROM \"chat_msg\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from chat_msg")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for chat_msg")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q chatMSGQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no chatMSGQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from chat_msg")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for chat_msg")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ChatMSGSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(chatMSGBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), chatMSGPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"chat_msg\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, chatMSGPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from chatMSG slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for chat_msg")
	}

	if len(chatMSGAfterDeleteHooks) != 0 {
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
func (o *ChatMSG) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindChatMSG(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChatMSGSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ChatMSGSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), chatMSGPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"chat_msg\".* FROM \"chat_msg\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, chatMSGPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ChatMSGSlice")
	}

	*o = slice

	return nil
}

// ChatMSGExists checks if the ChatMSG row exists.
func ChatMSGExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"chat_msg\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if chat_msg exists")
	}

	return exists, nil
}
