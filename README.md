# Gelv

Gelv is a simple package for getting elevated privileges in golang

Install with

```shell
$ go get github.com/jacobtread/gelv
```

## Example Usage

```go
func main() {
    if !gelv.IsElevated() { // Check the app isn't already elevated
        gelv.Elevate() // Elevate the app
        return // Stop execution
    }
    // TODO: Your normal application code
}
```