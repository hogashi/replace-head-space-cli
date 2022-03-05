# replace-head-space-cli

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
$ cat hoge.txt | go run .
ab
 cd
 ef
  gh
  ij
  kl
```
