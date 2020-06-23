GO=go

TARGET=rpgen.exe

FILES = \
	people.go \
	place.go \
	town.go \
	util.go \
	main.go

.PHONY: $(TARGET) clean

$(TARGET):
	$(GO) build -o $@ $(FILES)

clean:
	rm -rf $(TARGET)
