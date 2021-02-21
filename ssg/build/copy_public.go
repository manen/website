package build

import (
	"fmt"
	"os"

	"github.com/otiai10/copy"
)

func copyPublic() {
	copy.Copy("./public", "./build")
	copy.Copy("./public/secret", "./build")
	os.RemoveAll("./build/secret")

	fmt.Println("Public folder(s) copied")
}
