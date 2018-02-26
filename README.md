tempfile
========
An implementation of
[`ioutil.TempFile`](https://golang.org/pkg/io/ioutil/#TempFile), but also with
an additional parameter for a deterministic suffix.

This is a workaround for https://github.com/golang/go/issues/4896.
