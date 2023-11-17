package day7

import (
	"slices"
	"strings"

	"main.go/util"
)

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

	system := System{Dir: []Directory{}}

	for _, line := range systList {
		command := strings.Split(line, " ")

		if strings.Contains(line, "cd") {
			system.changeDirectory(command)
		} else if strings.Contains(line, "ls") {
			// TODO
		} else if strings.Contains(line, "dir") {
			// TODO
		} else {
			// TODO
		}
	}
}

type System struct {
	Dir []Directory
}

type Directory struct {
	Name     string
	Parent   *Directory
	Children []*string
	Files    []*File
}

type File struct {
	Name string
	Size int
}

func (syst *System) changeDirectory(command []string) {
	switch {
	case slices.Contains(command, "/"):
		// TODO
	case slices.Contains(command, ".."):
		syst.toParentDirectory()

	default:
		// TODO
	}
}

func (syst *System) toParentDirectory() {
	actualDirectory := syst.Dir[len(syst.Dir)-1]
	for i, directory := range syst.Dir {
		if directory.Name == actualDirectory.Parent.Name {
			// Actual is now the parent directory
			syst.Dir[len(syst.Dir)-1] = syst.Dir[i]
		}
	}
}
