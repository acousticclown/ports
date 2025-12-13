package ui

func PrintSuccess(msg string) {
	Success.Println("✅ " + msg)
}

func PrintError(msg string) {
	Error.Println("❌ " + msg)
}
