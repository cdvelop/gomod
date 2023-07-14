module github.com/cdvelop/module_client

go 1.20

require (
	github.com/cdvelop/model v0.0.32
)

require github.com/cdvelop/cutkey v0.6.0 // indirect

require (
	github.com/cdvelop/api v0.0.1
	github.com/cdvelop/js v0.0.0-20230710200247-0d9f5c8314da
	github.com/cdvelop/platform v0.0.2
	golang.org/x/text v0.10.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/godev => ../godev