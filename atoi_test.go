package atoi

import (
	"math"
	"testing"
)

func Test_Uint32(t *testing.T) {
	{
		var f float32
		_, err := Uint32(f)
		if err == nil {
			t.Fatalf("failed to Uint32: float32 not supported")
		}
	}

	{
		_, err := Uint32(nil)
		if err != nil {
			t.Fatalf("failed to convert nil")
		}
	}

	{
		actual, err := Uint32("4294967295")
		if err != nil {
			t.Fatalf("failed to Uint32: %s", err)
		}
		if expect := uint32(4294967295); actual != expect {
			t.Fatalf("failed to Uint32, expect %d, actual %d", expect, actual)
		}
	}
	{
		_, err := Uint32("2222233333323233332")
		if err == nil {
			t.Fatalf("cannot convert string with great than max uint32")
		}
	}
	{
		_, err := Uint32("-429496")
		if err == nil {
			t.Fatalf("cannot convert string leading with '-'")
		}
	}

	{
		expect := int8(127)
		actual, err := Uint32(expect)
		if err != nil {
			t.Fatalf("failed to Uint32: %s", err)
		}
		if actual != uint32(expect) {
			t.Fatalf("failed to Uint32, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := int8(-127)
		_, err := Uint32(expect)
		if err == nil {
			t.Fatalf("failed to Uint32: cannot convert negative number")
		}
	}
	{
		expect := uint8(128)
		actual, err := Uint32(expect)
		if err != nil {
			t.Fatalf("failed to Uint32: %s", err)
		}
		if actual != uint32(expect) {
			t.Fatalf("failed to Uint32, expect %d, actual %d", expect, actual)
		}
	}

	{
		expect := int16(3332)
		actual, err := Uint32(expect)
		if err != nil {
			t.Fatalf("failed to Uint32: %s", err)
		}
		if actual != uint32(expect) {
			t.Fatalf("failed to Uint32, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := int16(-3332)
		_, err := Uint32(expect)
		if err == nil {
			t.Fatalf("failed to Uint32: cannot convert negative number")
		}
	}
	{
		expect := uint16(13332)
		actual, err := Uint32(expect)
		if err != nil {
			t.Fatalf("failed to Uint32: %s", err)
		}
		if actual != uint32(expect) {
			t.Fatalf("failed to Uint32, expect %d, actual %d", expect, actual)
		}
	}

	{
		expect := int32(323233332)
		actual, err := Uint32(expect)
		if err != nil {
			t.Fatalf("failed to Uint32: %s", err)
		}
		if actual != uint32(expect) {
			t.Fatalf("failed to Uint32, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := int32(-323233332)
		_, err := Uint32(expect)
		if err == nil {
			t.Fatalf("failed to Uint32: cannot convert negative number")
		}
	}
	{
		expect := uint32(3323233332)
		actual, err := Uint32(expect)
		if err != nil {
			t.Fatalf("failed to Uint32: %s", err)
		}
		if actual != uint32(expect) {
			t.Fatalf("failed to Uint32, expect %d, actual %d", expect, actual)
		}
	}

	{
		expect := int64(3323233332)
		actual, err := Uint32(expect)
		if err != nil {
			t.Fatalf("failed to Uint32: %s", err)
		}
		if actual != uint32(expect) {
			t.Fatalf("failed to Uint32, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := int64(2222233333323233332)
		_, err := Uint32(expect)
		if err == nil {
			t.Fatalf("cannot convert max int64")
		}
	}
	{
		expect := int64(-2222233333323233332)
		_, err := Uint32(expect)
		if err == nil {
			t.Fatalf("failed to Uint32: cannot convert negative number")
		}
	}
	{
		expect := uint64(2323233332)
		actual, err := Uint32(expect)
		if err != nil {
			t.Fatalf("failed to Uint32: %s", err)
		}
		if actual != uint32(expect) {
			t.Fatalf("failed to Uint32, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := uint64(2222233333323233332)
		_, err := Uint32(expect)
		if err == nil {
			t.Fatalf("cannot convert max uint64")
		}
	}

	{
		expect := int(2323233332)
		actual, err := Uint32(expect)
		if err != nil {
			t.Fatalf("failed to Uint32: %s", err)
		}
		if actual != uint32(expect) {
			t.Fatalf("failed to Uint32, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := int(-2222233333323233332)
		_, err := Uint32(expect)
		if err == nil {
			t.Fatalf("failed to Uint32: cannot convert negative number")
		}
	}
	{
		expect := int(2222233333323233332)
		_, err := Uint32(expect)
		if err == nil {
			t.Fatalf("cannot convert max int64")
		}
	}
	{
		expect := uint(2323233332)
		actual, err := Uint32(expect)
		if err != nil {
			t.Fatalf("failed to Uint32: %s", err)
		}
		if actual != uint32(expect) {
			t.Fatalf("failed to Uint32, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := uint(2222233333323233332)
		_, err := Uint32(expect)
		if err == nil {
			t.Fatalf("cannot convert max uint")
		}
	}
}

func Test_Int32(t *testing.T) {
	{
		var f float32
		_, err := Int32(f)
		t.Logf("%s", err)
		if err == nil {
			t.Fatalf("failed to Int32: float32 not supported")
		}
	}

	{
		_, err := Int32(nil)
		if err != nil {
			t.Fatalf("failed to convert nil")
		}
	}

	{
		actual, err := Int32("429496729")
		if err != nil {
			t.Fatalf("failed to Int32: %s", err)
		}
		if expect := int32(429496729); actual != expect {
			t.Fatalf("failed to Int32, expect %d, actual %d", expect, actual)
		}
	}
	{
		actual, err := Int32("-429496729")
		if err != nil {
			t.Fatalf("failed to Int32: %s", err)
		}
		if expect := int32(-429496729); actual != expect {
			t.Fatalf("failed to Int32, expect %d, actual %d", expect, actual)
		}
	}
	{
		_, err := Int32("4294967295")
		if err == nil {
			t.Fatalf("cannot convert max int32")
		}
	}

	{
		expect := int8(127)
		actual, err := Int32(expect)
		if err != nil {
			t.Fatalf("failed to Int32: %s", err)
		}
		if actual != int32(expect) {
			t.Fatalf("failed to Int32, expect %d, actual %d", expect, actual)
		}
	}

	{
		expect := uint8(128)
		actual, err := Int32(expect)
		if err != nil {
			t.Fatalf("failed to Int32: %s", err)
		}
		if actual != int32(expect) {
			t.Fatalf("failed to Int32, expect %d, actual %d", expect, actual)
		}
	}

	{
		expect := int16(3332)
		actual, err := Int32(expect)
		if err != nil {
			t.Fatalf("failed to Int32: %s", err)
		}
		if actual != int32(expect) {
			t.Fatalf("failed to Int32, expect %d, actual %d", expect, actual)
		}
	}

	{
		expect := uint16(13332)
		actual, err := Int32(expect)
		if err != nil {
			t.Fatalf("failed to Int32: %s", err)
		}
		if actual != int32(expect) {
			t.Fatalf("failed to Int32, expect %d, actual %d", expect, actual)
		}
	}

	{
		expect := int32(math.MaxInt32)
		actual, err := Int32(expect)
		if err != nil {
			t.Fatalf("failed to Int32: %s", err)
		}
		if actual != int32(expect) {
			t.Fatalf("failed to Int32, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := int32(math.MinInt32)
		actual, err := Int32(expect)
		if err != nil {
			t.Fatalf("failed to Int32: %s", err)
		}
		if actual != int32(expect) {
			t.Fatalf("failed to Int32, expect %d, actual %d", expect, actual)
		}
	}

	{
		expect := uint32(math.MaxInt32)
		actual, err := Int32(expect)
		if err != nil {
			t.Fatalf("failed to Int32: %s", err)
		}
		if actual != int32(expect) {
			t.Fatalf("failed to Int32, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := uint32(math.MaxUint32)
		_, err := Int32(expect)
		if err == nil {
			t.Fatalf("cannot convert max uint32")
		}
	}

	{
		expect := int64(math.MaxInt32)
		actual, err := Int32(expect)
		if err != nil {
			t.Fatalf("failed to Int32: %s", err)
		}
		if actual != int32(expect) {
			t.Fatalf("failed to Int32, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := int64(math.MinInt32)
		actual, err := Int32(expect)
		if err != nil {
			t.Fatalf("failed to Int32: %s", err)
		}
		if actual != int32(expect) {
			t.Fatalf("failed to Int32, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := int64(math.MaxInt64)
		_, err := Int32(expect)
		if err == nil {
			t.Fatalf("cannot convert max int64")
		}
	}

	{
		expect := uint64(math.MaxInt32)
		actual, err := Int32(expect)
		if err != nil {
			t.Fatalf("failed to Int32: %s", err)
		}
		if actual != int32(expect) {
			t.Fatalf("failed to Int32, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := uint64(math.MaxUint64)
		_, err := Int32(expect)
		if err == nil {
			t.Fatalf("Uint() cannot convert max uint64")
		}
	}

	{
		expect := int(math.MaxInt32)
		actual, err := Int32(expect)
		if err != nil {
			t.Fatalf("failed to Int32: %s", err)
		}
		if actual != int32(expect) {
			t.Fatalf("failed to Int32, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := int(math.MinInt32)
		actual, err := Int32(expect)
		if err != nil {
			t.Fatalf("failed to Int32: %s", err)
		}
		if actual != int32(expect) {
			t.Fatalf("failed to Int32, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := uint(math.MaxInt32)
		actual, err := Int32(expect)
		if err != nil {
			t.Fatalf("failed to Int32: %s", err)
		}
		if actual != int32(expect) {
			t.Fatalf("failed to Int32, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := uint(math.MaxUint64)
		_, err := Int32(expect)
		if err == nil {
			t.Fatalf("Uint() cannot convert max uint")
		}
	}
}

func Test_Uint64(t *testing.T) {
	{
		var f float32
		_, err := Uint64(f)
		t.Logf("%s", err)
		if err == nil {
			t.Fatalf("failed to Uint64: float32 not supported")
		}
	}

	{
		_, err := Uint64(nil)
		if err != nil {
			t.Fatalf("failed to convert nil")
		}
	}

	{
		actual, err := Uint64("42949672950000")
		if err != nil {
			t.Fatalf("failed to Uint64: %s", err)
		}
		if expect := uint64(42949672950000); actual != expect {
			t.Fatalf("failed to Uint64, expect %d, actual %d", expect, actual)
		}
	}
	{
		_, err := Uint32("-429496")
		if err == nil {
			t.Fatalf("cannot convert string leading with '-'")
		}
	}

	{
		expect := int8(127)
		actual, err := Uint64(expect)
		if err != nil {
			t.Fatalf("failed to Uint64: %s", err)
		}
		if actual != uint64(expect) {
			t.Fatalf("failed to Uint64, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := int8(-127)
		_, err := Uint64(expect)
		if err == nil {
			t.Fatalf("failed to Uint64: cannot convert negative number")
		}
	}
	{
		expect := uint8(128)
		actual, err := Uint64(expect)
		if err != nil {
			t.Fatalf("failed to Uint64: %s", err)
		}
		if actual != uint64(expect) {
			t.Fatalf("failed to Uint64, expect %d, actual %d", expect, actual)
		}
	}

	{
		expect := int16(3332)
		actual, err := Uint64(expect)
		if err != nil {
			t.Fatalf("failed to Uint64: %s", err)
		}
		if actual != uint64(expect) {
			t.Fatalf("failed to Uint64, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := int16(-3332)
		_, err := Uint64(expect)
		if err == nil {
			t.Fatalf("failed to Uint64: cannot convert negative number")
		}
	}
	{
		expect := uint16(13332)
		actual, err := Uint64(expect)
		if err != nil {
			t.Fatalf("failed to Uint64: %s", err)
		}
		if actual != uint64(expect) {
			t.Fatalf("failed to Uint64, expect %d, actual %d", expect, actual)
		}
	}

	{
		expect := int32(323233332)
		actual, err := Uint64(expect)
		if err != nil {
			t.Fatalf("failed to Uint64: %s", err)
		}
		if actual != uint64(expect) {
			t.Fatalf("failed to Uint64, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := int32(-323233332)
		_, err := Uint64(expect)
		if err == nil {
			t.Fatalf("failed to Uint64: cannot convert negative number")
		}
	}
	{
		expect := uint32(3323233332)
		actual, err := Uint64(expect)
		if err != nil {
			t.Fatalf("failed to Uint64: %s", err)
		}
		if actual != uint64(expect) {
			t.Fatalf("failed to Uint64, expect %d, actual %d", expect, actual)
		}
	}

	{
		expect := int64(2222233333323233332)
		actual, err := Uint64(expect)
		if err != nil {
			t.Fatalf("failed to Uint64: %s", err)
		}
		if actual != uint64(expect) {
			t.Fatalf("failed to Uint64, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := int64(-2222233333323233332)
		_, err := Uint64(expect)
		if err == nil {
			t.Fatalf("failed to Uint64: cannot convert negative number")
		}
	}
	{
		expect := uint64(12222233333323233332)
		actual, err := Uint64(expect)
		if err != nil {
			t.Fatalf("failed to Uint64: %s", err)
		}
		if actual != uint64(expect) {
			t.Fatalf("failed to Uint64, expect %d, actual %d", expect, actual)
		}
	}

	{
		expect := int(2222233333323233332)
		actual, err := Uint64(expect)
		if err != nil {
			t.Fatalf("failed to Uint64: %s", err)
		}
		if actual != uint64(expect) {
			t.Fatalf("failed to Uint64, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := int(-2222233333323233332)
		_, err := Uint64(expect)
		if err == nil {
			t.Fatalf("failed to Uint64: cannot convert negative number")
		}
	}
	{
		expect := uint(12222233333323233332)
		actual, err := Uint64(expect)
		if err != nil {
			t.Fatalf("failed to Uint64: %s", err)
		}
		if actual != uint64(expect) {
			t.Fatalf("failed to Uint64, expect %d, actual %d", expect, actual)
		}
	}
}

func Test_Int64(t *testing.T) {
	{
		var f float32
		_, err := Int64(f)
		t.Logf("%s", err)
		if err == nil {
			t.Fatalf("failed to Int64: float32 not supported")
		}
	}

	{
		_, err := Int64(nil)
		if err != nil {
			t.Fatalf("failed to convert nil")
		}
	}

	{
		actual, err := Int64("42949672950000")
		if err != nil {
			t.Fatalf("failed to Int64: %s", err)
		}
		if expect := int64(42949672950000); actual != expect {
			t.Fatalf("failed to Int64, expect %d, actual %d", expect, actual)
		}
	}

	{
		expect := int8(127)
		actual, err := Int64(expect)
		if err != nil {
			t.Fatalf("failed to Int64: %s", err)
		}
		if actual != int64(expect) {
			t.Fatalf("failed to Int64, expect %d, actual %d", expect, actual)
		}
	}

	{
		expect := uint8(128)
		actual, err := Int64(expect)
		if err != nil {
			t.Fatalf("failed to Int64: %s", err)
		}
		if actual != int64(expect) {
			t.Fatalf("failed to Int64, expect %d, actual %d", expect, actual)
		}
	}

	{
		expect := int16(3332)
		actual, err := Int64(expect)
		if err != nil {
			t.Fatalf("failed to Int64: %s", err)
		}
		if actual != int64(expect) {
			t.Fatalf("failed to Int64, expect %d, actual %d", expect, actual)
		}
	}

	{
		expect := uint16(13332)
		actual, err := Int64(expect)
		if err != nil {
			t.Fatalf("failed to Int64: %s", err)
		}
		if actual != int64(expect) {
			t.Fatalf("failed to Int64, expect %d, actual %d", expect, actual)
		}
	}

	{
		expect := int32(323233332)
		actual, err := Int64(expect)
		if err != nil {
			t.Fatalf("failed to Int64: %s", err)
		}
		if actual != int64(expect) {
			t.Fatalf("failed to Int64, expect %d, actual %d", expect, actual)
		}
	}

	{
		expect := uint32(3323233332)
		actual, err := Int64(expect)
		if err != nil {
			t.Fatalf("failed to Int64: %s", err)
		}
		if actual != int64(expect) {
			t.Fatalf("failed to Int64, expect %d, actual %d", expect, actual)
		}
	}

	{
		expect := int64(2222233333323233332)
		actual, err := Int64(expect)
		if err != nil {
			t.Fatalf("failed to Int64: %s", err)
		}
		if actual != int64(expect) {
			t.Fatalf("failed to Int64, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := uint64(2222233333323233332)
		actual, err := Int64(expect)
		if err != nil {
			t.Fatalf("failed to Int64: %s", err)
		}
		if actual != int64(expect) {
			t.Fatalf("failed to Int64, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := uint64(math.MaxUint64)
		_, err := Int64(expect)
		if err == nil {
			t.Fatalf("Uint() cannot convert max uint64")
		}
	}

	{
		expect := int(2222233333323233332)
		actual, err := Int64(expect)
		if err != nil {
			t.Fatalf("failed to Int64: %s", err)
		}
		if actual != int64(expect) {
			t.Fatalf("failed to Int64, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := uint(2222233333323233332)
		actual, err := Int64(expect)
		if err != nil {
			t.Fatalf("failed to Int64: %s", err)
		}
		if actual != int64(expect) {
			t.Fatalf("failed to Int64, expect %d, actual %d", expect, actual)
		}
	}
	{
		expect := uint(math.MaxUint64)
		_, err := Int64(expect)
		if err == nil {
			t.Fatalf("Uint() cannot convert max uint")
		}
	}
}
