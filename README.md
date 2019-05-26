# riscv_sim
Simulate the RISC-V core to run with the binary file.

#What to do?

This project will support following functions:
1. Recognize the instruction, identify the illegal instruction.
2. Easy to support new instructions and user defined instructions.
3. Run as a RISC-V based core.
4. Run with some sample peripherals, such as UART and interrupt.
5. Build a unit test framework to test the software with the binary file compiled by RISC-V compiler.

#Why use GO langue?
1. Use goroutine could easily simulate the interrupt and peripherals.
2. The GO langue is very similar with the C langue, easy to add new instructions by C programmers.
3. Easy to do the uint test.

#Unit test
Use GoConvey to support Go's native testing package.
get GoConvey as follow:
	go get github.com/smartystreets/goconvey
run the goconvey.exe in this directory, then open browser to localhost:8080, tests will be run and
we can get the coverage.

