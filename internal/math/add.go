package math

import "golang.org/x/exp/constraints"

type AddAble interface {
	constraints.Integer | constraints.Float
}

func Add[T AddAble](a, b T) T {
	return a + b
}
