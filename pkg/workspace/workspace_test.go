package workspace

import (
	"fmt"
	"testing"
)

func TestInstallPathJoin(t *testing.T) {
	dir := InstallPathJoin(UnixSockSpace, "test.sock")
	fmt.Println(dir)
}
