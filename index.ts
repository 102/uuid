import { randomUUID } from "node:crypto";

Bun.serve({
  fetch: (req) => {
    return new Response(randomUUID());
  },
});
