package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// ReviewRepo
type ReviewRepo interface {
	Save(context.Context, *Greeter) (*Greeter, error)
}

// ReviewUsecase is a review usecase.
type ReviewUsecase struct {
	repo ReviewRepo
	log  *log.Helper
}

// NewReviewUsecase new a Review usecase.
func NewReviewUsecase(repo ReviewRepo, logger log.Logger) *ReviewUsecase {
	return &ReviewUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *ReviewUsecase) Save(ctx context.Context, g *Greeter) (*Greeter, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
