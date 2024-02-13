/*
 * @lc app=leetcode id=904 lang=golang
 *
 * [904] Fruit Into Baskets
 */

// @lc code=start
func totalFruit(fruits []int) int {
    left, right := 0,0
	total,count := 0,0
	types := make(map[int]int)
	for right < len(fruits) && left< len(fruits){
		types[fruits[right]]++
		// fmt.Println(fruits[right] ,":" ,types[fruits[right]])
		for len(types) > 2{
			// fmt.Println("right ",right,"left",left)

	
			// fmt.Println("total: ",total,"left: ",left)

			//update types
			types[fruits[left]] --
			// fmt.Println("fruits: ",fruits[left],"count",types[fruits[left]])

			if  types[fruits[left]]<=0{
				delete(types,fruits[left])
			}
			// keep moving starting position until only 2 types of fruit left
			left++
		}
		count = right - left + 1
			if total < count {
				total = count
			}
		right ++
	}

	return total
}
// @lc code=end

