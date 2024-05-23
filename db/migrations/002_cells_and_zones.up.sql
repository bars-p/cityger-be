-- *
-- * Cells and zones
-- *

CREATE TABLE IF NOT EXISTS  "base_zone" (
  "id" UUID NOT NULL DEFAULT gen_random_uuid(),
  "idx" INTEGER NOT NULL,
  "active" BOOLEAN NOT NULL DEFAULT TRUE,
  --! "centroid" point,
  --! "geom" polygon,
  "project_id" INTEGER NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "FK_project" FOREIGN KEY ("project_id") REFERENCES "project"("id") ON DELETE CASCADE
);
