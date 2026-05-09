bin := "radar"

default:
    just --list

run:
    go run .

build:
    go build -o {{ bin }} .

clean:
    rm -f {{ bin }}
