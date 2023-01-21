package service

import (
	"context"

	"github.com/heartBrokenGod/beer/model"
	"github.com/heartBrokenGod/beer/pkg/restutils/errors"
)

type Service interface {
	Method(ctx context.Context, reqModel *model.Model) (respModel *model.Model, err errors.Error)
}
