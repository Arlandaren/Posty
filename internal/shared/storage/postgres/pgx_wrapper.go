package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"reflect"
)

// DB - интерфейс для нашей обёртки, чтобы было удобно заменять реализацию (например, для моков в тестах)
type DB interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
}

// Wrapper - структура, которая содержит пул соединений с базой данных
type Wrapper struct {
	pool *pgxpool.Pool
}

// NewWrapper - функция для создания новой обёртки с пулом соединений
func NewWrapper(pool *pgxpool.Pool) *Wrapper {
	return &Wrapper{pool: pool}
}

// QueryRow - обёртка для метода QueryRow
func (w *Wrapper) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return w.pool.QueryRow(ctx, sql, args...)
}

// Query - обёртка для метода Query
func (w *Wrapper) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return w.pool.Query(ctx, sql, args...)
}

// Exec - обёртка для метода Exec
func (w *Wrapper) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	return w.pool.Exec(ctx, sql, args...)
}

// Get - функция для получения одного результата и его сканирования
func (w *Wrapper) Get(ctx context.Context, dest interface{}, sql string, args ...interface{}) error {
	row := w.pool.QueryRow(ctx, sql, args...)
	return row.Scan(dest)
}

// Select - функция для получения нескольких результатов и их сканирования в слайс
func (w *Wrapper) Select(ctx context.Context, dest interface{}, sql string, args ...interface{}) error {
	rows, err := w.pool.Query(ctx, sql, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	destVal := reflect.ValueOf(dest)
	if destVal.Kind() != reflect.Ptr || destVal.Elem().Kind() != reflect.Slice {
		return errors.New("dest must be a pointer to a slice")
	}

	sliceVal := destVal.Elem()
	elemType := sliceVal.Type().Elem()

	for rows.Next() {
		elemPtr := reflect.New(elemType)
		scanDest := elemPtr.Interface()

		if err := rows.Scan(scanDest); err != nil {
			return err
		}

		sliceVal.Set(reflect.Append(sliceVal, elemPtr.Elem()))
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}
