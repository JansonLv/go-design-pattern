package main

func main() {
	f1 := &file{name: "file1"}
	f2 := &file{name: "file2"}
	f3 := &file{name: "file3"}

	folder2 := &folder{name: "folder2"}

	folder1 := &folder{name: "folder1", components: []component{f1, f2, f3, folder2}}

	folder1.search("1")

}
