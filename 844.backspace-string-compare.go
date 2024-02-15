/*
 * @lc app=leetcode id=844 lang=golang
 *
 * [844] Backspace String Compare
 */

// @lc code=start

/* 
 * Is the input string a unicode type? if is -> use rune. If it are ASCII characters we can use byte
*/
func backspaceCompare(s string, t string) bool {
	return processStr(s) == processStr(t)
}

func processStr(s string) string{
	/*** Use two pointer ***/
	// Time O(m+n) Space O(max(len(s),len(t)))
	slow, fast:= 0,0
	res := []byte(s)

	for fast < len(s){
		if res[fast] != '#'{
			res[slow] = res[fast]
            slow++
            fast++
		}else{
            fast++
			if slow>0{
				slow --
			}
		}
	}
	return string(res[:slow])
	/***  Use Rune ***/
	// myStack := make([]rune,0)

	// for _,c := range s{  // type of c  is Rune not byte
	// 	if c == '#'{
	// 		if len(myStack)>0{
	// 			myStack = myStack[:len(myStack)-1]
	// 		}
	// 	}else{
	// 		myStack = append(myStack,c)
	// 	}
	// }
	// return string(myStack)
	/***  Use Byte ***/

	// myStack := make([]byte,0)

	// for i:=0; i<len(s);i++{
	// 	c:=s[i]  // type of c is Byte 
	// 	if c == '#'{
	// 		if len(myStack)>0{
	// 			myStack = myStack[:len(myStack)-1]
	// 		}
	// 	}else{
	// 		myStack = append(myStack,c)
	// 	}
	// }
	// return string(myStack)
}
// @lc code=end

