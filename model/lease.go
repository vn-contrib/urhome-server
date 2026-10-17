package model

import (
	"time"

	"github.com/uptrace/bun"
	"github.com/vn-contrib/urhome-server/pkg/bar"
)

var ModelLease = bar.ModelFor[Lease]()

type Lease struct {
	bun.BaseModel `bun:"alias:l"`

	ID         int32 `bun:",pk,autoincrement"`
	PropertyID int32
	StartTime  time.Time
	EndTime    *time.Time
}
