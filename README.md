# gqueue
Generic Queue implementation in Go using linked list

## Usage

```go
s := gqueue.New(1, 2, 3, 4)
fmt.Println(s.Len())
fmt.Println(s.Pop())
fmt.Println(s.Pop())
fmt.Println(s.Pop())
fmt.Println(s.Pop())
s.Push(5)
fmt.Println(s.Peek())
// Output: 4
// 1
// 2
// 3
// 4
// 5
```

## Changelog
### v0.1.0
- minimal supported version of Go is 1.20
- initial implementation

### v1.0.0
- minimal supported version of Go is 1.23
- added `Iter` method to iterate over the queue and consume all items
```go
s := gqueue.New(1, 2, 3, 4)
for v := range s.Iter() {
    fmt.Println(v)
}
fmt.Println("len:", s.Len())
// Output: 1
// 2
// 3
// 4
// len: 0
```

### v1.1.0
- added `IsEmpty` method to check if the queue is empty
