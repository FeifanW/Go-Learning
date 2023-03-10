package main

import "fmt"

func main() {
	nums1 := [3]int{1, 2, 3}
	nums2 := [3]int{7, 8, 9}

	sort := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0
	for {
		if nums1[i] < nums2[j] {
			sort = append(sort, nums1[i])
			i++
		} else {
			sort = append(sort, nums2[j])
			j++
		}
	}
	fmt.Println("数组", sort)
}
