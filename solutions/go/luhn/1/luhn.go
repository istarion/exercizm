package luhn

func Valid(id string) bool {
	numbers := make([]int, 0)
    for _, c := range id {
        if c == ' ' {
            continue
        }
        if c >= '0' && c <= '9' {
            digit := int(c - '0')
            numbers = append(numbers, digit)
        } else {
            return false
        }
    }
    if len(numbers) < 2 {
        return false
    }
    for i := len(numbers) - 2; i >= 0; i -= 2 {
        numbers[i] = numbers[i] * 2
        if numbers[i] > 9 {
            numbers[i] -= 9
        }
    }
    sum := 0
    for _, d := range numbers {
        sum += d
    }
    return sum % 10 == 0
}
