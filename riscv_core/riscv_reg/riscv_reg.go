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
	riscv_reg is the regfile of risc-v core
*/
package riscv_reg

type Regfile struct {
	x [32]uint    //x0~x31
	f [32]float32 //ft8~ft11
}

const RegDisable = 0xDEADBEAF

var regfile Regfile

func Init(option string) *Regfile {
	for n, _ := range regfile.x {
		regfile.x[n] = 0xfedcba98
	}
	for n, _ := range regfile.f {
		regfile.f[n] = float32(0x76543210)
	}
	for _, c := range option {
		switch c {
		case 'f':
			fallthrough
		case 'F':
			fallthrough
		case 'd':
			fallthrough
		case 'D':
			for n, _ := range regfile.f {
				regfile.f[n] = float32(RegDisable)
			}
		}
	}
	return &regfile
}
