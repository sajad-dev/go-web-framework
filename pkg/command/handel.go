package command


type CommandType struct {
	Commad  string
	Handler func(args []string)
}

var command = []CommandType{
	CommandType{Handler: Migrate,Commad: "migration"},
}

func Handel(args []string) {
	for _, c := range command{
		if c.Commad == args[1] && len(args) > 2 {
			c.Handler(args[1:])
		}
	}
}
