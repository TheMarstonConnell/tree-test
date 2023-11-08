# Tree Test

This repo is a cli tool for testing the creation of merkle trees on files.

## Install
```shell
go install github.com/TheMarstonConnell/tree-test@v0.0.0
```

## Usage

### Basic Usage
The default chunk size is 10240
```shell
tree-test [file-name]
```

### Changing the chunk size
```shell
tree-test [file-name] --size=[new-size]
```