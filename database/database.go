package database

import (
	"context"
	"github.com/Nimor111/grader/config"
	"github.com/go-pg/pg"
)

type Database interface {
	GetDB() *pg.DB
	Close() error
	CreateSchema(models []interface{}) error
	Transaction(ctx context.Context, operation func(ctx context.Context, tx *pg.Tx) error) error
}

type database struct {
	db *pg.DB
}

func (d *database) GetDB() *pg.DB {
	return d.db
}

func (d *database) Close() error {
	if err := d.db.Close(); err != nil {
		return err
	}

	return nil
}

func NewDB(cfg config.Config) Database {
	db := pg.Connect(&pg.Options{
		User:     cfg.Postgres.DBUser,
		Password: cfg.Postgres.DBPassword,
		Database: cfg.Postgres.DBName,
	})

	return &database{
		db: db,
	}
}

func (d *database) Transaction(ctx context.Context, operation func(ctx context.Context, tx *pg.Tx) error) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}

	if operationErr := operation(ctx, tx); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return operationErr
	}

	return tx.Commit()
}

func (d *database) CreateSchema(models []interface{}) error {
	for _, model := range models {
		err := d.db.CreateTable(model, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
