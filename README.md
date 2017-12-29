AesCtr
==========
[![GoDoc](https://godoc.org/github.com/bronze1man/AesCtr?status.png)](http://godoc.org/github.com/bronze1man/AesCtr)

* use the festest golang aes ctr implement with golang 1.9

### performance on my computer (mac)
* go1.9.2 aes ctr (BenchmarkAESCTR1K_goAesCtr)
```
1000000	      2336 ns/op	 436.15 MB/s
```

* CL 51670 (BenchmarkAESCTR1K_thisAesCtr)
```
3000000	       467 ns/op	2178.83 MB/s
```

### LICENSE
Unless otherwise noted, the Go source files are distributed under the BSD-style license found in the LICENSE file.

### how i get this package
* folk code from golang CL 51670 (https://go-review.googlesource.com/c/go/+/51670) (https://github.com/golang/go/issues/20967)
* change package name to AesCtr
* copy internal package from golang1.9.2
* add ctr test to this package. (aes_ctr_test.go)