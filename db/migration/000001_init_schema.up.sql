CREATE TABLE "urls" (
  "id" bigserial PRIMARY KEY,
  "code" text UNIQUE NOT NULL,
  "original_url" text NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE UNIQUE INDEX ON "urls" ("code");