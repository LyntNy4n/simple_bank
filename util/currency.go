package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

// IsValidCurrency checks if the currency is valid or not
func IsValidCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}