go build -o ./trans transmitter/main.go && chmod +x trans
go build -o ./receiv handler/main.go && chmod +x receiv
trans | receiv