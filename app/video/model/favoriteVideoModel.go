package model

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FavoriteVideoModel = (*customFavoriteVideoModel)(nil)

type (
	// FavoriteVideoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFavoriteVideoModel.
	FavoriteVideoModel interface {
		favoriteVideoModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*FavoriteVideo, error)
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*FavoriteVideo, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*FavoriteVideo, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*FavoriteVideo, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*FavoriteVideo, error)
	}

	customFavoriteVideoModel struct {
		*defaultFavoriteVideoModel
	}
)

// NewFavoriteVideoModel returns a model for the database table.
func NewFavoriteVideoModel(conn sqlx.SqlConn, c cache.CacheConf) FavoriteVideoModel {
	return &customFavoriteVideoModel{
		defaultFavoriteVideoModel: newFavoriteVideoModel(conn, c),
	}
}

func (m *defaultFavoriteVideoModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*FavoriteVideo, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp FavoriteVideo
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

func (m *defaultFavoriteVideoModel) FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error) {

	query, values, err := sumBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultFavoriteVideoModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

	query, values, err := countBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultFavoriteVideoModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*FavoriteVideo, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*FavoriteVideo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultFavoriteVideoModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*FavoriteVideo, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := rowBuilder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*FavoriteVideo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultFavoriteVideoModel) FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*FavoriteVideo, error) {

	if preMinId > 0 {
		rowBuilder = rowBuilder.Where(" id < ? ", preMinId)
	}

	query, values, err := rowBuilder.OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*FavoriteVideo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// 按照id升序分页查询数据，不支持排序
func (m *defaultFavoriteVideoModel) FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*FavoriteVideo, error) {

	if preMaxId > 0 {
		rowBuilder = rowBuilder.Where(" id > ? ", preMaxId)
	}

	query, values, err := rowBuilder.OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*FavoriteVideo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// export logic
func (m *defaultFavoriteVideoModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

// export logic
func (m *defaultFavoriteVideoModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(favoriteVideoRows).From(m.table)
}

// export logic
func (m *defaultFavoriteVideoModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

// export logic
func (m *defaultFavoriteVideoModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}
