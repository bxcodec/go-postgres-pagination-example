package offsetlimit

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	pagination "github.com/bxcodec/go-postgres-pagination-example"
)

type FetchParam struct {
	Limit  uint64
	OffSet uint64
}

func FetchPayment(ctx context.Context, db *sql.DB, params FetchParam) (res []pagination.Payment, nextOffset int64, err error) {
	queryBuilder := sq.Select("id", "amount", "name", "created_time").From("payment").PlaceholderFormat(sq.Dollar).OrderBy("id DESC")

	if params.Limit > 0 {
		queryBuilder = queryBuilder.Limit(params.Limit)
	}

	if params.OffSet > 0 {
		queryBuilder = queryBuilder.Offset(params.OffSet)
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
		nextOffset = int64(params.OffSet + params.Limit)
	}

	return
}
