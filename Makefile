all: day-1 day-2

day-1:
	@echo "Day 1"
	@docker run --rm -v $(PWD):/app -w /app golang:alpine go run days/day-1/main.go

day-2:
	@echo "Day 2"
	@docker run --rm -v $(PWD):/app -w /app golang:alpine go run days/day-2/main.go
