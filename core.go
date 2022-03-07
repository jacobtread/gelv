package gelv

import (
	"golang.org/x/sys/windows"
	"log"
	"os"
	"strings"
	"syscall"
)

// Elevate tells windows to elevate the current application. Will relaunch this app with the required perms
// make sure to check IsAdmin before running any app logic to be sure you don't continue executing
func Elevate() {
	exe, _ := os.Executable()              // Get the current executable name
	cwd, _ := os.Getwd()                   // Get the current working directory
	args := strings.Join(os.Args[1:], " ") // Take all the user arguments as a string
	// Converts the values to utf16 pointers for windows
	verbPointer, _ := syscall.UTF16PtrFromString("runas")
	exePointer, _ := syscall.UTF16PtrFromString(exe)
	cwdPointer, _ := syscall.UTF16PtrFromString(cwd)
	argsPointer, _ := syscall.UTF16PtrFromString(args)
	// Execute a new shell window with admin perms
	err := windows.ShellExecute(0, verbPointer, exePointer, argsPointer, cwdPointer, syscall.SW_NORMAL)
	if err != nil {
		log.Fatalln("Failed to elevate application", err)
	}
}

// IsElevated checks whether the application has elevated permissions or not. The user shouldn't be able
// to access the path if they are not elevated
func IsElevated() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	return err == nil
}
