package model

import (
	"bytes"
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/golang/protobuf/ptypes"
	"github.com/lib/pq"
	"github.com/piotrkowalczuk/ntypes"
	"github.com/piotrkowalczuk/qtypes"
)

const (
	TableUser                              = "charon.user"
	TableUserColumnConfirmationToken       = "confirmation_token"
	TableUserColumnCreatedAt               = "created_at"
	TableUserColumnCreatedBy               = "created_by"
	TableUserColumnFirstName               = "first_name"
	TableUserColumnID                      = "id"
	TableUserColumnIsActive                = "is_active"
	TableUserColumnIsConfirmed             = "is_confirmed"
	TableUserColumnIsStaff                 = "is_staff"
	TableUserColumnIsSuperuser             = "is_superuser"
	TableUserColumnLastLoginAt             = "last_login_at"
	TableUserColumnLastName                = "last_name"
	TableUserColumnPassword                = "password"
	TableUserColumnUpdatedAt               = "updated_at"
	TableUserColumnUpdatedBy               = "updated_by"
	TableUserColumnUsername                = "username"
	TableUserConstraintCreatedByForeignKey = "charon.user_created_by_fkey"

	TableUserConstraintPrimaryKey = "charon.user_id_pkey"

	TableUserConstraintUpdatedByForeignKey = "charon.user_updated_by_fkey"

	TableUserConstraintUsernameUnique = "charon.user_username_key"
)

var (
	TableUserColumns = []string{
		TableUserColumnConfirmationToken,
		TableUserColumnCreatedAt,
		TableUserColumnCreatedBy,
		TableUserColumnFirstName,
		TableUserColumnID,
		TableUserColumnIsActive,
		TableUserColumnIsConfirmed,
		TableUserColumnIsStaff,
		TableUserColumnIsSuperuser,
		TableUserColumnLastLoginAt,
		TableUserColumnLastName,
		TableUserColumnPassword,
		TableUserColumnUpdatedAt,
		TableUserColumnUpdatedBy,
		TableUserColumnUsername,
	}
)

// UserEntity ...
type UserEntity struct {
	// ConfirmationToken ...
	ConfirmationToken []byte
	// CreatedAt ...
	CreatedAt time.Time
	// CreatedBy ...
	CreatedBy ntypes.Int64
	// FirstName ...
	FirstName string
	// ID ...
	ID int64
	// IsActive ...
	IsActive bool
	// IsConfirmed ...
	IsConfirmed bool
	// IsStaff ...
	IsStaff bool
	// IsSuperuser ...
	IsSuperuser bool
	// LastLoginAt ...
	LastLoginAt pq.NullTime
	// LastName ...
	LastName string
	// Password ...
	Password []byte
	// UpdatedAt ...
	UpdatedAt pq.NullTime
	// UpdatedBy ...
	UpdatedBy ntypes.Int64
	// Username ...
	Username string
	// Author ...
	Author *UserEntity
	// Modifier ...
	Modifier *UserEntity
	// Permissions ...
	Permissions []*PermissionEntity
	// Groups ...
	Groups []*GroupEntity
}

func (e *UserEntity) Prop(cn string) (interface{}, bool) {
	switch cn {
	case TableUserColumnConfirmationToken:
		return &e.ConfirmationToken, true
	case TableUserColumnCreatedAt:
		return &e.CreatedAt, true
	case TableUserColumnCreatedBy:
		return &e.CreatedBy, true
	case TableUserColumnFirstName:
		return &e.FirstName, true
	case TableUserColumnID:
		return &e.ID, true
	case TableUserColumnIsActive:
		return &e.IsActive, true
	case TableUserColumnIsConfirmed:
		return &e.IsConfirmed, true
	case TableUserColumnIsStaff:
		return &e.IsStaff, true
	case TableUserColumnIsSuperuser:
		return &e.IsSuperuser, true
	case TableUserColumnLastLoginAt:
		return &e.LastLoginAt, true
	case TableUserColumnLastName:
		return &e.LastName, true
	case TableUserColumnPassword:
		return &e.Password, true
	case TableUserColumnUpdatedAt:
		return &e.UpdatedAt, true
	case TableUserColumnUpdatedBy:
		return &e.UpdatedBy, true
	case TableUserColumnUsername:
		return &e.Username, true
	default:
		return nil, false
	}
}

func (e *UserEntity) Props(cns ...string) ([]interface{}, error) {
	res := make([]interface{}, 0, len(cns))
	for _, cn := range cns {
		if prop, ok := e.Prop(cn); ok {
			res = append(res, prop)
		} else {
			return nil, fmt.Errorf("unexpected column provided: %s", cn)
		}
	}
	return res, nil
}

// UserIterator is not thread safe.
type UserIterator struct {
	rows *sql.Rows
	cols []string
}

func (i *UserIterator) Next() bool {
	return i.rows.Next()
}

func (i *UserIterator) Close() error {
	return i.rows.Close()
}

func (i *UserIterator) Err() error {
	return i.rows.Err()
}

// Columns is wrapper around sql.Rows.Columns method, that also cache outpu inside iterator.
func (i *UserIterator) Columns() ([]string, error) {
	if i.cols == nil {
		cols, err := i.rows.Columns()
		if err != nil {
			return nil, err
		}
		i.cols = cols
	}
	return i.cols, nil
}

// Ent is wrapper around User method that makes iterator more generic.
func (i *UserIterator) Ent() (interface{}, error) {
	return i.User()
}

func (i *UserIterator) User() (*UserEntity, error) {
	var ent UserEntity
	cols, err := i.rows.Columns()
	if err != nil {
		return nil, err
	}

	props, err := ent.Props(cols...)
	if err != nil {
		return nil, err
	}
	if err := i.rows.Scan(props...); err != nil {
		return nil, err
	}
	return &ent, nil
}

type UserCriteria struct {
	Offset, Limit     int64
	Sort              map[string]bool
	ConfirmationToken []byte
	CreatedAt         *qtypes.Timestamp
	CreatedBy         *qtypes.Int64
	FirstName         *qtypes.String
	ID                *qtypes.Int64
	IsActive          ntypes.Bool
	IsConfirmed       ntypes.Bool
	IsStaff           ntypes.Bool
	IsSuperuser       ntypes.Bool
	LastLoginAt       *qtypes.Timestamp
	LastName          *qtypes.String
	Password          []byte
	UpdatedAt         *qtypes.Timestamp
	UpdatedBy         *qtypes.Int64
	Username          *qtypes.String
}

type UserPatch struct {
	ConfirmationToken []byte
	CreatedAt         pq.NullTime
	CreatedBy         ntypes.Int64
	FirstName         ntypes.String
	IsActive          ntypes.Bool
	IsConfirmed       ntypes.Bool
	IsStaff           ntypes.Bool
	IsSuperuser       ntypes.Bool
	LastLoginAt       pq.NullTime
	LastName          ntypes.String
	Password          []byte
	UpdatedAt         pq.NullTime
	UpdatedBy         ntypes.Int64
	Username          ntypes.String
}

func ScanUserRows(rows *sql.Rows) (entities []*UserEntity, err error) {
	for rows.Next() {
		var ent UserEntity
		err = rows.Scan(&ent.ConfirmationToken,
			&ent.CreatedAt,
			&ent.CreatedBy,
			&ent.FirstName,
			&ent.ID,
			&ent.IsActive,
			&ent.IsConfirmed,
			&ent.IsStaff,
			&ent.IsSuperuser,
			&ent.LastLoginAt,
			&ent.LastName,
			&ent.Password,
			&ent.UpdatedAt,
			&ent.UpdatedBy,
			&ent.Username,
		)
		if err != nil {
			return
		}

		entities = append(entities, &ent)
	}
	if err = rows.Err(); err != nil {
		return
	}

	return
}

type UserRepositoryBase struct {
	Table   string
	Columns []string
	DB      *sql.DB
	Debug   bool
	Log     log.Logger
}

func (r *UserRepositoryBase) InsertQuery(e *UserEntity) (string, []interface{}, error) {
	insert := NewComposer(15)
	columns := bytes.NewBuffer(nil)
	buf := bytes.NewBufferString("INSERT INTO ")
	buf.WriteString(r.Table)

	if e.ConfirmationToken != nil {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableUserColumnConfirmationToken); err != nil {
			return "", nil, err
		}
		if insert.Dirty {
			if _, err := insert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := insert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		insert.Add(e.ConfirmationToken)
		insert.Dirty = true
	}

	if !e.CreatedAt.IsZero() {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableUserColumnCreatedAt); err != nil {
			return "", nil, err
		}
		if insert.Dirty {
			if _, err := insert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := insert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		insert.Add(e.CreatedAt)
		insert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserColumnCreatedBy); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.CreatedBy)
	insert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserColumnFirstName); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.FirstName)
	insert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserColumnIsActive); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.IsActive)
	insert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserColumnIsConfirmed); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.IsConfirmed)
	insert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserColumnIsStaff); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.IsStaff)
	insert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserColumnIsSuperuser); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.IsSuperuser)
	insert.Dirty = true

	if e.LastLoginAt.Valid {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableUserColumnLastLoginAt); err != nil {
			return "", nil, err
		}
		if insert.Dirty {
			if _, err := insert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := insert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		insert.Add(e.LastLoginAt)
		insert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserColumnLastName); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.LastName)
	insert.Dirty = true

	if e.Password != nil {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableUserColumnPassword); err != nil {
			return "", nil, err
		}
		if insert.Dirty {
			if _, err := insert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := insert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		insert.Add(e.Password)
		insert.Dirty = true
	}

	if e.UpdatedAt.Valid {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableUserColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if insert.Dirty {
			if _, err := insert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := insert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		insert.Add(e.UpdatedAt)
		insert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserColumnUpdatedBy); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.UpdatedBy)
	insert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserColumnUsername); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.Username)
	insert.Dirty = true
	if columns.Len() > 0 {
		buf.WriteString(" (")
		buf.ReadFrom(columns)
		buf.WriteString(") VALUES (")
		buf.ReadFrom(insert)
		buf.WriteString(") ")
		if len(r.Columns) > 0 {
			buf.WriteString("RETURNING ")
			buf.WriteString(strings.Join(r.Columns, ", "))
		}
	}
	return buf.String(), insert.Args(), nil
}

func (r *UserRepositoryBase) Insert(ctx context.Context, e *UserEntity) (*UserEntity, error) {
	query, args, err := r.InsertQuery(e)
	if err != nil {
		return nil, err
	}
	if err := r.DB.QueryRowContext(ctx, query, args...).Scan(&e.ConfirmationToken,
		&e.CreatedAt,
		&e.CreatedBy,
		&e.FirstName,
		&e.ID,
		&e.IsActive,
		&e.IsConfirmed,
		&e.IsStaff,
		&e.IsSuperuser,
		&e.LastLoginAt,
		&e.LastName,
		&e.Password,
		&e.UpdatedAt,
		&e.UpdatedBy,
		&e.Username,
	); err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}
	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query success", "query", query, "table", r.Table)
	}
	return e, nil
}

func (r *UserRepositoryBase) FindQuery(s []string, c *UserCriteria) (string, []interface{}, error) {
	where := NewComposer(15)
	buf := bytes.NewBufferString("SELECT ")
	buf.WriteString(strings.Join(s, ", "))
	buf.WriteString(" FROM ")
	buf.WriteString(r.Table)
	buf.WriteString(" ")
	if c.ConfirmationToken != nil {
		if where.Dirty {
			where.WriteString(" AND ")
		}
		if _, err := where.WriteString(TableUserColumnConfirmationToken); err != nil {
			return "", nil, err
		}
		if _, err := where.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := where.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		where.Add(c.ConfirmationToken)
		where.Dirty = true
	}

	QueryTimestampWhereClause(c.CreatedAt, TableUserColumnCreatedAt, where, And)

	QueryInt64WhereClause(c.CreatedBy, TableUserColumnCreatedBy, where, And)

	QueryStringWhereClause(c.FirstName, TableUserColumnFirstName, where, And)

	QueryInt64WhereClause(c.ID, TableUserColumnID, where, And)

	if c.IsActive.Valid {
		if where.Dirty {
			if _, err := where.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := where.WriteString(TableUserColumnIsActive); err != nil {
			return "", nil, err
		}
		if _, err := where.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := where.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		where.Add(c.IsActive)
		where.Dirty = true
	}

	if c.IsConfirmed.Valid {
		if where.Dirty {
			if _, err := where.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := where.WriteString(TableUserColumnIsConfirmed); err != nil {
			return "", nil, err
		}
		if _, err := where.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := where.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		where.Add(c.IsConfirmed)
		where.Dirty = true
	}

	if c.IsStaff.Valid {
		if where.Dirty {
			if _, err := where.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := where.WriteString(TableUserColumnIsStaff); err != nil {
			return "", nil, err
		}
		if _, err := where.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := where.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		where.Add(c.IsStaff)
		where.Dirty = true
	}

	if c.IsSuperuser.Valid {
		if where.Dirty {
			if _, err := where.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := where.WriteString(TableUserColumnIsSuperuser); err != nil {
			return "", nil, err
		}
		if _, err := where.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := where.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		where.Add(c.IsSuperuser)
		where.Dirty = true
	}

	QueryTimestampWhereClause(c.LastLoginAt, TableUserColumnLastLoginAt, where, And)

	QueryStringWhereClause(c.LastName, TableUserColumnLastName, where, And)

	if c.Password != nil {
		if where.Dirty {
			where.WriteString(" AND ")
		}
		if _, err := where.WriteString(TableUserColumnPassword); err != nil {
			return "", nil, err
		}
		if _, err := where.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := where.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		where.Add(c.Password)
		where.Dirty = true
	}

	QueryTimestampWhereClause(c.UpdatedAt, TableUserColumnUpdatedAt, where, And)

	QueryInt64WhereClause(c.UpdatedBy, TableUserColumnUpdatedBy, where, And)

	QueryStringWhereClause(c.Username, TableUserColumnUsername, where, And)
	if where.Dirty {
		if _, err := buf.WriteString("WHERE "); err != nil {
			return "", nil, err
		}
		buf.ReadFrom(where)
	}

	if len(c.Sort) > 0 {
		i := 0
		where.WriteString(" ORDER BY ")

		for cn, asc := range c.Sort {
			for _, tcn := range TableUserColumns {
				if cn == tcn {
					if i > 0 {
						if _, err := where.WriteString(", "); err != nil {
							return "", nil, err
						}
					}
					if _, err := where.WriteString(cn); err != nil {
						return "", nil, err
					}
					if !asc {
						if _, err := where.WriteString(" DESC "); err != nil {
							return "", nil, err
						}
					}
					i++
					break
				}
			}
		}
	}
	if c.Offset > 0 {
		if _, err := where.WriteString(" OFFSET "); err != nil {
			return "", nil, err
		}
		if err := where.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		if _, err := where.WriteString(" "); err != nil {
			return "", nil, err
		}
		where.Add(c.Offset)
	}
	if c.Limit > 0 {
		if _, err := where.WriteString(" LIMIT "); err != nil {
			return "", nil, err
		}
		if err := where.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		if _, err := where.WriteString(" "); err != nil {
			return "", nil, err
		}
		where.Add(c.Limit)
	}

	buf.ReadFrom(where)

	return buf.String(), where.Args(), nil
}

func (r *UserRepositoryBase) Find(ctx context.Context, c *UserCriteria) ([]*UserEntity, error) {
	query, args, err := r.FindQuery(r.Columns, c)
	if err != nil {
		return nil, err
	}
	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "find query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}
	defer rows.Close()

	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "find query success", "query", query, "table", r.Table)
	}

	return ScanUserRows(rows)
}

func (r *UserRepositoryBase) FindIter(ctx context.Context, c *UserCriteria) (*UserIterator, error) {
	query, args, err := r.FindQuery(r.Columns, c)
	if err != nil {
		return nil, err
	}
	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &UserIterator{rows: rows}, nil
}
func (r *UserRepositoryBase) FindOneByID(ctx context.Context, pk int64) (*UserEntity, error) {
	find := NewComposer(15)
	find.WriteString("SELECT ")
	find.WriteString(strings.Join(r.Columns, ", "))
	find.WriteString(" FROM ")
	find.WriteString(TableUser)
	find.WriteString(" WHERE ")
	find.WriteString(TableUserColumnID)
	find.WriteString("=")
	find.WritePlaceholder()
	find.Add(pk)
	var (
		ent UserEntity
	)
	props, err := ent.Props(r.Columns...)
	if err != nil {
		return nil, err
	}
	err = r.DB.QueryRowContext(ctx, find.String(), find.Args()...).Scan(props...)
	if err != nil {
		return nil, err
	}

	return &ent, nil
}
func (r *UserRepositoryBase) FindOneByUsername(ctx context.Context, userUsername string) (*UserEntity, error) {
	find := NewComposer(15)
	find.WriteString("SELECT ")
	find.WriteString(strings.Join(r.Columns, ", "))
	find.WriteString(" FROM ")
	find.WriteString(TableUser)
	find.WriteString(" WHERE ")
	find.WriteString(TableUserColumnUsername)
	find.WriteString("=")
	find.WritePlaceholder()
	find.Add(userUsername)

	var (
		ent UserEntity
	)
	props, err := ent.Props(r.Columns...)
	if err != nil {
		return nil, err
	}
	err = r.DB.QueryRowContext(ctx, find.String(), find.Args()...).Scan(props...)
	if err != nil {
		return nil, err
	}

	return &ent, nil
}
func (r *UserRepositoryBase) UpdateOneByIDQuery(pk int64, p *UserPatch) (string, []interface{}, error) {
	buf := bytes.NewBufferString("UPDATE ")
	buf.WriteString(r.Table)
	update := NewComposer(15)
	if p.ConfirmationToken != nil {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnConfirmationToken); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.ConfirmationToken)
		update.Dirty = true

	}

	if p.CreatedAt.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnCreatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.CreatedAt)
		update.Dirty = true

	}

	if p.CreatedBy.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnCreatedBy); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.CreatedBy)
		update.Dirty = true
	}

	if p.FirstName.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnFirstName); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.FirstName)
		update.Dirty = true
	}

	if p.IsActive.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnIsActive); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.IsActive)
		update.Dirty = true
	}

	if p.IsConfirmed.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnIsConfirmed); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.IsConfirmed)
		update.Dirty = true
	}

	if p.IsStaff.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnIsStaff); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.IsStaff)
		update.Dirty = true
	}

	if p.IsSuperuser.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnIsSuperuser); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.IsSuperuser)
		update.Dirty = true
	}

	if p.LastLoginAt.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnLastLoginAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.LastLoginAt)
		update.Dirty = true

	}

	if p.LastName.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnLastName); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.LastName)
		update.Dirty = true
	}

	if p.Password != nil {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnPassword); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.Password)
		update.Dirty = true

	}

	if p.UpdatedAt.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.UpdatedAt)
		update.Dirty = true

	} else {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("=NOW()"); err != nil {
			return "", nil, err
		}
	}

	if p.UpdatedBy.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnUpdatedBy); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.UpdatedBy)
		update.Dirty = true
	}

	if p.Username.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnUsername); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.Username)
		update.Dirty = true
	}

	if !update.Dirty {
		return "", nil, errors.New("User update failure, nothing to update")
	}
	buf.WriteString(" SET ")
	buf.ReadFrom(update)
	buf.WriteString(" WHERE ")

	update.WriteString(TableUserColumnID)
	update.WriteString("=")
	update.WritePlaceholder()
	update.Add(pk)

	buf.ReadFrom(update)
	buf.WriteString(" RETURNING ")
	buf.WriteString(strings.Join(r.Columns, ", "))

	return buf.String(), update.Args(), nil
}
func (r *UserRepositoryBase) UpdateOneByID(ctx context.Context, pk int64, p *UserPatch) (*UserEntity, error) {
	query, args, err := r.UpdateOneByIDQuery(pk, p)
	if err != nil {
		return nil, err
	}
	var ent UserEntity
	props, err := ent.Props(r.Columns...)
	if err != nil {
		return nil, err
	}
	err = r.DB.QueryRowContext(ctx, query, args...).Scan(props...)
	if err != nil {
		return nil, err
	}

	return &ent, nil
}
func (r *UserRepositoryBase) UpdateOneByUsernameQuery(userUsername string, p *UserPatch) (string, []interface{}, error) {
	buf := bytes.NewBufferString("UPDATE ")
	buf.WriteString(r.Table)
	update := NewComposer(1)
	if p.ConfirmationToken != nil {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnConfirmationToken); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.ConfirmationToken)
		update.Dirty = true

	}

	if p.CreatedAt.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnCreatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.CreatedAt)
		update.Dirty = true

	}

	if p.CreatedBy.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnCreatedBy); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.CreatedBy)
		update.Dirty = true
	}

	if p.FirstName.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnFirstName); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.FirstName)
		update.Dirty = true
	}

	if p.IsActive.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnIsActive); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.IsActive)
		update.Dirty = true
	}

	if p.IsConfirmed.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnIsConfirmed); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.IsConfirmed)
		update.Dirty = true
	}

	if p.IsStaff.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnIsStaff); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.IsStaff)
		update.Dirty = true
	}

	if p.IsSuperuser.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnIsSuperuser); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.IsSuperuser)
		update.Dirty = true
	}

	if p.LastLoginAt.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnLastLoginAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.LastLoginAt)
		update.Dirty = true

	}

	if p.LastName.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnLastName); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.LastName)
		update.Dirty = true
	}

	if p.Password != nil {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnPassword); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.Password)
		update.Dirty = true

	}

	if p.UpdatedAt.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.UpdatedAt)
		update.Dirty = true

	} else {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("=NOW()"); err != nil {
			return "", nil, err
		}
	}

	if p.UpdatedBy.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnUpdatedBy); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.UpdatedBy)
		update.Dirty = true
	}

	if p.Username.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserColumnUsername); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.Username)
		update.Dirty = true
	}

	if !update.Dirty {
		return "", nil, errors.New("User update failure, nothing to update")
	}
	buf.WriteString(" SET ")
	buf.ReadFrom(update)
	buf.WriteString(" WHERE ")
	update.WriteString(TableUserColumnUsername)
	update.WriteString("=")
	update.WritePlaceholder()
	update.Add(userUsername)
	buf.ReadFrom(update)
	buf.WriteString(" RETURNING ")
	buf.WriteString(strings.Join(r.Columns, ", "))
	return buf.String(), update.Args(), nil
}
func (r *UserRepositoryBase) UpdateOneByUsername(ctx context.Context, userUsername string, p *UserPatch) (*UserEntity, error) {
	query, args, err := r.UpdateOneByUsernameQuery(userUsername, p)
	if err != nil {
		return nil, err
	}
	var ent UserEntity
	props, err := ent.Props(r.Columns...)
	if err != nil {
		return nil, err
	}
	err = r.DB.QueryRowContext(ctx, query, args...).Scan(props...)
	if err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}

	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query success", "query", query, "table", r.Table)
	}

	return &ent, nil
}
func (r *UserRepositoryBase) UpsertQuery(e *UserEntity, p *UserPatch, inf ...string) (string, []interface{}, error) {
	upsert := NewComposer(30)
	columns := bytes.NewBuffer(nil)
	buf := bytes.NewBufferString("INSERT INTO ")
	buf.WriteString(r.Table)

	if e.ConfirmationToken != nil {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableUserColumnConfirmationToken); err != nil {
			return "", nil, err
		}
		if upsert.Dirty {
			if _, err := upsert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := upsert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		upsert.Add(e.ConfirmationToken)
		upsert.Dirty = true
	}

	if !e.CreatedAt.IsZero() {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableUserColumnCreatedAt); err != nil {
			return "", nil, err
		}
		if upsert.Dirty {
			if _, err := upsert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := upsert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		upsert.Add(e.CreatedAt)
		upsert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserColumnCreatedBy); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.CreatedBy)
	upsert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserColumnFirstName); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.FirstName)
	upsert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserColumnIsActive); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.IsActive)
	upsert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserColumnIsConfirmed); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.IsConfirmed)
	upsert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserColumnIsStaff); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.IsStaff)
	upsert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserColumnIsSuperuser); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.IsSuperuser)
	upsert.Dirty = true

	if e.LastLoginAt.Valid {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableUserColumnLastLoginAt); err != nil {
			return "", nil, err
		}
		if upsert.Dirty {
			if _, err := upsert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := upsert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		upsert.Add(e.LastLoginAt)
		upsert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserColumnLastName); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.LastName)
	upsert.Dirty = true

	if e.Password != nil {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableUserColumnPassword); err != nil {
			return "", nil, err
		}
		if upsert.Dirty {
			if _, err := upsert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := upsert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		upsert.Add(e.Password)
		upsert.Dirty = true
	}

	if e.UpdatedAt.Valid {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableUserColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if upsert.Dirty {
			if _, err := upsert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := upsert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		upsert.Add(e.UpdatedAt)
		upsert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserColumnUpdatedBy); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.UpdatedBy)
	upsert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserColumnUsername); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.Username)
	upsert.Dirty = true

	if upsert.Dirty {
		buf.WriteString(" (")
		buf.ReadFrom(columns)
		buf.WriteString(") VALUES (")
		buf.ReadFrom(upsert)
		buf.WriteString(")")
	}
	buf.WriteString(" ON CONFLICT ")
	if len(inf) > 0 {
		upsert.Dirty = false
		if p.ConfirmationToken != nil {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserColumnConfirmationToken); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.ConfirmationToken)
			upsert.Dirty = true

		}

		if p.CreatedAt.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserColumnCreatedAt); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.CreatedAt)
			upsert.Dirty = true

		}

		if p.CreatedBy.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserColumnCreatedBy); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.CreatedBy)
			upsert.Dirty = true
		}

		if p.FirstName.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserColumnFirstName); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.FirstName)
			upsert.Dirty = true
		}

		if p.IsActive.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserColumnIsActive); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.IsActive)
			upsert.Dirty = true
		}

		if p.IsConfirmed.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserColumnIsConfirmed); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.IsConfirmed)
			upsert.Dirty = true
		}

		if p.IsStaff.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserColumnIsStaff); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.IsStaff)
			upsert.Dirty = true
		}

		if p.IsSuperuser.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserColumnIsSuperuser); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.IsSuperuser)
			upsert.Dirty = true
		}

		if p.LastLoginAt.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserColumnLastLoginAt); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.LastLoginAt)
			upsert.Dirty = true

		}

		if p.LastName.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserColumnLastName); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.LastName)
			upsert.Dirty = true
		}

		if p.Password != nil {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserColumnPassword); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.Password)
			upsert.Dirty = true

		}

		if p.UpdatedAt.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserColumnUpdatedAt); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.UpdatedAt)
			upsert.Dirty = true

		} else {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserColumnUpdatedAt); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("=NOW()"); err != nil {
				return "", nil, err
			}
		}

		if p.UpdatedBy.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserColumnUpdatedBy); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.UpdatedBy)
			upsert.Dirty = true
		}

		if p.Username.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserColumnUsername); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.Username)
			upsert.Dirty = true
		}

	}

	if len(inf) > 0 && upsert.Dirty {
		buf.WriteString("(")
		for j, i := range inf {
			if j != 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(i)
		}
		buf.WriteString(")")
		buf.WriteString(" DO UPDATE SET ")
		buf.ReadFrom(upsert)
	} else {
		buf.WriteString(" DO NOTHING ")
	}
	if upsert.Dirty {
		if len(r.Columns) > 0 {
			buf.WriteString(" RETURNING ")
			buf.WriteString(strings.Join(r.Columns, ", "))
		}
	}
	return buf.String(), upsert.Args(), nil
}
func (r *UserRepositoryBase) Upsert(ctx context.Context, e *UserEntity, p *UserPatch, inf ...string) (*UserEntity, error) {
	query, args, err := r.UpsertQuery(e, p, inf...)
	if err != nil {
		return nil, err
	}
	if err := r.DB.QueryRowContext(ctx, query, args...).Scan(&e.ConfirmationToken,
		&e.CreatedAt,
		&e.CreatedBy,
		&e.FirstName,
		&e.ID,
		&e.IsActive,
		&e.IsConfirmed,
		&e.IsStaff,
		&e.IsSuperuser,
		&e.LastLoginAt,
		&e.LastName,
		&e.Password,
		&e.UpdatedAt,
		&e.UpdatedBy,
		&e.Username,
	); err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "upsert query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}
	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "upsert query success", "query", query, "table", r.Table)
	}
	return e, nil
}

