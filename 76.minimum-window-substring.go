/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */

// @lc code=start
func minWindow(s string, t string) string {
	if len(t)>len(s){
		return ""
	}
	count:=0
	tracker:=make(map[byte]int)
	for i := range t{
		tracker[t[i]] ++
		count++
	}
	fmt.Println("tracker: ", tracker)

	left,right:=0,0
	var res = string(make([]byte,len(s)))

	for right<len(s){
		if v,ok := tracker[s[right]]; ok{
			tracker[s[right]]--
			if v>0{ // 
				count--

			}
		}

		for count<=0{

			if len(res) >= len(s[left:right+1]){
				res = s[left:right+1]
			}
			if _,ok := tracker[s[left]]; ok{
				tracker[s[left]]++
				if tracker[s[left]] >0 {
					count++

				}
			}
			left++
		}

		right++
	}

	if res == string(make([]byte, len(s))) {
        return ""
    }
	/*
	left,right := 0,0
	for right<len(s) && left<=right{
		for count>0{
			if v,ok := tracker[s[right]]; ok{

			}
		}
	}
	*/
	return res
}

// @lc code=end

