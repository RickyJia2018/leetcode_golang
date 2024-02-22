/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start

func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

func removeExtraSpaces(b []byte) []byte {
	slow := 0
	fast := 0
	for fast < len(b) {
		if b[fast] == ' ' {
			// ski space first ( must )
			for fast < len(b) && b[fast] == ' ' {
				fast++
			}
			// add space (handle leading spaces and trailing spaces)
			if slow != 0 && fast != len(b) {
				b[slow] = ' '
				slow++
			}

		} else {
			b[slow] = b[fast]
			fast++
			slow++
		}
	}
	return b[:slow] // this line created a new slice, thus need to return the value
}
func reverseWords(s string) string {
	b := []byte(s)
	// remove extra spaces
	b = removeExtraSpaces(b)
	// reverse string
	reverse(b)
	//reverse words
	start := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			reverse(b[start:i])
			start = i + 1
		}
	}

	return string(b)
}

// func reverseWords(s string) string {
//     b := []byte(s)

//     // 移除前面、中间、后面存在的多余空格
//     slow := 0
//     for i := 0; i < len(b); i++ {
//         if b[i] != ' ' {
//             if slow != 0 {
//                 b[slow] = ' '
//                 slow++
//             }
//             for i < len(b) && b[i] != ' ' { // 复制逻辑
//                 b[slow] = b[i]
//                 slow++
//                 i++
//             }
//         }

//     }
//     b = b[0:slow]

//	    // 翻转整个字符串
//	    reverse(b)
//	    // 翻转每个单词
//	    last := 0
//	    for i := 0; i <= len(b); i++ {
//	        if i == len(b) || b[i] == ' ' {
//	            reverse(b[last:i])
//	            last = i + 1
//	        }
//	    }
//	    return string(b)
//	}
// @lc code=end

