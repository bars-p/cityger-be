-- *
-- * Projects environment
-- *

CREATE TABLE IF NOT EXISTS "role"(
    "id" SERIAL, 
    "name" TEXT UNIQUE NOT NULL,
    "info" TEXT NOT NULL DEFAULT '',
    "active" BOOLEAN NOT NULL DEFAULT TRUE,
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "client"(
    "id" SERIAL, 
    "name" TEXT UNIQUE NOT NULL,
    "info" TEXT NOT NULL DEFAULT '', 
    "created" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "user"(
  "id" SERIAL,
  "client_id" INTEGER NOT NULL,
  "role_id" INTEGER NOT NULL, 
  "login" TEXT NOT NULL,
  "name" TEXT NOT NULL DEFAULT '',
  "setup" JSONB NOT NULL DEFAULT '{}'::JSONB,
  "active" BOOLEAN NOT NULL DEFAULT TRUE,
  "created" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  UNIQUE ("client_id", "login"),
  CONSTRAINT "FK_client" FOREIGN KEY ("client_id") REFERENCES "client"("id") ON DELETE CASCADE,
  CONSTRAINT "FK_role" FOREIGN KEY ("role_id") REFERENCES "role"("id") 
);

CREATE TABLE IF NOT EXISTS "project" (
  "id" SERIAL,
  "client_id" INTEGER NOT NULL,
  "name" TEXT NOT NULL, 
  "info" TEXT NOT NULL DEFAULT '',
  "cell_size" INTEGER NOT NULL DEFAULT 100,
  "cell_area" NUMERIC NOT NULL,
  "cell_count" INTEGER NOT NULL DEFAULT 0,
  "zone_size" INTEGER NOT NULL DEFAULT 1000,
  "zone_area" NUMERIC NOT NULL,
  "zone_count" INTEGER NOT NULL DEFAULT 0,
  "created" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "by_user_id" INTEGER NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "FK_user" FOREIGN KEY ("by_user_id") REFERENCES "user"("id"),
  CONSTRAINT "FK_client" FOREIGN KEY ("client_id") REFERENCES "client"("id") ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "scenario"(
  "id" SERIAL,
  "project_id" INTEGER NOT NULL,
  "name" TEXT NOT NULL,
  "info" TEXT NOT NULL DEFAULT '',
  "created" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "by_user_id" INTEGER NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "FK_user" FOREIGN KEY ("by_user_id") REFERENCES "user"("id"),
  CONSTRAINT "FK_project" FOREIGN KEY ("project_id") REFERENCES "project"("id") ON DELETE CASCADE
);
