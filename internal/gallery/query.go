package gallery

import (
	"context"
	"database/sql"
	"fmt"
	comsmaDB "github.com/comsma/comsma/database"
	"os"
)

var ctx context.Context
var queries *comsmaDB.Queries

func Init() {
	ctx = context.Background()
	db, err := sql.Open("mysql", os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}
	queries = comsmaDB.New(db)
}
func GetGalleries(galleries []comsmaDB.Gallery) {
	var err error
	galleries, err = queries.GetGalleries(ctx)
	fmt.Println(err)
}