func (r *UserRepositoryBase) Count(ctx context.Context, c *UserCriteria) (int64, error) {
	query, args, err := r.FindQuery([]string{"COUNT(*)"}, c)
	if err != nil {
		return 0, err
	}
	var count int64
	if err := r.DB.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "count query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return 0, err
	}

	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "count query success", "query", query, "table", r.Table)
	}

	return count, nil
}
func (r *UserRepositoryBase) DeleteOneByID(ctx context.Context, pk int64) (int64, error) {
	find := NewComposer(15)
	find.WriteString("DELETE FROM ")
	find.WriteString(TableUser)
	find.WriteString(" WHERE ")
	find.WriteString(TableUserColumnID)
	find.WriteString("=")
	find.WritePlaceholder()
	find.Add(pk)
	res, err := r.DB.ExecContext(ctx, find.String(), find.Args()...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

const (
	TableGroup                              = "charon.group"
	TableGroupColumnCreatedAt               = "created_at"
	TableGroupColumnCreatedBy               = "created_by"
	TableGroupColumnDescription             = "description"
	TableGroupColumnID                      = "id"
	TableGroupColumnName                    = "name"
	TableGroupColumnUpdatedAt               = "updated_at"
	TableGroupColumnUpdatedBy               = "updated_by"
	TableGroupConstraintCreatedByForeignKey = "charon.group_created_by_fkey"

	TableGroupConstraintPrimaryKey = "charon.group_id_pkey"

	TableGroupConstraintNameUnique = "charon.group_name_key"

	TableGroupConstraintUpdatedByForeignKey = "charon.group_updated_by_fkey"
)

var (
	TableGroupColumns = []string{
		TableGroupColumnCreatedAt,
		TableGroupColumnCreatedBy,
		TableGroupColumnDescription,
		TableGroupColumnID,
		TableGroupColumnName,
		TableGroupColumnUpdatedAt,
		TableGroupColumnUpdatedBy,
	}
)

// GroupEntity ...
type GroupEntity struct {
	// CreatedAt ...
	CreatedAt time.Time
	// CreatedBy ...
	CreatedBy ntypes.Int64
	// Description ...
	Description ntypes.String
	// ID ...
	ID int64
	// Name ...
	Name string
	// UpdatedAt ...
	UpdatedAt pq.NullTime
	// UpdatedBy ...
	UpdatedBy ntypes.Int64
	// Author ...
	Author *UserEntity
	// Modifier ...
	Modifier *UserEntity
	// Permissions ...
	Permissions []*PermissionEntity
	// Users ...
	Users []*UserEntity
}

func (e *GroupEntity) Prop(cn string) (interface{}, bool) {
	switch cn {
	case TableGroupColumnCreatedAt:
		return &e.CreatedAt, true
	case TableGroupColumnCreatedBy:
		return &e.CreatedBy, true
	case TableGroupColumnDescription:
		return &e.Description, true
	case TableGroupColumnID:
		return &e.ID, true
	case TableGroupColumnName:
		return &e.Name, true
	case TableGroupColumnUpdatedAt:
		return &e.UpdatedAt, true
	case TableGroupColumnUpdatedBy:
		return &e.UpdatedBy, true
	default:
		return nil, false
	}
}

func (e *GroupEntity) Props(cns ...string) ([]interface{}, error) {
	res := make([]interface{}, 0, len(cns))
	for _, cn := range cns {
		if prop, ok := e.Prop(cn); ok {
			res = append(res, prop)
		} else {
			return nil, fmt.Errorf("unexpected column provided: %s", cn)
		}
	}
	return res, nil
}

// GroupIterator is not thread safe.
type GroupIterator struct {
	rows *sql.Rows
	cols []string
}

func (i *GroupIterator) Next() bool {
	return i.rows.Next()
}

func (i *GroupIterator) Close() error {
	return i.rows.Close()
}

func (i *GroupIterator) Err() error {
	return i.rows.Err()
}

// Columns is wrapper around sql.Rows.Columns method, that also cache outpu inside iterator.
func (i *GroupIterator) Columns() ([]string, error) {
	if i.cols == nil {
		cols, err := i.rows.Columns()
		if err != nil {
			return nil, err
		}
		i.cols = cols
	}
	return i.cols, nil
}

// Ent is wrapper around Group method that makes iterator more generic.
func (i *GroupIterator) Ent() (interface{}, error) {
	return i.Group()
}

func (i *GroupIterator) Group() (*GroupEntity, error) {
	var ent GroupEntity
	cols, err := i.rows.Columns()
	if err != nil {
		return nil, err
	}

	props, err := ent.Props(cols...)
	if err != nil {
		return nil, err
	}
	if err := i.rows.Scan(props...); err != nil {
		return nil, err
	}
	return &ent, nil
}

type GroupCriteria struct {
	Offset, Limit int64
	Sort          map[string]bool
	CreatedAt     *qtypes.Timestamp
	CreatedBy     *qtypes.Int64
	Description   *qtypes.String
	ID            *qtypes.Int64
	Name          *qtypes.String
	UpdatedAt     *qtypes.Timestamp
	UpdatedBy     *qtypes.Int64
}

type GroupPatch struct {
	CreatedAt   pq.NullTime
	CreatedBy   ntypes.Int64
	Description ntypes.String
	Name        ntypes.String
	UpdatedAt   pq.NullTime
	UpdatedBy   ntypes.Int64
}

func ScanGroupRows(rows *sql.Rows) (entities []*GroupEntity, err error) {
	for rows.Next() {
		var ent GroupEntity
		err = rows.Scan(&ent.CreatedAt,
			&ent.CreatedBy,
			&ent.Description,
			&ent.ID,
			&ent.Name,
			&ent.UpdatedAt,
			&ent.UpdatedBy,
		)
		if err != nil {
			return
		}

		entities = append(entities, &ent)
	}
	if err = rows.Err(); err != nil {
		return
	}

	return
}

type GroupRepositoryBase struct {
	Table   string
	Columns []string
	DB      *sql.DB
	Debug   bool
	Log     log.Logger
}

func (r *GroupRepositoryBase) InsertQuery(e *GroupEntity) (string, []interface{}, error) {
	insert := NewComposer(7)
	columns := bytes.NewBuffer(nil)
	buf := bytes.NewBufferString("INSERT INTO ")
	buf.WriteString(r.Table)

	if !e.CreatedAt.IsZero() {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableGroupColumnCreatedAt); err != nil {
			return "", nil, err
		}
		if insert.Dirty {
			if _, err := insert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := insert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		insert.Add(e.CreatedAt)
		insert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableGroupColumnCreatedBy); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.CreatedBy)
	insert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableGroupColumnDescription); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.Description)
	insert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableGroupColumnName); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.Name)
	insert.Dirty = true

	if e.UpdatedAt.Valid {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableGroupColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if insert.Dirty {
			if _, err := insert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := insert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		insert.Add(e.UpdatedAt)
		insert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableGroupColumnUpdatedBy); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.UpdatedBy)
	insert.Dirty = true
	if columns.Len() > 0 {
		buf.WriteString(" (")
		buf.ReadFrom(columns)
		buf.WriteString(") VALUES (")
		buf.ReadFrom(insert)
		buf.WriteString(") ")
		if len(r.Columns) > 0 {
			buf.WriteString("RETURNING ")
			buf.WriteString(strings.Join(r.Columns, ", "))
		}
	}
	return buf.String(), insert.Args(), nil
}

func (r *GroupRepositoryBase) Insert(ctx context.Context, e *GroupEntity) (*GroupEntity, error) {
	query, args, err := r.InsertQuery(e)
	if err != nil {
		return nil, err
	}
	if err := r.DB.QueryRowContext(ctx, query, args...).Scan(&e.CreatedAt,
		&e.CreatedBy,
		&e.Description,
		&e.ID,
		&e.Name,
		&e.UpdatedAt,
		&e.UpdatedBy,
	); err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}
	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query success", "query", query, "table", r.Table)
	}
	return e, nil
}

func (r *GroupRepositoryBase) FindQuery(s []string, c *GroupCriteria) (string, []interface{}, error) {
	where := NewComposer(7)
	buf := bytes.NewBufferString("SELECT ")
	buf.WriteString(strings.Join(s, ", "))
	buf.WriteString(" FROM ")
	buf.WriteString(r.Table)
	buf.WriteString(" ")
	QueryTimestampWhereClause(c.CreatedAt, TableGroupColumnCreatedAt, where, And)

	QueryInt64WhereClause(c.CreatedBy, TableGroupColumnCreatedBy, where, And)

	QueryStringWhereClause(c.Description, TableGroupColumnDescription, where, And)

	QueryInt64WhereClause(c.ID, TableGroupColumnID, where, And)

	QueryStringWhereClause(c.Name, TableGroupColumnName, where, And)

	QueryTimestampWhereClause(c.UpdatedAt, TableGroupColumnUpdatedAt, where, And)

	QueryInt64WhereClause(c.UpdatedBy, TableGroupColumnUpdatedBy, where, And)
	if where.Dirty {
		if _, err := buf.WriteString("WHERE "); err != nil {
			return "", nil, err
		}
		buf.ReadFrom(where)
	}

	if len(c.Sort) > 0 {
		i := 0
		where.WriteString(" ORDER BY ")

		for cn, asc := range c.Sort {
			for _, tcn := range TableGroupColumns {
				if cn == tcn {
					if i > 0 {
						if _, err := where.WriteString(", "); err != nil {
							return "", nil, err
						}
					}
					if _, err := where.WriteString(cn); err != nil {
						return "", nil, err
					}
					if !asc {
						if _, err := where.WriteString(" DESC "); err != nil {
							return "", nil, err
						}
					}
					i++
					break
				}
			}
		}
	}
	if c.Offset > 0 {
		if _, err := where.WriteString(" OFFSET "); err != nil {
			return "", nil, err
		}
		if err := where.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		if _, err := where.WriteString(" "); err != nil {
			return "", nil, err
		}
		where.Add(c.Offset)
	}
	if c.Limit > 0 {
		if _, err := where.WriteString(" LIMIT "); err != nil {
			return "", nil, err
		}
		if err := where.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		if _, err := where.WriteString(" "); err != nil {
			return "", nil, err
		}
		where.Add(c.Limit)
	}

	buf.ReadFrom(where)

	return buf.String(), where.Args(), nil
}

func (r *GroupRepositoryBase) Find(ctx context.Context, c *GroupCriteria) ([]*GroupEntity, error) {
	query, args, err := r.FindQuery(r.Columns, c)
	if err != nil {
		return nil, err
	}
	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "find query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}
	defer rows.Close()

	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "find query success", "query", query, "table", r.Table)
	}

	return ScanGroupRows(rows)
}

func (r *GroupRepositoryBase) FindIter(ctx context.Context, c *GroupCriteria) (*GroupIterator, error) {
	query, args, err := r.FindQuery(r.Columns, c)
	if err != nil {
		return nil, err
	}
	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &GroupIterator{rows: rows}, nil
}
func (r *GroupRepositoryBase) FindOneByID(ctx context.Context, pk int64) (*GroupEntity, error) {
	find := NewComposer(7)
	find.WriteString("SELECT ")
	find.WriteString(strings.Join(r.Columns, ", "))
	find.WriteString(" FROM ")
	find.WriteString(TableGroup)
	find.WriteString(" WHERE ")
	find.WriteString(TableGroupColumnID)
	find.WriteString("=")
	find.WritePlaceholder()
	find.Add(pk)
	var (
		ent GroupEntity
	)
	props, err := ent.Props(r.Columns...)
	if err != nil {
		return nil, err
	}
	err = r.DB.QueryRowContext(ctx, find.String(), find.Args()...).Scan(props...)
	if err != nil {
		return nil, err
	}

	return &ent, nil
}
func (r *GroupRepositoryBase) FindOneByName(ctx context.Context, groupName string) (*GroupEntity, error) {
	find := NewComposer(7)
	find.WriteString("SELECT ")
	find.WriteString(strings.Join(r.Columns, ", "))
	find.WriteString(" FROM ")
	find.WriteString(TableGroup)
	find.WriteString(" WHERE ")
	find.WriteString(TableGroupColumnName)
	find.WriteString("=")
	find.WritePlaceholder()
	find.Add(groupName)

	var (
		ent GroupEntity
	)
	props, err := ent.Props(r.Columns...)
	if err != nil {
		return nil, err
	}
	err = r.DB.QueryRowContext(ctx, find.String(), find.Args()...).Scan(props...)
	if err != nil {
		return nil, err
	}

	return &ent, nil
}
func (r *GroupRepositoryBase) UpdateOneByIDQuery(pk int64, p *GroupPatch) (string, []interface{}, error) {
	buf := bytes.NewBufferString("UPDATE ")
	buf.WriteString(r.Table)
	update := NewComposer(7)
	if p.CreatedAt.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupColumnCreatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.CreatedAt)
		update.Dirty = true

	}

	if p.CreatedBy.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupColumnCreatedBy); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.CreatedBy)
		update.Dirty = true
	}

	if p.Description.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupColumnDescription); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.Description)
		update.Dirty = true
	}

	if p.Name.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupColumnName); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.Name)
		update.Dirty = true
	}

	if p.UpdatedAt.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.UpdatedAt)
		update.Dirty = true

	} else {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("=NOW()"); err != nil {
			return "", nil, err
		}
	}

	if p.UpdatedBy.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupColumnUpdatedBy); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.UpdatedBy)
		update.Dirty = true
	}

	if !update.Dirty {
		return "", nil, errors.New("Group update failure, nothing to update")
	}
	buf.WriteString(" SET ")
	buf.ReadFrom(update)
	buf.WriteString(" WHERE ")

	update.WriteString(TableGroupColumnID)
	update.WriteString("=")
	update.WritePlaceholder()
	update.Add(pk)

	buf.ReadFrom(update)
	buf.WriteString(" RETURNING ")
	buf.WriteString(strings.Join(r.Columns, ", "))

	return buf.String(), update.Args(), nil
}
func (r *GroupRepositoryBase) UpdateOneByID(ctx context.Context, pk int64, p *GroupPatch) (*GroupEntity, error) {
	query, args, err := r.UpdateOneByIDQuery(pk, p)
	if err != nil {
		return nil, err
	}
	var ent GroupEntity
	props, err := ent.Props(r.Columns...)
	if err != nil {
		return nil, err
	}
	err = r.DB.QueryRowContext(ctx, query, args...).Scan(props...)
	if err != nil {
		return nil, err
	}

	return &ent, nil
}
func (r *GroupRepositoryBase) UpdateOneByNameQuery(groupName string, p *GroupPatch) (string, []interface{}, error) {
	buf := bytes.NewBufferString("UPDATE ")
	buf.WriteString(r.Table)
	update := NewComposer(1)
	if p.CreatedAt.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupColumnCreatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.CreatedAt)
		update.Dirty = true

	}

	if p.CreatedBy.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupColumnCreatedBy); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.CreatedBy)
		update.Dirty = true
	}

	if p.Description.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupColumnDescription); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.Description)
		update.Dirty = true
	}

	if p.Name.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupColumnName); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.Name)
		update.Dirty = true
	}

	if p.UpdatedAt.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.UpdatedAt)
		update.Dirty = true

	} else {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("=NOW()"); err != nil {
			return "", nil, err
		}
	}

	if p.UpdatedBy.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupColumnUpdatedBy); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.UpdatedBy)
		update.Dirty = true
	}

	if !update.Dirty {
		return "", nil, errors.New("Group update failure, nothing to update")
	}
	buf.WriteString(" SET ")
	buf.ReadFrom(update)
	buf.WriteString(" WHERE ")
	update.WriteString(TableGroupColumnName)
	update.WriteString("=")
	update.WritePlaceholder()
	update.Add(groupName)
	buf.ReadFrom(update)
	buf.WriteString(" RETURNING ")
	buf.WriteString(strings.Join(r.Columns, ", "))
	return buf.String(), update.Args(), nil
}
func (r *GroupRepositoryBase) UpdateOneByName(ctx context.Context, groupName string, p *GroupPatch) (*GroupEntity, error) {
	query, args, err := r.UpdateOneByNameQuery(groupName, p)
	if err != nil {
		return nil, err
	}
	var ent GroupEntity
	props, err := ent.Props(r.Columns...)
	if err != nil {
		return nil, err
	}
	err = r.DB.QueryRowContext(ctx, query, args...).Scan(props...)
	if err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}

	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query success", "query", query, "table", r.Table)
	}

	return &ent, nil
}
func (r *GroupRepositoryBase) UpsertQuery(e *GroupEntity, p *GroupPatch, inf ...string) (string, []interface{}, error) {
	upsert := NewComposer(14)
	columns := bytes.NewBuffer(nil)
	buf := bytes.NewBufferString("INSERT INTO ")
	buf.WriteString(r.Table)

	if !e.CreatedAt.IsZero() {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableGroupColumnCreatedAt); err != nil {
			return "", nil, err
		}
		if upsert.Dirty {
			if _, err := upsert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := upsert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		upsert.Add(e.CreatedAt)
		upsert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableGroupColumnCreatedBy); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.CreatedBy)
	upsert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableGroupColumnDescription); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.Description)
	upsert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableGroupColumnName); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.Name)
	upsert.Dirty = true

	if e.UpdatedAt.Valid {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableGroupColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if upsert.Dirty {
			if _, err := upsert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := upsert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		upsert.Add(e.UpdatedAt)
		upsert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableGroupColumnUpdatedBy); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.UpdatedBy)
	upsert.Dirty = true

	if upsert.Dirty {
		buf.WriteString(" (")
		buf.ReadFrom(columns)
		buf.WriteString(") VALUES (")
		buf.ReadFrom(upsert)
		buf.WriteString(")")
	}
	buf.WriteString(" ON CONFLICT ")
	if len(inf) > 0 {
		upsert.Dirty = false
		if p.CreatedAt.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableGroupColumnCreatedAt); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.CreatedAt)
			upsert.Dirty = true

		}

		if p.CreatedBy.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableGroupColumnCreatedBy); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.CreatedBy)
			upsert.Dirty = true
		}

		if p.Description.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableGroupColumnDescription); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.Description)
			upsert.Dirty = true
		}

		if p.Name.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableGroupColumnName); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.Name)
			upsert.Dirty = true
		}

		if p.UpdatedAt.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableGroupColumnUpdatedAt); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.UpdatedAt)
			upsert.Dirty = true

		} else {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableGroupColumnUpdatedAt); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("=NOW()"); err != nil {
				return "", nil, err
			}
		}

		if p.UpdatedBy.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableGroupColumnUpdatedBy); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.UpdatedBy)
			upsert.Dirty = true
		}

	}

	if len(inf) > 0 && upsert.Dirty {
		buf.WriteString("(")
		for j, i := range inf {
			if j != 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(i)
		}
		buf.WriteString(")")
		buf.WriteString(" DO UPDATE SET ")
		buf.ReadFrom(upsert)
	} else {
		buf.WriteString(" DO NOTHING ")
	}
	if upsert.Dirty {
		if len(r.Columns) > 0 {
			buf.WriteString(" RETURNING ")
			buf.WriteString(strings.Join(r.Columns, ", "))
		}
	}
	return buf.String(), upsert.Args(), nil
}
func (r *GroupRepositoryBase) Upsert(ctx context.Context, e *GroupEntity, p *GroupPatch, inf ...string) (*GroupEntity, error) {
	query, args, err := r.UpsertQuery(e, p, inf...)
	if err != nil {
		return nil, err
	}
	if err := r.DB.QueryRowContext(ctx, query, args...).Scan(&e.CreatedAt,
		&e.CreatedBy,
		&e.Description,
		&e.ID,
		&e.Name,
		&e.UpdatedAt,
		&e.UpdatedBy,
	); err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "upsert query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}
	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "upsert query success", "query", query, "table", r.Table)
	}
	return e, nil
}

