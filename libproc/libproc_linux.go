package libproc

import (
	"fmt"
	"os"
)

func ProcPath(pid int) (string, error) {
	return os.Readlink(fmt.Sprintf("/proc/%d/exe", pid))
}
