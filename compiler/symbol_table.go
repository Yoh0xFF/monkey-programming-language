package compiler

type SymbolScope string

const (
	LocalScope  SymbolScope = "LOCAL"
	GlobalScope SymbolScope = "GLOBAL"
)

type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
}

type SymbolTable struct {
	Outer *SymbolTable

	store   map[string]Symbol
	counter int
}

func NewEnclosedSymbolTable(outer *SymbolTable) *SymbolTable {
	s := NewSymbolTable()
	s.Outer = outer
	return s
}

func NewSymbolTable() *SymbolTable {
	store := make(map[string]Symbol)
	return &SymbolTable{store: store}
}

func (s *SymbolTable) Define(name string) Symbol {
	symbol := Symbol{Name: name, Index: s.counter}

	if s.Outer == nil {
		symbol.Scope = GlobalScope
	} else {
		symbol.Scope = LocalScope
	}

	s.store[name] = symbol
	s.counter++
	return symbol
}

func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	symbol, ok := s.store[name]

	if !ok && s.Outer != nil {
		symbol, ok = s.Outer.Resolve(name)
		return symbol, ok
	}

	return symbol, ok
}
