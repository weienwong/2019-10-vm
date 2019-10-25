package virtual

import (
	"fmt"

	"github.com/weienwong/2019-10-vm/is"
)

const (
	NumRegs      = 4
	InternalRegs = 3
	MemorySize   = 255

	R0 = 0
	R1 = 1
	R2 = 2
	R3 = 3
)

type Machine struct {
	Regs   [NumRegs]byte
	IRegs  [InternalRegs]byte
	Memory [MemorySize]byte

	PC     byte
	SP, FP byte
}

func (vm *Machine) Load(bin []byte) {
	if len(bin) > MemorySize {
		panic("binary too big")
	}

	copy(vm.Memory[:], bin)

	vm.PC = vm.Memory[0]
	vm.Memory[0] = 0
	vm.SP = MemorySize - 1
	vm.FP = MemorySize - 1
}

func (vm *Machine) IncrementSP() byte {
	b := vm.Memory[vm.SP]
	vm.SP++
	if vm.SP >= MemorySize {
		vm.SP = MemorySize - 1
	}
	return b
}

func (vm *Machine) DecrementSP() byte {
	b := vm.Memory[vm.SP]
	vm.SP--
	if vm.SP <= 0 {
		vm.SP = 0
	}
	return b
}

func (vm *Machine) IncrementPC(amount uint) error {
	vm.PC += byte(amount)

	if vm.PC >= MemorySize {
		return fmt.Errorf("out of memory")
	}

	return nil
}

func (vm *Machine) Execute() (int, error) {

	var inst is.Instruction

	for {
		if vm.PC >= MemorySize {
			return -1, fmt.Errorf("out of memory")
		}

		inst = is.Instruction(vm.Memory[vm.PC])

		fmt.Printf("opt: %v\n", inst.Op())

		switch inst.Op() {
		case is.LDR:
			vm.Regs[inst.R1()] = vm.Memory[vm.Regs[inst.R2()]]
		case is.STR:
			vm.Memory[vm.Regs[inst.R1()]] = vm.Regs[inst.R2()]
		case is.LDI:
			vm.Regs[R0] = inst.Imm()
		case is.MOV:
			vm.Regs[inst.R1()] = vm.Regs[inst.R2()]
		case is.POP:
			// always POP data from SP into R1
			vm.Regs[inst.R1()] = vm.DecrementSP()
		case is.PSH:
			vm.Memory[vm.SP] = vm.Regs[inst.R1()]
			vm.IncrementSP()
		case is.BNZ:
			if inst.R2() != 0 {
				vm.PC = byte(vm.Regs[inst.R2()])
				continue
			}
		case is.CAL:
			if vm.Regs[inst.R2()] != 0 {
				vm.PC = byte(vm.Memory[vm.FP])
				vm.SP = vm.FP + 2
				vm.FP = vm.Memory[vm.FP+1]
				continue
			}
			if vm.Regs[inst.R1()] != 0 {
				break
			}

			vm.DecrementSP()
			vm.Memory[vm.SP] = vm.FP
			vm.DecrementSP()
			vm.Memory[vm.SP] = vm.PC + 1

			vm.FP = vm.SP
			vm.PC = byte(inst.R1())

			continue
		case is.ADR:
		case is.ADD:
		case is.SHR:
		case is.AND:
		case is.ORR:
		case is.EOR:
		case is.OST:
		case is.IST:

		}

		vm.PC++

	}

	return int(vm.Memory[0]), nil
}
