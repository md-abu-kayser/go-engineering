package main

import "fmt"

func InstallAndVerifyGo() string {
	const topic = "Install And Verify Go"
	return topic
}

func main() {
	fmt.Println(InstallAndVerifyGo())
}

package main

import (
	"fmt"
	"runtime"
)

type envInfo struct {
	GoVersion string 
	OS        string 
	Arch      string 
	NumCPU    int    
	Compiler  string 
}

func collectEnvInfo() envInfo {
	return envInfo{
		GoVersion: runtime.Version(),
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		NumCPU:    runtime.NumCPU(),
		Compiler:  runtime.Compiler,
	}
}

func main() {
	info := collectEnvInfo()

	fmt.Println("Go is installed and working!")
	fmt.Println("----------------------------------")
	fmt.Printf("Go version : %s\n", info.GoVersion)
	fmt.Printf("Operating system : %s\n", info.OS)
	fmt.Printf("Architecture : %s\n", info.Arch)
	fmt.Printf("Logical CPUs : %d\n", info.NumCPU)
	fmt.Printf("Compiler : %s\n", info.Compiler)
}
