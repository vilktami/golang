package todo
import(
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
)

var index int
var tasks map[int]*Task = make(map[int]*Task)
type Task struct {
	Title string
	Done bool
}
type NewTaskTodo struct {
	Task string `json:"task"`
   }
func AddTask(rw http.ResponseWriter,r *http.Request) {
	fmt.Println("dd")
	defer r.Body.Close()
	var task NewTaskTodo
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	New(task.Task)

}

func ChangeDoneTask(rw http.ResponseWriter,r *http.Request) {
	vars := mux.Vars(r)
	index := vars["index"]
	i,err := strconv.Atoi(index)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	tasks[i].Done = true

}


func GetTask(rw http.ResponseWriter,r *http.Request) {
	
	var task NewTaskTodo
	if err := json.NewEncoder(rw).Encode(List()); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	New(task.Task)

}

func New(task string) {
	defer func() {
		index++
	}()

	tasks[index] = &Task{
		Title: task,
		Done: false,
	}
}

func List() map[int]*Task {
	return tasks
}