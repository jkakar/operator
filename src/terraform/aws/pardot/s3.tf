resource "aws_s3_bucket" "pardot_pdo" {
  bucket = "pardot-pdo"
  acl = "bucket-owner-full-control"
}
