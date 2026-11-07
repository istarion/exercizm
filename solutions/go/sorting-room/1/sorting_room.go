package sorting

import "fmt"
import "strconv"

func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

func ExtractFancyNumber(fnb FancyNumberBox) int {
	fn, ok := fnb.(FancyNumber)
    if !ok {
        return 0
    }
    d, err := strconv.Atoi(fn.n)
    if err != nil {
        return 0
    }
    return d
}

func DescribeFancyNumberBox(fnb FancyNumberBox) string {
    return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(ExtractFancyNumber(fnb)))
    
}

func DescribeAnything(i any) string {
	switch i.(type) {
    case int: return DescribeNumber(float64(i.(int)))
    case float64: return DescribeNumber(i.(float64))
    case NumberBox: return DescribeNumberBox(i.(NumberBox))
    case FancyNumberBox: return DescribeFancyNumberBox(i.(FancyNumberBox))
    default: return "Return to sender"
    }
}
