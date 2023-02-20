package version

import (
	"fmt"
	"github.com/aligang/Gophkeeper/internal/client/buildinfo"
)

func Print() {
	fmt.Printf("Build Version : %s \n", buildinfo.Version)
	fmt.Printf("Build Time    : %s \n", buildinfo.BuildTime)
}
