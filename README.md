# Haste.Heart Code Challenge

**NOTE:** This project uses [VSCode Remote Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) for local development.

The initial `docker build` takes 105.5s. Once launched run the command `task` to start `terratest` for all three parts of the code challenge; which takes 1m35s.

## Question 1 - Random Public IP

> One of our applications runs on a server with no fixed public IP. Instead, we have a pipeline that runs terraform periodically and updates the firewall rule, allowing access from this application to our primary cloud provider.
> 
> Code a terraform module that outputs the outgoing IP address of this server.
>
> **Note:** You can assume that the pipeline runs from the same server as the application.

http://ipv4.icanhazip.com is a simple site run by CloudFlare that returns your external IP address and nothing else. I used the [http data source](https://registry.terraform.io/providers/hashicorp/http/latest/docs/data-sources/http) to make an HTTP GET request to the given URL and exports information about the response. Nothing else was needed to complete this task. I didn't move the code into a module because it would have just been a variable, the data source, and the output.

## Question 2 - Modules dependency

> Create a terraform module that echos a string. Your main file should call the module twice with the following input, respectively:
> 
> - "I am message one"
> - "I am message two"
> 
> Execute the modules in order and print the messages with double quotes.

## Question 3 - Testing your code

Write any terraform code with a minimum of three resources including tests.

1. `resource "docker_container" "echo"`

2. `resource "local_file" "hello_world"`

3. `resource "docker_container" "database"`

4. `resource "postgresql_database" "test"`
sd