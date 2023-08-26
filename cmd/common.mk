MODULE = github.com/ozline/tiktok

.PHONY: target
target:
	sh build.sh

.PHONY: clean
clean:
	@find . -type d -name "output" -exec rm -rf {} + -print