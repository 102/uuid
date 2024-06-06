import { randomUUID } from "node:crypto";

Bun.serve({
  port: process.env.PORT,
  fetch: () => {
    return new Response(randomUUID());
  },
});
