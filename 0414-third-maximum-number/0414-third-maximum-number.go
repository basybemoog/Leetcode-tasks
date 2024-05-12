import "math"
func thirdMax(nums []int) int {
        first := math.MinInt
    second, third := first, first

    for _, value := range nums {
        if value > first {
            first, second, third = value, first, second
        } else if value > second && value != first {
            second, third = value, second
        } else if value > third && value != second && value != first {
            third = value
        }
    }

    if third == math.MinInt {
        return first
    }
    return third
}