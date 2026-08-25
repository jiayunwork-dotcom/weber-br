package drop

import "errors"

var (
	errZeroDenominator = errors.New("Ohnesorge denominator is zero")
	errNegativeTarget  = errors.New("target Weber number must be non-negative")
	errZeroVelocity    = errors.New("velocity must be positive to invert diameter")
)

func IsWeberError(err error) bool {
	return errors.Is(err, errZeroDenominator) ||
		errors.Is(err, errNegativeTarget) ||
		errors.Is(err, errZeroVelocity)
}
