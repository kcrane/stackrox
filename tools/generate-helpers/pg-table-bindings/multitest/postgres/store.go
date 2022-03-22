// Code generated by pg-bindings generator. DO NOT EDIT.

package postgres

import (
	"context"
	"reflect"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/central/metrics"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	log = logging.LoggerForModule()
)

const (
	baseTable  = "multikey"
	countStmt  = "SELECT COUNT(*) FROM multikey"
	existsStmt = "SELECT EXISTS(SELECT 1 FROM multikey WHERE Key1 = $1 AND Key2 = $2)"

	getStmt    = "SELECT serialized FROM multikey WHERE Key1 = $1 AND Key2 = $2"
	deleteStmt = "DELETE FROM multikey WHERE Key1 = $1 AND Key2 = $2"
	walkStmt   = "SELECT serialized FROM multikey"
)

var (
	schema = walker.Walk(reflect.TypeOf((*storage.TestMultiKeyStruct)(nil)), baseTable)
)

func init() {
	globaldb.RegisterTable(schema)
}

type Store interface {
	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, key1 string, key2 string) (bool, error)
	Get(ctx context.Context, key1 string, key2 string) (*storage.TestMultiKeyStruct, bool, error)
	Upsert(ctx context.Context, obj *storage.TestMultiKeyStruct) error
	UpsertMany(ctx context.Context, objs []*storage.TestMultiKeyStruct) error
	Delete(ctx context.Context, key1 string, key2 string) error

	Walk(ctx context.Context, fn func(obj *storage.TestMultiKeyStruct) error) error

	AckKeysIndexed(ctx context.Context, keys ...string) error
	GetKeysToIndex(ctx context.Context) ([]string, error)
}

type storeImpl struct {
	db *pgxpool.Pool
}

func createTableMultikey(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists multikey (
    Key1 varchar,
    Key2 varchar,
    StringSlice text[],
    Bool bool,
    Uint64 integer,
    Int64 integer,
    Float numeric,
    Labels jsonb,
    Timestamp timestamp,
    Enum integer,
    Enums int[],
    String_ varchar,
    IntSlice int[],
    Embedded_Embedded varchar,
    Oneofstring varchar,
    Oneofnested_Nested varchar,
    serialized bytea,
    PRIMARY KEY(Key1, Key2)
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

	createTableMultikeyNested(ctx, db)
}

func createTableMultikeyNested(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists multikey_Nested (
    multikey_Key1 varchar,
    multikey_Key2 varchar,
    idx integer,
    Nested varchar,
    IsNested bool,
    Int64 integer,
    Nested2_Nested2 varchar,
    Nested2_IsNested bool,
    Nested2_Int64 integer,
    PRIMARY KEY(multikey_Key1, multikey_Key2, idx),
    CONSTRAINT fk_parent_table FOREIGN KEY (multikey_Key1, multikey_Key2) REFERENCES multikey(Key1, Key2) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{

		"create index if not exists multikeyNested_idx on multikey_Nested using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

}

func insertIntoMultikey(ctx context.Context, tx pgx.Tx, obj *storage.TestMultiKeyStruct) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		obj.GetKey1(),
		obj.GetKey2(),
		obj.GetStringSlice(),
		obj.GetBool(),
		obj.GetUint64(),
		obj.GetInt64(),
		obj.GetFloat(),
		obj.GetLabels(),
		pgutils.NilOrStringTimestamp(obj.GetTimestamp()),
		obj.GetEnum(),
		obj.GetEnums(),
		obj.GetString_(),
		obj.GetIntSlice(),
		obj.GetEmbedded().GetEmbedded(),
		obj.GetOneofstring(),
		obj.GetOneofnested().GetNested(),
		serialized,
	}

	finalStr := "INSERT INTO multikey (Key1, Key2, StringSlice, Bool, Uint64, Int64, Float, Labels, Timestamp, Enum, Enums, String_, IntSlice, Embedded_Embedded, Oneofstring, Oneofnested_Nested, serialized) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) ON CONFLICT(Key1, Key2) DO UPDATE SET Key1 = EXCLUDED.Key1, Key2 = EXCLUDED.Key2, StringSlice = EXCLUDED.StringSlice, Bool = EXCLUDED.Bool, Uint64 = EXCLUDED.Uint64, Int64 = EXCLUDED.Int64, Float = EXCLUDED.Float, Labels = EXCLUDED.Labels, Timestamp = EXCLUDED.Timestamp, Enum = EXCLUDED.Enum, Enums = EXCLUDED.Enums, String_ = EXCLUDED.String_, IntSlice = EXCLUDED.IntSlice, Embedded_Embedded = EXCLUDED.Embedded_Embedded, Oneofstring = EXCLUDED.Oneofstring, Oneofnested_Nested = EXCLUDED.Oneofnested_Nested, serialized = EXCLUDED.serialized"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	var query string

	for childIdx, child := range obj.GetNested() {
		if err := insertIntoMultikeyNested(ctx, tx, child, obj.GetKey1(), obj.GetKey2(), childIdx); err != nil {
			return err
		}
	}

	query = "delete from multikey_Nested where multikey_Key1 = $1 AND multikey_Key2 = $2 AND idx >= $3"
	_, err = tx.Exec(ctx, query, obj.GetKey1(), obj.GetKey2(), len(obj.GetNested()))
	if err != nil {
		return err
	}
	return nil
}

func insertIntoMultikeyNested(ctx context.Context, tx pgx.Tx, obj *storage.TestMultiKeyStruct_Nested, multikey_Key1 string, multikey_Key2 string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		multikey_Key1,
		multikey_Key2,
		idx,
		obj.GetNested(),
		obj.GetIsNested(),
		obj.GetInt64(),
		obj.GetNested2().GetNested2(),
		obj.GetNested2().GetIsNested(),
		obj.GetNested2().GetInt64(),
	}

	finalStr := "INSERT INTO multikey_Nested (multikey_Key1, multikey_Key2, idx, Nested, IsNested, Int64, Nested2_Nested2, Nested2_IsNested, Nested2_Int64) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) ON CONFLICT(multikey_Key1, multikey_Key2, idx) DO UPDATE SET multikey_Key1 = EXCLUDED.multikey_Key1, multikey_Key2 = EXCLUDED.multikey_Key2, idx = EXCLUDED.idx, Nested = EXCLUDED.Nested, IsNested = EXCLUDED.IsNested, Int64 = EXCLUDED.Int64, Nested2_Nested2 = EXCLUDED.Nested2_Nested2, Nested2_IsNested = EXCLUDED.Nested2_IsNested, Nested2_Int64 = EXCLUDED.Nested2_Int64"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

// New returns a new Store instance using the provided sql instance.
func New(ctx context.Context, db *pgxpool.Pool) Store {
	createTableMultikey(ctx, db)

	return &storeImpl{
		db: db,
	}
}

func (s *storeImpl) upsert(ctx context.Context, objs ...*storage.TestMultiKeyStruct) error {
	conn, release := s.acquireConn(ctx, ops.Get, "TestMultiKeyStruct")
	defer release()

	for _, obj := range objs {
		tx, err := conn.Begin(ctx)
		if err != nil {
			return err
		}

		if err := insertIntoMultikey(ctx, tx, obj); err != nil {
			if err := tx.Rollback(ctx); err != nil {
				return err
			}
			return err
		}
		if err := tx.Commit(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (s *storeImpl) Upsert(ctx context.Context, obj *storage.TestMultiKeyStruct) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Upsert, "TestMultiKeyStruct")

	return s.upsert(ctx, obj)
}

func (s *storeImpl) UpsertMany(ctx context.Context, objs []*storage.TestMultiKeyStruct) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.UpdateMany, "TestMultiKeyStruct")

	return s.upsert(ctx, objs...)
}

