run:
	go run .

run-autoreload:
	nodemon -w . -e go -x go run . --signal SIGTERM
