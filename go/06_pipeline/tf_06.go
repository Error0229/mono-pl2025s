package pipeline

func CurriedAdd(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x*y + z
		}
	}
}
