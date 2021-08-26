package main

import "fmt"

func twoSum(nums []int, target int) []int {
    tmp := make(map[int]int)
    for i :=0; i <len(nums); i ++ {
        another := target - nums[i]
        j, ok := tmp[another]
        if ok {
            return []int{j, i}
        }
        tmp[nums[i]] = i
    }
    return nil
}

func main() {
    l := []int{1, 3, 5, 7}
    target := 8
    result := twoSum(l, target)
    fmt.Println(result)
}