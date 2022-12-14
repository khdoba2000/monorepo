# setup
    install direnv
    add 'eval "$(direnv hook zsh)"' to your ~/.zshrc file. Other shell options at - https://direnv.net/docs/hook.html

# To run jeager tracer
    docker run -d -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest

 # To see traces open and hit "Find Traces"
    http://localhost:16686

# To run gazelle
    bazel run //:gazelle

# To generate protos
    1. get all proto target names
        make list-proto-targets
    2. run the target you want
        bazel run (targetName)

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

# To show a list of collable targets
    bazel query //src/... --output label_kind | sort | column -t




# To disable auto proto generation add '# gazelle:proto disable_global' to root BUILD file