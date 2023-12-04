GO=go
GORUN=$(GO) run

.PHONY: main
main:
	$(GORUN) main.go

.PHONY: 1
1:
	$(GORUN) main.go --num=1

.PHONY: 2
2:
	$(GORUN) main.go --num=2

.PHONY: 3
3:
	$(GORUN) main.go --num=3
