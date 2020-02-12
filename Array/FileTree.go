package main

import (
	"AlgorithmGolang/Array/queue"
	"AlgorithmGolang/Array/stack"
	"fmt"
	"io/ioutil"
	"os"
)

func FileTreeRecursion(path string, files []string) ([]string, error) {
	var (
		r   []os.FileInfo
		err error
	)
	if r, err = ioutil.ReadDir(path); err != nil {
		return []string{}, err
	}
	for _, fi := range r {
		fulldir := path + "/" + fi.Name()
		files = append(files, fulldir)
		if fi.IsDir() {
			files, _ = FileTreeRecursion(fulldir, files)
		}
	}
	return files, nil
}

func FileTreeStack(path string) (files []string, err error) {
	var (
		mystack = stack.NewArrayListStack()
	)
	mystack.Push(path)
	for !mystack.IsEmpty() {
		var (
			curPath string
			r       []os.FileInfo
		)
		curPath = mystack.Pop().(string)

		if r, err = ioutil.ReadDir(curPath); err != nil {
			return []string{}, err
		}
		for _, fi := range r {
			fulldir := curPath + "/" + fi.Name()
			files = append(files, fulldir)
			if fi.IsDir() {
				mystack.Push(fulldir)
			}
		}
	}
	return
}

func FileTreeQueue(path string) (files []string, err error) {
	var (
		q = queue.NewQueueArray()
	)
	q.EnQueue(path)
	for !q.IsEmpty() {
		var (
			r       []os.FileInfo
			info    os.FileInfo
			curPath string
		)
		curPath = q.DeQueue().(string)

		if info, err = os.Stat(curPath); err != nil {
			return
		}
		files = append(files, curPath)
		if info.IsDir() {
			if r, err = ioutil.ReadDir(curPath); err != nil {
				return
			}
			for _, fi := range r {
				fulldir := curPath + "/" + fi.Name()
				q.EnQueue(fulldir)
			}
		} else {

		}

	}
	return
}

func main() {
	//getFileTreeStack()
	getFileTreeQueue()
}

func getFileTreeQueue() {
	var (
		path  = "."
		files = []string{}
		err   error
	)
	if files, err = FileTreeQueue(path); err != nil {
		panic(fmt.Sprintf("display file tree err %v", err))
	}
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}

}

func getFileTreeStack() {
	var (
		path  = "."
		files = []string{}
		err   error
	)
	if files, err = FileTreeStack(path); err != nil {
		panic(fmt.Sprintf("display file tree err %v", err))
	}
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}

func getFileTree() {
	var (
		path  = "."
		files = []string{}
		err   error
	)
	if files, err = FileTreeRecursion(path, files); err != nil {
		panic(fmt.Sprintf("display file tree err %v", err))
	}
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}