// Count returns the number of objects in the store
func (s *storeImpl) Count(ctx context.Context) (int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Count, "TestMultiKeyStruct")

	row := s.db.QueryRow(ctx, countStmt)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// Exists returns if the id exists in the store
func (s *storeImpl) Exists(ctx context.Context, key1 string, key2 string) (bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Exists, "TestMultiKeyStruct")

	row := s.db.QueryRow(ctx, existsStmt, key1, key2)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, pgutils.ErrNilIfNoRows(err)
	}
	return exists, nil
}

// Get returns the object, if it exists from the store
func (s *storeImpl) Get(ctx context.Context, key1 string, key2 string) (*storage.TestMultiKeyStruct, bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Get, "TestMultiKeyStruct")

	conn, release := s.acquireConn(ctx, ops.Get, "TestMultiKeyStruct")
	defer release()

	row := conn.QueryRow(ctx, getStmt, key1, key2)
	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, false, pgutils.ErrNilIfNoRows(err)
	}

	var msg storage.TestMultiKeyStruct
	if err := proto.Unmarshal(data, &msg); err != nil {
		return nil, false, err
	}
	return &msg, true, nil
}

func (s *storeImpl) acquireConn(ctx context.Context, op ops.Op, typ string) (*pgxpool.Conn, func()) {
	defer metrics.SetAcquireDBConnDuration(time.Now(), op, typ)
	conn, err := s.db.Acquire(ctx)
	if err != nil {
		panic(err)
	}
	return conn, conn.Release
}

// Delete removes the specified ID from the store
func (s *storeImpl) Delete(ctx context.Context, key1 string, key2 string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Remove, "TestMultiKeyStruct")

	conn, release := s.acquireConn(ctx, ops.Remove, "TestMultiKeyStruct")
	defer release()

	if _, err := conn.Exec(ctx, deleteStmt, key1, key2); err != nil {
		return err
	}
	return nil
}

// Walk iterates over all of the objects in the store and applies the closure
func (s *storeImpl) Walk(ctx context.Context, fn func(obj *storage.TestMultiKeyStruct) error) error {
	rows, err := s.db.Query(ctx, walkStmt)
	if err != nil {
		return pgutils.ErrNilIfNoRows(err)
	}
	defer rows.Close()
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return err
		}
		var msg storage.TestMultiKeyStruct
		if err := proto.Unmarshal(data, &msg); err != nil {
			return err
		}
		if err := fn(&msg); err != nil {
			return err
		}
	}
	return nil
}

//// Used for testing

func dropTableMultikey(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS multikey CASCADE")
	dropTableMultikeyNested(ctx, db)

}

func dropTableMultikeyNested(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS multikey_Nested CASCADE")

}

func Destroy(ctx context.Context, db *pgxpool.Pool) {
	dropTableMultikey(ctx, db)
}

//// Stubs for satisfying legacy interfaces

// AckKeysIndexed acknowledges the passed keys were indexed
func (s *storeImpl) AckKeysIndexed(ctx context.Context, keys ...string) error {
	return nil
}

// GetKeysToIndex returns the keys that need to be indexed
func (s *storeImpl) GetKeysToIndex(ctx context.Context) ([]string, error) {
	return nil, nil
}
