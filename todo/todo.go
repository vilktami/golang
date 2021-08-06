package todo
import(
	"io"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	
)

var index int
var tasks map[int]*Task = make(map[int]*Task)

type NewTaskTodo struct {
	Task string `json:"task"`
}

type Inserter interface {
	Insert(interface{}) error
}

type Repository interface {
	NewTask(*Task) error
}


type Todo struct{
	db *gorm.DB
	repo Repository
}

func (todo Todo) Add(c *gin.Context){
	var task NewTaskTodo
	if err := c.Bind(&task); err!= nil {
		c.JSON(http.StatusBadRequest,nil)
		return 
	}

	
	if err := todo.repo.NewTask(&task).Error; err!= nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message":"ok",
	})
}


// func (j JSONSerializer) Decode(r io.Reader, v interface{}) error {
// 	return json.NewDecoder(r).Decode(v)
// }
// func (j JSONSerializer) Encode(w io.Writer, v interface{}) error {
// 	return json.NewEncoder(w).Encode(v)
// }


// type App struct {
// 	serialize Serializer
// }
   
// type Serializer interface {
// 	Decode(io.Reader, interface{}) error
// 	Encode(io.Writer, interface{}) error
// }


// func NewApp(serialize Serializer) *App {
// 	return &App{
// 		serialize: serialize,
// 	}
// }

type App struct {
	// db *gorm.DB
	db Inserter
}

type Insert *gorm.DB
type memDB map[int]*Task
func (gdb Insert) Insert(interface{}) error {
		return gdb.Create(v).Error
}

type nodb struct{}
func (db memDB) Insert(interface{}) error {
	if cache,ok := v.(*Task); ok {
		tasks[index] = cache
	}
	return nil
}

type Inserter interface {
	Insert(interface{}) error
}

func NewApp(db *gorm.DB) *App {
	return &app
}

func (app *App) AddTask(c *gin.Context) {
	
	var task NewTaskTodo
	// if err := app.serialize.Decode(c,&task)
	if err := c.Bind(&task); err != nil {
		c.JSON(http.StatusBadRequest,nil)
		return
	}
	// New(task.Task)
	app.db.Create(&Task{Title:task.Task,Doe:false})
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

	tasks[index] = &Task{
		Title: task,
		Done: false,
	}
}

func List() map[int]*Task {
	return tasks
}