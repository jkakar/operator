# App.dev environment

# managed by pd-bread@salesforce.com

variable "environment_appdev" {
  type = "map"

  default = {
    env_name                  = "appdev"
    pardot_env_id             = "pardot2"
    dc_id                     = "ue1"
    lightweight_instance_type = "c4.large"
    app_instance_type         = "m4.large"
    job_instance_type         = "m4.large"
    metrics_instance_type     = "m4.xlarge"
    db_instance_type          = "m4.2xlarge"
    kafka_instance_type       = "m4.2xlarge"
    storm_instance_type       = "m4.4xlarge"
    ceph_instance_type        = "m4.xlarge"
    db_volume_device_name     = "/dev/xvdf"
    num_globaldb1_hosts       = 2
    num_dbshard1_hosts        = 4
    num_whoisdb1_hosts        = 2
    num_app1_hosts            = 2
    num_thumbs1_hosts         = 1
    num_redisjob1_hosts       = 2
    num_jobmanager1_hosts     = 2
    num_push1_hosts           = 2
    num_provisioning1_hosts   = 1
    num_rabbit1_hosts         = 3
    num_rabbit2_hosts         = 3
    num_redisrules1_hosts     = 2
    num_autojob1_hosts        = 4
    num_storm1_hosts          = 3
    num_kafka1_hosts          = 3
    num_zkkafka1_hosts        = 1
    num_pubsub1_hosts         = 2
    num_zkstorm1_hosts        = 1
    num_nimbus1_hosts         = 1
    num_appcache1_hosts       = 2
    num_discovery1_hosts      = 3
    num_proxyout1_hosts       = 1
    num_toolsproxy1_hosts     = 1
    num_vault1_hosts          = 3
    num_consul1_hosts         = 2
    num_indexer1_hosts        = 1
    num_cephrgw1_hosts        = 1
    num_cephosd1_hosts        = 3
    num_cephrgw2_hosts        = 1
    num_cephosd2_hosts        = 3
    num_ns1_hosts             = 1
    num_metrics1_hosts        = 1
    num_lba1_hosts            = 1
  }
}

variable "appdev_globaldb1_ips" {
  type = "map"

  default = {
    "0" = "172.26.80.125"
    "1" = "172.26.81.47"
  }
}

variable "appdev_dbshard1_ips" {
  type = "map"

  default = {
    "0" = "172.26.92.23"
    "1" = "172.26.93.40"
    "2" = "172.26.75.74"
    "3" = "172.26.69.79"
  }
}

variable "appdev_whoisdb1_ips" {
  type = "map"

  default = {
    "0" = "172.26.66.54"
    "1" = "172.26.66.55"
  }
}

