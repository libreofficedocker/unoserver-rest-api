package unoconvert

var Unoconvert = &UnoconvertOption{
	Interface:  "127.0.0.1",
	Port:       "2002",
	Executable: "unoconvert",
}

func SetExecutable(executable string) {
	Unoconvert.SetExecutable(executable)
}

func SetInterface(interf string) {
	Unoconvert.SetInterface(interf)
}

func SetPort(port string) {
	Unoconvert.SetPort(port)
}

func Run(infile string, outfile string, opts ...string) {
	Unoconvert.Run(infile, outfile, opts...)
}

type UnoconvertOption struct {
	Interface  string
	Port       string
	Executable string
}

func (u *UnoconvertOption) SetExecutable(executable string) {
	u.Executable = executable
}

func (u *UnoconvertOption) SetInterface(interf string) {
	u.Interface = interf
}

func (u *UnoconvertOption) SetPort(port string) {
	u.Port = port
}

func (u *UnoconvertOption) Run(infile string, outfile string, opts ...string) {
	// Run
}
