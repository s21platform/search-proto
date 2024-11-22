protoc  --go_out=./  \
        --go-grpc_out=./ \
        search.proto

protoc --doc_out=. --doc_opt=markdown,README.md ./search.proto
