# Testing in go
The standard toolchain of go contains the go test tool. This tool provides functionality to implement unit tests. It scans a package or a set of packages for go functions in go files that have been marked to be tests and executes them. 

To implement a test the file which holds the tests needs to have the suffix `_test`. This package implements a simple terminal programm that can add two numbers. The programm is tested with the standard go test toolchain. Just execute `test ../...` in the directory of this README.