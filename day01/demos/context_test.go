package demos

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func testContext(t *testing.T) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	//time.Sleep(2 * time.Second)
	dl, ok := timeoutCtx.Deadline()
	//err := timeoutCtx.Err()
	fmt.Println(dl, ok)
}

func SomeBusiness() {
	ctx := context.Background()
	Step1(ctx)
}

func Step1(ctx context.Context) {
	var db sql.DB
	db.ExecContext(ctx, "UPDATE xxxx", 1)
}
