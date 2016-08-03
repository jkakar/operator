//// App.dev environment
//// managed by pd-bread@salesforce.com


variable "environment_appdev" {
  type = "map"

  default = {
    env_name = "appdev"
    pardot_env_id = "pardot2"
    dc_id = "ue1"
    app_instance_type = "m4.large"
    job_instance_type = "m4.large"
    db_instance_type = "m4.2xlarge"
    db_volume_device_name = "/dev/xvdf"
    num_globaldb1_hosts = 2
    num_dbshard1_hosts = 4
    num_whoisdb1_hosts = 2
    num_app1_hosts = 2
    num_thumbs1_hosts = 1
    num_redisjob1_hosts = 2
    num_jobmanager1_hosts = 2
    num_push1_hosts = 2
    num_provisioning1_hosts = 1
    num_rabbit1_hosts = 3
    num_redisrules1_hosts = 2
    num_autojob1_hosts = 1
    num_storm1_hosts = 1
    num_kafka1_hosts = 1
    num_zkkafka1_hosts = 1
    num_pubsub1_hosts = 2
    num_zkstorm1_hosts = 3
    num_nimbus1_hosts = 1
    num_appcache1_hosts = 2
    num_discovery1_hosts = 3
    num_proxyout1_hosts = 1
  }
}

variable "appdev_globaldb1_ips" {
  default = {
    "0" = "172.26.80.125"
    "1" = "172.26.81.47"
  }
}

variable "appdev_dbshard1_ips" {
  default = {
    "0" = "172.26.92.23"
    "1" = "172.26.93.40"
    "2" = "172.26.75.74"
    "3" = "172.26.69.79"
  }
}

////
//// TEMPLATES
////
//
//// EC2 DB INSTANCE: replace "lbl" w/ "servicename"
//resource "aws_instance" "appdev_lbl1" {
//  key_name = "internal_apps"
//  count = "${var.environment_appdev["num_lbl1_hosts"]}"
//  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
//  instance_type = "${var.environment_appdev["db_instance_type}"
//  subnet_id = "${var.environment_appdev["subnet_id}"
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
//    "${aws_security_group.appdev_apphost.id}"
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
//  instance_type = "${var.environment_appdev["db_instance_type}"
//  subnet_id = "${var.environment_appdev["subnet_id}"
//  root_block_device {
//    volume_type = "gp2"
//    volume_size = "50"
//    delete_on_termination = true
//  }
//  vpc_security_group_ids = [
//    "${aws_security_group.appdev_vpc_default.id}",
//    "${aws_security_group.appdev_dbhost.id}",
//    "${aws_security_group.appdev_apphost.id}"
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
//  instance_type = "${var.environment_appdev["db_instance_type}"
//  subnet_id = "${var.environment_appdev["subnet_id}"
//  ebs_optimized = "true"
//  root_block_device {
//    volume_type = "gp2"
//    volume_size = "50"
//    delete_on_termination = true
//  }
//  vpc_security_group_ids = [
//    "${aws_security_group.appdev_vpc_default.id}",
//    "${aws_security_group.appdev_dbhost.id}",
//    "${aws_security_group.appdev_apphost.id}"
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


