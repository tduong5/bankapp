package util

// Constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

// IsSupporterCurrency returns true if currency is supported
func IsSupporterCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
