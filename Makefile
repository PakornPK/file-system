run:
	go run main.go
push:
	git add . && git commit -m "$(commit)" && git push origin main