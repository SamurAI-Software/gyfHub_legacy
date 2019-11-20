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
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Gif is an object representing the database table.
type Gif struct {
	ID      string `boil:"id" json:"id" toml:"id" yaml:"id"`
	GifData []byte `boil:"gif_data" json:"gif_data" toml:"gif_data" yaml:"gif_data"`

	R *gifR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L gifL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GifColumns = struct {
	ID      string
	GifData string
}{
	ID:      "id",
	GifData: "gif_data",
}

// Generated where

type whereHelper__byte struct{ field string }

func (w whereHelper__byte) EQ(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelper__byte) NEQ(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelper__byte) LT(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelper__byte) LTE(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelper__byte) GT(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelper__byte) GTE(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var GifWhere = struct {
	ID      whereHelperstring
	GifData whereHelper__byte
}{
	ID:      whereHelperstring{field: "\"gifs\".\"id\""},
	GifData: whereHelper__byte{field: "\"gifs\".\"gif_data\""},
}

// GifRels is where relationship names are stored.
var GifRels = struct {
	UserGif       string
	ChatMSGS      string
	UserFavorites string
}{
	UserGif:       "UserGif",
	ChatMSGS:      "ChatMSGS",
	UserFavorites: "UserFavorites",
}

// gifR is where relationships are stored.
type gifR struct {
	UserGif       *UserGif
	ChatMSGS      ChatMSGSlice
	UserFavorites UserFavoriteSlice
}

// NewStruct creates a new relationship struct
func (*gifR) NewStruct() *gifR {
	return &gifR{}
}

// gifL is where Load methods for each relationship are stored.
type gifL struct{}

var (
	gifAllColumns            = []string{"id", "gif_data"}
	gifColumnsWithoutDefault = []string{"id", "gif_data"}
	gifColumnsWithDefault    = []string{}
	gifPrimaryKeyColumns     = []string{"id"}
)

type (
	// GifSlice is an alias for a slice of pointers to Gif.
	// This should generally be used opposed to []Gif.
	GifSlice []*Gif
	// GifHook is the signature for custom Gif hook methods
	GifHook func(context.Context, boil.ContextExecutor, *Gif) error

	gifQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	gifType                 = reflect.TypeOf(&Gif{})
	gifMapping              = queries.MakeStructMapping(gifType)
	gifPrimaryKeyMapping, _ = queries.BindMapping(gifType, gifMapping, gifPrimaryKeyColumns)
	gifInsertCacheMut       sync.RWMutex
	gifInsertCache          = make(map[string]insertCache)
	gifUpdateCacheMut       sync.RWMutex
	gifUpdateCache          = make(map[string]updateCache)
	gifUpsertCacheMut       sync.RWMutex
	gifUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var gifBeforeInsertHooks []GifHook
var gifBeforeUpdateHooks []GifHook
var gifBeforeDeleteHooks []GifHook
var gifBeforeUpsertHooks []GifHook

var gifAfterInsertHooks []GifHook
var gifAfterSelectHooks []GifHook
var gifAfterUpdateHooks []GifHook
var gifAfterDeleteHooks []GifHook
var gifAfterUpsertHooks []GifHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Gif) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gifBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Gif) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gifBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Gif) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gifBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Gif) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gifBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Gif) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gifAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Gif) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gifAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Gif) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gifAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Gif) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gifAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Gif) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gifAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGifHook registers your hook function for all future operations.
func AddGifHook(hookPoint boil.HookPoint, gifHook GifHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		gifBeforeInsertHooks = append(gifBeforeInsertHooks, gifHook)
	case boil.BeforeUpdateHook:
		gifBeforeUpdateHooks = append(gifBeforeUpdateHooks, gifHook)
	case boil.BeforeDeleteHook:
		gifBeforeDeleteHooks = append(gifBeforeDeleteHooks, gifHook)
	case boil.BeforeUpsertHook:
		gifBeforeUpsertHooks = append(gifBeforeUpsertHooks, gifHook)
	case boil.AfterInsertHook:
		gifAfterInsertHooks = append(gifAfterInsertHooks, gifHook)
	case boil.AfterSelectHook:
		gifAfterSelectHooks = append(gifAfterSelectHooks, gifHook)
	case boil.AfterUpdateHook:
		gifAfterUpdateHooks = append(gifAfterUpdateHooks, gifHook)
	case boil.AfterDeleteHook:
		gifAfterDeleteHooks = append(gifAfterDeleteHooks, gifHook)
	case boil.AfterUpsertHook:
		gifAfterUpsertHooks = append(gifAfterUpsertHooks, gifHook)
	}
}

