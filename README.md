# Mural

Mural is a CLI library write in go allow you to implement CLI logic on your go application.
Mural implement multiple way to handle CLI logic in your application.
Must usefull feature of Mural is **Description as Code** pattern. Describe your command and tag as a simple documentation (as help common command can do) and Mural, from the description, can construct you command logic.
Any code need, just description!!!

## Installation

Make only a go get command:

```bash
go get -u github.com/jakofys/mural
```

And it ready to be used

## Features

- [ ] Description as Code
    - [ ] Command description
    - [ ] Flag description
    - [ ] Variable
    - [ ] Environment variable
- [ ] Decode from args and flags
- [ ] Command line pipeline from file
- [ ] Code own CLI
- [ ] Extensible flag logic
- [ ] Subcommand supports
- [ ] Autocompletion command line

## Use

### Create command

Initialize basic information about CLI application as:
Mural can support semantic version, and is recommended to use version to enable to use `-v` or `--version` flag to know version of the CLI application. (see [semantic version standard](https://semver.org))

```go
cli,err := mural.NewCLI("myapp", "1.12.3-alpha+6f8ca4bfb2", "A simple application")
if err != nil{
    panic(err)
}
```

As see before, you can have some keyword flag already blocked by default, see keyword references [here](#keywords-references).

At this step, use of `-v` return given output on standard output:

```
myapp 1.12.3-alpha+6f8ca4bfb2 (arm64)
```

You can know add a new command handler.
You have to **implement** `mural.CommandHandler` interface to be capable of handle a command.

```go
mural.AddCommands(
    &Command{
        Name: "hello",
        Description: "This command get a hello world",
        Handler: HelloCommandHandler,
    }
)
```

Then implement `CommandHandler` to handle incomming command:

```go
func HelloCommandHandler (args *Args, flags *Flags)error{
    fmt.Print("Hellow world")
    return nil
}
```

>  `CommandHandlerFunc` is a function signature that implement `CommandHandler` interface

The `args` and `flags` input reflect all values use in command line. 
It allow to reach data using `index` or `tagname` according to how to describe your command.
It have `Decode` method that apply all value to a `struct` that use `mural` tagname. (see [here](#args-and-flags-decode) how it works).

Don't forget to specify to CLI to build is intern logic using:

```go
_, err := cli.Compile()
```

It needed to `Compile` the CLI that construct **schema** and defined all scenarios about command, subcommand, optional and required flags and args
(see [architecture]() of build).
It create a file named by default `<myapp.cli>` containing op-code for cli. It will be modify when CLI resources changed.
> Op-code file can be renamed using `cli.SetCachePath(mypath)`.

> This function compile given information, so after called `Compile`, you cannot neither adding or removing, neither changing ressources behaviour. <br>
Can returns error if duplicated resources found, none sense logic retrieve or missing informations. 

Finally for start running CLI, you can use multiple ways.
The first is the common way to handle input, can use standard input from `os.Args` value (see [os](https://pkg.go.dev/os) library)

<!-- Explain the signature and add new signature for stream input -->

```go
cli.Run(os.Args)
```

> It use `strings.Join` and pass the result to `cli.RunString`

Another way is to use stream the input as `os.Stdin` or another `os.Reader` can do using following methods:

```go
cli.RunStream(os.Stdin)
```

Mural use `EOL` as separator between command line, and a `EOF` stop running instruction.
With it, you can chain multiple command line in stream.

We can run this application to making:

```bash
./myapp hello # Output 'Hello world!'
```

## Description as Code

This feature allow developers to describe how work they command line, just describing alls command line.

With this way, you defined specification about your CLI tools.
Instead of coding, you have the choice to make it more readable as:

```go
cli.AddPatternCommand("hello", "This command get a hello world")    
```

We see here a command that named `hello` and his description. To implement this command, just bind it with your command handler:

```go
cli.BindCommand("hello", HelloCommandHandler)
```

And we have the same result as previous use case.

### Introduction

A **description** instance is a set of `@annotating` information.
A block start with type of information, and then, list description's content respecting data format.

### Command

A command has it own pattern, and it composed of:

- Command name
- Argument name (only string parameter)
- Flag name or tag
- Description (used in `help` command)

The structure of description pattern is simple, the first line is to define command line composition.

```go
cli.AddPatternCommand("command [arg1] [?arg2] --flag1 --flag2", "Command description for helping command")    
```

Like docker or kubernetes, supports subcommand for delegate data flow to another command handler.

```go
cli.AddPatternCommand("command subcommand [arg]", "Command description for helping command")
```

You can use only tag name for your flag declaration.
Why is possible ?
In Flag block, if a flag has it name or tag, Mural know how to bind them.

```go
cli.AddPatternCommand("command -f1", "Command description for helping command")
```

And a argument is not simple a location in command, it has name, default property, and take advantage of three features:

**Required argument** require you to indicate string argument:

`[arg]`

**Optional argument** allow you to not declare argument value on command line, without return error:

`[?arg]`

**Default argument value** can set default value to argument if it null:

`[arg='default value']`

**Argument as list** create a pattern that list string argument:

`[arg...]`

> Be carefull using **Argument as list**,  it has the same behaviour than `params...` golang syntax, except, by default is required if you not use **Optional argument**. So make this type as last argument of the command  else must panic.

> You can have multiple argument list seperate by command as `command subcommand [?arg2...]`

### Flag

A flag is a way to add behaviour metadata to create context to the command.
You can add flag definition as:

```go
cli.AddPatternFlag("flag | f1: string", "Flag one description as string")
```

Use it into tapped command line:

```shell
myapp hello --flag='flag text'
# same as
myapp hello --flag 'flag text'
# same as
myapp hello -f 'flag text'
```

### File description

You can also create a documentation file use specific language to descibe your entirely application with:

```go
cli, err := mural.From(io.Reader) // Output a cli importing file description
```
<!-- Stop 1 -->
```txt
@commands:
    hello
    > This command get a simple hello world!
```

We see here a command that named `hello` with a simple description at next line starting with operator `>`.

<!-- Stop 1 -->

Then in this content, we have:

```
@version: 1.12.3-pre
@name: myapp

@description:

    A simple application

@commands:

    command subcommand [arg1] [?arg2] --flag1 --flag2
    > Command description has coming soon

    command [arg1] nextcommand [arg2] -f1 -f2
    > Command for deleguate to next command value

@flags:

    flag1 | f1: string
    > Flag one description as string

    flag2 | f2: bool=true
    > Flag two description as boolean with true as default value
```

Operator `>` determine a description line.

## Keywords references

Mural use specific keyword corresponding to standard use and cannot be rewriting.
Here a list and explanation of they use:

- help flag `-h, --help`: return description of how use the CLI interface

example:

```txt
$ myapp --help

A simple application.

Usage:
  myapp [command]

Available Commands:
  hello       This command get a hello world

Flags:
  -f1, --flag1  string  Flag one description as string
  -f2, --flag2          Flag two description as string

Use "myapp [command] --help" for more information about a command.
```

And get with command help flag:

```txt
$ myapp command --help

A simple application.

Usage:
  myapp command [arg1] [subcommand]

Available Subcommands:
  subcommand       Command for deleguate to next command value

Flags:
  -f1, --flag1  string  Flag one description as string
  -f2, --flag2          Flag two description as string

Use "myapp [command] --help" for more information about a command.
```

- version flag `-v, --version`: show the version of the CLI, can be templated to add multiple version information

example: 
```txt
$ myapp --version
  __  __                         
 |  \/  |_   _  __ _ _ __  _ __  
 | |\/| | | | |/ _` | '_ \| '_ \ 
 | |  | | |_| | (_| | |_) | |_) |
 |_|  |_|\__, |\__,_| .__/| .__/ 
         |___/      |_|   |_|    

Version: 1.12.3-alpha+6f8ca4bfb2
OS: macOS 12.3
Arch: arm64
Device: Apple M1 Pro

Dependancies                    Version
------------------------------------------
api-app                         1.32.2-release
go                              1.19.3
```

- doc command `doc`:


## Architecture

Mural use two important pattern.

For file description as code, Mural treat sequentially (line by line) each information, it allowing to read correctly pattern and know if you respected syntax.

Another is the binding information from description (example in `flag`) allow to interpreting for example a flag tag without describe it in the command, but just to create a documentation of it. 

## Sources

- [Official console documentation](mural.dev) of Mural library available
- Go on [Golang package documentation](pkg.go.dev/github.com/jakofys/mural) to get package's profile.

## Reference

- [Cobra](github.com/spf13/cobra) inspiration
- [Description as code]() inspiration to [Terraform](terraform.io)
- [Command line interface (CLI)]() composition and schema
- [ASCII Art](https://patorjk.com/software/taag) generate ASCII Art from cli app name

## Author and contributor

- [Jacques COFIS](github.com/jakofys) Software engineer as **Owner**
