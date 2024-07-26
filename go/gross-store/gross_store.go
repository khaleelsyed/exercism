package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if unitValue, unitExists := units[unit]; unitExists {
		_, itemExists := bill[item]
		if itemExists {
			bill[item] += unitValue
		} else {
			bill[item] = unitValue
		}
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemQuantity, itemExists := bill[item]
	unitValue, unitExists := units[unit]
	finalQuantity := itemQuantity - unitValue
	validRemoval := itemExists && unitExists && (finalQuantity >= 0)
	if validRemoval && finalQuantity > 0 {
		bill[item] -= unitValue
	} else if validRemoval && finalQuantity == 0 {
		delete(bill, item)
	}
	return validRemoval
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemQuantity, itemExists := bill[item]
	if !itemExists {
		return 0, itemExists
	}
	return itemQuantity, itemExists
}
