package main

import (
    "os"
    "fmt"
    "strings"
    "os/exec"
)

func get_docker_creds() map[string]string {
    var docker_creds = map[string]string{
        "username": os.Getenv("DOCKER_USERNAME"),
        "password": os.Getenv("DOCKER_PASSWORD"),
    }
    return docker_creds
}

func get_docker_image_id(tag string) string {
    out, err := exec.Command(
        "docker",
        "image",
        "ls",
        "|grep",
        strings.Split(tag, ":")[0],
    ).Output()
    if err != nil {
        fmt.Println(err)
    }
    return string(out)
}

func docker_login(username string, password string) {
    cmd := exec.Command(
        "docker",
        "login",
        "-u",
        username,
        "-p",
        password,
    )
    err := cmd.Run()
    if err != nil {
        fmt.Println(err)
    }
}

func docker_build(dockerfile string, tag string) {
    cmd := exec.Command(
        "docker",
        "build",
        dockerfile,
        "-t",
        tag,
    )
    err := cmd.Run()
    if err != nil {
        fmt.Println(err)
    }
}

func docker_push(tag string) {
    var docker_creds = get_docker_creds()
    docker_login(docker_creds["username"], docker_creds["password"])

    cmd := exec.Command(
        "docker",
        "push",
        get_docker_image_id(tag),
    )
    err := cmd.Run()
    if err != nil {
        fmt.Println(err)
    }
}