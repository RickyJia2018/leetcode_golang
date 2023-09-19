/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
    myHashTable := make(map[byte]int)

    for i := 0; i < len(magazine); i++ {
        myHashTable[magazine[i]]++
    }

    for i := 0; i < len(ransomNote); i++ {
        count, ok := myHashTable[ransomNote[i]]
        if count <= 0 || !ok {
            return false
        } else {
            myHashTable[ransomNote[i]]--
        }
    }

    return true
}

// @lc code=end

