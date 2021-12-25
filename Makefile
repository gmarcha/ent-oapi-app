NAME := bin/api/api.out

all:
		go build -o $(NAME) cmd/api/main.go && ./$(NAME)