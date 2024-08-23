package e

import "errors"

var (
	ErrNotFound              = errors.New("not found")
	ErrUserNotFound          = errors.New("user not found")
	ErrUserConflict          = errors.New("user already exist")
	ErrIsEmpty               = errors.New("is empty")
	ErrDb                    = errors.New("db error")
	ErrInvalidToken          = errors.New("invalid token")
	ErrInvalidPassword       = errors.New("invalid password")
	ErrInvalidStatusValue    = errors.New("invalid status value")
	ErrNotAvailable          = errors.New("not available")
	ErrQuantityExceeds       = errors.New("selected quantity not available")
	ErrInvalidCoupon         = errors.New("invalid coupon")
	ErrCouponNotApplicable   = errors.New("coupon doesn't meet terms")
	ErrCouponAlreadyRedeemed = errors.New("coupon already redeemed")
)
