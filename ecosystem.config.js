module.exports = {
    apps: [
        {
            name: "gateway",
            script: "go run cmd/gateway/main.go",
        },
        {
            name: "catalog",
            script: "go run cmd/catalog/main.go",
        },
        {
            name: "reading",
            script: "go run cmd/reading/main.go",
        },
        {
            name: "static",
            script: "go run cmd/static/main.go",
        },
    ]
}