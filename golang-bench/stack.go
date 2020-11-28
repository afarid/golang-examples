package stack

import (
	"errors"
	"sync"
)

type Stack interface {
	Pop() (Stack, error)
	Push(elem interface{}) Stack
	Top() (interface{}, error)
}

var (
	TopEmptyStackError = errors.New("can't use top on empty stack")
	PopEmptyStackError = errors.New("can't pop top on empty stack")
)
type MutableStack struct {
	mu sync.RWMutex
	data []interface{}
}

func (m *MutableStack) Top() (interface{}, error){
  m.mu.Lock()
  defer m.mu.Unlock()
  if len(m.data) == 0 {
  	return nil, TopEmptyStackError
  }
  return m.data[0], nil
}

func (m *MutableStack) Push(elem interface{}) (Stack) {
	m.mu.Lock()
	defer  m.mu.Unlock()
	m.data = append(m.data, elem)
	return m
}

func (m *MutableStack) Pop() (Stack, error){
	m.mu.Lock()
	defer m.mu.Unlock()
	if len(m.data) == 0 {
		return nil, PopEmptyStackError
	}
	m.data = m.data[:len(m.data) -1]
	return m, nil
}

func NewMutableStack() *MutableStack {
	return &MutableStack{}
}