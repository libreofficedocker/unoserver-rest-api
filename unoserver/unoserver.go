package unoserver

import (
	"fmt"
	"log"
	"os/exec"
)

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
	u.UserInstallation = userInstallation
}

func (u *Unoserver) Run() error {
	connections := fmt.Sprintf(
		"socket,host=%s,port=%s,tcpNoDelay=1;urp;StarOffice.ComponentContext",
		u.Interface, u.Port,
	)

	var args = []string{
		"--headless",
		"--invisible",
		"--nocrashreport",
		"--nodefault",
		"--nologo",
		"--nofirststartwizard",
		"--norestore",
	}

	// Set UserInstallation path
	if u.UserInstallation != "" {
		args = append(
			args,
			fmt.Sprintf("-env:UserInstallation=%s", u.UserInstallation),
		)
	}

	// Add uno connection parameters
	args = append(
		args,
		fmt.Sprintf("--accept=%s", connections),
	)

	libreoffice, err := exec.LookPath(u.Executable)
	if err != nil {
		return err
	}

	log.Printf("Command: %s %s", libreoffice, args)
	cmd := exec.Command(libreoffice, args...)
	return cmd.Start()
}