resource "aws_security_group" "appdev_apphost" {
  name = "appdev_apphost"
  description = "Allow HTTP/HTTPS traffic from appdev vpc"
  vpc_id = "${aws_vpc.appdev.id}"

  ingress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    self = true
  }

  # allow health check from ELBs
  ingress {
    from_port = 80
    to_port = 80
    protocol = "tcp"
    security_groups = [
      "${aws_security_group.appdev_sfdc_vpn_http_https.id}"
    ]
  }

  egress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "appdev_dbhost" {
  name = "appdev_dbhost"
  description = "Allow MYSQL traffic from appdev apphosts"
  vpc_id = "${aws_vpc.appdev.id}"

  ingress {
    from_port = 3306
    to_port = 3306
    protocol = "tcp"
    security_groups = [
      "${aws_security_group.appdev_apphost.id}"
    ]
  }

  egress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_elb" "appdev_app_elb" {
  name = "${var.environment_appdev["env_name"]}-app-elb"
  security_groups = [
    "${aws_security_group.appdev_sfdc_vpn_http_https.id}"
  ]
  subnets = [
    "${aws_subnet.appdev_us_east_1a.id}",
    "${aws_subnet.appdev_us_east_1c.id}",
    "${aws_subnet.appdev_us_east_1d.id}",
    "${aws_subnet.appdev_us_east_1e.id}"
  ]
  cross_zone_load_balancing = true
  connection_draining = true
  connection_draining_timeout = 30
  instances = ["${aws_instance.appdev_app1.*.id}"]

  listener {
    lb_port = 443
    lb_protocol = "https"
    instance_port = 80
    instance_protocol = "http"
    ssl_certificate_id = "arn:aws:iam::${var.pardotops_account_number}:server-certificate/dev.pardot.com-2016-with-intermediate"
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
    target = "HTTP:80/home/ping"
    interval = 5
  }

  tags {
    Name = "appdev_public_elb"
  }
}

#TODO: rename/copy this from appdev.dev.pardot.com to app.dev.pardot.com
resource "aws_route53_record" "appdev_dev_pardot_com_CNAMErecord" {
  zone_id = "${aws_route53_zone.dev_pardot_com.zone_id}"
  name = "appdev.${aws_route53_zone.dev_pardot_com.name}"
  records = ["${aws_elb.appdev_app_elb.dns_name}"]
  type = "CNAME"
  ttl = "900"
}

resource "aws_instance" "appdev_globaldb1" {
  key_name = "internal_apps"
  count = "${var.environment_appdev["num_globaldb1_hosts"]}"
  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["db_instance_type"]}"
  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
  ebs_optimized = "true"
  private_ip = "${lookup(var.appdev_globaldb1_ips,count.index)}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "50"
    delete_on_termination = true
  }
  ebs_block_device {
    device_name = "${var.environment_appdev["db_volume_device_name"]}"
    volume_type = "gp2"
    volume_size = "512"
    delete_on_termination = true
  }
  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_apphost.id}",
    "${aws_security_group.appdev_dbhost.id}"
  ]
  tags {
    Name = "${var.environment_appdev["pardot_env_id"]}-globaldb1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_globaldb1_arecord" {
  count = "${var.environment_appdev["num_globaldb1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name = "${var.environment_appdev["pardot_env_id"]}-globaldb1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_globaldb1.*.private_ip, count.index)}"]
  type = "A"
  ttl = "900"
}

resource "aws_instance" "appdev_dbshard1" {
  key_name = "internal_apps"
  count = "${var.environment_appdev["num_dbshard1_hosts"]}"
  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["db_instance_type"]}"
  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
  ebs_optimized = "true"
  private_ip = "${lookup(var.appdev_dbshard1_ips,count.index)}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "50"
    delete_on_termination = true
  }
  ebs_block_device {
    device_name = "${var.environment_appdev["db_volume_device_name"]}"
    volume_type = "gp2"
    volume_size = "512"
    delete_on_termination = true
  }
  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_apphost.id}",
    "${aws_security_group.appdev_dbhost.id}"
  ]
  tags {
    Name = "${var.environment_appdev["pardot_env_id"]}-dbshard1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_dbshard1_arecord" {
  count = "${var.environment_appdev["num_dbshard1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name = "${var.environment_appdev["pardot_env_id"]}-dbshard1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_dbshard1.*.private_ip, count.index)}"]
  type = "A"
  ttl = "900"
}

resource "aws_instance" "appdev_app1" {
  key_name = "internal_apps"
  count = "${var.environment_appdev["num_app1_hosts"]}"
  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "50"
    delete_on_termination = true
  }
  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_apphost.id}"
  ]
  tags {
    Name = "${var.environment_appdev["pardot_env_id"]}-app1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_app1_arecord" {
  count = "${var.environment_appdev["num_app1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name = "${var.environment_appdev["pardot_env_id"]}-app1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_app1.*.private_ip, count.index)}"]
  type = "A"
  ttl = "900"
}

