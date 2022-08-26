terraform {
  required_version = "~> 1.2.7"
  required_providers {
    # https://github.com/kreuzwerker/terraform-provider-docker
    docker = {
      source  = "kreuzwerker/docker"
      version = "2.20.2"
    }
    local = {
      source  = "hashicorp/local"
      version = "2.2.3"
    }
    postgresql = {
      source  = "cyrilgdn/postgresql"
      version = "1.16.0"
    }
  }
}

provider "local" {}
provider "docker" {}
provider "postgresql" {
  host             = "localhost"
  port             = 5432
  database         = "hasteheart"
  username         = "user"
  password         = "password"
  sslmode          = "disable"
  superuser        = true
  expected_version = 14.7
  connect_timeout  = 15
}

#
# resource 1
#
resource "docker_container" "echo" {
  name  = "echo"
  image = "mendhak/http-https-echo:24"

  ports {
    external = 8080
    internal = 8080
  }
}
output "image_tag" {
  value = docker_container.echo.id
}

#
# resource 2
#

resource "local_file" "hello_world" {
  content  = "foo!"
  filename = "${path.module}/helloWorld"
}

#
# resource 3
#

resource "docker_container" "database" {
  name  = "postgres"
  image = "postgres:14.5"
  ports {
    external = 5432
    internal = 5432
  }
  env = [
    "POSTGRES_DB=hasteheart",
    "POSTGRES_PASSWORD=password",
    "POSTGRES_USER=user"
  ]
}

resource "null_resource" "wait_for_postgres" {
  provisioner "local-exec" {
    command = "./check.sh"
  }
  triggers = {
    always_run = "${timestamp()}"
  }
}

#
# resource 4
#

resource "postgresql_database" "test" {
  depends_on        = [
    null_resource.wait_for_postgres,
    docker_container.database
  ]
  name              = "test"
  template          = "template0"
  lc_collate        = "C"
  connection_limit  = -1
  allow_connections = true
}
