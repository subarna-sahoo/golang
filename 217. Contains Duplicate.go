//217. Contains Duplicate

func containsDuplicate(nums []int) bool {
  temp_nums := make(map[int]int)
  for _, num := range nums {
    if _, isExist := temp_nums[num]; isExist {
      return true
    } else {
      temp_nums[num] = 0
    }
  }
  return false
}
