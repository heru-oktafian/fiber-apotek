/*
 Navicat Premium Dump SQL

 Source Server         : PostgresSQL-2
 Source Server Type    : PostgreSQL
 Source Server Version : 170009 (170009)
 Source Host           : 200.200.200.2:5440
 Source Catalog        : heru_db
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 170009 (170009)
 File Encoding         : 65001

 Date: 15/03/2026 03:29:48
*/


-- ----------------------------
-- Table structure for admins
-- ----------------------------
DROP TABLE IF EXISTS "public"."admins";
CREATE TABLE "public"."admins" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "email" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "password_hash" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."admins" OWNER TO "postgres";

-- ----------------------------
-- Table structure for contacts
-- ----------------------------
DROP TABLE IF EXISTS "public"."contacts";
CREATE TABLE "public"."contacts" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "email" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "message" text COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."contacts" OWNER TO "postgres";

-- ----------------------------
-- Table structure for projects
-- ----------------------------
DROP TABLE IF EXISTS "public"."projects";
CREATE TABLE "public"."projects" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "title" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "slug" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "description" text COLLATE "pg_catalog"."default" NOT NULL,
  "image_url" varchar(255) COLLATE "pg_catalog"."default",
  "is_published" bool DEFAULT false,
  "created_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP,
  "tags" varchar(255) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."projects" OWNER TO "postgres";

-- ----------------------------
-- Table structure for seo_metadata
-- ----------------------------
DROP TABLE IF EXISTS "public"."seo_metadata";
CREATE TABLE "public"."seo_metadata" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "reference_id" uuid NOT NULL,
  "reference_type" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "meta_title" varchar(255) COLLATE "pg_catalog"."default",
  "meta_desc" text COLLATE "pg_catalog"."default",
  "og_image_url" varchar(255) COLLATE "pg_catalog"."default",
  "alt_text" varchar(255) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."seo_metadata" OWNER TO "postgres";

-- ----------------------------
-- Table structure for services
-- ----------------------------
DROP TABLE IF EXISTS "public"."services";
CREATE TABLE "public"."services" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "title" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "description" text COLLATE "pg_catalog"."default" NOT NULL,
  "icon_name" varchar(100) COLLATE "pg_catalog"."default",
  "order_num" int4 DEFAULT 0,
  "created_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."services" OWNER TO "postgres";

