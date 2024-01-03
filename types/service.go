package types

import (
	"context"
)

type IService interface {
	SetContext(c context.Context)
}
