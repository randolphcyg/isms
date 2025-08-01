// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                    = new(Query)
	IsmsCountry          *ismsCountry
	IsmsDeveloper        *ismsDeveloper
	IsmsIndustry         *ismsIndustry
	IsmsOS               *ismsOS
	IsmsSoftware         *ismsSoftware
	IsmsSoftwareIndustry *ismsSoftwareIndustry
	IsmsSoftwareOS       *ismsSoftwareOS
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	IsmsCountry = &Q.IsmsCountry
	IsmsDeveloper = &Q.IsmsDeveloper
	IsmsIndustry = &Q.IsmsIndustry
	IsmsOS = &Q.IsmsOS
	IsmsSoftware = &Q.IsmsSoftware
	IsmsSoftwareIndustry = &Q.IsmsSoftwareIndustry
	IsmsSoftwareOS = &Q.IsmsSoftwareOS
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                   db,
		IsmsCountry:          newIsmsCountry(db, opts...),
		IsmsDeveloper:        newIsmsDeveloper(db, opts...),
		IsmsIndustry:         newIsmsIndustry(db, opts...),
		IsmsOS:               newIsmsOS(db, opts...),
		IsmsSoftware:         newIsmsSoftware(db, opts...),
		IsmsSoftwareIndustry: newIsmsSoftwareIndustry(db, opts...),
		IsmsSoftwareOS:       newIsmsSoftwareOS(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	IsmsCountry          ismsCountry
	IsmsDeveloper        ismsDeveloper
	IsmsIndustry         ismsIndustry
	IsmsOS               ismsOS
	IsmsSoftware         ismsSoftware
	IsmsSoftwareIndustry ismsSoftwareIndustry
	IsmsSoftwareOS       ismsSoftwareOS
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                   db,
		IsmsCountry:          q.IsmsCountry.clone(db),
		IsmsDeveloper:        q.IsmsDeveloper.clone(db),
		IsmsIndustry:         q.IsmsIndustry.clone(db),
		IsmsOS:               q.IsmsOS.clone(db),
		IsmsSoftware:         q.IsmsSoftware.clone(db),
		IsmsSoftwareIndustry: q.IsmsSoftwareIndustry.clone(db),
		IsmsSoftwareOS:       q.IsmsSoftwareOS.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                   db,
		IsmsCountry:          q.IsmsCountry.replaceDB(db),
		IsmsDeveloper:        q.IsmsDeveloper.replaceDB(db),
		IsmsIndustry:         q.IsmsIndustry.replaceDB(db),
		IsmsOS:               q.IsmsOS.replaceDB(db),
		IsmsSoftware:         q.IsmsSoftware.replaceDB(db),
		IsmsSoftwareIndustry: q.IsmsSoftwareIndustry.replaceDB(db),
		IsmsSoftwareOS:       q.IsmsSoftwareOS.replaceDB(db),
	}
}

type queryCtx struct {
	IsmsCountry          IIsmsCountryDo
	IsmsDeveloper        IIsmsDeveloperDo
	IsmsIndustry         IIsmsIndustryDo
	IsmsOS               IIsmsOSDo
	IsmsSoftware         IIsmsSoftwareDo
	IsmsSoftwareIndustry IIsmsSoftwareIndustryDo
	IsmsSoftwareOS       IIsmsSoftwareOSDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		IsmsCountry:          q.IsmsCountry.WithContext(ctx),
		IsmsDeveloper:        q.IsmsDeveloper.WithContext(ctx),
		IsmsIndustry:         q.IsmsIndustry.WithContext(ctx),
		IsmsOS:               q.IsmsOS.WithContext(ctx),
		IsmsSoftware:         q.IsmsSoftware.WithContext(ctx),
		IsmsSoftwareIndustry: q.IsmsSoftwareIndustry.WithContext(ctx),
		IsmsSoftwareOS:       q.IsmsSoftwareOS.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
