{% raw %}

<!-- TODO: uniformare {% end .. %} -->

<!-- <style>
pre.example hr {
    border-width: 3px;
    margin-top: 25px;
}
pre.example {
    border-style: solid;
    border-width: 2px;
    border-radius: 2px;
    border-color: #A8E6FF;
    background-color: #ECFAFF;
}
</style> -->


# Scriggo Virtual Machine
{: .no_toc}

La macchina virtuale di Scriggo è una macchina virtuale a registri che esegue il bytecode generato dal compilatore di Scriggo a partire da un programma, uno script o un template di Scriggo. 


1. TOC
{:toc}

## Registri

La macchina virtuale ha 500 registri suddivisi nei quattro gruppi int, float, string e generic:

* `i1`, `i2`, ..., `i125`: 125 registri di tipo `int64`
* `f1`, `f2`, ..., `f125`: 125 registri di tipo `float64`
* `s1`, `s2`, ..., `s125`: 125 registri di tipo `string`
* `g1`, `g2`, ..., `g125`: 125 registri di tipo `interface{}`

Nei registri int vengono memorizzate tutte le variabili locali con tipo base intero o booleano:

```go
// variabili locali memorizzate nei registri int
var a int
var b uint32
var c rune
var d time.Duration = 5 * time.Second
var e bool // false è memorizzato nei registri con int64(0)
var f MyBool = true // true è memorizzato nei registri con int64(1)
```
Nei registri float vengono memorizzate tutte le variabili locali con tipo base floating point:

```go
// variabili locali memorizzate nei registri float
var a float32
var b float64
var d MyFloat = 3.2
```
Nei registri string vengono memorizzate tutte le variabili locali con tipo base string:

```go
// variabili locali memorizzate nei registri string
var a string
var b MyString = "hello"
```
Nei registri generic vengono memorizzate tutte le altre variabili locali:
```go
// variabili locali memorizzate nei registri generic
var a []int
var b map[string]string
var c interface{}
var d time.Time
var e *S
```

## Istruzioni

La macchina virtuale ha 120 istruzioni macchina. Il formato più comune per le istruzioni prevede una lunghezza di 32 bit:

| Opcode     | Operando A  | Operando B  | Operando C |
| ------------- | ------------- | ------------- | ------------- |
| 8 bit      | 8 bit       | 8 bit       | 8 bit      |

Ad esempio:

| OpAddInt32 | 2 | 3 | 7 |
| ------- | --- | --- | --- |

somma i valori nei registri `i2` e `i3`, trattandoli come `int32`, e memorizza il risultato nel registro `i7`.

Nel bytecode (il codice disassemblato) diverse istruzioni macchina vengono rappresentate come una singola istruzione e si possono distinguere dal tipo dei registri su cui operano. Ad esempio le istruzioni macchina `OpAddInt64` e `OpAddFloat64` vengono rappresentate con la singola istruzione `Add`:

```go
Add i2 i3 i7 // istruzione macchina OpAddInt64 2 3 7
Add f5 f1 f3 // istruzione macchina OpAddFloat64 5 1 3
```

Alcune istruzioni hanno una corrispettiva istruzione negata che tratta il secondo operando come immediato, ad esempio:

| -OpAddInt32 | 2 | 3 | 7 |
| -------- | --- | --- | --- |

somma il valore nel registro `i2` con il valore costante `3`, trattandoli come `int32`, e memorizza il risultato nel registro `i7`.

La precedente istruzione macchina `-OpAddInt32` è rappresentata in bytecode nel seguente modo:

```go
Add32 i2 3 i7 // istruzione macchina -OpAddInt32 2 3 7
```

Si noti che in questo caso il secondo operando è la costante `3` e non il registro `i3`.

### Add, Add8, Add16 and Add32

The `Add` instructions sum the operands addressed by `i1` and `i2` (or `f1` and `f2`) and stores the result in `i3` (or `f3`).

```go
Syntax:  Add   i1 i2 i3   ; Function:  i3 = i1 + i2
         Add   f1 f2 f3   ;            f3 = f1 + f2
         Add8  i1 i2 i3   ;            i3 = i1 + i2
         Add16 i1 i2 i3   ;            i3 = i1 + i2
         Add32 i1 i2 i3   ;            i3 = i1 + i2
         Add32 f1 f2 f3   ;            f3 = f1 + f2
```
The second operand can be an integer constant `c` between -127 and 126:
```go
Syntax:  Add   i1 c i3   ; Function:  i3 = i1 + c
         Add   f1 c f3   ;            f3 = f1 + c
         Add8  i1 c i3   ;            i3 = i1 + c
         Add16 i1 c i3   ;            i3 = i1 + c
         Add32 i1 c i3   ;            i3 = i1 + c
         Add32 f1 c f3   ;            f3 = f1 + c
```

