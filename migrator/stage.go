package migrator

import "github.com/go-xorm/builder"

type BuildSQLFunc func(b *builder.Builder) (string, error)

type Stage struct {
	ID       string
	Migrate  BuildSQLFunc
	Rollback BuildSQLFunc
}

type ByID []Stage

func (a ByID) Len() int           { return len(a) }
func (a ByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByID) Less(i, j int) bool { return a[i].ID < a[j].ID }