resource "aws_instance" "appdev_thumbs1" {
  key_name = "internal_apps"
  count = "${var.environment_appdev["num_thumbs1_hosts"]}"
  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "50"
    delete_on_termination = true
  }

  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_apphost.id}"
  ]
  tags {
    Name = "${var.environment_appdev["pardot_env_id"]}-thumbs1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_thumbs1_arecord" {
  count = "${var.environment_appdev["num_thumbs1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name = "${var.environment_appdev["pardot_env_id"]}-thumbs1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_thumbs1.*.private_ip, count.index)}"]
  type = "A"
  ttl = "900"
}

resource "aws_instance" "appdev_redisjob1" {
  key_name = "internal_apps"
  count = "${var.environment_appdev["num_redisjob1_hosts"]}"
  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["job_instance_type"]}"
  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "50"
    delete_on_termination = true
  }
  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_apphost.id}"
  ]
  tags {
    Name = "${var.environment_appdev["pardot_env_id"]}-redisjob1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_redisjob1_arecord" {
  count = "${var.environment_appdev["num_redisjob1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name = "${var.environment_appdev["pardot_env_id"]}-redisjob1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_redisjob1.*.private_ip, count.index)}"]
  type = "A"
  ttl = "900"
}

resource "aws_instance" "appdev_jobmanager1" {
  key_name = "internal_apps"
  count = "${var.environment_appdev["num_jobmanager1_hosts"]}"
  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "50"
    delete_on_termination = true
  }
  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_apphost.id}"
  ]
  tags {
    Name = "${var.environment_appdev["pardot_env_id"]}-jobmanager1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_jobmanager1_arecord" {
  count = "${var.environment_appdev["num_jobmanager1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name = "${var.environment_appdev["pardot_env_id"]}-jobmanager1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_jobmanager1.*.private_ip, count.index)}"]
  type = "A"
  ttl = "900"
}

resource "aws_instance" "appdev_push1" {
  key_name = "internal_apps"
  count = "${var.environment_appdev["num_push1_hosts"]}"
  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "50"
    delete_on_termination = true
  }
  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_apphost.id}"
  ]
  tags {
    Name = "${var.environment_appdev["pardot_env_id"]}-push1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_push1_arecord" {
  count = "${var.environment_appdev["num_push1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name = "${var.environment_appdev["pardot_env_id"]}-push1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_push1.*.private_ip, count.index)}"]
  type = "A"
  ttl = "900"
}

resource "aws_instance" "appdev_provisioning1" {
  key_name = "internal_apps"
  count = "${var.environment_appdev["num_provisioning1_hosts"]}"
  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "50"
    delete_on_termination = true
  }
  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_apphost.id}"
  ]
  tags {
    Name = "${var.environment_appdev["pardot_env_id"]}-provisioning1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_provisioning1_arecord" {
  count = "${var.environment_appdev["num_provisioning1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name = "${var.environment_appdev["pardot_env_id"]}-provisioning1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_provisioning1.*.private_ip, count.index)}"]
  type = "A"
  ttl = "900"
}

resource "aws_instance" "appdev_rabbit1" {
  key_name = "internal_apps"
  count = "${var.environment_appdev["num_rabbit1_hosts"]}"
  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "50"
    delete_on_termination = true
  }
  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_apphost.id}"
  ]
  tags {
    Name = "${var.environment_appdev["pardot_env_id"]}-rabbit1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_rabbit1_arecord" {
  count = "${var.environment_appdev["num_rabbit1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name = "${var.environment_appdev["pardot_env_id"]}-rabbit1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_rabbit1.*.private_ip, count.index)}"]
  type = "A"
  ttl = "900"
}

