package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

type Node struct {
	Name     string
	IsDir    bool
	Children []*Node
}

var root = &Node{
	Name:  ".",
	IsDir: true,
}

func myWalkDirFunc(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if path == "." {
		return nil
	}

	parts := strings.Split(path, "/")
	current := root

	for i, part := range parts {
		isLast := i == len(parts)-1

		var child *Node
		for _, c := range current.Children {
			if part == c.Name {
				child = c
				break
			}
		}

		if child == nil {
			child = &Node{
				Name:  part,
				IsDir: d.IsDir() && isLast,
			}

			current.Children = append(current.Children, child)
		}
		current = child
	}

	return nil
}

func printTree(n *Node, prefix string, isLast bool) {
	branch := "├── "
	if isLast {
		branch = "└── "
	}

	var name string
	if n.IsDir {
		name = n.Name + "/"
	} else {
		name = n.Name
	}

	fmt.Println(prefix + branch + name)

	childPrefix := prefix
	if isLast {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}

	for i, child := range n.Children {
		printTree(child, childPrefix, i == len(n.Children)-1)
	}
}

func main() {
	fileSystem := os.DirFS(".")

	err := fs.WalkDir(fileSystem, ".", myWalkDirFunc)
	if err != nil {
		log.Fatal(err)
	}

	printTree(root, "", true)
}
