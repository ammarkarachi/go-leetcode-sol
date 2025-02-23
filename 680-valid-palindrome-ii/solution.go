// https://leetcode.com/problems/valid-palindrome-ii/description/?envType=company&envId=facebook&favoriteSlug=facebook-thirty-days
package solution

func validPalindrome(s string) bool {
	l := 0
	r := len(s)

	for l <= r {
		if s[l] == s[r] {
			l++
			r--
			continue
		} else {
			return isPalindrome(s, l+1, r) || isPalindrome(s, l, r-1)
		}

	}
	return true
}

func isPalindrome(s string, l int, r int) bool {
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true

}
