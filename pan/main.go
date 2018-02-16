package main

// CTRL + p + q to quit a attached Docker container without killing it.

import (
    "os"
    "fmt"
    //"log"

    "github.com/urfave/cli"

    "github.com/docker/docker/api/types"
    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/client"
    "golang.org/x/net/context"

    //"time"
    //"github.com/yhat/go-docker"
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

    /*if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
        panic(err)
    }*/

    return resp.ID[0:12]
}

func DockerCommit(containerID string, containerRef string) string {
    ctx := context.Background()
    cli, err := client.NewEnvClient()
    if err != nil {
        panic(err)
    }

    resp, err := cli.ContainerCommit(
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

func DockerStart(containerID string) {
    ctx := context.Background()
    cli, err := client.NewEnvClient()
    if err != nil {
        panic(err)
    }
    if err := cli.ContainerStart(
        ctx,
        containerID,
        types.ContainerStartOptions{},
    ); err != nil {
        panic(err)
    }
}

/*func main() {
    if len(os.Args) < 2 {
        log.Println("Usage: pan [option]")
        os.Exit(1)
    } else {
        if os.Args[1] == "dev" {
            containerID := DockerCreate()
            log.Println(containerID)
        } else if os.Args[1] == "build" {
            if os.Args[2] != "" && os.Args[3] != "" {
                DockerCommit(os.Args[2], os.Args[3])
            }
        }
    }
}*/

/*func DockerAttach(containerID string) {
    timeout := 3 * time.Second
    cli, err := docker.NewDefaultClient(timeout)
    streamOpts := &docker.AttachOptions{Stream: true, Stdout: true, Stderr: true}
    stream, err := cli.Attach(containerID, streamOpts)
    if err != nil {
        panic(err)
    }
    defer stream.Close()
    go docker.SplitStream(stream, os.Stdout, os.Stderr)
}*/

/*func Test() error {
    timeout := 3 * time.Second

    cli, err := docker.NewDefaultClient(timeout)
    if err != nil {
        return err
    }

    // create a container
    config := &docker.ContainerConfig{
        Image: "alpine",
        Cmd:   []string{"/bin/sh"},
        Tty:            true,
        OpenStdin:      true,
        StdinOnce:      false,
    }
    cid, err := cli.CreateContainer(config, "myimage")
    if err != nil {
        return err
    }

    // always remember to clean up after yourself
    //defer cli.RemoveContainer(cid, true, false)

    // attach to the container
    streamOpts := &docker.AttachOptions{
        Stream: true,
        Stdout: true,
        Stderr: true,
    }
    stream, err := cli.Attach(cid, streamOpts)
    if err != nil {
        return err
    }
    defer stream.Close()

    // concurrently write stream to stdout and stderr
    go docker.SplitStream(stream, os.Stdout, os.Stderr)

    // start the container
    err = cli.StartContainer(cid, &docker.HostConfig{})
    if err != nil {
        return err
    }

    // wait for the container to exit
    statusCode, err := cli.Wait(cid)
    if err != nil {
        return err
    }
    if statusCode != 0 {
        return fmt.Errorf("process returned bad status code: %d", statusCode)
    }

    return nil
}*/

func DockerRemove(containerID string) {
    //ctx := context.Background()
    cli, err := client.NewEnvClient()
    if err != nil {
        panic(err)
    }
    cli.ContainerRemove(
        context.Background(),
        containerID,
        types.ContainerRemoveOptions{
            RemoveVolumes: true,
            Force:         true,
        },
    )
}

func main() {
  app := cli.NewApp()
  app.Name = "pan"
  app.Usage = "fight the loneliness!"

  app.Commands = []cli.Command{
    {
        Name: "dev",
        Aliases: []string{"dev"},
        Usage: "options for dev environment",
        Subcommands: []cli.Command{
            {
                Name: "add",
                Usage: "create a new dev environment",
                Action: func(c *cli.Context) error {
                    /*if err := Test(); err != nil {
                        panic(err)
                    }*/
                    containerID := DockerCreate()
                    //DockerAttach(containerID)
                    DockerStart(containerID)
                    fmt.Println("New dev container created: " + containerID)
                    return nil
                },
            },
            {
                Name: "rm",
                Usage: "delete a dev environment",
                Action: func(c *cli.Context) error {
                    DockerRemove(c.Args().First())
                    fmt.Println("Deleted dev container: " + c.Args().First())
                    return nil
                },
            },
        },
    },
  }

  app.Run(os.Args)
}