package tally

// UTally is a counter of type uint.
type UTally uint

// Cur returns the current uint value of this UTally.
func (t UTally) Cur() uint {
	return uint(t)
}

// Add increases this counter by the given value, returning the previous
// uint value of this UTally.
func (t *UTally) Add(i uint) (cur uint) {
	cur = uint(*t)
	*t += UTally(i)

	return
}

// Inc increases this counter by 1, returning the previous uint value of
// this UTally.
func (t *UTally) Inc() (cur uint) {
	cur = uint(*t)
	*t++

	return
}

// Dec decreases this counter by 1, returning the previous uint value of
// this UTally.
func (t *UTally) Dec() (cur uint) {
	cur = uint(*t)
	*t--

	return
}

// Sub decreases this counter by the given value, returning the previous
// uint value of this UTally.
func (t *UTally) Sub(i uint) (cur uint) {
	cur = uint(*t)
	*t -= UTally(i)

	return
}

// Zero sets the current counter value to 0, returning the previous uint
// value of this UTally.
func (t *UTally) Zero() (cur uint) {
	cur = uint(*t)
	*t = 0

	return
}

// UTally8 is a counter of type uint8.
type UTally8 uint8

// Cur returns the current uint8 value of this UTally8.
func (t UTally8) Cur() uint8 {
	return uint8(t)
}

// Add increases this counter by the given value, returning the previous
// uint8 value of this UTally8.
func (t *UTally8) Add(i uint8) (cur uint8) {
	cur = uint8(*t)
	*t += UTally8(i)

	return
}

// Inc increases this counter by 1, returning the previous uint8 value of
// this UTally8.
func (t *UTally8) Inc() (cur uint8) {
	cur = uint8(*t)
	*t++

	return
}

// Dec decreases this counter by 1, returning the previous uint8 value of
// this UTally8.
func (t *UTally8) Dec() (cur uint8) {
	cur = uint8(*t)
	*t--

	return
}

// Sub decreases this counter by the given value, returning the previous
// uint8 value of this UTally8.
func (t *UTally8) Sub(i uint8) (cur uint8) {
	cur = uint8(*t)
	*t -= UTally8(i)

	return
}

// Zero sets the current counter value to 0, returning the previous uint8
// value of this UTally8.
func (t *UTally8) Zero() (cur uint8) {
	cur = uint8(*t)
	*t = 0

	return
}

// UTally16 is a counter of type uint16.
type UTally16 uint16

// Cur returns the current uint16 value of this UTally16.
func (t UTally16) Cur() uint16 {
	return uint16(t)
}

// Add increases this counter by the given value, returning the previous
// uint16 value of this UTally16.
func (t *UTally16) Add(i uint16) (cur uint16) {
	cur = uint16(*t)
	*t += UTally16(i)

	return
}

// Inc increases this counter by 1, returning the previous uint16 value of
// this UTally16.
func (t *UTally16) Inc() (cur uint16) {
	cur = uint16(*t)
	*t++

	return
}

// Dec decreases this counter by 1, returning the previous uint16 value of
// this UTally16.
func (t *UTally16) Dec() (cur uint16) {
	cur = uint16(*t)
	*t--

	return
}

// Sub decreases this counter by the given value, returning the previous
// uint16 value of this UTally16.
func (t *UTally16) Sub(i uint16) (cur uint16) {
	cur = uint16(*t)
	*t -= UTally16(i)

	return
}

// Zero sets the current counter value to 0, returning the previous uint16
// value of this UTally16.
func (t *UTally16) Zero() (cur uint16) {
	cur = uint16(*t)
	*t = 0

	return
}

// UTally32 is a counter of type uint32.
type UTally32 uint32

// Cur returns the current uint32 value of this UTally32.
func (t UTally32) Cur() uint32 {
	return uint32(t)
}

// Add increases this counter by the given value, returning the previous
// uint32 value of this UTally32.
func (t *UTally32) Add(i uint32) (cur uint32) {
	cur = uint32(*t)
	*t += UTally32(i)

	return
}

// Inc increases this counter by 1, returning the previous uint32 value of
// this UTally32.
func (t *UTally32) Inc() (cur uint32) {
	cur = uint32(*t)
	*t++

	return
}

// Dec decreases this counter by 1, returning the previous uint32 value of
// this UTally32.
func (t *UTally32) Dec() (cur uint32) {
	cur = uint32(*t)
	*t--

	return
}

// Sub decreases this counter by the given value, returning the previous
// uint32 value of this UTally32.
func (t *UTally32) Sub(i uint32) (cur uint32) {
	cur = uint32(*t)
	*t -= UTally32(i)

	return
}

// Zero sets the current counter value to 0, returning the previous uint32
// value of this UTally32.
func (t *UTally32) Zero() (cur uint32) {
	cur = uint32(*t)
	*t = 0

	return
}

// UTally64 is a counter of type uint64.
type UTally64 uint64

// Cur returns the current uint64 value of this UTally64.
func (t UTally64) Cur() uint64 {
	return uint64(t)
}

// Add increases this counter by the given value, returning the previous
// uint64 value of this UTally64.
func (t *UTally64) Add(i uint64) (cur uint64) {
	cur = uint64(*t)
	*t += UTally64(i)

	return
}

// Inc increases this counter by 1, returning the previous uint64 value of
// this UTally64.
func (t *UTally64) Inc() (cur uint64) {
	cur = uint64(*t)
	*t++

	return
}

// Dec decreases this counter by 1, returning the previous uint64 value of
// this UTally64.
func (t *UTally64) Dec() (cur uint64) {
	cur = uint64(*t)
	*t--

	return
}

// Sub decreases this counter by the given value, returning the previous
// uint64 value of this UTally64.
func (t *UTally64) Sub(i uint64) (cur uint64) {
	cur = uint64(*t)
	*t -= UTally64(i)

	return
}

// Zero sets the current counter value to 0, returning the previous uint64
// value of this UTally64.
func (t *UTally64) Zero() (cur uint64) {
	cur = uint64(*t)
	*t = 0

	return
}