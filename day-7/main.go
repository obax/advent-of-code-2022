package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const DIR_TYPE string = "dir"
const FILE_TYPE string = "file"

type FS map[string][]*Node

type Node struct {
	T       string
	Name    string
	Parent  *Node
	Size    int
	Files   []*Node
	SubDirs []*Node
}

func add_node(fs FS, node *Node) (FS, error) {
	for _, f := range fs[node.T] {
		if f.Parent == node.Parent && f.Name == node.Name {
			return fs, errors.New("Already present on the list")
		}
	}
	fs[node.T] = append(fs[node.T], node)
	return fs, nil
}

func update_cursor(fs FS, cursor *Node, arg string) *Node {
	if arg == ".." {
		return cursor.Parent
	}
	for _, f := range fs[DIR_TYPE] {
		if f.Name == arg && f.Parent == cursor {
			return f
		}
	}
	return nil
}

func get_all_descendant_dirs(dirs_to_visit []*Node, dirs_visited []*Node) (tv, v []*Node) {
	var next_visit []*Node
	if len(dirs_to_visit) == 0 {
		return nil, dirs_visited
	}

	for _, tv := range dirs_to_visit {
		if tv.T == DIR_TYPE {
			dirs_visited = append(dirs_visited, tv)
			next_visit = append(next_visit, tv.SubDirs...)
		}
	}
	return get_all_descendant_dirs(next_visit, dirs_visited)
}

func solution1(dir_sizes map[*Node]Node) int {
	combined_size := 0
	for _, v := range dir_sizes {
		current := v.Size
		for _, dc := range v.SubDirs {
			current += dir_sizes[dc].Size
		}
		if 100000 >= current {
			combined_size += current
		}
	}
	return combined_size
}

func get_combined_sizes(n []*Node, info map[*Node]Node) (combined int) {
	for _, k := range n {
		combined += info[k].Size
	}
	return
}

// 1243694 too low
// 3529788 too high3268249
func solution2(dir_nodes map[*Node]Node) int {
	const DISK_SPACE int = 70000000
	const SIZE_UPDATE int = 30000000

	var dir_sizes []int

	for _, v := range dir_nodes {
		_, descendants := get_all_descendant_dirs(v.SubDirs, []*Node{})
		combined_sizes := get_combined_sizes(descendants, dir_nodes)
		dir_sizes = append(dir_sizes, combined_sizes)
	}

	sort.Sort(sort.IntSlice(dir_sizes))

	root_dir := dir_sizes[len(dir_sizes)-1]
	unused := DISK_SPACE - root_dir
	to_free_up := -(unused - SIZE_UPDATE)

	for _, i := range dir_sizes {
		if i >= to_free_up {
			return i
		}
	}
	return 0
}

func main() {
	var file_system = FS{
		DIR_TYPE:  {},
		FILE_TYPE: {},
	}
	var cursor *Node

	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(input)
	}

	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	for _, l := range lines {
		items := strings.Split(strings.ReplaceAll(l, "$ ", ""), " ")

		switch items[0] {
		case "ls":
			break
		case "dir":
			node := &Node{
				T:      DIR_TYPE,
				Parent: cursor,
				Name:   items[1],
			}
			file_system, err = add_node(file_system, node)
			if err != nil {
				break
			}
		case "cd":
			cursor = update_cursor(file_system, cursor, items[1])
		default:
			sizeStr, name := items[0], items[1]
			size, err := strconv.Atoi(sizeStr)
			if err != nil {
				panic(err)
			}
			node := &Node{
				T:      FILE_TYPE,
				Parent: cursor,
				Name:   name,
				Size:   size,
			}
			file_system, err = add_node(file_system, node)
			if err != nil {
				break
			}
		}
	}

	dir_nodes := make(map[*Node]Node, len(file_system[DIR_TYPE]))

	for _, d := range file_system[DIR_TYPE] {
		dir_nodes[d] = Node{
			Size:    0,
			Parent:  d.Parent,
			Name:    d.Name,
			Files:   []*Node{},
			SubDirs: []*Node{},
		}

		for _, f := range file_system[FILE_TYPE] {
			if f.Parent == d {
				dir_nodes[d] = Node{
					Size:    dir_nodes[d].Size + f.Size,
					Parent:  dir_nodes[d].Parent,
					Name:    dir_nodes[d].Name,
					Files:   append(dir_nodes[d].Files, f),
					SubDirs: dir_nodes[d].SubDirs,
				}
			}
		}
		for _, d1 := range file_system[DIR_TYPE] {
			if d1.Parent == d {
				dir_nodes[d] = Node{
					Size:    dir_nodes[d].Size,
					Parent:  dir_nodes[d].Parent,
					Name:    dir_nodes[d].Name,
					Files:   dir_nodes[d].Files,
					SubDirs: append(dir_nodes[d].SubDirs, d1),
				}
			}
		}
	}

	fmt.Println("size_solution1", solution1(dir_nodes))
	fmt.Println("size_solution2", solution2(dir_nodes))
}
