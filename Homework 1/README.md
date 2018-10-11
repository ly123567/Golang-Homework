### 编译运行流程
```shell
go get github.com/spf13/pflag
go install selpg.go
./selpg [flags]
```
命令相关格式见测试结果，测试环境为MacOS，可能对于不同系统稍有变动。

### 测试结果

```shell
MacBookPro:~ Administrator$ cd Desktop/
MacBookPro:Desktop Administrator$ ./selpg -s 1 -e 2 -l 2 test.txt
01
02
03
04
MacBookPro:Desktop Administrator$ ./selpg -s 1 -e 2 -l 2 < test.txt
01
02
03
04
MacBookPro:Desktop Administrator$ ./selpg -s 1 -e 2 -l 2 test.txt > out.txt
MacBookPro:Desktop Administrator$ cat out.txt
01
02
03
04
MacBookPro:Desktop Administrator$ rm out.txt
MacBookPro:Desktop Administrator$ ./selpg -s 1 -e 0 -l 2 test.txt 2> error.txt
MacBookPro:Desktop Administrator$ cat error.txt 
Usage: ./selpg [-s start] [-e end] [optional...] [filepath]
  -d, --dest string   Select the output file path
  -e, --end int       End page of showing pages (default -1)
  -h, --help          Show this message
  -l, --length int    The line-length of each pages (default 72)
  -s, --start int     Start page of showing pages (default -1)
  -f, --type          Divide pages by page break

Start page should less than end page!
MacBookPro:Desktop Administrator$ rm error.txt 
MacBookPro:Desktop Administrator$ ./selpg -s 1 -e 2 -l 2 test.txt | grep -i "04"
04
MacBookPro:Desktop Administrator$ ./selpg -s 1 -e 1 -f pages.txt
there have a page break
```