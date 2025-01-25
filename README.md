# Monkey Interpreter

An implementation of the Monkey Interpreter from Thorston Ball's Writing an Interpreter in Go and Writing a Compiler in Go.

Extended to support running a script from a file.

## Building

```bsah
go build
```

## Testing

```bash
go test ./...
```

# Running

### Run The REPL

```bash
go run .
```

### Run A Source File

```bash
go run . monkey-scripts/fact.monkey
```