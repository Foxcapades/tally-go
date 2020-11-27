package tally_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/foxcapades/tally-go/v1/tally"
)

func TestUTally_Cur(t *testing.T) {
	Convey("UTally.Cur", t, func() {
		Convey("Should return the current tally value", func() {
			test := tally.UTally(13)
			So(test.Cur(), ShouldEqual, 13)
			test = 35
			So(test.Cur(), ShouldEqual, 35)
		})
	})
}

func TestUTally_Add(t *testing.T) {
	Convey("UTally.Add", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally(13)
			So(test.Add(10), ShouldEqual, 13)
		})

		Convey("Should increment the tally counter by the given value", func() {
			test := tally.UTally(13)
			test.Add(10)
			So(test, ShouldEqual, 23)
		})
	})
}

func TestUTally_Inc(t *testing.T) {
	Convey("UTally.Inc", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally(13)
			So(test.Inc(), ShouldEqual, 13)
		})

		Convey("Should increment the tally counter by 1", func() {
			test := tally.UTally(13)
			test.Inc()
			So(test, ShouldEqual, 14)
		})
	})
}

func TestUTally_Dec(t *testing.T) {
	Convey("UTally.Dec", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally(13)
			So(test.Dec(), ShouldEqual, 13)
		})

		Convey("Should decrement the tally counter by 1", func() {
			test := tally.UTally(13)
			test.Dec()
			So(test, ShouldEqual, 12)
		})
	})
}

func TestUTally_Sub(t *testing.T) {
	Convey("UTally.Sub", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally(13)
			So(test.Sub(10), ShouldEqual, 13)
		})

		Convey("Should decrement the tally counter by 1", func() {
			test := tally.UTally(13)
			test.Sub(10)
			So(test, ShouldEqual, 3)
		})
	})
}

func TestUTally_Zero(t *testing.T) {
	Convey("UTally.Zero", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally(13)
			So(test.Zero(), ShouldEqual, 13)
		})

		Convey("Should set the tally value to 0", func() {
			test := tally.UTally(13)
			test.Zero()
			So(test, ShouldEqual, 0)
		})
	})
}

func TestUTally8_Cur(t *testing.T) {
	Convey("UTally8.Cur", t, func() {
		Convey("Should return the current tally value", func() {
			test := tally.UTally8(13)
			So(test.Cur(), ShouldEqual, 13)
			test = 35
			So(test.Cur(), ShouldEqual, 35)
		})
	})
}

func TestUTally8_Add(t *testing.T) {
	Convey("UTally8.Add", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally8(13)
			So(test.Add(10), ShouldEqual, 13)
		})

		Convey("Should increment the tally counter by the given value", func() {
			test := tally.UTally8(13)
			test.Add(10)
			So(test, ShouldEqual, 23)
		})
	})
}

func TestUTally8_Inc(t *testing.T) {
	Convey("UTally8.Inc", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally8(13)
			So(test.Inc(), ShouldEqual, 13)
		})

		Convey("Should increment the tally counter by 1", func() {
			test := tally.UTally8(13)
			test.Inc()
			So(test, ShouldEqual, 14)
		})
	})
}

func TestUTally8_Dec(t *testing.T) {
	Convey("UTally8.Dec", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally8(13)
			So(test.Dec(), ShouldEqual, 13)
		})

		Convey("Should decrement the tally counter by 1", func() {
			test := tally.UTally8(13)
			test.Dec()
			So(test, ShouldEqual, 12)
		})
	})
}

func TestUTally8_Sub(t *testing.T) {
	Convey("UTally8.Sub", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally8(13)
			So(test.Sub(10), ShouldEqual, 13)
		})

		Convey("Should decrement the tally counter by 1", func() {
			test := tally.UTally8(13)
			test.Sub(10)
			So(test, ShouldEqual, 3)
		})
	})
}

func TestUTally8_Zero(t *testing.T) {
	Convey("UTally8.Zero", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally8(13)
			So(test.Zero(), ShouldEqual, 13)
		})

		Convey("Should set the tally value to 0", func() {
			test := tally.UTally8(13)
			test.Zero()
			So(test, ShouldEqual, 0)
		})
	})
}

func TestUTally16_Cur(t *testing.T) {
	Convey("UTally16.Cur", t, func() {
		Convey("Should return the current tally value", func() {
			test := tally.UTally16(13)
			So(test.Cur(), ShouldEqual, 13)
			test = 35
			So(test.Cur(), ShouldEqual, 35)
		})
	})
}

func TestUTally16_Add(t *testing.T) {
	Convey("UTally16.Add", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally16(13)
			So(test.Add(10), ShouldEqual, 13)
		})

		Convey("Should increment the tally counter by the given value", func() {
			test := tally.UTally16(13)
			test.Add(10)
			So(test, ShouldEqual, 23)
		})
	})
}

func TestUTally16_Inc(t *testing.T) {
	Convey("UTally16.Inc", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally16(13)
			So(test.Inc(), ShouldEqual, 13)
		})

		Convey("Should increment the tally counter by 1", func() {
			test := tally.UTally16(13)
			test.Inc()
			So(test, ShouldEqual, 14)
		})
	})
}

