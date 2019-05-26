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
