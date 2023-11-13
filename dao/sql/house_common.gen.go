// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package sql

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"netsFrame/dao/sqlmodel"
)

func newHouseCommon(db *gorm.DB, opts ...gen.DOOption) houseCommon {
	_houseCommon := houseCommon{}

	_houseCommon.houseCommonDo.UseDB(db, opts...)
	_houseCommon.houseCommonDo.UseModel(&sqlmodel.HouseCommon{})

	tableName := _houseCommon.houseCommonDo.TableName()
	_houseCommon.ALL = field.NewAsterisk(tableName)
	_houseCommon.ID = field.NewInt32(tableName, "id")
	_houseCommon.UserID = field.NewInt32(tableName, "user_id")
	_houseCommon.Level = field.NewInt32(tableName, "level")
	_houseCommon.GoldCoin = field.NewInt32(tableName, "gold_coin")
	_houseCommon.Diamond = field.NewInt32(tableName, "diamond")
	_houseCommon.Strength = field.NewInt32(tableName, "strength")
	_houseCommon.Experience = field.NewInt32(tableName, "experience")

	_houseCommon.fillFieldMap()

	return _houseCommon
}

type houseCommon struct {
	houseCommonDo houseCommonDo

	ALL        field.Asterisk
	ID         field.Int32 // 唯一ID
	UserID     field.Int32 // 用户ID
	Level      field.Int32 // 等级
	GoldCoin   field.Int32 // 金币
	Diamond    field.Int32 // 钻石
	Strength   field.Int32 // 体力
	Experience field.Int32 // 经验

	fieldMap map[string]field.Expr
}

func (h houseCommon) Table(newTableName string) *houseCommon {
	h.houseCommonDo.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h houseCommon) As(alias string) *houseCommon {
	h.houseCommonDo.DO = *(h.houseCommonDo.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *houseCommon) updateTableName(table string) *houseCommon {
	h.ALL = field.NewAsterisk(table)
	h.ID = field.NewInt32(table, "id")
	h.UserID = field.NewInt32(table, "user_id")
	h.Level = field.NewInt32(table, "level")
	h.GoldCoin = field.NewInt32(table, "gold_coin")
	h.Diamond = field.NewInt32(table, "diamond")
	h.Strength = field.NewInt32(table, "strength")
	h.Experience = field.NewInt32(table, "experience")

	h.fillFieldMap()

	return h
}

func (h *houseCommon) WithContext(ctx context.Context) *houseCommonDo {
	return h.houseCommonDo.WithContext(ctx)
}

func (h houseCommon) TableName() string { return h.houseCommonDo.TableName() }

func (h houseCommon) Alias() string { return h.houseCommonDo.Alias() }

func (h houseCommon) Columns(cols ...field.Expr) gen.Columns { return h.houseCommonDo.Columns(cols...) }

func (h *houseCommon) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *houseCommon) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 7)
	h.fieldMap["id"] = h.ID
	h.fieldMap["user_id"] = h.UserID
	h.fieldMap["level"] = h.Level
	h.fieldMap["gold_coin"] = h.GoldCoin
	h.fieldMap["diamond"] = h.Diamond
	h.fieldMap["strength"] = h.Strength
	h.fieldMap["experience"] = h.Experience
}

func (h houseCommon) clone(db *gorm.DB) houseCommon {
	h.houseCommonDo.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h houseCommon) replaceDB(db *gorm.DB) houseCommon {
	h.houseCommonDo.ReplaceDB(db)
	return h
}

type houseCommonDo struct{ gen.DO }

func (h houseCommonDo) Debug() *houseCommonDo {
	return h.withDO(h.DO.Debug())
}

func (h houseCommonDo) WithContext(ctx context.Context) *houseCommonDo {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h houseCommonDo) ReadDB() *houseCommonDo {
	return h.Clauses(dbresolver.Read)
}

func (h houseCommonDo) WriteDB() *houseCommonDo {
	return h.Clauses(dbresolver.Write)
}

func (h houseCommonDo) Session(config *gorm.Session) *houseCommonDo {
	return h.withDO(h.DO.Session(config))
}

func (h houseCommonDo) Clauses(conds ...clause.Expression) *houseCommonDo {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h houseCommonDo) Returning(value interface{}, columns ...string) *houseCommonDo {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h houseCommonDo) Not(conds ...gen.Condition) *houseCommonDo {
	return h.withDO(h.DO.Not(conds...))
}

