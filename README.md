# StateRoutine

Simple in-memory cache for Go. Implemented as a loop running inside a goroutine that reads from a channel to perform operations. See `main.go` for example usage.

### Initialize Cache

```
// Replace `string` with desired value type
state := stateroutine.Go[string]()
```

### Store a new key-value pair into cache

```
stateroutine.Set(state, "key", "value")
```

### Get the value by key from cache

```
value := stateroutine.Get(state, "key")
if value != nil {
  fmt.Println(*value)
}
```

### Delete a key-value pair from cache

```
stateroutine.Delete(state, "key")
```
