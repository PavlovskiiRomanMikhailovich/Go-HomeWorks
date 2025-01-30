package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	out := &bytes.Buffer{}
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(for_out *bytes.Buffer, filename string, need_print bool) error {
	if need_print {
		print_info(for_out, filename, "")
	} else {
		print_only_folder(for_out, filename, "")
	}
	return nil
}

func print_only_folder(for_write *bytes.Buffer, filename string, head string) {
	files, err := os.ReadDir(filename)
	os.Chdir(filename)
	if err != nil {
		fmt.Println("error")
	}
	for _, file := range files {
		if file.IsDir() {
			io.WriteString(for_write, head)
			flag := 0
			flag_2 := 0
			for _, file_2 := range files {
				if file_2 == file {
					flag_2 = 1
					continue
				}
				if file_2.IsDir() && flag_2 == 1 {
					flag = 1
					break
				}
			}
			if files[len(files)-1] == file || flag != 1 {
				io.WriteString(for_write, "└───"+file.Name()+"\n")
			} else {
				io.WriteString(for_write, "├───"+file.Name()+"\n")
			}
			if files[len(files)-1] == file || flag == 0 {
				current := head + "    "
				print_only_folder(for_write, file.Name(), current)
			} else {
				current := head + "│   "
				print_only_folder(for_write, file.Name(), current)
			}
			os.Chdir("../")
		}
	}
}

func print_info(for_write *bytes.Buffer, filename string, head string) {
	files, err := os.ReadDir(filename)
	os.Chdir(filename)
	if err != nil {
		fmt.Println("error")
	}
	for _, file := range files {
		io.WriteString(for_write, head)
		if files[len(files)-1] == file {
			io.WriteString(for_write, "└───"+file.Name()+get_size(file.Name())+"\n")
		} else {
			io.WriteString(for_write, "├───"+file.Name()+get_size(file.Name())+"\n")
		}
		if file.IsDir() {
			if files[len(files)-1] == file {
				current := head + "        "
				print_info(for_write, file.Name(), current)
			} else {
				current := head + "│       "
				print_info(for_write, file.Name(), current)
			}
			os.Chdir("../")
		}
	}
}

func get_size(name string) string {
	fi, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	f, err := fi.Stat()
	if f.IsDir() {
		return ""
	} else {
		if f.Size() == 0 {
			return " (empty)"
		} else {
			return " (" + fmt.Sprint(f.Size()) + "b) "
		}
	}
}
