import { Router } from "express";

const router = Router();

router.post("/upload-url", async (_req, res) => {
  res.status(501).json({
    message: "Not implemented yet"
  });
});

export default router;
