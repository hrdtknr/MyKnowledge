package infra

import(
	"fmt"
	"time"
	"github.com/test/layered/repo"
	"github.com/go-xorm/xorm"
)


type sampleInfra struct{
	db *xorm.Engine
}

// あとで移動する
type (
	Todo struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Todo string `json:"todo"`
	}
)


func NewSampleInfra(db *xorm.Engine) repo.SampleInterface{ // 戻り値のSampleInterfaceはrepoにあるので、repo.interfaceを返す
	return &sampleInfra{db: db}
}

func (s *sampleInfra) Method() time.Time{
	fmt.Println("infra")
	var todoList []Todo
	table := s.db.Table("todo")
	if err := table.Find(&todoList); err != nil {
		fmt.Println("infra err:", err)
	}

	fmt.Println(todoList)



	return time.Now()
}