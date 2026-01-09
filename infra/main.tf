provider "aws" {
  region = var.aws_region
}

resource "aws_s3_bucket" "media_bucket" {
  bucket = var.media_bucket_name

  tags = {
    Service = "media-service"
    Env     = var.environment
  }
}