func (r *GroupRepositoryBase) Count(ctx context.Context, c *GroupCriteria) (int64, error) {
	query, args, err := r.FindQuery([]string{"COUNT(*)"}, c)
	if err != nil {
		return 0, err
	}
	var count int64
	if err := r.DB.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "count query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return 0, err
	}

	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "count query success", "query", query, "table", r.Table)
	}

	return count, nil
}
func (r *GroupRepositoryBase) DeleteOneByID(ctx context.Context, pk int64) (int64, error) {
	find := NewComposer(7)
	find.WriteString("DELETE FROM ")
	find.WriteString(TableGroup)
	find.WriteString(" WHERE ")
	find.WriteString(TableGroupColumnID)
	find.WriteString("=")
	find.WritePlaceholder()
	find.Add(pk)
	res, err := r.DB.ExecContext(ctx, find.String(), find.Args()...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

const (
	TablePermission                     = "charon.permission"
	TablePermissionColumnAction         = "action"
	TablePermissionColumnCreatedAt      = "created_at"
	TablePermissionColumnID             = "id"
	TablePermissionColumnModule         = "module"
	TablePermissionColumnSubsystem      = "subsystem"
	TablePermissionColumnUpdatedAt      = "updated_at"
	TablePermissionConstraintPrimaryKey = "charon.permission_id_pkey"

	TablePermissionConstraintSubsystemModuleActionUnique = "charon.permission_subsystem_module_action_key"
)

var (
	TablePermissionColumns = []string{
		TablePermissionColumnAction,
		TablePermissionColumnCreatedAt,
		TablePermissionColumnID,
		TablePermissionColumnModule,
		TablePermissionColumnSubsystem,
		TablePermissionColumnUpdatedAt,
	}
)

// PermissionEntity ...
type PermissionEntity struct {
	// Action ...
	Action string
	// CreatedAt ...
	CreatedAt time.Time
	// ID ...
	ID int64
	// Module ...
	Module string
	// Subsystem ...
	Subsystem string
	// UpdatedAt ...
	UpdatedAt pq.NullTime
	// Groups ...
	Groups []*GroupEntity
	// Users ...
	Users []*UserEntity
}

func (e *PermissionEntity) Prop(cn string) (interface{}, bool) {
	switch cn {
	case TablePermissionColumnAction:
		return &e.Action, true
	case TablePermissionColumnCreatedAt:
		return &e.CreatedAt, true
	case TablePermissionColumnID:
		return &e.ID, true
	case TablePermissionColumnModule:
		return &e.Module, true
	case TablePermissionColumnSubsystem:
		return &e.Subsystem, true
	case TablePermissionColumnUpdatedAt:
		return &e.UpdatedAt, true
	default:
		return nil, false
	}
}

func (e *PermissionEntity) Props(cns ...string) ([]interface{}, error) {
	res := make([]interface{}, 0, len(cns))
	for _, cn := range cns {
		if prop, ok := e.Prop(cn); ok {
			res = append(res, prop)
		} else {
			return nil, fmt.Errorf("unexpected column provided: %s", cn)
		}
	}
	return res, nil
}

// PermissionIterator is not thread safe.
type PermissionIterator struct {
	rows *sql.Rows
	cols []string
}

func (i *PermissionIterator) Next() bool {
	return i.rows.Next()
}

func (i *PermissionIterator) Close() error {
	return i.rows.Close()
}

func (i *PermissionIterator) Err() error {
	return i.rows.Err()
}

// Columns is wrapper around sql.Rows.Columns method, that also cache outpu inside iterator.
func (i *PermissionIterator) Columns() ([]string, error) {
	if i.cols == nil {
		cols, err := i.rows.Columns()
		if err != nil {
			return nil, err
		}
		i.cols = cols
	}
	return i.cols, nil
}

// Ent is wrapper around Permission method that makes iterator more generic.
func (i *PermissionIterator) Ent() (interface{}, error) {
	return i.Permission()
}

func (i *PermissionIterator) Permission() (*PermissionEntity, error) {
	var ent PermissionEntity
	cols, err := i.rows.Columns()
	if err != nil {
		return nil, err
	}

	props, err := ent.Props(cols...)
	if err != nil {
		return nil, err
	}
	if err := i.rows.Scan(props...); err != nil {
		return nil, err
	}
	return &ent, nil
}

type PermissionCriteria struct {
	Offset, Limit int64
	Sort          map[string]bool
	Action        *qtypes.String
	CreatedAt     *qtypes.Timestamp
	ID            *qtypes.Int64
	Module        *qtypes.String
	Subsystem     *qtypes.String
	UpdatedAt     *qtypes.Timestamp
}

type PermissionPatch struct {
	Action    ntypes.String
	CreatedAt pq.NullTime
	Module    ntypes.String
	Subsystem ntypes.String
	UpdatedAt pq.NullTime
}

func ScanPermissionRows(rows *sql.Rows) (entities []*PermissionEntity, err error) {
	for rows.Next() {
		var ent PermissionEntity
		err = rows.Scan(&ent.Action,
			&ent.CreatedAt,
			&ent.ID,
			&ent.Module,
			&ent.Subsystem,
			&ent.UpdatedAt,
		)
		if err != nil {
			return
		}

		entities = append(entities, &ent)
	}
	if err = rows.Err(); err != nil {
		return
	}

	return
}

type PermissionRepositoryBase struct {
	Table   string
	Columns []string
	DB      *sql.DB
	Debug   bool
	Log     log.Logger
}

func (r *PermissionRepositoryBase) InsertQuery(e *PermissionEntity) (string, []interface{}, error) {
	insert := NewComposer(6)
	columns := bytes.NewBuffer(nil)
	buf := bytes.NewBufferString("INSERT INTO ")
	buf.WriteString(r.Table)

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TablePermissionColumnAction); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.Action)
	insert.Dirty = true

	if !e.CreatedAt.IsZero() {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TablePermissionColumnCreatedAt); err != nil {
			return "", nil, err
		}
		if insert.Dirty {
			if _, err := insert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := insert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		insert.Add(e.CreatedAt)
		insert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TablePermissionColumnModule); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.Module)
	insert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TablePermissionColumnSubsystem); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.Subsystem)
	insert.Dirty = true

	if e.UpdatedAt.Valid {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TablePermissionColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if insert.Dirty {
			if _, err := insert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := insert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		insert.Add(e.UpdatedAt)
		insert.Dirty = true
	}

	if columns.Len() > 0 {
		buf.WriteString(" (")
		buf.ReadFrom(columns)
		buf.WriteString(") VALUES (")
		buf.ReadFrom(insert)
		buf.WriteString(") ")
		if len(r.Columns) > 0 {
			buf.WriteString("RETURNING ")
			buf.WriteString(strings.Join(r.Columns, ", "))
		}
	}
	return buf.String(), insert.Args(), nil
}

func (r *PermissionRepositoryBase) Insert(ctx context.Context, e *PermissionEntity) (*PermissionEntity, error) {
	query, args, err := r.InsertQuery(e)
	if err != nil {
		return nil, err
	}
	if err := r.DB.QueryRowContext(ctx, query, args...).Scan(&e.Action,
		&e.CreatedAt,
		&e.ID,
		&e.Module,
		&e.Subsystem,
		&e.UpdatedAt,
	); err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}
	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query success", "query", query, "table", r.Table)
	}
	return e, nil
}

func (r *PermissionRepositoryBase) FindQuery(s []string, c *PermissionCriteria) (string, []interface{}, error) {
	where := NewComposer(6)
	buf := bytes.NewBufferString("SELECT ")
	buf.WriteString(strings.Join(s, ", "))
	buf.WriteString(" FROM ")
	buf.WriteString(r.Table)
	buf.WriteString(" ")
	QueryStringWhereClause(c.Action, TablePermissionColumnAction, where, And)

	QueryTimestampWhereClause(c.CreatedAt, TablePermissionColumnCreatedAt, where, And)

	QueryInt64WhereClause(c.ID, TablePermissionColumnID, where, And)

	QueryStringWhereClause(c.Module, TablePermissionColumnModule, where, And)

	QueryStringWhereClause(c.Subsystem, TablePermissionColumnSubsystem, where, And)

	QueryTimestampWhereClause(c.UpdatedAt, TablePermissionColumnUpdatedAt, where, And)
	if where.Dirty {
		if _, err := buf.WriteString("WHERE "); err != nil {
			return "", nil, err
		}
		buf.ReadFrom(where)
	}

	if len(c.Sort) > 0 {
		i := 0
		where.WriteString(" ORDER BY ")

		for cn, asc := range c.Sort {
			for _, tcn := range TablePermissionColumns {
				if cn == tcn {
					if i > 0 {
						if _, err := where.WriteString(", "); err != nil {
							return "", nil, err
						}
					}
					if _, err := where.WriteString(cn); err != nil {
						return "", nil, err
					}
					if !asc {
						if _, err := where.WriteString(" DESC "); err != nil {
							return "", nil, err
						}
					}
					i++
					break
				}
			}
		}
	}
	if c.Offset > 0 {
		if _, err := where.WriteString(" OFFSET "); err != nil {
			return "", nil, err
		}
		if err := where.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		if _, err := where.WriteString(" "); err != nil {
			return "", nil, err
		}
		where.Add(c.Offset)
	}
	if c.Limit > 0 {
		if _, err := where.WriteString(" LIMIT "); err != nil {
			return "", nil, err
		}
		if err := where.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		if _, err := where.WriteString(" "); err != nil {
			return "", nil, err
		}
		where.Add(c.Limit)
	}

	buf.ReadFrom(where)

	return buf.String(), where.Args(), nil
}

func (r *PermissionRepositoryBase) Find(ctx context.Context, c *PermissionCriteria) ([]*PermissionEntity, error) {
	query, args, err := r.FindQuery(r.Columns, c)
	if err != nil {
		return nil, err
	}
	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "find query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}
	defer rows.Close()

	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "find query success", "query", query, "table", r.Table)
	}

	return ScanPermissionRows(rows)
}

func (r *PermissionRepositoryBase) FindIter(ctx context.Context, c *PermissionCriteria) (*PermissionIterator, error) {
	query, args, err := r.FindQuery(r.Columns, c)
	if err != nil {
		return nil, err
	}
	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &PermissionIterator{rows: rows}, nil
}
func (r *PermissionRepositoryBase) FindOneByID(ctx context.Context, pk int64) (*PermissionEntity, error) {
	find := NewComposer(6)
	find.WriteString("SELECT ")
	find.WriteString(strings.Join(r.Columns, ", "))
	find.WriteString(" FROM ")
	find.WriteString(TablePermission)
	find.WriteString(" WHERE ")
	find.WriteString(TablePermissionColumnID)
	find.WriteString("=")
	find.WritePlaceholder()
	find.Add(pk)
	var (
		ent PermissionEntity
	)
	props, err := ent.Props(r.Columns...)
	if err != nil {
		return nil, err
	}
	err = r.DB.QueryRowContext(ctx, find.String(), find.Args()...).Scan(props...)
	if err != nil {
		return nil, err
	}

	return &ent, nil
}
func (r *PermissionRepositoryBase) FindOneBySubsystemAndModuleAndAction(ctx context.Context, permissionSubsystem string, permissionModule string, permissionAction string) (*PermissionEntity, error) {
	find := NewComposer(6)
	find.WriteString("SELECT ")
	find.WriteString(strings.Join(r.Columns, ", "))
	find.WriteString(" FROM ")
	find.WriteString(TablePermission)
	find.WriteString(" WHERE ")
	find.WriteString(TablePermissionColumnSubsystem)
	find.WriteString("=")
	find.WritePlaceholder()
	find.Add(permissionSubsystem)
	find.WriteString(" AND ")
	find.WriteString(TablePermissionColumnModule)
	find.WriteString("=")
	find.WritePlaceholder()
	find.Add(permissionModule)
	find.WriteString(" AND ")
	find.WriteString(TablePermissionColumnAction)
	find.WriteString("=")
	find.WritePlaceholder()
	find.Add(permissionAction)

	var (
		ent PermissionEntity
	)
	props, err := ent.Props(r.Columns...)
	if err != nil {
		return nil, err
	}
	err = r.DB.QueryRowContext(ctx, find.String(), find.Args()...).Scan(props...)
	if err != nil {
		return nil, err
	}

	return &ent, nil
}
func (r *PermissionRepositoryBase) UpdateOneByIDQuery(pk int64, p *PermissionPatch) (string, []interface{}, error) {
	buf := bytes.NewBufferString("UPDATE ")
	buf.WriteString(r.Table)
	update := NewComposer(6)
	if p.Action.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TablePermissionColumnAction); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.Action)
		update.Dirty = true
	}

	if p.CreatedAt.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TablePermissionColumnCreatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.CreatedAt)
		update.Dirty = true

	}

	if p.Module.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TablePermissionColumnModule); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.Module)
		update.Dirty = true
	}

	if p.Subsystem.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TablePermissionColumnSubsystem); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.Subsystem)
		update.Dirty = true
	}

	if p.UpdatedAt.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TablePermissionColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.UpdatedAt)
		update.Dirty = true

	} else {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TablePermissionColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("=NOW()"); err != nil {
			return "", nil, err
		}
	}

	if !update.Dirty {
		return "", nil, errors.New("Permission update failure, nothing to update")
	}
	buf.WriteString(" SET ")
	buf.ReadFrom(update)
	buf.WriteString(" WHERE ")

	update.WriteString(TablePermissionColumnID)
	update.WriteString("=")
	update.WritePlaceholder()
	update.Add(pk)

	buf.ReadFrom(update)
	buf.WriteString(" RETURNING ")
	buf.WriteString(strings.Join(r.Columns, ", "))

	return buf.String(), update.Args(), nil
}
func (r *PermissionRepositoryBase) UpdateOneByID(ctx context.Context, pk int64, p *PermissionPatch) (*PermissionEntity, error) {
	query, args, err := r.UpdateOneByIDQuery(pk, p)
	if err != nil {
		return nil, err
	}
	var ent PermissionEntity
	props, err := ent.Props(r.Columns...)
	if err != nil {
		return nil, err
	}
	err = r.DB.QueryRowContext(ctx, query, args...).Scan(props...)
	if err != nil {
		return nil, err
	}

	return &ent, nil
}
func (r *PermissionRepositoryBase) UpdateOneBySubsystemAndModuleAndActionQuery(permissionSubsystem string, permissionModule string, permissionAction string, p *PermissionPatch) (string, []interface{}, error) {
	buf := bytes.NewBufferString("UPDATE ")
	buf.WriteString(r.Table)
	update := NewComposer(3)
	if p.Action.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TablePermissionColumnAction); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.Action)
		update.Dirty = true
	}

	if p.CreatedAt.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TablePermissionColumnCreatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.CreatedAt)
		update.Dirty = true

	}

	if p.Module.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TablePermissionColumnModule); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.Module)
		update.Dirty = true
	}

	if p.Subsystem.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TablePermissionColumnSubsystem); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.Subsystem)
		update.Dirty = true
	}

	if p.UpdatedAt.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TablePermissionColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.UpdatedAt)
		update.Dirty = true

	} else {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TablePermissionColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("=NOW()"); err != nil {
			return "", nil, err
		}
	}

	if !update.Dirty {
		return "", nil, errors.New("Permission update failure, nothing to update")
	}
	buf.WriteString(" SET ")
	buf.ReadFrom(update)
	buf.WriteString(" WHERE ")
	update.WriteString(TablePermissionColumnSubsystem)
	update.WriteString("=")
	update.WritePlaceholder()
	update.Add(permissionSubsystem)
	update.WriteString(" AND ")
	update.WriteString(TablePermissionColumnModule)
	update.WriteString("=")
	update.WritePlaceholder()
	update.Add(permissionModule)
	update.WriteString(" AND ")
	update.WriteString(TablePermissionColumnAction)
	update.WriteString("=")
	update.WritePlaceholder()
	update.Add(permissionAction)
	buf.ReadFrom(update)
	buf.WriteString(" RETURNING ")
	buf.WriteString(strings.Join(r.Columns, ", "))
	return buf.String(), update.Args(), nil
}
func (r *PermissionRepositoryBase) UpdateOneBySubsystemAndModuleAndAction(ctx context.Context, permissionSubsystem string, permissionModule string, permissionAction string, p *PermissionPatch) (*PermissionEntity, error) {
	query, args, err := r.UpdateOneBySubsystemAndModuleAndActionQuery(permissionSubsystem, permissionModule, permissionAction, p)
	if err != nil {
		return nil, err
	}
	var ent PermissionEntity
	props, err := ent.Props(r.Columns...)
	if err != nil {
		return nil, err
	}
	err = r.DB.QueryRowContext(ctx, query, args...).Scan(props...)
	if err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}

	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query success", "query", query, "table", r.Table)
	}

	return &ent, nil
}
func (r *PermissionRepositoryBase) UpsertQuery(e *PermissionEntity, p *PermissionPatch, inf ...string) (string, []interface{}, error) {
	upsert := NewComposer(12)
	columns := bytes.NewBuffer(nil)
	buf := bytes.NewBufferString("INSERT INTO ")
	buf.WriteString(r.Table)

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TablePermissionColumnAction); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.Action)
	upsert.Dirty = true

	if !e.CreatedAt.IsZero() {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TablePermissionColumnCreatedAt); err != nil {
			return "", nil, err
		}
		if upsert.Dirty {
			if _, err := upsert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := upsert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		upsert.Add(e.CreatedAt)
		upsert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TablePermissionColumnModule); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.Module)
	upsert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TablePermissionColumnSubsystem); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.Subsystem)
	upsert.Dirty = true

	if e.UpdatedAt.Valid {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TablePermissionColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if upsert.Dirty {
			if _, err := upsert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := upsert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		upsert.Add(e.UpdatedAt)
		upsert.Dirty = true
	}

	if upsert.Dirty {
		buf.WriteString(" (")
		buf.ReadFrom(columns)
		buf.WriteString(") VALUES (")
		buf.ReadFrom(upsert)
		buf.WriteString(")")
	}
	buf.WriteString(" ON CONFLICT ")
	if len(inf) > 0 {
		upsert.Dirty = false
		if p.Action.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TablePermissionColumnAction); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.Action)
			upsert.Dirty = true
		}

		if p.CreatedAt.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TablePermissionColumnCreatedAt); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.CreatedAt)
			upsert.Dirty = true

		}

		if p.Module.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TablePermissionColumnModule); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.Module)
			upsert.Dirty = true
		}

		if p.Subsystem.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TablePermissionColumnSubsystem); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.Subsystem)
			upsert.Dirty = true
		}

		if p.UpdatedAt.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TablePermissionColumnUpdatedAt); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.UpdatedAt)
			upsert.Dirty = true

		} else {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TablePermissionColumnUpdatedAt); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("=NOW()"); err != nil {
				return "", nil, err
			}
		}

	}

	if len(inf) > 0 && upsert.Dirty {
		buf.WriteString("(")
		for j, i := range inf {
			if j != 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(i)
		}
		buf.WriteString(")")
		buf.WriteString(" DO UPDATE SET ")
		buf.ReadFrom(upsert)
	} else {
		buf.WriteString(" DO NOTHING ")
	}
	if upsert.Dirty {
		if len(r.Columns) > 0 {
			buf.WriteString(" RETURNING ")
			buf.WriteString(strings.Join(r.Columns, ", "))
		}
	}
	return buf.String(), upsert.Args(), nil
}
func (r *PermissionRepositoryBase) Upsert(ctx context.Context, e *PermissionEntity, p *PermissionPatch, inf ...string) (*PermissionEntity, error) {
	query, args, err := r.UpsertQuery(e, p, inf...)
	if err != nil {
		return nil, err
	}
	if err := r.DB.QueryRowContext(ctx, query, args...).Scan(&e.Action,
		&e.CreatedAt,
		&e.ID,
		&e.Module,
		&e.Subsystem,
		&e.UpdatedAt,
	); err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "upsert query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}
	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "upsert query success", "query", query, "table", r.Table)
	}
	return e, nil
}

