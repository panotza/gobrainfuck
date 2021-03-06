package main

import (
	"bufio"
	"io"
)

type Parser struct {
	scanner *Scanner
	ast     AST
}

func (p *Parser) parse(reader io.Reader) *AST {
	r := bufio.NewReader(reader)
	p.scanner = NewScanner(r)
	return &AST{nodes: p.parseNode()}
}

func (p *Parser) parseNode() []Node {
	var ns []Node
	for {
		tok := p.scanner.Scan()
		if tok == EOF {
			break
		}

		n := p.parseToken(tok)
		if n == nil {
			break
		}
		ns = append(ns, n)
	}
	return ns
}

func (p *Parser) parseToken(tok Token) Node {
	switch tok {
	case FORWARD:
		return CommandForward{}
	case BACKWARD:
		return CommandBackward{}
	case ADD:
		return CommandAdd{}
	case SUB:
		return CommandSub{}
	case OUTPUT:
		return CommandOutput{}
	case INPUT:
		return CommandInput{}
	case JMPL:
		return CommandJump{
			Nodes: p.parseNode(),
		}
	case JMPR:
		return nil
	}
	return nil
}
