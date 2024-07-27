package diffsquares

func SquareOfSum(n int) int  { sumSequence := (n * (1 + n)) / 2; return sumSequence * sumSequence }
func SumOfSquares(n int) int { return (n * (n + 1) * (1 + 2*n)) / 6 }
func Difference(n int) int   { return SquareOfSum(n) - SumOfSquares(n) }
