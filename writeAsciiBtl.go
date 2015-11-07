// writeAsciiBtl
package main

import (
	"fmt"
)

func (nc *Nc) WriteAsciiBTL(map_format map[string]string, hdr []string,cfg ConfigBTL) {
	fmt.Println("Hello from WriteAscii for bottle !")
}

func (nc *Nc) WriteHeader(map_format map[string]string, hdr []string) {
	fmt.Println("Hello from WriteHeader for bootle !")
}
