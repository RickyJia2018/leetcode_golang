/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start
func reverseWords(s string) string {
    b := []byte(s)

    // 移除前面、中间、后面存在的多余空格
    slow := 0
    for i := 0; i < len(b); i++ {
        if b[i] != ' ' {
            if slow != 0 {
                b[slow] = ' '
                slow++
            }
            for i < len(b) && b[i] != ' ' { // 复制逻辑
                b[slow] = b[i]
                slow++
                i++
            }
        }
    }
    b = b[0:slow]
    
    // 翻转整个字符串
    reverse(b)
    // 翻转每个单词
    pre := 0
    for i := 0; i <= len(b); i++ {
        if i == len(b) || b[i] == ' ' {
            reverse(b[last:i])
            pre = i + 1
        }
    }
    return string(b)
}

func reverse(b []byte) {
    for left,right:=0,len(b) -1; left < right;left++,right-- {
        b[left], b[right] = b[right], b[left]
    }
}
// @lc code=end

