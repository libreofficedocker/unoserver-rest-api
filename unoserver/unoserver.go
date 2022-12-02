package unoserver

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os/exec"
)

var DefaultConnection = "socket,host=%s,port=%s,tcpNoDelay=1;urp;StarOffice.ComponentContext"

var DefaultLibreofficeOptions = []string{
	"--headless",
	"--invisible",
	"--nocrashreport",
	"--nodefault",
	"--nologo",
	"--nofirststartwizard",
	"--norestore",
}

var unoserver = &Unoserver{
	Interface:  "127.0.0.1",
	Port:       "2002",
	Executable: "libreoffice",
}

func SetExecutable(executable string) {
	unoserver.SetExecutable(executable)
}

func SetInterface(interf string) {
	unoserver.SetInterface(interf)
}

func SetPort(port string) {
	unoserver.SetPort(port)
}

func SetUserInstallation(userInstallation string) {
	unoserver.SetUserInstallation(userInstallation)
}

func Run() error {
	return unoserver.Run()
}

func RunContext(ctx context.Context) error {
	return unoserver.RunContext(ctx)
}

type Unoserver struct {
	Interface        string
	Port             string
	Executable       string
	UserInstallation string
}

func (u *Unoserver) SetExecutable(executable string) {
	u.Executable = executable
}

func (u *Unoserver) SetInterface(interf string) {
	u.Interface = interf
}

func (u *Unoserver) SetPort(port string) {
	u.Port = port
}

func (u *Unoserver) SetUserInstallation(userInstallation string) {
	path, _ := url.JoinPath("file://", userInstallation)
	u.UserInstallation = path
}

func (u *Unoserver) Run() error {
	connections := fmt.Sprintf(DefaultConnection, u.Interface, u.Port)

	var args = []string{}
	args = append(args, DefaultLibreofficeOptions...)

	// Set UserInstallation path
	args = WithUserInstallation(u, args)

	// Add uno connection parameters
	args = append(
		args,
		fmt.Sprintf("--accept=%s", connections),
	)

	log.Printf("Command: %s %s", u.Executable, args)
	cmd := exec.Command(u.Executable, args...)
	return cmd.Run()
}

func (u *Unoserver) RunContext(ctx context.Context) error {
	connections := fmt.Sprintf(DefaultConnection, u.Interface, u.Port)

	var args = []string{}
	args = append(args, DefaultLibreofficeOptions...)

	// Set UserInstallation path
	args = WithUserInstallation(u, args)

	// Add uno connection parameters
	args = append(
		args,
		fmt.Sprintf("--accept=%s", connections),
	)

	log.Printf("Command: %s %s", u.Executable, args)
	cmd := exec.CommandContext(ctx, u.Executable, args...)
	return cmd.Run()
}

func WithUserInstallation(u *Unoserver, args []string) []string {
	if u.UserInstallation != "" {
		args = append(
			args,
			fmt.Sprintf("-env:UserInstallation=%s", u.UserInstallation),
		)
	}

	return args
}
