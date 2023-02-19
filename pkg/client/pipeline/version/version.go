package version

import (
	"fmt"
	"github.com/aligang/Gophkeeper/pkg/client/buildinfo"
)

func Print() {
	fmt.Printf("Build Version : %s \n", buildinfo.Version)
	fmt.Printf("Build Time    : %s \n", buildinfo.BuildTime)
}
