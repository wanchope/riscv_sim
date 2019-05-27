/*
	Copyright 2019 Wencheng Li

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/
/*
	riscv_decode is used to decode instruction as the RISBUJ format,
	this will be imported by every instruction family.
*/
package riscv_decode

import ()

func bit(n byte) uint {
	return 1<<n
}

func common_decode(instr uint) (rs1, rs2, rd, funct3, funct7 byte) {
	rd = byte((instr >> 7) & 0x1f)
	rs1 = byte((instr >> 15) & 0x1f)
	rs2 = byte((instr >> 20) & 0x1f)
	funct3 = byte((instr >> 12) & 0x07)
	funct7 = byte((instr >> 25) & 0x7f)
	return
}

func Rformat(instr uint) (rs1, rs2, rd, funct3, funct7 byte) {
	rs1, rs2, rd, funct3, funct7 = common_decode(instr)
	return
}

func Iformat(instr uint) (rs1, rd, funct3 byte, imm uint) {
	rs1, _, rd, funct3, _ = common_decode(instr)
	imm = instr >> 20
	return
}

func Sformat(instr uint) (rs1, rs2, funct3 byte, imm uint) {
	rs1, rs2, _, funct3, _ = common_decode(instr)
	imm = (instr>>25 << 5) | (instr>>7 & 0x1f)
	return
}

func Bformat(instr uint) (rs1, rs2, funct3 byte, imm uint) {
	rs1, rs2, _, funct3, _ = common_decode(instr)
	imm = ((instr>>25 & 0x3f)<<5) | (instr>>7 & 0x1e) | (instr&bit(31)>>(31-12)) | (instr&bit(7)<<(11-7))
	return
}

func Uformat(instr uint) (rd byte, imm uint) {
	rd = byte(instr>>7 & 0x1f)
	imm = instr & 0xfffff000
	return
}

func Jformat(instr uint) (rd byte, imm uint) {
	ins := instr >> 12
	rd = byte(instr>>7 & 0x1f)
	imm = (ins&0xff << 12) | (ins&0x100 << 3) | (ins&0x7fe00 >> 8) | (ins>>19 << 20)
	return
}

