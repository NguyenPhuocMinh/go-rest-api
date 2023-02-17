package utils

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	constants "fast-food-api-client/constants"

	"github.com/gosimple/slug"
)

// Make slug
func SlugifyString(s string) string {
	return slug.Make(s)
}

// Get File By Error
func GetFileByError(err error) (string, string) {
	pc, filename, line, _ := runtime.Caller(2)

	compName := fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
	fileName := fmt.Sprintf("%s:%d", filename, line)

	fmt.Printf("Error File = [%s:%d]", filename, line)
	fmt.Println()
	fmt.Printf("Error Comp = [%s]", compName)
	fmt.Println()

	return compName, fileName
}

// Audit attribute
func AuditAttribute(mode string) (createdAt time.Time, updatedAt time.Time) {
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	now := time.Now().In(loc)

	if mode == constants.CreateMode {
		createdAt = now
		updatedAt = now
	}

	updatedAt = now

	return createdAt, updatedAt
}

func JoinString(s []string) string {
	return strings.Join(s, "-")
}
