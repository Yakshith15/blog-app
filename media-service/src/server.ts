import app from "./app";
import { env } from "./config/env";

const port = env.PORT;


console.log("AWS REGION:", env.AWS_REGION);
app.listen(port, () => {
  console.log(`Media service running on port ${port}`);
});
