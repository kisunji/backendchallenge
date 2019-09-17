# League Backend Challenge

## Comments
The instructions specified a square matrix input (NxN) all operations can technically handle MxN inputs.

`matrix.IsSquareMatrix` is unused but can be used to validate if required.

Download dependencies
```
go get
```

Run web server
```
go run .
```

Send request
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
```

## What we're looking for

- The solution runs
- The solution performs all cases correctly
- The code is easy to read
- The code is reasonably documented
- The code is tested
- The code is robust and handles invalid input and provides helpful error messages