/*
//// TEMPLATES
////
//
//// EC2 DB INSTANCE: replace "lbl" w/ "servicename"
//resource "aws_instance" "appdev_lbl1" {
//  key_name = "internal_apps"
//  count = "${var.environment_appdev["num_lbl1_hosts"]}"
//  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
//  instance_type = "${var.environment_appdev["db_instance_type}"
//  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
//  ebs_optimized = "true"
//  root_block_device {
//    volume_type = "gp2"
//    volume_size = "50"
//    delete_on_termination = true
//  }
//  ebs_block_device {
//    device_name = "${var.environment_appdev["db_volume_device_name"]}"
//    volume_type = "gp2"
//    volume_size = "512"
//    delete_on_termination = true
//  }
//  vpc_security_group_ids = [
//    "${aws_security_group.appdev_vpc_default.id}",
//    "${aws_security_group.appdev_dbhost.id}",
//    "${aws_security_group.appdev.id}"
//  ]
//  tags {
//    Name = "${var.environment_appdev["pardot_env_id"]}-lbl1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
//    terraform = "true"
//  }
//}
//
//// EC2 APP INSTANCE: replace "lbl" w/ "servicename"
//resource "aws_instance" "appdev_lbl1" {
//  key_name = "internal_apps"
//  count = "${var.environment_appdev["num_lbl1_hosts"]}"
//  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
//  instance_type = "${var.environment_appdev["app_instance_type"]}"
//  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
//  root_block_device {
//    volume_type = "gp2"
//    volume_size = "50"
//    delete_on_termination = true
//  }
//  vpc_security_group_ids = [
//    "${aws_security_group.appdev_vpc_default.id}",
//    "${aws_security_group.appdev_dbhost.id}",
//    "${aws_security_group.appdev.id}"
//  ]
//  tags {
//    Name = "${var.environment_appdev["pardot_env_id"]}-lbl1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
//    terraform = "true"
//  }
//}
//
//// EC2 JOB INSTANCE: replace "lbl" w/ "servicename"
//resource "aws_instance" "appdev_lbl1" {
//  key_name = "internal_apps"
//  count = "${var.environment_appdev["num_lbl1_hosts"]}"
//  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
//  instance_type = "${var.environment_appdev["job_instance_type"]}"
//  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
//  ebs_optimized = "true"
//  root_block_device {
//    volume_type = "gp2"
//    volume_size = "50"
//    delete_on_termination = true
//  }
//  vpc_security_group_ids = [
//    "${aws_security_group.appdev_vpc_default.id}",
//    "${aws_security_group.appdev_dbhost.id}",
//    "${aws_security_group.appdev.id}"
//  ]
//  tags {
//    Name = "${var.environment_appdev["pardot_env_id"]}-lbl1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
//    terraform = "true"
//  }
//}
//
//// EC2 ELASTIC IP: replace "lbl" w/ "servicename"
//resource "aws_eip" "appdev_lbl1" {
//  count = "${var.environment_appdev["num_lbl1_hosts"]}"
//  instance = "${element(aws_instance.appdev_lbl1.*.id, count.index)}"
//  vpc = true
//}
//
//// ROUTE53 A-RECORD: replace "lbl" w/ "servicename"
//resource "aws_route53_record" "appdev_lbl1_arecord" {
//  count = "${var.environment_appdev["num_lbl1_hosts"]}"
//  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
//  name = "${var.environment_appdev["pardot_env_id"]}-lbl1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
//  records = ["${var.environment_appdev["pardot_env_id"]}-lbl1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"]
//  type = "A"
//  ttl = 900
//}
*/

