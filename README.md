# go-modules-playground

learn using [golang modules](https://blog.golang.org/using-go-modules)

## Resources

* <https://blog.golang.org/using-go-modules>

## Session

```sh
# run tests recursively on change
fswatch -o . | xargs -n1 -I{} go test -v ./...

# run main on change
fswatch -o . | xargs -n1 -I{} go run main.go
```