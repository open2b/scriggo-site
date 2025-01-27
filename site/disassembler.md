{% extends "/layouts/doc.html" %}
{% macro Title string %}Scriggo command{% end %}
{% Article %}

{% raw code %}

# Disassembler

When Scriggo runs a program or template, first compiles it into a bytecode and then runs it on the Scriggo Virtual Machine.

## How to disassemble

To disassemble a program or a template you can:

* in the [Scriggo Playground](https://play.scriggo.com/) click on the `Disassemble` button to show the disassembled program.

* with the scriggo package, call the method `Disassemble` on the value returned by the `Build` method to disassemble a named package:

    ```
    asm, err = program.Disassemble("main")
    ``` 

* with the scriggo package, call the method `Disassemble` on the value returned by the `BuildTemplate` method to disassemble the template:

    ```
    asm = template.Disassemble(-1)
    ``` 

## Registers

The Scriggo Virtual Machine has 500 registers for each called function to store its local variables:  

* 125 integer registers: `i1`, `i2`, ..., `i125`
* 125 floating-point registers: `f1`, `f2`, ..., `f125`
* 125 string registers: `s1`, `s2`, ..., `s125`
* 125 general registers: `g1`, `g2`, ..., `g125`

Local variables with basic type an integer or boolean type are stored in the int registers:
```go
var a int
var b uint32
var c rune
var d time.Duration = 5 * time.Second
var e bool // false is stored as 0
var f MyBool = true // true is stored as 1
```
Local variables with basic type a floating-point type are stored in the float registers:
```go
var a float32
var b float64
var d MyFloat = 3.2
```
Local variables with basic type string are stored in the string registers:
```go
var a string
var b MyString = "hello"
```
All other local variables are stored in the general registers:
```go
var a []int
var b map[string]string
var c interface{}
var d time.Time
var e *S
``` 

## <em>ok</em> flag

In addition to the registers for called functions, there is special boolean `ok` flag used by `Assert`, `If`, `MapIndex`, `Select`, `Range` and `Receive` instructions.

## Package clause

The clause `Package` begins a disassembled code and defines the package to which the code belongs.

```go
Syntax:  Package name ; description: package name
```

## Import declaration

The declaration `Import` states that the disassembled package uses exported identifiers of the imported package.   

```go
Syntax:  Import "pkg" ; description: import "pkg"
```

## Func declaration

The declaration `Func` declares a function with its body and specifies the number of registers that it uses.

```go
Syntax:  Func f(a1, ...an) (r1, ...rn) ; func f(a1, ...an) (r1, ...rn) 
                 ; regs(ni,nf,ns,ng)
```

```go
Example: Func sum(i1, i2 int) (i3 int)
                 ; regs(3,0,0,0)
```

## Macro declaration

The declaration `Macro` declares a macro with its body and specifies the number of registers that it uses.

```go
Syntax:  Macro m(a1, ...an) T ; {% macro m(a1, ...an) T %} ... {% end %}
                 ; regs(ni,nf,ns,ng)
```

```go
Example: Macro Head(s5 markdown) (s7 markdown)
                 ; regs(1,0,3,0)
```

## Assembly instructions

The Scriggo assembly is an abstraction above the virtual machine instructions. Some assembly instructions have a direct representation in a virtual machine instruction but some do not.

There are 72 assembly instructions:

<ol>
  <li><a href="#add">Add</a></li>
  <li><a href="#addr">Addr</a></li>
  <li><a href="#and">And</a></li>
  <li><a href="#andnot">AndNot</a></li>
  <li><a href="#append">Append</a></li>
  <li><a href="#appendslice">AppendSlice</a></li>
  <li><a href="#assert">Assert</a></li>
  <li><a href="#break">Break</a></li>
  <li><a href="#call">Call</a></li>
  <li><a href="#cap">Cap</a></li>
  <li><a href="#case">Case</a></li>
  <li><a href="#close">Close</a></li>
  <li><a href="#complex64">Complex64</a></li>
  <li><a href="#complex128">Complex128</a></li>
  <li><a href="#continue">Continue</a></li>
  <li><a href="#convertnumber">ConvertNumber</a></li>
  <li><a href="#convertslice">ConvertSlice</a></li>
  <li><a href="#convertstring">ConvertString</a></li>
  <li><a href="#concat">Concat</a></li>
  <li><a href="#copy">Copy</a></li>
  <li><a href="#defer">Defer</a></li>
  <li><a href="#delete">Delete</a></li>
  <li><a href="#div">Div</a></li>
  <li><a href="#field">Field</a></li>
  <li><a href="#func">Func</a></li>
  <li><a href="#getvar">GetVar</a></li>
  <li><a href="#getvaraddr">GetVarAddr</a></li>
  <li><a href="#go">Go</a></li>
  <li><a href="#goto">Goto</a></li>
  <li><a href="#if">If</a></li>
  <li><a href="#index">Index</a></li>
  <li><a href="#len">Len</a></li>
  <li><a href="#load">Load</a></li>
  <li><a href="#loadfunc">LoadFunc</a></li>
  <li><a href="#makearray">MakeArray</a></li>
  <li><a href="#makechan">MakeChan</a></li>
  <li><a href="#makemap">MakeMap</a></li>
  <li><a href="#makeslice">MakeSlice</a></li>
  <li><a href="#makestruct">MakeStruct</a></li>
  <li><a href="#mapindex">MapIndex</a></li>
  <li><a href="#methodvalue">MethodValue</a></li>
  <li><a href="#move">Move</a></li>
  <li><a href="#mul">Mul</a></li>
  <li><a href="#neg">Neg</a></li>
  <li><a href="#new">New</a></li>
  <li><a href="#notzero">NotZero</a></li>
  <li><a href="#or">Or</a></li>
  <li><a href="#panic">Panic</a></li>
  <li><a href="#print">Print</a></li>
  <li><a href="#range">Range</a></li>
  <li><a href="#realimag">RealImag</a></li>
  <li><a href="#receive">Receive</a></li>
  <li><a href="#recover">Recover</a></li>
  <li><a href="#rem">Rem</a></li>
  <li><a href="#return">Return</a></li>
  <li><a href="#select">Select</a></li>
  <li><a href="#send">Send</a></li>
  <li><a href="#setfield">SetField</a></li>
  <li><a href="#setmap">SetMap</a></li>
  <li><a href="#setslice">SetSlice</a></li>
  <li><a href="#setvar">SetVar</a></li>
  <li><a href="#shl">Shl</a></li>
  <li><a href="#show">Show</a></li>
  <li><a href="#shr">Shr</a></li>
  <li><a href="#slice">Slice</a></li>
  <li><a href="#sub">Sub</a></li>
  <li><a href="#subinv">SubInv</a></li>
  <li><a href="#tailcall">TailCall</a></li>
  <li><a href="#text">Text</a></li>
  <li><a href="#typify">Typify</a></li>
  <li><a href="#xor">Xor</a></li>
  <li><a href="#zero">Zero</a></li>
</ol>

### Add

The instruction `Add` sums two integers or two floats. `Add` has two forms.

The first form sums the operands addressed by `a` and `b` and stores the result in `c`, the type of operands is `int` or `float64`.

The second form adds the operand addressed by `b` and `c` and stores the result in `c`. `type` is the type of the operands.

`b` can be an integer constant between 0 and 255.

```go
Syntax:  Add a b c    ; description: c = a + b // int and float64 types
         Add type b c ; description: c += b
```

```go  
Example:  Add i12 8 i7
          Add f2 f9 f5
          Add uint8 i2 i6
          Add float32 50 f2
```

### Addr

The instruction `Addr` takes the address of a slice element or struct pointer field and stores it in `p`. For slices, the slice is addressed by `s` and the index of the element is addressed by `i`. For struct pointers, the struct pointer is addressed by `s` and the field index is addressed by `i`.    

```go
Syntax:  Addr s i p ; description: p = &s[i]
         Addr s i p ;              p = &s.f // where f if the field of s at index i
```

### And

The instruction `And` computes the bitwise AND of the operands addressed by `a` and `b` and stores the result in `c`.

```go
Syntax:  And a b c  ; description: c = a & b
```

### AndNot

The instruction `AndNot` computes the AND NOT (bit clear) of the operands addressed by `a` and `b` and stores the result in `c`.

```go
Syntax:  AndNot a b c ; description: c = a &^ b
```

### Append

The instruction `Append` appends values to the slice addressed by `s` and store the resulting slice in `s`. The appended values are the values in the registers starting from the register `start` and ending to the register `end`. The registers with the values to append are consecutive and their type depends on the type of the slice's element.

```go
Syntax:  Append start end s ; description: s = append(s, start, ..., end)
```

```go
Example:  Append s3 s5 g2 ; append the values in the s3, s4 and s5 registers 
```

### AppendSlice

The instruction `AppendSlice` appends the slice addressed by `s1` to the slice addressed by `s2` and stores the resulting slice in `s2`.  

```go
Syntax:  AppendSlice s1 s2 ; description: s2 = append(s2, s1...)
```

### Assert

The instruction `Assert` asserts that the value addressed by `x` is not `nil` and it is of type `T`.

```go
Syntax:  Assert x T v ; description: v, ok = x.(T)
```

If the type assertion holds, it stores the resulting value into the register `v` (the register type depends on the type `T`), sets the `ok` flag to `true` and skips the next instruction.

If the type assertion is false, it sets the `ok` flag to `false` and, if the next instruction is a `Panic`, does a run-time panic.

### Break

The instruction `Break` breaks a previous executed `Range` instruction at label `label`.

```go
Syntax:  Break label ; description: break label
```

### Call

The instruction `Call` calls the function, or macro, with name `name` or the function, or macro, addressed by `fn` with arguments addressed by `i`, `f`, `s` and `g`.

1. `i` is the first integer register of the integer parameters.
2. `f` is the first floating point register of the floating point parameters.
3. `s` is the first string register of the string parameters.
4. `g` is the first general register of the other parameters.

`i`, `f`, `s` and `g` may be the blank identifier if there are no such arguments.

```go
Syntax:  Call name i f s g ; description: name(...)
         Call (fn) i f s g ;              fn(...)
```

```go
Example: Call fmt.Printf _ _ s3 g2
         Call (g8) i2 _ s4 _
```

In a function call, consecutive registers are used to store the return values and the arguments to the function. For example for a call to the function `sum`:

```go
func sum(a, b int) int {
    return a + b
}
```

three consecutive integer register are used for the parameters. The first is reserved for the return parameter, the second and third for the parameters `a` and `b`. So the following example:

```go
s := sum(1, 2)
print(s)
```
could be compiled to:
```go
Move 1 i6          ; store a into the register i6
Move 2 i7          ; store b into the register i7
Call sum i5 _ _ _  ; call sum
Print i5           ; print the return value stored in the register i5 
```
where the register `i5` is reserved for the return parameter, the registers `i6` and `i7` respectively for the parameters `a` and `b`.

This little more complex example:
```go
split := strings.SplitN("a,b,c", ",", 2)
n := len(split)
```
could be compiled to:
```go
Move "a,b,c" s5                 ; store the string to split into s5
Move "," s6                     ; store the separator into s6
Move 2 i3                       ; store the count into i3
Call strings.SplitN i3 _ s5 g2  ; call strings.SplitN
Len g2 i2                       ; get the length of the returned slice
```

### Cap

The instruction `Cap` gets the cap of the slice or channel addresses by `s` and stores it in `c`. 

```go
Syntax:  Cap s c ; description: c = cap(s)
```

### Case

The instruction `Case` defines a new case that can be chosen by a subsequent `Select` instruction.

If the instruction `Select` chooses a `Send` case, it sends the value addressed by `v` to the channel `ch`.
If the instruction `Select` chooses a `Recv` case, it receives from the channel addressed by `ch` and stores the value in `v`.
If the instruction `Select` chooses a `Default` case, it does nothing.

See the [Select](#select) instruction for how to use `Case` with `Select`.

```go
Syntax:  Case Send v ch ; description: case ch <- v:
         Case Recv ch v ; description: case v = <-ch:
         Case Default   ; description: default:
``` 

### Close

The instruction `Close` closes the channel addressed by `ch`. Closing a closed channel or a `nil` channel causes a run-time panic.  

```go
Syntax:  Close ch ; description: close(ch)
```

### Complex64

The instruction `Complex64` assembles a `complex64` value from the `float32` values addressed by `re` and `im` and stores the resulting complex in `c`.

```go
Syntax:  Complex64 re im c ; description: c = complex(re, im)
```

### Complex128

The instruction `Complex128` assembles a `complex128` value from the `float64` values addressed by `re` and `im` and stores the resulting complex in `c`.

```go
Syntax:  Complex128 re im c ; description: c = complex(re, im)
```

### Continue

The instruction `Continue` begins the next iteration of the previous executed `Range` instruction at label `label`.  
 
```go
Syntax:  Continue label ; description: continue label
```

### ConvertNumber

The instruction `ConvertNumber` explicitly converts the numeric operand addressed by `x` from the kind `xKind` to the kind `yKind` and stores it in `y`.

```go
Syntax:  ConvertNumber x xKind yKind y
```

`xKind` and `yKind` can be `Int`, `Int8`, `Int16`, `Int32`, `Uint`, `Uint8`, `Uint16`, `Uint32`, `Float32`, `Float64`, `Complex64` or `Complex128`. `yKind` can also be `String` if `xKind` is a integer kind. 

```go
Example: ConvertNumber i7 Int32 Uint i4 
         ConvertNumber f2 Float64 Float32 f6
         ConvertNumber g5 Complex128 Complex64 g4
         ConvertNumber f11 Float32 Complex128 g2
         ConvertNumber i3 Int String s7
```

### ConvertSlice

The instruction `ConvertSlice` explicitly converts the slice of bytes or runes addressed by `x` to a value of type `string` and stores it in `y`.

```go
Syntax:  ConvertSlice x y ; description: y = string(x)
```

### ConvertString

The instruction `ConvertString` explicitly converts the string operand addressed by `x` to a slice of bytes or runes value of type `T` and stores it in `y`.

As a special case, if `T` is the format type `html`, `ConvertString` explicitly converts a value of type `markdown` to a value of type `html`.

```go
Syntax:  ConvertString x T y ; description: y = T(x)
```

```go
Example:  ConvertString s3 []byte g5
          ConvertString s8 []rune g2
          ConvertString s6 html s7
```

### Concat

The instruction `Concat` concatenates the string operands addressed by `s1` and `s2` and stores the resulting string in `s3`.

```go
Syntax:  Concat s1 s2 s3 ; description: s3 = s1 + s2
```

### Copy

The instruction `Copy` copies the elements of the slice addressed by `src` to the slice addressed by `dst`. If the operand `n` is not zero, `Copy` stores in the integer register addressed by `n` the number of elements copied .         

```go
Syntax:  Copy src n dst ; description: n = copy(dst, src)
```

### Defer

The instruction `Defer` defers the execution of the function addressed by `f`.        

```go
Syntax:  Defer f ; description: defer f() { ... }
```

TODO: to be completed.

### Delete

The instruction `Delete` removes from the map addressed by `m` the element with key addressed by `k`.

```go
  Syntax:  Delete m k ; description: delete(m, k)
```

### Div

The instruction `Div` divides two integers or two floats. `Div` has two forms.

The first form divides the operands addressed by `a` and the operand addressed by `b` and stores the result in `c`, the type of operands is `int` or `float64`.

The second form divides the operand addressed by `c` and the operand addressed by `b` and stores the result in `c`. `type` is the type of the operands.

`b` can be an integer constant between -128 and 127, excluding zero.

```go
Syntax:  Div a b c    ; description: c = a / b // int and float64 types
         Div type b c ; description: c /= b
```

```go
Example:  Div i10 4 i7
          Div f3 f4 f5
          Div int8 -6 i3
          Div float64 f2 f1
```

### Field

The instruction `Field` gets from the struct pointer referred by `s` the field at the index referred by `i` and stores its value in `v`.

```go
  Syntax:  Field s i v ; description: v = s.f // where f is the field of s at index i
```

### Func

The instruction `Func` combines a function literal declaration with a load of the function (a closure to be precise) into `fn`.

```go
Syntax:  Func fn ; description: fn = func(...) { ... }
             <function body>
``` 

```go
Example: Move 2 i1
         Func g3
             Move 2 i3
             Add i1 1 i2
             Return
         Call (g3) i1 _ _ _
```

See the [LoadFunc](#loadfunc) instruction for how to load a non-literal function.

### GetVar

The instruction `GetVar` gets (TODO).

```go
Syntax:  GetVar i c ; description:
```

### GetVarAddr

The instruction `GetVarAddr` gets (TODO).

```go
Syntax:  GetVarAddr i c ; description:
```

### Go

The instruction `Go` runs the following `Call` or `CallIndirect` instruction in a new goroutine.

```go
Syntax:  Go ; description: go
```

Example:

```go
Go
Call Serve i2 _ s1 _
```

### Goto

The instruction `Goto` transfers control to the instruction with the corresponding label within the same function.

```go
Syntax:  Goto label ; description: goto label
```

### If

The instruction `If` checks its condition and skips the next instruction if it is satisfied.

```go
Syntax:  If Zero a                   ; description: if a == 0
         If NotZero a                ;              if a != 0

         If Nil a                    ;              if a == nil
         If NotNil a                 ;              if a != nil

         If NilInterface a           ;              if a == nil // a has an interface type
         If NotNilInterface a        ;              if a != nil // a has an interface type

         If a Equal b                ;              if a == b
         If a NotEqual b             ;              if a != b
         If a Less b                 ;              if a < b
         If a LessEqual b            ;              if a <= b
         If a Greater b              ;              if a > b
         If a GreaterEqual b         ;              if a >= b

         If a LenEqual b             ;              if len(a) == b
         If a LenNotEqual b          ;              if len(a) != b
         If a LenLess b              ;              if len(a) < b
         If a LenLessEqual b         ;              if len(a) <= b
         If a LenGreater b           ;              if len(a) > b
         If a LenGreaterEqual b      ;              if len(a) >= b

         If a ContainsSubstring b    ;              if a contains b   // as substring
         If a ContainsRune b         ;              if a contains b   // as rune
         If a ContainsElement b      ;              if a contains b   // as element
         If a ContainsKey b          ;              if a contains b   // as key
         If a ContainsNil            ;              if a contains nil // as element or key

         If a NotContainsSubstring b ;              if a not contains b   // as substring
         If a NotContainsRune b      ;              if a not contains b   // as rune
         If a NotContainsElement b   ;              if a not contains b   // as element
         If a NotContainsKey b       ;              if a not contains b   // as key
         If a NotContainsNil         ;              if a not contains nil // as element or key

         If OK                       ;              if ok  // ok is the ok flag
         If NotOK                    ;              if !ok // ok is the ok flag
```


### Index

The instruction `Index` gets, from the slice or string referred by `s`, the element at the index referred by `i` and stores its value in `v`.

```go
Syntax:  Index s i v ; description: v = s[i]
```

### Len

The instruction `Len` gets the length of the slice, string or channel addressed by `s` and stores it in `n`. 

```go
Syntax:  Len s n ; description: n = len(s)
```


### Load

The instruction `Load` loads a value `v` and stores it by copy into `dst`.

```go
Syntax:  Load v dst ; description: dst = v
```

```go
Example: Load 64 i3
         Load 3.14 f7
         Load "a" s1
         Load [2]int{} g3
         Load nil g5
```

### LoadFunc

The instruction `LoadFunc` loads the function with name `name` into `fn`.

```go
Syntax:  LoadFunc name fn ; description: fn = name
```

```go
Example: LoadFunc strings.HasPrefix g2
         Call (g2) i1 _ s3 _
```

See the [Func](#func) instruction for how to declare and load a function literal.

### MakeArray

The instruction `MakeArray` makes an array with type `T` and stores it in `a`.

```go
Syntax: MakeArray T a ; description: var a T
```

### MakeChan

The instruction `MakeChan` makes a channel with type `T` and buffer size addressed by `s` and stores it in `ch`.

```go
Syntax:  MakeChan T s ch ; description: ch = make(T, s)
```

### MakeMap

The instruction `MakeMap` makes a map with type `T` and size addressed by `s` and stores it in `m`.

```go
Syntax:  MakeMap T s m ; description: m = make(T, s)
```

### MakeSlice

The instruction `MakeSlice` makes a slice with element type `T`, length addressed by `n`, cap addressed by `c` and stores it in `s`.

```go
Syntax:  MakeSlice T n c s ; description: s = make(T, n, c)
```

### MakeStruct

The instruction `MakeStruct` makes a struct with type `T` and stores it in `s`.

```go
Syntax: MakeStruct T s ; description: var s T
```

### MapIndex

The instruction `MapIndex` gets from the map addressed by `m` the element with key addressed by `k` and stores its value in `v`. Also, it sets the `ok` flag to `true` if the key exists or `false` if key does not exist.

```go
Syntax:  MapIndex m k v ; description: v = m[k]
```

### MethodValue

The instruction `MethodValue` retrieves the method of the value addressed by `s`, named `n`, and stores the corresponding function value into `d`. The function value stored in `d` will always use the value addressed by `s` as the receiver. A call to this function does not require the receiver as an argument.

```go
Syntax:  MethodValue s n d ; description: d = s.n
```

### Move

The instruction `Move` copies the value addressed by `s` to `d`.

```go
Syntax:  Move s d ; description: d = s
```

### Mul

The instruction `Mul` multiplies two integers or two floats. `Mul` has two forms.

The first form multiplies the operands addressed by `a` and the operand addressed by `b` and stores the result in `c`, the type of operands is `int` or `float64`.

The second form multiplies the operand addressed by `c` and the operand addressed by `b` and stores the result in `c`. `type` is the type of the operands.

`b` can be an integer constant between -128 and 127.

```go
Syntax:  Mul a b c    ; description: c = a * b // int and float64 types
         Mul type b c ; description: c *= b
```

```go
Example:  Mul i2 i4 i3
          Mul f9 12 f2
          Mul uint16 i3 i1
          Mul float32 f5 f6
```

### Neg

The instruction `Neg` negates the operand addresses by `b`, with an integer or floating point type, and stores the result into `c`.

If `b` does not have `int` or `float64` type, `type` indicates its type.

```go
Syntax:  Neg b c      ; description: c = -b // for int and float64 types
         Neg type b c ; description: c = -b // for the other integer and floating point types
```

```go
Example:  Neg f2 f3
          Neg i4 i4
          Neg float32 f5 f8
```

### New

The instruction `New` allocates storage for a variable of type `T` and stores a pointer to this variable in `v`.

```go
Syntax:  New T v ; description: v = new(T)
```

### NotZero

The instruction `NotZero` checks if the value addressed by the register `b` is not a zero value; if so, then stores `1` in the integer register `c`, otherwise stores `0`.

As special cases, if the value addressed by `b` is an empty slice or an empty
channel the instruction `NotZero` stores `0`.

Moreover, if the value addressed by `b` has type interface and it is not `nil`,
the instruction `NotZero` evaluates its dynamic value.

```
Syntax: NotZero b c ; description: c = 1 if b is not the zero value for its type
                                   c = 0 otherwise
```

```
Example:  NotZero i1 i2
          NotZero s1 i3
```

See also the [Zero](#zero) instruction.

### Or

The instruction `Or` computes the bitwise OR of the operands addressed by `a` and `b` and stores the result in `c`.

```go
Syntax:  Or a b c ; description: c = a | b
```

### Panic

The instruction `Panic` panics with the value addressed by `v`.

```go
Syntax:  Panic v ; description: panic(v)
```

### Print

The instruction `Print` formats the operand addressed by `v` as does the built-in `print` and writes the result to standard error.  

```go
Syntax:  Print v ; description: print(v)
```

### Range

The instruction `Range` does an iteration through the entries of a slice, string, map or channel.
1. For a slice, the iteration is on the slice addressed by `s` and the instruction `Range` stores the iteration index in `i` and the element in `e`.     
2. For a string, the iteration is on the string addressed by `s` and the instruction `Range` stores the iteration index in `i` and the rune in `r`.
3. For a map, the iteration is on the map addressed by `m` and the instruction `Range` stores the key in `k` and the value in `v`.
4. For a channel, the iteration is on the values received on the channel addressed by `ch` and the instruction `Range` stores the value in `v`.

If there are no elements to iterate over, it sets the `vm.ok` flag to `false`, otherwise it sets the flag to `true`.

The instruction that follow `Range` is executed only when there are no more values to iterate over, during the iteration this instruction is skipped. 

The second and third operands can be the blank identifier.

```go
Syntax:  Range s i e ; description: for i, e = range s // for a slice
         Range s i r ; description: for i, r = range s // for a string 
         Range m k v ; description: for k, v = range m // for a map 
         Range ch v  ; description: for v = range ch   // for a channel
```

The instruction `Range` is used in combination to the instructions `Goto`, `Continue` and `Break` to implement a `for range` statement.

For example a range over the first five elements of a slice:

```go
for i, e := range s {
    if i == 5 {
        break
    }
    print(i)
    print(e)
}
return
```

could be compiled to assembly:

```go
1: Range g2 i1 s7
   Goto 3
   If i1 Equal 5
   Goto 2
   Break 1
2: Print i1
   Print s7
   Continue 1
3: Return
```

### RealImag

The instruction `RealImag` extracts the real and imaginary parts of the complex number addressed by `c` and stores the real part in `re` and the imaginary part in `im`.

```go
Syntax:  RealImag c re im ; description: re, im = real(c), imag(c)
```

### Receive

The instruction `Receive` receive a value from the channel addressed by `ch` and stores the value in `v`. Also it sets the `ok` flag to `true` if the value received was delivered by a successful send operation, or `false` if a zero value was stored in `v` because the channel is closed and empty. 

```go
Syntax:  Receive ch v ; description: v, ok = <-ch
```

### Recover

The instruction `Recover` recovers a panicking goroutine and stores in `v` the value passed to the call of panic. If the goroutine is not panicking or `Recover` was not executed directly by a deferred function it stores `nil` in `v`. `v` can be the blank identifier.

```go
Syntax:  Recover v ; description: v = recover()
```

As a special case, compiling the statement `defer recover()` the following assembly line is generated:

```go
Recover DownTheStack v
```

### Rem

The instruction `Rem` computes the remainder of two integers. `Rem` has two forms.

The first form computes the remainder of the operands addressed by `a` and the operand addressed by `b` and stores the result in `c`, the type of operands is `int`.

The second form computes the remainder of the operand addressed by `c` and the operand addressed by `b` and stores the result in `c`. `type` is the type of the operands.

`b` can be an integer constant between 1 and 255.

```go
Syntax:  Rem a b c    ; description: c = a % b // int type
         Rem type b c ; description: c %= b
```

```go
Example:  Rem i10 12 i2
          Rem uint32 i7 i8
```

### Return 

The instruction `Return` returns from the current running function. Any functions deferred with the `Defer` instruction in the current function are executed before returning.

```go
Syntax:  Return ; description: return
```

### Select 

The instruction `Select` selects a case from one of the cases defined by previously executed `Case` instructions, executes a send or receive depending on the case, clears all the cases previously defined and jumps to the instruction that follows the selected `Case` instruction.   

```go
Syntax:  Select ; description: select
```

The instruction `Select` is used in combination to the instruction [Case](#case) to implement a `select` statement.

For example the select statement:

```go
select {
case a <-tick:
	print(a)
case b <-boom:
	print(b)
default:
	print(3)
}
return
```

could be compiled to this assembly:

```go
   Case Recv i5 g2  ; case a <-tick:
   Goto 1
   Case Recv s3 g1  ; case b <-boom:
   Goto 2
   Case Default     ; default
   Goto 3
   Select           ; select
1: Print i5         ; print(a)
   Goto 4
2: Print s3         ; print(b)
   Goto 4
3: Print 3          ; print(3)
   Goto 4
4: Return           ; return
```

### Send

The instruction `Send` sends the value addressed by `v` on the channel `ch`. 

```go
Syntax:  Send v ch ; description: ch <- v
```

### SetField

The instruction `SetField` stores the value addressed by `v` into the field with index `i` of the struct addressed by `s`.

```go
Syntax:  SetField v s i ; description: s.f = v // where f is the field of s at index i
```

### SetMap

The instruction `SetMap` sets the value of the map addressed by `m` and indexed by the key `k` with the value addressed by `v`.

```go
Syntax:  SetMap v m k ; description: m[k] = v
```

### SetSlice

The instruction `SetSlice` sets the value of the slice addressed by `s` at the index `i` with the value addressed by `v`.

```go
Syntax:  SetSlice v s i ; description: s[i] = v
```

### SetVar

The instruction `SetVar` sets the value of the global or closure variable at index `i` with the value addressed by `v`.

```go
Syntax:  SetVar v i ; description: vars[i] = v
```

### Shl

The instruction `Shl` computes the left shift of an integer. `Shl` has two forms.

The first form computes the left shift of the operands addressed by `a` with shift count `n` and stores the result in `c`, the type of operands `a` and `c` is `int`.

The second form computes the left shift of the operand addressed by `c` with shift count `n` and stores the result in `c`. `type` is the type of `c`.

`n` can be an integer constant between 0 and 255.

```go
Syntax:  Shl a n c    ; description: c = a << n // int type
         Shl type n c ; description: c <<= n
```

```go
Example:  Shl i3 5 i2
          Shl uint16 i5 i6
```

### Show

The instruction `Show` formats the value addresses by `v` based on the context `ctx` and writes the result to the template out writer.


```go
Syntax:  Show T v (ctx); description: out.Write(format(v, ctx))
```

### Shr

The instruction `Shr` computes the right shift of an integer. `Shr` has two forms.

The first form computes the right shift of the operands addressed by `a` with shift count `n` and stores the result in `c`, the type of operands `a` and `c` is `int`.

The second form computes the right shift of the operand addressed by `c` with shift count `n` and stores the result in `c`. `type` is the type of `c`.

`n` can be an integer constant between 0 and 255.

```go
Syntax:  Shr a n c    ; description: c = a >> n // int type
         Shr type n c ; description: c >>= n
```

```go
Example:  Shr i7 2 i3
          Shr int8 i9 i3
```

### Slice

The instruction `Slice` slices the slice or the string addressed by `s1` and stores the resulting slice or string into `s2`. The values addressed by `low`, `high` and `max` are the low, high and max indices. `max` is optional for slices and it is never present for strings.

```go
Syntax:  Slice s1 low high s2     ; description: s2 = s1[low:high]
         Slice s1 low high max s2 ;              s2 = s1[low:high:max]
```

### Sub

The instruction `Sub` subtracts two integers or two floats. `Sub` has two forms.

The first form subtracts the operands addressed by `b` from the operand addressed by `a` and stores the result in `c`, the type of operands is `int` or `float64`.

The second form subtracts the operand addressed by `b` from the operand addressed by `c` and stores the result in `c`. `type` is the type of the operands.

`b` can be an integer constant between 0 and 255.

```go
Syntax:  Sub a b c    ; description: c = a - b // int and float64 types
         Sub type b c ; description: c -= b
```

```go
Example:  Sub i3 i9 i2
          Sub f8 23 f2
          Sub int64 i7 i1
          Sub float64 f14 f21
```

### SubInv

The instruction `SubInv` subtracts two integers or two floats. `SubInv` has two forms.

The first form subtracts the operand addressed by `a` from the operand addressed by `b` and stores the result in `c`, the type of operands is `int` or `float64`.

The second form subtracts the constant addressed by `b` from the operand addressed by `c` and stores the result in `c`. `type` is the type of the operands.

The operand `b` is an integer constant between -128 and 127.

```go
Syntax:  SubInv a b c    ; description: c = b - a // int and float64 types
         SubInv type b c ; description: c = b - c
```

```go  
Example:  SubInv i3 i9 i2
          SubInv f8 -23 f2
          SubInv int64 i7 i1
          SubInv float64 f14 f21
```

### TailCall

TODO

### Text

The instruction `Text` writes the text `txt` to the template out writer.

```go
Syntax:  Text txt ; description: out.Write(txt)
```

```go
Example: Text "<h1>About us</h1>"
```

### Typify

The instruction `Typify` gets the value addressed by `v1`, convert it to type `T` and stores the resulting value into `v2`.

```go
Syntax:  Typify T v1 v2 ; description: var v2 T = v1
```

### Xor

The instruction `Xor` computes the bitwise XOR of the operands addressed by `a` and `b` and stores the result in `c`.

```go
Syntax:  Xor a b c ; description: c = a ^ b
```

### Zero

The instruction `Zero` checks if the value addressed by the register `b` is a zero value; if so, then stores `1` in the integer register `c`, otherwise stores `0`.

As special cases, if the value addressed by `b` is an empty slice or an empty
channel the instruction `Zero` stores `1`.

Moreover, if the value addressed by `b` has type interface and it is not `nil`,
the instruction `Zero` evaluates its dynamic value.

```
Syntax: Zero b c ; description: c = 1 if b is the zero value for its type
                                c = 0 otherwise
```

```
Example:  Zero i1 i2
          Zero s1 i3
```

See also the [NotZero](#notzero) instruction.

{% end raw code %}
