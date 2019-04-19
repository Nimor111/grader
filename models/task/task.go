package task

import (
	"fmt"
	"github.com/Nimor111/grader/database"
	log "github.com/sirupsen/logrus"
)

type Task struct {
	Id    int64
	Name  string
	Code  string
	Tests string
}

func (t *Task) String() string {
	return fmt.Sprintf("Task<%d %s>", t.Id, t.Name)
}

func (t *Task) Insert() error {
	err := t.DB.GetDB().Insert(t)
	if err != nil {
		log.Errorf("Failed to insert task %v into database.", t.String())
	}

	log.Infof("Inserted task %v into database.", t.String())

	return nil
}
