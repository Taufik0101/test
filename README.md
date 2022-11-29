git clone https://github.com/Taufik0101/test.git
cd test/
cp .env.example .env
Update .env and set your database credentials
go get ./...
go run server.go