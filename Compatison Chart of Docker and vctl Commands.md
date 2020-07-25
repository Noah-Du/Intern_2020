| Docker  | vctl     | Help                                                         |
| ------- | -------- | ------------------------------------------------------------ |
| build   | build    | Build a container image from a Dockerfile.                   |
| create  | create   | Create a new container from a container image.               |
| exec    | exec     | Run a command in a running container                         |
| images  | images   | List container images.                                       |
| ps      | ps       | List containers.                                             |
| pull    | pull     | Pull a container image from a registry.                      |
| push    | ush      | Push a container image to a registry.                        |
| rm      | rm       | Remove one or more containers.                               |
| rmi     | rmi      | Remove one or more container images.                         |
| run     | run      | Run a new container from a container image.                  |
| start   | start    | Start an existing container.                                 |
| stop    | stop     | Stop a container.                                            |
| tag     | tag      | Tag container images.                                        |
| ersion  | version  | Show the version of docker/vctl.                             |
|         | describe | Show details of a container.                                 |
|         | execvm   | Execute a command within a running virtual machine that hosts container. |
|         | help     | Help about any command.                                      |
|         | system   | Manage the Nautilus Container Engine.                        |
| attach  |          | Attach local standard input, output, and error streams to a running container |
| commit  |          | Create a new image from a container's changes                |
| cp      |          | Copy files/folders between a container and the local filesystem |
| diff    |          | Inspect changes to files or directories on a container's filesystem |
| events  |          | Get real time events from the server                         |
| export  |          | Export a container's filesystem as a tar archive             |
| history |          | Show the history of an image                                 |
| import  |          | Import the contents from a tarball to create a filesystem image |
| info    |          | Display system-wide information                              |
| inspect |          | Return low-level information on Docker objects               |
| kill    |          | Kill one or more running containers                          |
| load    |          | Load an image from a tar archive or STDIN                    |
| login   |          | Log in to a Docker registry                                  |
| logout  |          | Log out from a Docker registry                               |
| logs    |          | Fetch the logs of a container                                |
| pause   |          | Pause all processes within one or more containers            |
| port    |          | List port mappings or a specific mapping for the container   |
| rename  |          | Rename a container                                           |
| restart |          | Restart one or more containers                               |
| save    |          | Save one or more images to a tar archive (streamed to STDOUT by default) |
| search  |          | Search the Docker Hub for images                             |
| stats   |          | Display a live stream of container(s) resource usage statistics |
| top     |          | Display the running processes of a container                 |
| unpause |          | Unpause all processes within one or more containers          |
| update  |          | Update configuration of one or more containers               |
| wait    |          | Block until one or more containers stop, then print their exit codes |



Docker Options:

|      |                    |                                                              |
| ---- | ------------------ | ------------------------------------------------------------ |
|      | --config string    | Location of client config files                              |
| -c,  | --context string   | Name of the context to use to connect to the daemon (overrides DOCKER_HOST env var and default context set with "docker context use") |
| -D,  | --debug            | Enable dubug mode                                            |
| -H,  | --host list        | Daemon socket(s) to connect to                               |
| -l,  | --log-level string | Set the logging level("debug"\|"info"\|"warn"\|"error"\|"fatal")(default "info") |
|      | --tls              | Use TLS; implied by --tlsverify                              |
|      | --tlscacert string | Trust certs signed only by this CA (default "/root/.docker/ca.pem") |
|      | --tlscert string   | Path to TLS certificate file (default "/root/.docker/cert.pem") |
|      | --tlskey string    | Path to TLS key file (default "/root/.docker/key.pem")       |
|      | --tlsverify        | Use TLS and verify the remote                                |
| -v,  | --version          | Print version information and quit                           |



Docker Management Commands:

| builder   | Manage builds                 |
| --------- | ----------------------------- |
| config    | Manage Docker configs         |
| container | Manager containers            |
| context   | Manage contexts               |
| engine    | Manage the docker engine      |
| image     | Manage images                 |
| network   | Manage networks               |
| node      | Manage Swarm nodes            |
| plugin    | Manage plugins                |
| secret    | Manage Docker secrets         |
| service   | Manage Docker stackes         |
| swarm     | Manage Docker                 |
| trust     | Manage trust on Docker images |
| volume    | Manage volumes                |

