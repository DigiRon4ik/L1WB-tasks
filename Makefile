.DEFAULT_GOAL := run
.PHONY: run l1 l2 l3 l4 l5 l6 l7 l8 l9 l10 l11 l12 l13 l14 l15 l16 l17 l18 l19 l20 l21 l22 l23 l24 l25 l26

l1:
	@go run L1.1/main.go
l2:
	@go run L1.2/main.go
l3:
	@go run L1.3/main.go
l4:
	@go run L1.4/main.go 4
l5:
	@go run L1.5/main.go
l6: l6-4
	@go run L1.6/Solution-5/main.go
l7:
	@go run L1.7/main.go
l8:
	@go run L1.8/main.go
l9:
	@go run L1.9/main.go
l10:
	@go run L1.10/main.go
l11:
	@go run L1.11/main.go
l12:
	@go run L1.12/main.go
l13:
	@go run L1.13/main.go
l14:
	@go run L1.14/main.go
l15:
	@go run L1.15/main.go
l16:
	@go run L1.16/main.go
l17:
	@go run L1.17/main.go
l18:
	@go run L1.18/main.go
l19:
	@go run L1.19/main.go
l20:
	@go run L1.20/main.go
l21:
	@go run L1.21/main.go
l22:
	@go run L1.22/main.go
l23:
	@go run L1.23/main.go
l24:
	@go run L1.24/main.go
l25:
	@go run L1.25/main.go
l26:
	@go run L1.26/main.go

run:
	echo "Run one of the tasks. Example: make l1"

l6-1:
	@go run L1.6/Solution-1/main.go
l6-2: l6-1
	@go run L1.6/Solution-2/main.go
l6-3: l6-2
	@go run L1.6/Solution-3/main.go
l6-4: l6-3
	@go run L1.6/Solution-4/main.go