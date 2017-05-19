# riddles
Riddles via the command line.

## Installation

### Via Go

```sh
go get github.com/tjkemper/riddle
```

## Usage

* **`--id, -i`:** print riddle with id
* **`--random, -r`:** print random riddle *(default true)*
* **`--answer, -a`:** show answer

### Examples

```sh
# print random riddle
$ riddle

# print riddle with id 0
$ riddle --id 0

# show answer
$ riddle --id 0 --answer
```
