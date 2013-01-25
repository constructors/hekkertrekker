ht: *.go
	go build -o ht *.go

clean:
	rm -f ht

all: ht

.PHONY: clean all
