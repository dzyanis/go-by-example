go build rand/main.go && mv main transmitter
go build main.go && mv main receiver
./transmitter | ./receiver
rm transmitter
rm receiver