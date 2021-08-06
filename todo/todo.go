package todo
import(
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"todo/entities"
)

var index int
var tasks map[int]*entities.Task = make(map[int]*entities.Task)



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

func (todo Todo) Add(c *gin.Context){
	var task entities.NewTaskTodo
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
// type memDB map[int]*entities.Task
// func (gdb Insert) Insert(interface{}) error {
// 		return gdb.Create(v).Error
// }

// type nodb struct{}
// func (db memDB) Insert(interface{}) error {
// 	if cache,ok := v.(*entities.Task); ok {
// 		tasks[index] = cache
// 	}
// 	return nil
// }


func NewApp(db *gorm.DB) *App {
	return &App
}

func (todo *Todo) AddTask(c *gin.Context) {
	
	var task entities.NewTaskTodo
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