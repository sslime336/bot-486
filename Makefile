.DEFAULT_GOAL=bot

BUILD=bot-486.exe

bot:
	go build -o $(BUILD) .

run: bot
	./$(BUILD)

release-linux-amd64:
	GOARCH=amd64 GOOS=linux go build -o build/$(BUILD) .

clean:
	rm -rf ./log
	rm -rf ./build
	rm -rf $(BUILD)

.PHONY: run bot