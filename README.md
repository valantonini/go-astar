# go-astar

[![Build and test](https://github.com/valantonini/go-astar/actions/workflows/go.yml/badge.svg)](https://github.com/valantonini/go-astar/actions/workflows/go.yml)

A 2d AStar implementation in Go.

```go
// will traverse the 1
w := []int{
    0, 0, 0, 0, 0, 0, 0, 0,
    0, 1, 0, 1, 1, 1, 0, 0,
    0, 1, 1, 1, 0, 1, 1, 0,
    0, 0, 0, 0, 0, 0, 0, 0,
}
grid := NewGridFromSlice(8, 4, w)

pathfinder := NewPathfinder(grid)
got := pathfinder.Find(Vec2{1, 1}, Vec2{6, 2})

want := []Vec2{
    {1, 1},
    {1, 2},
    {2, 2},
    {3, 2},
    {3, 1},
    {4, 1},
    {5, 1},
    {5, 2},
    {6, 2},
}
```

To Do:
- diagonals
- different heuristic calcs
- punish change direction

see [my c# implementation](https://github.com/valantonini/AStar)
