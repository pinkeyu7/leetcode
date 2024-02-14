package main

func getSumAbsoluteDifferences(nums []int) []int {
	l := len(nums)
	var res []int
	sum, prev := 0, 0
	for _, v := range nums {
		sum += v
	}

	for i, v := range nums {
		temp := v - prev

		sum += temp * (2*i - l)
		res = append(res, sum)
		prev = v
	}

	return res
}
