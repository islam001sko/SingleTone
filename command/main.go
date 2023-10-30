package main

func main() {
	micro := &micro{}

	onCommand := &OnCommand{
		device: micro,
	}

	offCommand := &OffCommand{
		device: micro,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()
}