### Addr

The instruction `Addr` takes the address of a slice element, array element, struct field or pointer to struct field and store the value in `dst`. 

```go
Syntax:  Addr src i dst   ; Function:  dst = &src[i1]
         Addr src i dst   ;            dst = &src.f   // i is the index of the field src.f
```

### Alloc

The instruction `Alloc` allocates the memory used by the next instruction or allocates a fixed bytes of memory. If the memory is not limited, `Alloc` is a no-op.

```go
Syntax:  Alloc     ; Function: allocates memory used by the next instruction
         Alloc n   ;           allocates n bytes of memory
```
Allocation of memory is necessary only to control the memory usage during the execution. The compiler adds `Alloc` instructions if the `LimitMemorySize` option is used when loading a program, script or template.

### And

The instruction `And` computes the bitwise AND of the two operands addressed by i1 and i2 and stores the result in the i3 operand.

```go
Syntax:  And i1 i2 i3  ; Function: i3 = i1 & i2
```

### AndNot

The instruction `AndNot` computes the AND NOT (bit clear) of the two operands addressed by i1 and i2 and stores the result in the i3 operand.

```go
Syntax:  And i1 i2 i3  ; Function: i3 = i1 &^ i2
```

### Append

The instruction `Append` appends `length` values to the slice in the register `dst` and store the resulting slice in the same register `dst`. The appended values are the values starting from the register at index `start` and ending to the register at index `start+length-1`. The type of the values register depends on the type of the slice element.

`start` and `length` are integer constant between 1 and 126.

```go
Syntax:  Append start length dst  ; Function: dst = append(dst, regs[start:start+length]...)
```

### AppendSlice

The instruction `AppendSlice` appends the slice in the register `src` to the slice in the register `dst` and store the resulting slice in the register `dst`. The two slices have the same type.  

```go
Syntax:  AppendSlice src dst  ; Function: dst = append(dst, src...)
```

### Assert

The instruction `Assert` does the type assertion `s.(T)` where `s` is the value of the general register at index `src` and `T` is the type in the type register at index `typ`.

```go
Syntax:  Assert src typ dst  ; Function: dst, ok = s.(T)
```

If the type assertion successes, it stores the the resulting value into the register at index `dst` (the register type depends on the type `T`), sets the VM field `ok` to `true` and skips the next instruction.

If the type assertion fails, sets the VM field `ok` to `false` and, if the next instruction is a `Panic`, does a run-time panic. TODO: document the `Panic` operand.     

### Bind

TODO: to be removed?

### Break

The instruction `Break` breaks a previous executed `Range` instruction at label `label`.

```go
Syntax:  Break label  ; Function: break label
```

### Call

The instruction `Call` calls the function at `current.Functions[f]` where `current` is the current executed function. 

```go
Syntax:  Call f  ; Function: f(...)
```

The instruction `Call` is 8 bytes long and the last 4 bytes stores... (TODO)

### CallIndirect

TODO

### CallPredefined

TODO

### Cap

The instruction `Cap` gets the cap of the slice at the general register `src` and stores it in the integer register `dst`. 

```go
Syntax:  Cap src dst  ; Function: dst = cap(src)
```

### Case

TODO

### Close

The instruction `Close` closes the channel at the general register `ch`. Closing a closed channel or a `nil` channel causes a run-time panic.  

```go
Syntax:  Close ch  ; Function: close(ch)
```

### Complex64

The instruction `Complex64` assembles a `complex64` value from the `float32` values in the float registers `re` and `im` and stores the resulting complex in the general register `cmpx`.

```go
Syntax:  Complex64 re im cmpx  ; Function: cmpx = complex(re, im)
```

### Complex128

The instruction `Complex128` assembles a `complex128` value from the `float64` values in the float registers `re` and `im` and stores the resulting complex in the general register `cmpx`.

```go
Syntax:  Complex128 re im cmpx  ; Function: cmpx = complex(re, im)
```

### Continue

The instruction `Continue` begins the next iteration of the previous executed `Range` instruction at label `label`.  
 
```go
Syntax:  Continue label  ; Function: continue label
```

### ConvertGeneral

The instruction `ConvertGeneral` converts the value at general register `src` to the type `T` at `current.Types[typ]`, where `current` is the current running function. The converted value is stored in the general register `dst` or, if the converted value is a string, in the string register `dst`.     

