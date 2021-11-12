# Mural

Mural is a CLI library write in go allow you to implement CLI logic on your go application.
It use multiple pattern to describe how a CLI can be work.
Must usefull feature of Mural is **Description as Code** pattern that allow you to just describe your command and tag as a simple documentation (as help common command can do) and Mural, from the description, can construct you command logic.
Any code need, just description!!!

## Installation

Make only a go get command:

```bash
go get -u github.com/jakofys/mural
```

## Features

- [ ] Description as Code
    - [ ] Command description
    - [ ] Flag description
    - [ ] Variable
    - [ ] Environment variable
- [ ] Code own CLI
- [ ] Extensible flag logic
- [ ] Subcommand supports
- [ ] Autocompletion command line

## Use

### Create command

Initialize basic information about CLI application as:

```go
version := ParseVersion("1.12.3-alpha")
cli := NewCLI("myapp", version, "A simple application")
```

> CLI supports semantic version pattern

And just add a command like:

```go
cli.AddCommand(
    &Command{
        Name: "hello",
        Description: "This command get a hello world",
        Handler: HelloCommandHandler,
    }
)
```

Then implement `CommandHandler` to create a handler function:

```go
func HelloCommandHandler (cmd Command, args map[string]Argument, flags map[string]Flag){
    fmt.Print("Hellow world")
}
```

Don't forget to specify to CLI to build is intern logic using:

```go
_, err := cli.Build()
```

> This function allow cli to build logic with given ressources, and (par conséquent) locked ressources access. You cannot neither adding or removing, neither changing ressources behaviour. <br>
It returns error if exist duplication, no sense logic or missing informations. 

Finally for start running CLI, just make:

```go
cli.Run(os.Args)
```

We can run this application to making:

```bash
myapp hello # Output 'Hello world!'
```

## Description as Code

You have three way to declare command, the first, we saw it in previous chapter as basic code, the second is to describe in function as a command line already tap.

Instead of coding, you have the choice to make it shorter as:

```go
cli.AddPatternCommand("hello", "This command get a hello world", HelloCommandHandler)
```

And then last and better way is to describe your commands and flags to  import as command logic tool as:

```txt
@commands:
    hello
    > This command get a simple hello world!
```

And make a bind to refer description code to an existing handler making:

```go
cli.BindCommand("command", HelloCommandHandler)
```

And we have the same result as previous use cases.

### Introduction

A **description** instance is a set of `@annotating` information.
A block start with type of information, and then, list description's content respecting data format.

### Command

A command has it own pattern, and it composed of:

- Command name
- Argument name (only string parameter)
- Flag name or tag
- Description (for `help` command)

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

**Optional argument** allow you to not declare argument value on command line, without return error.

`[?arg]`

**Default argument value** can set default value to argument if it null.

`[arg='default value']`

**Argument as list** create a pattern that list string argument.

`[arg...]`

Or define minimum and/or maximum string using: `[arg min-max]`

> Be carefull using **Argument as list**,  it has the same behaviour than `params...` golang syntax., and then, can be null by default. So make this type at last of a command because you can have **only one argument list by command**.

> You can have multiple argument list seperate by command as `command [arg1...] subcommand [arg2...]`


### Flag

A flag is a way to add behaviour metadata to create context to the command.
You can add flag definition as:

```go
cli.AddPatternFlag("flag | f1: string", "Flag one description as string")
```

You can create you own flag logic registering an extractor as (who implements `FlagExtractor` interface):

```go
cli.RegisterFlagExtractor(&StringFlag{})
```

### File description

You can also create a documentation file use specific language to descibe your entirely application with:

```go
cli := mural.From(io.Reader) // Output a cli importing file description
```

Then in this content, we have:

```
@version: 1.12.3-pre
@name: myapp

@commands:

    command subcommand [arg1] [?arg2] --flag1 --flag2
    > Command description has coming soon

    command command2 [arg1] nextcommant [arg2] -f1 -f2
    > Command for deleguate to next command value

@flags:

    flag1 | f1: string
    > Flag one description as string

    flag2 | f2: bool=true
    > Flag two description as boolean with true as default value
```

Operator `>` determine a description line.

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

## Author and contributor

- [Jacques COFIS](github.com/jaokfys) Software engineer as **Owner**
