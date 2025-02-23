package solution

import "strings"

// https://leetcode.com/problems/simplify-path/submissions/1552440573/?envType=company&envId=facebook&favoriteSlug=facebook-thirty-days

func simplifyPath(path string) string {
	locations := strings.Split(path, "/")

	stack := []string{}

	for _, location := range locations {

		if location == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if location != "." && len(location) > 0 {
			stack = append(stack, location)
		}
	}
	return "/" + strings.Join(stack, "/")

}
