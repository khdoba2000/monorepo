# To run gazelle
    bazel run //:gazelle

# To build
    bazel build //...

# To create docker image of a service
    bazel run //src/service-one:docker
    bazel run //src/service-two:docker    

# To run services in docker container
    docker run --rm -it -p8000:8080 bazel/src/service-one:docker   
    docker run --rm -it -p8001:8082 bazel/src/service-two:docker   

# To run the whole app
    docker compose up
    
# To sync with go mod imports
    bazel run //:gazelle-update-repos

# To run service1
    bazel run //src/service-one

# To run service2
    bazel run //src/service-two 


# To run all tests
    bazel test //... --test_output=errors  



# To disable auto proto generation add '# gazelle:proto disable_global' to root BUILD file