// One returns a single gif record from the query.
func (q gifQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Gif, error) {
	o := &Gif{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for gifs")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Gif records from the query.
func (q gifQuery) All(ctx context.Context, exec boil.ContextExecutor) (GifSlice, error) {
	var o []*Gif

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Gif slice")
	}

	if len(gifAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Gif records in the query.
func (q gifQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count gifs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q gifQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if gifs exists")
	}

	return count > 0, nil
}

// UserGif pointed to by the foreign key.
func (o *Gif) UserGif(mods ...qm.QueryMod) userGifQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"gif_id\" = ?", o.ID),
	}

	queryMods = append(queryMods, mods...)

	query := UserGifs(queryMods...)
	queries.SetFrom(query.Query, "\"user_gif\"")

	return query
}

// ChatMSGS retrieves all the chat_msg's ChatMSGS with an executor.
func (o *Gif) ChatMSGS(mods ...qm.QueryMod) chatMSGQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"chat_msg\".\"gif_id\"=?", o.ID),
	)

	query := ChatMSGS(queryMods...)
	queries.SetFrom(query.Query, "\"chat_msg\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"chat_msg\".*"})
	}

	return query
}

// UserFavorites retrieves all the user_favorite's UserFavorites with an executor.
func (o *Gif) UserFavorites(mods ...qm.QueryMod) userFavoriteQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"user_favorite\".\"gif_id\"=?", o.ID),
	)

	query := UserFavorites(queryMods...)
	queries.SetFrom(query.Query, "\"user_favorite\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"user_favorite\".*"})
	}

	return query
}

// LoadUserGif allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-1 relationship.
func (gifL) LoadUserGif(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGif interface{}, mods queries.Applicator) error {
	var slice []*Gif
	var object *Gif

	if singular {
		object = maybeGif.(*Gif)
	} else {
		slice = *maybeGif.(*[]*Gif)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &gifR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &gifR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`user_gif`), qm.WhereIn(`user_gif.gif_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load UserGif")
	}

	var resultSlice []*UserGif
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice UserGif")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for user_gif")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_gif")
	}

	if len(gifAfterSelectHooks) != 0 {
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
		object.R.UserGif = foreign
		if foreign.R == nil {
			foreign.R = &userGifR{}
		}
		foreign.R.Gif = object
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ID == foreign.GifID {
				local.R.UserGif = foreign
				if foreign.R == nil {
					foreign.R = &userGifR{}
				}
				foreign.R.Gif = local
				break
			}
		}
	}

	return nil
}

// LoadChatMSGS allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (gifL) LoadChatMSGS(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGif interface{}, mods queries.Applicator) error {
	var slice []*Gif
	var object *Gif

	if singular {
		object = maybeGif.(*Gif)
	} else {
		slice = *maybeGif.(*[]*Gif)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &gifR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &gifR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`chat_msg`), qm.WhereIn(`chat_msg.gif_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load chat_msg")
	}

	var resultSlice []*ChatMSG
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice chat_msg")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on chat_msg")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for chat_msg")
	}

	if len(chatMSGAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ChatMSGS = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &chatMSGR{}
			}
			foreign.R.Gif = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.GifID {
				local.R.ChatMSGS = append(local.R.ChatMSGS, foreign)
				if foreign.R == nil {
					foreign.R = &chatMSGR{}
				}
				foreign.R.Gif = local
				break
			}
		}
	}

	return nil
}

