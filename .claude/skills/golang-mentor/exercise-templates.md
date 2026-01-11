# Example Exercise Templates

This file contains templates for creating exercises at different difficulty levels. Use these patterns when designing new exercises.

## Beginner Exercise Template

```
Exercise: [Pokemon-themed Title]
Difficulty: Beginner
Estimated Time: 5-10 minutes

**Concept**: [One sentence explaining what this teaches]

**Setup**: Create a file called `exercise_name.go`

**Starting Code**:
```go
package main

import "fmt"

func main() {
    // Your code here
}
```

**Your Task**:
1. [Specific instruction 1 with Pokemon context]
2. [Specific instruction 2 with Pokemon context]
3. [Specific instruction 3 with Pokemon context]

**Hints**:
- [Hint about syntax or approach]
- [Reminder about a relevant concept]

**Expected Output**:
```
[Exact output they should see - Pokemon themed]
```

**Test Your Code**:
Run: `go run exercise_name.go`
```

**Example - Pikachu Stats Display:**
```
Exercise: Display Pikachu's Battle Stats
Difficulty: Beginner
Estimated Time: 5 minutes

**Concept**: Variables and formatted output

**Setup**: Create `pikachu_stats.go`

**Starting Code**:
```go
package main

import "fmt"

func main() {
    name := "Pikachu"
    level := 25
    hp := 95
    attack := 55
    
    // Print formatted stats here
}
```

**Your Task**:
1. Print each stat on its own line
2. Format should be: "Stat: Value"
3. Add a title line at the top

**Expected Output**:
```
=== Pikachu's Stats ===
Name: Pikachu
Level: 25
HP: 95
Attack: 55
```
```

## Intermediate Exercise Template

```
Exercise: [Pokemon-themed Title]
Difficulty: Intermediate
Estimated Time: 10-15 minutes

**Goal**: [What they're building with Pokemon context]

**Requirements**:
- [Pokemon feature 1]
- [Pokemon feature 2]
- [Pokemon feature 3]

**Function Signature**:
```go
func FunctionName(params types) returnType {
    // Implement this
}
```

**Test Cases**:
```go
Input: [Pokemon example input]
Expected: [expected output]

Input: [Pokemon edge case]
Expected: [expected output]
```

**Expected Output**:
```
[Sample output showing it works]
```
```

**Example - Type Effectiveness Calculator:**
```
Exercise: Calculate Type Effectiveness
Difficulty: Intermediate
Estimated Time: 15 minutes

**Goal**: Build a function that calculates damage multiplier based on attack and defense types

**Requirements**:
- Handle all basic types (Fire, Water, Grass, Electric, etc.)
- Return correct multiplier (0.5x, 1x, 2x)
- Handle same-type matchups

**Function Signature**:
```go
func GetTypeEffectiveness(attackType, defenseType string) float64 {
    // Implement type chart logic
}
```

**Test Cases**:
```go
Input: GetTypeEffectiveness("Fire", "Grass")
Expected: 2.0

Input: GetTypeEffectiveness("Water", "Fire") 
Expected: 2.0

Input: GetTypeEffectiveness("Electric", "Ground")
Expected: 0.0

Input: GetTypeEffectiveness("Normal", "Normal")
Expected: 1.0
```

**Expected Output**:
```
Fire vs Grass: 2.0x damage
Water vs Fire: 2.0x damage
Electric vs Ground: No effect!
Normal vs Normal: 1.0x damage
```
```

## Advanced Exercise Template

```
Exercise: [Descriptive Title]
Difficulty: Advanced
Estimated Time: 20-30 minutes

**Problem**: [Brief description of the problem to solve]

**Constraints**:
- [Technical constraint 1]
- [Technical constraint 2]

**Example Usage**:
```go
// How the finished code should be used
result := YourSolution(input)
fmt.Println(result)
```

**Expected Behavior**:
- [Behavior 1]
- [Behavior 2]
- [Edge case handling]

**Consider**:
- Error handling
- Performance
- Code organization
```

## Challenge Problem Template

