package services

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"reflect"
)

type Service[T any] interface {
	FindAll(ctx context.Context) ([]T, error)
	FindOne(ctx context.Context, id int) (T, error)
}

type mySql[T any] struct {
	mySql  *sql.DB
	dbName string
}

func NewService[T any](mySqlConn *sql.DB, dbName string) Service[T] {
	return &mySql[T]{mySql: mySqlConn, dbName: dbName}
}

func (m *mySql[T]) FindAll(ctx context.Context) ([]T, error) {
	rows, err := m.mySql.QueryContext(ctx, fmt.Sprintf("SELECT * FROM %s", m.dbName))
	if err != nil {
		log.Println("Failed to query:", err)
		return nil, err
	}
	defer rows.Close()

	var results []T
	for rows.Next() {
		// Create a new instance of T
		t := reflect.New(reflect.TypeOf((*T)(nil)).Elem()).Interface()

		// Get the fields of the struct
		val := reflect.ValueOf(t).Elem()
		fields := make([]interface{}, val.NumField())
		for i := 0; i < val.NumField(); i++ {
			fields[i] = val.Field(i).Addr().Interface()
		}

		// Scan the row into the struct fields
		if err := rows.Scan(fields...); err != nil {
			log.Println("Failed to scan row:", err)
			return nil, err
		}

		// Append the result to the results slice
		results = append(results, *t.(*T))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (m *mySql[T]) FindOne(ctx context.Context, id int) (T, error) {
	row := m.mySql.QueryRowContext(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id = ?", m.dbName), id)

	// Create a new instance of T
	t := reflect.New(reflect.TypeOf((*T)(nil)).Elem()).Interface()

	// Get the fields of the struct
	val := reflect.ValueOf(t).Elem()
	fields := make([]interface{}, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		fields[i] = val.Field(i).Addr().Interface()
	}

	// Scan the row into the struct fields
	if err := row.Scan(fields...); err != nil {
		log.Println("Failed to scan row:", err)
		var zero T
		return zero, err
	}

	return *t.(*T), nil
}
