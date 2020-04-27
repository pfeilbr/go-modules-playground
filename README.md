# go-modules-playground

learn using [golang modules](https://blog.golang.org/using-go-modules)

## Resources

* <https://blog.golang.org/using-go-modules>
* [golang/go/wiki/Modules](https://github.com/golang/go/wiki/Modules)

## Session

```sh
# run tests recursively on change
fswatch -o . | xargs -n1 -I{} go test -v ./...

# run main on change
fswatch -o . | xargs -n1 -I{} go run main.go
```