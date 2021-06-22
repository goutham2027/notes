provider "aws" {
  region     = "us-east-1"
  access_key = "AKIAJJHJHXZG7O4P6H6Q"
  secret_key = "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"
}

# resource "aws_instance" "my-first-aws-instance" {
#   ami           = "ami-06b263d6ceff0b3dd"
#   instance_type = "t2.micro"

#   tags = {
#     Name = "ubuntu"
#   }
# }


# resource "aws_subnet" "subnet" {
#   vpc_id     = aws_vpc.first-vpc.id
#   cidr_block = "10.0.1.0/24"

#   tags = {
#     Name = "prod-subnet"
#   }
# }

# 1. Create VPC
 resource "aws_vpc" "prod-vpc" {
   cidr_block = "10.0.0.0/16"
   tags = {
     Name = "production"
   }
 }
# 2. Create Internet Gateway
resource "aws_internet_gateway", "gw" {
   vpc_id     = aws_vpc.prod-vpc.id
}
# 3. Create Custom Route Table
# 4. Create a Subnet
# 5. Associate Subnet with Route Table
# 6. Create Security Group to allow port 22, 80, 443
# 7. Create a network interface with an ip in the subnet that was created in step-4
# 8. Assign an elastic IP to the network interface created in step-7
# 9. Create Ubuntu server and install/enable apache2
