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
	riscv_i is the base integer instructions handler
*/
package riscv_i

import (
	"errors"
	"riscv/riscv_core/riscv_reg"
)

var instr_enable bool

func Init(enable bool) {
	instr_enable = enable
}

func Instr_handle(instr uint, reg *riscv_reg.Regfile) error {
	if false == instr_enable {
		return errors.New("RVI is not supported.")
	}
	return nil
}
