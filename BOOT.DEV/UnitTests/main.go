package main

func getMonthlyPrice(tier string) int {
	switch tier {
	case "basic":
		return 10000 // $100.00 converted to pennies
	case "premium":
		return 15000 // $150.00 converted to pennies
	case "enterprise":
		return 50000 // $500.00 converted to pennies
	default:
		return 0 // Return 0 pennies for any other tier
	}
}
