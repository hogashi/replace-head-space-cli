# replace-head-space-cli

## Install

```
go install github.com/hogashi/replace-head-space-cli@latest
```

## Usage

spaces (hard tab, full-width space) on head of lines will be replaced into half-width space

```
$ cat hoge.txt
0 spaces
 1 half-width space
　1 full-width space
	1 hard tab
	 　3 mixed spaces
$ cat hoge.txt | replace-head-space-cli
0 spaces
 1 half-width space
 1 full-width space
 1 hard tab
   3 mixed spaces
```
