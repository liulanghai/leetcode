/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 *
 * https://leetcode.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (26.42%)
 * Total Accepted:    469.7K
 * Total Submissions: 1.8M
 * Testcase Example:  '"babad"'
 *
 * Given a string s, find the longest palindromic substring in s. You may
 * assume that the maximum length of s is 1000.
 *
 * Example 1:
 *
 *
 * Input: "babad"
 * Output: "bab"
 * Note: "aba" is also a valid answer.
 *
 *
 * Example 2:
 *
 *
 * Input: "cbbd"
 * Output: "bb"
 *
 *
 */

func longestPalindrome(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	result := s[0:1]
	for i := 0; i < l; i++ {
		step := 0 //已s[i]为对称轴
		for i-step >= 0 && i+step < l {
			if s[i-step] == s[i+step] {
				tmp := s[i-step : i+step+1]
				step = step + 1
				if len(tmp) > len(result) {
					result = tmp
				}
			} else {
				break
			}
		}

		step2 := 0 //cbbc 类型，已i i+1对称
		for i-step2 >= 0 && i+step2+1 < l {
			if s[i-step2] == s[i+step2+1] {
				tmp := s[i-step2 : i+step2+1+1]
				step2 = step2 + 1
				if len(tmp) > len(result) {
					result = tmp
				}
			} else {
				break
			}
		}
	}
	return result
}
