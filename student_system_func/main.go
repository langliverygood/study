package main

import "fmt"
import "os"

type student struct {
	id  int64
	name string
}

func newStudent(id int64, name string) *student {
	return &student {
		id : id,
		name : name,
	}
}

var (
	allstudent map[int64]*student
)

func showAllStudent() {
	for k, v := range allstudent {
		fmt.Printf("学号：%d 姓名：%s\n", k, v.name)
	}
}

func addStudent() {
	 var (
		 id int64
		 name string
	 )
	 fmt.Print("请输入学生学号:")
	 fmt.Scanln(&id)
	 fmt.Print("请输入学生姓名:")
	 fmt.Scanln(&name)

	 newStu := newStudent(id, name)
	 allstudent[id] = newStu
}

func deleteStudent() {
	var id int64

	fmt.Print("请输入要删除的学生学号:")
	fmt.Scanln(&id)
	delete(allstudent, id)
}

func main() {
	allstudent = make(map[int64]*student, 100)
	for {
		fmt.Println("欢迎使用学生管理系统！")
		fmt.Println(`
		1.查看所有学生
		2.新增学生
		3.删除学生
		4.退出
		`)

		fmt.Print("请输入你的选项：")
		var choice int64
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项\n", choice)

		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("滚~")
		}

	}
}