ifeq ($(OS), Windows_NT)
	OS_IS_WINDOWS = true
else
	OS_IS_WINDOWS = false
endif

all: main.go
	go build $(if OS_IS_WINDOWS,,-ldflags="-s -w" )-buildmode=pie -o msattack$(if OS_IS_WINDOWS,.exe,) main.go
	chmod +x msattack$(if OS_IS_WINDOWS,.exe,)

clean:
	rm -f msattack$(if OS_IS_WINDOWS,.exe,)