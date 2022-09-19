package main

type Task struct {
	id     int
	name   string
	status bool
}

type TodoList struct {
	list [100]Task
}

func (t *Task) AddTask(TodoList *TodoList) {
	var x int = 0
	for i := 0; TodoList.list[i].name != ""; i++ {
		x++
	}
	t.status = true
	t.id = x + 1
	TodoList.list[x] = *t

}

func (t *Task) DoneTask(todoList *TodoList) {
	for i := 0; i < len(todoList.list); i++ {
		if todoList.list[i].id == t.id {
			t.status = false
			t.id = 0
			t.name = ""

			todoList.list[i] = *t
		}
	}
}

func (l *TodoList) GetTasksNotDone() TodoList {
	var list TodoList
	var k int = 0
	for i := 0; i < len(l.list); i++ {
		if l.list[i].status == true {
			list.list[k] = l.list[i]
			k++
		}
	}
	return list
}

func main() {

}
