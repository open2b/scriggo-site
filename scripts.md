{% extends "/layouts/doc.html" %}
{% macro Title string %}Scripts{% end %}
{% Article %}

# Scripts

Scripts a scripting language based on Go but with the following differences that make them more suitable for a scripting context:
 
* files can start with a shebang: `#!/usr/bin/env scriggo`
* there is no `package main` declaration
* there is no `func main` declaration
* execution exits with the script’s exit code
* there is a new shell command literal
* there are a new `run` and `exit` builtin functions

This is a script:

```go
println("Hello Word!")
```

and this is another script with the `run` builtin function with a shell command literal as argument:

```go
hello := "Hello Word!"
run $ echo #{hello}
```

The following is a more complete example of script:

```go
#!/usr/bin/env scriggo

import (
    "fmt"
    "os"
)

const limit = 10000

fi, err := os.Open("requests.log")
if err != nil {
    exit(1)
}
defer fi.Close()

cat := $ cat "a.txt" | grep "n" | sort 
err = cat.Run()
if err != nil {
    exit(err)
}

repo := "github.com/open2b/scriggo"

_, err = run $ git checkout #{repo}
if err != nil {
    exit(1)
}

_, err := run(checkout)
if err != nil {
    exit(err)
}

_, err = run $ while [[ $i -le 3 ]]; do echo $i; i=$((i+1)); done
if err != nil {
    exit(1)
}

run $ alias | grep -i git | sort -R | head -10

out, err = run("docker", "run", "nmp")

apk := $ apk --no-cache add \
    bash \
    btrfs-progs-dev \
    build-base \
    curl \
    lvm2-dev \
    jq
apk.Stdin = in
err = apk.Run()

_, err = run $ git branch dev
if err != nil {
    exit(err)
}

_, err = run(cmd, checkout)

if len(requests) > limit {
    fmt.Println("How many requests!")
}

rx``

```

A Scriggo interpreter runs scripts with performances comparable or superior to those of a programming language like Python.  

## Make a script interpreter

A script interpreter is a program that can run scripts. It can be a command line executable or a web application, with
an editor to code and run scripts, or a server that executes scripts for routine operations and plugins.

Creating an interpreter starts from writing a [Scriggofile](scriggofile) that defines which packages a script can import and
which globals will be defined for the executed scripts.

The [Scriggofile](scriggofile) is then saved in a module directory. The module's `go.mod` will specify the versions of
packages's modules imported in the [Scriggofile](scriggofile).

And finally the `scriggo` command is executed to make an executable of the interpreter or to generate a package importer
to use in an existing application to load and run scripts. 

### A simple interpreter

To create a simple interpreter, as a command line executable, that runs scripts that can import the packages of the Go standard library, start
creating a new Go module:

```
$ mkdir foo
$ cd foo
$ go mod init foo
```

Create a file named `Scriggofile` in the module directory `foo` with this content:

```
TARGET SCRIPTS
IMPORT STANDARD LIBRARY
```

Build the interpreter from the module directory with the [scriggo command](/scriggo-command):

```
$ scriggo build
```

It builds the interpreter from the instructions in `Scriggofile` and creates the `foo`, or `foo.exe`, executable in
the current directory.

Try to run a script with the new `foo` interpreter. Create a file named `script.sg` with the following content: 

```go
import "fmt"
fmt.Println("hello world!")
```

and run the script: 

```
$ ./foo script.sg
hello world!
```

or on Windows:

```
$ foo.exe script.sgo
hello world!
```

### A more complex interpreter

Starting fro  the  

   


 


The following is an example of [Scriggofile](scriggofile) for scripts to create and manage containers:

```
# Scriggofile for create and manage containers
TARGET SCRIPTS
IMPORT context INCLUDING Background NOT CAPITALIZED
IMPORT fmt
IMPORT github.com/docker/docker/client AS main INCLUDING NewClientWithOpts NOT CAPITALIZED
IMPORT github.com/docker/docker/client AS main INCLUDING FromEnv Image
IMPORT github.com/docker/docker/types AS main INCLUDING ContainerListOptions 
```

The following is an example of a script that can be executed with an interpreter made from [Scriggofile](scriggofile):

```go
import (
    "context"
    "fmt"
)

cli, err := newClientWithOpts(FromEnv)
if err != nil {
    panic(err)
}

containers, err := cli.ContainerList(background(), ContainerListOptions{})
if err != nil {
    panic(err)
}

for _, container := range containers {
    fmt.Printf("%s %s\n", container.ID[:10], Image)
}
```



* a 
Building of a script interpreter starts from writing a Scriggofile for the interpreter and running one of the following commands:

* `scriggo build` or `scriggo install` to create a standalone interpreter.
* `scriggo embed` to generate a package importer to use with an existing program.

### Scriggofile for a script interpreter


  