package main

import "fmt"

type CompiledCode struct {
  Name string
  Regs int
  Opcodes []int
}

type CallFrame struct {
  Ip int
  Regs []int
}

type Instruction struct {
  Name string
  Size int
  Opcode int
  Parameters int
  Interpreter func(state *State, cf *CallFrame, opcodes []int) int
}

type Instructions map[int]Instruction

type State struct {
  insns Instructions
}

func next_insn(state *State, cf *CallFrame, opcodes []int) Instruction {
  return state.insns[opcodes[cf.Ip]]
}

func arg(cf *CallFrame, arg int, opcodes []int) int {
  return opcodes[cf.Ip+arg]
}

func load_int(state *State, cf *CallFrame, opcodes []int) int {
  cf.Regs[arg(cf, 0, opcodes)] = arg(cf, 1, opcodes)

  cf.Ip += 3
  return next_insn(state, cf, opcodes).Interpreter(state, cf, opcodes)
}

// func jgte(state *State, cf *CallFrame, opcodes []int) int {
// }

func ret(state *State, cf *CallFrame, opcodes []int) int {
  return 0
}

func interpreter(state *State, opcodes []int) int {
  regs := make([]int, 3)
  cf := CallFrame{ 0, regs }

  return next_insn(state, &cf, opcodes).Interpreter(state, &cf, opcodes)
}

func main() {
  opcodes := []int {
    0, 0, 3,        // load_int reg0 3
    1,              // ret
  }

  state := State {
    Instructions {
      0: { "load_int", 3, 0, 2, load_int },
      1: { "ret", 1, 1, 0, ret },
    },
  }

  fmt.Printf("result: %v\n", interpreter(&state, opcodes))
}
