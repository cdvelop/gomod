module github.com/cdvelop/gomod

go 1.20

require (
	github.com/cdvelop/fileserver v0.0.60
	github.com/cdvelop/strings v0.0.9
)

require (
	github.com/cdvelop/filehandler v0.0.41 // indirect
	github.com/cdvelop/input v0.0.83 // indirect
	github.com/cdvelop/maps v0.0.8 // indirect
	github.com/cdvelop/model v0.0.113 // indirect
	github.com/cdvelop/object v0.0.71 // indirect
	github.com/cdvelop/structs v0.0.1 // indirect
	github.com/cdvelop/timetools v0.0.34 // indirect
	github.com/cdvelop/unixid v0.0.49 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	golang.org/x/net v0.19.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/unixid => ../unixid

replace github.com/cdvelop/object => ../object

replace github.com/cdvelop/maps => ../maps

replace github.com/cdvelop/fileserver => ../fileserver

replace github.com/cdvelop/filehandler => ../filehandler

replace github.com/cdvelop/input => ../input

replace github.com/cdvelop/gotools => ../gotools

replace github.com/cdvelop/strings => ../strings

replace github.com/cdvelop/timetools => ../timetools
