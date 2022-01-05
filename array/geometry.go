package array

// MaxPoints 从point中找出在一条线上的最多点数（一条线上最多可以串多少个点）
// no 149
func MaxPoints(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}

	tanMap := map[float32]int{}
	max := 2
	for i := 1; i < len(points); i++ {
		for j := 0; j < i; j++ {
			a := float32(points[i][0] - points[j][0])
			b := float32(points[i][1] - points[j][1])
			tan := a / b
			if _, ok := tanMap[tan]; !ok {
				tanMap[tan] = 1
			} else {
				tanMap[tan]++
				if tanMap[tan] > max {
					max = tanMap[tan]
				}
			}
		}
	}

	return max
}
