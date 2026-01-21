group "default" {
    targets = ["tokenapp", "issuelogger"]
}

target "auth" {
    context = "./services/auth"
    dockerfile = "Dockerfile"
    tags = ["auth:latest"]
    args = {}
}

target "api" {
    context = "./services/api"
    dockerfile = "Dockerfile"
    tags = ["api:latest"]
    args = {}
}