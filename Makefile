.DEFAULT_GOAL=paper-airplane

BUILD=paper-airplane.exe

paper-airplane:
	go build -o $(BUILD) .

run: paper-airplane
	./$(BUILD)

clean:
	rm -rf ./log
	rm -rf $(BUILD)

.PHONY: run paper-airplane