// LoadUserFavorites allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (gifL) LoadUserFavorites(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGif interface{}, mods queries.Applicator) error {
	var slice []*Gif
	var object *Gif

	if singular {
		object = maybeGif.(*Gif)
	} else {
		slice = *maybeGif.(*[]*Gif)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &gifR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &gifR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`user_favorite`), qm.WhereIn(`user_favorite.gif_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load user_favorite")
	}

	var resultSlice []*UserFavorite
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice user_favorite")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on user_favorite")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_favorite")
	}

	if len(userFavoriteAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.UserFavorites = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userFavoriteR{}
			}
			foreign.R.Gif = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.GifID {
				local.R.UserFavorites = append(local.R.UserFavorites, foreign)
				if foreign.R == nil {
					foreign.R = &userFavoriteR{}
				}
				foreign.R.Gif = local
				break
			}
		}
	}

	return nil
}

// SetUserGif of the gif to the related item.
// Sets o.R.UserGif to related.
// Adds o to related.R.Gif.
func (o *Gif) SetUserGif(ctx context.Context, exec boil.ContextExecutor, insert bool, related *UserGif) error {
	var err error

	if insert {
		related.GifID = o.ID

		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"user_gif\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"gif_id"}),
			strmangle.WhereClause("\"", "\"", 2, userGifPrimaryKeyColumns),
		)
		values := []interface{}{o.ID, related.ID}

		if boil.IsDebug(ctx) {
			writer := boil.DebugWriterFrom(ctx)
			fmt.Fprintln(writer, updateQuery)
			fmt.Fprintln(writer, values)
		}
		if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.GifID = o.ID

	}

	if o.R == nil {
		o.R = &gifR{
			UserGif: related,
		}
	} else {
		o.R.UserGif = related
	}

	if related.R == nil {
		related.R = &userGifR{
			Gif: o,
		}
	} else {
		related.R.Gif = o
	}
	return nil
}

// AddChatMSGS adds the given related objects to the existing relationships
// of the gif, optionally inserting them as new records.
// Appends related to o.R.ChatMSGS.
// Sets related.R.Gif appropriately.
func (o *Gif) AddChatMSGS(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*ChatMSG) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.GifID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"chat_msg\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"gif_id"}),
				strmangle.WhereClause("\"", "\"", 2, chatMSGPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.GifID = o.ID
		}
	}

	if o.R == nil {
		o.R = &gifR{
			ChatMSGS: related,
		}
	} else {
		o.R.ChatMSGS = append(o.R.ChatMSGS, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &chatMSGR{
				Gif: o,
			}
		} else {
			rel.R.Gif = o
		}
	}
	return nil
}

// AddUserFavorites adds the given related objects to the existing relationships
// of the gif, optionally inserting them as new records.
// Appends related to o.R.UserFavorites.
// Sets related.R.Gif appropriately.
func (o *Gif) AddUserFavorites(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*UserFavorite) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.GifID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"user_favorite\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"gif_id"}),
				strmangle.WhereClause("\"", "\"", 2, userFavoritePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.GifID = o.ID
		}
	}

	if o.R == nil {
		o.R = &gifR{
			UserFavorites: related,
		}
	} else {
		o.R.UserFavorites = append(o.R.UserFavorites, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userFavoriteR{
				Gif: o,
			}
		} else {
			rel.R.Gif = o
		}
	}
	return nil
}

// Gifs retrieves all the records using an executor.
func Gifs(mods ...qm.QueryMod) gifQuery {
	mods = append(mods, qm.From("\"gifs\""))
	return gifQuery{NewQuery(mods...)}
}

// FindGif retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGif(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Gif, error) {
	gifObj := &Gif{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"gifs\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, gifObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from gifs")
	}

	return gifObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Gif) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no gifs provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(gifColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	gifInsertCacheMut.RLock()
	cache, cached := gifInsertCache[key]
	gifInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			gifAllColumns,
			gifColumnsWithDefault,
			gifColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(gifType, gifMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(gifType, gifMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"gifs\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"gifs\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into gifs")
	}

	if !cached {
		gifInsertCacheMut.Lock()
		gifInsertCache[key] = cache
		gifInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Gif.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Gif) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	gifUpdateCacheMut.RLock()
	cache, cached := gifUpdateCache[key]
	gifUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			gifAllColumns,
			gifPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update gifs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"gifs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, gifPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(gifType, gifMapping, append(wl, gifPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update gifs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for gifs")
	}

	if !cached {
		gifUpdateCacheMut.Lock()
		gifUpdateCache[key] = cache
		gifUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q gifQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for gifs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for gifs")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GifSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gifPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"gifs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, gifPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in gif slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all gif")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Gif) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no gifs provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(gifColumnsWithDefault, o)

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

	gifUpsertCacheMut.RLock()
	cache, cached := gifUpsertCache[key]
	gifUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			gifAllColumns,
			gifColumnsWithDefault,
			gifColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			gifAllColumns,
			gifPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert gifs, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(gifPrimaryKeyColumns))
			copy(conflict, gifPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"gifs\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(gifType, gifMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(gifType, gifMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert gifs")
	}

	if !cached {
		gifUpsertCacheMut.Lock()
		gifUpsertCache[key] = cache
		gifUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Gif record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Gif) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Gif provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), gifPrimaryKeyMapping)
	sql := "DELETE FROM \"gifs\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from gifs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for gifs")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q gifQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no gifQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from gifs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for gifs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GifSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(gifBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gifPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"gifs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, gifPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from gif slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for gifs")
	}

	if len(gifAfterDeleteHooks) != 0 {
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
func (o *Gif) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGif(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GifSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GifSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gifPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"gifs\".* FROM \"gifs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, gifPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in GifSlice")
	}

	*o = slice

	return nil
}

// GifExists checks if the Gif row exists.
func GifExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"gifs\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if gifs exists")
	}

	return exists, nil
}
