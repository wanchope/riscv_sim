package riscv_decode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const test_rs1 = 0x10
const test_rs2 = 0x12
const test_rd = 0x14
const test_funct3 = 4
const test_funct7 = 58

const test_instr = test_rd<<7 | test_funct3<<12 | test_rs1<<15 | test_funct7<<25 | test_rs2<<20

func Test_common_decode_0(t *testing.T) {
	rs1, rs2, rd, funct3, funct7 := common_decode(test_instr)
	Convey("rs1 decode from instruction", t, func() {
		So(rs1, ShouldEqual, test_rs1)
	})
	Convey("rs2 decode from instruction", t, func() {
		So(rs2, ShouldEqual, test_rs2)
	})
	Convey("rd decode from instruction", t, func() {
		So(rd, ShouldEqual, test_rd)
	})
	Convey("funct3 decode from instruction", t, func() {
		So(funct3, ShouldEqual, test_funct3)
	})
	Convey("funct7 decode from instruction", t, func() {
		So(funct7, ShouldEqual, test_funct7)
	})
}
