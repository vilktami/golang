package todo
import(
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"todo/entities"
	// "errors"
)

var index int
var tasks map[int]*entities.Task = make(map[int]*entities.Task)



type NewTaskTodo struct {
	Task string `json:"task"`
}


type Inserter interface {
	Insert(interface{}) error
}

type Repository interface {
	NewTask(*entities.Task) error
}


type Todo struct{
	db *gorm.DB
	repo Repository
}

func NewTodo(db *gorm.DB) *Todo {
	return &Todo{db:db}
}


func (todo Todo) Add(c *gin.Context){
	var task entities.Task
	if err := c.Bind(&task); err!= nil {
		c.JSON(http.StatusBadRequest,nil)
		return 
	}

	
	if err := todo.repo.NewTask(&task).Error; err!= nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error": "err.Error()",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message":"ok",
	})
}


type GormInsert struct {
	db *gorm.DB
}

// func (insert GormInsert) Insert(v interface{}) error {
// 	return errors.WithMessage(insert.db.Create(v).Error,"gorm insert")
// }




func (todo *Todo) AddTask(c *gin.Context) {
	
	var task NewTaskTodo
	// if err := app.serialize.Decode(c,&task)
	if err := c.Bind(&task); err != nil {
		c.JSON(http.StatusBadRequest,nil)
		return
	}
	// New(task.Task)
	todo.db.Create(&entities.Task{Title:task.Task,Done:false})
}

func ChangeDoneTask(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	tasks[i].Done = true
}


func GetTask(c *gin.Context) {
	c.JSON(http.StatusOK,List())

}

func New(task string) {
	defer func() {
		index++
	}()

	tasks[index] = &entities.Task{
		Title: task,
		Done: false,
	}
}

func List() map[int]*entities.Task {
	return tasks
}