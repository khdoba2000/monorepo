PROJECT_NAME = src/customer_service

run:
	bazel run //$(PROJECT_NAME)

build:  
	bazel build //$(PROJECT_NAME)

gazelle:
	bazel run //:gazelle -- $(PROJECT_NAME)

gazelle-synch:
	bazel run //:gazelle-update-repos -- $(PROJECT_NAME)

synch:
	go mod tidy
	bazel run //:gazelle -- $(PROJECT_NAME)
	bazel run //:gazelle-update-repos -- $(PROJECT_NAME)


test:
	bazel test //$(PROJECT_NAME) --test_output=errors  