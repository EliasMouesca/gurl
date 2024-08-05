
EXEC=gurl

$(EXEC): src/*
	go build -o $@ $^

