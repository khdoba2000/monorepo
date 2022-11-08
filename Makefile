
build:  
	bazel build //...

gazelle:
	bazel run //:gazelle

gazelle-synch:
	bazel run //:gazelle-update-repos

protogen: #https://medium.com/goc0de/a-cute-bazel-proto-hack-for-golang-ides-2a4ef0415a7f
	bazel query 'kind("proto_link", //...)'  | xargs bazel run

test:
	bazel test //... --test_output=errors  
