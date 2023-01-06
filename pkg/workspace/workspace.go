package workspace

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

var (
	filePath, _ = exec.LookPath(os.Args[0])
	absFilePath, _ = filepath.Abs(filePath)
	baseDir = path.Dir(absFilePath)
	UnixSockSpace = InstallPathJoin()
)

func InstallPathJoin(elem ...string) string {
	joinDir := make([]string, 0)
	joinDir = append(joinDir, baseDir)
	joinDir = append(joinDir, elem...)
	return filepath.Join(joinDir...)
}