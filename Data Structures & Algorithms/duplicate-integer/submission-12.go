func hasDuplicate(nums []int) bool {
    for i ,n := range nums {
        for j,m := range nums{
            if (m == n && j != i){
                return true
            }
        }
    }
    return false
}
