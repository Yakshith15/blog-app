import { Router } from "express";
import { generateUploadUrl } from "../services/media.service";

const router = Router();

router.post("/upload-url", async (req, res) => {
  try {
    const { fileName, contentType, entityType } = req.body;

    if (!fileName || !contentType || !entityType) {
      return res.status(400).json({
        message: "fileName, contentType and entityType are required"
      });
    }

    const result = await generateUploadUrl({
      fileName,
      contentType,
      entityType
    });

    res.status(200).json(result);
  } catch (error: any) {
    res.status(400).json({
      message: error.message || "Failed to generate upload URL"
    });
  }
});

export default router;