func (r *PermissionRepositoryBase) Count(ctx context.Context, c *PermissionCriteria) (int64, error) {
	query, args, err := r.FindQuery([]string{"COUNT(*)"}, c)
	if err != nil {
		return 0, err
	}
	var count int64
	if err := r.DB.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "count query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return 0, err
	}

	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "count query success", "query", query, "table", r.Table)
	}

	return count, nil
}
func (r *PermissionRepositoryBase) DeleteOneByID(ctx context.Context, pk int64) (int64, error) {
	find := NewComposer(6)
	find.WriteString("DELETE FROM ")
	find.WriteString(TablePermission)
	find.WriteString(" WHERE ")
	find.WriteString(TablePermissionColumnID)
	find.WriteString("=")
	find.WritePlaceholder()
	find.Add(pk)
	res, err := r.DB.ExecContext(ctx, find.String(), find.Args()...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

const (
	TableUserGroups                              = "charon.user_groups"
	TableUserGroupsColumnCreatedAt               = "created_at"
	TableUserGroupsColumnCreatedBy               = "created_by"
	TableUserGroupsColumnGroupID                 = "group_id"
	TableUserGroupsColumnUpdatedAt               = "updated_at"
	TableUserGroupsColumnUpdatedBy               = "updated_by"
	TableUserGroupsColumnUserID                  = "user_id"
	TableUserGroupsConstraintCreatedByForeignKey = "charon.user_groups_created_by_fkey"

	TableUserGroupsConstraintUpdatedByForeignKey = "charon.user_groups_updated_by_fkey"

	TableUserGroupsConstraintUserIDForeignKey = "charon.user_groups_user_id_fkey"

	TableUserGroupsConstraintGroupIDForeignKey = "charon.user_groups_group_id_fkey"

	TableUserGroupsConstraintUserIDGroupIDUnique = "charon.user_groups_user_id_group_id_key"
)

var (
	TableUserGroupsColumns = []string{
		TableUserGroupsColumnCreatedAt,
		TableUserGroupsColumnCreatedBy,
		TableUserGroupsColumnGroupID,
		TableUserGroupsColumnUpdatedAt,
		TableUserGroupsColumnUpdatedBy,
		TableUserGroupsColumnUserID,
	}
)

// UserGroupsEntity ...
type UserGroupsEntity struct {
	// CreatedAt ...
	CreatedAt time.Time
	// CreatedBy ...
	CreatedBy ntypes.Int64
	// GroupID ...
	GroupID int64
	// UpdatedAt ...
	UpdatedAt pq.NullTime
	// UpdatedBy ...
	UpdatedBy ntypes.Int64
	// UserID ...
	UserID int64
	// User ...
	User *UserEntity
	// Group ...
	Group *GroupEntity
	// Author ...
	Author *UserEntity
	// Modifier ...
	Modifier *UserEntity
}

func (e *UserGroupsEntity) Prop(cn string) (interface{}, bool) {
	switch cn {
	case TableUserGroupsColumnCreatedAt:
		return &e.CreatedAt, true
	case TableUserGroupsColumnCreatedBy:
		return &e.CreatedBy, true
	case TableUserGroupsColumnGroupID:
		return &e.GroupID, true
	case TableUserGroupsColumnUpdatedAt:
		return &e.UpdatedAt, true
	case TableUserGroupsColumnUpdatedBy:
		return &e.UpdatedBy, true
	case TableUserGroupsColumnUserID:
		return &e.UserID, true
	default:
		return nil, false
	}
}

func (e *UserGroupsEntity) Props(cns ...string) ([]interface{}, error) {
	res := make([]interface{}, 0, len(cns))
	for _, cn := range cns {
		if prop, ok := e.Prop(cn); ok {
			res = append(res, prop)
		} else {
			return nil, fmt.Errorf("unexpected column provided: %s", cn)
		}
	}
	return res, nil
}

// UserGroupsIterator is not thread safe.
type UserGroupsIterator struct {
	rows *sql.Rows
	cols []string
}

func (i *UserGroupsIterator) Next() bool {
	return i.rows.Next()
}

func (i *UserGroupsIterator) Close() error {
	return i.rows.Close()
}

func (i *UserGroupsIterator) Err() error {
	return i.rows.Err()
}

// Columns is wrapper around sql.Rows.Columns method, that also cache outpu inside iterator.
func (i *UserGroupsIterator) Columns() ([]string, error) {
	if i.cols == nil {
		cols, err := i.rows.Columns()
		if err != nil {
			return nil, err
		}
		i.cols = cols
	}
	return i.cols, nil
}

// Ent is wrapper around UserGroups method that makes iterator more generic.
func (i *UserGroupsIterator) Ent() (interface{}, error) {
	return i.UserGroups()
}

func (i *UserGroupsIterator) UserGroups() (*UserGroupsEntity, error) {
	var ent UserGroupsEntity
	cols, err := i.rows.Columns()
	if err != nil {
		return nil, err
	}

	props, err := ent.Props(cols...)
	if err != nil {
		return nil, err
	}
	if err := i.rows.Scan(props...); err != nil {
		return nil, err
	}
	return &ent, nil
}

type UserGroupsCriteria struct {
	Offset, Limit int64
	Sort          map[string]bool
	CreatedAt     *qtypes.Timestamp
	CreatedBy     *qtypes.Int64
	GroupID       *qtypes.Int64
	UpdatedAt     *qtypes.Timestamp
	UpdatedBy     *qtypes.Int64
	UserID        *qtypes.Int64
}

type UserGroupsPatch struct {
	CreatedAt pq.NullTime
	CreatedBy ntypes.Int64
	GroupID   ntypes.Int64
	UpdatedAt pq.NullTime
	UpdatedBy ntypes.Int64
	UserID    ntypes.Int64
}

func ScanUserGroupsRows(rows *sql.Rows) (entities []*UserGroupsEntity, err error) {
	for rows.Next() {
		var ent UserGroupsEntity
		err = rows.Scan(&ent.CreatedAt,
			&ent.CreatedBy,
			&ent.GroupID,
			&ent.UpdatedAt,
			&ent.UpdatedBy,
			&ent.UserID,
		)
		if err != nil {
			return
		}

		entities = append(entities, &ent)
	}
	if err = rows.Err(); err != nil {
		return
	}

	return
}

type UserGroupsRepositoryBase struct {
	Table   string
	Columns []string
	DB      *sql.DB
	Debug   bool
	Log     log.Logger
}

func (r *UserGroupsRepositoryBase) InsertQuery(e *UserGroupsEntity) (string, []interface{}, error) {
	insert := NewComposer(6)
	columns := bytes.NewBuffer(nil)
	buf := bytes.NewBufferString("INSERT INTO ")
	buf.WriteString(r.Table)

	if !e.CreatedAt.IsZero() {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableUserGroupsColumnCreatedAt); err != nil {
			return "", nil, err
		}
		if insert.Dirty {
			if _, err := insert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := insert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		insert.Add(e.CreatedAt)
		insert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserGroupsColumnCreatedBy); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.CreatedBy)
	insert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserGroupsColumnGroupID); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.GroupID)
	insert.Dirty = true

	if e.UpdatedAt.Valid {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableUserGroupsColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if insert.Dirty {
			if _, err := insert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := insert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		insert.Add(e.UpdatedAt)
		insert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserGroupsColumnUpdatedBy); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.UpdatedBy)
	insert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserGroupsColumnUserID); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.UserID)
	insert.Dirty = true
	if columns.Len() > 0 {
		buf.WriteString(" (")
		buf.ReadFrom(columns)
		buf.WriteString(") VALUES (")
		buf.ReadFrom(insert)
		buf.WriteString(") ")
		if len(r.Columns) > 0 {
			buf.WriteString("RETURNING ")
			buf.WriteString(strings.Join(r.Columns, ", "))
		}
	}
	return buf.String(), insert.Args(), nil
}

func (r *UserGroupsRepositoryBase) Insert(ctx context.Context, e *UserGroupsEntity) (*UserGroupsEntity, error) {
	query, args, err := r.InsertQuery(e)
	if err != nil {
		return nil, err
	}
	if err := r.DB.QueryRowContext(ctx, query, args...).Scan(&e.CreatedAt,
		&e.CreatedBy,
		&e.GroupID,
		&e.UpdatedAt,
		&e.UpdatedBy,
		&e.UserID,
	); err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}
	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query success", "query", query, "table", r.Table)
	}
	return e, nil
}

func (r *UserGroupsRepositoryBase) FindQuery(s []string, c *UserGroupsCriteria) (string, []interface{}, error) {
	where := NewComposer(6)
	buf := bytes.NewBufferString("SELECT ")
	buf.WriteString(strings.Join(s, ", "))
	buf.WriteString(" FROM ")
	buf.WriteString(r.Table)
	buf.WriteString(" ")
	QueryTimestampWhereClause(c.CreatedAt, TableUserGroupsColumnCreatedAt, where, And)

	QueryInt64WhereClause(c.CreatedBy, TableUserGroupsColumnCreatedBy, where, And)

	QueryInt64WhereClause(c.GroupID, TableUserGroupsColumnGroupID, where, And)

	QueryTimestampWhereClause(c.UpdatedAt, TableUserGroupsColumnUpdatedAt, where, And)

	QueryInt64WhereClause(c.UpdatedBy, TableUserGroupsColumnUpdatedBy, where, And)

	QueryInt64WhereClause(c.UserID, TableUserGroupsColumnUserID, where, And)
	if where.Dirty {
		if _, err := buf.WriteString("WHERE "); err != nil {
			return "", nil, err
		}
		buf.ReadFrom(where)
	}

	if len(c.Sort) > 0 {
		i := 0
		where.WriteString(" ORDER BY ")

		for cn, asc := range c.Sort {
			for _, tcn := range TableUserGroupsColumns {
				if cn == tcn {
					if i > 0 {
						if _, err := where.WriteString(", "); err != nil {
							return "", nil, err
						}
					}
					if _, err := where.WriteString(cn); err != nil {
						return "", nil, err
					}
					if !asc {
						if _, err := where.WriteString(" DESC "); err != nil {
							return "", nil, err
						}
					}
					i++
					break
				}
			}
		}
	}
	if c.Offset > 0 {
		if _, err := where.WriteString(" OFFSET "); err != nil {
			return "", nil, err
		}
		if err := where.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		if _, err := where.WriteString(" "); err != nil {
			return "", nil, err
		}
		where.Add(c.Offset)
	}
	if c.Limit > 0 {
		if _, err := where.WriteString(" LIMIT "); err != nil {
			return "", nil, err
		}
		if err := where.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		if _, err := where.WriteString(" "); err != nil {
			return "", nil, err
		}
		where.Add(c.Limit)
	}

	buf.ReadFrom(where)

	return buf.String(), where.Args(), nil
}

func (r *UserGroupsRepositoryBase) Find(ctx context.Context, c *UserGroupsCriteria) ([]*UserGroupsEntity, error) {
	query, args, err := r.FindQuery(r.Columns, c)
	if err != nil {
		return nil, err
	}
	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "find query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}
	defer rows.Close()

	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "find query success", "query", query, "table", r.Table)
	}

	return ScanUserGroupsRows(rows)
}

func (r *UserGroupsRepositoryBase) FindIter(ctx context.Context, c *UserGroupsCriteria) (*UserGroupsIterator, error) {
	query, args, err := r.FindQuery(r.Columns, c)
	if err != nil {
		return nil, err
	}
	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &UserGroupsIterator{rows: rows}, nil
}
func (r *UserGroupsRepositoryBase) FindOneByUserIDAndGroupID(ctx context.Context, userGroupsUserID int64, userGroupsGroupID int64) (*UserGroupsEntity, error) {
	find := NewComposer(6)
	find.WriteString("SELECT ")
	find.WriteString(strings.Join(r.Columns, ", "))
	find.WriteString(" FROM ")
	find.WriteString(TableUserGroups)
	find.WriteString(" WHERE ")
	find.WriteString(TableUserGroupsColumnUserID)
	find.WriteString("=")
	find.WritePlaceholder()
	find.Add(userGroupsUserID)
	find.WriteString(" AND ")
	find.WriteString(TableUserGroupsColumnGroupID)
	find.WriteString("=")
	find.WritePlaceholder()
	find.Add(userGroupsGroupID)

	var (
		ent UserGroupsEntity
	)
	props, err := ent.Props(r.Columns...)
	if err != nil {
		return nil, err
	}
	err = r.DB.QueryRowContext(ctx, find.String(), find.Args()...).Scan(props...)
	if err != nil {
		return nil, err
	}

	return &ent, nil
}
func (r *UserGroupsRepositoryBase) UpdateOneByUserIDAndGroupIDQuery(userGroupsUserID int64, userGroupsGroupID int64, p *UserGroupsPatch) (string, []interface{}, error) {
	buf := bytes.NewBufferString("UPDATE ")
	buf.WriteString(r.Table)
	update := NewComposer(2)
	if p.CreatedAt.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserGroupsColumnCreatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.CreatedAt)
		update.Dirty = true

	}

	if p.CreatedBy.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserGroupsColumnCreatedBy); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.CreatedBy)
		update.Dirty = true
	}

	if p.GroupID.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserGroupsColumnGroupID); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.GroupID)
		update.Dirty = true
	}

	if p.UpdatedAt.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserGroupsColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.UpdatedAt)
		update.Dirty = true

	} else {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserGroupsColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("=NOW()"); err != nil {
			return "", nil, err
		}
	}

	if p.UpdatedBy.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserGroupsColumnUpdatedBy); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.UpdatedBy)
		update.Dirty = true
	}

	if p.UserID.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserGroupsColumnUserID); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.UserID)
		update.Dirty = true
	}

	if !update.Dirty {
		return "", nil, errors.New("UserGroups update failure, nothing to update")
	}
	buf.WriteString(" SET ")
	buf.ReadFrom(update)
	buf.WriteString(" WHERE ")
	update.WriteString(TableUserGroupsColumnUserID)
	update.WriteString("=")
	update.WritePlaceholder()
	update.Add(userGroupsUserID)
	update.WriteString(" AND ")
	update.WriteString(TableUserGroupsColumnGroupID)
	update.WriteString("=")
	update.WritePlaceholder()
	update.Add(userGroupsGroupID)
	buf.ReadFrom(update)
	buf.WriteString(" RETURNING ")
	buf.WriteString(strings.Join(r.Columns, ", "))
	return buf.String(), update.Args(), nil
}
func (r *UserGroupsRepositoryBase) UpdateOneByUserIDAndGroupID(ctx context.Context, userGroupsUserID int64, userGroupsGroupID int64, p *UserGroupsPatch) (*UserGroupsEntity, error) {
	query, args, err := r.UpdateOneByUserIDAndGroupIDQuery(userGroupsUserID, userGroupsGroupID, p)
	if err != nil {
		return nil, err
	}
	var ent UserGroupsEntity
	props, err := ent.Props(r.Columns...)
	if err != nil {
		return nil, err
	}
	err = r.DB.QueryRowContext(ctx, query, args...).Scan(props...)
	if err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}

	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query success", "query", query, "table", r.Table)
	}

	return &ent, nil
}
func (r *UserGroupsRepositoryBase) UpsertQuery(e *UserGroupsEntity, p *UserGroupsPatch, inf ...string) (string, []interface{}, error) {
	upsert := NewComposer(12)
	columns := bytes.NewBuffer(nil)
	buf := bytes.NewBufferString("INSERT INTO ")
	buf.WriteString(r.Table)

	if !e.CreatedAt.IsZero() {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableUserGroupsColumnCreatedAt); err != nil {
			return "", nil, err
		}
		if upsert.Dirty {
			if _, err := upsert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := upsert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		upsert.Add(e.CreatedAt)
		upsert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserGroupsColumnCreatedBy); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.CreatedBy)
	upsert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserGroupsColumnGroupID); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.GroupID)
	upsert.Dirty = true

	if e.UpdatedAt.Valid {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableUserGroupsColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if upsert.Dirty {
			if _, err := upsert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := upsert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		upsert.Add(e.UpdatedAt)
		upsert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserGroupsColumnUpdatedBy); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.UpdatedBy)
	upsert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserGroupsColumnUserID); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.UserID)
	upsert.Dirty = true

	if upsert.Dirty {
		buf.WriteString(" (")
		buf.ReadFrom(columns)
		buf.WriteString(") VALUES (")
		buf.ReadFrom(upsert)
		buf.WriteString(")")
	}
	buf.WriteString(" ON CONFLICT ")
	if len(inf) > 0 {
		upsert.Dirty = false
		if p.CreatedAt.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserGroupsColumnCreatedAt); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.CreatedAt)
			upsert.Dirty = true

		}

		if p.CreatedBy.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserGroupsColumnCreatedBy); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.CreatedBy)
			upsert.Dirty = true
		}

		if p.GroupID.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserGroupsColumnGroupID); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.GroupID)
			upsert.Dirty = true
		}

		if p.UpdatedAt.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserGroupsColumnUpdatedAt); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.UpdatedAt)
			upsert.Dirty = true

		} else {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserGroupsColumnUpdatedAt); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("=NOW()"); err != nil {
				return "", nil, err
			}
		}

		if p.UpdatedBy.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserGroupsColumnUpdatedBy); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.UpdatedBy)
			upsert.Dirty = true
		}

		if p.UserID.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserGroupsColumnUserID); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.UserID)
			upsert.Dirty = true
		}

	}

	if len(inf) > 0 && upsert.Dirty {
		buf.WriteString("(")
		for j, i := range inf {
			if j != 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(i)
		}
		buf.WriteString(")")
		buf.WriteString(" DO UPDATE SET ")
		buf.ReadFrom(upsert)
	} else {
		buf.WriteString(" DO NOTHING ")
	}
	if upsert.Dirty {
		if len(r.Columns) > 0 {
			buf.WriteString(" RETURNING ")
			buf.WriteString(strings.Join(r.Columns, ", "))
		}
	}
	return buf.String(), upsert.Args(), nil
}
func (r *UserGroupsRepositoryBase) Upsert(ctx context.Context, e *UserGroupsEntity, p *UserGroupsPatch, inf ...string) (*UserGroupsEntity, error) {
	query, args, err := r.UpsertQuery(e, p, inf...)
	if err != nil {
		return nil, err
	}
	if err := r.DB.QueryRowContext(ctx, query, args...).Scan(&e.CreatedAt,
		&e.CreatedBy,
		&e.GroupID,
		&e.UpdatedAt,
		&e.UpdatedBy,
		&e.UserID,
	); err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "upsert query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}
	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "upsert query success", "query", query, "table", r.Table)
	}
	return e, nil
}

func (r *UserGroupsRepositoryBase) Count(ctx context.Context, c *UserGroupsCriteria) (int64, error) {
	query, args, err := r.FindQuery([]string{"COUNT(*)"}, c)
	if err != nil {
		return 0, err
	}
	var count int64
	if err := r.DB.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "count query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return 0, err
	}

	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "count query success", "query", query, "table", r.Table)
	}

	return count, nil
}

const (
	TableGroupPermissions                              = "charon.group_permissions"
	TableGroupPermissionsColumnCreatedAt               = "created_at"
	TableGroupPermissionsColumnCreatedBy               = "created_by"
	TableGroupPermissionsColumnGroupID                 = "group_id"
	TableGroupPermissionsColumnPermissionAction        = "permission_action"
	TableGroupPermissionsColumnPermissionModule        = "permission_module"
	TableGroupPermissionsColumnPermissionSubsystem     = "permission_subsystem"
	TableGroupPermissionsColumnUpdatedAt               = "updated_at"
	TableGroupPermissionsColumnUpdatedBy               = "updated_by"
	TableGroupPermissionsConstraintCreatedByForeignKey = "charon.group_permissions_created_by_fkey"

	TableGroupPermissionsConstraintUpdatedByForeignKey = "charon.group_permissions_updated_by_fkey"

	TableGroupPermissionsConstraintGroupIDForeignKey = "charon.group_permissions_group_id_fkey"

	TableGroupPermissionsConstraintPermissionSubsystemPermissionModulePermissionActionForeignKey = "charon.group_permissions_subsystem_module_action_fkey"

	TableGroupPermissionsConstraintGroupIDPermissionSubsystemPermissionModulePermissionActionUnique = "charon.group_permissions_group_id_subsystem_module_action_key"
)

