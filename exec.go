package main

import (
	"fmt"
)

type Executor struct {
	ast     *AST
	mem     []int
	pointer int64
}

func NewExecutor(ast *AST) *Executor {
	return &Executor{
		mem: make([]int, 3000),
		ast: ast,
	}
}

func (e *Executor) exec(n Node) {
	switch b := n.(type) {
	case CommandForward:
		e.pointer++
	case CommandBackward:
		e.pointer--
	case CommandAdd:
		e.mem[e.pointer]++
	case CommandSub:
		e.mem[e.pointer]--
	case CommandOutput:
		fmt.Printf(string([]byte{byte(e.mem[e.pointer])}))
	case CommandInput:
	// TODO: impl
	case CommandJump:
		e.execNodes(b.Nodes)
	}
}

func (e *Executor) execNodes(nodes []Node) {
	for e.mem[e.pointer] > 0 {
		for _, node := range nodes {
			e.exec(node)
		}
	}
}

func (e *Executor) run() {
	for _, node := range e.ast.nodes {
		e.exec(node)
	}
}
