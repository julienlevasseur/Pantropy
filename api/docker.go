package main

// TODO : Add tests for docker routes

import (
    "io"
    "os"
    "fmt"
    "strings"
    "net/http"
    "archive/tar"
    "bytes"
    "io/ioutil"
    "encoding/json"
    "encoding/base64"
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

    if err := json.NewEncoder(w).Encode(images); err != nil {
        panic(err)
    }
}

func getDockerImageId(tag string) string {
    cli, err := client.NewEnvClient()
    if err != nil {
        panic(err)
    }

    images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
    if err != nil {
        panic(err)
    }

    var id string

    for _, image := range images {
        if image.RepoTags[0] == tag {
            id = strings.Split(image.ID, ":")[1]
            if err != nil {
                panic(err)
            }
        }
    }
    return id
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

    os.Remove("Dockerfile") // Remove the Dockerfile
}

func dockerPush(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    defer r.Body.Close()

    m := make(map[string]string)
    json.Unmarshal(body, &m)

    ctx := context.Background()
    cli, err := client.NewEnvClient()
    if err != nil {
        panic(err)
    }

    if os.Getenv("DOCKER_USERNAME") == "" || os.Getenv("DOCKER_PASSWORD") == "" {
        w.WriteHeader(http.StatusUnauthorized)
    }
    authConfig := types.AuthConfig{
        Username: os.Getenv("DOCKER_USERNAME"),
        Password: os.Getenv("DOCKER_PASSWORD"),
    }
    encodedJSON, err := json.Marshal(authConfig)
    if err != nil {
        panic(err)
    }
    authStr := base64.URLEncoding.EncodeToString(encodedJSON)

    imagePushResponse, err := cli.ImagePush(
        ctx,
        m["tag"],
        types.ImagePushOptions{
            RegistryAuth: authStr,
            })
    if err != nil {
        log.Fatal(err, " :unable to build docker image")
    }
    _, err = io.Copy(os.Stdout, imagePushResponse)
    if err != nil {
        log.Fatal(err, " :unable to read image build response")
    }
    defer imagePushResponse.Close()
    w.WriteHeader(http.StatusCreated)
}