func TestUTally16_Dec(t *testing.T) {
	Convey("UTally16.Dec", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally16(13)
			So(test.Dec(), ShouldEqual, 13)
		})

		Convey("Should decrement the tally counter by 1", func() {
			test := tally.UTally16(13)
			test.Dec()
			So(test, ShouldEqual, 12)
		})
	})
}

func TestUTally16_Sub(t *testing.T) {
	Convey("UTally16.Sub", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally16(13)
			So(test.Sub(10), ShouldEqual, 13)
		})

		Convey("Should decrement the tally counter by 1", func() {
			test := tally.UTally16(13)
			test.Sub(10)
			So(test, ShouldEqual, 3)
		})
	})
}

func TestUTally16_Zero(t *testing.T) {
	Convey("UTally16.Zero", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally16(13)
			So(test.Zero(), ShouldEqual, 13)
		})

		Convey("Should set the tally value to 0", func() {
			test := tally.UTally16(13)
			test.Zero()
			So(test, ShouldEqual, 0)
		})
	})
}

func TestUTally32_Cur(t *testing.T) {
	Convey("UTally32.Cur", t, func() {
		Convey("Should return the current tally value", func() {
			test := tally.UTally32(13)
			So(test.Cur(), ShouldEqual, 13)
			test = 35
			So(test.Cur(), ShouldEqual, 35)
		})
	})
}

func TestUTally32_Add(t *testing.T) {
	Convey("UTally32.Add", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally32(13)
			So(test.Add(10), ShouldEqual, 13)
		})

		Convey("Should increment the tally counter by the given value", func() {
			test := tally.UTally32(13)
			test.Add(10)
			So(test, ShouldEqual, 23)
		})
	})
}

func TestUTally32_Inc(t *testing.T) {
	Convey("UTally32.Inc", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally32(13)
			So(test.Inc(), ShouldEqual, 13)
		})

		Convey("Should increment the tally counter by 1", func() {
			test := tally.UTally32(13)
			test.Inc()
			So(test, ShouldEqual, 14)
		})
	})
}

func TestUTally32_Dec(t *testing.T) {
	Convey("UTally32.Dec", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally32(13)
			So(test.Dec(), ShouldEqual, 13)
		})

		Convey("Should decrement the tally counter by 1", func() {
			test := tally.UTally32(13)
			test.Dec()
			So(test, ShouldEqual, 12)
		})
	})
}

func TestUTally32_Sub(t *testing.T) {
	Convey("UTally32.Sub", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally32(13)
			So(test.Sub(10), ShouldEqual, 13)
		})

		Convey("Should decrement the tally counter by 1", func() {
			test := tally.UTally32(13)
			test.Sub(10)
			So(test, ShouldEqual, 3)
		})
	})
}

func TestUTally32_Zero(t *testing.T) {
	Convey("UTally32.Zero", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally32(13)
			So(test.Zero(), ShouldEqual, 13)
		})

		Convey("Should set the tally value to 0", func() {
			test := tally.UTally32(13)
			test.Zero()
			So(test, ShouldEqual, 0)
		})
	})
}

func TestUTally64_Cur(t *testing.T) {
	Convey("UTally64.Cur", t, func() {
		Convey("Should return the current tally value", func() {
			test := tally.UTally64(13)
			So(test.Cur(), ShouldEqual, 13)
			test = 35
			So(test.Cur(), ShouldEqual, 35)
		})
	})
}

func TestUTally64_Add(t *testing.T) {
	Convey("UTally64.Add", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally64(13)
			So(test.Add(10), ShouldEqual, 13)
		})

		Convey("Should increment the tally counter by the given value", func() {
			test := tally.UTally64(13)
			test.Add(10)
			So(test, ShouldEqual, 23)
		})
	})
}

func TestUTally64_Inc(t *testing.T) {
	Convey("UTally64.Inc", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally64(13)
			So(test.Inc(), ShouldEqual, 13)
		})

		Convey("Should increment the tally counter by 1", func() {
			test := tally.UTally64(13)
			test.Inc()
			So(test, ShouldEqual, 14)
		})
	})
}

func TestUTally64_Dec(t *testing.T) {
	Convey("UTally64.Dec", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally64(13)
			So(test.Dec(), ShouldEqual, 13)
		})

		Convey("Should decrement the tally counter by 1", func() {
			test := tally.UTally64(13)
			test.Dec()
			So(test, ShouldEqual, 12)
		})
	})
}

func TestUTally64_Sub(t *testing.T) {
	Convey("UTally64.Sub", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally64(13)
			So(test.Sub(10), ShouldEqual, 13)
		})

		Convey("Should decrement the tally counter by 1", func() {
			test := tally.UTally64(13)
			test.Sub(10)
			So(test, ShouldEqual, 3)
		})
	})
}

func TestUTally64_Zero(t *testing.T) {
	Convey("UTally64.Zero", t, func() {
		Convey("Should return the previous tally value on call", func() {
			test := tally.UTally64(13)
			So(test.Zero(), ShouldEqual, 13)
		})

		Convey("Should set the tally value to 0", func() {
			test := tally.UTally64(13)
			test.Zero()
			So(test, ShouldEqual, 0)
		})
	})
}