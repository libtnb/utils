package env

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// IsDirectlyRun checks if the application is running using go run.
// IsDirectlyRun 检查应用程序是否使用 go run 运行。
func IsDirectlyRun() bool {
	executable, err := os.Executable()
	if err != nil {
		return false
	}

	return strings.Contains(filepath.Base(executable), os.TempDir()) ||
		(strings.Contains(filepath.ToSlash(executable), "/var/folders") && strings.Contains(filepath.ToSlash(executable), "/T/go-build")) // macOS
}

// IsGithub returns whether the current environment is GitHub Action.
// IsGithub 返回当前系统环境是否为 GitHub Action。
func IsGithub() bool {
	_, exists := os.LookupEnv("GITHUB_ACTION")

	return exists
}

// IsWindows returns whether the current operating system is Windows.
// IsWindows 返回当前操作系统是否为 Windows。
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsLinux returns whether the current operating system is Linux.
// IsLinux 返回当前操作系统是否为 Linux。
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsDarwin returns whether the current operating system is Darwin.
// IsDarwin 返回当前操作系统是否为 Darwin。
func IsDarwin() bool {
	return runtime.GOOS == "darwin"
}

// IsArm returns whether the current CPU architecture is ARM.
// IsArm 返回当前 CPU 架构是否为 ARM。
func IsArm() bool {
	return runtime.GOARCH == "arm" || runtime.GOARCH == "arm64"
}

// IsX86 returns whether the current CPU architecture is X86.
// IsX86 返回当前 CPU 架构是否为 X86。
func IsX86() bool {
	return runtime.GOARCH == "386" || runtime.GOARCH == "amd64"
}

// Is64Bit returns whether the current CPU architecture is 64-bit.
// Is64Bit 返回当前 CPU 架构是否为 64 位。
func Is64Bit() bool {
	return runtime.GOARCH == "amd64" || runtime.GOARCH == "arm64"
}