```
Challenge: [Engaging Title]
Estimated Time: 30-45 minutes

**Scenario**: [Real-world context for the problem]

**Build**:
[High-level description of what to create]

**Must Have**:
- [ ] [Core feature 1]
- [ ] [Core feature 2]
- [ ] [Core feature 3]

**Nice to Have**:
- [ ] [Enhancement 1]
- [ ] [Enhancement 2]

**Evaluation Criteria**:
- Does it work correctly?
- Is the code readable?
- Are errors handled appropriately?

**Example Interaction**:
```
$ go run challenge.go [args]
[Expected output]
```
```

## Exercise Pattern: Progressive Difficulty

### Level 1: Direct Application
```
Exercise: Display Pokemon Name
Create variables for pokemon name and type.
Print them in the format: "Pokemon is a Type-type Pokemon"
Example: "Charizard is a Fire-type Pokemon"
```

### Level 2: Small Logic Required
```
Exercise: Pokeball Recommendation
Get Pokemon's current HP percentage as input.
Recommend ball: Great Ball (>50%), Ultra Ball (20-50%), Master Ball (<20%)
```

### Level 3: Combining Concepts
```
Exercise: Team Statistics
Create a slice of Pokemon levels: []int{25, 42, 18, 65, 33, 51}
Calculate: average level, highest level, lowest level, how many ready for Elite Four (50+)
```

### Level 4: Multi-Step Problem
```
Exercise: Tournament Betting Analyzer
Read bet amounts from users (until they enter -1)
Track bets on different Pokemon
Calculate total pot, most popular Pokemon, payout odds
Determine if there's insider trading (suspicious patterns)
```

## Quick Exercise Patterns

### Pattern: Input → Process → Output

```go
// Basic
input := getUserInput()
result := process(input)
fmt.Println(result)
```

### Pattern: Loop + Accumulate

```go
// Running total/max/min
for _, value := range data {
    // Update accumulator
}
fmt.Println(accumulated)
```

### Pattern: Validation

```go
// Check conditions
if !isValid(input) {
    return error
}
// Process valid input
```

### Pattern: Transform Data

```go
// Convert one format to another
for _, item := range source {
    transformed := transform(item)
    dest = append(dest, transformed)
}
```

## Common Exercise Types by Topic

### Variables & Types
- Type conversion exercises (Level to string, HP to float)
- Multiple variable assignments (Pokemon stats)
- Const vs var practice (Base stats vs current stats)

### Control Flow
- If/else decision trees (Type matchups, status effects)
- Switch statement scenarios (Move selection, weather effects)
- Loop pattern practice (Wild encounter attempts, XP grinding)

### Data Structures
- Slice operations (Team management, move lists)
- Map CRUD operations (Pokedex entries, trainer inventory)
- Struct creation and access (Pokemon, Trainer, Move data)
- Nested structure navigation (Trainer→Team→Pokemon→Moves)

### Functions
- Pure function implementation (Damage calculations)
- Multiple return values (Catch attempt returns Pokemon and error)
- Error handling (Failed evolution, invalid move)
- Variadic functions (Multi-hit moves, splitting rewards)

### Pointers & Methods
- Pass by value vs pointer (Trading vs evolving)
- Method receivers (Pokemon.TakeDamage(), Pokemon.Heal())
- Modifying through pointers (Level up original Pokemon)

### Interfaces
- Multiple type implementations (Different Pokemon implementing Battler)
- Interface composition (Swimmer, Flyer interfaces)
- Type assertions (Checking if Pokemon can Mega Evolve)

### Concurrency
- Basic goroutine spawning (Multiple wild encounters)
- Channel communication (Pokemon Center healing queue)
- Select with multiple channels (Battle or flee decision)
- WaitGroup coordination (Tournament with multiple matches)

## Exercise Modifications for Adaptation

### Make Easier:
- Provide more of the code structure
- Add explicit hints in comments
- Break into smaller sub-tasks
- Show expected output format

### Make Harder:
- Remove scaffolding code
- Add edge cases to handle
- Require optimization
- Add feature requirements
- No hints, just goal

## Red Flags in Exercise Design

❌ Too much starter code (they're not learning)
❌ Solution is copy-paste from explanation
❌ Unclear success criteria
❌ Requires concepts not yet taught
❌ Too many steps (break into multiple exercises)
❌ Boring problem (add context or make it playful)

✅ Clear what to build
✅ Can verify correctness
✅ Requires applying the concept
✅ Achievable but not trivial
✅ Builds toward practical skills