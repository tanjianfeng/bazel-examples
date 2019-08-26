load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "main",
    srcs = ["main.go"],
    cgo = True,
    gc_linkopts = [
        "-linkmode",
        "external",
    ],
    cdeps = [
	"@libclib//:libclib",
    ],
    #clinkopts = [
    #    "-L./bazel-out/k8-fastbuild/genfiles/external/libclib/",
    #    "-L./bazel-out/k8-fastbuild/bin/external/libclib/",
    #    "-L./bazel-out/k8-dbg/genfiles/external/libclib/",
    #    "-L./bazel-out/k8-dbg/bin/external/libclib/",
    #    "-Wl,--whole-archive -lclib -Wl,--no-whole-archive",
    #],
)
