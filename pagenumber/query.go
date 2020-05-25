package pagenumber

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	pagination "github.com/bxcodec/go-postgres-pagination-example"
)

const (
	DefaultLimit = 10
)

type FetchParam struct {
	PageNumber uint64
}

func FetchPayment(ctx context.Context, db *sql.DB, params FetchParam) (res []pagination.Payment, nextPage int64, err error) {
	queryBuilder := sq.Select("id", "amount", "name", "created_time").From("payment").PlaceholderFormat(sq.Dollar).OrderBy("id DESC")
	queryBuilder = queryBuilder.Limit(DefaultLimit)
	if params.PageNumber >= 2 {
		queryBuilder = queryBuilder.Offset(params.PageNumber * DefaultLimit)
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return
	}

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return
	}

	res = []pagination.Payment{}
	for rows.Next() {
		var item pagination.Payment
		err = rows.Scan(
			&item.ID,
			&item.Amount,
			&item.Name,
			&item.CreatedTime,
		)
		if err != nil {
			return
		}
		res = append(res, item)
	}

	if len(res) > 0 {
		nextPage = int64(params.PageNumber + 1)
	}
	return
}
