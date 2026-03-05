module github.com/lattice-substrate/jcs-bench-lab

go 1.22.5

require (
	github.com/lattice-substrate/jcs-schubfach v0.0.0
	github.com/lattice-substrate/json-canon v0.0.0
)

replace github.com/lattice-substrate/jcs-schubfach => ./impl-schubfach

replace github.com/lattice-substrate/json-canon => ./impl-json-canon
