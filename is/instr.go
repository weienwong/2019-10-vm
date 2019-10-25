package is

type Instruction byte

func (instr Instruction) Op() Opcode {
	return Opcode(instr >> 4)
}

func (instr Instruction) R1() Reg {
	return Reg((instr >> 2) & 3)
}

func (instr Instruction) R2() Reg {
	return Reg(instr & 3)
}

func (instr Instruction) IReg() IReg {
	return IReg(instr.R2())
}

func (instr Instruction) Imm() byte {
	return byte(instr & 0b00001111)
}

/*
This could be helpful for debugging. Given a binary file it would return
the assmebly code used to create it.
func (instr Instr) String() string {

}
*/

func (instr Instruction) Encode() byte {
	// TODO
	return 0
}

func DecodeInstruction(b byte) Instruction {
	// TODO
	return 0
}

type Reg byte

const (
	R0 Reg = iota
	R1
	R2
	R3
)

type IReg byte

const (
	PC IReg = iota
	FP
	SP
)
