package isEven

import (
	"github.com/GOKaraketir/isOdd"
	"golang.org/x/exp/constraints"
)

func IsEven[T constraints.Integer](val T) bool {
	return !isOdd.IsOdd(val)
}