var (
	TableGroupPermissionsColumns = []string{
		TableGroupPermissionsColumnCreatedAt,
		TableGroupPermissionsColumnCreatedBy,
		TableGroupPermissionsColumnGroupID,
		TableGroupPermissionsColumnPermissionAction,
		TableGroupPermissionsColumnPermissionModule,
		TableGroupPermissionsColumnPermissionSubsystem,
		TableGroupPermissionsColumnUpdatedAt,
		TableGroupPermissionsColumnUpdatedBy,
	}
)

// GroupPermissionsEntity ...
type GroupPermissionsEntity struct {
	// CreatedAt ...
	CreatedAt time.Time
	// CreatedBy ...
	CreatedBy ntypes.Int64
	// GroupID ...
	GroupID int64
	// PermissionAction ...
	PermissionAction string
	// PermissionModule ...
	PermissionModule string
	// PermissionSubsystem ...
	PermissionSubsystem string
	// UpdatedAt ...
	UpdatedAt pq.NullTime
	// UpdatedBy ...
	UpdatedBy ntypes.Int64
	// Group ...
	Group *GroupEntity
	// Author ...
	Author *UserEntity
	// Modifier ...
	Modifier *UserEntity
}

func (e *GroupPermissionsEntity) Prop(cn string) (interface{}, bool) {
	switch cn {
	case TableGroupPermissionsColumnCreatedAt:
		return &e.CreatedAt, true
	case TableGroupPermissionsColumnCreatedBy:
		return &e.CreatedBy, true
	case TableGroupPermissionsColumnGroupID:
		return &e.GroupID, true
	case TableGroupPermissionsColumnPermissionAction:
		return &e.PermissionAction, true
	case TableGroupPermissionsColumnPermissionModule:
		return &e.PermissionModule, true
	case TableGroupPermissionsColumnPermissionSubsystem:
		return &e.PermissionSubsystem, true
	case TableGroupPermissionsColumnUpdatedAt:
		return &e.UpdatedAt, true
	case TableGroupPermissionsColumnUpdatedBy:
		return &e.UpdatedBy, true
	default:
		return nil, false
	}
}

func (e *GroupPermissionsEntity) Props(cns ...string) ([]interface{}, error) {
	res := make([]interface{}, 0, len(cns))
	for _, cn := range cns {
		if prop, ok := e.Prop(cn); ok {
			res = append(res, prop)
		} else {
			return nil, fmt.Errorf("unexpected column provided: %s", cn)
		}
	}
	return res, nil
}

// GroupPermissionsIterator is not thread safe.
type GroupPermissionsIterator struct {
	rows *sql.Rows
	cols []string
}

func (i *GroupPermissionsIterator) Next() bool {
	return i.rows.Next()
}

func (i *GroupPermissionsIterator) Close() error {
	return i.rows.Close()
}

func (i *GroupPermissionsIterator) Err() error {
	return i.rows.Err()
}

// Columns is wrapper around sql.Rows.Columns method, that also cache outpu inside iterator.
func (i *GroupPermissionsIterator) Columns() ([]string, error) {
	if i.cols == nil {
		cols, err := i.rows.Columns()
		if err != nil {
			return nil, err
		}
		i.cols = cols
	}
	return i.cols, nil
}

// Ent is wrapper around GroupPermissions method that makes iterator more generic.
func (i *GroupPermissionsIterator) Ent() (interface{}, error) {
	return i.GroupPermissions()
}

func (i *GroupPermissionsIterator) GroupPermissions() (*GroupPermissionsEntity, error) {
	var ent GroupPermissionsEntity
	cols, err := i.rows.Columns()
	if err != nil {
		return nil, err
	}

	props, err := ent.Props(cols...)
	if err != nil {
		return nil, err
	}
	if err := i.rows.Scan(props...); err != nil {
		return nil, err
	}
	return &ent, nil
}

type GroupPermissionsCriteria struct {
	Offset, Limit       int64
	Sort                map[string]bool
	CreatedAt           *qtypes.Timestamp
	CreatedBy           *qtypes.Int64
	GroupID             *qtypes.Int64
	PermissionAction    *qtypes.String
	PermissionModule    *qtypes.String
	PermissionSubsystem *qtypes.String
	UpdatedAt           *qtypes.Timestamp
	UpdatedBy           *qtypes.Int64
}

type GroupPermissionsPatch struct {
	CreatedAt           pq.NullTime
	CreatedBy           ntypes.Int64
	GroupID             ntypes.Int64
	PermissionAction    ntypes.String
	PermissionModule    ntypes.String
	PermissionSubsystem ntypes.String
	UpdatedAt           pq.NullTime
	UpdatedBy           ntypes.Int64
}

func ScanGroupPermissionsRows(rows *sql.Rows) (entities []*GroupPermissionsEntity, err error) {
	for rows.Next() {
		var ent GroupPermissionsEntity
		err = rows.Scan(&ent.CreatedAt,
			&ent.CreatedBy,
			&ent.GroupID,
			&ent.PermissionAction,
			&ent.PermissionModule,
			&ent.PermissionSubsystem,
			&ent.UpdatedAt,
			&ent.UpdatedBy,
		)
		if err != nil {
			return
		}

		entities = append(entities, &ent)
	}
	if err = rows.Err(); err != nil {
		return
	}

	return
}

type GroupPermissionsRepositoryBase struct {
	Table   string
	Columns []string
	DB      *sql.DB
	Debug   bool
	Log     log.Logger
}

func (r *GroupPermissionsRepositoryBase) InsertQuery(e *GroupPermissionsEntity) (string, []interface{}, error) {
	insert := NewComposer(8)
	columns := bytes.NewBuffer(nil)
	buf := bytes.NewBufferString("INSERT INTO ")
	buf.WriteString(r.Table)

	if !e.CreatedAt.IsZero() {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableGroupPermissionsColumnCreatedAt); err != nil {
			return "", nil, err
		}
		if insert.Dirty {
			if _, err := insert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := insert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		insert.Add(e.CreatedAt)
		insert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableGroupPermissionsColumnCreatedBy); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.CreatedBy)
	insert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableGroupPermissionsColumnGroupID); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.GroupID)
	insert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableGroupPermissionsColumnPermissionAction); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.PermissionAction)
	insert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableGroupPermissionsColumnPermissionModule); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.PermissionModule)
	insert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableGroupPermissionsColumnPermissionSubsystem); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.PermissionSubsystem)
	insert.Dirty = true

	if e.UpdatedAt.Valid {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableGroupPermissionsColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if insert.Dirty {
			if _, err := insert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := insert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		insert.Add(e.UpdatedAt)
		insert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableGroupPermissionsColumnUpdatedBy); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.UpdatedBy)
	insert.Dirty = true
	if columns.Len() > 0 {
		buf.WriteString(" (")
		buf.ReadFrom(columns)
		buf.WriteString(") VALUES (")
		buf.ReadFrom(insert)
		buf.WriteString(") ")
		if len(r.Columns) > 0 {
			buf.WriteString("RETURNING ")
			buf.WriteString(strings.Join(r.Columns, ", "))
		}
	}
	return buf.String(), insert.Args(), nil
}

func (r *GroupPermissionsRepositoryBase) Insert(ctx context.Context, e *GroupPermissionsEntity) (*GroupPermissionsEntity, error) {
	query, args, err := r.InsertQuery(e)
	if err != nil {
		return nil, err
	}
	if err := r.DB.QueryRowContext(ctx, query, args...).Scan(&e.CreatedAt,
		&e.CreatedBy,
		&e.GroupID,
		&e.PermissionAction,
		&e.PermissionModule,
		&e.PermissionSubsystem,
		&e.UpdatedAt,
		&e.UpdatedBy,
	); err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}
	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query success", "query", query, "table", r.Table)
	}
	return e, nil
}

func (r *GroupPermissionsRepositoryBase) FindQuery(s []string, c *GroupPermissionsCriteria) (string, []interface{}, error) {
	where := NewComposer(8)
	buf := bytes.NewBufferString("SELECT ")
	buf.WriteString(strings.Join(s, ", "))
	buf.WriteString(" FROM ")
	buf.WriteString(r.Table)
	buf.WriteString(" ")
	QueryTimestampWhereClause(c.CreatedAt, TableGroupPermissionsColumnCreatedAt, where, And)

	QueryInt64WhereClause(c.CreatedBy, TableGroupPermissionsColumnCreatedBy, where, And)

	QueryInt64WhereClause(c.GroupID, TableGroupPermissionsColumnGroupID, where, And)

	QueryStringWhereClause(c.PermissionAction, TableGroupPermissionsColumnPermissionAction, where, And)

	QueryStringWhereClause(c.PermissionModule, TableGroupPermissionsColumnPermissionModule, where, And)

	QueryStringWhereClause(c.PermissionSubsystem, TableGroupPermissionsColumnPermissionSubsystem, where, And)

	QueryTimestampWhereClause(c.UpdatedAt, TableGroupPermissionsColumnUpdatedAt, where, And)

	QueryInt64WhereClause(c.UpdatedBy, TableGroupPermissionsColumnUpdatedBy, where, And)
	if where.Dirty {
		if _, err := buf.WriteString("WHERE "); err != nil {
			return "", nil, err
		}
		buf.ReadFrom(where)
	}

	if len(c.Sort) > 0 {
		i := 0
		where.WriteString(" ORDER BY ")

		for cn, asc := range c.Sort {
			for _, tcn := range TableGroupPermissionsColumns {
				if cn == tcn {
					if i > 0 {
						if _, err := where.WriteString(", "); err != nil {
							return "", nil, err
						}
					}
					if _, err := where.WriteString(cn); err != nil {
						return "", nil, err
					}
					if !asc {
						if _, err := where.WriteString(" DESC "); err != nil {
							return "", nil, err
						}
					}
					i++
					break
				}
			}
		}
	}
	if c.Offset > 0 {
		if _, err := where.WriteString(" OFFSET "); err != nil {
			return "", nil, err
		}
		if err := where.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		if _, err := where.WriteString(" "); err != nil {
			return "", nil, err
		}
		where.Add(c.Offset)
	}
	if c.Limit > 0 {
		if _, err := where.WriteString(" LIMIT "); err != nil {
			return "", nil, err
		}
		if err := where.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		if _, err := where.WriteString(" "); err != nil {
			return "", nil, err
		}
		where.Add(c.Limit)
	}

	buf.ReadFrom(where)

	return buf.String(), where.Args(), nil
}

func (r *GroupPermissionsRepositoryBase) Find(ctx context.Context, c *GroupPermissionsCriteria) ([]*GroupPermissionsEntity, error) {
	query, args, err := r.FindQuery(r.Columns, c)
	if err != nil {
		return nil, err
	}
	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "find query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}
	defer rows.Close()

	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "find query success", "query", query, "table", r.Table)
	}

	return ScanGroupPermissionsRows(rows)
}

func (r *GroupPermissionsRepositoryBase) FindIter(ctx context.Context, c *GroupPermissionsCriteria) (*GroupPermissionsIterator, error) {
	query, args, err := r.FindQuery(r.Columns, c)
	if err != nil {
		return nil, err
	}
	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &GroupPermissionsIterator{rows: rows}, nil
}
func (r *GroupPermissionsRepositoryBase) FindOneByGroupIDAndPermissionSubsystemAndPermissionModuleAndPermissionAction(ctx context.Context, groupPermissionsGroupID int64, groupPermissionsPermissionSubsystem string, groupPermissionsPermissionModule string, groupPermissionsPermissionAction string) (*GroupPermissionsEntity, error) {
	find := NewComposer(8)
	find.WriteString("SELECT ")
	find.WriteString(strings.Join(r.Columns, ", "))
	find.WriteString(" FROM ")
	find.WriteString(TableGroupPermissions)
	find.WriteString(" WHERE ")
	find.WriteString(TableGroupPermissionsColumnGroupID)
	find.WriteString("=")
	find.WritePlaceholder()
	find.Add(groupPermissionsGroupID)
	find.WriteString(" AND ")
	find.WriteString(TableGroupPermissionsColumnPermissionSubsystem)
	find.WriteString("=")
	find.WritePlaceholder()
	find.Add(groupPermissionsPermissionSubsystem)
	find.WriteString(" AND ")
	find.WriteString(TableGroupPermissionsColumnPermissionModule)
	find.WriteString("=")
	find.WritePlaceholder()
	find.Add(groupPermissionsPermissionModule)
	find.WriteString(" AND ")
	find.WriteString(TableGroupPermissionsColumnPermissionAction)
	find.WriteString("=")
	find.WritePlaceholder()
	find.Add(groupPermissionsPermissionAction)

	var (
		ent GroupPermissionsEntity
	)
	props, err := ent.Props(r.Columns...)
	if err != nil {
		return nil, err
	}
	err = r.DB.QueryRowContext(ctx, find.String(), find.Args()...).Scan(props...)
	if err != nil {
		return nil, err
	}

	return &ent, nil
}
func (r *GroupPermissionsRepositoryBase) UpdateOneByGroupIDAndPermissionSubsystemAndPermissionModuleAndPermissionActionQuery(groupPermissionsGroupID int64, groupPermissionsPermissionSubsystem string, groupPermissionsPermissionModule string, groupPermissionsPermissionAction string, p *GroupPermissionsPatch) (string, []interface{}, error) {
	buf := bytes.NewBufferString("UPDATE ")
	buf.WriteString(r.Table)
	update := NewComposer(4)
	if p.CreatedAt.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupPermissionsColumnCreatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.CreatedAt)
		update.Dirty = true

	}

	if p.CreatedBy.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupPermissionsColumnCreatedBy); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.CreatedBy)
		update.Dirty = true
	}

	if p.GroupID.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupPermissionsColumnGroupID); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.GroupID)
		update.Dirty = true
	}

	if p.PermissionAction.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupPermissionsColumnPermissionAction); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.PermissionAction)
		update.Dirty = true
	}

	if p.PermissionModule.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupPermissionsColumnPermissionModule); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.PermissionModule)
		update.Dirty = true
	}

	if p.PermissionSubsystem.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupPermissionsColumnPermissionSubsystem); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.PermissionSubsystem)
		update.Dirty = true
	}

	if p.UpdatedAt.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupPermissionsColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.UpdatedAt)
		update.Dirty = true

	} else {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupPermissionsColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("=NOW()"); err != nil {
			return "", nil, err
		}
	}

	if p.UpdatedBy.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableGroupPermissionsColumnUpdatedBy); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.UpdatedBy)
		update.Dirty = true
	}

	if !update.Dirty {
		return "", nil, errors.New("GroupPermissions update failure, nothing to update")
	}
	buf.WriteString(" SET ")
	buf.ReadFrom(update)
	buf.WriteString(" WHERE ")
	update.WriteString(TableGroupPermissionsColumnGroupID)
	update.WriteString("=")
	update.WritePlaceholder()
	update.Add(groupPermissionsGroupID)
	update.WriteString(" AND ")
	update.WriteString(TableGroupPermissionsColumnPermissionSubsystem)
	update.WriteString("=")
	update.WritePlaceholder()
	update.Add(groupPermissionsPermissionSubsystem)
	update.WriteString(" AND ")
	update.WriteString(TableGroupPermissionsColumnPermissionModule)
	update.WriteString("=")
	update.WritePlaceholder()
	update.Add(groupPermissionsPermissionModule)
	update.WriteString(" AND ")
	update.WriteString(TableGroupPermissionsColumnPermissionAction)
	update.WriteString("=")
	update.WritePlaceholder()
	update.Add(groupPermissionsPermissionAction)
	buf.ReadFrom(update)
	buf.WriteString(" RETURNING ")
	buf.WriteString(strings.Join(r.Columns, ", "))
	return buf.String(), update.Args(), nil
}
func (r *GroupPermissionsRepositoryBase) UpdateOneByGroupIDAndPermissionSubsystemAndPermissionModuleAndPermissionAction(ctx context.Context, groupPermissionsGroupID int64, groupPermissionsPermissionSubsystem string, groupPermissionsPermissionModule string, groupPermissionsPermissionAction string, p *GroupPermissionsPatch) (*GroupPermissionsEntity, error) {
	query, args, err := r.UpdateOneByGroupIDAndPermissionSubsystemAndPermissionModuleAndPermissionActionQuery(groupPermissionsGroupID, groupPermissionsPermissionSubsystem, groupPermissionsPermissionModule, groupPermissionsPermissionAction, p)
	if err != nil {
		return nil, err
	}
	var ent GroupPermissionsEntity
	props, err := ent.Props(r.Columns...)
	if err != nil {
		return nil, err
	}
	err = r.DB.QueryRowContext(ctx, query, args...).Scan(props...)
	if err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}

	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query success", "query", query, "table", r.Table)
	}

	return &ent, nil
}
func (r *GroupPermissionsRepositoryBase) UpsertQuery(e *GroupPermissionsEntity, p *GroupPermissionsPatch, inf ...string) (string, []interface{}, error) {
	upsert := NewComposer(16)
	columns := bytes.NewBuffer(nil)
	buf := bytes.NewBufferString("INSERT INTO ")
	buf.WriteString(r.Table)

	if !e.CreatedAt.IsZero() {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableGroupPermissionsColumnCreatedAt); err != nil {
			return "", nil, err
		}
		if upsert.Dirty {
			if _, err := upsert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := upsert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		upsert.Add(e.CreatedAt)
		upsert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableGroupPermissionsColumnCreatedBy); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.CreatedBy)
	upsert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableGroupPermissionsColumnGroupID); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.GroupID)
	upsert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableGroupPermissionsColumnPermissionAction); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.PermissionAction)
	upsert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableGroupPermissionsColumnPermissionModule); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.PermissionModule)
	upsert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableGroupPermissionsColumnPermissionSubsystem); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.PermissionSubsystem)
	upsert.Dirty = true

	if e.UpdatedAt.Valid {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableGroupPermissionsColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if upsert.Dirty {
			if _, err := upsert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := upsert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		upsert.Add(e.UpdatedAt)
		upsert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableGroupPermissionsColumnUpdatedBy); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.UpdatedBy)
	upsert.Dirty = true

	if upsert.Dirty {
		buf.WriteString(" (")
		buf.ReadFrom(columns)
		buf.WriteString(") VALUES (")
		buf.ReadFrom(upsert)
		buf.WriteString(")")
	}
	buf.WriteString(" ON CONFLICT ")
	if len(inf) > 0 {
		upsert.Dirty = false
		if p.CreatedAt.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableGroupPermissionsColumnCreatedAt); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.CreatedAt)
			upsert.Dirty = true

		}

		if p.CreatedBy.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableGroupPermissionsColumnCreatedBy); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.CreatedBy)
			upsert.Dirty = true
		}

		if p.GroupID.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableGroupPermissionsColumnGroupID); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.GroupID)
			upsert.Dirty = true
		}

		if p.PermissionAction.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableGroupPermissionsColumnPermissionAction); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.PermissionAction)
			upsert.Dirty = true
		}

		if p.PermissionModule.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableGroupPermissionsColumnPermissionModule); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.PermissionModule)
			upsert.Dirty = true
		}

		if p.PermissionSubsystem.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableGroupPermissionsColumnPermissionSubsystem); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.PermissionSubsystem)
			upsert.Dirty = true
		}

		if p.UpdatedAt.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableGroupPermissionsColumnUpdatedAt); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.UpdatedAt)
			upsert.Dirty = true

		} else {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableGroupPermissionsColumnUpdatedAt); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("=NOW()"); err != nil {
				return "", nil, err
			}
		}

		if p.UpdatedBy.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableGroupPermissionsColumnUpdatedBy); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.UpdatedBy)
			upsert.Dirty = true
		}

	}

	if len(inf) > 0 && upsert.Dirty {
		buf.WriteString("(")
		for j, i := range inf {
			if j != 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(i)
		}
		buf.WriteString(")")
		buf.WriteString(" DO UPDATE SET ")
		buf.ReadFrom(upsert)
	} else {
		buf.WriteString(" DO NOTHING ")
	}
	if upsert.Dirty {
		if len(r.Columns) > 0 {
			buf.WriteString(" RETURNING ")
			buf.WriteString(strings.Join(r.Columns, ", "))
		}
	}
	return buf.String(), upsert.Args(), nil
}
func (r *GroupPermissionsRepositoryBase) Upsert(ctx context.Context, e *GroupPermissionsEntity, p *GroupPermissionsPatch, inf ...string) (*GroupPermissionsEntity, error) {
	query, args, err := r.UpsertQuery(e, p, inf...)
	if err != nil {
		return nil, err
	}
	if err := r.DB.QueryRowContext(ctx, query, args...).Scan(&e.CreatedAt,
		&e.CreatedBy,
		&e.GroupID,
		&e.PermissionAction,
		&e.PermissionModule,
		&e.PermissionSubsystem,
		&e.UpdatedAt,
		&e.UpdatedBy,
	); err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "upsert query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}
	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "upsert query success", "query", query, "table", r.Table)
	}
	return e, nil
}

func (r *GroupPermissionsRepositoryBase) Count(ctx context.Context, c *GroupPermissionsCriteria) (int64, error) {
	query, args, err := r.FindQuery([]string{"COUNT(*)"}, c)
	if err != nil {
		return 0, err
	}
	var count int64
	if err := r.DB.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "count query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return 0, err
	}

	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "count query success", "query", query, "table", r.Table)
	}

	return count, nil
}

const (
	TableUserPermissions                              = "charon.user_permissions"
	TableUserPermissionsColumnCreatedAt               = "created_at"
	TableUserPermissionsColumnCreatedBy               = "created_by"
	TableUserPermissionsColumnPermissionAction        = "permission_action"
	TableUserPermissionsColumnPermissionModule        = "permission_module"
	TableUserPermissionsColumnPermissionSubsystem     = "permission_subsystem"
	TableUserPermissionsColumnUpdatedAt               = "updated_at"
	TableUserPermissionsColumnUpdatedBy               = "updated_by"
	TableUserPermissionsColumnUserID                  = "user_id"
	TableUserPermissionsConstraintCreatedByForeignKey = "charon.user_permissions_created_by_fkey"

	TableUserPermissionsConstraintUpdatedByForeignKey = "charon.user_permissions_updated_by_fkey"

	TableUserPermissionsConstraintUserIDForeignKey = "charon.user_permissions_user_id_fkey"

	TableUserPermissionsConstraintPermissionSubsystemPermissionModulePermissionActionForeignKey = "charon.user_permissions_subsystem_module_action_fkey"

	TableUserPermissionsConstraintUserIDPermissionSubsystemPermissionModulePermissionActionUnique = "charon.user_permissions_user_id_subsystem_module_action_key"
)

