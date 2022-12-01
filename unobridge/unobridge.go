package unobridge

import "runtime"

var unobridge = &Unobridge{
	Interface: "127.0.0.1",
	Port:      "2002",
}

func GetInterface() string {
	return unobridge.Interface
}

func SetInterface(interf string) {
	unobridge.SetInterface(interf)
}

func GetPort() string {
	return unobridge.Port
}

func SetPort(port string) {
	unobridge.SetPort(port)
}

func GetLibreofficeBinary() string {
	if runtime.GOOS == "darwin" {
		return "soffice"
	}

	return "libreoffice"
}

type Unobridge struct {
	Interface string
	Port      string
}

func (u *Unobridge) SetInterface(interf string) {
	u.Interface = interf
}

func (u *Unobridge) SetPort(port string) {
	u.Port = port
}
