resource "aws_security_group" "artifactory_instance_secgroup" {
  name = "artifactory_instance_secgroup"
  vpc_id = "${aws_vpc.artifactory_integration.id}"

  ingress {
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = [
      "${aws_instance.internal_apps_bastion.public_ip}/32",
      "${aws_instance.internal_apps_bastion_2.public_ip}/32"
    ]
  }

  ingress {
    from_port = 443
    to_port = 443
    protocol = "tcp"
    cidr_blocks = [
      "${aws_vpc.pardot_ci.cidr_block}",
      "${aws_vpc.internal_apps.cidr_block}"
    ]
  }

  ingress {
    from_port = 80
    to_port = 80
    protocol = "tcp"
    cidr_blocks = [
      "${aws_vpc.pardot_ci.cidr_block}",
      "${aws_vpc.internal_apps.cidr_block}"
    ]
  }

  egress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "artifactory_http_lb" {
  name = "internal_apps_http_lb"
  description = "Allow HTTP/HTTPS from SFDC VPN only"
  vpc_id = "${aws_vpc.internal_apps.id}"

  ingress {
    from_port = 80
    to_port = 80
    protocol = "tcp"
    cidr_blocks = [
      "204.14.236.0/24",    # aloha-east
      "204.14.239.0/24",    # aloha-west
      "62.17.146.140/30",   # aloha-emea
      "62.17.146.144/28",   # aloha-emea
      "62.17.146.160/27",   # aloha-emea
      "173.192.141.222/32", # tools-s1 (prodbot)
      "174.37.191.2/32",    # proxy.dev
      "169.45.0.88/32",     # squid-d4
      "136.147.104.20/30",  # pardot-proxyout1-{1,2,3,4}-dfw
      "136.147.96.20/30"    # pardot-proxyout1-{1,2,3,4}-phx
    ]
  }

  ingress {
    from_port = 443
    to_port = 443
    protocol = "tcp"
    cidr_blocks = [
      "204.14.236.0/24",    # aloha-east
      "204.14.239.0/24",    # aloha-west
      "62.17.146.140/30",   # aloha-emea
      "62.17.146.144/28",   # aloha-emea
      "62.17.146.160/27",   # aloha-emea
      "173.192.141.222/32", # tools-s1 (prodbot)
      "174.37.191.2/32",    # proxy.dev
      "169.45.0.88/32",     # squid-d4
      "136.147.104.20/30",  # pardot-proxyout1-{1,2,3,4}-dfw
      "136.147.96.20/30",   # pardot-proxyout1-{1,2,3,4}-phx
      "50.22.140.200/32"    # tools-s1.dev
    ]
  }

  egress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "artifactory_dc_only_http_lb" {
  name = "internal_apps_dc_only_http_lb"
  description = "Allow HTTP/HTTPS from SFDC datacenters only"
  vpc_id = "${aws_vpc.internal_apps.id}"

  ingress {
    from_port = 80
    to_port = 80
    protocol = "tcp"
    cidr_blocks = [
      "173.192.141.222/32", # tools-s1 (prodbot)
      "174.37.191.2/32",    # proxy.dev
      "169.45.0.88/32",     # squid-d4
      "136.147.104.20/30",  # pardot-proxyout1-{1,2,3,4}-dfw
      "136.147.96.20/30"    # pardot-proxyout1-{1,2,3,4}-phx
    ]
  }

  ingress {
    from_port = 443
    to_port = 443
    protocol = "tcp"
    cidr_blocks = [
      "173.192.141.222/32", # tools-s1 (prodbot)
      "208.43.203.134/32",  # email-d1 (replication check)
      "174.37.191.2/32",    # proxy.dev
      "169.45.0.88/32",     # squid-d4
      "136.147.104.20/30",  # pardot-proxyout1-{1,2,3,4}-dfw
      "136.147.96.20/30"    # pardot-proxyout1-{1,2,3,4}-phx
    ]
  }

  egress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}


resource "aws_instance" "pardot0-artifactory1-1-ue1" {
  ami = "${var.centos_7_hvm_ebs_ami}"
  instance_type = "c4.4xlarge"
  key_name = "internal_apps"
  subnet_id = "${aws_subnet.artifactory_integration_us_east_1a_dmz.id}"
  vpc_security_group_ids = ["${aws_security_group.artifactory_instance_secgroup.id}"]
  security_groups = [
    "${aws_security_group.artifactory_dc_only_http_lb.id}",
    "${aws_security_group.artifactory_http_lb.id}"
  ]
  root_block_device {
    volume_type = "gp2"
    volume_size = "2047"
    delete_on_termination = true
  }
  tags {
    Name = "pardot0-artifactory1-1-ue1"
    terraform = "true"
  }
}

resource "aws_eip" "elasticip_pardot0-artifactory1-1-ue1" {
  vpc = true
  instance = "${aws_instance.pardot0-artifactory1-1-ue1.id}"
}

resource "aws_instance" "pardot0-artifactory1-2-ue1" {
  ami = "${var.centos_7_hvm_ebs_ami}"
  instance_type = "c4.4xlarge"
  key_name = "internal_apps"
  subnet_id = "${aws_subnet.artifactory_integration_us_east_1d_dmz.id}"
  vpc_security_group_ids = ["${aws_security_group.artifactory_instance_secgroup.id}"]
  security_groups = [
    "${aws_security_group.artifactory_dc_only_http_lb.id}",
    "${aws_security_group.artifactory_http_lb.id}"
  ]
  root_block_device {
    volume_type = "gp2"
    volume_size = "2047"
    delete_on_termination = true
  }
  tags {
    Name = "pardot0-artifactory1-2-ue1"
    terraform = "true"
  }
}

resource "aws_eip" "elasticip_pardot0-artifactory1-2-ue1" {
  vpc = true
  instance = "${aws_instance.pardot0-artifactory1-2-ue1.id}"
}

resource "aws_instance" "pardot0-artifactory1-3-ue1" {
  ami = "${var.centos_7_hvm_ebs_ami}"
  instance_type = "c4.4xlarge"
  key_name = "internal_apps"
  subnet_id = "${aws_subnet.artifactory_integration_us_east_1c_dmz.id}"
  vpc_security_group_ids = ["${aws_security_group.artifactory_instance_secgroup.id}"]
  security_groups = [
    "${aws_security_group.artifactory_dc_only_http_lb.id}",
    "${aws_security_group.artifactory_http_lb.id}"
  ]
  root_block_device {
    volume_type = "gp2"
    volume_size = "2047"
    delete_on_termination = true
  }
  tags {
    Name = "pardot0-artifactory1-3-ue1"
    terraform = "true"
  }
}

resource "aws_eip" "elasticip_pardot0-artifactory1-3-ue1" {
  vpc = true
  instance = "${aws_instance.pardot0-artifactory1-3-ue1.id}"
}

resource "aws_elb" "artifactory_ops_elb" {
  name = "artifactory-elb"
  security_groups = [
    "${aws_security_group.artifactory_dc_only_http_lb.id}",
    "${aws_security_group.artifactory_http_lb.id}"
  ]
  subnets = [
    "${aws_subnet.internal_apps_us_east_1a_dmz.id}",
    "${aws_subnet.internal_apps_us_east_1c_dmz.id}",
    "${aws_subnet.internal_apps_us_east_1d_dmz.id}",
    "${aws_subnet.internal_apps_us_east_1e_dmz.id}"
  ]
  cross_zone_load_balancing = true
  connection_draining = true
  connection_draining_timeout = 30
  instances = ["${aws_instance.pardot0-artifactory1-1-ue1.id}","${aws_instance.pardot0-artifactory1-2-ue1.id}"]

  listener {
    lb_port = 443
    lb_protocol = "https"
    instance_port = 80
    instance_protocol = "http"
    ssl_certificate_id = "arn:aws:iam::364709603225:server-certificate/ops.pardot.com"
  }

  listener {
    lb_port = 80
    lb_protocol = "http"
    instance_port = 80
    instance_protocol = "http"
  }

  health_check {
    healthy_threshold = 4
    unhealthy_threshold = 2
    timeout = 3
    target = "HTTP:80/artifactory/api/system/ping"
    interval = 5
  }

  tags {
    Name = "artifactory"
  }
}

resource "aws_iam_user" "artifactory_sysacct" {
  name = "artifactorysyscct"
}

resource "aws_s3_bucket" "artifactory_s3_filestore" {
  bucket = "artifactory_s3_filestore"
  acl = "private"

  tags {
    Name = "artifactory_s3_filestore"
    terraform = "true"
  }

//  policy = <<EOF{
//    "Version": "2012-10-17",
//    "Statement": [
//      {
//        "Sid": "allow artifactory sysacct",
//        "Effect": "Allow",
//        "Principal": {
//        "AWS": "${aws_iam_user.artifactory_sysacct.arn}"
//        },
//        "Action": "s3:*",
//        "Resource": [
//          "arn:aws:s3:::artifactory_s3_filestore",
//          "arn:aws:s3:::artifactory_s3_filestore/*"
//        ]
//      },
//      {
//        "Sid": "DenyIncorrectEncryptionHeader",
//        "Effect": "Deny",
//        "Principal": "*",
//        "Action": "s3:PutObject",
//        "Resource": "arn:aws:s3:::YourBucket/*",
//        "Condition": {
//          "StringNotEquals": {
//            "s3:x-amz-server-side-encryption": "AES256"
//          }
//        }
//      },
//      {
//        "Sid": "DenyUnEncryptedObjectUploads",
//        "Effect": "Deny",
//        "Principal": "*",
//        "Action": "s3:PutObject",
//        "Resource": "arn:aws:s3:::YourBucket/*",
//        "Condition": {
//          "Null": {
//            "s3:x-amz-server-side-encryption": "true"
//          }
//        }
//      }
//    ]
//  }
//EOF
}

resource "aws_vpc" "artifactory_integration" {
cidr_block = "172.28.0.0/24"
  enable_dns_support = true
  enable_dns_hostnames = true
  tags {
    Name = "artifactory_integration"
  }
}

resource "aws_subnet" "artifactory_integration_us_east_1a" {
  vpc_id = "${aws_vpc.artifactory_integration.id}"
  availability_zone = "us-east-1a"
  cidr_block = "172.28.0.0/27"
  map_public_ip_on_launch = false
}

resource "aws_subnet" "artifactory_integration_us_east_1c" {
  vpc_id = "${aws_vpc.artifactory_integration.id}"
  availability_zone = "us-east-1c"
  cidr_block = "172.28.0.32/27"
  map_public_ip_on_launch = false
}

resource "aws_subnet" "artifactory_integration_us_east_1d" {
  vpc_id = "${aws_vpc.artifactory_integration.id}"
  availability_zone = "us-east-1d"
  cidr_block = "172.28.0.64/27"
  map_public_ip_on_launch = false
}

resource "aws_subnet" "artifactory_integration_us_east_1e" {
  vpc_id = "${aws_vpc.artifactory_integration.id}"
  availability_zone = "us-east-1e"
  cidr_block = "172.28.0.96/27"
  map_public_ip_on_launch = false
}

resource "aws_subnet" "artifactory_integration_us_east_1a_dmz" {
  vpc_id = "${aws_vpc.artifactory_integration.id}"
  availability_zone = "us-east-1a"
  cidr_block = "172.28.0.128/27"
  map_public_ip_on_launch = true
}

resource "aws_subnet" "artifactory_integration_us_east_1c_dmz" {
  vpc_id = "${aws_vpc.artifactory_integration.id}"
  availability_zone = "us-east-1c"
  cidr_block = "172.28.0.160/27"
  map_public_ip_on_launch = true
}

resource "aws_subnet" "artifactory_integration_us_east_1d_dmz" {
  vpc_id = "${aws_vpc.artifactory_integration.id}"
  availability_zone = "us-east-1d"
  cidr_block = "172.28.0.192/27"
  map_public_ip_on_launch = true
}

resource "aws_subnet" "artifactory_integration_us_east_1e_dmz" {
  vpc_id = "${aws_vpc.artifactory_integration.id}"
  availability_zone = "us-east-1e"
  cidr_block = "172.28.0.224/27"
  map_public_ip_on_launch = true
}

resource "aws_internet_gateway" "artifactory_integration_internet_gw" {
  vpc_id = "${aws_vpc.artifactory_integration.id}"
}

resource "aws_eip" "artifactory_integration_nat_gw" {
  vpc = true
}

resource "aws_nat_gateway" "artifactory_integration_nat_gw" {
  allocation_id = "${aws_eip.artifactory_integration_nat_gw.id}"
  subnet_id = "${aws_subnet.artifactory_integration_us_east_1a_dmz.id}"
}

resource "aws_route" "artifactory_integration_route" {
  route_table_id = "${aws_vpc.artifactory_integration.main_route_table_id}"
  destination_cidr_block = "0.0.0.0/0"
  nat_gateway_id = "${aws_nat_gateway.artifactory_integration_nat_gw.id}"
}

resource "aws_route_table" "artifactory_integration_route_dmz" {
  vpc_id = "${aws_vpc.artifactory_integration.id}"
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = "${aws_internet_gateway.artifactory_integration_internet_gw.id}"
  }
}

resource "aws_route_table_association" "artifactory_integration_us_east_1a" {
  subnet_id = "${aws_subnet.artifactory_integration_us_east_1a.id}"
  route_table_id = "${aws_vpc.artifactory_integration.main_route_table_id}"
}

resource "aws_route_table_association" "artifactory_integration_us_east_1c" {
  subnet_id = "${aws_subnet.artifactory_integration_us_east_1c.id}"
  route_table_id = "${aws_vpc.artifactory_integration.main_route_table_id}"
}

resource "aws_route_table_association" "artifactory_integration_us_east_1d" {
  subnet_id = "${aws_subnet.artifactory_integration_us_east_1d.id}"
  route_table_id = "${aws_vpc.artifactory_integration.main_route_table_id}"
}

resource "aws_route_table_association" "artifactory_integration_us_east_1e" {
  subnet_id = "${aws_subnet.artifactory_integration_us_east_1e.id}"
  route_table_id = "${aws_vpc.artifactory_integration.main_route_table_id}"
}

resource "aws_route_table_association" "artifactory_integration_us_east_1a_dmz" {
  subnet_id = "${aws_subnet.artifactory_integration_us_east_1a_dmz.id}"
  route_table_id = "${aws_route_table.artifactory_integration_route_dmz.id}"
}

resource "aws_route_table_association" "artifactory_integration_us_east_1c_dmz" {
  subnet_id = "${aws_subnet.artifactory_integration_us_east_1c_dmz.id}"
  route_table_id = "${aws_route_table.artifactory_integration_route_dmz.id}"
}

resource "aws_route_table_association" "artifactory_integration_us_east_1d_dmz" {
  subnet_id = "${aws_subnet.artifactory_integration_us_east_1d_dmz.id}"
  route_table_id = "${aws_route_table.artifactory_integration_route_dmz.id}"
}

resource "aws_route_table_association" "artifactory_integration_us_east_1e_dmz" {
  subnet_id = "${aws_subnet.artifactory_integration_us_east_1e_dmz.id}"
  route_table_id = "${aws_route_table.artifactory_integration_route_dmz.id}"
}

resource "aws_security_group" "artifactory_integration_mysql_ingress" {
  name = "artifactory_integration_mysql_ingress"
  description = "Allow mysql from artifactory instances only"
  vpc_id = "${aws_vpc.artifactory_integration.id}"

  ingress {
    from_port = 3306
    to_port = 3306
    protocol = "tcp"
    cidr_blocks = [
      "${aws_subnet.artifactory_integration_us_east_1a.cidr_block}",
      "${aws_subnet.artifactory_integration_us_east_1a_dmz.cidr_block}",
      "${aws_subnet.artifactory_integration_us_east_1c.cidr_block}",
      "${aws_subnet.artifactory_integration_us_east_1c_dmz.cidr_block}",
      "${aws_subnet.artifactory_integration_us_east_1d.cidr_block}",
      "${aws_subnet.artifactory_integration_us_east_1d_dmz.cidr_block}"
    ]
  }

  egress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_db_subnet_group" "artifactory_integration" {
  name = "artifactory_integration"
  description = "Pardot CI DB Subnet"
  subnet_ids = [
    "${aws_subnet.artifactory_integration_us_east_1a.id}",
    "${aws_subnet.artifactory_integration_us_east_1c.id}",
    "${aws_subnet.artifactory_integration_us_east_1d.id}",
    "${aws_subnet.artifactory_integration_us_east_1e.id}"
  ]
}
resource "aws_vpc_peering_connection" "internal_apps_and_artifactory_integration_vpc_peering" {
  peer_owner_id = "364709603225" # pardotops
  peer_vpc_id = "${aws_vpc.internal_apps.id}"
  vpc_id = "${aws_vpc.artifactory_integration.id}"
}

resource "aws_vpc_peering_connection" "pardot_ci_and_artifactory_integration_vpc_peering" {
  peer_owner_id = "364709603225" # pardotops
  peer_vpc_id = "${aws_vpc.pardot_ci.id}"
  vpc_id = "${aws_vpc.artifactory_integration.id}"
}