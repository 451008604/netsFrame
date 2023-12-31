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

func newHouseAccount(db *gorm.DB, opts ...gen.DOOption) houseAccount {
	_houseAccount := houseAccount{}

	_houseAccount.houseAccountDo.UseDB(db, opts...)
	_houseAccount.houseAccountDo.UseModel(&sqlmodel.HouseAccount{})

	tableName := _houseAccount.houseAccountDo.TableName()
	_houseAccount.ALL = field.NewAsterisk(tableName)
	_houseAccount.ID = field.NewInt32(tableName, "id")
	_houseAccount.UserID = field.NewInt32(tableName, "user_id")
	_houseAccount.Account = field.NewString(tableName, "account")
	_houseAccount.Password = field.NewString(tableName, "password")

	_houseAccount.fillFieldMap()

	return _houseAccount
}

type houseAccount struct {
	houseAccountDo houseAccountDo

	ALL      field.Asterisk
	ID       field.Int32  // 唯一ID
	UserID   field.Int32  // 用户ID
	Account  field.String // 账号
	Password field.String // MD5加密后的密码

	fieldMap map[string]field.Expr
}

func (h houseAccount) Table(newTableName string) *houseAccount {
	h.houseAccountDo.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h houseAccount) As(alias string) *houseAccount {
	h.houseAccountDo.DO = *(h.houseAccountDo.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *houseAccount) updateTableName(table string) *houseAccount {
	h.ALL = field.NewAsterisk(table)
	h.ID = field.NewInt32(table, "id")
	h.UserID = field.NewInt32(table, "user_id")
	h.Account = field.NewString(table, "account")
	h.Password = field.NewString(table, "password")

	h.fillFieldMap()

	return h
}

func (h *houseAccount) WithContext(ctx context.Context) *houseAccountDo {
	return h.houseAccountDo.WithContext(ctx)
}

func (h houseAccount) TableName() string { return h.houseAccountDo.TableName() }

func (h houseAccount) Alias() string { return h.houseAccountDo.Alias() }

func (h houseAccount) Columns(cols ...field.Expr) gen.Columns {
	return h.houseAccountDo.Columns(cols...)
}

func (h *houseAccount) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *houseAccount) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 4)
	h.fieldMap["id"] = h.ID
	h.fieldMap["user_id"] = h.UserID
	h.fieldMap["account"] = h.Account
	h.fieldMap["password"] = h.Password
}

func (h houseAccount) clone(db *gorm.DB) houseAccount {
	h.houseAccountDo.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h houseAccount) replaceDB(db *gorm.DB) houseAccount {
	h.houseAccountDo.ReplaceDB(db)
	return h
}

type houseAccountDo struct{ gen.DO }

func (h houseAccountDo) Debug() *houseAccountDo {
	return h.withDO(h.DO.Debug())
}

func (h houseAccountDo) WithContext(ctx context.Context) *houseAccountDo {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h houseAccountDo) ReadDB() *houseAccountDo {
	return h.Clauses(dbresolver.Read)
}

func (h houseAccountDo) WriteDB() *houseAccountDo {
	return h.Clauses(dbresolver.Write)
}

func (h houseAccountDo) Session(config *gorm.Session) *houseAccountDo {
	return h.withDO(h.DO.Session(config))
}

func (h houseAccountDo) Clauses(conds ...clause.Expression) *houseAccountDo {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h houseAccountDo) Returning(value interface{}, columns ...string) *houseAccountDo {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h houseAccountDo) Not(conds ...gen.Condition) *houseAccountDo {
	return h.withDO(h.DO.Not(conds...))
}

func (h houseAccountDo) Or(conds ...gen.Condition) *houseAccountDo {
	return h.withDO(h.DO.Or(conds...))
}

func (h houseAccountDo) Select(conds ...field.Expr) *houseAccountDo {
	return h.withDO(h.DO.Select(conds...))
}

func (h houseAccountDo) Where(conds ...gen.Condition) *houseAccountDo {
	return h.withDO(h.DO.Where(conds...))
}

func (h houseAccountDo) Order(conds ...field.Expr) *houseAccountDo {
	return h.withDO(h.DO.Order(conds...))
}

func (h houseAccountDo) Distinct(cols ...field.Expr) *houseAccountDo {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h houseAccountDo) Omit(cols ...field.Expr) *houseAccountDo {
	return h.withDO(h.DO.Omit(cols...))
}

func (h houseAccountDo) Join(table schema.Tabler, on ...field.Expr) *houseAccountDo {
	return h.withDO(h.DO.Join(table, on...))
}

func (h houseAccountDo) LeftJoin(table schema.Tabler, on ...field.Expr) *houseAccountDo {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h houseAccountDo) RightJoin(table schema.Tabler, on ...field.Expr) *houseAccountDo {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h houseAccountDo) Group(cols ...field.Expr) *houseAccountDo {
	return h.withDO(h.DO.Group(cols...))
}

func (h houseAccountDo) Having(conds ...gen.Condition) *houseAccountDo {
	return h.withDO(h.DO.Having(conds...))
}

func (h houseAccountDo) Limit(limit int) *houseAccountDo {
	return h.withDO(h.DO.Limit(limit))
}

func (h houseAccountDo) Offset(offset int) *houseAccountDo {
	return h.withDO(h.DO.Offset(offset))
}

func (h houseAccountDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *houseAccountDo {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h houseAccountDo) Unscoped() *houseAccountDo {
	return h.withDO(h.DO.Unscoped())
}

func (h houseAccountDo) Create(values ...*sqlmodel.HouseAccount) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h houseAccountDo) CreateInBatches(values []*sqlmodel.HouseAccount, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h houseAccountDo) Save(values ...*sqlmodel.HouseAccount) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h houseAccountDo) First() (*sqlmodel.HouseAccount, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*sqlmodel.HouseAccount), nil
	}
}

func (h houseAccountDo) Take() (*sqlmodel.HouseAccount, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*sqlmodel.HouseAccount), nil
	}
}

func (h houseAccountDo) Last() (*sqlmodel.HouseAccount, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*sqlmodel.HouseAccount), nil
	}
}

func (h houseAccountDo) Find() ([]*sqlmodel.HouseAccount, error) {
	result, err := h.DO.Find()
	return result.([]*sqlmodel.HouseAccount), err
}

func (h houseAccountDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*sqlmodel.HouseAccount, err error) {
	buf := make([]*sqlmodel.HouseAccount, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h houseAccountDo) FindInBatches(result *[]*sqlmodel.HouseAccount, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h houseAccountDo) Attrs(attrs ...field.AssignExpr) *houseAccountDo {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h houseAccountDo) Assign(attrs ...field.AssignExpr) *houseAccountDo {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h houseAccountDo) Joins(fields ...field.RelationField) *houseAccountDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h houseAccountDo) Preload(fields ...field.RelationField) *houseAccountDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h houseAccountDo) FirstOrInit() (*sqlmodel.HouseAccount, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*sqlmodel.HouseAccount), nil
	}
}

func (h houseAccountDo) FirstOrCreate() (*sqlmodel.HouseAccount, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*sqlmodel.HouseAccount), nil
	}
}

func (h houseAccountDo) FindByPage(offset int, limit int) (result []*sqlmodel.HouseAccount, count int64, err error) {
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

func (h houseAccountDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h houseAccountDo) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h houseAccountDo) Delete(models ...*sqlmodel.HouseAccount) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *houseAccountDo) withDO(do gen.Dao) *houseAccountDo {
	h.DO = *do.(*gen.DO)
	return h
}
