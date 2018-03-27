# Big Earl's CLI

`big-earl` is a CLI for interacting with the Big Earl short URL service.

  * Provides shell access to the shrink and expand services provided by the Big Earl web service
    
    ```sh-session
    $ big-earl shrink "http://google.com/search?..."
    https://big-earl.herokuapp.com/lezd
    ```

  * Copy the results directly to the clipboard

    ```sh-session
    # The new short URL will be copied to the clipboard
    $ big-earl shrink "https://google/com/search?..." -c

    # The expanded URL will be copied to the clipboard
    $ big-earl grow https://big-earl.herokuapp.com/flcf -c 

    # The URL to the short linke preview will be copied to the clipboard   
    $ big-earl preview https://big-earl.herokuapp.com/flcf -c 
    ```

  * Open the results immediately in the browser
    ```sh-session
    # The URL for the preview page will be opened in the default browser 
    $ big-earl preview https://big-earl.herokuapp.com/flcf -o
    ```

## Getting Started

The easiest way to try out the `big-earl` cli is to use the Docker image:

```sh-session
$ docker run -it --rm micahlee/big-earl-cli
```

To get the full benefit of the browser and clipboard helpers, you will want
to install `big-earl` locally on your machine. You can grab the latest build
from releases, or build it yourself using the instructions below.

## Building

To build `big-earl` you will need git, Docker, and Docker Compose.

  1. Compile the binaries with:
        ```sh-session
        $ ./build.sh
        ```

        This will place the compiled binaries for each platform in `./output`

  2. Copy the binary for your platform to a location on your PATH. For example:
        ```sh-session
        $ cp output/big-earl-darwin-amd64 /usr/local/bin/big-earl
        ```
  3. The `big-earl` command is now ready for you to use!
        ```sh-session
        $ big-earl 
        NAME:
        Big Earl's CLI - interact with Big Earl's Lean URLs API

        USAGE:
        big-earl [global options] command [command options] [arguments...]

        VERSION:
        0.0.0

        COMMANDS:
            grow, g     expands a short URL to its original URL
            preview, p  provides the preview URL for a given short URL
            shrink, s   gets a new short URL for the given input URL
            help, h     Shows a list of commands or help for one command

        GLOBAL OPTIONS:
        --big_earl_api value  api server to use (default: "https://big-earl.herokuapp.com/") [$BIG_EARL_API]
        --help, -h            show help
        --version, -v         print the version
        ```
        
    