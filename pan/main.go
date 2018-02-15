package main

// CTRL + p + q to quit a attached Docker container without killing it.

import (
    "os"
    "log"

    "github.com/docker/docker/api/types"
    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/client"
    "golang.org/x/net/context"
)

func DockerCreate() string {
    ctx := context.Background()
    cli, err := client.NewEnvClient()
    if err != nil {
        panic(err)
    }

    _, err = cli.ImagePull(
        ctx,
        "docker.io/library/alpine",
        types.ImagePullOptions{},
    )
    if err != nil {
        panic(err)
    }

    resp, err := cli.ContainerCreate(
        ctx,
        &container.Config{
            Image:          "alpine",
            Cmd:            []string{"/bin/sh"},
            Tty:            true,
            OpenStdin:      true,
            StdinOnce:      false,
        },
        &container.HostConfig{
            NetworkMode:    "host",
        },
        nil,
        "",
    )
    if err != nil {
        panic(err)
    }

    if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
        panic(err)
    }

    return resp.ID[0:12]
}

func DockerCommit(containerID string, containerRef string) {
    ctx := context.Background()
    cli, err := client.NewEnvClient()
    if err != nil {
        panic(err)
    }

    Resp, err := cli.ContainerCommit(
        ctx,
        containerID,
        types.ContainerCommitOptions{
            Reference: containerRef,
        },
    )
    if err != nil {
        panic(err)
    }

    //log.Println(Resp.ID[0:12])
    return resp.ID[0:12]
}

func DockerAttach(containerID string) {
    /*ctx := context.Background()
    cli, err := client.NewEnvClient()
    if err != nil {
        panic(err)
    }

    conn, err := cli.ContainerAttach(ctx, containerID, types.ContainerAttachOptions{})
    if err != nil {
        panic(err)
    }

    writer := bufio.NewScanner(conn)
    log.Println(writer.Text())*/
}

func main() {
    if len(os.Args) < 2 {
        log.Println("Usage: pan [option]")
        os.Exit(1)
    } else {
        if os.Args[1] == "dev" {
            containerID := DockerCreate()
            //DockerAttach(containerID)
            log.Println(containerID)
        } else if os.Args[1] == "build" {
            if os.Args[2] != "" && os.Args[3] != "" {
                DockerCommit(os.Args[2], os.Args[3])
            }
        }
    }
}