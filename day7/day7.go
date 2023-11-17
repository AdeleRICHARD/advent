package day7

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"main.go/util"
)

var totalSizeDirectories int

const minimumSizeRequired = 100000

func Day7() {
	/*

		cd means change directory. This changes which directory is the current directory, but the specific result depends on the argument:
		cd x moves in one level: it looks in the current directory for the directory named x and makes it the current directory.
		cd .. moves out one level: it finds the directory that contains the current directory, then makes that directory the current directory.
		cd / switches the current directory to the outermost directory, /.
		ls means list. It prints out all of the files and directories immediately contained by the current directory:
		123 abc means that the current directory contains a file named abc with size 123.
		dir xyz means that the current directory contains a directory named xyz.
	*/
	systTxt := util.ReadFile("day7/system.txt")
	systList := strings.Split(string(systTxt), "\n")

	system := System{Dir: []*Directory{}}

	for _, line := range systList {
		command := strings.Split(line, " ")
		fmt.Printf("Command: %s\n", command)
		if strings.Contains(line, "cd") {
			system.changeDirectory(command)
			continue
		}

		if command[0] != "$" {
			system.addToCurrentDirectory(command)
		}
		if system.Dir[len(system.Dir)-1].TotalSize >= minimumSizeRequired {
			totalSizeDirectories += system.Dir[len(system.Dir)-1].TotalSize
			println(system.Dir[len(system.Dir)-1].TotalSize, " TOTAL : ", totalSizeDirectories)
		}
	}
	println(totalSizeDirectories)
}

type System struct {
	Dir []*Directory
}

type Directory struct {
	Name      string
	Parent    *Directory
	Children  []*string
	Files     []*File
	TotalSize int
}

type File struct {
	Name string
	Size int
}

func (syst *System) changeDirectory(command []string) {
	switch {
	case slices.Contains(command, "/"):
		syst.Dir = []*Directory{
			{
				Name:     "/",
				Parent:   nil,
				Children: []*string{},
				Files:    []*File{},
			},
		}
	case slices.Contains(command, ".."):
		syst.toParentDirectory()
	default:
		syst.toChildDirectory(command[len(command)-1])
	}
}

func (syst *System) toParentDirectory() {
	actualDirectory := syst.Dir[len(syst.Dir)-1]
	for i, directory := range syst.Dir {
		if actualDirectory.Name == "/" {
			syst.Dir = syst.Dir[:1]
			break
		}
		if directory.Name == actualDirectory.Parent.Name {
			// Actual is now the parent directory
			syst.Dir = syst.Dir[:i+1]
			break
		}
	}
}

func (syst *System) addToCurrentDirectory(command []string) {
	actualDirectory := syst.Dir[len(syst.Dir)-1]
	if slices.Contains(command, "dir") {
		actualDirectory.Children = append(actualDirectory.Children, &command[1])
	} else {
		size, err := strconv.Atoi(command[0])
		if err != nil {
			panic(err)
		}
		actualDirectory.TotalSize += size
		actualDirectory.Files = append(actualDirectory.Files, &File{Name: command[1], Size: size})
	}
}

func (syst *System) toChildDirectory(name string) {
	actualDirectory := syst.Dir[len(syst.Dir)-1]
	for _, directory := range actualDirectory.Children {
		if *directory == name {
			syst.Dir = append(syst.Dir, &Directory{
				Name:     name,
				Parent:   actualDirectory,
				Children: []*string{},
				Files:    []*File{},
			})
		}
	}
}
