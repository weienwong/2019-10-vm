# VM

For this challange we're writing a virtual machine(vm) and assembler. The
virtual machine is not replicating any specific, existing, machine. The goal
is to learn how a computer works at a lower level. In the scope of Go, this
kind of project will show that it does have the capability of being a
"systems" language.


## The challange

The challanage is to implement a vm or an assembler for the architecture
written here.

## The architecture

The architecture describes features and attributes of the machine and its
components (memory, paging, etc).

* 8 bit memory addresses, instructions, and registers.
* 4 general purpose registers
* 2 specific registers that cannot be accessed by user code
    * `PC` (program counter) a pointer to the memory location of the
        currently executing instruction.
    * `SP` (stack pointer) a pointer to the memory location at the top of the
        stack. A caveat is that the stack grows from high to low memory so,
        the "top of the stack" actually has the lowest address.
* A program is copied into memory starting at address 0. The first byte
    (address 0) of a program is also a pointer to where execution should start.
    * 0 is not a valid pointer


## The instruction set

Now that we know roughly the architecture, we can describe the instructions
and their encoding. There are 5 classes of instructions: memory,
data processing, branching, stack, and IO.

Once again, the instructions are 8 bits. And for our instruction set,
the first fout bits of the isntruction denote the op code. The rest of the
bits are used to denote 1 or 2 registers, or an immediate constant.

| mnemonic | meaning | op code | r1 | r2 | imm | c code |
| - | - | - | - | - | - | - |
| LDR | load register         | 0  | dest | source | | `dest = *source`         |
| STR | store register        | 1  | dest | source | | `*dest = source`         |
| LDI | load immediate        | 2  |      |        | const | `r0 = const`       |
| MOV | move, register        | 3  | dest | source | | `dest = source`          |
| POP | pop from stack        | 4  | reg  | exit   | | `reg = *--SP`            |
| PSH | push to stack         | 5  | reg  |        | | `*SP++ = reg`            |
| BIZ | branch if zero        | 6  | pred | loc    | | `if (pred) { PC = loc }` |
| CAL | call function         | 7  | loc  | ex     | | `if (ex) { exit(loc) } else { *--SP = (PC = loc) + 1` } |
| RET | return from function  | 8  |      |        | | `PC = *SP++`             |
| ADD | add                   | 9  | dest | source | | `dest += source`         |
| SHR | shift right           | 10 | dest | source | | `dest <<= source`        |
| AND | bitwise and           | 11 | dest | source | | `dest &= source`         |
| ORR | bitwise or            | 12 | dest | source | | `dest \|= source`        |
| EOR | bitwise exclusive or  | 13 | dest | source | | `dest ^= source`         |
| OST | output string         | 14 | ptr  | len    | | `write(1, ptr, len)`     |
| IST | input string          | 15 | ptr  | len    | | `read(0, ptr, len)`      |

*note*: CAL

The data structure for an instruction is a byte and it has methods to extract
each of the possible parameters.

## The VM

The vm is a pretty simple structure. It contains the registers and memory
for the architecture. The vm type will also have methods
for `Load`-ing and then `Execute`-ing a binary. It is recommended that
evaluation of the intructions is done in a loop with the `Execute` method.

## The Assembly language

The assembly language should be familiar to anyone who's worked with assembly
before. It has the following syntax:
```text
(; [^\n]*)*
start (label)
(section (data | text)
((label:)? (directive arg (, arg)*)? (; [^\n]*))*
)+
```
In english:
* a comment starts with ';' and goes until the end of the line
* the file may start with 0 or more comment lines
* the first functional statement in the file must be `start` and a
    label with denotes the first instruction to execute
* the rest of the file is 1 or more `section` blocks
    * a section block must be either `text` or `data`
    * a section block is 1 or more `line`s
        * a `line` consists of a `label`, `directive`, or `comment` or
            any combination of the three in this order

One thing to note is that a `directive` does not necessarilly align with an
instruction. For example, data processing directives might be accepted by
the assembler with an immediate value argument for the assembler to then
generate `LDI` instructions.

