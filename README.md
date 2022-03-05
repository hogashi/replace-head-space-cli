# replace-head-space-cli

## Install

```
go install github.com/hogashi/replace-head-space-cli@latest
```

## Usage

spaces (hard tab, full-width space) on head of lines will be replaced into half-width space

```
$ cat hoge.txt
ab
 cd
	ef
		gh
	 ij
	　kl
$ cat hoge.txt | replace-head-space-cli
ab
 cd
 ef
  gh
  ij
  kl
```
