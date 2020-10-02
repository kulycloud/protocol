all: spec/*
	mkdir -p $(notdir $(basename $^))
	$(foreach file, $^, protoc -I spec $(file) --go_out=plugins=grpc:$(notdir $(basename $(file))) --go_opt=paths=source_relative;)
