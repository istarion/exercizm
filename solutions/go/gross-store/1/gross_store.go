package gross

func Units() map[string]int {
	return map[string]int{
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen": 6,
        "dozen": 12,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
    }
}

func NewBill() map[string]int {
	return make(map[string]int, 0)
}

func AddItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := units[unit]; !ok {
        return false
    }
    bill[item] += units[unit]
    return true
}

func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := bill[item]; !ok {
        return false
    }
    if _, ok := units[unit]; !ok {
        return false
    }
    quantity := bill[item] - units[unit]
    if quantity < 0 {
        return false
    }
    if quantity == 0 {
        delete(bill, item)
        return true
    }
	bill[item] = quantity
    return true
}

func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, ok := bill[item]
    return quantity, ok
}
