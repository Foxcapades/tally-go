package tally

// ITally is a counter of type int.
type ITally int

// Cur returns the current int value of this ITally.
func (t ITally) Cur() int {
	return int(t)
}

// Add increases this counter by the given value, returning the previous
// int value of this ITally.
func (t *ITally) Add(i int) (cur int) {
	cur = int(*t)
	*t += ITally(i)

	return
}

// Inc increases this counter by 1, returning the previous int value of
// this ITally.
func (t *ITally) Inc() (cur int) {
	cur = int(*t)
	*t++

	return
}

// Dec decreases this counter by 1, returning the previous int value of
// this ITally.
func (t *ITally) Dec() (cur int) {
	cur = int(*t)
	*t--

	return
}

// Sub decreases this counter by the given value, returning the previous
// int value of this ITally.
func (t *ITally) Sub(i int) (cur int) {
	cur = int(*t)
	*t -= ITally(i)

	return
}

// Zero sets the current counter value to 0, returning the previous int
// value of this ITally.
func (t *ITally) Zero() (cur int) {
	cur = int(*t)
	*t = 0

	return
}

// ITally8 is a counter of type int8.
type ITally8 int8

// Cur returns the current int8 value of this ITally8.
func (t ITally8) Cur() int8 {
	return int8(t)
}

// Add increases this counter by the given value, returning the previous
// int8 value of this ITally8.
func (t *ITally8) Add(i int8) (cur int8) {
	cur = int8(*t)
	*t += ITally8(i)

	return
}

// Inc increases this counter by 1, returning the previous int8 value of
// this ITally8.
func (t *ITally8) Inc() (cur int8) {
	cur = int8(*t)
	*t++

	return
}

// Dec decreases this counter by 1, returning the previous int8 value of
// this ITally8.
func (t *ITally8) Dec() (cur int8) {
	cur = int8(*t)
	*t--

	return
}

// Sub decreases this counter by the given value, returning the previous
// int8 value of this ITally8.
func (t *ITally8) Sub(i int8) (cur int8) {
	cur = int8(*t)
	*t -= ITally8(i)

	return
}

// Zero sets the current counter value to 0, returning the previous int8
// value of this ITally8.
func (t *ITally8) Zero() (cur int8) {
	cur = int8(*t)
	*t = 0

	return
}

// ITally16 is a counter of type int16.
type ITally16 int16

// Cur returns the current int16 value of this ITally16.
func (t ITally16) Cur() int16 {
	return int16(t)
}

// Add increases this counter by the given value, returning the previous
// int16 value of this ITally16.
func (t *ITally16) Add(i int16) (cur int16) {
	cur = int16(*t)
	*t += ITally16(i)

	return
}

// Inc increases this counter by 1, returning the previous int16 value of
// this ITally16.
func (t *ITally16) Inc() (cur int16) {
	cur = int16(*t)
	*t++

	return
}

// Dec decreases this counter by 1, returning the previous int16 value of
// this ITally16.
func (t *ITally16) Dec() (cur int16) {
	cur = int16(*t)
	*t--

	return
}

// Sub decreases this counter by the given value, returning the previous
// int16 value of this ITally16.
func (t *ITally16) Sub(i int16) (cur int16) {
	cur = int16(*t)
	*t -= ITally16(i)

	return
}

// Zero sets the current counter value to 0, returning the previous int16
// value of this ITally16.
func (t *ITally16) Zero() (cur int16) {
	cur = int16(*t)
	*t = 0

	return
}

// ITally32 is a counter of type int32.
type ITally32 int32

// Cur returns the current int32 value of this ITally32.
func (t ITally32) Cur() int32 {
	return int32(t)
}

// Add increases this counter by the given value, returning the previous
// int32 value of this ITally32.
func (t *ITally32) Add(i int32) (cur int32) {
	cur = int32(*t)
	*t += ITally32(i)

	return
}

// Inc increases this counter by 1, returning the previous int32 value of
// this ITally32.
func (t *ITally32) Inc() (cur int32) {
	cur = int32(*t)
	*t++

	return
}

// Dec decreases this counter by 1, returning the previous int32 value of
// this ITally32.
func (t *ITally32) Dec() (cur int32) {
	cur = int32(*t)
	*t--

	return
}

// Sub decreases this counter by the given value, returning the previous
// int32 value of this ITally32.
func (t *ITally32) Sub(i int32) (cur int32) {
	cur = int32(*t)
	*t -= ITally32(i)

	return
}

// Zero sets the current counter value to 0, returning the previous int32
// value of this ITally32.
func (t *ITally32) Zero() (cur int32) {
	cur = int32(*t)
	*t = 0

	return
}

// ITally64 is a counter of type int64.
type ITally64 int64

// Cur returns the current int64 value of this ITally64.
func (t ITally64) Cur() int64 {
	return int64(t)
}

// Add increases this counter by the given value, returning the previous
// int64 value of this ITally64.
func (t *ITally64) Add(i int64) (cur int64) {
	cur = int64(*t)
	*t += ITally64(i)

	return
}

// Inc increases this counter by 1, returning the previous int64 value of
// this ITally64.
func (t *ITally64) Inc() (cur int64) {
	cur = int64(*t)
	*t++

	return
}

// Dec decreases this counter by 1, returning the previous int64 value of
// this ITally64.
func (t *ITally64) Dec() (cur int64) {
	cur = int64(*t)
	*t--

	return
}

// Sub decreases this counter by the given value, returning the previous
// int64 value of this ITally64.
func (t *ITally64) Sub(i int64) (cur int64) {
	cur = int64(*t)
	*t -= ITally64(i)

	return
}

// Zero sets the current counter value to 0, returning the previous int64
// value of this ITally64.
func (t *ITally64) Zero() (cur int64) {
	cur = int64(*t)
	*t = 0

	return
}