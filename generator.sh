oapi-codegen -package api -generate client -o $(pwd)/typesense/api/client_gen.go $(pwd)/typesense/api/generator/generator.yml
oapi-codegen -package api -generate types -o $(pwd)/typesense/api/types_gen.go $(pwd)/typesense/api/generator/generator.yml