-- ----------------------------
-- Table structure for site_settings
-- ----------------------------
DROP TABLE IF EXISTS "public"."site_settings";
CREATE TABLE "public"."site_settings" (
  "key" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "value" jsonb NOT NULL,
  "updated_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."site_settings" OWNER TO "postgres";

-- ----------------------------
-- Table structure for skills
-- ----------------------------
DROP TABLE IF EXISTS "public"."skills";
CREATE TABLE "public"."skills" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "category" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "percentage" int4 DEFAULT 0,
  "icon_url" varchar(255) COLLATE "pg_catalog"."default",
  "order_num" int4 DEFAULT 0,
  "created_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "public"."skills" OWNER TO "postgres";

-- ----------------------------
-- Function structure for armor
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."armor"(bytea);
CREATE OR REPLACE FUNCTION "public"."armor"(bytea)
  RETURNS "pg_catalog"."text" AS '$libdir/pgcrypto', 'pg_armor'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."armor"(bytea) OWNER TO "postgres";

-- ----------------------------
-- Function structure for armor
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."armor"(bytea, _text, _text);
CREATE OR REPLACE FUNCTION "public"."armor"(bytea, _text, _text)
  RETURNS "pg_catalog"."text" AS '$libdir/pgcrypto', 'pg_armor'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."armor"(bytea, _text, _text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for crypt
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."crypt"(text, text);
CREATE OR REPLACE FUNCTION "public"."crypt"(text, text)
  RETURNS "pg_catalog"."text" AS '$libdir/pgcrypto', 'pg_crypt'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."crypt"(text, text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for dearmor
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."dearmor"(text);
CREATE OR REPLACE FUNCTION "public"."dearmor"(text)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pg_dearmor'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."dearmor"(text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for decrypt
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."decrypt"(bytea, bytea, text);
CREATE OR REPLACE FUNCTION "public"."decrypt"(bytea, bytea, text)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pg_decrypt'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."decrypt"(bytea, bytea, text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for decrypt_iv
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."decrypt_iv"(bytea, bytea, bytea, text);
CREATE OR REPLACE FUNCTION "public"."decrypt_iv"(bytea, bytea, bytea, text)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pg_decrypt_iv'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."decrypt_iv"(bytea, bytea, bytea, text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for digest
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."digest"(bytea, text);
CREATE OR REPLACE FUNCTION "public"."digest"(bytea, text)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pg_digest'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."digest"(bytea, text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for digest
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."digest"(text, text);
CREATE OR REPLACE FUNCTION "public"."digest"(text, text)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pg_digest'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."digest"(text, text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for encrypt
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."encrypt"(bytea, bytea, text);
CREATE OR REPLACE FUNCTION "public"."encrypt"(bytea, bytea, text)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pg_encrypt'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."encrypt"(bytea, bytea, text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for encrypt_iv
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."encrypt_iv"(bytea, bytea, bytea, text);
CREATE OR REPLACE FUNCTION "public"."encrypt_iv"(bytea, bytea, bytea, text)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pg_encrypt_iv'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."encrypt_iv"(bytea, bytea, bytea, text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for gen_random_bytes
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."gen_random_bytes"(int4);
CREATE OR REPLACE FUNCTION "public"."gen_random_bytes"(int4)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pg_random_bytes'
  LANGUAGE c VOLATILE STRICT
  COST 1;
ALTER FUNCTION "public"."gen_random_bytes"(int4) OWNER TO "postgres";

-- ----------------------------
-- Function structure for gen_random_uuid
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."gen_random_uuid"();
CREATE OR REPLACE FUNCTION "public"."gen_random_uuid"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/pgcrypto', 'pg_random_uuid'
  LANGUAGE c VOLATILE
  COST 1;
ALTER FUNCTION "public"."gen_random_uuid"() OWNER TO "postgres";

-- ----------------------------
-- Function structure for gen_salt
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."gen_salt"(text, int4);
CREATE OR REPLACE FUNCTION "public"."gen_salt"(text, int4)
  RETURNS "pg_catalog"."text" AS '$libdir/pgcrypto', 'pg_gen_salt_rounds'
  LANGUAGE c VOLATILE STRICT
  COST 1;
ALTER FUNCTION "public"."gen_salt"(text, int4) OWNER TO "postgres";

-- ----------------------------
-- Function structure for gen_salt
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."gen_salt"(text);
CREATE OR REPLACE FUNCTION "public"."gen_salt"(text)
  RETURNS "pg_catalog"."text" AS '$libdir/pgcrypto', 'pg_gen_salt'
  LANGUAGE c VOLATILE STRICT
  COST 1;
ALTER FUNCTION "public"."gen_salt"(text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for hmac
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."hmac"(text, text, text);
CREATE OR REPLACE FUNCTION "public"."hmac"(text, text, text)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pg_hmac'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."hmac"(text, text, text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for hmac
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."hmac"(bytea, bytea, text);
CREATE OR REPLACE FUNCTION "public"."hmac"(bytea, bytea, text)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pg_hmac'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."hmac"(bytea, bytea, text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for pgp_armor_headers
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."pgp_armor_headers"(text, OUT "key" text, OUT "value" text);
CREATE OR REPLACE FUNCTION "public"."pgp_armor_headers"(IN text, OUT "key" text, OUT "value" text)
  RETURNS SETOF "pg_catalog"."record" AS '$libdir/pgcrypto', 'pgp_armor_headers'
  LANGUAGE c IMMUTABLE STRICT
  COST 1
  ROWS 1000;
ALTER FUNCTION "public"."pgp_armor_headers"(text, OUT "key" text, OUT "value" text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for pgp_key_id
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."pgp_key_id"(bytea);
CREATE OR REPLACE FUNCTION "public"."pgp_key_id"(bytea)
  RETURNS "pg_catalog"."text" AS '$libdir/pgcrypto', 'pgp_key_id_w'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."pgp_key_id"(bytea) OWNER TO "postgres";

-- ----------------------------
-- Function structure for pgp_pub_decrypt
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."pgp_pub_decrypt"(bytea, bytea, text);
CREATE OR REPLACE FUNCTION "public"."pgp_pub_decrypt"(bytea, bytea, text)
  RETURNS "pg_catalog"."text" AS '$libdir/pgcrypto', 'pgp_pub_decrypt_text'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."pgp_pub_decrypt"(bytea, bytea, text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for pgp_pub_decrypt
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."pgp_pub_decrypt"(bytea, bytea);
CREATE OR REPLACE FUNCTION "public"."pgp_pub_decrypt"(bytea, bytea)
  RETURNS "pg_catalog"."text" AS '$libdir/pgcrypto', 'pgp_pub_decrypt_text'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."pgp_pub_decrypt"(bytea, bytea) OWNER TO "postgres";

-- ----------------------------
-- Function structure for pgp_pub_decrypt
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."pgp_pub_decrypt"(bytea, bytea, text, text);
CREATE OR REPLACE FUNCTION "public"."pgp_pub_decrypt"(bytea, bytea, text, text)
  RETURNS "pg_catalog"."text" AS '$libdir/pgcrypto', 'pgp_pub_decrypt_text'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."pgp_pub_decrypt"(bytea, bytea, text, text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for pgp_pub_decrypt_bytea
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."pgp_pub_decrypt_bytea"(bytea, bytea, text);
CREATE OR REPLACE FUNCTION "public"."pgp_pub_decrypt_bytea"(bytea, bytea, text)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pgp_pub_decrypt_bytea'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."pgp_pub_decrypt_bytea"(bytea, bytea, text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for pgp_pub_decrypt_bytea
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."pgp_pub_decrypt_bytea"(bytea, bytea, text, text);
CREATE OR REPLACE FUNCTION "public"."pgp_pub_decrypt_bytea"(bytea, bytea, text, text)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pgp_pub_decrypt_bytea'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."pgp_pub_decrypt_bytea"(bytea, bytea, text, text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for pgp_pub_decrypt_bytea
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."pgp_pub_decrypt_bytea"(bytea, bytea);
CREATE OR REPLACE FUNCTION "public"."pgp_pub_decrypt_bytea"(bytea, bytea)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pgp_pub_decrypt_bytea'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."pgp_pub_decrypt_bytea"(bytea, bytea) OWNER TO "postgres";

-- ----------------------------
-- Function structure for pgp_pub_encrypt
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."pgp_pub_encrypt"(text, bytea);
CREATE OR REPLACE FUNCTION "public"."pgp_pub_encrypt"(text, bytea)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pgp_pub_encrypt_text'
  LANGUAGE c VOLATILE STRICT
  COST 1;
ALTER FUNCTION "public"."pgp_pub_encrypt"(text, bytea) OWNER TO "postgres";

-- ----------------------------
-- Function structure for pgp_pub_encrypt
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."pgp_pub_encrypt"(text, bytea, text);
CREATE OR REPLACE FUNCTION "public"."pgp_pub_encrypt"(text, bytea, text)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pgp_pub_encrypt_text'
  LANGUAGE c VOLATILE STRICT
  COST 1;
ALTER FUNCTION "public"."pgp_pub_encrypt"(text, bytea, text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for pgp_pub_encrypt_bytea
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."pgp_pub_encrypt_bytea"(bytea, bytea, text);
CREATE OR REPLACE FUNCTION "public"."pgp_pub_encrypt_bytea"(bytea, bytea, text)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pgp_pub_encrypt_bytea'
  LANGUAGE c VOLATILE STRICT
  COST 1;
ALTER FUNCTION "public"."pgp_pub_encrypt_bytea"(bytea, bytea, text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for pgp_pub_encrypt_bytea
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."pgp_pub_encrypt_bytea"(bytea, bytea);
CREATE OR REPLACE FUNCTION "public"."pgp_pub_encrypt_bytea"(bytea, bytea)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pgp_pub_encrypt_bytea'
  LANGUAGE c VOLATILE STRICT
  COST 1;
ALTER FUNCTION "public"."pgp_pub_encrypt_bytea"(bytea, bytea) OWNER TO "postgres";

-- ----------------------------
-- Function structure for pgp_sym_decrypt
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."pgp_sym_decrypt"(bytea, text);
CREATE OR REPLACE FUNCTION "public"."pgp_sym_decrypt"(bytea, text)
  RETURNS "pg_catalog"."text" AS '$libdir/pgcrypto', 'pgp_sym_decrypt_text'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."pgp_sym_decrypt"(bytea, text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for pgp_sym_decrypt
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."pgp_sym_decrypt"(bytea, text, text);
CREATE OR REPLACE FUNCTION "public"."pgp_sym_decrypt"(bytea, text, text)
  RETURNS "pg_catalog"."text" AS '$libdir/pgcrypto', 'pgp_sym_decrypt_text'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."pgp_sym_decrypt"(bytea, text, text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for pgp_sym_decrypt_bytea
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."pgp_sym_decrypt_bytea"(bytea, text, text);
CREATE OR REPLACE FUNCTION "public"."pgp_sym_decrypt_bytea"(bytea, text, text)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pgp_sym_decrypt_bytea'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."pgp_sym_decrypt_bytea"(bytea, text, text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for pgp_sym_decrypt_bytea
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."pgp_sym_decrypt_bytea"(bytea, text);
CREATE OR REPLACE FUNCTION "public"."pgp_sym_decrypt_bytea"(bytea, text)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pgp_sym_decrypt_bytea'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."pgp_sym_decrypt_bytea"(bytea, text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for pgp_sym_encrypt
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."pgp_sym_encrypt"(text, text, text);
CREATE OR REPLACE FUNCTION "public"."pgp_sym_encrypt"(text, text, text)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pgp_sym_encrypt_text'
  LANGUAGE c VOLATILE STRICT
  COST 1;
ALTER FUNCTION "public"."pgp_sym_encrypt"(text, text, text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for pgp_sym_encrypt
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."pgp_sym_encrypt"(text, text);
CREATE OR REPLACE FUNCTION "public"."pgp_sym_encrypt"(text, text)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pgp_sym_encrypt_text'
  LANGUAGE c VOLATILE STRICT
  COST 1;
ALTER FUNCTION "public"."pgp_sym_encrypt"(text, text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for pgp_sym_encrypt_bytea
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."pgp_sym_encrypt_bytea"(bytea, text, text);
CREATE OR REPLACE FUNCTION "public"."pgp_sym_encrypt_bytea"(bytea, text, text)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pgp_sym_encrypt_bytea'
  LANGUAGE c VOLATILE STRICT
  COST 1;
ALTER FUNCTION "public"."pgp_sym_encrypt_bytea"(bytea, text, text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for pgp_sym_encrypt_bytea
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."pgp_sym_encrypt_bytea"(bytea, text);
CREATE OR REPLACE FUNCTION "public"."pgp_sym_encrypt_bytea"(bytea, text)
  RETURNS "pg_catalog"."bytea" AS '$libdir/pgcrypto', 'pgp_sym_encrypt_bytea'
  LANGUAGE c VOLATILE STRICT
  COST 1;
ALTER FUNCTION "public"."pgp_sym_encrypt_bytea"(bytea, text) OWNER TO "postgres";

-- ----------------------------
-- Uniques structure for table admins
-- ----------------------------
ALTER TABLE "public"."admins" ADD CONSTRAINT "admins_email_key" UNIQUE ("email");

-- ----------------------------
-- Primary Key structure for table admins
-- ----------------------------
ALTER TABLE "public"."admins" ADD CONSTRAINT "admins_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table contacts
-- ----------------------------
ALTER TABLE "public"."contacts" ADD CONSTRAINT "contacts_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table projects
-- ----------------------------
CREATE INDEX "idx_projects_slug" ON "public"."projects" USING btree (
  "slug" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Uniques structure for table projects
-- ----------------------------
ALTER TABLE "public"."projects" ADD CONSTRAINT "projects_slug_key" UNIQUE ("slug");

-- ----------------------------
-- Primary Key structure for table projects
-- ----------------------------
ALTER TABLE "public"."projects" ADD CONSTRAINT "projects_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table seo_metadata
-- ----------------------------
CREATE INDEX "idx_seo_metadata_reference" ON "public"."seo_metadata" USING btree (
  "reference_type" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "reference_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);

-- ----------------------------
-- Uniques structure for table seo_metadata
-- ----------------------------
ALTER TABLE "public"."seo_metadata" ADD CONSTRAINT "seo_metadata_reference_id_reference_type_key" UNIQUE ("reference_id", "reference_type");

-- ----------------------------
-- Primary Key structure for table seo_metadata
-- ----------------------------
ALTER TABLE "public"."seo_metadata" ADD CONSTRAINT "seo_metadata_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table services
-- ----------------------------
ALTER TABLE "public"."services" ADD CONSTRAINT "services_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table site_settings
-- ----------------------------
ALTER TABLE "public"."site_settings" ADD CONSTRAINT "site_settings_pkey" PRIMARY KEY ("key");

-- ----------------------------
-- Primary Key structure for table skills
-- ----------------------------
ALTER TABLE "public"."skills" ADD CONSTRAINT "skills_pkey" PRIMARY KEY ("id");
