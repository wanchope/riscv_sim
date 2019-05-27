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
	riscv_core is the top level of the risc-v core.
*/
package riscv_core

import (
	//"fmt"
	"riscv/riscv_core/riscv_i"
	"riscv/riscv_core/riscv_reg"
)

var regfile *riscv_reg.Regfile

func Init(option string) {
	regfile = riscv_reg.Init(option)
	riscv_i_enable := false
	for _, c := range option {
		switch c {
		case 'I':
			fallthrough
		case 'i':
			riscv_i_enable = true
		}
	}
	riscv_i.Init(riscv_i_enable)
}
