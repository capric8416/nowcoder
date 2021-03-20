package main

/**
 * find median in two sorted array
 * @param arr1 int整型一维数组 the array1
 * @param arr2 int整型一维数组 the array2
 * @return int整型
 */
func findMedianinTwoSortedAray(arr1 []int, arr2 []int) int {
	// write code here

	length := len(arr1)
	if length == 1 {
		if arr1[0] < arr2[0] {
			return arr1[0]
		} else {
			return arr2[0]
		}
	}

	var odd bool
	var left1, center1, right1, left2, center2, right2 int
	for left1, right1, left2, right2 = 0, length-1, 0, length-1; left1 < right1; {
		odd = (right1-left1+1)&1 == 1
		center1, center2 = left1+(right1-left1)>>1, left2+(right2-left2)>>1
		if arr1[center1] == arr2[center2] {
			return arr1[center1]
		} else if arr1[center1] > arr2[center2] {
			if odd {
				right1, left2 = center1, center2
			} else {
				right1, left2 = center1, center2+1
			}
		} else {
			if odd {
				left1, right2 = center1, center2
			} else {
				left1, right2 = center1+1, center2
			}
		}
	}

	if arr1[left1] < arr2[left2] {
		return arr1[left1]
	} else {
		return arr2[left2]
	}
}

func main() {

}
