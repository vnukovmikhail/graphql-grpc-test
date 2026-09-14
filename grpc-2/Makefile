run-client:
	go run ./cmd/client/

run-server:
	go run ./cmd/server/

test-by-grpcurl:
	grpcurl -plaintext -d '{"email":"test@test.com","password":"123456","username":"test"}' \
  localhost:50051 user.v1.UserService/Register
