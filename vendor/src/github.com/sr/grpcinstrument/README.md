grpcinstrument
==============

This is a golang package that helps instrumenting [gRPC][] servers. It includes
a [Protocol Buffer][pb] compiler plugin for generating code that wraps and
instruments gRPC servers and an implementation of the `Instrumentator` interface
that logs requests to any of the backends supported by [protolog][] and exposes
metrics about RPC calls via [Prometheus][].

[gRPC]: http://www.grpc.io
[pb]: https://developers.google.com/protocol-buffers/?hl=en
[protolog]: https://github.com/peter-edge/go-protolog
[Prometheus]: http://prometheus.io

[Checkout the godoc][godoc] online or download this package and view them
locally:

    $ go get github.com/sr/grpcinstrument/...

[godoc]: https://godoc.org/github.com/sr/grpcinstrument

Assuming your proto files are located under `src/`, invoke the `protoc` command
like so:

    $ protoc --grpcinstrument_out=src/ src/*.proto
    $ ls src/*-gen.go
    src/instrumented_buildkite-gen.go

Then update your main package to use the generated instrumented code:

```diff
--- a.go        2015-12-15 13:01:03.307634424 +0000
+++ b.go        2015-12-15 13:00:40.386861334 +0000
@@ -1,3 +1,5 @@
 server := grpc.NewServer()
+instrumentator := myapp.NewInstrumentator()
 buildkiteServer, := buildkite.NewAPIServer()
-buildkite.RegisterBuildkiteServiceServer(server, buildkiteServer)
+instrumented := buildkite.NewInstrumentedBuildkiteServiceServer(instrumentator, buildkiteServer)
+buildkite.RegisterBuildkiteServiceServer(server, instrumented)
```