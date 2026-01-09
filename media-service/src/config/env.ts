import dotenv from "dotenv";

dotenv.config();

export const env = {
  PORT: process.env.PORT || "8085",
  AWS_REGION: process.env.AWS_REGION!,
  S3_BUCKET_NAME: process.env.S3_BUCKET_NAME!,
  CDN_BASE_URL: process.env.CDN_BASE_URL!
};
