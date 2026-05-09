# Go Lang Practice

A collection of Go language learning projects covering fundamental concepts and practical implementations. This repository serves as a practice ground for mastering Go's core features and concurrency patterns.

## 📚 Projects Overview

### CardsProject
A practical project that demonstrates file I/O, slices, and type methods in Go.
- **Features:**
  - Deck creation and shuffling
  - Dealing cards to hands
  - Saving and reading deck data from files
  - Unit tests for deck operations
- **Key Concepts:** Receiver functions, slicing, file operations, testing

**Files:**
- `main.go` - Main program logic
- `deck.go` - Deck implementation
- `deck_test.go` - Unit tests
- `my_cards` - Sample saved deck file

### ChannelsAndGoRoutines
Demonstrates Go's concurrency model with goroutines and channels.
- **Features:**
  - HTTP link checking with goroutines
  - Channel-based communication between goroutines
  - Concurrent request handling
- **Key Concepts:** Goroutines, channels, concurrent programming, HTTP requests

**File:** `main.go`

### EvenOdd
A simple utility project for classifying numbers.
- **Features:**
  - Identify even and odd numbers
  - Process number collections
- **Key Concepts:** Structs, methods, slices

**Files:**
- `main.go` - Entry point
- `evenodd.go` - Even/odd logic implementation

### HTTP
HTTP client implementation demonstrating web requests and response handling.
- **Features:**
  - Making HTTP GET requests
  - Reading and processing response bodies
  - Custom writer implementation
  - Error handling for network operations
- **Key Concepts:** HTTP clients, interfaces (io.Writer), error handling

**File:** `main.go`

### Interface
Demonstrates Go's interface system with practical examples.
- **Features:**
  - Defining and implementing interfaces
  - Language-specific implementations (English, Spanish)
  - Polymorphic behavior through interfaces
- **Key Concepts:** Interface definition, type assertions, polymorphism

**File:** `main.go`

### InterfaceTestAreaOfShapes
Advanced interface usage for calculating shape areas.
- **Features:**
  - Shape interface with multiple implementations
  - Area calculation for different shape types
- **Key Concepts:** Interface contracts, polymorphic functions

**File:** `main.go`

### Map
Data structure exploration focusing on Go maps.
- **Features:**
  - Map creation and manipulation
  - Iteration over key-value pairs
  - Common map operations
- **Key Concepts:** Maps, iteration, data storage

**File:** `main.go`

### Struct
Foundation project exploring Go structs.
- **Features:**
  - Struct definition and instantiation
  - Field access and modification
  - Embedded structs
- **Key Concepts:** Struct definition, composition, value types

**File:** `main.go`

### Variable
Basic variable declarations and type system.
- **Features:**
  - Variable declaration syntax
  - Type inference
  - Constants and variables
- **Key Concepts:** Go type system, variable declaration, constants

**File:** `variable.go`

### ReadFile
File reading and text processing.
- **Features:**
  - Reading files from disk
  - Text processing and output
  - Error handling for file operations
- **Key Concepts:** File I/O, string manipulation, error handling

**Files:**
- `main.go` - File reading logic
- `myFile.txt` - Sample text file

## 🚀 Getting Started

### Prerequisites
- Go 1.16 or higher installed on your system
- Basic command line knowledge

### Running Projects

Navigate to any project directory and run:

```bash
cd ProjectName
go run main.go
```

For example:
```bash
cd CardsProject
go run main.go deck.go
```

### Running Tests

For projects with test files (like CardsProject):

```bash
cd CardsProject
go test
```

Or with verbose output:
```bash
go test -v
```

## 📁 Repository Structure

```
GoLang-Practice/
├── CardsProject/           # Card deck operations with file I/O
├── ChannelsAndGoRoutines/  # Concurrent programming examples
├── EvenOdd/                # Even/odd classification
├── HTTP/                   # HTTP client requests
├── Interface/              # Interface implementation
├── InterfaceTestAreaOfShapes/  # Advanced interface usage
├── Map/                    # Map data structure
├── ReadFile/               # File reading operations
├── Struct/                 # Struct definition and usage
├── Variable/               # Variable declarations
├── go.mod                  # Module definition
├── main.go                 # Root entry point
└── README.md              # This file
```

## 🎯 Learning Path

Recommended order to explore concepts:

1. **Variable** - Start with basic syntax and type system
2. **Struct** - Understand data structures
3. **Map** - Learn about key-value storage
4. **EvenOdd** - Practice with methods and structs
5. **Interface** - Understand polymorphism
6. **InterfaceTestAreaOfShapes** - Advanced interface patterns
7. **HTTP** - Real-world web operations
8. **ReadFile** - File I/O operations
9. **CardsProject** - Combine concepts in a practical project
10. **ChannelsAndGoRoutines** - Master concurrency

## 🔑 Key Go Concepts Covered

- **Variables & Types** - Declaration, inference, type system
- **Collections** - Slices, maps, arrays
- **Structs** - Composition, embedded structs, methods
- **Interfaces** - Interface definition, polymorphism, type assertions
- **Concurrency** - Goroutines, channels, select statements
- **File I/O** - Reading and writing files
- **HTTP** - Client requests, response handling
- **Testing** - Unit tests with Go's testing package
- **Error Handling** - Error interfaces and custom error types

## 💡 Tips for Learning

- Modify the examples to experiment with different approaches
- Add more test cases to deepen understanding
- Combine concepts from different projects
- Use `go fmt` to format code consistently
- Run `go vet` to catch common mistakes

## 📝 Notes

- Each project can be run independently
- Some projects may make external HTTP requests (ensure network connectivity)
- The CardsProject includes file operations that create/modify `my_cards` file

## 🤝 Contributing

Feel free to expand these examples with:
- Additional features
- More comprehensive tests
- Documentation improvements
- Performance optimizations

## 📚 Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Go Tour](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go by Example](https://gobyexample.com/)

---

**Last Updated:** May 2026

Happy learning! 🎉