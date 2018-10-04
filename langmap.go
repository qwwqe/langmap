package langmap

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
)

var (
	Version = "0.0.0"
)

type Injectable interface {
	Inject(map[string]interface{})
}

type Identifiable interface {
	GetId() uint
}

type DatabaseWriter interface {
	Db() *gorp.DbMap
}

type IdentifiableTable interface {
	Identifiable
	TableName() string
}

type RoutableResource interface {
	Create(*gin.Context)
	Delete(*gin.Context)
	Get(*gin.Context)
	GetOne(*gin.Context)
	Update(*gin.Context)
}

type Preloadable interface {
	Preload(*gorp.DbMap) error
}

func LoadOne(db *gorp.DbMap, r interface {
	IdentifiableTable
	Preloadable
}, id uint) error {
	if err := db.SelectOne(r, "select * from "+r.TableName()+" where id = $1", id); err != nil {
		return err
	}
	r.Preload(db)
	return nil
}

func InsertOne(db *gorp.DbMap, r interface{}) error {
	if err := db.Insert(r); err != nil {
		return err
	}
	return nil
}

func AddForeignKey(db *gorp.DbMap, table *gorp.TableMap, key string, reference *gorp.TableMap, column string, ordinal uint) error {
	if _, err := db.Exec(fmt.Sprintf(
		"alter table %s add constraint %s foreign key (%s) references %s(%s);",
		db.Dialect.QuoteField(table.TableName),
		db.Dialect.QuoteField(fmt.Sprintf("fk_%s_%s_%d", table.TableName, reference.TableName, ordinal)),
		db.Dialect.QuoteField(key),
		db.Dialect.QuoteField(reference.TableName),
		db.Dialect.QuoteField(column),
	)); err != nil {
		return err
	}

	return nil
}

func AddTable(db *gorp.DbMap, i IdentifiableTable) *gorp.TableMap {
	return db.AddTableWithName(i, i.TableName()).SetKeys(true, "id")
}
