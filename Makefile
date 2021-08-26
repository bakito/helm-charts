index:
	cr index
	go run ./cmd/index2md/ > docs/README.md

push: index
	git add docs
	git commit -m"update index"
	git push