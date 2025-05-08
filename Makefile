.PHONY: gen-demo-proto
gen-demo-proto:
	@cd tutorial/demo_proto && cwgo server -I ../../idl --type RPC --module github.com/charleschile/TikTok-E-Commerce-Gomall/tutorial/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd tutorial/demo_thrift && cwgo server -I ../../idl --type RPC --module github.com/charleschile/TikTok-E-Commerce-Gomall/tutorial/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift

