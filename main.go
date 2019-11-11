package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jszwec/csvutil"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Email string `csv:"email" validate:"email"`
}

var (
	out        = os.Stdout
	targetFile = flag.String("t", "", "target csv file")
)

func main() {
	flag.Parse()

	if len(strings.TrimSpace(*targetFile)) == 0 {
		fmt.Fprintf(out, "empty target file\n")
		os.Exit(1)
	}

	b, err := ioutil.ReadFile(*targetFile)
	if err != nil {
		fmt.Fprintln(out, err)
		os.Exit(1)
	}

	var users []*User
	if err := csvutil.Unmarshal(b, &users); err != nil {
		fmt.Fprintln(out, err)
		os.Exit(1)
	}

	v := validator.New()
	for i, u := range users {
		if err := v.Struct(u); err != nil {
			fmt.Fprintf(out, "%d行目: %s\n", i+2, u.Email)
		}
	}
}
