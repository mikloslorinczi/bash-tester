package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	if err := ioutil.WriteFile("/tmp/sajt.sh", []byte(`
	echo "asdf" >> /tmp/test.txt
		`), 0777); err != nil {
		panic(err)
	}
	cmd := exec.Command("/bin/sh", "-c", "/tmp/sajt.sh")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	b, err := ioutil.ReadFile("/tmp/test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("/tmp/test.txt:\n%s", b)
}
