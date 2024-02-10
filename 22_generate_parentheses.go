package main

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	generateParenthesisRecursive(&result, n, 0, 0, "")
	return result
}

func generateParenthesisRecursive(result *[]string, target int, left int, right int, current string) {
	if len(current) == target*2 {
		*result = append(*result, current)
	}

	if left < target {
		generateParenthesisRecursive(result, target, left+1, right, current+"(")
	}
	if right < left {
		generateParenthesisRecursive(result, target, left, right+1, current+")")
	}
}
