variable "ami" {
  default = "foo"
  type    = "string"
}

variable "list" {
  default = []
  type    = "list"
}

resource "aws_instance" "bar" {
  foo = "${var.ami}"
  bar = "${join(",", var.list)}"
}
