package jsii

import (
	_      "embed"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

//go:embed mexc-sdk-1.0.0.tgz
var tarball []byte

// Initialize loads the necessary packages in the @jsii/kernel to support the enclosing module.
func Initialize() {
	// Explicitly set Node.js path (fixes "node not found in PATH")
	// Use the correct path for your system (e.g., /usr/bin/node or NVM path)
	_jsii_.SetNodePath("/usr/bin/node")  // or /root/.nvm/versions/node/v20.16.0/bin/node

	// Load this library into the kernel
	_jsii_.Load("mexc-sdk", "1.0.0", tarball)
}