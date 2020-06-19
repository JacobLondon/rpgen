GO=go

TARGET = city

FILES = \
	city.go \
	main.go

.PHONY: clean

$(TARGET):
	$(GO) build $(FILES)

clean:
	rm -rf $(TARGET)
