LOCAL_BIN:=$(CURDIR)/bin

docker-build:
	docker buildx build --no-cache --platform linux/amd64 -t http_test_server:v0.0.1 .

http-load-test:
	wrk -t1 -c1 -d30s -s programs.lua http://localhost:8080/programs/5 --latency

http-load-ab:
	ab -n 10 -k https://be3e-93-100-98-132.ngrok-free.app/programs/21

grpc-load-test:
	ghz \
		--proto api/program_v3/program_v3.proto \
		--call program_v3.ProgramV3.Get \
		--data '{"Count": 5}' \
		--rps 100 \
		--total 3000 \
		--insecure \
		localhost:50051