var (
	TableUserPermissionsColumns = []string{
		TableUserPermissionsColumnCreatedAt,
		TableUserPermissionsColumnCreatedBy,
		TableUserPermissionsColumnPermissionAction,
		TableUserPermissionsColumnPermissionModule,
		TableUserPermissionsColumnPermissionSubsystem,
		TableUserPermissionsColumnUpdatedAt,
		TableUserPermissionsColumnUpdatedBy,
		TableUserPermissionsColumnUserID,
	}
)

// UserPermissionsEntity ...
type UserPermissionsEntity struct {
	// CreatedAt ...
	CreatedAt time.Time
	// CreatedBy ...
	CreatedBy ntypes.Int64
	// PermissionAction ...
	PermissionAction string
	// PermissionModule ...
	PermissionModule string
	// PermissionSubsystem ...
	PermissionSubsystem string
	// UpdatedAt ...
	UpdatedAt pq.NullTime
	// UpdatedBy ...
	UpdatedBy ntypes.Int64
	// UserID ...
	UserID int64
	// User ...
	User *UserEntity
	// Author ...
	Author *UserEntity
	// Modifier ...
	Modifier *UserEntity
}

func (e *UserPermissionsEntity) Prop(cn string) (interface{}, bool) {
	switch cn {
	case TableUserPermissionsColumnCreatedAt:
		return &e.CreatedAt, true
	case TableUserPermissionsColumnCreatedBy:
		return &e.CreatedBy, true
	case TableUserPermissionsColumnPermissionAction:
		return &e.PermissionAction, true
	case TableUserPermissionsColumnPermissionModule:
		return &e.PermissionModule, true
	case TableUserPermissionsColumnPermissionSubsystem:
		return &e.PermissionSubsystem, true
	case TableUserPermissionsColumnUpdatedAt:
		return &e.UpdatedAt, true
	case TableUserPermissionsColumnUpdatedBy:
		return &e.UpdatedBy, true
	case TableUserPermissionsColumnUserID:
		return &e.UserID, true
	default:
		return nil, false
	}
}

func (e *UserPermissionsEntity) Props(cns ...string) ([]interface{}, error) {
	res := make([]interface{}, 0, len(cns))
	for _, cn := range cns {
		if prop, ok := e.Prop(cn); ok {
			res = append(res, prop)
		} else {
			return nil, fmt.Errorf("unexpected column provided: %s", cn)
		}
	}
	return res, nil
}

// UserPermissionsIterator is not thread safe.
type UserPermissionsIterator struct {
	rows *sql.Rows
	cols []string
}

func (i *UserPermissionsIterator) Next() bool {
	return i.rows.Next()
}

func (i *UserPermissionsIterator) Close() error {
	return i.rows.Close()
}

func (i *UserPermissionsIterator) Err() error {
	return i.rows.Err()
}

// Columns is wrapper around sql.Rows.Columns method, that also cache outpu inside iterator.
func (i *UserPermissionsIterator) Columns() ([]string, error) {
	if i.cols == nil {
		cols, err := i.rows.Columns()
		if err != nil {
			return nil, err
		}
		i.cols = cols
	}
	return i.cols, nil
}

// Ent is wrapper around UserPermissions method that makes iterator more generic.
func (i *UserPermissionsIterator) Ent() (interface{}, error) {
	return i.UserPermissions()
}

func (i *UserPermissionsIterator) UserPermissions() (*UserPermissionsEntity, error) {
	var ent UserPermissionsEntity
	cols, err := i.rows.Columns()
	if err != nil {
		return nil, err
	}

	props, err := ent.Props(cols...)
	if err != nil {
		return nil, err
	}
	if err := i.rows.Scan(props...); err != nil {
		return nil, err
	}
	return &ent, nil
}

type UserPermissionsCriteria struct {
	Offset, Limit       int64
	Sort                map[string]bool
	CreatedAt           *qtypes.Timestamp
	CreatedBy           *qtypes.Int64
	PermissionAction    *qtypes.String
	PermissionModule    *qtypes.String
	PermissionSubsystem *qtypes.String
	UpdatedAt           *qtypes.Timestamp
	UpdatedBy           *qtypes.Int64
	UserID              *qtypes.Int64
}

type UserPermissionsPatch struct {
	CreatedAt           pq.NullTime
	CreatedBy           ntypes.Int64
	PermissionAction    ntypes.String
	PermissionModule    ntypes.String
	PermissionSubsystem ntypes.String
	UpdatedAt           pq.NullTime
	UpdatedBy           ntypes.Int64
	UserID              ntypes.Int64
}

func ScanUserPermissionsRows(rows *sql.Rows) (entities []*UserPermissionsEntity, err error) {
	for rows.Next() {
		var ent UserPermissionsEntity
		err = rows.Scan(&ent.CreatedAt,
			&ent.CreatedBy,
			&ent.PermissionAction,
			&ent.PermissionModule,
			&ent.PermissionSubsystem,
			&ent.UpdatedAt,
			&ent.UpdatedBy,
			&ent.UserID,
		)
		if err != nil {
			return
		}

		entities = append(entities, &ent)
	}
	if err = rows.Err(); err != nil {
		return
	}

	return
}

type UserPermissionsRepositoryBase struct {
	Table   string
	Columns []string
	DB      *sql.DB
	Debug   bool
	Log     log.Logger
}

func (r *UserPermissionsRepositoryBase) InsertQuery(e *UserPermissionsEntity) (string, []interface{}, error) {
	insert := NewComposer(8)
	columns := bytes.NewBuffer(nil)
	buf := bytes.NewBufferString("INSERT INTO ")
	buf.WriteString(r.Table)

	if !e.CreatedAt.IsZero() {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableUserPermissionsColumnCreatedAt); err != nil {
			return "", nil, err
		}
		if insert.Dirty {
			if _, err := insert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := insert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		insert.Add(e.CreatedAt)
		insert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserPermissionsColumnCreatedBy); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.CreatedBy)
	insert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserPermissionsColumnPermissionAction); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.PermissionAction)
	insert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserPermissionsColumnPermissionModule); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.PermissionModule)
	insert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserPermissionsColumnPermissionSubsystem); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.PermissionSubsystem)
	insert.Dirty = true

	if e.UpdatedAt.Valid {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableUserPermissionsColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if insert.Dirty {
			if _, err := insert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := insert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		insert.Add(e.UpdatedAt)
		insert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserPermissionsColumnUpdatedBy); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.UpdatedBy)
	insert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserPermissionsColumnUserID); err != nil {
		return "", nil, err
	}
	if insert.Dirty {
		if _, err := insert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := insert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	insert.Add(e.UserID)
	insert.Dirty = true
	if columns.Len() > 0 {
		buf.WriteString(" (")
		buf.ReadFrom(columns)
		buf.WriteString(") VALUES (")
		buf.ReadFrom(insert)
		buf.WriteString(") ")
		if len(r.Columns) > 0 {
			buf.WriteString("RETURNING ")
			buf.WriteString(strings.Join(r.Columns, ", "))
		}
	}
	return buf.String(), insert.Args(), nil
}

func (r *UserPermissionsRepositoryBase) Insert(ctx context.Context, e *UserPermissionsEntity) (*UserPermissionsEntity, error) {
	query, args, err := r.InsertQuery(e)
	if err != nil {
		return nil, err
	}
	if err := r.DB.QueryRowContext(ctx, query, args...).Scan(&e.CreatedAt,
		&e.CreatedBy,
		&e.PermissionAction,
		&e.PermissionModule,
		&e.PermissionSubsystem,
		&e.UpdatedAt,
		&e.UpdatedBy,
		&e.UserID,
	); err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}
	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query success", "query", query, "table", r.Table)
	}
	return e, nil
}

func (r *UserPermissionsRepositoryBase) FindQuery(s []string, c *UserPermissionsCriteria) (string, []interface{}, error) {
	where := NewComposer(8)
	buf := bytes.NewBufferString("SELECT ")
	buf.WriteString(strings.Join(s, ", "))
	buf.WriteString(" FROM ")
	buf.WriteString(r.Table)
	buf.WriteString(" ")
	QueryTimestampWhereClause(c.CreatedAt, TableUserPermissionsColumnCreatedAt, where, And)

	QueryInt64WhereClause(c.CreatedBy, TableUserPermissionsColumnCreatedBy, where, And)

	QueryStringWhereClause(c.PermissionAction, TableUserPermissionsColumnPermissionAction, where, And)

	QueryStringWhereClause(c.PermissionModule, TableUserPermissionsColumnPermissionModule, where, And)

	QueryStringWhereClause(c.PermissionSubsystem, TableUserPermissionsColumnPermissionSubsystem, where, And)

	QueryTimestampWhereClause(c.UpdatedAt, TableUserPermissionsColumnUpdatedAt, where, And)

	QueryInt64WhereClause(c.UpdatedBy, TableUserPermissionsColumnUpdatedBy, where, And)

	QueryInt64WhereClause(c.UserID, TableUserPermissionsColumnUserID, where, And)
	if where.Dirty {
		if _, err := buf.WriteString("WHERE "); err != nil {
			return "", nil, err
		}
		buf.ReadFrom(where)
	}

	if len(c.Sort) > 0 {
		i := 0
		where.WriteString(" ORDER BY ")

		for cn, asc := range c.Sort {
			for _, tcn := range TableUserPermissionsColumns {
				if cn == tcn {
					if i > 0 {
						if _, err := where.WriteString(", "); err != nil {
							return "", nil, err
						}
					}
					if _, err := where.WriteString(cn); err != nil {
						return "", nil, err
					}
					if !asc {
						if _, err := where.WriteString(" DESC "); err != nil {
							return "", nil, err
						}
					}
					i++
					break
				}
			}
		}
	}
	if c.Offset > 0 {
		if _, err := where.WriteString(" OFFSET "); err != nil {
			return "", nil, err
		}
		if err := where.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		if _, err := where.WriteString(" "); err != nil {
			return "", nil, err
		}
		where.Add(c.Offset)
	}
	if c.Limit > 0 {
		if _, err := where.WriteString(" LIMIT "); err != nil {
			return "", nil, err
		}
		if err := where.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		if _, err := where.WriteString(" "); err != nil {
			return "", nil, err
		}
		where.Add(c.Limit)
	}

	buf.ReadFrom(where)

	return buf.String(), where.Args(), nil
}

func (r *UserPermissionsRepositoryBase) Find(ctx context.Context, c *UserPermissionsCriteria) ([]*UserPermissionsEntity, error) {
	query, args, err := r.FindQuery(r.Columns, c)
	if err != nil {
		return nil, err
	}
	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "find query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}
	defer rows.Close()

	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "find query success", "query", query, "table", r.Table)
	}

	return ScanUserPermissionsRows(rows)
}

func (r *UserPermissionsRepositoryBase) FindIter(ctx context.Context, c *UserPermissionsCriteria) (*UserPermissionsIterator, error) {
	query, args, err := r.FindQuery(r.Columns, c)
	if err != nil {
		return nil, err
	}
	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &UserPermissionsIterator{rows: rows}, nil
}
func (r *UserPermissionsRepositoryBase) FindOneByUserIDAndPermissionSubsystemAndPermissionModuleAndPermissionAction(ctx context.Context, userPermissionsUserID int64, userPermissionsPermissionSubsystem string, userPermissionsPermissionModule string, userPermissionsPermissionAction string) (*UserPermissionsEntity, error) {
	find := NewComposer(8)
	find.WriteString("SELECT ")
	find.WriteString(strings.Join(r.Columns, ", "))
	find.WriteString(" FROM ")
	find.WriteString(TableUserPermissions)
	find.WriteString(" WHERE ")
	find.WriteString(TableUserPermissionsColumnUserID)
	find.WriteString("=")
	find.WritePlaceholder()
	find.Add(userPermissionsUserID)
	find.WriteString(" AND ")
	find.WriteString(TableUserPermissionsColumnPermissionSubsystem)
	find.WriteString("=")
	find.WritePlaceholder()
	find.Add(userPermissionsPermissionSubsystem)
	find.WriteString(" AND ")
	find.WriteString(TableUserPermissionsColumnPermissionModule)
	find.WriteString("=")
	find.WritePlaceholder()
	find.Add(userPermissionsPermissionModule)
	find.WriteString(" AND ")
	find.WriteString(TableUserPermissionsColumnPermissionAction)
	find.WriteString("=")
	find.WritePlaceholder()
	find.Add(userPermissionsPermissionAction)

	var (
		ent UserPermissionsEntity
	)
	props, err := ent.Props(r.Columns...)
	if err != nil {
		return nil, err
	}
	err = r.DB.QueryRowContext(ctx, find.String(), find.Args()...).Scan(props...)
	if err != nil {
		return nil, err
	}

	return &ent, nil
}
func (r *UserPermissionsRepositoryBase) UpdateOneByUserIDAndPermissionSubsystemAndPermissionModuleAndPermissionActionQuery(userPermissionsUserID int64, userPermissionsPermissionSubsystem string, userPermissionsPermissionModule string, userPermissionsPermissionAction string, p *UserPermissionsPatch) (string, []interface{}, error) {
	buf := bytes.NewBufferString("UPDATE ")
	buf.WriteString(r.Table)
	update := NewComposer(4)
	if p.CreatedAt.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserPermissionsColumnCreatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.CreatedAt)
		update.Dirty = true

	}

	if p.CreatedBy.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserPermissionsColumnCreatedBy); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.CreatedBy)
		update.Dirty = true
	}

	if p.PermissionAction.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserPermissionsColumnPermissionAction); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.PermissionAction)
		update.Dirty = true
	}

	if p.PermissionModule.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserPermissionsColumnPermissionModule); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.PermissionModule)
		update.Dirty = true
	}

	if p.PermissionSubsystem.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserPermissionsColumnPermissionSubsystem); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.PermissionSubsystem)
		update.Dirty = true
	}

	if p.UpdatedAt.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserPermissionsColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.UpdatedAt)
		update.Dirty = true

	} else {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserPermissionsColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("=NOW()"); err != nil {
			return "", nil, err
		}
	}

	if p.UpdatedBy.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserPermissionsColumnUpdatedBy); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.UpdatedBy)
		update.Dirty = true
	}

	if p.UserID.Valid {
		if update.Dirty {
			if _, err := update.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := update.WriteString(TableUserPermissionsColumnUserID); err != nil {
			return "", nil, err
		}
		if _, err := update.WriteString("="); err != nil {
			return "", nil, err
		}
		if err := update.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		update.Add(p.UserID)
		update.Dirty = true
	}

	if !update.Dirty {
		return "", nil, errors.New("UserPermissions update failure, nothing to update")
	}
	buf.WriteString(" SET ")
	buf.ReadFrom(update)
	buf.WriteString(" WHERE ")
	update.WriteString(TableUserPermissionsColumnUserID)
	update.WriteString("=")
	update.WritePlaceholder()
	update.Add(userPermissionsUserID)
	update.WriteString(" AND ")
	update.WriteString(TableUserPermissionsColumnPermissionSubsystem)
	update.WriteString("=")
	update.WritePlaceholder()
	update.Add(userPermissionsPermissionSubsystem)
	update.WriteString(" AND ")
	update.WriteString(TableUserPermissionsColumnPermissionModule)
	update.WriteString("=")
	update.WritePlaceholder()
	update.Add(userPermissionsPermissionModule)
	update.WriteString(" AND ")
	update.WriteString(TableUserPermissionsColumnPermissionAction)
	update.WriteString("=")
	update.WritePlaceholder()
	update.Add(userPermissionsPermissionAction)
	buf.ReadFrom(update)
	buf.WriteString(" RETURNING ")
	buf.WriteString(strings.Join(r.Columns, ", "))
	return buf.String(), update.Args(), nil
}
func (r *UserPermissionsRepositoryBase) UpdateOneByUserIDAndPermissionSubsystemAndPermissionModuleAndPermissionAction(ctx context.Context, userPermissionsUserID int64, userPermissionsPermissionSubsystem string, userPermissionsPermissionModule string, userPermissionsPermissionAction string, p *UserPermissionsPatch) (*UserPermissionsEntity, error) {
	query, args, err := r.UpdateOneByUserIDAndPermissionSubsystemAndPermissionModuleAndPermissionActionQuery(userPermissionsUserID, userPermissionsPermissionSubsystem, userPermissionsPermissionModule, userPermissionsPermissionAction, p)
	if err != nil {
		return nil, err
	}
	var ent UserPermissionsEntity
	props, err := ent.Props(r.Columns...)
	if err != nil {
		return nil, err
	}
	err = r.DB.QueryRowContext(ctx, query, args...).Scan(props...)
	if err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}

	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "insert query success", "query", query, "table", r.Table)
	}

	return &ent, nil
}
func (r *UserPermissionsRepositoryBase) UpsertQuery(e *UserPermissionsEntity, p *UserPermissionsPatch, inf ...string) (string, []interface{}, error) {
	upsert := NewComposer(16)
	columns := bytes.NewBuffer(nil)
	buf := bytes.NewBufferString("INSERT INTO ")
	buf.WriteString(r.Table)

	if !e.CreatedAt.IsZero() {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableUserPermissionsColumnCreatedAt); err != nil {
			return "", nil, err
		}
		if upsert.Dirty {
			if _, err := upsert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := upsert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		upsert.Add(e.CreatedAt)
		upsert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserPermissionsColumnCreatedBy); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.CreatedBy)
	upsert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserPermissionsColumnPermissionAction); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.PermissionAction)
	upsert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserPermissionsColumnPermissionModule); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.PermissionModule)
	upsert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserPermissionsColumnPermissionSubsystem); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.PermissionSubsystem)
	upsert.Dirty = true

	if e.UpdatedAt.Valid {
		if columns.Len() > 0 {
			if _, err := columns.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if _, err := columns.WriteString(TableUserPermissionsColumnUpdatedAt); err != nil {
			return "", nil, err
		}
		if upsert.Dirty {
			if _, err := upsert.WriteString(", "); err != nil {
				return "", nil, err
			}
		}
		if err := upsert.WritePlaceholder(); err != nil {
			return "", nil, err
		}
		upsert.Add(e.UpdatedAt)
		upsert.Dirty = true
	}

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserPermissionsColumnUpdatedBy); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.UpdatedBy)
	upsert.Dirty = true

	if columns.Len() > 0 {
		if _, err := columns.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if _, err := columns.WriteString(TableUserPermissionsColumnUserID); err != nil {
		return "", nil, err
	}
	if upsert.Dirty {
		if _, err := upsert.WriteString(", "); err != nil {
			return "", nil, err
		}
	}
	if err := upsert.WritePlaceholder(); err != nil {
		return "", nil, err
	}
	upsert.Add(e.UserID)
	upsert.Dirty = true

	if upsert.Dirty {
		buf.WriteString(" (")
		buf.ReadFrom(columns)
		buf.WriteString(") VALUES (")
		buf.ReadFrom(upsert)
		buf.WriteString(")")
	}
	buf.WriteString(" ON CONFLICT ")
	if len(inf) > 0 {
		upsert.Dirty = false
		if p.CreatedAt.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserPermissionsColumnCreatedAt); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.CreatedAt)
			upsert.Dirty = true

		}

		if p.CreatedBy.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserPermissionsColumnCreatedBy); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.CreatedBy)
			upsert.Dirty = true
		}

		if p.PermissionAction.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserPermissionsColumnPermissionAction); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.PermissionAction)
			upsert.Dirty = true
		}

		if p.PermissionModule.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserPermissionsColumnPermissionModule); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.PermissionModule)
			upsert.Dirty = true
		}

		if p.PermissionSubsystem.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserPermissionsColumnPermissionSubsystem); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.PermissionSubsystem)
			upsert.Dirty = true
		}

		if p.UpdatedAt.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserPermissionsColumnUpdatedAt); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.UpdatedAt)
			upsert.Dirty = true

		} else {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserPermissionsColumnUpdatedAt); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("=NOW()"); err != nil {
				return "", nil, err
			}
		}

		if p.UpdatedBy.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserPermissionsColumnUpdatedBy); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.UpdatedBy)
			upsert.Dirty = true
		}

		if p.UserID.Valid {
			if upsert.Dirty {
				if _, err := upsert.WriteString(", "); err != nil {
					return "", nil, err
				}
			}
			if _, err := upsert.WriteString(TableUserPermissionsColumnUserID); err != nil {
				return "", nil, err
			}
			if _, err := upsert.WriteString("="); err != nil {
				return "", nil, err
			}
			if err := upsert.WritePlaceholder(); err != nil {
				return "", nil, err
			}
			upsert.Add(p.UserID)
			upsert.Dirty = true
		}

	}

	if len(inf) > 0 && upsert.Dirty {
		buf.WriteString("(")
		for j, i := range inf {
			if j != 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(i)
		}
		buf.WriteString(")")
		buf.WriteString(" DO UPDATE SET ")
		buf.ReadFrom(upsert)
	} else {
		buf.WriteString(" DO NOTHING ")
	}
	if upsert.Dirty {
		if len(r.Columns) > 0 {
			buf.WriteString(" RETURNING ")
			buf.WriteString(strings.Join(r.Columns, ", "))
		}
	}
	return buf.String(), upsert.Args(), nil
}
func (r *UserPermissionsRepositoryBase) Upsert(ctx context.Context, e *UserPermissionsEntity, p *UserPermissionsPatch, inf ...string) (*UserPermissionsEntity, error) {
	query, args, err := r.UpsertQuery(e, p, inf...)
	if err != nil {
		return nil, err
	}
	if err := r.DB.QueryRowContext(ctx, query, args...).Scan(&e.CreatedAt,
		&e.CreatedBy,
		&e.PermissionAction,
		&e.PermissionModule,
		&e.PermissionSubsystem,
		&e.UpdatedAt,
		&e.UpdatedBy,
		&e.UserID,
	); err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "upsert query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return nil, err
	}
	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "upsert query success", "query", query, "table", r.Table)
	}
	return e, nil
}

