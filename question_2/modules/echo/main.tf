variable "input" {
  type        = string
  description = "value"
  validation {
    condition     = substr(var.input, 0, 12) == "I am message"
    error_message = "The input value must start with \"I am message\"."
  }
}

output "input" {
  value = var.input
}
