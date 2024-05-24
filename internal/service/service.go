package service

func Exchange(amount int, banknotes []int, currentCombination []int, startIndex int, result *[][]int) {
	if amount == 0 {
		combination := make([]int, len(currentCombination))
		copy(combination, currentCombination)
		*result = append(*result, combination)
		return
	}

	for i := startIndex; i < len(banknotes); i++ {
		if banknotes[i] <= amount {
			currentCombination = append(currentCombination, banknotes[i])
			Exchange(amount-banknotes[i], banknotes, currentCombination, i, result)
			currentCombination = currentCombination[:len(currentCombination)-1]
		}
	}

}
