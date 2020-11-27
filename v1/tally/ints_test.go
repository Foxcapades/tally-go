package tally_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/foxcapades/tally-go/v1/tally"
)

func TestITally_Cur(t *testing.T) {
	Convey("ITally.Cur", t, func() {
		Convey("Should return the current tally value", func() {
			test := tally.ITally(13)
			So(test.Cur(), ShouldEqual, 13)
			test = 35
			So(test.Cur(), ShouldEqual, 35)
		})
	})
}

func TestITally_Add(t *testing.T) {
	Convey("ITally.Add", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally(13)
			So(test.Add(10), ShouldEqual, 13)
		})

		Convey("Should increment the tally counter by the given value", func() {
			test := tally.ITally(13)
			test.Add(10)
			So(test, ShouldEqual, 23)
		})
	})
}

func TestITally_Inc(t *testing.T) {
	Convey("ITally.Inc", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally(13)
			So(test.Inc(), ShouldEqual, 13)
		})

		Convey("Should increment the tally counter by 1", func() {
			test := tally.ITally(13)
			test.Inc()
			So(test, ShouldEqual, 14)
		})
	})
}

func TestITally_Dec(t *testing.T) {
	Convey("ITally.Dec", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally(13)
			So(test.Dec(), ShouldEqual, 13)
		})

		Convey("Should decrement the tally counter by 1", func() {
			test := tally.ITally(13)
			test.Dec()
			So(test, ShouldEqual, 12)
		})
	})
}

func TestITally_Sub(t *testing.T) {
	Convey("ITally.Sub", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally(13)
			So(test.Sub(10), ShouldEqual, 13)
		})

		Convey("Should decrement the tally counter by 1", func() {
			test := tally.ITally(13)
			test.Sub(10)
			So(test, ShouldEqual, 3)
		})
	})
}

func TestITally_Zero(t *testing.T) {
	Convey("ITally.Zero", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally(13)
			So(test.Zero(), ShouldEqual, 13)
		})

		Convey("Should set the tally value to 0", func() {
			test := tally.ITally(13)
			test.Zero()
			So(test, ShouldEqual, 0)
		})
	})
}

func TestITally8_Cur(t *testing.T) {
	Convey("ITally8.Cur", t, func() {
		Convey("Should return the current tally value", func() {
			test := tally.ITally8(13)
			So(test.Cur(), ShouldEqual, 13)
			test = 35
			So(test.Cur(), ShouldEqual, 35)
		})
	})
}

func TestITally8_Add(t *testing.T) {
	Convey("ITally8.Add", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally8(13)
			So(test.Add(10), ShouldEqual, 13)
		})

		Convey("Should increment the tally counter by the given value", func() {
			test := tally.ITally8(13)
			test.Add(10)
			So(test, ShouldEqual, 23)
		})
	})
}

func TestITally8_Inc(t *testing.T) {
	Convey("ITally8.Inc", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally8(13)
			So(test.Inc(), ShouldEqual, 13)
		})

		Convey("Should increment the tally counter by 1", func() {
			test := tally.ITally8(13)
			test.Inc()
			So(test, ShouldEqual, 14)
		})
	})
}

func TestITally8_Dec(t *testing.T) {
	Convey("ITally8.Dec", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally8(13)
			So(test.Dec(), ShouldEqual, 13)
		})

		Convey("Should decrement the tally counter by 1", func() {
			test := tally.ITally8(13)
			test.Dec()
			So(test, ShouldEqual, 12)
		})
	})
}

func TestITally8_Sub(t *testing.T) {
	Convey("ITally8.Sub", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally8(13)
			So(test.Sub(10), ShouldEqual, 13)
		})

		Convey("Should decrement the tally counter by 1", func() {
			test := tally.ITally8(13)
			test.Sub(10)
			So(test, ShouldEqual, 3)
		})
	})
}

func TestITally8_Zero(t *testing.T) {
	Convey("ITally8.Zero", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally8(13)
			So(test.Zero(), ShouldEqual, 13)
		})

		Convey("Should set the tally value to 0", func() {
			test := tally.ITally8(13)
			test.Zero()
			So(test, ShouldEqual, 0)
		})
	})
}

func TestITally16_Cur(t *testing.T) {
	Convey("ITally16.Cur", t, func() {
		Convey("Should return the current tally value", func() {
			test := tally.ITally16(13)
			So(test.Cur(), ShouldEqual, 13)
			test = 35
			So(test.Cur(), ShouldEqual, 35)
		})
	})
}

func TestITally16_Add(t *testing.T) {
	Convey("ITally16.Add", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally16(13)
			So(test.Add(10), ShouldEqual, 13)
		})

		Convey("Should increment the tally counter by the given value", func() {
			test := tally.ITally16(13)
			test.Add(10)
			So(test, ShouldEqual, 23)
		})
	})
}

func TestITally16_Inc(t *testing.T) {
	Convey("ITally16.Inc", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally16(13)
			So(test.Inc(), ShouldEqual, 13)
		})

		Convey("Should increment the tally counter by 1", func() {
			test := tally.ITally16(13)
			test.Inc()
			So(test, ShouldEqual, 14)
		})
	})
}

