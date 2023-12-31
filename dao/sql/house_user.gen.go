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

func newHouseUser(db *gorm.DB, opts ...gen.DOOption) houseUser {
	_houseUser := houseUser{}

	_houseUser.houseUserDo.UseDB(db, opts...)
	_houseUser.houseUserDo.UseModel(&sqlmodel.HouseUser{})

	tableName := _houseUser.houseUserDo.TableName()
	_houseUser.ALL = field.NewAsterisk(tableName)
	_houseUser.ID = field.NewInt32(tableName, "id")
	_houseUser.UniID = field.NewInt64(tableName, "uni_id")
	_houseUser.Nickname = field.NewString(tableName, "nickname")
	_houseUser.HeadImage = field.NewString(tableName, "head_image")
	_houseUser.RegisterTime = field.NewInt32(tableName, "register_time")

	_houseUser.fillFieldMap()

	return _houseUser
}

type houseUser struct {
	houseUserDo houseUserDo

	ALL          field.Asterisk
	ID           field.Int32  // 自增ID
	UniID        field.Int64  // 唯一ID
	Nickname     field.String // 昵称
	HeadImage    field.String // 头像
	RegisterTime field.Int32  // 注册时间

	fieldMap map[string]field.Expr
}

func (h houseUser) Table(newTableName string) *houseUser {
	h.houseUserDo.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h houseUser) As(alias string) *houseUser {
	h.houseUserDo.DO = *(h.houseUserDo.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *houseUser) updateTableName(table string) *houseUser {
	h.ALL = field.NewAsterisk(table)
	h.ID = field.NewInt32(table, "id")
	h.UniID = field.NewInt64(table, "uni_id")
	h.Nickname = field.NewString(table, "nickname")
	h.HeadImage = field.NewString(table, "head_image")
	h.RegisterTime = field.NewInt32(table, "register_time")

	h.fillFieldMap()

	return h
}

func (h *houseUser) WithContext(ctx context.Context) *houseUserDo {
	return h.houseUserDo.WithContext(ctx)
}

func (h houseUser) TableName() string { return h.houseUserDo.TableName() }

func (h houseUser) Alias() string { return h.houseUserDo.Alias() }

func (h houseUser) Columns(cols ...field.Expr) gen.Columns { return h.houseUserDo.Columns(cols...) }

func (h *houseUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *houseUser) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 5)
	h.fieldMap["id"] = h.ID
	h.fieldMap["uni_id"] = h.UniID
	h.fieldMap["nickname"] = h.Nickname
	h.fieldMap["head_image"] = h.HeadImage
	h.fieldMap["register_time"] = h.RegisterTime
}

func (h houseUser) clone(db *gorm.DB) houseUser {
	h.houseUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h houseUser) replaceDB(db *gorm.DB) houseUser {
	h.houseUserDo.ReplaceDB(db)
	return h
}

type houseUserDo struct{ gen.DO }

func (h houseUserDo) Debug() *houseUserDo {
	return h.withDO(h.DO.Debug())
}

func (h houseUserDo) WithContext(ctx context.Context) *houseUserDo {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h houseUserDo) ReadDB() *houseUserDo {
	return h.Clauses(dbresolver.Read)
}

func (h houseUserDo) WriteDB() *houseUserDo {
	return h.Clauses(dbresolver.Write)
}

func (h houseUserDo) Session(config *gorm.Session) *houseUserDo {
	return h.withDO(h.DO.Session(config))
}

func (h houseUserDo) Clauses(conds ...clause.Expression) *houseUserDo {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h houseUserDo) Returning(value interface{}, columns ...string) *houseUserDo {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h houseUserDo) Not(conds ...gen.Condition) *houseUserDo {
	return h.withDO(h.DO.Not(conds...))
}

func (h houseUserDo) Or(conds ...gen.Condition) *houseUserDo {
	return h.withDO(h.DO.Or(conds...))
}

func (h houseUserDo) Select(conds ...field.Expr) *houseUserDo {
	return h.withDO(h.DO.Select(conds...))
}

func (h houseUserDo) Where(conds ...gen.Condition) *houseUserDo {
	return h.withDO(h.DO.Where(conds...))
}

func (h houseUserDo) Order(conds ...field.Expr) *houseUserDo {
	return h.withDO(h.DO.Order(conds...))
}

func (h houseUserDo) Distinct(cols ...field.Expr) *houseUserDo {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h houseUserDo) Omit(cols ...field.Expr) *houseUserDo {
	return h.withDO(h.DO.Omit(cols...))
}

func (h houseUserDo) Join(table schema.Tabler, on ...field.Expr) *houseUserDo {
	return h.withDO(h.DO.Join(table, on...))
}

func (h houseUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) *houseUserDo {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h houseUserDo) RightJoin(table schema.Tabler, on ...field.Expr) *houseUserDo {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h houseUserDo) Group(cols ...field.Expr) *houseUserDo {
	return h.withDO(h.DO.Group(cols...))
}

func (h houseUserDo) Having(conds ...gen.Condition) *houseUserDo {
	return h.withDO(h.DO.Having(conds...))
}

func (h houseUserDo) Limit(limit int) *houseUserDo {
	return h.withDO(h.DO.Limit(limit))
}

func (h houseUserDo) Offset(offset int) *houseUserDo {
	return h.withDO(h.DO.Offset(offset))
}

func (h houseUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *houseUserDo {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h houseUserDo) Unscoped() *houseUserDo {
	return h.withDO(h.DO.Unscoped())
}

func (h houseUserDo) Create(values ...*sqlmodel.HouseUser) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h houseUserDo) CreateInBatches(values []*sqlmodel.HouseUser, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h houseUserDo) Save(values ...*sqlmodel.HouseUser) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h houseUserDo) First() (*sqlmodel.HouseUser, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*sqlmodel.HouseUser), nil
	}
}

func (h houseUserDo) Take() (*sqlmodel.HouseUser, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*sqlmodel.HouseUser), nil
	}
}

func (h houseUserDo) Last() (*sqlmodel.HouseUser, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*sqlmodel.HouseUser), nil
	}
}

func (h houseUserDo) Find() ([]*sqlmodel.HouseUser, error) {
	result, err := h.DO.Find()
	return result.([]*sqlmodel.HouseUser), err
}

func (h houseUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*sqlmodel.HouseUser, err error) {
	buf := make([]*sqlmodel.HouseUser, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h houseUserDo) FindInBatches(result *[]*sqlmodel.HouseUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h houseUserDo) Attrs(attrs ...field.AssignExpr) *houseUserDo {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h houseUserDo) Assign(attrs ...field.AssignExpr) *houseUserDo {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h houseUserDo) Joins(fields ...field.RelationField) *houseUserDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h houseUserDo) Preload(fields ...field.RelationField) *houseUserDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h houseUserDo) FirstOrInit() (*sqlmodel.HouseUser, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*sqlmodel.HouseUser), nil
	}
}

func (h houseUserDo) FirstOrCreate() (*sqlmodel.HouseUser, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*sqlmodel.HouseUser), nil
	}
}

func (h houseUserDo) FindByPage(offset int, limit int) (result []*sqlmodel.HouseUser, count int64, err error) {
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

func (h houseUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h houseUserDo) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h houseUserDo) Delete(models ...*sqlmodel.HouseUser) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *houseUserDo) withDO(do gen.Dao) *houseUserDo {
	h.DO = *do.(*gen.DO)
	return h
}
