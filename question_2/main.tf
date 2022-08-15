module "message_one" {
  source = "./modules/echo"
  input  = "I am message one"
}

module "message_two" {
  depends_on = [module.message_one]
  source     = "./modules/echo"
  input      = "I am message two"
}

output "message_one_input" {
  value = module.message_one.input
}

output "message_two_input" {
  value = module.message_two.input
}
