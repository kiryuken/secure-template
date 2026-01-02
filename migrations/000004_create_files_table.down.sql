-- Drop files table
DROP INDEX IF EXISTS idx_files_checksum;
DROP INDEX IF EXISTS idx_files_uploaded_at;
DROP INDEX IF EXISTS idx_files_user_id;
DROP TABLE IF EXISTS files;