resource "aws_security_group" "appdev" {
  name = "appdev"

  # This description is not accurate, but we can't change it. Here's what it should be:
  # "Allow all traffic from appdev vpc"
  description = "Allow HTTP/HTTPS traffic from appdev vpc"

  vpc_id = "${aws_vpc.appdev.id}"

  ingress {
    from_port = 0
    to_port   = 0
    protocol  = "-1"
    self      = true
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "appdev_consulhost" {
  name        = "appdev_consulhost"
  description = "Allows communication among Vault and Consul hosts"
  vpc_id      = "${aws_vpc.appdev.id}"

  ingress {
    from_port = 0
    to_port   = 0
    protocol  = "-1"
    self      = true
  }

  # allow app.dev hosts access to vault port
  ingress {
    from_port = 8200
    to_port   = 8200
    protocol  = "tcp"

    cidr_blocks = [
      "${aws_vpc.appdev.cidr_block}",
    ]
  }

  # allow app.dev hosts access to consul server port
  ingress {
    from_port = 8300
    to_port   = 8300
    protocol  = "tcp"

    cidr_blocks = [
      "${aws_vpc.appdev.cidr_block}",
    ]
  }

  # allow toolsproxy to access consul UI
  ingress {
    from_port = 8500
    to_port   = 8500
    protocol  = "tcp"

    security_groups = [
      "${aws_security_group.appdev_toolsproxy.id}",
    ]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "appdev_dbhost" {
  name        = "appdev_dbhost"
  description = "Allow MYSQL traffic from appdev apphosts"
  vpc_id      = "${aws_vpc.appdev.id}"

  ingress {
    from_port = 3306
    to_port   = 3306
    protocol  = "tcp"

    security_groups = [
      "${aws_security_group.appdev.id}",
    ]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "appdev_proxyout_host" {
  name        = "appdev_proxyout_host"
  description = "Allow Squid proxy traffic from appdev apphosts"
  vpc_id      = "${aws_vpc.appdev.id}"

  ingress {
    from_port = 3128
    to_port   = 3128
    protocol  = "tcp"

    cidr_blocks = [
      "${aws_vpc.appdev.cidr_block}",
    ]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_route53_record" "app_dev_pardot_com_Arecord" {
  zone_id = "${aws_route53_zone.dev_pardot_com.zone_id}"
  name    = "app.${aws_route53_zone.dev_pardot_com.name}"
  records = ["${aws_eip.appdev_lba1_eip.public_ip}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_globaldb1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_globaldb1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["db_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"
  ebs_optimized = "true"
  private_ip    = "${lookup(var.appdev_globaldb1_ips,count.index)}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  ebs_block_device {
    device_name           = "${var.environment_appdev["db_volume_device_name"]}"
    volume_type           = "gp2"
    volume_size           = "512"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev.id}",
    "${aws_security_group.appdev_dbhost.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-globaldb1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_globaldb1_arecord" {
  count   = "${var.environment_appdev["num_globaldb1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-globaldb1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_globaldb1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_dbshard1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_dbshard1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["db_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"
  ebs_optimized = "true"
  private_ip    = "${lookup(var.appdev_dbshard1_ips,count.index)}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  ebs_block_device {
    device_name           = "${var.environment_appdev["db_volume_device_name"]}"
    volume_type           = "gp2"
    volume_size           = "512"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev.id}",
    "${aws_security_group.appdev_dbhost.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-dbshard1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_dbshard1_arecord" {
  count   = "${var.environment_appdev["num_dbshard1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-dbshard1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_dbshard1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_app1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_app1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev.id}",
    "${aws_security_group.appdev_app1host.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-app1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_security_group" "appdev_app1host" {
  name        = "appdev_app1host"
  description = "Allow HTTP/HTTPS traffic from appdev vpc"
  vpc_id      = "${aws_vpc.appdev.id}"

  # allow health check from ELBs
  ingress {
    from_port = 80
    to_port   = 80
    protocol  = "tcp"

    security_groups = [
      "${aws_security_group.appdev_app_lb.id}",
    ]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_route53_record" "appdev_app1_arecord" {
  count   = "${var.environment_appdev["num_app1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-app1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_app1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_thumbs1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_thumbs1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-thumbs1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_thumbs1_arecord" {
  count   = "${var.environment_appdev["num_thumbs1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-thumbs1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_thumbs1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_redisjob1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_redisjob1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["job_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-redisjob1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_redisjob1_arecord" {
  count   = "${var.environment_appdev["num_redisjob1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-redisjob1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_redisjob1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_jobmanager1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_jobmanager1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-jobmanager1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_jobmanager1_arecord" {
  count   = "${var.environment_appdev["num_jobmanager1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-jobmanager1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_jobmanager1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_push1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_push1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-push1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_push1_arecord" {
  count   = "${var.environment_appdev["num_push1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-push1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_push1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_provisioning1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_provisioning1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev.id}",
    "${aws_security_group.appdev_provisioning1host.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-provisioning1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_security_group" "appdev_provisioning1host" {
  name        = "appdev_provisioning1host"
  description = "Allow HTTPS traffic from appdev vpc"
  vpc_id      = "${aws_vpc.appdev.id}"

  # allow health check from ELBs
  ingress {
    from_port = 8095
    to_port   = 8095
    protocol  = "tcp"

    security_groups = [
      "${aws_security_group.appdev_sfdc_provisioning_https.id}",
    ]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_elb" "appdev_provisioning_elb" {
  name = "${var.environment_appdev["env_name"]}-provisioning-elb"

  security_groups = [
    "${aws_security_group.appdev_sfdc_provisioning_https.id}",
  ]

  subnets = [
    "${aws_subnet.appdev_us_east_1a_dmz.id}",
    "${aws_subnet.appdev_us_east_1c_dmz.id}",
    "${aws_subnet.appdev_us_east_1d_dmz.id}",
    "${aws_subnet.appdev_us_east_1e_dmz.id}",
  ]

  cross_zone_load_balancing   = true
  connection_draining         = true
  connection_draining_timeout = 30
  instances                   = ["${aws_instance.appdev_provisioning1.*.id}"]

  listener {
    lb_port            = 443
    lb_protocol        = "https"
    instance_port      = 8095
    instance_protocol  = "http"
    ssl_certificate_id = "arn:aws:iam::${var.pardotops_account_number}:server-certificate/dev.pardot.com-2016-with-intermediate"
  }

  listener {
    lb_port           = 80
    lb_protocol       = "http"
    instance_port     = 8095
    instance_protocol = "http"
  }

  health_check {
    healthy_threshold   = 4
    unhealthy_threshold = 2
    timeout             = 3
    target              = "HTTP:8095/services/data/v30.0/test"
    interval            = 5
  }

  tags {
    Name = "appdev_provisioning_elb"
  }
}

resource "aws_route53_record" "provisioning_dev_pardot_com_CNAMErecord" {
  zone_id = "${aws_route53_zone.dev_pardot_com.zone_id}"
  name    = "provisioning.${aws_route53_zone.dev_pardot_com.name}"
  records = ["${aws_elb.appdev_provisioning_elb.dns_name}"]
  type    = "CNAME"
  ttl     = "900"
}

resource "aws_route53_record" "appdev_provisioning1_arecord" {
  count   = "${var.environment_appdev["num_provisioning1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-provisioning1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_provisioning1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_rabbit1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_rabbit1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_rabbithost.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-rabbit1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_security_group" "appdev_rabbithost" {
  name        = "appdev_rabbithost"
  description = "Allow access through the toolsproxy and from apphosts"
  vpc_id      = "${aws_vpc.appdev.id}"

  ingress {
    from_port = 15672
    to_port   = 15672
    protocol  = "tcp"
    self      = true

    security_groups = [
      "${aws_security_group.appdev_toolsproxy.id}",
      "${aws_security_group.appdev.id}",
    ]
  }

  ingress {
    from_port = 5671
    to_port   = 5672
    protocol  = "tcp"
    self      = true

    security_groups = [
      "${aws_security_group.appdev.id}",
    ]
  }

  ingress {
    from_port = 4369
    to_port   = 4369
    protocol  = "tcp"
    self      = true
  }

  ingress {
    from_port = 25672
    to_port   = 25672
    protocol  = "tcp"
    self      = true
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_route53_record" "appdev_rabbit_webapp_arecord" {
  count   = "2"
  zone_id = "${aws_route53_zone.dev_pardot_com.zone_id}"
  name    = "rabbit${count.index + 1}-app.${aws_route53_zone.dev_pardot_com.name}"
  records = ["${aws_eip.appdev_toolsproxy1.public_ip}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_route53_record" "appdev_rabbit1_arecord" {
  count   = "${var.environment_appdev["num_rabbit1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-rabbit1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_rabbit1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_rabbit2" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_rabbit2_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_rabbithost.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-rabbit2-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_rabbit2_arecord" {
  count   = "${var.environment_appdev["num_rabbit2_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-rabbit2-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_rabbit2.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_redisrules1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_redisrules1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-redisrules1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_redisrules1_arecord" {
  count   = "${var.environment_appdev["num_redisrules1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-redisrules1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_redisrules1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_autojob1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_autojob1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["job_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-autojob1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_autojob1_arecord" {
  count   = "${var.environment_appdev["num_autojob1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-autojob1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_autojob1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_storm1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_storm1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["storm_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-storm1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_storm1_arecord" {
  count   = "${var.environment_appdev["num_storm1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-storm1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_storm1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_kafka1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_kafka1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["kafka_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-kafka1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_kafka1_arecord" {
  count   = "${var.environment_appdev["num_kafka1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-kafka1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_kafka1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_zkkafka1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_zkkafka1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-zkkafka1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_zkkafka1_arecord" {
  count   = "${var.environment_appdev["num_zkkafka1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-zkkafka1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_zkkafka1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_pubsub1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_pubsub1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-pubsub1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_pubsub1_arecord" {
  count   = "${var.environment_appdev["num_pubsub1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-pubsub1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_pubsub1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_zkstorm1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_zkstorm1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-zkstorm1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_zkstorm1_arecord" {
  count   = "${var.environment_appdev["num_zkstorm1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-zkstorm1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_zkstorm1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_nimbus1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_nimbus1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-nimbus1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_nimbus1_arecord" {
  count   = "${var.environment_appdev["num_nimbus1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-nimbus1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_nimbus1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_appcache1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_appcache1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-appcache1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_appcache1_arecord" {
  count   = "${var.environment_appdev["num_appcache1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-appcache1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_appcache1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_discovery1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_discovery1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-discovery1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_discovery1_arecord" {
  count   = "${var.environment_appdev["num_discovery1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-discovery1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_discovery1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_proxyout1" {
  key_name                    = "internal_apps"
  count                       = "${var.environment_appdev["num_proxyout1_hosts"]}"
  ami                         = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type               = "${var.environment_appdev["app_instance_type"]}"
  subnet_id                   = "${aws_subnet.appdev_us_east_1d_dmz.id}"
  associate_public_ip_address = true

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev.id}",
    "${aws_security_group.appdev_proxyout_host.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-proxyout1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_eip" "appdev_proxyout1_eip" {
  vpc      = true
  instance = "${aws_instance.appdev_proxyout1.id}"
}

resource "aws_route53_record" "appdev_proxyout1_arecord" {
  count   = "${var.environment_appdev["num_proxyout1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-proxyout1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_proxyout1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_whoisdb1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_whoisdb1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["db_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"
  ebs_optimized = "true"
  private_ip    = "${lookup(var.appdev_whoisdb1_ips,count.index)}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  ebs_block_device {
    device_name           = "${var.environment_appdev["db_volume_device_name"]}"
    volume_type           = "gp2"
    volume_size           = "512"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev.id}",
    "${aws_security_group.appdev_dbhost.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-whoisdb1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_whoisdb1_arecord" {
  count   = "${var.environment_appdev["num_whoisdb1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-whoisdb1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_whoisdb1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_security_group" "appdev_toolsproxy" {
  name        = "appdev_toolsproxy"
  description = "Allow access to rabbit servers"
  vpc_id      = "${aws_vpc.appdev.id}"

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = "${var.aloha_vpn_cidr_blocks}"
  }

  ingress {
    from_port = 22
    to_port   = 22
    protocol  = "tcp"

    cidr_blocks = [
      "${aws_instance.appdev_bastion.private_ip}/32",
      "${aws_instance.appdev_bastion.public_ip}/32",
    ]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "appdev_toolsproxy1" {
  key_name                    = "internal_apps"
  count                       = "${var.environment_appdev["num_toolsproxy1_hosts"]}"
  ami                         = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type               = "${var.environment_appdev["lightweight_instance_type"]}"
  subnet_id                   = "${aws_subnet.appdev_us_east_1d_dmz.id}"
  associate_public_ip_address = true

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_toolsproxy.id}",
    "${aws_security_group.appdev.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-toolsproxy1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_eip" "appdev_toolsproxy1" {
  count    = "${var.environment_appdev["num_toolsproxy1_hosts"]}"
  instance = "${element(aws_instance.appdev_toolsproxy1.*.id, count.index)}"
  vpc      = true
}

resource "aws_route53_record" "appdev_toolsproxy1_arecord" {
  count   = "${var.environment_appdev["num_toolsproxy1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-toolsproxy1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_toolsproxy1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_vault1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_vault1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_consulhost.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-vault1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_vault1_arecord" {
  count   = "${var.environment_appdev["num_vault1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-vault1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_vault1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_consul1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_consul1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_consulhost.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-consul1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_consul1_arecord" {
  count   = "${var.environment_appdev["num_consul1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-consul1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_consul1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_route53_record" "consul_dev_pardot_com_Arecord" {
  zone_id = "${aws_route53_zone.dev_pardot_com.zone_id}"
  name    = "consul.${aws_route53_zone.dev_pardot_com.name}"
  records = ["${aws_instance.appdev_toolsproxy1.public_ip}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_iam_user" "cephthumbs_sysacct" {
  name = "sa_cephthumbs"
}

resource "aws_s3_bucket" "cephthumbs-s3-filestore" {
  bucket = "cephthumbs-s3-filestore"

  lifecycle_rule {
    prefix  = "/"
    enabled = true

    expiration {
      days = 30
    }
  }

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "allow cephthumbs sysacct",
      "Effect": "Allow",
      "Principal": {
        "AWS": "${aws_iam_user.cephthumbs_sysacct.arn}"
      },
      "Action": "s3:*",
      "Resource": [
        "arn:aws:s3:::cephthumbs-s3-filestore",
        "arn:aws:s3:::cephthumbs-s3-filestore/*"
      ]
    }
  ]
}
EOF

  tags {
    Name      = "cephthumbs-s3-filestore"
    terraform = "true"
  }
}

resource "aws_s3_bucket" "cephthumbs-s3-filestore-long" {
  bucket = "cephthumbs-s3-filestore-long"
  acl    = "private"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "allow cephthumbs sysacct",
      "Effect": "Allow",
      "Principal": {
        "AWS": "${aws_iam_user.cephthumbs_sysacct.arn}"
      },
      "Action": "s3:*",
      "Resource": [
        "arn:aws:s3:::cephthumbs-s3-filestore-long",
        "arn:aws:s3:::cephthumbs-s3-filestore-long/*"
      ]
    }
  ]
}
EOF

  tags {
    Name      = "cephthumbs-s3-filestore"
    terraform = "true"
  }
}

resource "aws_route53_record" "logs_dev_pardot_com_arecord" {
  zone_id = "${aws_route53_zone.dev_pardot_com.zone_id}"
  name    = "logs.${aws_route53_zone.dev_pardot_com.name}"
  records = ["${aws_eip.appdev_toolsproxy1.public_ip}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_indexer1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_indexer1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_indexerhost.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-indexer1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_security_group" "appdev_indexerhost" {
  name        = "appdev_indexerhost"
  description = "Allow access through the toolsproxy and from apphosts"
  vpc_id      = "${aws_vpc.appdev.id}"

  ingress {
    from_port = 5601
    to_port   = 5601
    protocol  = "tcp"

    security_groups = [
      "${aws_security_group.appdev_toolsproxy.id}",
    ]
  }

  ingress {
    from_port = 6379
    to_port   = 6379
    protocol  = "tcp"

    cidr_blocks = [
      "${aws_vpc.appdev.cidr_block}",
    ]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_route53_record" "appdev_indexer1_arecord" {
  count   = "${var.environment_appdev["num_indexer1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-indexer1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_indexer1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_ns1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_ns1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_nshost.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-ns1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_security_group" "appdev_nshost" {
  name        = "appdev_nshost"
  description = "Allow access from apphosts"
  vpc_id      = "${aws_vpc.appdev.id}"

  # allow app.dev hosts access to tcp/53
  ingress {
    from_port = 53
    to_port   = 53
    protocol  = "tcp"

    cidr_blocks = [
      "${aws_vpc.appdev.cidr_block}",
    ]
  }

  # allow app.dev hosts access to udp/53
  ingress {
    from_port = 53
    to_port   = 53
    protocol  = "udp"

    cidr_blocks = [
      "${aws_vpc.appdev.cidr_block}",
    ]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_route53_record" "appdev_ns1_arecord" {
  count   = "${var.environment_appdev["num_ns1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-ns1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_ns1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_security_group" "appdev_metricshost" {
  name        = "appdev_metricshost"
  description = "app.dev metrics cluster"
  vpc_id      = "${aws_vpc.appdev.id}"

  # anywhere internal -> statsd/brubeck udp/8125
  ingress {
    from_port = 8125
    to_port   = 8125
    protocol  = "udp"

    cidr_blocks = [
      "${aws_vpc.appdev.cidr_block}",
    ]
  }

  # anywhere internal -> graphite tcp/2003
  ingress {
    from_port = 2003
    to_port   = 2003
    protocol  = "tcp"

    cidr_blocks = [
      "${aws_vpc.appdev.cidr_block}",
    ]
  }

  # toolsproxy -> graphite tcp/8080
  ingress {
    from_port = 8080
    to_port   = 8080
    protocol  = "tcp"

    security_groups = [
      "${aws_security_group.appdev_toolsproxy.id}",
    ]
  }

  # toolsproxy -> grafana tcp/3000
  ingress {
    from_port = 3000
    to_port   = 3000
    protocol  = "tcp"

    security_groups = [
      "${aws_security_group.appdev_toolsproxy.id}",
    ]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "appdev_metrics1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_metrics1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["metrics_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_metricshost.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-metrics1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_ebs_volume" "appdev_metrics1_opt_volume" {
  count             = "${var.environment_appdev["num_metrics1_hosts"]}"
  availability_zone = "us-east-1d"
  type              = "gp2"
  size              = "512"
}

resource "aws_volume_attachment" "appdev_metrics1_opt_volume_attachment" {
  count       = "${var.environment_appdev["num_metrics1_hosts"]}"
  device_name = "/dev/xvdf"
  volume_id   = "${element(aws_ebs_volume.appdev_metrics1_opt_volume.*.id, count.index)}"
  instance_id = "${element(aws_instance.appdev_metrics1.*.id, count.index)}"
}

resource "aws_route53_record" "appdev_metrics1_arecord" {
  count   = "${var.environment_appdev["num_metrics1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-metrics1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_metrics1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_route53_record" "graphite_dev_pardot_com_Arecord" {
  zone_id = "${aws_route53_zone.dev_pardot_com.zone_id}"
  name    = "graphite.${aws_route53_zone.dev_pardot_com.name}"
  records = ["${aws_instance.appdev_toolsproxy1.public_ip}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_route53_record" "grafana_dev_pardot_com_Arecord" {
  zone_id = "${aws_route53_zone.dev_pardot_com.zone_id}"
  name    = "grafana.${aws_route53_zone.dev_pardot_com.name}"
  records = ["${aws_instance.appdev_toolsproxy1.public_ip}"]
  type    = "A"
  ttl     = "900"
}

resource "aws_instance" "appdev_lba1" {
  key_name      = "internal_apps"
  count         = "${var.environment_appdev["num_lba1_hosts"]}"
  ami           = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id     = "${aws_subnet.appdev_us_east_1d_dmz.id}"

  root_block_device {
    volume_type           = "gp2"
    volume_size           = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev.id}",
    "${aws_security_group.appdev_app_lb.id}",
  ]

  tags {
    Name      = "${var.environment_appdev["pardot_env_id"]}-lba1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_eip" "appdev_lba1_eip" {
  vpc      = true
  instance = "${aws_instance.appdev_lba1.id}"
}

resource "aws_route53_record" "appdev_lba1_arecord" {
  count   = "${var.environment_appdev["num_lba1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name    = "${var.environment_appdev["pardot_env_id"]}-lba1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_lba1.*.private_ip, count.index)}"]
  type    = "A"
  ttl     = "900"
}