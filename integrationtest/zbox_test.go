package integrationtest

import (
	"fmt"
	"log"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegister(t *testing.T) {
	cmd := exec.Command("./zbox", "register")         // or whatever the program is
	cmd.Dir = "/Users/dev/Desktop/0chain/zboxcli" // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s", out)
	}
	require.Equal(t, "Wallet registered\n", string(out))
}


func TestRegister(t *testing.T) {
	cmd := exec.Command("./zbox", "register")         // or whatever the program is
	cmd.Dir = "/Users/dev/Desktop/0chain/zboxcli" // or whatever directory it's in
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s", out)
	}
	require.Equal(t, "Wallet registered\n", string(out))
}

