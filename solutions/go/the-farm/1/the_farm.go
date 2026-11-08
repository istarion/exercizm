package thefarm

import "errors"
import "fmt"

//type FodderCalculator interface {
//	FodderAmount(int) (float64, error)
//	FatteningFactor() (float64, error)
//}

var ErrInvalidNumberOfCows = errors.New("invalid number of cows")

type InvalidCowsError struct {
    cows int
    msg string
}

func (e InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.msg)
}

func DivideFood(fc FodderCalculator, cows int) (float64, error) {
    allFodder, err := fc.FodderAmount(cows)
    if err != nil {
        return 0.0, err
    }
    ff, err := fc.FatteningFactor()
    if err != nil {
        return 0.0, err
    }
    return allFodder / float64(cows) * ff, nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
    if cows <= 0 {
        return 0.0, ErrInvalidNumberOfCows
    }
    return DivideFood(fc, cows)
}

func ValidateNumberOfCows(cows int) error {
    if cows < 0 {
        return &InvalidCowsError{cows: cows, msg: "there are no negative cows"}
    }
    if cows == 0 {
        return &InvalidCowsError{cows: cows, msg: "no cows don't need food"}
    }
    return nil
}
