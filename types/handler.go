package types

import (
	"context"
)

type IHandler interface {
	SetContext(c context.Context)
}
