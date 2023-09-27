module test04

go 1.20


require (
	test04/gee v0.0.0
)

replace (
	test04/gee => ./gee
)