func (r *UserPermissionsRepositoryBase) Count(ctx context.Context, c *UserPermissionsCriteria) (int64, error) {
	query, args, err := r.FindQuery([]string{"COUNT(*)"}, c)
	if err != nil {
		return 0, err
	}
	var count int64
	if err := r.DB.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		if r.Debug {
			r.Log.Log("level", "error", "timestamp", time.Now().Format(time.RFC3339), "msg", "count query failure", "query", query, "table", r.Table, "error", err.Error())
		}
		return 0, err
	}

	if r.Debug {
		r.Log.Log("level", "debug", "timestamp", time.Now().Format(time.RFC3339), "msg", "count query success", "query", query, "table", r.Table)
	}

	return count, nil
}

// ErrorConstraint returns the error constraint of err if it was produced by the pq library.
// Otherwise, it returns empty string.
func ErrorConstraint(err error) string {
	if err == nil {
		return ""
	}
	if pqerr, ok := err.(*pq.Error); ok {
		return pqerr.Constraint
	}

	return ""
}

type NullInt64Array struct {
	pq.Int64Array
	Valid bool
}

func (n *NullInt64Array) Scan(value interface{}) error {
	if value == nil {
		n.Int64Array, n.Valid = nil, false
		return nil
	}
	n.Valid = true
	return n.Int64Array.Scan(value)
}

type NullFloat64Array struct {
	pq.Float64Array
	Valid bool
}

func (n *NullFloat64Array) Scan(value interface{}) error {
	if value == nil {
		n.Float64Array, n.Valid = nil, false
		return nil
	}
	n.Valid = true
	return n.Float64Array.Scan(value)
}

type NullBoolArray struct {
	pq.BoolArray
	Valid bool
}

func (n *NullBoolArray) Scan(value interface{}) error {
	if value == nil {
		n.BoolArray, n.Valid = nil, false
		return nil
	}
	n.Valid = true
	return n.BoolArray.Scan(value)
}

type NullStringArray struct {
	pq.StringArray
	Valid bool
}

func (n *NullStringArray) Scan(value interface{}) error {
	if value == nil {
		n.StringArray, n.Valid = nil, false
		return nil
	}
	n.Valid = true
	return n.StringArray.Scan(value)
}

type NullByteaArray struct {
	pq.ByteaArray
	Valid bool
}

func (n *NullByteaArray) Scan(value interface{}) error {
	if value == nil {
		n.ByteaArray, n.Valid = nil, false
		return nil
	}
	n.Valid = true
	return n.ByteaArray.Scan(value)
}

const (
	jsonArraySeparator     = ","
	jsonArrayBeginningChar = "["
	jsonArrayEndChar       = "]"
)

// JSONArrayInt64 is a slice of int64s that implements necessary interfaces.
type JSONArrayInt64 []int64

// Scan satisfy sql.Scanner interface.
func (a *JSONArrayInt64) Scan(src interface{}) error {
	if src == nil {
		if a == nil {
			*a = make(JSONArrayInt64, 0)
		}
		return nil
	}

	var tmp []string
	var srcs string

	switch t := src.(type) {
	case []byte:
		srcs = string(t)
	case string:
		srcs = t
	default:
		return fmt.Errorf("pqt: expected slice of bytes or string as a source argument in Scan, not %T", src)
	}

	l := len(srcs)

	if l < 2 {
		return fmt.Errorf("pqt: expected to get source argument in format '[1,2,...,N]', but got %s", srcs)
	}

	if l == 2 {
		*a = make(JSONArrayInt64, 0)
		return nil
	}

	if string(srcs[0]) != jsonArrayBeginningChar || string(srcs[l-1]) != jsonArrayEndChar {
		return fmt.Errorf("pqt: expected to get source argument in format '[1,2,...,N]', but got %s", srcs)
	}

	tmp = strings.Split(string(srcs[1:l-1]), jsonArraySeparator)
	*a = make(JSONArrayInt64, 0, len(tmp))
	for i, v := range tmp {
		j, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return fmt.Errorf("pqt: expected to get source argument in format '[1,2,...,N]', but got %s at index %d", v, i)
		}

		*a = append(*a, j)
	}

	return nil
}

// Value satisfy driver.Valuer interface.
func (a JSONArrayInt64) Value() (driver.Value, error) {
	var (
		buffer bytes.Buffer
		err    error
	)

	if _, err = buffer.WriteString(jsonArrayBeginningChar); err != nil {
		return nil, err
	}

	for i, v := range a {
		if i > 0 {
			if _, err := buffer.WriteString(jsonArraySeparator); err != nil {
				return nil, err
			}
		}
		if _, err := buffer.WriteString(strconv.FormatInt(v, 10)); err != nil {
			return nil, err
		}
	}

	if _, err = buffer.WriteString(jsonArrayEndChar); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// JSONArrayString is a slice of strings that implements necessary interfaces.
type JSONArrayString []string

// Scan satisfy sql.Scanner interface.
func (a *JSONArrayString) Scan(src interface{}) error {
	if src == nil {
		if a == nil {
			*a = make(JSONArrayString, 0)
		}
		return nil
	}

	var srcs string

	switch t := src.(type) {
	case []byte:
		srcs = string(t)
	case string:
		srcs = t
	default:
		return fmt.Errorf("pqt: expected slice of bytes or string as a source argument in Scan, not %T", src)
	}

	l := len(srcs)

	if l < 2 {
		return fmt.Errorf("pqt: expected to get source argument in format '[text1,text2,...,textN]', but got %s", srcs)
	}

	if string(srcs[0]) != jsonArrayBeginningChar || string(srcs[l-1]) != jsonArrayEndChar {
		return fmt.Errorf("pqt: expected to get source argument in format '[text1,text2,...,textN]', but got %s", srcs)
	}

	*a = strings.Split(string(srcs[1:l-1]), jsonArraySeparator)

	return nil
}

// Value satisfy driver.Valuer interface.
func (a JSONArrayString) Value() (driver.Value, error) {
	var (
		buffer bytes.Buffer
		err    error
	)

	if _, err = buffer.WriteString(jsonArrayBeginningChar); err != nil {
		return nil, err
	}

	for i, v := range a {
		if i > 0 {
			if _, err := buffer.WriteString(jsonArraySeparator); err != nil {
				return nil, err
			}
		}
		if _, err = buffer.WriteString(v); err != nil {
			return nil, err
		}
	}

	if _, err = buffer.WriteString(jsonArrayEndChar); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// JSONArrayFloat64 is a slice of int64s that implements necessary interfaces.
type JSONArrayFloat64 []float64

// Scan satisfy sql.Scanner interface.
func (a *JSONArrayFloat64) Scan(src interface{}) error {
	if src == nil {
		if a == nil {
			*a = make(JSONArrayFloat64, 0)
		}
		return nil
	}

	var tmp []string
	var srcs string

	switch t := src.(type) {
	case []byte:
		srcs = string(t)
	case string:
		srcs = t
	default:
		return fmt.Errorf("pqt: expected slice of bytes or string as a source argument in Scan, not %T", src)
	}

	l := len(srcs)

	if l < 2 {
		return fmt.Errorf("pqt: expected to get source argument in format '[1.3,2.4,...,N.M]', but got %s", srcs)
	}

	if l == 2 {
		*a = make(JSONArrayFloat64, 0)
		return nil
	}

	if string(srcs[0]) != jsonArrayBeginningChar || string(srcs[l-1]) != jsonArrayEndChar {
		return fmt.Errorf("pqt: expected to get source argument in format '[1.3,2.4,...,N.M]', but got %s", srcs)
	}

	tmp = strings.Split(string(srcs[1:l-1]), jsonArraySeparator)
	*a = make(JSONArrayFloat64, 0, len(tmp))
	for i, v := range tmp {
		j, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return fmt.Errorf("pqt: expected to get source argument in format '[1.3,2.4,...,N.M]', but got %s at index %d", v, i)
		}

		*a = append(*a, j)
	}

	return nil
}

// Value satisfy driver.Valuer interface.
func (a JSONArrayFloat64) Value() (driver.Value, error) {
	var (
		buffer bytes.Buffer
		err    error
	)

	if _, err = buffer.WriteString(jsonArrayBeginningChar); err != nil {
		return nil, err
	}

	for i, v := range a {
		if i > 0 {
			if _, err := buffer.WriteString(jsonArraySeparator); err != nil {
				return nil, err
			}
		}
		if _, err := buffer.WriteString(strconv.FormatFloat(v, 'f', -1, 64)); err != nil {
			return nil, err
		}
	}

	if _, err = buffer.WriteString(jsonArrayEndChar); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

var (
	// Space is a shorthand composition option that holds space.
	Space = &CompositionOpts{
		Joint: " ",
	}
	// And is a shorthand composition option that holds AND operator.
	And = &CompositionOpts{
		Joint: " AND ",
	}
	// Or is a shorthand composition option that holds OR operator.
	Or = &CompositionOpts{
		Joint: " OR ",
	}
	// Comma is a shorthand composition option that holds comma.
	Comma = &CompositionOpts{
		Joint: ", ",
	}
)

// CompositionOpts is a container for modification that can be applied.
type CompositionOpts struct {
	Joint                         string
	PlaceholderFunc, SelectorFunc string
	Cast                          string
	IsJSON                        bool
}

// CompositionWriter is a simple wrapper for WriteComposition function.
type CompositionWriter interface {
	// WriteComposition is a function that allow custom struct type to be used as a part of criteria.
	// It gives possibility to write custom query based on object that implements this interface.
	WriteComposition(string, *Composer, *CompositionOpts) error
}

// Composer holds buffer, arguments and placeholders count.
// In combination with external buffet can be also used to also generate sub-queries.
// To do that simply write buffer to the parent buffer, composer will hold all arguments and remember number of last placeholder.
type Composer struct {
	buf     bytes.Buffer
	args    []interface{}
	counter int
	Dirty   bool
}

// NewComposer allocates new Composer with inner slice of arguments of given size.
func NewComposer(size int64) *Composer {
	return &Composer{
		counter: 1,
		args:    make([]interface{}, 0, size),
	}
}

// WriteString appends the contents of s to the query buffer, growing the buffer as
// needed. The return value n is the length of s; err is always nil. If the
// buffer becomes too large, WriteString will panic with bytes ErrTooLarge.
func (c *Composer) WriteString(s string) (int, error) {
	return c.buf.WriteString(s)
}

// Write implements io Writer interface.
func (c *Composer) Write(b []byte) (int, error) {
	return c.buf.Write(b)
}

// Read implements io Reader interface.
func (c *Composer) Read(b []byte) (int, error) {
	return c.buf.Read(b)
}

// ResetBuf resets internal buffer.
func (c *Composer) ResetBuf() {
	c.buf.Reset()
}

// String implements fmt Stringer interface.
func (c *Composer) String() string {
	return c.buf.String()
}

// WritePlaceholder writes appropriate placeholder to the query buffer based on current state of the composer.
func (c *Composer) WritePlaceholder() error {
	if _, err := c.buf.WriteString("$"); err != nil {
		return err
	}
	if _, err := c.buf.WriteString(strconv.Itoa(c.counter)); err != nil {
		return err
	}

	c.counter++
	return nil
}

// Len returns number of arguments.
func (c *Composer) Len() int {
	return c.counter
}

// Add appends list with new element.
func (c *Composer) Add(arg interface{}) {
	c.args = append(c.args, arg)
}

// Args returns all arguments stored as a slice.
func (c *Composer) Args() []interface{} {
	return c.args
}

func QueryInt64WhereClause(i *qtypes.Int64, sel string, com *Composer, opt *CompositionOpts) (err error) {
	if i == nil || !i.Valid {
		return nil
	}
	switch i.Type {
	case qtypes.QueryType_NULL:
		if com.Dirty {
			if _, err = com.WriteString(opt.Joint); err != nil {
				return
			}
		}
		com.WriteString(sel)
		if i.Negation {
			if _, err = com.WriteString(" IS NOT NULL"); err != nil {
				return
			}
		} else {
			if _, err = com.WriteString(" IS NULL"); err != nil {
				return
			}
		}
		com.Dirty = true
		return
	case qtypes.QueryType_EQUAL:
		if com.Dirty {
			if _, err = com.WriteString(opt.Joint); err != nil {
				return
			}
		}
		if _, err = com.WriteString(sel); err != nil {
			return
		}
		if i.Negation {
			if _, err = com.WriteString(" <> "); err != nil {
				return
			}
		} else {
			if _, err = com.WriteString(" = "); err != nil {
				return
			}
		}
		if err = com.WritePlaceholder(); err != nil {
			return
		}
		com.Add(i.Value())
		com.Dirty = true
	case qtypes.QueryType_GREATER:
		if com.Dirty {
			if _, err = com.WriteString(opt.Joint); err != nil {
				return
			}
		}
		if _, err = com.WriteString(sel); err != nil {
			return
		}
		if i.Negation {
			if _, err = com.WriteString(" <= "); err != nil {
				return
			}
		} else {
			if _, err = com.WriteString(" > "); err != nil {
				return
			}
		}
		if err = com.WritePlaceholder(); err != nil {
			return
		}
		com.Add(i.Value())
		com.Dirty = true
	case qtypes.QueryType_GREATER_EQUAL:
		if com.Dirty {
			if _, err = com.WriteString(opt.Joint); err != nil {
				return
			}
		}
		if _, err = com.WriteString(sel); err != nil {
			return
		}
		if i.Negation {
			if _, err = com.WriteString(" < "); err != nil {
				return
			}
		} else {
			if _, err = com.WriteString(" >= "); err != nil {
				return
			}
		}
		if err = com.WritePlaceholder(); err != nil {
			return
		}
		com.Add(i.Value())
		com.Dirty = true
	case qtypes.QueryType_LESS:
		if com.Dirty {
			if _, err = com.WriteString(opt.Joint); err != nil {
				return
			}
		}
		if _, err = com.WriteString(sel); err != nil {
			return
		}
		if i.Negation {
			if _, err = com.WriteString(" >= "); err != nil {
				return
			}
		} else {
			if _, err = com.WriteString(" < "); err != nil {
				return
			}
		}
		if err = com.WritePlaceholder(); err != nil {
			return
		}
		com.Add(i.Value())
		com.Dirty = true
	case qtypes.QueryType_LESS_EQUAL:
		if com.Dirty {
			if _, err = com.WriteString(opt.Joint); err != nil {
				return
			}
		}
		if _, err = com.WriteString(sel); err != nil {
			return
		}
		if i.Negation {
			if _, err = com.WriteString(" > "); err != nil {
				return
			}
		} else {
			if _, err = com.WriteString(" <= "); err != nil {
				return
			}
		}
		if err = com.WritePlaceholder(); err != nil {
			return
		}
		com.Add(i.Value())
		com.Dirty = true
	case qtypes.QueryType_CONTAINS:
		if !i.Negation {
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return
				}
			}
			if _, err = com.WriteString(sel); err != nil {
				return
			}
			if _, err = com.WriteString(" @> "); err != nil {
				return
			}
			if err = com.WritePlaceholder(); err != nil {
				return
			}
			switch opt.IsJSON {
			case true:
				com.Add(JSONArrayInt64(i.Values))
			case false:
				com.Add(pq.Int64Array(i.Values))
			}
			com.Dirty = true
		}
	case qtypes.QueryType_IS_CONTAINED_BY:
		if !i.Negation {
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return
				}
			}
			if _, err = com.WriteString(sel); err != nil {
				return
			}
			if _, err = com.WriteString(" <@ "); err != nil {
				return
			}
			if err = com.WritePlaceholder(); err != nil {
				return
			}
			switch opt.IsJSON {
			case true:
				com.Add(JSONArrayInt64(i.Values))
			case false:
				com.Add(pq.Int64Array(i.Values))
			}
			com.Dirty = true
		}
	case qtypes.QueryType_OVERLAP:
		if !i.Negation {
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return
				}
			}
			if _, err = com.WriteString(sel); err != nil {
				return
			}
			if _, err = com.WriteString(" && "); err != nil {
				return
			}
			if err = com.WritePlaceholder(); err != nil {
				return
			}
			switch opt.IsJSON {
			case true:
				com.Add(JSONArrayInt64(i.Values))
			case false:
				com.Add(pq.Int64Array(i.Values))
			}
			com.Dirty = true
		}
	case qtypes.QueryType_HAS_ANY_ELEMENT:
		if !i.Negation {
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return
				}
			}
			if _, err = com.WriteString(sel); err != nil {
				return
			}
			if _, err = com.WriteString(" ?| "); err != nil {
				return
			}
			if err = com.WritePlaceholder(); err != nil {
				return
			}
			switch opt.IsJSON {
			case true:
				com.Add(JSONArrayInt64(i.Values))
			case false:
				com.Add(pq.Int64Array(i.Values))
			}
			com.Dirty = true
		}
	case qtypes.QueryType_HAS_ALL_ELEMENTS:
		if !i.Negation {
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return
				}
			}
			if _, err = com.WriteString(sel); err != nil {
				return
			}
			if _, err = com.WriteString(" ?& "); err != nil {
				return
			}
			if err = com.WritePlaceholder(); err != nil {
				return
			}
			switch opt.IsJSON {
			case true:
				com.Add(JSONArrayInt64(i.Values))
			case false:
				com.Add(pq.Int64Array(i.Values))
			}
			com.Dirty = true
		}
	case qtypes.QueryType_HAS_ELEMENT:
		if !i.Negation {
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return
				}
			}
			if _, err = com.WriteString(sel); err != nil {
				return
			}
			if _, err = com.WriteString(" ? "); err != nil {
				return
			}
			if err = com.WritePlaceholder(); err != nil {
				return
			}
			com.Add(i.Value())
			com.Dirty = true
		}
	case qtypes.QueryType_IN:
		if len(i.Values) > 0 {
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return
				}
			}
			if _, err = com.WriteString(sel); err != nil {
				return
			}
			if i.Negation {
				if _, err = com.WriteString(" NOT IN ("); err != nil {
					return
				}
			} else {
				if _, err = com.WriteString(" IN ("); err != nil {
					return
				}
			}
			for i, v := range i.Values {
				if i != 0 {
					if _, err = com.WriteString(","); err != nil {
						return
					}
				}
				if err = com.WritePlaceholder(); err != nil {
					return
				}
				com.Add(v)
				com.Dirty = true
			}
			if _, err = com.WriteString(")"); err != nil {
				return
			}
		}
	case qtypes.QueryType_BETWEEN:
		if com.Dirty {
			if _, err = com.WriteString(opt.Joint); err != nil {
				return
			}
		}
		if _, err = com.WriteString(sel); err != nil {
			return
		}
		if i.Negation {
			if _, err = com.WriteString(" <= "); err != nil {
				return
			}
		} else {
			if _, err = com.WriteString(" > "); err != nil {
				return
			}
		}
		if err = com.WritePlaceholder(); err != nil {
			return
		}
		com.Add(i.Values[0])
		if _, err = com.WriteString(" AND "); err != nil {
			return
		}
		if _, err = com.WriteString(sel); err != nil {
			return
		}
		if i.Negation {
			if _, err = com.WriteString(" >= "); err != nil {
				return
			}
		} else {
			if _, err = com.WriteString(" < "); err != nil {
				return
			}
		}
		if err = com.WritePlaceholder(); err != nil {
			return
		}
		com.Add(i.Values[1])
		com.Dirty = true
	default:
		return fmt.Errorf("pqtgo: unknown int64 query type %!s(MISSING)", i.Type.String())
	}

	if com.Dirty {
		if opt.Cast != "" {
			if _, err = com.WriteString(opt.Cast); err != nil {
				return
			}
		}
	}

	return
}
func QueryFloat64WhereClause(i *qtypes.Float64, sel string, com *Composer, opt *CompositionOpts) (err error) {
	if i == nil || !i.Valid {
		return nil
	}
	switch i.Type {
	case qtypes.QueryType_NULL:
		if com.Dirty {
			if _, err = com.WriteString(opt.Joint); err != nil {
				return
			}
		}
		com.WriteString(sel)
		if i.Negation {
			if _, err = com.WriteString(" IS NOT NULL"); err != nil {
				return
			}
		} else {
			if _, err = com.WriteString(" IS NULL"); err != nil {
				return
			}
		}
		com.Dirty = true
		return
	case qtypes.QueryType_EQUAL:
		if com.Dirty {
			if _, err = com.WriteString(opt.Joint); err != nil {
				return
			}
		}
		if _, err = com.WriteString(sel); err != nil {
			return
		}
		if i.Negation {
			if _, err = com.WriteString(" <> "); err != nil {
				return
			}
		} else {
			if _, err = com.WriteString(" = "); err != nil {
				return
			}
		}
		if err = com.WritePlaceholder(); err != nil {
			return
		}
		com.Add(i.Value())
		com.Dirty = true
	case qtypes.QueryType_GREATER:
		if com.Dirty {
			if _, err = com.WriteString(opt.Joint); err != nil {
				return
			}
		}
		if _, err = com.WriteString(sel); err != nil {
			return
		}
		if i.Negation {
			if _, err = com.WriteString(" <= "); err != nil {
				return
			}
		} else {
			if _, err = com.WriteString(" > "); err != nil {
				return
			}
		}
		if err = com.WritePlaceholder(); err != nil {
			return
		}
		com.Add(i.Value())
		com.Dirty = true
	case qtypes.QueryType_GREATER_EQUAL:
		if com.Dirty {
			if _, err = com.WriteString(opt.Joint); err != nil {
				return
			}
		}
		if _, err = com.WriteString(sel); err != nil {
			return
		}
		if i.Negation {
			if _, err = com.WriteString(" < "); err != nil {
				return
			}
		} else {
			if _, err = com.WriteString(" >= "); err != nil {
				return
			}
		}
		if err = com.WritePlaceholder(); err != nil {
			return
		}
		com.Add(i.Value())
		com.Dirty = true
	case qtypes.QueryType_LESS:
		if com.Dirty {
			if _, err = com.WriteString(opt.Joint); err != nil {
				return
			}
		}
		if _, err = com.WriteString(sel); err != nil {
			return
		}
		if i.Negation {
			if _, err = com.WriteString(" >= "); err != nil {
				return
			}
		} else {
			if _, err = com.WriteString(" < "); err != nil {
				return
			}
		}
		if err = com.WritePlaceholder(); err != nil {
			return
		}
		com.Add(i.Value())
		com.Dirty = true
	case qtypes.QueryType_LESS_EQUAL:
		if com.Dirty {
			if _, err = com.WriteString(opt.Joint); err != nil {
				return
			}
		}
		if _, err = com.WriteString(sel); err != nil {
			return
		}
		if i.Negation {
			if _, err = com.WriteString(" > "); err != nil {
				return
			}
		} else {
			if _, err = com.WriteString(" <= "); err != nil {
				return
			}
		}
		if err = com.WritePlaceholder(); err != nil {
			return
		}
		com.Add(i.Value())
		com.Dirty = true
	case qtypes.QueryType_CONTAINS:
		if !i.Negation {
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return
				}
			}
			if _, err = com.WriteString(sel); err != nil {
				return
			}
			if _, err = com.WriteString(" @> "); err != nil {
				return
			}
			if err = com.WritePlaceholder(); err != nil {
				return
			}
			switch opt.IsJSON {
			case true:
				com.Add(JSONArrayFloat64(i.Values))
			case false:
				com.Add(pq.Float64Array(i.Values))
			}
			com.Dirty = true
		}
	case qtypes.QueryType_IS_CONTAINED_BY:
		if !i.Negation {
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return
				}
			}
			if _, err = com.WriteString(sel); err != nil {
				return
			}
			if _, err = com.WriteString(" <@ "); err != nil {
				return
			}
			if err = com.WritePlaceholder(); err != nil {
				return
			}
			switch opt.IsJSON {
			case true:
				com.Add(JSONArrayFloat64(i.Values))
			case false:
				com.Add(pq.Float64Array(i.Values))
			}
			com.Dirty = true
		}
	case qtypes.QueryType_OVERLAP:
		if !i.Negation {
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return
				}
			}
			if _, err = com.WriteString(sel); err != nil {
				return
			}
			if _, err = com.WriteString(" && "); err != nil {
				return
			}
			if err = com.WritePlaceholder(); err != nil {
				return
			}
			switch opt.IsJSON {
			case true:
				com.Add(JSONArrayFloat64(i.Values))
			case false:
				com.Add(pq.Float64Array(i.Values))
			}
			com.Dirty = true
		}
	case qtypes.QueryType_HAS_ANY_ELEMENT:
		if !i.Negation {
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return
				}
			}
			if _, err = com.WriteString(sel); err != nil {
				return
			}
			if _, err = com.WriteString(" ?| "); err != nil {
				return
			}
			if err = com.WritePlaceholder(); err != nil {
				return
			}
			switch opt.IsJSON {
			case true:
				com.Add(JSONArrayFloat64(i.Values))
			case false:
				com.Add(pq.Float64Array(i.Values))
			}
			com.Dirty = true
		}
	case qtypes.QueryType_HAS_ALL_ELEMENTS:
		if !i.Negation {
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return
				}
			}
			if _, err = com.WriteString(sel); err != nil {
				return
			}
			if _, err = com.WriteString(" ?& "); err != nil {
				return
			}
			if err = com.WritePlaceholder(); err != nil {
				return
			}
			switch opt.IsJSON {
			case true:
				com.Add(JSONArrayFloat64(i.Values))
			case false:
				com.Add(pq.Float64Array(i.Values))
			}
			com.Dirty = true
		}
	case qtypes.QueryType_HAS_ELEMENT:
		if !i.Negation {
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return
				}
			}
			if _, err = com.WriteString(sel); err != nil {
				return
			}
			if _, err = com.WriteString(" ? "); err != nil {
				return
			}
			if err = com.WritePlaceholder(); err != nil {
				return
			}
			com.Add(i.Value())
			com.Dirty = true
		}
	case qtypes.QueryType_IN:
		if len(i.Values) > 0 {
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return
				}
			}
			if _, err = com.WriteString(sel); err != nil {
				return
			}
			if i.Negation {
				if _, err = com.WriteString(" NOT IN ("); err != nil {
					return
				}
			} else {
				if _, err = com.WriteString(" IN ("); err != nil {
					return
				}
			}
			for i, v := range i.Values {
				if i != 0 {
					if _, err = com.WriteString(","); err != nil {
						return
					}
				}
				if err = com.WritePlaceholder(); err != nil {
					return
				}
				com.Add(v)
				com.Dirty = true
			}
			if _, err = com.WriteString(")"); err != nil {
				return
			}
		}
	case qtypes.QueryType_BETWEEN:
		if com.Dirty {
			if _, err = com.WriteString(opt.Joint); err != nil {
				return
			}
		}
		if _, err = com.WriteString(sel); err != nil {
			return
		}
		if i.Negation {
			if _, err = com.WriteString(" <= "); err != nil {
				return
			}
		} else {
			if _, err = com.WriteString(" > "); err != nil {
				return
			}
		}
		if err = com.WritePlaceholder(); err != nil {
			return
		}
		com.Add(i.Values[0])
		if _, err = com.WriteString(" AND "); err != nil {
			return
		}
		if _, err = com.WriteString(sel); err != nil {
			return
		}
		if i.Negation {
			if _, err = com.WriteString(" >= "); err != nil {
				return
			}
		} else {
			if _, err = com.WriteString(" < "); err != nil {
				return
			}
		}
		if err = com.WritePlaceholder(); err != nil {
			return
		}
		com.Add(i.Values[1])
		com.Dirty = true
	default:
		return fmt.Errorf("pqtgo: unknown int64 query type %!s(MISSING)", i.Type.String())
	}

	if com.Dirty {
		if opt.Cast != "" {
			if _, err = com.WriteString(opt.Cast); err != nil {
				return
			}
		}
	}

	return
}
func QueryTimestampWhereClause(t *qtypes.Timestamp, sel string, com *Composer, opt *CompositionOpts) error {
	if t == nil || !t.Valid {
		return nil
	}
	v := t.Value()
	if v != nil {
		vv1, err := ptypes.Timestamp(v)
		if err != nil {
			return err
		}
		switch t.Type {
		case qtypes.QueryType_NULL:
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return err
				}
			}
			com.WriteString(sel)
			if t.Negation {
				com.WriteString(" IS NOT NULL ")
			} else {
				com.WriteString(" IS NULL ")
			}
		case qtypes.QueryType_EQUAL:
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return err
				}
			}
			com.WriteString(sel)
			if t.Negation {
				com.WriteString(" <> ")
			} else {
				com.WriteString(" = ")
			}
			com.WritePlaceholder()
			com.Add(t.Value())
		case qtypes.QueryType_GREATER:
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return err
				}
			}
			com.WriteString(sel)
			com.WriteString(">")
			com.WritePlaceholder()
			com.Add(t.Value())
		case qtypes.QueryType_GREATER_EQUAL:
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return err
				}
			}
			com.WriteString(sel)
			com.WriteString(">=")
			com.WritePlaceholder()
			com.Add(t.Value())
		case qtypes.QueryType_LESS:
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return err
				}
			}
			com.WriteString(sel)
			com.WriteString(" < ")
			com.WritePlaceholder()
			com.Add(t.Value())
		case qtypes.QueryType_LESS_EQUAL:
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return err
				}
			}
			com.WriteString(sel)
			com.WriteString(" <= ")
			com.WritePlaceholder()
			com.Add(t.Value())
		case qtypes.QueryType_IN:
			if len(t.Values) > 0 {
				if com.Dirty {
					if _, err = com.WriteString(opt.Joint); err != nil {
						return err
					}
				}
				com.WriteString(sel)
				com.WriteString(" IN (")
				for i, v := range t.Values {
					if i != 0 {
						com.WriteString(", ")
					}
					com.WritePlaceholder()
					com.Add(v)
				}
				com.WriteString(") ")
			}
		case qtypes.QueryType_BETWEEN:
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return err
				}
			}
			v2 := t.Values[1]
			if v2 != nil {
				vv2, err := ptypes.Timestamp(v2)
				if err != nil {
					return err
				}
				com.WriteString(sel)
				com.WriteString(" > ")
				com.WritePlaceholder()
				com.Add(vv1)
				com.WriteString(" AND ")
				com.WriteString(sel)
				com.WriteString(" < ")
				com.WritePlaceholder()
				com.Add(vv2)
			}
		}
	}
	return nil
}
func QueryStringWhereClause(s *qtypes.String, sel string, com *Composer, opt *CompositionOpts) (err error) {
	if s == nil || !s.Valid {
		return
	}
	switch s.Type {
	case qtypes.QueryType_NULL:
		if com.Dirty {
			if _, err = com.WriteString(opt.Joint); err != nil {
				return
			}
		}
		if _, err = com.WriteString(sel); err != nil {
			return
		}
		if s.Negation {
			if _, err = com.WriteString(" IS NOT NULL"); err != nil {
				return
			}
		} else {
			if _, err = com.WriteString(" IS NULL"); err != nil {
				return
			}
		}
		com.Dirty = true
		return // cannot be casted so simply return
	case qtypes.QueryType_EQUAL:
		if com.Dirty {
			if _, err = com.WriteString(opt.Joint); err != nil {
				return
			}
		}
		if _, err = com.WriteString(sel); err != nil {
			return
		}
		if s.Negation {
			if _, err = com.WriteString(" <> "); err != nil {
				return
			}
		} else {
			if _, err = com.WriteString(" = "); err != nil {
				return
			}
		}
		if opt.PlaceholderFunc != "" {
			if _, err = com.WriteString(opt.PlaceholderFunc); err != nil {
				return
			}
			if _, err = com.WriteString("("); err != nil {
				return
			}
		}
		if err = com.WritePlaceholder(); err != nil {
			return
		}
		com.Add(s.Value())
		com.Dirty = true
	case qtypes.QueryType_SUBSTRING:
		if com.Dirty {
			if _, err = com.WriteString(opt.Joint); err != nil {
				return
			}
		}
		if _, err = com.WriteString(sel); err != nil {
			return
		}
		if s.Negation {
			if _, err = com.WriteString(" NOT "); err != nil {
				return
			}
		}
		if s.Insensitive {
			if _, err = com.WriteString(" ILIKE "); err != nil {
				return
			}
		} else {
			if _, err = com.WriteString(" LIKE "); err != nil {
				return
			}
		}

		if opt.PlaceholderFunc != "" {
			if _, err = com.WriteString(opt.PlaceholderFunc); err != nil {
				return
			}
			if _, err = com.WriteString("("); err != nil {
				return
			}
		}
		if err = com.WritePlaceholder(); err != nil {
			return
		}

		com.Add(fmt.Sprintf("%%!s(MISSING)%", s.Value()))
		com.Dirty = true
	case qtypes.QueryType_HAS_PREFIX:
		if com.Dirty {
			if _, err = com.WriteString(opt.Joint); err != nil {
				return
			}
		}
		if _, err = com.WriteString(sel); err != nil {
			return
		}
		if s.Negation {
			if _, err = com.WriteString(" NOT "); err != nil {
				return
			}
		}
		if s.Insensitive {
			if _, err = com.WriteString(" ILIKE "); err != nil {
				return
			}
		} else {
			if _, err = com.WriteString(" LIKE "); err != nil {
				return
			}
		}
		if opt.PlaceholderFunc != "" {
			if _, err = com.WriteString(opt.PlaceholderFunc); err != nil {
				return
			}
			if _, err = com.WriteString("("); err != nil {
				return
			}
		}
		if err = com.WritePlaceholder(); err != nil {
			return
		}

		com.Add(fmt.Sprintf("%!s(MISSING)%", s.Value()))
		com.Dirty = true
	case qtypes.QueryType_HAS_SUFFIX:
		if com.Dirty {
			if _, err = com.WriteString(opt.Joint); err != nil {
				return
			}
		}
		if _, err = com.WriteString(sel); err != nil {
			return
		}
		if s.Negation {
			if _, err = com.WriteString(" NOT "); err != nil {
				return
			}

		}
		if s.Insensitive {
			if _, err = com.WriteString(" ILIKE "); err != nil {
				return
			}
		} else {
			if _, err = com.WriteString(" LIKE "); err != nil {
				return
			}
		}
		if opt.PlaceholderFunc != "" {
			if _, err = com.WriteString(opt.PlaceholderFunc); err != nil {
				return
			}
			if _, err = com.WriteString("("); err != nil {
				return
			}
		}
		if err = com.WritePlaceholder(); err != nil {
			return
		}

		com.Add(fmt.Sprintf("%%!s(MISSING)", s.Value()))
		com.Dirty = true
	case qtypes.QueryType_CONTAINS:
		if !s.Negation {
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return
				}
			}
			if _, err = com.WriteString(sel); err != nil {
				return
			}
			if _, err = com.WriteString(" @> "); err != nil {
				return
			}
			if opt.PlaceholderFunc != "" {
				if _, err = com.WriteString(opt.PlaceholderFunc); err != nil {
					return
				}
				if _, err = com.WriteString("("); err != nil {
					return
				}
			}
			if err = com.WritePlaceholder(); err != nil {
				return
			}

			com.Add(JSONArrayString(s.Values))
			com.Dirty = true
		}
	case qtypes.QueryType_IS_CONTAINED_BY:
		if !s.Negation {
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return
				}
			}
			if _, err = com.WriteString(sel); err != nil {
				return
			}
			if _, err = com.WriteString(" <@ "); err != nil {
				return
			}
			if opt.PlaceholderFunc != "" {
				if _, err = com.WriteString(opt.PlaceholderFunc); err != nil {
					return
				}
				if _, err = com.WriteString("("); err != nil {
					return
				}
			}
			if err = com.WritePlaceholder(); err != nil {
				return
			}

			com.Add(s.Value())
			com.Dirty = true
		}
	case qtypes.QueryType_OVERLAP:
		if !s.Negation {
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return
				}
			}
			if _, err = com.WriteString(sel); err != nil {
				return
			}
			if _, err = com.WriteString(" && "); err != nil {
				return
			}
			if opt.PlaceholderFunc != "" {
				if _, err = com.WriteString(opt.PlaceholderFunc); err != nil {
					return
				}
				if _, err = com.WriteString("("); err != nil {
					return
				}
			}
			if err = com.WritePlaceholder(); err != nil {
				return
			}

			com.Add(s.Value())
			com.Dirty = true
		}
	case qtypes.QueryType_HAS_ANY_ELEMENT:
		if !s.Negation {
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return
				}
			}
			if _, err = com.WriteString(sel); err != nil {
				return
			}
			if _, err = com.WriteString(" ?| "); err != nil {
				return
			}
			if err = com.WritePlaceholder(); err != nil {
				return
			}
			com.Add(pq.StringArray(s.Values))
			com.Dirty = true
		}
	case qtypes.QueryType_HAS_ALL_ELEMENTS:
		if !s.Negation {
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return
				}
			}
			if _, err = com.WriteString(sel); err != nil {
				return
			}
			if _, err = com.WriteString(" ?& "); err != nil {
				return
			}
			if err = com.WritePlaceholder(); err != nil {
				return
			}
			com.Add(pq.StringArray(s.Values))
			com.Dirty = true
		}
	case qtypes.QueryType_IN:
		if len(s.Values) > 0 {
			if com.Dirty {
				if _, err = com.WriteString(opt.Joint); err != nil {
					return
				}
			}
			if _, err = com.WriteString(sel); err != nil {
				return
			}
			if s.Negation {
				if _, err = com.WriteString(" NOT IN ("); err != nil {
					return
				}
			} else {
				if _, err = com.WriteString(" IN ("); err != nil {
					return
				}
			}
			for i, v := range s.Values {
				if i != 0 {
					if _, err = com.WriteString(","); err != nil {
						return
					}
				}
				if err = com.WritePlaceholder(); err != nil {
					return
				}
				com.Add(v)
				com.Dirty = true
			}
			if _, err = com.WriteString(")"); err != nil {
				return
			}
		}
	default:
		return fmt.Errorf("pqtgo: unknown string query type %!s(MISSING)", s.Type.String())
	}

	switch {
	case com.Dirty && opt.Cast != "":
		if _, err = com.WriteString(opt.Cast); err != nil {
			return
		}
	case com.Dirty && opt.PlaceholderFunc != "":
		if _, err = com.WriteString(")"); err != nil {
			return
		}
	case com.Dirty:
	}
	return
}

