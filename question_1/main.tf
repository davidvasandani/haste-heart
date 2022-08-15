data "http" "myip" {
  url = "http://ipv4.icanhazip.com"
}

output "myip" {
  value = chomp(data.http.myip.response_body)
}
