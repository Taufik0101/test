git clone https://github.com/Taufik0101/test.git
cd test/
cp .env.example .env
Update .env and set your database credentials (set APP_ENV as DEVELOPMENT if you first use / clone)
go get ./...
go run server.go