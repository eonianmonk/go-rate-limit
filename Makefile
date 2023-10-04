build:
	go build -o ratel backend/cmd/main.go 

serve:
	./ratel run

migrate-up:
	./ratel migrate up
	
