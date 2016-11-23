variable "github_enterprise_ami_us_east_1" {
  default = "ami-6f92bf78"
}

variable "github_enterprise_instance_type" {
  default = "c4.4xlarge"
}

variable "github_enterprise_xvdf_size" {
  default = "250"
}

resource "aws_security_group" "github_enterprise_server_admin_management" {
  name   = "github_enterprise_server_admin_management"
  vpc_id = "${aws_vpc.internal_tools_integration.id}"

  # Administrative SSH port
  ingress {
    from_port = 122
    to_port   = 122
    protocol  = "tcp"

    cidr_blocks = [
      "${var.aloha_vpn_cidr_blocks}",
      "52.4.132.69/32",               # 1.git.dev.pardot.com
    ]

    security_groups = [
      "${aws_security_group.github_enterprise_server_backups.id}",
    ]

    self = true
  }

  # Administrative web interface
  ingress {
    from_port = 8443
    to_port   = 8443
    protocol  = "tcp"

    cidr_blocks = [
      "${var.aloha_vpn_cidr_blocks}",
    ]
  }

  # VPN for secure replication from primary
  ingress {
    from_port = 1194
    to_port   = 1194
    protocol  = "udp"

    cidr_blocks = [
      "52.4.132.69/32", # 1.git.dev.pardot.com
    ]

    self = true
  }
}

resource "aws_security_group" "github_enterprise_server_ssh" {
  name   = "github_enterprise_server_ssh"
  vpc_id = "${aws_vpc.internal_tools_integration.id}"

  ingress {
    from_port = 22
    to_port   = 22
    protocol  = "tcp"

    cidr_blocks = [
      "${var.aloha_vpn_cidr_blocks}",
      "${var.sfdc_proxyout_cidr_blocks}",
      "${aws_vpc.internal_tools_integration.cidr_block}",
      "${aws_vpc.appdev.cidr_block}",
      "${aws_nat_gateway.appdev_nat_gw.public_ip}/32",
      "${aws_vpc.internal_apps.cidr_block}",
      "${aws_nat_gateway.internal_apps_nat_gw.public_ip}/32",
      "${var.pardot_ci_vpc_cidr}",
      "${var.pardot_ci_nat_gw_public_ip}/32",
      "${var.sfdc_pardot_tools_production_heroku_space_cidr_blocks}",
    ]
  }
}

resource "aws_security_group" "github_enterprise_server_http" {
  name   = "github_enterprise_server_http"
  vpc_id = "${aws_vpc.internal_tools_integration.id}"

  ingress {
    from_port = 80
    to_port   = 80
    protocol  = "tcp"

    cidr_blocks = [
      "${var.aloha_vpn_cidr_blocks}",
      "${var.sfdc_proxyout_cidr_blocks}",
      "${aws_vpc.internal_tools_integration.cidr_block}",
      "${aws_vpc.appdev.cidr_block}",
      "${aws_nat_gateway.appdev_nat_gw.public_ip}/32",
      "${aws_vpc.internal_apps.cidr_block}",
      "${aws_nat_gateway.internal_apps_nat_gw.public_ip}/32",
      "${var.pardot_ci_vpc_cidr}",
      "${var.pardot_ci_nat_gw_public_ip}/32",
      "${var.sfdc_pardot_tools_production_heroku_space_cidr_blocks}",
    ]
  }

  ingress {
    from_port = 443
    to_port   = 443
    protocol  = "tcp"

    cidr_blocks = [
      "${var.aloha_vpn_cidr_blocks}",
      "${var.sfdc_proxyout_cidr_blocks}",
      "${aws_vpc.internal_tools_integration.cidr_block}",
      "${aws_vpc.appdev.cidr_block}",
      "${aws_nat_gateway.appdev_nat_gw.public_ip}/32",
      "${aws_vpc.internal_apps.cidr_block}",
      "${aws_nat_gateway.internal_apps_nat_gw.public_ip}/32",
      "${var.pardot_ci_vpc_cidr}",
      "${var.pardot_ci_nat_gw_public_ip}/32",
      "${var.sfdc_pardot_tools_production_heroku_space_cidr_blocks}",
    ]
  }
}

resource "aws_eip" "github_enterprise_server_1" {
  vpc      = true
  instance = "${aws_instance.github_enterprise_server_1.id}"
}

resource "aws_eip" "github_enterprise_server_2" {
  vpc      = true
  instance = "${aws_instance.github_enterprise_server_2.id}"
}

