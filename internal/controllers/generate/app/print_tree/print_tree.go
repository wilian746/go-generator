package printtree

import (
	"fmt"
	"github.com/google/uuid"
	prettyList "github.com/jedib0t/go-pretty/v6/list"
	"path"
	"sort"
	"strings"
)

type File struct {
	ID   string
	Name string
}

type Folder struct {
	Name    string
	Files   []File
	Folders map[string]*Folder
}

func newFolder(name string) *Folder {
	return &Folder{name, []File{}, make(map[string]*Folder)}
}

func (f *Folder) getFolder(name string) *Folder {
	if nextF, ok := f.Folders[name]; ok {
		return nextF
	}
	if f.Name == name {
		return f
	}
	return &Folder{}
}

func (f *Folder) existFolder(name string) bool {
	for _, v := range f.Folders {
		if v.Name == name {
			return true
		}
	}
	return false
}

func (f *Folder) addFolder(folderName string) {
	if !f.existFolder(folderName) {
		f.Folders[folderName] = newFolder(folderName)
	}
}

func (f *Folder) addFile(fileName, fileID string) {
	f.Files = append(f.Files, File{fileID, fileName})
}

type Interface interface {
	Print()
}

type Tree struct {
	ListPaths []string
	Destiny   string
}

func NewTree(list []string, destiny string) Interface {
	return &Tree{
		ListPaths: list,
		Destiny:   destiny,
	}
}

func (t *Tree) Print() {
	currentFolder := t.getCurrentFolderByListPaths()

	l := prettyList.NewWriter()
	lTemp := prettyList.List{}
	lTemp.Render()
	l.Reset()
	l.SetStyle(prettyList.StyleConnectedRounded)

	t.addToTreeFilesFromFolder(currentFolder, l)

	fmt.Println(l.Render())
}

func (t *Tree) getCurrentFolderByListPaths() *Folder {
	rootFolder := newFolder(t.Destiny)

	t.loopToAddFolderAndFilesInRootFolder(rootFolder)

	return rootFolder.getFolder(t.Destiny)
}

func (t *Tree) loopToAddFolderAndFilesInRootFolder(rootFolder *Folder) {
	for _, filePath := range t.ListPaths {
		splitPath := t.deleteEmptyElements(strings.Split(filePath, "/"))
		tmpFolder := rootFolder
		for _, item := range splitPath {
			if t.isFile(item) {
				tmpFolder.addFile(item, uuid.New().String())
			} else {
				if item != t.Destiny {
					tmpFolder.addFolder(item)
				}
				tmpFolder = tmpFolder.getFolder(item)
			}
		}
	}
}

func (t *Tree) isFile(str string) bool {
	if path.Ext(str) != "" && str != "Makefile" {
		return true
	}
	return false
}

func (t *Tree) deleteEmptyElements(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func (t *Tree) addToTreeFilesFromFolder(currentFolder *Folder, l prettyList.Writer) {
	l.AppendItem(currentFolder.Name)
	l.Indent()
	for _, file := range t.getAlphabeticFiles(currentFolder.Files) {
		l.AppendItem(file.Name)
	}
	if len(currentFolder.Folders) > 0 {
		for _, internalFolder := range t.getAlphabeticFolders(currentFolder.Folders) {
			t.addToTreeFilesFromFolder(internalFolder, l)
			l.UnIndent()
		}
	}
}

func (t *Tree) getAlphabeticFiles(files []File) (newFiles []File) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name < files[j].Name
	})
	return files
}

func (t *Tree) getAlphabeticFolders(folders map[string]*Folder) (newFolders []*Folder) {
	for _, folder := range folders {
		newFolders = append(newFolders, folder)
	}
	sort.Slice(newFolders, func(i, j int) bool {
		return newFolders[i].Name < newFolders[j].Name
	})
	return newFolders
}
