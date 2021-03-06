package platform

// Code generated by sangkuriang protoc-gen-go. DO NOT EDIT.
// source: platform.proto_platform
// File Location: Platform.model.go
import "time"
// Platform for struct info
type Platform struct {
	Id			int64 ` validate:"required" decorator:"EQUAL"`
	Name			string ` validate:"required" decorator:"EQUAL"`
	Email			string ` validate:"required,email" decorator:"EQUAL"`
	CreatedBy			string `decorator:"EQUAL"`
	CreatedAt			time.Time `decorator:"EQUAL"`
}
// ResponseCreatePlatform for struct info
type ResponseCreatePlatform struct {
	Id			int64 `decorator:"EQUAL"`
	Name			string `decorator:"EQUAL"`
	Email			string `decorator:"EQUAL"`
	CreatedAt			time.Time `decorator:"EQUAL"`
}