```go
Syntax:  ConvertGeneral src typ dst ; Function: dst = T(src)
```

### ConvertInt

The instruction `ConvertInt` converts the signed integer value at int register `src` to the type `T` at `current.Types[typ]`, where `current` is the current running function. The converted value is stored in the int, float or string register `dst` depending of the type `T`.     

```go
Syntax:  ConvertInt src typ dst ; Function: dst = T(src)
```

### ConvertUint

The instruction `ConvertInt` converts the unsigned integer value at int register `src` to the type `T` at `current.Types[typ]`, where `current` is the current running function. The converted value is stored in the int, float or string register `dst` depending of the type `T`.     

```go
Syntax:  ConvertUint src typ dst ; Function: dst = T(src)
```

### ConvertFloat

The instruction `ConvertFloat` converts the floating point value at float register `src` to the type `T` at `current.Types[typ]`, where `current` is the current running function. The converted value is stored in the int or float register `dst` depending of the type `T`.     

```go
Syntax:  ConvertFloat src typ dst ; Function: dst = T(src)
```

### ConvertString

The instruction `ConvertString` converts the string value at string register `src` to the type `T` at `current.Types[typ]`, where `current` is the current running function. The converted value is stored in the general register `dst`.

```go
Syntax:  ConvertString src typ dst ; Function: dst = T(src)
```

### Concat

The instruction `Concat` concatenates the string values at string registers `src1` and `src2` and stores the resulting string at register `dst`.

```go
Syntax:  Concat src1 src2 dst ; Function: dst = src1 + src2
```

### Copy

The instruction `Copy` copies the elements of the slice at general register `src` to the slice at general register `dst`. If `n` is not zero, `Copy` store in the integer register `n` the number of elements copied.         

```go
Syntax:  Copy src n dst ; Function: n = copy(dst, src)
```

### Defer

The instruction `Defer` defers the execution of the function at general register `f`.         

```go
Syntax:  Defer f ; Function: defer f() { ... }
```

TODO: to be completed.

### Delete

The instruction `Delete` removes the element with key in the general register `k` from the map in the general register `m`.

```go
  Syntax:  Delete m k; Function: delete(m, k)
```

### Div, Div8, Div16, Div32

The `Div` instructions divide the operands at int registers `i1` and `i2` ( or float register `f1` and `f2` ) and stores the result in the int register `i3` ( or float register `f3` ). Dividing by zero causes a run-time error.

```go
Syntax:  Div   i1 i2 i3   ; Function:  i3 = i1 / i2
         Div   f1 f2 f3   ;            i3 = i1 / i2
         Div8  i1 i2 i3   ;            i3 = i1 / i2
         Div16 i1 i2 i3   ;            i3 = i1 / i2
         Div32 i1 i2 i3   ;            i3 = i1 / i2
         Div32 f1 f2 f3   ;            f3 = f1 / f2
```
The second operand can be an integer constant `c` between -127 and 126, excluding 0.
```go
Syntax:  Div   i1 c i3   ; Function:  i3 = i1 / c
         Div   f1 c f3   ;            f3 = f1 / c
         Div8  i1 c i3   ;            i3 = i1 / c
         Div16 i1 c i3   ;            i3 = i1 / c
         Div32 i1 c i3   ;            i3 = i1 / c
         Div32 f1 c f3   ;            f3 = f1 / c
```

### DivU, DivU8, DivU16, DivU32

The `DivU` instructions divide the unsigned integer operands at int registers `i1` and `i2` and stores the result in the int register `i3`. Dividing by zero causes a run-time error.   

```go
Syntax:  DivU   i1 i2 i3   ; Function:  i3 = i1 / i2
         DivU8  i1 i2 i3   ;            i3 = i1 / i2
         DivU16 i1 i2 i3   ;            i3 = i1 / i2
         DivU32 i1 i2 i3   ;            i3 = i1 / i2
```
The second operand can be an integer constant `c` between 1 and 256:
```go
Syntax:  DivU   i1 c i3   ; Function:  i3 = i1 / c
         DivU8  i1 c i3   ;            i3 = i1 / c
         DivU16 i1 c i3   ;            i3 = i1 / c
         DivU32 i1 c i3   ;            i3 = i1 / c
```

### Field

The instruction `Field` gets the value the field at index `idx` of the pointer to struct in the global register `src` and stores the value into the register `dst` which type depends on the field's type.

```go
  Syntax:  Field src idx dst; Function: dst = src.F // where F is the field at index idx
```

{% endraw %}