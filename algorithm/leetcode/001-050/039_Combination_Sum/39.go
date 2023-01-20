package _39_Combination_Sum

func combinationSum(candidates []int, target int) [][]int {
	var dfs func(result *[][]int, candidates []int, target int, selection *[]int, index int)
	dfs = func(result *[][]int, candidates []int, target int, selection *[]int, index int) {
		if target < 0 {
			return
		} else if target == 0 {
			temp := make([]int, len(*selection), len(*selection))
			copy(temp, *selection)
			*result = append(*result, temp)
			return
		}
		for i := index; i < len(candidates); i++ {
			*selection = append(*selection, candidates[i])
			dfs(result, candidates, target-candidates[i], selection, i)
			*selection = (*selection)[:len(*selection)-1]
		}
	}
	var result [][]int
	dfs(&result, candidates, target, &[]int{}, 0)
	return result
}

func combinationSum1(candidates []int, target int) [][]int {
	var dfs func(result *[][]int, candidates []int, target int, selection *[]int, index int)
	dfs = func(result *[][]int, candidates []int, target int, selection *[]int, index int) {
		if target < 0 {
			return
		} else if target == 0 {
			temp := make([]int, len(*selection), len(*selection))
			copy(temp, *selection)
			*result = append(*result, temp)
			return
		}
		for i := index; i < len(candidates); i++ {
			tempTarget := target - candidates[i]
			*selection = append(*selection, candidates[i])
			dfs(result, candidates, tempTarget, selection, i)
			*selection = (*selection)[:len(*selection)-1]
		}
	}
	var result [][]int
	dfs(&result, candidates, target, &[]int{}, 0)
	return result
}

//func combinationSum1(candidates []int, target int) [][]int {
//	var result [][]int
//	dfs(&result, &[]int{}, candidates, target, 0)
//	return result
//}
//func dfs(result *[][]int, selection *[]int, candidates []int, target int, index int) {
//	if target < 0 {
//		return
//	} else if target == 0 {
//		*result = append(*result, append([]int{}, *selection...))
//		return
//	}
//	for i := index; i < len(candidates); i++ {
//		tempTarget := target - candidates[i]
//		*selection = append(*selection, candidates[i])
//		dfs(result, selection, candidates, tempTarget, i)
//		*selection = (*selection)[:len(*selection)-1]
//	}
//}
