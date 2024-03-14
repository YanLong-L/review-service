package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"review-service/internal/biz"
	"review-service/internal/data/model"
)

type reviewRepo struct {
	data *Data
	log  *log.Helper
}

// NewReviewRepo
func NewReviewRepo(data *Data, logger log.Logger) biz.ReviewRepo {
	return &reviewRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r reviewRepo) SaveReview(ctx context.Context, review *model.ReviewInfo) (*model.ReviewInfo, error) {
	//TODO implement me
	panic("implement me")
}
