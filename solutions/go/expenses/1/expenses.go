package expenses

import "errors"

type Record struct {
	Day      int
	Amount   float64
	Category string
}

type DaysPeriod struct {
	From int
	To   int
}

func Filter(in []Record, predicate func(Record) bool) []Record {
	result := make([]Record, 0)
    for _, r := range in {
        if predicate(r) {
        	result = append(result, r)
        }
    }
	return result
}

func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
        return r.Day >= p.From && r.Day <= p.To
    }
}

func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
        return r.Category == c
    }
}

func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	sum := 0.0
    for _, r := range Filter(in, ByDaysPeriod(p)) {
        sum += float64(r.Amount)
    }
    return sum
}

func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	catFiltered := Filter(in, ByCategory(c))
    if len(catFiltered) == 0 {
        return 0.0, errors.New("unknown category " + c)
    }
    return TotalByPeriod(catFiltered, p), nil
}
