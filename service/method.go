package service

import (
	"context"

	"github.com/heartBrokenGod/beer/model"
	"github.com/heartBrokenGod/beer/pkg/restutils/errors"
)

func (s *ServiceImpl) Method(ctx context.Context, reqModel *model.Model) (respModel *model.Model, err errors.Error) {
	// debug log for more context
	s.logger.Debug().Msg("method executing...")
	defer s.logger.Debug().Msg("method executed")

	// write the business logic

	return &model.Model{}, nil

}
