PLATFORMS := linux/amd64 windows/amd64 linux/arm linux/arm64

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

LDFLAGS = -s -w

.PHONY: all
all: frontend backend

.PHONY: backend
backend: frontend release

.PHONY: release
release: $(PLATFORMS)

$(PLATFORMS):
	cd backend && GOOS=$(os) GOARCH=$(arch) go build -ldflags '$(LDFLAGS)' -o 'dist/$(os)-$(arch)/'

.PHONY: frontend-build frontend-install
frontend: frontend-install frontend-build

frontend-install:
	cd tacticalclock && npm install

frontend-build:
	cd tacticalclock && npm run build

.PHONY: clean
clean:
	cd backend && rm -rf dist/ && rm -rf frontend/