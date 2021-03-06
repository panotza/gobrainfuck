package main

type AST struct {
	nodes []Node
}

type Node interface {
	node()
}

type CommandForward struct{}

func (c CommandForward) node() {}

type CommandBackward struct{}

func (c CommandBackward) node() {}

type CommandAdd struct{}

func (c CommandAdd) node() {}

type CommandSub struct{}

func (c CommandSub) node() {}

type CommandOutput struct{}

func (c CommandOutput) node() {}

type CommandInput struct{}

func (c CommandInput) node() {}

type CommandJump struct {
	Nodes []Node
}

func (c CommandJump) node() {}
