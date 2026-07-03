alias rc := rust-clean
alias cc := c-clean
alias gc := go-clean
alias tc := ts-clean

default:
    just --list

# 在指定子目录中执行命令
_run dir cmd:
    cd {{dir}} && {{cmd}}

rust-clean:
    just _run rust "cargo clean"

c-clean:
    just _run c "make clean"

go-clean:
    just _run go "go clean"

ts-clean:
    just _run ts "rm -rf node_modules dist"

clean: rust-clean c-clean go-clean ts-clean
