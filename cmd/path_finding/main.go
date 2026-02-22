package main

type point struct {
	x int
	y int
}

// TODO!!
// func walk(maze []string, wall string, curr point, end point, seen [][]bool) bool {

// 	if curr.x < 0 || curr.x >= len(maze[0]) || curr.y < 0 || curr.y >= len(maze) {
// 		return false
// 	}

// 	if maze[curr.y][curr.x] == wall {
// 		return false
// 	}

// 	if curr.x == end.x && curr.y == end.y {
// 		return true
// 	}

// 	if seen[curr.y][curr.x] {
// 		return false
// 	}
// }

// func solve(maze []string, wall string, start point, end point) []point {

// }
