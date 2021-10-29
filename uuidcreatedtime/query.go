package uuidcreatedtime

import (
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	pagination "github.com/bxcodec/go-postgres-pagination-example"
)

type FetchParam struct {
	Limit  uint64
	Cursor string
}

func decodeCursor(encodedCursor string) (res time.Time, uuid string, err error) {
	byt, err := base64.StdEncoding.DecodeString(encodedCursor)
	if err != nil {
		return
	}

	arrStr := strings.Split(string(byt), ",")
	if len(arrStr) != 2 {
		err = errors.New("cursor is invalid")
		return
	}

	res, err = time.Parse(time.RFC3339Nano, arrStr[0])
	if err != nil {
		return
	}
	uuid = arrStr[1]
	return
}

func encodeCursor(t time.Time, uuid string) string {
	key := fmt.Sprintf("%s,%s", t.Format(time.RFC3339Nano), uuid)
	return base64.StdEncoding.EncodeToString([]byte(key))
}

func FetchPayment(ctx context.Context, db *sql.DB, params FetchParam) (res []pagination.Payment, nextCursor string, err error) {
	queryBuilder := sq.Select("id", "amount", "name", "created_time").From("payment_with_uuid").PlaceholderFormat(sq.Dollar).OrderBy("created_time DESC")

	if params.Limit > 0 {
		queryBuilder = queryBuilder.Limit(params.Limit)
	}

	if params.Cursor != "" {
		createdCursor, paymentID, errCsr := decodeCursor(params.Cursor)
		if errCsr != nil {
			err = errors.New("invalid-cursor")
			return
		}
		queryBuilder = queryBuilder.Where(sq.LtOrEq{
			"created_time": createdCursor,
		})
		queryBuilder = queryBuilder.Where(sq.NotEq{
			"id": paymentID,
		})
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
	var createdTime time.Time // only using one for all loops, we only need the latest one in the end
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
		createdTime = item.CreatedTime
		res = append(res, item)
	}

	if len(res) > 0 {
		nextCursor = encodeCursor(createdTime, res[len(res)-1].ID)
	}

	return
}
