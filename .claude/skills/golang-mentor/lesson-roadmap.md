# Lesson Progression Roadmap

This reference provides the structured progression of lessons. Load this when planning which lesson to teach next or when the user asks about the curriculum.

## Phase 1: Foundations (Lessons 1-5)

### Lesson 1: Hello World & Basic Syntax
- Package main, import, func main
- fmt.Println
- Running Go programs (go run)
**Exercise**: Print your favorite Pokemon and its level

### Lesson 2: Variables & Types
- var declarations
- Short declaration (:=)
- Basic types (int, string, bool, float64)
**Exercise**: Create variables for Pokemon's name, level, HP, and if it's shiny

### Lesson 3: User Input & Conversions
- fmt.Scan / fmt.Scanln
- Type conversions
- String formatting
**Exercise**: Pokedollar to Bitcoin converter (shady black market rates)

### Lesson 4: Control Flow - If/Else
- If statements
- If with initialization
- Comparison operators
**Exercise**: Pokeball selector (Pokemon HP determines which ball to use)

### Lesson 5: Loops
- For loop (the only loop in Go)
- While-style loops
- Infinite loops with break
- Continue keyword
**Exercise**: Wild Pokemon encounter simulator (1-100 encounters)

## Phase 2: Data Structures (Lessons 6-10)

### Lesson 6: Arrays & Slices
- Array declaration and initialization
- Slice creation and operations
- append, len, cap
**Exercise**: Find strongest and weakest Pokemon in your team

### Lesson 7: Slice Operations
- Slicing syntax [start:end]
- Copy and range
- Multi-dimensional slices
**Exercise**: Type effectiveness chart (2D slice with multipliers)

### Lesson 8: Maps
- Map declaration and initialization
- Adding, accessing, deleting
- Checking existence with ok
**Exercise**: Pokedex - track which Pokemon you've caught by location

### Lesson 9: Structs
- Defining structs
- Creating instances
- Accessing fields
- Anonymous structs
**Exercise**: Pokemon struct with Name, Type, Level, HP, Moves

### Lesson 10: Struct Embedding
- Nested structs
- Anonymous fields
- Struct tags (brief intro)
**Exercise**: Trainer struct containing Pokemon team and badge collection

## Phase 3: Functions & Methods (Lessons 11-15)

### Lesson 11: Functions
- Function declaration
- Parameters and return values
- Multiple return values
**Exercise**: Battle damage calculator functions (physical, special, critical hit)

### Lesson 12: Error Handling
- error type
- Creating errors
- Checking and handling errors
**Exercise**: Pokeball throw function (returns caught Pokemon or error)

### Lesson 13: Variadic Functions
- ... syntax
- Variable number of arguments
- Passing slices
**Exercise**: Experience points distributor (split XP among team)

### Lesson 14: Pointers Basics
- & and * operators
- Pass by value vs reference
- Nil pointers
**Exercise**: Pokemon evolution function (modifies original Pokemon stats)

### Lesson 15: Methods
- Method receivers
- Value vs pointer receivers
- Methods on any type
**Exercise**: Pokemon methods - TakeDamage, Heal, LevelUp (pointer receivers)

## Phase 4: Interfaces & Advanced (Lessons 16-20)

### Lesson 16: Interfaces
- Interface definition
- Implicit implementation
- Empty interface
**Exercise**: Battler interface implemented by different Pokemon types

### Lesson 17: Type Assertions & Switches
- Type assertion
- Type switches
- Interface values
**Exercise**: Move effectiveness checker (different move types)

### Lesson 18: Defer, Panic, Recover
- Defer statement
- Panic for errors
- Recover mechanism
**Exercise**: Pokemon Center healing system with proper cleanup

### Lesson 19: File I/O
- os package
- Reading files (ioutil, bufio)
- Writing files
**Exercise**: Save/load trainer's Pokemon team to file

### Lesson 20: JSON & Data Serialization
- json.Marshal and Unmarshal
- Struct tags
- Working with APIs
**Exercise**: Fetch and parse Pokemon data from PokeAPI

## Phase 5: Concurrency (Lessons 21-25)

### Lesson 21: Goroutines
- go keyword
- Concurrent execution
- time.Sleep (temporary sync)
**Exercise**: Multiple trainers encountering wild Pokemon simultaneously

### Lesson 22: Channels
- Channel creation
- Sending and receiving
- Blocking behavior
**Exercise**: Pokemon Center healing queue (send Pokemon, receive healed)

### Lesson 23: Channel Direction & Buffering
- Buffered channels
- Channel direction (send-only, receive-only)
- Close channels
**Exercise**: Black market Pokemon trading system with dealers and buyers

### Lesson 24: Select Statement
- Multiple channel operations
- Default case
- Timeouts
**Exercise**: Battle timeout system (Pokemon must attack within time limit)

### Lesson 25: Sync Package
- WaitGroups
- Mutexes
- Once
**Exercise**: Tournament betting system (safe concurrent bet placement)

## Phase 6: Real-World Projects (Lessons 26-30)

### Lesson 26: HTTP Server Basics
- net/http package
- HandleFunc
- ResponseWriter and Request
**Exercise**: PokeAPI proxy server (Hello World endpoint)

### Lesson 27: HTTP Routing & Methods
- Multiple routes
- GET vs POST
- Query parameters
**Exercise**: Pokemon tournament registration API (CRUD)

### Lesson 28: Command-Line Tools
- flag package
- os.Args
- Exit codes
**Exercise**: Pokedex CLI tool with search, add, list commands

### Lesson 29: Testing
- testing package
- Table-driven tests
- Benchmarks
**Exercise**: Test damage calculator and type effectiveness

### Lesson 30: Final Project
Choose one:
- **Battle Simulator**: Text-based Pokemon battle with AI opponent
- **Tournament Manager**: HTTP server for tournament brackets and betting
- **Shiny Hunter**: Concurrent Pokemon encounter simulator with odds
- **Trading Platform**: WebSocket-based real-time Pokemon trading system
- **Black Market API**: REST API for underground Pokemon deals with auth

User chooses project, implements with guidance.

## Progression Logic

- Linear through Phases 1-2 (foundations needed)
- Flexible in Phases 3-4 (can adjust based on interest)
- Phase 5 requires Phase 4 completion
- Phase 6 is application of all previous learning

## Skill Checks

After every 5 lessons, give a challenge that combines previous concepts:
- After L5: Wild Pokemon spawn system with user input and functions
- After L10: Pokemon team tracker with maps and structs  
- After L15: Pokemon battle system with methods and error handling
- After L20: Pokedex with JSON persistence and file I/O
- After L25: Concurrent tournament betting with worker pools and mutexes