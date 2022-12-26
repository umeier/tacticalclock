PLATFORMS := linux/amd64 windows/amd64 linux/arm linux/arm64

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

LDFLAGS = -s -w

.PHONY: all
all: frontend backend

.PHONY: backend
backend: frontend release compress

.PHONY: release
release: $(PLATFORMS)

$(PLATFORMS):
	cd backend && GOOS=$(os) GOARCH=$(arch) go build -ldflags '$(LDFLAGS)' -o 'dist/$(os)-$(arch)/'

.PHONY: compress
compress:
	find backend/dist/ -type f -exec upx {} \;

.PHONY: frontend frontend-build frontend-install
frontend: frontend-install frontend-build

frontend-install:
	cd tacticalclock && npm install

frontend-build:
	cd tacticalclock && npm run build

.PHONY: clean
clean:
	cd backend && rm -rf dist/ && rm -rf frontend/