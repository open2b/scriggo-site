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

The `Add` instructions sums the operands addressed by `i1` and `i2` (or `f1` and `f2`) and stores the result in `i3` (or `f3`).

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

{% endraw %}