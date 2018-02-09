package main

// TODO : Add tests for docker routes

import (
    "io"
    "os"
    "fmt"
    "strings"
    "net/http"
    "os/exec"
    "archive/tar"
    "bytes"
    "io/ioutil"
    "encoding/json"
    "log"

    "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
    "golang.org/x/net/context"
)

func downloadFile(url string) {
    tokens := strings.Split(url, "/")
    fileName := tokens[len(tokens)-1]

    output, err := os.Create(fileName)
    if err != nil {
        fmt.Println("Error while creating", fileName, "-", err)
        return
    }
    defer output.Close()

    response, err := http.Get(url)
    if err != nil {
        fmt.Println("Error while downloading", url, "-", err)
        return
    }
    defer response.Body.Close()

    n, err := io.Copy(output, response.Body)
    if err != nil {
        fmt.Println("Error while downloading", url, "-", err)
        return
    }
    fmt.Println(n)
}

func deleteFile(path string) {
    os.Remove(path)
}

func get_docker_creds() map[string]string {
    var docker_creds = map[string]string{
        "username": os.Getenv("DOCKER_USERNAME"),
        "password": os.Getenv("DOCKER_PASSWORD"),
    }
    return docker_creds
}

func dockerImageList(w http.ResponseWriter, r *http.Request) {
    cli, err := client.NewEnvClient()
    if err != nil {
        panic(err)
    }

    images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
    if err != nil {
        panic(err)
    }

    //for _, image := range images {
    //    fmt.Println(image.ID)
    //}
    if err := json.NewEncoder(w).Encode(images); err != nil {
        panic(err)
    }
}

// TODO : Implement it in Go
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

// TODO : Implement it in Go
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

func DockerBuild(w http.ResponseWriter, r *http.Request) {

    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    defer r.Body.Close()

    m := make(map[string]string)
    json.Unmarshal(body, &m)

    downloadFile(m["url"])

    ctx := context.Background()
    cli, err := client.NewEnvClient()
    if err != nil {
        log.Fatal(err, " :unable to init client")
    }

    buf := new(bytes.Buffer)
    tw := tar.NewWriter(buf)
    defer tw.Close()

    dockerFile := "Dockerfile"
    dockerFileReader, err := os.Open("Dockerfile")
    if err != nil {
        log.Fatal(err, " :unable to open Dockerfile")
    }
    readDockerFile, err := ioutil.ReadAll(dockerFileReader)
    if err != nil {
        log.Fatal(err, " :unable to read dockerfile")
    }

    tarHeader := &tar.Header{
        Name: dockerFile,
        Size: int64(len(readDockerFile)),
    }
    err = tw.WriteHeader(tarHeader)
    if err != nil {
        log.Fatal(err, " :unable to write tar header")
    }
    _, err = tw.Write(readDockerFile)
    if err != nil {
        log.Fatal(err, " :unable to write tar body")
    }
    dockerFileTarReader := bytes.NewReader(buf.Bytes())

    imageBuildResponse, err := cli.ImageBuild(
        ctx,
        dockerFileTarReader,
        types.ImageBuildOptions{
            Context:    dockerFileTarReader,
            Dockerfile: dockerFile,
            Remove:     true})
    if err != nil {
        log.Fatal(err, " :unable to build docker image")
    }
    defer imageBuildResponse.Body.Close()
    _, err = io.Copy(os.Stdout, imageBuildResponse.Body)
    if err != nil {
        log.Fatal(err, " :unable to read image build response")
    }

    deleteFile("Dockerfile")
}

// TODO : Implement it in Go
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