package main

import "fmt"

func main() {

	sort()
	keyValuePairs()
}

func sort() {
	nums := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	end := len(nums) - 1
	for {

		if end == 0 {
			break
		}

		for i := 0; i < len(nums)-1; i++ {

			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}

		}
		fmt.Println(nums)
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")

		end -= 1

	}
}

func keyValuePairs() {
	pairs := map[string]string{"username": "Reverend Billy Horton Heat", "password": "Definitely Encrypted", "email": "yahoo@gmail.com"}

	for key, value := range pairs {
		fmt.Println("Key: " + key)
		fmt.Println("Value: " + value)
	}
}
