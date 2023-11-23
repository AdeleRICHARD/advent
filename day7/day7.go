package day7

import (
	"fmt"
	"strconv"
	"strings"

	"main.go/util"
)

func Day7() {
	data := util.ReadFile("day7/system.txt")
	lines := strings.Split(string(data), "\n")

	root := &Directory{Path: "/", Size: 0, Parent: nil}
	current := root
	var directories []*Directory
	directories = append(directories, root)

	for _, line := range lines {
		command := strings.Fields(line)
		// Print actualDirectories

		fmt.Printf("Command: %s\n", command)

		if command[0] == "$" {
			if command[1] != "ls" {
				switch command[2] {
				case "/":
					continue
				case "..":
					if current.Parent != nil {
						current = current.Parent
					}
				default:
					newDir := &Directory{Path: current.Path + "/" + command[2], Size: 0, Parent: current}
					current = newDir
					directories = append(directories, newDir)
				}
			}
		} else {
			size, err := strconv.Atoi(command[0])
			if err != nil {
				fmt.Println("Error parsing file size:", err)
				continue
			}
			addSize(current, uint64(size))
		}
	}

	totalSize := uint64(0)
	// Taille demandée
	// Taille minimum requise - (taille totale du disque - taille totale de chaque dossier)
	minRequired := uint64(30000000)
	spaceMinimumNeeded := minRequired - (70000000 - directories[0].Size)
	for _, dir := range directories {
		if dir.Size <= 100000 {
			totalSize += dir.Size
		}
		// Cherche l'espace suplémentaire nécessaire à supprimer
		// Doit être plus grand ou égal à l'espace nécessaire
		// Mais on veut le plus petit des fichiers à supprimer
		if dir.Size >= spaceMinimumNeeded && dir.Size < minRequired {
			minRequired = dir.Size
		}

	}

	fmt.Println("Total size of directories <= 100000: ", totalSize)
	fmt.Println("Minimum size required: ", minRequired)
}

// uint64 meilleur pour calculer des gros nombres
type Directory struct {
	Path   string
	Size   uint64
	Parent *Directory
}

func addSize(current *Directory, size uint64) {
	// Will go back to the parent until it reaches parent == nil
	for c := current; c != nil; c = c.Parent {
		c.Size += size
	}
}
