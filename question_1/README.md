## Random Public IP

One of our applications runs on a server with no fixed public IP. Instead, we have a pipeline that runs terraform periodically and updates the firewall rule, allowing access from this application to our primary cloud provider.

Code a terraform module that outputs the outgoing IP address of this server.

**Note:** You can assume that the pipeline runs from the same server as the application.
