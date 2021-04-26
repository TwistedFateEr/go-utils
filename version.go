package go_utils

import "os"

var(
	BUILD_VERSION = os.Getenv("BUILD_VERSION")
	BUILD_NAME=os.Getenv("BUILD_NAME")
)