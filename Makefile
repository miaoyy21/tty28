
release:
	# Windows
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o tty28_windows_386.exe

	# Linux
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o  tty28_linux_386

	# Mac M1
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o tty28_darwin_arm64