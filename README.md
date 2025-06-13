# Programming Languages, Spring 2025

- Instructor: Prof. Y C Cheng
- Class meeting time: Mon 2-3-4
- Class meeting place:
  - Second Academic Building-207 (in person)
  - [Microsoft Teams](https://teams.microsoft.com/l/team/19%3ACxxkswKurkJFiUujdCtAfg7t9IO6ZwPGO2EiY4AyicA1%40thread.tacv2/conversations?groupId=763aba6a-09fe-4b95-9837-b7cf8d08d47d&tenantId=dfb5e216-2b8a-4b32-b1cb-e786a1095218)
- Office hours: Mon 5-6-7, Tue 2-3-4
- Course Repo: Course repository: http://140.124.181.100/yccheng/pl2025s
- TA:
  - Benny Wang \<benny870704@gmail.com\>, Regina Yen\<ycycchre@gmail.com\> (Hong-Yue technology Research Building, Room 1324)
  - Office hours: 10-12 am, Wed
- TA's homework repository: [pl2025_TA](http://140.124.181.100/course/pl2025s_ta)
- Please hand in the homework on [Microsoft Teams](https://teams.microsoft.com/l/team/19%3ACxxkswKurkJFiUujdCtAfg7t9IO6ZwPGO2EiY4AyicA1%40thread.tacv2/conversations?groupId=763aba6a-09fe-4b95-9837-b7cf8d08d47d&tenantId=dfb5e216-2b8a-4b32-b1cb-e786a1095218).

- HW1 is announced, deadline: 3/3(Mon.) 23:59.

## Synopsis

We will take a practice-first approach to learn the most frequently seen language features topic by topic.

Each topic consists of three parts given a problem:

1. A program we will write in **a style satisfying the constraints** in solving the problem. We will discuss the consequences of the style
2. The principles in language design that implements the constraints
3. The same program in Go written either in class or as an exercise to illustrate how the constraints are satisfied.

Some problems we will solve:

- [Term frequencies](https://github.com/crista/exercises-in-programming-style)

We will repeat the three-parts for a number of styles under the following programming paradigms:

- structured programming
  - Cookbook
  - Monolithic
  - Pipeline
- functional programming
  - Infinite Mirror
  - Kick Forward
  - The One
- object-oriented programming
  - Things
  - Letter Box
  - Abstract Things
  - Hollywood
  - Bulletin Board
- concurrency
  - Map Reduce

## Reference Books

- Donovan, Alan AA, and Brian W. Kernighan. The Go programming language. Addison-Wesley Professional, 2015. [Online resource](https://www.gopl.io)
- Cristina Videira Lopes, Exercises in Programming Style, First Edition, Chapman and Hall/CRC, 2014.
- Michael L. Scott, Programming Language Pragmatics, Fourth Edition, Morgan Kaufmann (Elsevier), 2016.

## Lectures

**Week 1**

- Program: Cookbook
  - run the program
  - unit-tests it
- Principle: Names, Bindings, and scope
- Program in Go and constraints resolution
  - TDD the Go program
  - Reexamine the Principle

**Week 2**

- pipeline style and mathematical function
- why single input and output?
- Currying
- idempotence
- runtime environment - stack needed (Pragmatics).

== HW1, deadline: 3/3(Mon.) 23:59

1. (30%) Do exercise 5.2 of **Exercises in Programming Style** (see below.) Hint: apply currying to remove_stop_words. Write unit tests for this and change the main program to use this function.

Example program: https://github.com/crista/exercises-in-programming-style/blob/master/06-pipeline/tf-06.py

In the example program, the name of the file containing the list of stop words is hard-coded (line #36). Modify the program so that the name of the stop words file is given as the second argument in the command line. You must observe the following additional stylistic constraints: (1) no function can have more than 1 argument, and (2) the only function that can be changed is remove_stop_words; it’s ok to change the call chain in the main block in order to reflect the changes in remove_stop_words.

2. (30%) Show with unit testing for the functions in the Cookbook style that are **not idempotent** and describe the reason in the comment.

3. (40%) Rewrite the Pipeline style in Go, including unit tests.

**Week 3**

1. Names, bindings, lifetimes, and scopes
   - Chapter 3, Pragmatics (p1-35, pptx.)
   - Chapter 2, Go
     - _keywords_ (language design time)
     - _predeclared names_ (language implementation time)
       - constants: e.g., true, false, nil
       - types: e.g., int, int64
       - functions: e.g., len, new, close
     - name declarations (program writing time)
       - var, const, type, func
       - package-level declarations
       - local declarations
     - static allocation, stack allocation, and heap allocation
     - lifetime management and garbage collection
       - a local variable escapes from a function
         - C/C++ vs. Go
     - stack vs. heap allocation depending on reachability to entities, not the use of "new"
       - curried three-arg addition, closure, capture and escape
2. Exercise 11, Things: Style
3. Exercise 12, Letterbox: Style

**Week 4**

- object closure
  - C++ functor and STL
    - overloading call operator
- object-oriented programming in Go in shapesapp
  - module, directory structure, and packages
  - type, type struct, type interface
  - Encapsulation
    - Cap for public
    - Lowercase for private
  - Factory functions NewTypeName for constructing object
  - Methods: function with receivers
    - can be defined for any **named type** other than a pointer or an interface
    - value receiver
    - pointer receiver

== HW2, deadline: 3/17(Mon.) 23:59

1. 10.1 with Go. You must include unit tests. (without the info method)
   - Link: https://github.com/crista/exercises-in-programming-style/tree/master/11-things

**Week 5**

Go OOP language features

- structure embedding
  - field promotion (anonymous field)
  - method promotion
  - struct embedding for reuse (vs inheritance)
  - interface (implicit) to enable polymorphism
- composition with struct embedding vs. inheritance
  - "Favor object composition over class inheritance."
  - Composite pattern
    - shapesapp: CompositeShape
    - Add(): use struct embedding to add default add behavior
  - What happens with design patterns of Class scope, e.g., Template method

**Week 6**

## Progressing from ad hoc polymorphism to parametric polymorphism: C++, Java, and Go

### Introduction

In this lecture, we will explore two types of polymorphisms in C++, Java, and Go:

1. **Ad hoc Polymorphism (Function Overloading & Operator Overloading)**
2. **Parametric Polymorphism (Generics/Templates)**

We will use a simple function that adds two arguments as a running example to compare how the three languages handle these polymorphisms. Additionally, we will discuss key concepts such as **universal types (void \* in C/C++, Object in Java and interface{} in Go)**, **monomorphization**, **type erasure**, and **boxing/unboxing**. The lecture will also cover how parametric polymorphism was handled before and after Java 5 (generics introduction) and Go 1.18 (generics introduction).

---

## Ad Hoc Polymorphism (Function Overloading & Operator Overloading)

### C++ (Function & Operator Overloading)

C++ supports ad hoc polymorphism through function overloading and operator overloading.

**Example:**

```cpp
#include <iostream>
using namespace std;

int Add(int a, int b) { return a + b; }
double Add(double a, double b) { return a + b; }

int main() {
    cout << Add(2, 3) << endl;        // Outputs: 5
    cout << Add(2.5, 3.14) << endl;   // Outputs: 5.64
    return 0;
}
```

- **Compile-time Resolution:** The compiler determines which version of `Add()` to call based on the argument types.
  - type of `Add`: `int (int, int)`
- **Type Safety:** Enforced at compile time.
- **Performance:** No runtime overhead.

### Java (Pre-Java 5)

Java supports ad hoc polymorphism through **method overloading** but **not operator overloading**.

**Example:**

```java
public class Add {
    static int add(int a, int b) { return a + b; }
    static double add(double a, double b) { return a + b; }

    public static void main(String[] args) {
        System.out.println(add(2, 3));       // Outputs: 5
        System.out.println(add(2.5, 3.14));  // Outputs: 5.64
    }
}
```

- **Compile-time Resolution:** The compiler selects the appropriate method based on argument types.
- **Type Safety:** Enforced at compile time.
- **Performance:** No runtime overhead.

### Go (Pre-Go 1.18)

Go does not support function overloading or operator overloading.

**Example:**

```go
func AddInt(a, b int) int {
    return a + b
}

func AddFloat(a, b float64) float64 {
    return a + b
}
```

- **No Overloading:** Functions must have unique names.
- **Explicit Function Naming:** Developers must define separate functions for different types.
- **Performance:** No runtime overhead.

---

## Simulate ad hoc polymorphism with Universal type

### C++: `void *`

```cpp
#include <stdio.h>

void* Add(void* a, void* b, int type) {
    if (type == 1) {
        int* result = new int;
        *result = *(int*)a + *(int*)b;
        return result;
    }
    else if (type == 2) {
        float* result = new float;
        *result = *(float*)a + *(float*)b;
        return result;
    }
    else if (type == 3) {
        char* result = new char[100];
        sprintf(result, "%s%s", (char*)a, (char*)b);
        return result;
    }

    return NULL;
}

void printValue(void *value, char type) {
    switch (type) {
        case 'i':
            printf("Integer: %d\n", *(int*)value);
            break;
        case 'f':
            printf("Float: %.2f\n", *(float*)value);
            break;
        case 's':
            printf("String: %s\n", (char*)value);
            break;
        default:
            printf("Unknown Type\n");
    }
}

int main() {
    int a = 1;
    int b = 2;
    void* result = Add(&a, &b, 1);
    printf("sum = %d\n", *(int*)result);
    delete (int*)result;

    float c = 3.14f;
    float d = 2.71f;
    result = Add(&c, &d, 2);
    printf("sum = %.2f\n", *(float*)result);
    printf("sum = %d\n", *(int*)result); // compiles and runs but a bug: no type checking
    delete (float*)result;

    char* e = "Hello, ";
    char* f = "C";
    result = Add(e, f, 3);
    printf("sum = %s\n", (char*)result);
    delete[] (char*)result;

    // int i = 42;
    // float f = 3.14f;
    // char *s = "Hello, C";

    // printValue(&i, 'i');
    // printValue(&f, 'f');
    // printValue(s, 's');

    return 0;
}
```

- Only the memory address is stored.
- The programmer must manually remember and cast to the correct type.
- No automatic checking — if the cast is incorrect, it results in undefined behavior.

---

### Go: the empty interface type `interface{}`

```go line_numbers
package main

import (
	"fmt"
)

func PrintValue(val interface{}) {
	switch v := val.(type) { // type assertion and type switch
	case int:
		fmt.Println("Integer:", v)
	case float64:
		fmt.Println("Float:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown Type")
	}
}

func main() {
	PrintValue(42)
	PrintValue(3.14)
	PrintValue("Hello, Go")
}
```

- The empty interface type `interface{}` declares no method, a variable of `interface{}` type can hold a value of any type because any value implicitly satisfies it.
- interface{} uses **structural typing** (structural equivalence)

```go
package main

import "fmt"

func main() {
	var any interface{}
	any = 1
	fmt.Println(any)
	any = "hello"
	fmt.Println(any)
	any = true
	fmt.Println(any)
	any = 1.0
	fmt.Println(any)
	any = []int{1, 2, 3}
	fmt.Println(any)
	any = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(any)
	any = 1
	// x := any + 1
}
```

- Stores a **type descriptor** and a **pointer to the actual data**.
- When performing a type assertion (val.(int)), Go checks if the stored type descriptor matches int.
- If the type assertion fails, an error (panic) or a nil return value is provided depending on how the assertion is made.
- variadic function `fmt.println(x, y)`, `fmt.println(1.0, "hello", []{1,2,3})`

```go
func Add(a, b interface{}) interface{} {
	switch v := a.(type) { // type switch
	case int:
		if bInt, ok := b.(int); ok { // type assertion
			return v + bInt
		}
	case string:
		if bString, ok := b.(string); ok {
			return v + bString
		}
	case float64:
		if bFloat, ok := b.(float64); ok {
			return v + bFloat
		}
	default:
	}
	return nil
}
```

### type assertion x.(T) and type switch switch x.(type) {...}

- if T is a concrete type: get dynamic value of x if its type is identical to T.
  - succeed: get value
  - failure: panic
- if T is an interface type: check if x _satisfies_ T
  - result of type T but x's dynamic value is unchanged
- `switch v := a.(type)`: v is a variable that holds the value of a cast to the matched type in the case.

### Java: `Object` with `instanceof`

```java
public class Main {
    public static void printValue(Object val) {
        if (val instanceof Integer) {
            System.out.println("Integer: " + val);
        } else if (val instanceof Double) {
            System.out.println("Float: " + val);
        } else if (val instanceof String) {
            System.out.println("String: " + val);
        } else {
            System.out.println("Unknown Type");
        }
    }

    public static void main(String[] args) {
        printValue(42);
        printValue(3.14);
        printValue("Hello, Java");
    }
}
```

```java
public class UniversalAdd {
    public static Object add(Object a, Object b) {
        if (a instanceof Integer && b instanceof Integer) {
            return (Integer) a + (Integer) b;
        } else if (a instanceof Double && b instanceof Double) {
            return (Double) a + (Double) b;
        }
        return null;
    }
}

UniversalAdd.add(1, 2);
```

- Object uses **Nominal Typing** and represents any **reference type** (referential equivalence)
- In Java, instanceof is used to check the type, followed by casting.

## Parametric Polymorphism (Generics/Templates)

### C++ (Templates)

C++ supports parametric polymorphism through **templates**.

**Example:**

```cpp
template <typename T>
T Add(T a, T b) {
    return a + b;
}

int main() {
    std::cout << Add(2, 3) << std::endl;         // Outputs: 5
    std::cout << Add(2.5, 3.14) << std::endl;    // Outputs: 5.64

    Add(1, 2);
    Add(1.0, 2.0);
    // Add(1.0, 2);
    printf("%f\n", Add<double>(1.0, 2));
    printf("%d\n", Add<int>(1.0, 2)); // bug: no type checking

    return 0;
}
```

- **Monomorphization:** The compiler generates a separate instance of the function for each type used (e.g., `Add<int>` and `Add<double>`).
- **Compile-time Resolution:** No runtime overhead.
- **Type Safety:** Checked at compile time.

---

### Java (Post-Java 5 Generics)

Java supports parametric polymorphism through **generics**.

**Example:**

```java
public class GenericAdder {

    // Generic method to add two numbers of the same type
    public static <T extends Number> Number add(T a, T b) {
        // Handle different numeric types
        if (a instanceof Integer) {
            return a.intValue() + b.intValue();
        } else if (a instanceof Double || b instanceof Double) {
            return a.doubleValue() + b.doubleValue();
        } else if (a instanceof Float || b instanceof Float) {
            return a.floatValue() + b.floatValue();
        } else if (a instanceof Long) {
            return a.longValue() + b.longValue();
        } else {
            // Default to double for other number types
            return a.doubleValue() + b.doubleValue();
        }
    }

    // Example usage
    public static void main(String[] args) {
        System.out.println(add(1, 2));         // Integer addition: 3
        System.out.println(add(3.0, 4.5));     // Double addition: 7.5
        System.out.println(add(5, 6.7));       // Mixed types will use double: 11.7
    }
}
```

- **Type Erasure:** Generics are implemented by erasing type information at runtime.
- **Boxing/Unboxing:** Primitive types are converted to their wrapper classes (`Integer`, `Double`, etc.) during this process.
- **Type Safety:** Ensured at compile-time, but type information is lost at runtime.
- **Performance:** Some runtime overhead due to boxing/unboxing.

---

### Go (Post-Go 1.18 Generics)

Go supports parametric polymorphism via **type parameters**.

**Example:**

```go
func Add[T int | float64](a, b T) T {
    return a + b
}
```

- **Type Constraints:** `T int | float64` allows only `int` and `float64` types.
- **Compile-time Resolution:** Type checking occurs during compilation.
- **No Boxing/Unboxing:** Unlike Java, Go operates directly on primitive types.
- **Performance:** Similar to C++ monomorphization

```go
func TestAddGIFF(t *testing.T) {
	s := go_interface.AddG(go_interface.AddG(1.0, 2), 3)

	if s != 6 {
		t.Errorf("Expected 6, got %v", s)
	}

	// how do I get s.'s type?
	fmt.Println(reflect.TypeOf(s))
}
```

---

## Summary

- **C++:** Powerful static polymorphism using templates with monomorphization. Universal type (`void*`) is dangerous and not type-safe.
- **Java:** Uses type erasure for generics, resulting in some runtime overhead. Universal type (`Object`) requires casting and boxing/unboxing.
- **Go:** Newly introduced generics are powerful and efficient. Universal type (`interface{}`) allows for flexibility but requires type assertion.

### ### More examples in Java Generics

```java
interface Chooser<T> {
    public boolean better(T a, T b);
}

class Arbiter<T> {
    T bestSoFar;
    Chooser<? super T> comp;
    public Arbiter(Chooser<? super T> c) {
        comp = c;
    }

    public void consider(T t) {
        if (bestSoFar == null || comp.better(t, bestSoFar)) bestSoFar = t;
    }

    public T best() {
        return bestSoFar;
    }
}
```

```java
class StringLengthChooser implements Chooser<String> {
    public boolean better(String a, String b) {
        return a.length() > b.length();
    }
}

public class Main {
    public static void main(String[] args) {
        Arbiter<String> arbiter = new Arbiter<>(new StringLengthChooser());

        arbiter.consider("apple");
        arbiter.consider("banana");
        arbiter.consider("kiwi");

        System.out.println("Best: " + arbiter.best());  // Outputs: "banana"
    }
}
```

### Why Chooser<? super T>?

```java
class Animal {
    String name;
    int weight;

    public Animal(String name, int weight) {
        this.name = name;
        this.weight = weight;
    }
}

class Dog extends Animal {
    public Dog(String name, int weight) {
        super(name, weight);
    }
}

class AnimalWeightChooser implements Chooser<Animal> {
    public boolean better(Animal a, Animal b) {
        return a.weight > b.weight; // Heavier animal is considered better
    }
}

public class Main {
    public static void main(String[] args) {
        Arbiter<Dog> dogArbiter = new Arbiter<>(new AnimalWeightChooser()); // Note: AnimalWeightChooser is a Chooser<Animal>

        Dog dog1 = new Dog("Buddy", 25);
        Dog dog2 = new Dog("Rex", 30);
        Dog dog3 = new Dog("Bella", 20);

        dogArbiter.consider(dog1);
        dogArbiter.consider(dog2);
        dogArbiter.consider(dog3);

        Dog bestDog = dogArbiter.best();
        System.out.println("Best Dog: " + bestDog.name);  // Outputs: "Rex"
    }
}
```

```java
class DogNameLengthChooser implements Chooser<Dog> {
    public boolean better(Dog a, Dog b) {
        return a.name.length() > b.name.length(); // Longer name is considered better
    }
}

public class Main {
    public static void main(String[] args) {
        // Creating an Arbiter that uses DogNameLengthChooser for comparison
        Arbiter<Dog> dogArbiter = new Arbiter<>(new DogNameLengthChooser());

        Dog dog1 = new Dog("Max", 25);
        Dog dog2 = new Dog("Alexander", 30);
        Dog dog3 = new Dog("Bella", 20);

        dogArbiter.consider(dog1);
        dogArbiter.consider(dog2);
        dogArbiter.consider(dog3);

        Dog bestDog = dogArbiter.best();
        System.out.println("Best Dog: " + bestDog.name);  // Outputs: "Alexander"
    }
}
```

**week7**

HW3: deadline 4/7(Mon.) 23:59, Implementing the Template Method Pattern in Go

Background

In software development, many applications need to generate documents in different formats, such as plain text documents and HTML documents. The document generation process typically follows these common steps:

1. Prepare the raw data
2. Format the content
3. Save the formatted content

Although these steps are fixed in sequence, different document types may have variations in implementation. For example:

HTML documents require tags like `<html>` and `<body>`.
Plain text documents may only need simple strings.

To unify this process, we can use the Template Method Pattern, which encapsulates these steps within a Base Document Generator while allowing different document types to define their specific behaviors.

_Assignment Requirements_

Based on the above concept, design a Document Generator Framework and complete the following tasks.

1. Define the DocumentGenerator Interface

The DocumentGenerator interface defines the fundamental steps for document generation. Each document type must implement the following methods:

PrepareData() string // Prepare the raw content of the document
FormatContent(data string) string // Format the content
Save(content string) string // Save the formatted content

Each step’s implementation will depend on the specific document type.

2. Implement the BaseGenerator Class

Design a Base Document Generator (BaseGenerator) that contains the Generate() method, which executes the three steps defined in the DocumentGenerator interface sequentially.
The Generate() method does not implement specific logic; instead, it relies on the document type’s implementation of the DocumentGenerator interface.

3. Implement TextDocument and HTMLDocument Classes

Implement the following two document types and ensure they conform to the DocumentGenerator interface:

(1) Plain Text Document (TextDocument)

PrepareData(): Returns `This is the raw text data`.
FormatContent(data string): Returns `Formatted Text: {data}`
Save(content string): Returns `Saving text document: {formatted_content}`

- Expected Output Example:

`Saving text document: Formatted Text: This is the raw text data.`

(2) HTML Document (HTMLDocument)

PrepareData(): Returns `<html><body>This is raw HTML data.</body></html>`
FormatContent(data string): Returns `<div>{data}</div>`
Save(content string): Returns `Saving HTML document: {formatted_content}`

- Expected Output Example:

`Saving HTML document: <div><html><body>This is raw HTML data.</body></html></div>`

4. Write Unit Tests

Write test cases for TextDocument and HTMLDocument to ensure the Generate() method executes correctly.

Grading Criteria

1. Correct implementation of TextDocument and HTMLDocument: 40%
2. Generate() method follows Template Method Pattern: 30%
3. Completeness and correctness of test cases: 30%

**Week 8**

### Style: Bulletin board

- publish-subscribe
  - events: immutable objects carrying information for action
  - tuple: a composite type
    - Ordered: Elements maintain their position.
    - Immutable: Once created, elements can’t be changed.
    - Heterogeneous: Can contain different data types.
    - Indexable: Accessed via position (e.g. t[0]).
    - Iterable: Can be looped over

```python
# Basic tuple
t = (1, 2, 3)

# Without parentheses (works due to commas)
t = 1, 2, 3

# Single-element tuple needs a comma
t = (1,)       # Tuple
not_a_tuple = (1)   # Just an int

# Nested tuple
t = (1, (2, 3), 4)

t = (10, 20, 30)
print(t[1])     # 20
print(t[-1])    # 30

len(t)              # Length
t + (40, 50)        # Concatenation
t * 2               # Repetition
20 in t             # Membership

for i, value in enumerate(['a', 'b']):
    print(i, value)


```

### Style: Introspective

- Obtain information about self at runtime
  - runtime stack
    - named tuple: a subtype of tuple with selector syntax
    - inspect module:
      - inspect.stack() to get stack frames

```python
FrameInfo = namedtuple('FrameInfo',
    ['frame', 'filename', 'lineno', 'function', 'code_context', 'index']
)

<class 'inspect.FrameInfo'>
```

    - local variables: with locals()
        - returns a dictionary of the current local symbol table — i.e., all local variables and their values in the current scope.

### Style: Reflection

- Change program behaviors at runtime
  - exec(): takes a string and interprets it as a python program text

```python
exec(object[, globals[, locals]])

```

- object: A string of Python code, or a compiled code object.
- globals (optional): A dictionary to use as the global namespace.
- locals (optional): A dictionary to use as the local namespace.

```python
code = """
def greet(name):
    print("Hello,", name)

greet("Alice")
"""

exec(code)
```

```python
global_env = {}
local_env = {}

exec("x = 5", global_env, local_env)

print(local_env["x"])   # 5
print(global_env)       # Doesn't contain 'x'
```

**Week 9**

### Data objects

### Matching

### List and list operations

### Lambda in Prolog

Creates an anonymous predicate that can be passed to higher-order predicates like maplist/2, include/3, exclude/3, etc.

```prolog
{CapturedVars}/[Args]>>Goal
```

- CapturedVars: Variables from the outer scope you want to “bring in”.

- [Args]: Arguments passed when the lambda is called (like parameters).

- Goal: The body (a Prolog goal that gets executed).

- Example:

```prolog
exclude({StopWords}/[W]>>(member(W, StopWords)), Words, Filtered).
```

- exclude/3 is a built-in that removes elements matching a predicate.

- The predicate is defined inline as a lambda:
  - {StopWords} captures the list of stop words from the outer scope.
  - [W] is the input word to check.
  - member(W, StopWords) succeeds if W is a stop word — i.e., one we want to exclude.

Since exclude removes items where the predicate succeeds, this keeps only the words not in StopWords.

**Week 10**
Midterm exam.

**Week 11**

### Prolog for AI, chapter 3, 4, 5, and 6.

lists and list processing

member/2, append/3, length/2

Built-in predicates for meta-programming
findall(X,P,L), maplist

the dynamic directive

:- dynamic instance/2.

assert, asserta, assertz: add new facts and rules to database (set of relations)
retract

T =..[Functor|Args], call(T)
call(T) treat term T as a goal (goal has the syntax of a term)

Cut !: green and red cuts.
fail

reading and writing from files
exception

**Week 12**
Prolog continued.

**Week 13, Prolog continued.**
problem representation as state space

- a state space, s(X,Y)
- a start node
- goal nodes defined by a goal condition
- depth-first search (fig11_7.pl)
- breadth-first search (fig_11_10.pl)
- computational complexity

heuristic search and A\*

- pick the least expensive subtree to expand f = g+h
  - g actual cost to current node
  - h estimated cast to a goal node
- s(X,Y,C)
- h(X,H)v

- a tree can be a leaf or a tree
  l(N, F/G)
  t(N, F/G, Subs)
- Subtrees of a node N are ordered by f-values
- The subtress with the smalles f-value is expanded within bound.
- expand stops when you find a goal; continue stops when you find a solution
- expand continues for leaf (add new fringe trees), tree
- expand no if bounded exceed, tree go back to waiting trees ranked by f-value
- expand never if no trees to expand
- continue gives yes,
- continue expand the next best tree for no (tree goes back to waiting), never (tree pruned)
- s(N,M,C) and h(N) embbedded into A\*
- state space search with A\*

Homework assignment #6
Problem 12.5 Bratko 4th ed.

Modify the best-first search program fig12_3.pl to count the number of nodes generated in the search. One easy way to do this is to keep the current number of nodes asserted as a fact, and update it by **retract** and **assert** whenever new nodes are generated. Experiment with various heuristic functions for the eight puzzle with respect to their heuristic power, which is relfected by the number of nodes generated.

**Answers to Midterm**

1.  text - executable code of programs,
    data - globals, literal constants
    stack - local variables, temps, return address, pointer to access names in enclosing scope
    heap - allocated

Entity | Memory Segment | Reason
global | Data | Package-level global variable
note | Stack | Local pointer variable
Note object | Heap | Escapes local scope; assigned to global
"Go" | Data | string literal

2.

(a) data-object - structure date(4,21, Year) - simp nle object - variable X, \_, \_123 - constant - atom a, b, 'it rains' - number 1, 3,14

(b) Yes.
The swap keeps the declarative meaning.

(c)
infinite recursion that causes the stack to overflow. Procedural meaning prevents it to complete the recursion caused by applying the walk before any push move.

3.  (a) C++ void \*, Go interface{}

(b)
func UniversalAdd(a, b interface{}) interface{} {
switch v := a.(type) { // type switch
case int:
if bInt, ok := b.(int); ok { // type assertion
return v + bInt
}
case string:
if bString, ok := b.(string); ok {
return v + bString
}
case float64:
if bFloat, ok := b.(float64); ok {
return v + bFloat
}
default:
}
return nil
}

(c)
func Add[T int | float64 | string](a, b T) T {
return a + b
}

4.

func CurriedThreeArgAddition(a int) func(int) func(int) int {
// This function takes 'a' and returns a function closure remembering 'a' that takes 'b'
return func(b int) func(int) int {
return func(c int) int {
return a + b + c
}
}
}

5.  (a) 25, capture by ref, so x = 20. in bar x = 5, f(5) = 20+5 = 25.
    (b) 21, x is captured by ref x++ => x = 21

6.  (a) Idempotence: A function, or procedure, is idempotent when multiple applications of it produce always the same observable eﬀects as the first application.

(b) Side eﬀects: A piece of code is said to have side eﬀects when it modifies existing state or has an observable interaction with the world. Examples of side eﬀects: modifying the value of a non-local variable or of an argument,
reading/writing data from/to a file or the network or the display, raising
an exception, and calling a function that has side eﬀects.

(c) Computational reflection: The ability for programs to access information
about themselves and modify themselves.

(d) Struct embedding: Use one named struct type as an anonymous field of another struct type to provide a syntactic short cut.

**Week 14, Prolog continued**

- 8-puzzle with A\*
- scheduling with A\*
- the state space search
  - key abstractions: successor relation s(X,Y,C), heuristic function h(X,H)
- admissibility

- more computational examples
  - bubble sort
  - insertion sort
  - quicksort

**Week 15**

A\* task scheduling - an admissible heuristic function h

Concurrency

- mutual exclusion
  - operational approach
  - axiomatic approach: {P} S {Q}
    If S statement is executed in a state satisfying the predicate P, and if it terminates, then it ends in a state satisfying Q.
- communicating sequential process (csp)
  - go routines and channels
  - Chapter 8, Goroutines and Channels
    - spinner
    - clock1, clock2
    - netcat1, netcat2, netcat3
    - reverb1, reverb2
    - pipeline1, pipeline2, pipeline3
- shared variable (Go, Chapter 9)
  - Concurrency with Shared Variables
    - mutex, etc.

**Week 16** - goroutines - reverb1, reverb2 - pipeline1, pipeline2, pipeline3 - channels (8.4) - unbuffered - pipeline - unidirectional - buffered - looping in parallel (8.5) - capture in functional literals (section 5.6.1)

- shared variable (Go, Chapter 9)
  - Concurrency with Shared Variables
    - mutex, etc.

Homework. Exercise 8.5 Take the CPU-bound sequential program, the Madelbrot program of Section 3.3, and execute its main loop in parallel using channels for communication. How much faster does it run on a multiprocessor machine? What is the optimal number of goroutines to use?

**Week 17**

Section 3, Concepts and Notations for Concurrent Programming, Section 3.1-3.2
Go: Section 8.7, 9.1-9.3

## Final Exam coverage: Close book, notes, and devices.

- Prolog: state space approach, heuristic search with A\*
  - ppt for Weeks 13-15 on Teams
- Concurrency:
  - paper "Concepts and Notations for Concurrent Programming", Section 1, Section 3.1-3.2
  - Go book Section 8.1-8.5, 8.7, 9.1-9.2
