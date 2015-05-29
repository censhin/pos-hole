package util

import "log"

const (
    MAX_STRONT_SMALL = 4166
    MAX_STRONT_MEDIUM = 8333
    MAX_STRONT_LARGE = 16666
    MAX_HOURS = 41.7
)

// Formula - (hr/stront) + time received evemail = time POS comes out of reinforced
// TODO: change time from float64 to a time library value
func PosRefFormula(units float64, time float64, pos_type string) (float64, error) {
    if time > MAX_HOURS {
        log.Panic("Time entered is greater than the max allowed.")
        // TODO: Create a custom error message to return
        return 0.0, nil
    }

    switch pos_type {
    case "small":
        if units > MAX_STRONT_SMALL {
            log.Panic("Stront entered is greater than the max allowed.")
            // TODO: Create a custom error message to return
            return 0.0, nil
        }

        units = units/100
    case "medium":
        if units > MAX_STRONT_MEDIUM {
            log.Panic("Stront entered is greater than the max allowed.")
            // TODO: Create a custom error message to return
            return 0.0, nil
        }

        units = units/200
    case "large":
        if units > MAX_STRONT_LARGE {
            log.Panic("Stront entered is greater than the max allowed.")
            // TODO: Create a custom error message to return
            return 0.0, nil
        }

        units = units/400
    }

    return (units + time), nil
}

func SmallPosRefFormula(unit float64, time float64) (float64, error) {
    return PosRefFormula(unit, time, "small")
}

func MediumPosRefFormula(unit float64, time float64) (float64, error) {
    return PosRefFormula(unit, time, "medium")
}

func LargePosRefFormula(unit float64, time float64) (float64, error) {
    return PosRefFormula(unit, time, "large")
}

