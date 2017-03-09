resource "aws_security_group" "pardot0_ue1_ldap_server" {
  name        = "internal_apps_ldap_server"
  description = "Allow LDAP and LDAPS from SFDC datacenters and internal apps"
  vpc_id      = "${aws_vpc.pardot0_ue1.id}"

  # We run LDAP over port 443 to allow SFDC datacenters to connect to us, since
  # only 80, 443, and 25 are allowed outbound.
  # LDAPS is run on port 443.
  ingress {
    from_port = 80
    to_port   = 80
    protocol  = "tcp"

    cidr_blocks = [
      "${var.sfdc_proxyout_cidr_blocks}",
    ]
  }

  ingress {
    from_port = 443
    to_port   = 443
    protocol  = "tcp"

    cidr_blocks = [
      "${var.sfdc_proxyout_cidr_blocks}",
    ]
  }

  ingress {
    from_port = 389
    to_port   = 389
    protocol  = "tcp"

    cidr_blocks = [
      "${aws_vpc.pardot0_ue1.cidr_block}",
      "${aws_eip.dc_access_00.public_ip}/32",
      "${aws_vpc.internal_tools_integration.cidr_block}",
      "${aws_eip.internal_tools_integration_nat_gw.public_ip}/32",
      "${aws_eip.appdev_ldap_host_eip.public_ip}/32",
      "${aws_eip.github_enterprise_server_1.public_ip}/32",
      "${aws_eip.github_enterprise_server_2.public_ip}/32",
      "52.4.132.69/32",                                            # 1.git.dev.pardot.com
      "52.3.83.197/32",                                            # 2.git.dev.pardot.com
    ]
  }

  ingress {
    from_port = 636
    to_port   = 636
    protocol  = "tcp"

    cidr_blocks = [
      "${aws_vpc.pardot0_ue1.cidr_block}",
      "${aws_eip.dc_access_00.public_ip}/32",
      "${aws_vpc.internal_tools_integration.cidr_block}",
      "${aws_eip.internal_tools_integration_nat_gw.public_ip}/32",
      "${aws_eip.appdev_ldap_host_eip.public_ip}/32",
      "${aws_eip.github_enterprise_server_1.public_ip}/32",
      "${aws_eip.github_enterprise_server_2.public_ip}/32",
      "52.4.132.69/32",                                            # 1.git.dev.pardot.com
      "52.3.83.197/32",                                            # 2.git.dev.pardot.com
    ]
  }

  # SSH from bastion
  ingress {
    from_port = 22
    to_port   = 22
    protocol  = "tcp"

    security_groups = [
      "${aws_security_group.pardot0_ue1_bastion.id}",
    ]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_iam_role" "pardot0_ue1_ldap_master" {
  name               = "internal_apps_ldap_master"
  assume_role_policy = "${file("ec2_instance_trust_relationship.json")}"
}

resource "aws_iam_instance_profile" "pardot0_ue1_ldap_master" {
  name  = "internal_apps_ldap_master"
  roles = ["${aws_iam_role.pardot0_ue1_ldap_master.id}"]
}

resource "aws_iam_role_policy" "pardot0_ue1_ldap_master_policy" {
  name = "internal_apps_ldap_master_policy"
  role = "${aws_iam_role.pardot0_ue1_ldap_master.id}"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ses:SendEmail"
      ],
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_instance" "pardot0_ue1_ldap_master" {
  ami                  = "${var.centos_6_hvm_ebs_ami}"
  instance_type        = "t2.medium"
  iam_instance_profile = "${aws_iam_instance_profile.pardot0_ue1_ldap_master.id}"
  key_name             = "internal_apps"
  private_ip           = "172.30.132.212"
  subnet_id            = "${aws_subnet.pardot0_ue1_1a_dmz.id}"

  vpc_security_group_ids = [
    "${aws_security_group.pardot0_ue1_ldap_server.id}",
  ]

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "40"
    delete_on_termination = false
  }

  tags {
    Name = "pardot0-auth1-1-ue1"
  }
}

resource "aws_instance" "pardot0_ue1_ldap_replica" {
  ami           = "${var.centos_6_hvm_ebs_ami}"
  instance_type = "t2.medium"
  key_name      = "internal_apps"
  private_ip    = "172.30.213.2"
  subnet_id     = "${aws_subnet.pardot0_ue1_1d_dmz.id}"

  vpc_security_group_ids = [
    "${aws_security_group.pardot0_ue1_ldap_server.id}",
  ]

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "40"
    delete_on_termination = false
  }

  tags {
    Name = "pardot0-auth1-2-ue1"
  }
}

resource "aws_eip" "pardot0_ue1_ldap_master" {
  vpc      = true
  instance = "${aws_instance.pardot0_ue1_ldap_master.id}"
}

resource "aws_eip" "pardot0_ue1_ldap_replica" {
  vpc      = true
  instance = "${aws_instance.pardot0_ue1_ldap_replica.id}"
}

// THE FOLLOWING FOUR RECORDS MUST STAY SYNCHRONIZED BETWEEN PRIVATE AND PUBLIC VERSIONS! SEEK BREAD-TEAM FOR ASSISTANCE

resource "aws_route53_record" "pardot0_ue1_auth1-1_Arecord" {
  zone_id = "${aws_route53_zone.pardot0_ue1_aws_pardot_com_hosted_zone.zone_id}"
  name    = "pardot0-auth1-1-ue1.aws.pardot.com"
  records = ["${aws_eip.pardot0_ue1_ldap_master.public_ip}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_route53_record" "pardot0_ue1_auth1-2_Arecord" {
  zone_id = "${aws_route53_zone.pardot0_ue1_aws_pardot_com_hosted_zone.zone_id}"
  name    = "pardot0-auth1-2-ue1.aws.pardot.com"
  records = ["${aws_eip.pardot0_ue1_ldap_replica.public_ip}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_route53_record" "pardot0_ue1_auth1-1_Arecord_PUBLIC" {
  zone_id = "${aws_route53_zone.aws_pardot_com_restricted_use_public_zone.zone_id}"
  name    = "pardot0-auth1-1-ue1.aws.pardot.com"
  records = ["${aws_eip.pardot0_ue1_ldap_master.public_ip}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_route53_record" "pardot0_ue1_auth1-2_Arecord_PUBLIC" {
  zone_id = "${aws_route53_zone.aws_pardot_com_restricted_use_public_zone.zone_id}"
  name    = "pardot0-auth1-2-ue1.aws.pardot.com"
  records = ["${aws_eip.pardot0_ue1_ldap_replica.public_ip}"]
  type    = "A"
  ttl     = "900"
}