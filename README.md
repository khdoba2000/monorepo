# To run gazelle
    bazel run //:gazelle

# To build
    bazel build //...

# To create docker image of a service
    bazel run //src/service-one:image
    bazel run //src/service-two:image
    bazel run //src/api_gateway:image
    bazel run //src/auth_service:image    

# To run all services
    docker compose up

# To run services in docker container
    docker run --rm -it -p8000:8080 khdoba/src/service-one:image   
    docker run --rm -it -p8001:8082 khdoba/src/service-two:image
    docker run --rm -it -p8002:8083 khdoba/src/api_gateway:image
    docker run --rm -it -p8003:8084 khdoba/src/auth_service:image   

# To run the whole app
    docker compose up
    
# To sync with go mod imports
    bazel run //:gazelle-update-repos

# To run service1
    bazel run //src/service-one

# To run service2
    bazel run //src/service-two 

# To run api_gateway
    bazel run //src/api_gateway

# To run auth_service:
    bazel run //src/auth_service 


# To run all tests
    bazel test //... --test_output=errors  



# To disable auto proto generation add '# gazelle:proto disable_global' to root BUILD file