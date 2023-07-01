## Couple notes while learning Go 

- Functions cannot have parameters with default values
- Arrays have always size encoded in its type
- Use Slices for dynamically-sized array
- equality operators do not work with slices. Use reflect.DeepEqual
- variables can be assigned to functions inside another function "scope"
    - the function/variable is hidden from other functions in the package
    - e.g privateFunc := func(){}


- `go test -cover` checks the test coverage
- Test files can also include Examples