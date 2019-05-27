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

func Test_Rformat_0(t *testing.T) {
	rs1, rs2, rd, funct3, funct7 := Rformat(test_instr)
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

func Test_Iformat_0(t *testing.T) {
	rs1, rd, funct3, imm := Iformat(test_instr)
	Convey("rs1 decode from instruction", t, func() {
		So(rs1, ShouldEqual, test_rs1)
	})
	Convey("rd decode from instruction", t, func() {
		So(rd, ShouldEqual, test_rd)
	})
	Convey("funct3 decode from instruction", t, func() {
		So(funct3, ShouldEqual, test_funct3)
	})
	Convey("imm decode from instruction", t, func() {
		So(imm, ShouldEqual, test_funct7<<5 | test_rs2)
	})
}

func Test_Sformat_0(t *testing.T) {
	rs1, rs2, funct3, imm := Sformat(test_instr)
	Convey("rs1 decode from instruction", t, func() {
		So(rs1, ShouldEqual, test_rs1)
	})
	Convey("rs2 decode from instruction", t, func() {
		So(rs2, ShouldEqual, test_rs2)
	})
	Convey("funct3 decode from instruction", t, func() {
		So(funct3, ShouldEqual, test_funct3)
	})
	Convey("imm decode from instruction", t, func() {
		So(imm, ShouldEqual, test_funct7<<5 | test_rd)
	})
}

func Test_Bformat_0(t *testing.T) {
	rs1, rs2, funct3, imm := Bformat(test_instr)
	Convey("rs1 decode from instruction", t, func() {
		So(rs1, ShouldEqual, test_rs1)
	})
	Convey("rs2 decode from instruction", t, func() {
		So(rs2, ShouldEqual, test_rs2)
	})
	Convey("funct3 decode from instruction", t, func() {
		So(funct3, ShouldEqual, test_funct3)
	})
	Convey("imm decode from instruction", t, func() {
		So(imm, ShouldEqual, (test_funct7&0x3f)<<5 | test_rd&0x1e | (test_rd&1 << 11) | (test_funct7>>6 << 12))
	})
}

func Test_Uformat_0(t *testing.T) {
	rd, imm := Uformat(test_instr)
	Convey("rd decode from instruction", t, func() {
		So(rd, ShouldEqual, test_rd)
	})
	Convey("imm decode from instruction", t, func() {
		So(imm, ShouldEqual, (test_funct3 | test_rs1<<3 | test_rs2<<8 | test_funct7<<13)<<12 )
	})
}

func Test_Jformat_0(t *testing.T) {
	rd, imm := Jformat(test_instr)
	Convey("rd decode from instruction", t, func() {
		So(rd, ShouldEqual, test_rd)
	})
	Convey("imm decode from instruction", t, func() {
		So(imm, ShouldEqual, test_rs2&0x1e | (test_funct7&0x3f << 5) | (test_rs2&1 << 11) | test_funct3<<12 | test_rs1<<15 | (test_funct7>>6 << 20))
	})
}

