package atoi

import (
	"errors"
	"math"
	"strconv"
)

// error define
var (
	ErrNegativeNumber = errors.New("negative number")
	ErrOverflow       = errors.New("overflow")
	ErrInvalidType    = errors.New("invalid type")
)

// Uint32 converts string or integer to uint32
// return zero is i is nil
func Uint32(i interface{}) (uint32, error) {
	switch v := i.(type) {
	case nil:
		return 0, nil
	case string:
		n, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(n), nil
	case int:
		if v < 0 {
			return 0, ErrNegativeNumber
		}
		if v > math.MaxUint32 {
			return 0, ErrOverflow
		}
		return uint32(v), nil
	case int64:
		if v < 0 {
			return 0, ErrNegativeNumber
		}
		if v > math.MaxUint32 {
			return 0, ErrOverflow
		}
		return uint32(v), nil
	case int32:
		if v < 0 {
			return 0, ErrNegativeNumber
		}
		return uint32(v), nil
	case int16:
		if v < 0 {
			return 0, ErrNegativeNumber
		}
		return uint32(v), nil
	case int8:
		if v < 0 {
			return 0, ErrNegativeNumber
		}
		return uint32(v), nil

	case uint:
		if v > math.MaxUint32 {
			return 0, ErrOverflow
		}
		return uint32(v), nil
	case uint64:
		if v > math.MaxUint32 {
			return 0, ErrOverflow
		}
		return uint32(v), nil
	case uint32:
		return uint32(v), nil // 这里保持队形，方便批量替换
	case uint16:
		return uint32(v), nil
	case uint8:
		return uint32(v), nil
	default:
		return 0, ErrInvalidType
	}
}

// Int32 converts string or integer to int32
// return zero is i is nil
func Int32(i interface{}) (int32, error) {
	switch v := i.(type) {
	case nil:
		return 0, nil
	case string:
		n, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return 0, err
		}
		return int32(n), nil
	case int:
		if v < 0 && v < math.MinInt32 || v > 0 && v > math.MaxInt32 {
			return 0, ErrOverflow
		}
		return int32(v), nil
	case int64:
		if v < 0 && v < math.MinInt32 || v > 0 && v > math.MaxInt32 {
			return 0, ErrOverflow
		}
		return int32(v), nil
	case int32:
		return int32(v), nil
	case int16:
		return int32(v), nil
	case int8:
		return int32(v), nil
	case uint:
		if v > math.MaxInt32 {
			return 0, ErrOverflow
		}
		return int32(v), nil
	case uint64:
		if v > math.MaxInt32 {
			return 0, ErrOverflow
		}
		return int32(v), nil
	case uint32:
		if v > math.MaxInt32 {
			return 0, ErrOverflow
		}
		return int32(v), nil
	case uint16:
		return int32(v), nil
	case uint8:
		return int32(v), nil
	default:
		return 0, ErrInvalidType
	}
}

// Uint64 converts string or integer to uint64
// return zero is i is nil
func Uint64(i interface{}) (uint64, error) {
	switch v := i.(type) {
	case nil:
		return 0, nil
	case string:
		n, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return 0, err
		}
		return uint64(n), nil
	case int:
		if v < 0 {
			return 0, ErrNegativeNumber
		}
		return uint64(v), nil
	case int64:
		if v < 0 {
			return 0, ErrNegativeNumber
		}
		return uint64(v), nil
	case int32:
		if v < 0 {
			return 0, ErrNegativeNumber
		}
		return uint64(v), nil
	case int16:
		if v < 0 {
			return 0, ErrNegativeNumber
		}
		return uint64(v), nil
	case int8:
		if v < 0 {
			return 0, ErrNegativeNumber
		}
		return uint64(v), nil
	case uint:
		return uint64(v), nil
	case uint64:
		return uint64(v), nil
	case uint32:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint8:
		return uint64(v), nil
	default:
		return 0, ErrInvalidType
	}
}

// Int64 converts string or integer to int64
// return zero is i is nil
func Int64(i interface{}) (int64, error) {
	switch v := i.(type) {
	case nil:
		return 0, nil
	case string:
		n, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, err
		}
		return int64(n), nil
	case int:
		return int64(v), nil
	case int64:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int8:
		return int64(v), nil

	case uint:
		if v > math.MaxInt64 {
			return 0, ErrOverflow
		}
		return int64(v), nil
	case uint64:
		if v > math.MaxInt64 {
			return 0, ErrOverflow
		}
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	default:
		return 0, ErrInvalidType
	}
}
