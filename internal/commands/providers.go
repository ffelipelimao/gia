package commands

func getProvider(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return "default"
}
