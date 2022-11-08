ALTER TABLE "public"."players" ADD COLUMN "transfered" bool NOT NULL DEFAULT 'false';
UPDATE metadata SET value = 'v2', updated_at = NOW() WHERE key = 'version';