all: spec/*
	mkdir -p $(notdir $(basename $^))
	$(foreach file, $^, protoc -I spec $(file) --go_out=$(notdir $(basename $(file))) --go_opt=paths=source_relative --go-grpc_out=$(notdir $(basename $(file))) --go-grpc_opt=paths=source_relative;)
