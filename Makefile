all:
	go build -o msattack$(if $(findstring Windows,$(shell uname -s)),.exe,) main.go