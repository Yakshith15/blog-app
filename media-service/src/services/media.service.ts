import { PutObjectCommand } from "@aws-sdk/client-s3";
import { getSignedUrl } from "@aws-sdk/s3-request-presigner";
import { randomUUID } from "crypto";
import { s3Client } from "./s3.client";
import { env } from "../config/env";

type GenerateUploadUrlInput = {
  fileName: string;
  contentType: string;
  entityType: "BLOG" | "COMMENT" | "PROFILE";
};

export async function generateUploadUrl(input: GenerateUploadUrlInput) {
  const { fileName, contentType, entityType } = input;

  if (!contentType.startsWith("image/")) {
    throw new Error("Only image uploads are allowed");
  }

  const mediaId = randomUUID();

  const objectKey = `${entityType.toLowerCase()}/${mediaId}/${fileName}`;

  const command = new PutObjectCommand({
    Bucket: env.S3_BUCKET_NAME,
    Key: objectKey,
    ContentType: contentType
  });

  const uploadUrl = await getSignedUrl(s3Client, command, {
    expiresIn: 300
  });

  const publicUrl = `${env.CDN_BASE_URL}/${objectKey}`;

  return {
    mediaId,
    uploadUrl,
    publicUrl
  };
}
