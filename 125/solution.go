// https://leetcode.com/problems/valid-palindrome/description/?envType=company&envId=facebook&favoriteSlug=facebook-thirty-days

package solution

func isPalindrome(s string) bool {
	length := len(s)
	l := 0
	r := length - 1

	for l <= r {
		a := toLower(s[l])
		b := toLower(s[r])
		if !isLetterOrDigit(a) {
			l++
			continue
		}
		if !isLetterOrDigit(b) {
			r--
			continue
		}

		if a != b {
			return false
		}
		l++
		r--

	}

	return true

}

func isLetterOrDigit(c byte) bool {

	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9')
}

func toLower(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		return c - 'A' + 'a'
	}

	return c
}
