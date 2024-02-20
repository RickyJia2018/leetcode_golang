/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start

func isAnagram(s string, t string) bool {
	temp:= [26]int{}  // cannot use make, because slice can only comapre to nil

	for _,c := range s{
		temp[c - rune('a')]++
	}
	for _,c := range t{
		temp[c - rune('a')]--
	}
	return temp == [26]int{}
}

// func isAnagram(s string, t string) bool {
//     sByte := []byte(s)
// 	tByte := []byte(t)

// 	temp := make([]int,256)

// 	for i:=0; i<len(sByte);i++{
// 		temp[sByte[i]]++
// 	}
// 	for i:=0; i<len(tByte);i++{
// 		temp[tByte[i]]--
// 	}
// 	fmt.Println(temp)
// 	for i:=0; i<len(sByte);i++{
// 		if temp[i] != 0{
// 			return false
// 		}
// 	}
// 	return true
// }
// @lc code=end

