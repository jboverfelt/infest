package main

// Store represents a store of Closures
// Store must be safe for use by concurrent goroutines
type Store interface {
	Insert(closure Closure) error
	GetAll() ([]*Closure, error)
}
