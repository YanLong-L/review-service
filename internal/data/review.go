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

// GetReviewByOrderID 通过订单ID查找评价数据
func (r *reviewRepo) GetReviewByOrderID(ctx context.Context, orderID int64) ([]*model.ReviewInfo, error) {
	res, err := r.data.query.ReviewInfo.WithContext(ctx).
		Where(r.data.query.ReviewInfo.OrderID.Eq(orderID)).
		Find()
	return res, err
}

// SaveReview 创建评价
func (r *reviewRepo) SaveReview(ctx context.Context, review *model.ReviewInfo) (*model.ReviewInfo, error) {
	err := r.data.query.ReviewInfo.WithContext(ctx).
		Save(review)
	return review, err
}
