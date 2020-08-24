package packageA

//go:generate go run github.com/ebauman/oapi-codegen/cmd/oapi-codegen -generate types,skip-prune --package=packageA -o externalref.gen.go --import-mapping=../packageB/spec.yaml:github.com/ebauman/oapi-codegen/internal/test/externalref/packageB spec.yaml