resource "aws_instance" "appdev_redisrules1" {
  key_name = "internal_apps"
  count = "${var.environment_appdev["num_redisrules1_hosts"]}"
  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "50"
    delete_on_termination = true
  }
  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_apphost.id}"
  ]
  tags {
    Name = "${var.environment_appdev["pardot_env_id"]}-redisrules1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_redisrules1_arecord" {
  count = "${var.environment_appdev["num_redisrules1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name = "${var.environment_appdev["pardot_env_id"]}-redisrules1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_redisrules1.*.private_ip, count.index)}"]
  type = "A"
  ttl = "900"
}

resource "aws_instance" "appdev_autojob1" {
  key_name = "internal_apps"
  count = "${var.environment_appdev["num_autojob1_hosts"]}"
  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["job_instance_type"]}"
  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "50"
    delete_on_termination = true
  }
  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_apphost.id}"
  ]
  tags {
    Name = "${var.environment_appdev["pardot_env_id"]}-autojob1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_autojob1_arecord" {
  count = "${var.environment_appdev["num_autojob1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name = "${var.environment_appdev["pardot_env_id"]}-autojob1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_autojob1.*.private_ip, count.index)}"]
  type = "A"
  ttl = "900"
}

resource "aws_instance" "appdev_storm1" {
  key_name = "internal_apps"
  count = "${var.environment_appdev["num_storm1_hosts"]}"
  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "50"
    delete_on_termination = true
  }
  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_apphost.id}"
  ]
  tags {
    Name = "${var.environment_appdev["pardot_env_id"]}-storm1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_storm1_arecord" {
  count = "${var.environment_appdev["num_storm1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name = "${var.environment_appdev["pardot_env_id"]}-storm1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_storm1.*.private_ip, count.index)}"]
  type = "A"
  ttl = "900"
}

resource "aws_instance" "appdev_kafka1" {
  key_name = "internal_apps"
  count = "${var.environment_appdev["num_kafka1_hosts"]}"
  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "50"
    delete_on_termination = true
  }
  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_apphost.id}"
  ]
  tags {
    Name = "${var.environment_appdev["pardot_env_id"]}-kafka1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_kafka1_arecord" {
  count = "${var.environment_appdev["num_kafka1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name = "${var.environment_appdev["pardot_env_id"]}-kafka1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_kafka1.*.private_ip, count.index)}"]
  type = "A"
  ttl = "900"
}

resource "aws_instance" "appdev_zkkafka1" {
  key_name = "internal_apps"
  count = "${var.environment_appdev["num_zkkafka1_hosts"]}"
  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "50"
    delete_on_termination = true
  }
  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_apphost.id}"
  ]
  tags {
    Name = "${var.environment_appdev["pardot_env_id"]}-zkkafka1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_zkkafka1_arecord" {
  count = "${var.environment_appdev["num_zkkafka1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name = "${var.environment_appdev["pardot_env_id"]}-zkkafka1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_zkkafka1.*.private_ip, count.index)}"]
  type = "A"
  ttl = "900"
}

resource "aws_instance" "appdev_pubsub1" {
  key_name = "internal_apps"
  count = "${var.environment_appdev["num_pubsub1_hosts"]}"
  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "50"
    delete_on_termination = true
  }
  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_apphost.id}"
  ]
  tags {
    Name = "${var.environment_appdev["pardot_env_id"]}-pubsub1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_pubsub1_arecord" {
  count = "${var.environment_appdev["num_pubsub1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name = "${var.environment_appdev["pardot_env_id"]}-pubsub1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_pubsub1.*.private_ip, count.index)}"]
  type = "A"
  ttl = "900"
}

resource "aws_instance" "appdev_zkstorm1" {
  key_name = "internal_apps"
  count = "${var.environment_appdev["num_zkstorm1_hosts"]}"
  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "50"
    delete_on_termination = true
  }
  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_apphost.id}"
  ]
  tags {
    Name = "${var.environment_appdev["pardot_env_id"]}-zkstorm1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_zkstorm1_arecord" {
  count = "${var.environment_appdev["num_zkstorm1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name = "${var.environment_appdev["pardot_env_id"]}-zkstorm1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_zkstorm1.*.private_ip, count.index)}"]
  type = "A"
  ttl = "900"
}

