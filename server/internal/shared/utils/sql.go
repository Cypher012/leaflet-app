package utils

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func PgText(s string) pgtype.Text {
	return pgtype.Text{
		String: s,
		Valid:  s != "",
	}
}

func PgInt8(i int64) pgtype.Int8 {
	return pgtype.Int8{
		Int64: i,
		Valid: true,
	}
}

func PgFloat8(f float64) pgtype.Float8 {
	return pgtype.Float8{
		Float64: f,
		Valid:   true,
	}
}

func PgTimestamptz(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{
		Time:  t,
		Valid: !t.IsZero(),
	}
}

func PgCursorTime(t *time.Time) pgtype.Timestamptz {
	if t == nil {
		return pgtype.Timestamptz{Valid: false}
	}
	return PgTimestamptz(*t)
}
