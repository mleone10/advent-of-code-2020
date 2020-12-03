# Advent of Code 2020
> Inputs and solutions to 2020's daily challenges

## Setup
All executables contained within this repo can be built with a single `make` invocation:
```bash
$ make build
```
Resultant executables will be located in `./bin`.

### Day Initialization
A helper script is included to easily create directories for new problems as they are released.  Running it will create a directory with the given suffix in the `cmd` directory and create an initial `main.go` file:
```bash
$ ./newday <day number, e.g. 03>
```

## Execution
All solutions are executed directly from the command line.  If a problem requires an input, it is accepted from STDIN:
```bash
$ ./cmd/day11/input.txt | ./bin/day11
```