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