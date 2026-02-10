package postgresql

import (
	"context"
	"time"

	"github.com/mpsdantas/bottle/pkg/core/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Expression interface {
	Build(builder Builder)
}

type Builder interface {
	Writer
	WriteQuoted(field interface{})
	AddVar(Writer, ...interface{})
	AddError(error) error
}

type Writer interface {
	WriteByte(byte) error
	WriteString(string) (int, error)
}

type Returning struct {
	clause.Returning
}

type Postgresql struct {
	db *gorm.DB
}

func New(opt ...OptionFunc) *Postgresql {
	ctx := context.Background()

	opts := options{}

	for _, optionFunc := range opt {
		optionFunc(&opts)
	}

	dbg, err := gorm.Open(postgres.New(postgres.Config{
		PreferSimpleProtocol: opts.preferSimpleProtocol,
		Conn:                 opts.db,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(ctx, "could not start gorm db",
			log.Err(err),
		)
	}

	sqlDb, err := dbg.DB()
	if err != nil {
		log.Fatal(ctx, "could not open db",
			log.Err(err),
		)
	}

	sqlDb.SetMaxIdleConns(2)
	sqlDb.SetMaxOpenConns(10)
	sqlDb.SetConnMaxLifetime(time.Hour)

	return &Postgresql{db: dbg}
}

type Result struct {
	err          error
	rowsAffected int64
}

func (r *Result) Error() error {
	return r.err
}

func (r *Result) RowsAffected() int64 {
	return r.rowsAffected
}

func (w *Postgresql) withContext(ctx context.Context) *gorm.DB {
	return w.db.WithContext(ctx)
}

func (w *Postgresql) Create(ctx context.Context, value any) *Result {
	tx := w.withContext(ctx).Create(value)
	return &Result{err: tx.Error, rowsAffected: tx.RowsAffected}
}

func (w *Postgresql) Save(ctx context.Context, value any) *Result {
	tx := w.withContext(ctx).Save(value)
	return &Result{err: tx.Error, rowsAffected: tx.RowsAffected}
}

func (w *Postgresql) Find(ctx context.Context, dest any, conds ...any) *Result {
	tx := w.withContext(ctx).Find(dest, conds...)
	return &Result{err: tx.Error, rowsAffected: tx.RowsAffected}
}

func (w *Postgresql) First(ctx context.Context, dest any, conds ...any) *Result {
	tx := w.withContext(ctx).First(dest, conds...)
	return &Result{err: tx.Error, rowsAffected: tx.RowsAffected}
}

func (w *Postgresql) Update(ctx context.Context, column string, value any) *Result {
	tx := w.withContext(ctx).Model(value).Update(column, value)
	return &Result{err: tx.Error, rowsAffected: tx.RowsAffected}
}

func (w *Postgresql) Updates(values interface{}) *Postgresql {
	tx := w.db.Updates(values)
	return &Postgresql{db: tx}
}

func (w *Postgresql) Delete(ctx context.Context, value any, conds ...any) *Result {
	tx := w.withContext(ctx).Delete(value, conds...)
	return &Result{err: tx.Error, rowsAffected: tx.RowsAffected}
}

func (w *Postgresql) Exec(ctx context.Context, sql string, values ...any) *Result {
	tx := w.withContext(ctx).Exec(sql, values...)
	return &Result{err: tx.Error, rowsAffected: tx.RowsAffected}
}

func (w *Postgresql) Transaction(ctx context.Context, fn func(*Postgresql) error) error {
	return w.withContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(&Postgresql{db: tx})
	})
}

func (w *Postgresql) Where(query interface{}, args ...interface{}) *Postgresql {
	tx := w.db.Where(query, args...)
	return &Postgresql{db: tx}
}

func (w *Postgresql) Model(ctx context.Context, value any) *Postgresql {
	tx := w.withContext(ctx).Model(value)
	return &Postgresql{db: tx}
}

func (w *Postgresql) Clauses(ctx context.Context, conds ...clause.Expression) *Postgresql {
	tx := w.withContext(ctx).Clauses(conds...)
	return &Postgresql{db: tx}
}

func (w *Postgresql) Migrate(dst ...interface{}) error {
	return w.db.AutoMigrate(dst...)
}