resource "aws_instance" "github_enterprise_server_1" {
  ami                     = "${var.github_enterprise_ami_us_east_1}"
  instance_type           = "${var.github_enterprise_instance_type}"
  key_name                = "github_enterprise"
  subnet_id               = "${aws_subnet.internal_tools_integration_us_east_1a_dmz.id}"
  ebs_optimized           = true
  disable_api_termination = true

  vpc_security_group_ids = [
    "${aws_security_group.internal_tools_integration_default.id}",
    "${aws_security_group.github_enterprise_server_admin_management.id}",
    "${aws_security_group.github_enterprise_server_ssh.id}",
    "${aws_security_group.github_enterprise_server_http.id}",
  ]

  root_block_device {
    volume_type           = "gp2"
    delete_on_termination = false
  }

  ebs_block_device {
    device_name           = "/dev/xvdf"
    volume_type           = "gp2"
    volume_size           = "${var.github_enterprise_xvdf_size}"
    delete_on_termination = false
    encrypted             = true
  }

  tags {
    terraform = "true"
    Name      = "pardot0-github1-1-ue1"
  }
}

resource "aws_route53_record" "github_enterprise_server_1_Arecord" {
  zone_id = "${aws_route53_zone.internal_apps_aws_pardot_com_hosted_zone.zone_id}"
  name    = "pardot0-github1-1-ue1.${aws_route53_zone.internal_apps_aws_pardot_com_hosted_zone.name}"
  records = ["${aws_instance.github_enterprise_server_1.private_ip}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "github_enterprise_server_2" {
  ami                     = "${var.github_enterprise_ami_us_east_1}"
  instance_type           = "${var.github_enterprise_instance_type}"
  key_name                = "github_enterprise"
  subnet_id               = "${aws_subnet.internal_tools_integration_us_east_1d_dmz.id}"
  ebs_optimized           = true
  disable_api_termination = true

  vpc_security_group_ids = [
    "${aws_security_group.internal_tools_integration_default.id}",
    "${aws_security_group.github_enterprise_server_admin_management.id}",
    "${aws_security_group.github_enterprise_server_ssh.id}",
    "${aws_security_group.github_enterprise_server_http.id}",
  ]

  root_block_device {
    volume_type           = "gp2"
    delete_on_termination = false
  }

  ebs_block_device {
    device_name           = "/dev/xvdf"
    volume_type           = "gp2"
    volume_size           = "${var.github_enterprise_xvdf_size}"
    delete_on_termination = false
    encrypted             = true
  }

  tags {
    terraform = "true"
    Name      = "pardot0-github1-2-ue1"
  }
}

resource "aws_route53_record" "github_enterprise_server_2_Arecord" {
  zone_id = "${aws_route53_zone.internal_apps_aws_pardot_com_hosted_zone.zone_id}"
  name    = "pardot0-github1-2-ue1.${aws_route53_zone.internal_apps_aws_pardot_com_hosted_zone.name}"
  records = ["${aws_instance.github_enterprise_server_2.private_ip}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_security_group" "github_enterprise_server_backups" {
  name   = "github_enterprise_server_backups"
  vpc_id = "${aws_vpc.internal_apps.id}"

  # SSH from bastion
  ingress {
    from_port = 22
    to_port   = 22
    protocol  = "tcp"

    security_groups = [
      "${aws_security_group.internal_apps_bastion.id}",
    ]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "github_enterprise_server_backups" {
  ami                     = "${var.centos_6_hvm_ebs_ami}"
  instance_type           = "t2.medium"
  key_name                = "internal_apps"
  subnet_id               = "${aws_subnet.internal_apps_us_east_1c.id}"
  disable_api_termination = true

  vpc_security_group_ids = [
    "${aws_security_group.github_enterprise_server_backups.id}",
  ]

  root_block_device {
    volume_type           = "gp2"
    delete_on_termination = true
  }

  ebs_block_device {
    device_name           = "/dev/xvdf"
    volume_type           = "standard"
    volume_size           = "1024"
    delete_on_termination = false
    encrypted             = true
  }

  tags {
    terraform = "true"
    Name      = "pardot0-githubbackup1-1-ue1"
  }
}

resource "aws_route53_record" "github_enterprise_server_backups_Arecord" {
  zone_id = "${aws_route53_zone.internal_apps_aws_pardot_com_hosted_zone.zone_id}"
  name    = "pardot0-githubbackup1-1-ue1.${aws_route53_zone.internal_apps_aws_pardot_com_hosted_zone.name}"
  records = ["${aws_instance.github_enterprise_server_backups.private_ip}"]
  type    = "A"
  ttl     = "900"
}
