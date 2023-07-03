## Variable in GO

### What is a variable?

A variable is a storage location for holding a value. The value of a variable can be changed throughout the program. The value of a variable is retrieved by referring to the variable name.

### Declaring a variable

To declare a variable in Go, we use the **var** keyword followed by the variable name and the type of the variable. For example, to declare a variable named age of type int, we write:

``` go
var age int
```

### Initializing a variable

To initialize a variable in Go, we use the **var** keyword followed by the variable name, the type of the variable, and the initial value of the variable. For example, to initialize a variable named age of type int with the value 20, we write:

``` go
var age int = 20
```

### Declaring and initializing a variable

To declare and initialize a variable in Go, we use the **var** keyword followed by the variable name, the type of the variable, and the initial value of the variable. For example, to declare and initialize a variable named age of type int with the value 20, we write:

``` go
var age int = 20 

or 

age := 20 // This is just another way to declare and initialize a variable in Go.

or

var age = 20 // This is another way to declare and initialize a variable in Go without type declaration.
```

### Declaring multiple variables

To declare multiple variables in Go, we use the **var** keyword followed by the variable names and the type of the variables. For example, to declare two variables named age and weight of type int, we write:

``` go

var age, weight int
    
    ```

### Initializing and Declaring multiple variables

To initialize multiple variables in Go, we use the **var** keyword followed by the variable names, the type of the variables, and the initial values of the variables. For example, to initialize two variables named age and weight of type int with the values 20 and 50 respectively, we write:

``` go

var age, weight int = 20, 50

or

age, weight := 20, 50 // This is just another way to initialize multiple variables in Go.

or

var age, weight = 20, 50 // This is another way to initialize multiple variables in Go without type declaration.
    
    ```
### Declaring and Initializing constants

To declare and initialize a constant in Go, we use the **const** keyword followed by the constant name, the type of the constant, and the initial value of the constant. For example, to declare and initialize a constant named pi of type float64 with the value 3.14, we write:

``` go

const pi float64 = 3.14 

or

const pi = 3.14 // This is another way to declare and initialize a constant in Go without type declaration.
    
    ```

Note: The value of a constant cannot be changed throughout the program.


