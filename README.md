# go-astar

[![Build and test](https://github.com/valantonini/go-astar/actions/workflows/go.yml/badge.svg)](https://github.com/valantonini/go-astar/actions/workflows/go.yml)

A 2d AStar implementation in Go.

```go
// will traverse the 1
weights := []int{
    0, 0, 0, 0, 0, 0, 0, 0,
    0, 1, 0, 1, 1, 1, 0, 0,
    0, 1, 1, 1, 0, 1, 1, 0,
    0, 0, 0, 0, 0, 0, 0, 0,
}
grid := astar.NewGridFromSlice(8, 4, weights)

pathfinder := astar.NewPathfinder(grid)
got := pathfinder.Find(astar.Vec2{1, 1}, astar.Vec2{6, 2})

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

fmt.Println(slices.Equal(got, want)) // true
```

To Do:
- punish change direction (in progress)

see [my c# implementation](https://github.com/valantonini/AStar)
