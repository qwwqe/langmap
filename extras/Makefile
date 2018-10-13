all: database.svg database.png

.PHONY: database.svg
database.svg:
	dot database.dot -Tsvg -o$@

.PHONY: database.png
database.png:
	dot database.dot -Tpng -o$@

.PHONY: clean
clean:
	rm database.png database.svg