func TestITally16_Dec(t *testing.T) {
	Convey("ITally16.Dec", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally16(13)
			So(test.Dec(), ShouldEqual, 13)
		})

		Convey("Should decrement the tally counter by 1", func() {
			test := tally.ITally16(13)
			test.Dec()
			So(test, ShouldEqual, 12)
		})
	})
}

func TestITally16_Sub(t *testing.T) {
	Convey("ITally16.Sub", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally16(13)
			So(test.Sub(10), ShouldEqual, 13)
		})

		Convey("Should decrement the tally counter by 1", func() {
			test := tally.ITally16(13)
			test.Sub(10)
			So(test, ShouldEqual, 3)
		})
	})
}

func TestITally16_Zero(t *testing.T) {
	Convey("ITally16.Zero", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally16(13)
			So(test.Zero(), ShouldEqual, 13)
		})

		Convey("Should set the tally value to 0", func() {
			test := tally.ITally16(13)
			test.Zero()
			So(test, ShouldEqual, 0)
		})
	})
}

func TestITally32_Cur(t *testing.T) {
	Convey("ITally32.Cur", t, func() {
		Convey("Should return the current tally value", func() {
			test := tally.ITally32(13)
			So(test.Cur(), ShouldEqual, 13)
			test = 35
			So(test.Cur(), ShouldEqual, 35)
		})
	})
}

func TestITally32_Add(t *testing.T) {
	Convey("ITally32.Add", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally32(13)
			So(test.Add(10), ShouldEqual, 13)
		})

		Convey("Should increment the tally counter by the given value", func() {
			test := tally.ITally32(13)
			test.Add(10)
			So(test, ShouldEqual, 23)
		})
	})
}

func TestITally32_Inc(t *testing.T) {
	Convey("ITally32.Inc", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally32(13)
			So(test.Inc(), ShouldEqual, 13)
		})

		Convey("Should increment the tally counter by 1", func() {
			test := tally.ITally32(13)
			test.Inc()
			So(test, ShouldEqual, 14)
		})
	})
}

func TestITally32_Dec(t *testing.T) {
	Convey("ITally32.Dec", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally32(13)
			So(test.Dec(), ShouldEqual, 13)
		})

		Convey("Should decrement the tally counter by 1", func() {
			test := tally.ITally32(13)
			test.Dec()
			So(test, ShouldEqual, 12)
		})
	})
}

func TestITally32_Sub(t *testing.T) {
	Convey("ITally32.Sub", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally32(13)
			So(test.Sub(10), ShouldEqual, 13)
		})

		Convey("Should decrement the tally counter by 1", func() {
			test := tally.ITally32(13)
			test.Sub(10)
			So(test, ShouldEqual, 3)
		})
	})
}

func TestITally32_Zero(t *testing.T) {
	Convey("ITally32.Zero", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally32(13)
			So(test.Zero(), ShouldEqual, 13)
		})

		Convey("Should set the tally value to 0", func() {
			test := tally.ITally32(13)
			test.Zero()
			So(test, ShouldEqual, 0)
		})
	})
}

func TestITally64_Cur(t *testing.T) {
	Convey("ITally64.Cur", t, func() {
		Convey("Should return the current tally value", func() {
			test := tally.ITally64(13)
			So(test.Cur(), ShouldEqual, 13)
			test = 35
			So(test.Cur(), ShouldEqual, 35)
		})
	})
}

func TestITally64_Add(t *testing.T) {
	Convey("ITally64.Add", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally64(13)
			So(test.Add(10), ShouldEqual, 13)
		})

		Convey("Should increment the tally counter by the given value", func() {
			test := tally.ITally64(13)
			test.Add(10)
			So(test, ShouldEqual, 23)
		})
	})
}

func TestITally64_Inc(t *testing.T) {
	Convey("ITally64.Inc", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally64(13)
			So(test.Inc(), ShouldEqual, 13)
		})

		Convey("Should increment the tally counter by 1", func() {
			test := tally.ITally64(13)
			test.Inc()
			So(test, ShouldEqual, 14)
		})
	})
}

func TestITally64_Dec(t *testing.T) {
	Convey("ITally64.Dec", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally64(13)
			So(test.Dec(), ShouldEqual, 13)
		})

		Convey("Should decrement the tally counter by 1", func() {
			test := tally.ITally64(13)
			test.Dec()
			So(test, ShouldEqual, 12)
		})
	})
}

func TestITally64_Sub(t *testing.T) {
	Convey("ITally64.Sub", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally64(13)
			So(test.Sub(10), ShouldEqual, 13)
		})

		Convey("Should decrement the tally counter by 1", func() {
			test := tally.ITally64(13)
			test.Sub(10)
			So(test, ShouldEqual, 3)
		})
	})
}

func TestITally64_Zero(t *testing.T) {
	Convey("ITally64.Zero", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.ITally64(13)
			So(test.Zero(), ShouldEqual, 13)
		})

		Convey("Should set the tally value to 0", func() {
			test := tally.ITally64(13)
			test.Zero()
			So(test, ShouldEqual, 0)
		})
	})
}