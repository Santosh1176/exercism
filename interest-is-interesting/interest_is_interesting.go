package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance < 0 {
		return 3.213
	} else if balance >= 0 && balance < 1000 {
		return 0.5
	} else if balance >= 1000 && balance < 5000 {
		return 1.621
	} else if balance >= 5000 {
		return 2.475
	}
	return 0.0
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	if balance < 0 {
		return balance * (3.213 / 100)
	} else if balance >= 0 && balance < 1000 {
		return balance * (0.5 / 100)
	} else if balance >= 1000 && balance < 5000 {
		return balance * (1.621 / 100)
	} else if balance >= 5000 {
		return balance * (2.475 / 100)
	}
	return 0.0
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	if balance < 0 {
		return balance + (balance * (3.213 / 100))
	} else if balance >= 0 && balance < 1000 {
		return balance + (balance * (0.5 / 100))
	} else if balance >= 1000 && balance < 5000 {
		return balance + (balance * (1.621 / 100))
	} else if balance >= 5000 {
		return balance + (balance * (2.475 / 100))
	}
	return 0.0
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	for year := 0; ; year++ {
		if balance >= targetBalance {
			return year
		}
		balance = AnnualBalanceUpdate(balance)
	}
}
