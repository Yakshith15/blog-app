import express from "express";
import mediaRoutes from "./routes/media.routes";

const app = express();

app.use(express.json());
app.use("/media", mediaRoutes);

app.get("/health", (_, res) => {
  res.json({ status: "UP" });
});

export default app;