const SQL = `
-- do not modify, generated by pqt

CREATE SCHEMA IF NOT EXISTS charon; 

CREATE TABLE IF NOT EXISTS charon.user (
	confirmation_token BYTEA,
	created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
	created_by BIGINT,
	first_name TEXT NOT NULL,
	id BIGSERIAL,
	is_active BOOL DEFAULT FALSE NOT NULL,
	is_confirmed BOOL DEFAULT FALSE NOT NULL,
	is_staff BOOL DEFAULT FALSE NOT NULL,
	is_superuser BOOL DEFAULT FALSE NOT NULL,
	last_login_at TIMESTAMPTZ,
	last_name TEXT NOT NULL,
	password BYTEA NOT NULL,
	updated_at TIMESTAMPTZ,
	updated_by BIGINT,
	username TEXT NOT NULL,

	CONSTRAINT "charon.user_created_by_fkey" FOREIGN KEY (created_by) REFERENCES charon.user (id),
	CONSTRAINT "charon.user_id_pkey" PRIMARY KEY (id),
	CONSTRAINT "charon.user_updated_by_fkey" FOREIGN KEY (updated_by) REFERENCES charon.user (id),
	CONSTRAINT "charon.user_username_key" UNIQUE (username)
);

CREATE TABLE IF NOT EXISTS charon.group (
	created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
	created_by BIGINT,
	description TEXT,
	id BIGSERIAL,
	name TEXT NOT NULL,
	updated_at TIMESTAMPTZ,
	updated_by BIGINT,

	CONSTRAINT "charon.group_created_by_fkey" FOREIGN KEY (created_by) REFERENCES charon.user (id),
	CONSTRAINT "charon.group_id_pkey" PRIMARY KEY (id),
	CONSTRAINT "charon.group_name_key" UNIQUE (name),
	CONSTRAINT "charon.group_updated_by_fkey" FOREIGN KEY (updated_by) REFERENCES charon.user (id)
);

CREATE TABLE IF NOT EXISTS charon.permission (
	action TEXT NOT NULL,
	created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
	id BIGSERIAL,
	module TEXT NOT NULL,
	subsystem TEXT NOT NULL,
	updated_at TIMESTAMPTZ,

	CONSTRAINT "charon.permission_id_pkey" PRIMARY KEY (id),
	CONSTRAINT "charon.permission_subsystem_module_action_key" UNIQUE (subsystem, module, action)
);

CREATE TABLE IF NOT EXISTS charon.user_groups (
	created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
	created_by BIGINT,
	group_id BIGINT NOT NULL,
	updated_at TIMESTAMPTZ,
	updated_by BIGINT,
	user_id BIGINT NOT NULL,

	CONSTRAINT "charon.user_groups_created_by_fkey" FOREIGN KEY (created_by) REFERENCES charon.user (id),
	CONSTRAINT "charon.user_groups_updated_by_fkey" FOREIGN KEY (updated_by) REFERENCES charon.user (id),
	CONSTRAINT "charon.user_groups_user_id_fkey" FOREIGN KEY (user_id) REFERENCES charon.user (id),
	CONSTRAINT "charon.user_groups_group_id_fkey" FOREIGN KEY (group_id) REFERENCES charon.group (id),
	CONSTRAINT "charon.user_groups_user_id_group_id_key" UNIQUE (user_id, group_id)
);

CREATE TABLE IF NOT EXISTS charon.group_permissions (
	created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
	created_by BIGINT,
	group_id BIGINT NOT NULL,
	permission_action TEXT NOT NULL,
	permission_module TEXT NOT NULL,
	permission_subsystem TEXT NOT NULL,
	updated_at TIMESTAMPTZ,
	updated_by BIGINT,

	CONSTRAINT "charon.group_permissions_created_by_fkey" FOREIGN KEY (created_by) REFERENCES charon.user (id),
	CONSTRAINT "charon.group_permissions_updated_by_fkey" FOREIGN KEY (updated_by) REFERENCES charon.user (id),
	CONSTRAINT "charon.group_permissions_group_id_fkey" FOREIGN KEY (group_id) REFERENCES charon.group (id),
	CONSTRAINT "charon.group_permissions_subsystem_module_action_fkey" FOREIGN KEY (permission_subsystem, permission_module, permission_action) REFERENCES charon.permission (subsystem, module, action),
	CONSTRAINT "charon.group_permissions_group_id_subsystem_module_action_key" UNIQUE (group_id, permission_subsystem, permission_module, permission_action)
);

CREATE TABLE IF NOT EXISTS charon.user_permissions (
	created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
	created_by BIGINT,
	permission_action TEXT NOT NULL,
	permission_module TEXT NOT NULL,
	permission_subsystem TEXT NOT NULL,
	updated_at TIMESTAMPTZ,
	updated_by BIGINT,
	user_id BIGINT NOT NULL,

	CONSTRAINT "charon.user_permissions_created_by_fkey" FOREIGN KEY (created_by) REFERENCES charon.user (id),
	CONSTRAINT "charon.user_permissions_updated_by_fkey" FOREIGN KEY (updated_by) REFERENCES charon.user (id),
	CONSTRAINT "charon.user_permissions_user_id_fkey" FOREIGN KEY (user_id) REFERENCES charon.user (id),
	CONSTRAINT "charon.user_permissions_subsystem_module_action_fkey" FOREIGN KEY (permission_subsystem, permission_module, permission_action) REFERENCES charon.permission (subsystem, module, action),
	CONSTRAINT "charon.user_permissions_user_id_subsystem_module_action_key" UNIQUE (user_id, permission_subsystem, permission_module, permission_action)
);

`