resource "aws_instance" "appdev_nimbus1" {
  key_name = "internal_apps"
  count = "${var.environment_appdev["num_nimbus1_hosts"]}"
  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "50"
    delete_on_termination = true
  }
  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_apphost.id}"
  ]
  tags {
    Name = "${var.environment_appdev["pardot_env_id"]}-nimbus1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_nimbus1_arecord" {
  count = "${var.environment_appdev["num_nimbus1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name = "${var.environment_appdev["pardot_env_id"]}-nimbus1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_nimbus1.*.private_ip, count.index)}"]
  type = "A"
  ttl = "900"
}

resource "aws_instance" "appdev_appcache1" {
  key_name = "internal_apps"
  count = "${var.environment_appdev["num_appcache1_hosts"]}"
  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "50"
    delete_on_termination = true
  }
  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_apphost.id}"
  ]
  tags {
    Name = "${var.environment_appdev["pardot_env_id"]}-appcache1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_appcache1_arecord" {
  count = "${var.environment_appdev["num_appcache1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name = "${var.environment_appdev["pardot_env_id"]}-appcache1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_appcache1.*.private_ip, count.index)}"]
  type = "A"
  ttl = "900"
}

resource "aws_instance" "appdev_discovery1" {
  key_name = "internal_apps"
  count = "${var.environment_appdev["num_discovery1_hosts"]}"
  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "50"
    delete_on_termination = true
  }
  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_apphost.id}"
  ]
  tags {
    Name = "${var.environment_appdev["pardot_env_id"]}-discovery1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_discovery1_arecord" {
  count = "${var.environment_appdev["num_discovery1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name = "${var.environment_appdev["pardot_env_id"]}-discovery1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_discovery1.*.private_ip, count.index)}"]
  type = "A"
  ttl = "900"
}

resource "aws_instance" "appdev_proxyout1" {
  key_name = "internal_apps"
  count = "${var.environment_appdev["num_proxyout1_hosts"]}"
  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "50"
    delete_on_termination = true
  }
  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_apphost.id}"
  ]
  tags {
    Name = "${var.environment_appdev["pardot_env_id"]}-proxyout1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_proxyout1_arecord" {
  count = "${var.environment_appdev["num_proxyout1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name = "${var.environment_appdev["pardot_env_id"]}-proxyout1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_proxyout1.*.private_ip, count.index)}"]
  type = "A"
  ttl = "900"
}

resource "aws_instance" "appdev_whoisdb1" {
  key_name = "internal_apps"
  count = "${var.environment_appdev["num_whoisdb1_hosts"]}"
  ami = "${var.centos_6_hvm_50gb_chefdev_ami}"
  instance_type = "${var.environment_appdev["app_instance_type"]}"
  subnet_id = "${aws_subnet.appdev_us_east_1d.id}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "50"
    delete_on_termination = true
  }
  vpc_security_group_ids = [
    "${aws_security_group.appdev_vpc_default.id}",
    "${aws_security_group.appdev_apphost.id}"
  ]
  tags {
    Name = "${var.environment_appdev["pardot_env_id"]}-whoisdb1-${count.index + 1}-${var.environment_appdev["dc_id"]}"
    terraform = "true"
  }
}

resource "aws_route53_record" "appdev_whoisdb1_arecord" {
  count = "${var.environment_appdev["num_whoisdb1_hosts"]}"
  zone_id = "${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.zone_id}"
  name = "${var.environment_appdev["pardot_env_id"]}-whoisdb1-${count.index + 1}-${var.environment_appdev["dc_id"]}.${aws_route53_zone.appdev_aws_pardot_com_hosted_zone.name}"
  records = ["${element(aws_instance.appdev_whoisdb1.*.private_ip, count.index)}"]
  type = "A"
  ttl = "900"
}
