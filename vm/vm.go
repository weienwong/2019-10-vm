package vm

type VM struct {
	Regs   [4]byte
	IRegs  [3]byte
	Memory [256]byte
}

func (vm *VM) Load(bin []byte) {
	if len(bin) > 256 {
		panic("binary too big")
	}

	copy(vm.Memory[:], bin)

	vm.PC = vm.Memory[0]
	vm.Memory[0] = 0
	vm.SP = len(vm.Memory) - 1
	vm.FP = len(vm.Memory) - 1
}

func (vm *VM) Execute() error {
	return nil
}
