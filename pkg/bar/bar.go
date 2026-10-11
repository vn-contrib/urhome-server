package bar

import (
	"context"
	"database/sql"
)

func Create(ctx context.Context, model any) error {
	_, err := DBFromContext(ctx).NewInsert().Model(model).Exec(ctx)
	return err
}

func FindBy(ctx context.Context, model any, cols ...string) error {
	return DBFromContext(ctx).NewSelect().Model(model).WherePK(cols...).Scan(ctx)
}

func FindOrCreateBy(ctx context.Context, model any, cols ...string) error {
	err := FindBy(ctx, model, cols...)
	if err == sql.ErrNoRows {
		return Create(ctx, model)
	}
	return err
}
