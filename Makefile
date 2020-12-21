VERSION := $(shell grep "version" plugin.yaml | cut -d '"' -f 2)

build_mac: export GOARCH=amd64
build_mac:
	@GOOS=darwin go build delete_chart.go

release: build_mac
	tar -czf helm-kill_${VERSION}_darwin_amd64.tar.gz delete_chart

clean:
	rm delete_chart
	rm helm-kill_${VERSION}_darwin_amd64.tar.gz