func (h houseCommonDo) Or(conds ...gen.Condition) *houseCommonDo {
	return h.withDO(h.DO.Or(conds...))
}

func (h houseCommonDo) Select(conds ...field.Expr) *houseCommonDo {
	return h.withDO(h.DO.Select(conds...))
}

func (h houseCommonDo) Where(conds ...gen.Condition) *houseCommonDo {
	return h.withDO(h.DO.Where(conds...))
}

func (h houseCommonDo) Order(conds ...field.Expr) *houseCommonDo {
	return h.withDO(h.DO.Order(conds...))
}

func (h houseCommonDo) Distinct(cols ...field.Expr) *houseCommonDo {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h houseCommonDo) Omit(cols ...field.Expr) *houseCommonDo {
	return h.withDO(h.DO.Omit(cols...))
}

func (h houseCommonDo) Join(table schema.Tabler, on ...field.Expr) *houseCommonDo {
	return h.withDO(h.DO.Join(table, on...))
}

func (h houseCommonDo) LeftJoin(table schema.Tabler, on ...field.Expr) *houseCommonDo {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h houseCommonDo) RightJoin(table schema.Tabler, on ...field.Expr) *houseCommonDo {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h houseCommonDo) Group(cols ...field.Expr) *houseCommonDo {
	return h.withDO(h.DO.Group(cols...))
}

func (h houseCommonDo) Having(conds ...gen.Condition) *houseCommonDo {
	return h.withDO(h.DO.Having(conds...))
}

func (h houseCommonDo) Limit(limit int) *houseCommonDo {
	return h.withDO(h.DO.Limit(limit))
}

func (h houseCommonDo) Offset(offset int) *houseCommonDo {
	return h.withDO(h.DO.Offset(offset))
}

func (h houseCommonDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *houseCommonDo {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h houseCommonDo) Unscoped() *houseCommonDo {
	return h.withDO(h.DO.Unscoped())
}

func (h houseCommonDo) Create(values ...*sqlmodel.HouseCommon) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h houseCommonDo) CreateInBatches(values []*sqlmodel.HouseCommon, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h houseCommonDo) Save(values ...*sqlmodel.HouseCommon) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h houseCommonDo) First() (*sqlmodel.HouseCommon, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*sqlmodel.HouseCommon), nil
	}
}

func (h houseCommonDo) Take() (*sqlmodel.HouseCommon, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*sqlmodel.HouseCommon), nil
	}
}

func (h houseCommonDo) Last() (*sqlmodel.HouseCommon, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*sqlmodel.HouseCommon), nil
	}
}

func (h houseCommonDo) Find() ([]*sqlmodel.HouseCommon, error) {
	result, err := h.DO.Find()
	return result.([]*sqlmodel.HouseCommon), err
}

func (h houseCommonDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*sqlmodel.HouseCommon, err error) {
	buf := make([]*sqlmodel.HouseCommon, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h houseCommonDo) FindInBatches(result *[]*sqlmodel.HouseCommon, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h houseCommonDo) Attrs(attrs ...field.AssignExpr) *houseCommonDo {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h houseCommonDo) Assign(attrs ...field.AssignExpr) *houseCommonDo {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h houseCommonDo) Joins(fields ...field.RelationField) *houseCommonDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h houseCommonDo) Preload(fields ...field.RelationField) *houseCommonDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h houseCommonDo) FirstOrInit() (*sqlmodel.HouseCommon, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*sqlmodel.HouseCommon), nil
	}
}

func (h houseCommonDo) FirstOrCreate() (*sqlmodel.HouseCommon, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*sqlmodel.HouseCommon), nil
	}
}

func (h houseCommonDo) FindByPage(offset int, limit int) (result []*sqlmodel.HouseCommon, count int64, err error) {
	result, err = h.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = h.Offset(-1).Limit(-1).Count()
	return
}

func (h houseCommonDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h houseCommonDo) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h houseCommonDo) Delete(models ...*sqlmodel.HouseCommon) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *houseCommonDo) withDO(do gen.Dao) *houseCommonDo {
	h.DO = *do.(*gen.DO)
	return h
}
