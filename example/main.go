package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ehrktia/sensitive"
)

func main() {
	d := struct {
		Name, Sensitive string
		Age             int
	}{
		Name:      "Testing",
		Sensitive: "my-secret-value",
		Age:       21,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	out, err := sensitive.CreateStruct(d, "Sensitive")
	if err != nil {
		panic(fmt.Sprintf("%q", err.Error()))
	}

	logger.Info("message", "data", out)

}
