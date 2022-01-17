package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema
//go:generate bash -c "OPTION_GO_PACKAGE=github.com/pigfall/curdboy_playground/api/v1/contacts go run -mod=mod entgo.io/contrib/entproto/cmd/entproto -path ./schema -targetPath ../api/v1"
