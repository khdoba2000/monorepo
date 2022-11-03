# To run gazelle
    bazel run //:gazelle

# To build
    bazel build //...

# To sync with go mod imports
    bazel run //:gazelle-update-repos

# To run service1
    bazel run //src/service-one/cmd:cmd          

# To run service2
    bazel run //src/service-two/cmd:cmd 


# To run all tests
    bazel test //... --test_output=errors  



# To disable auto proto generation add '# gazelle:proto disable_global' to root BUILD file