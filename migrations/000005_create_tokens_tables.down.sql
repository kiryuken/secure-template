-- Drop tokens tables
DROP INDEX IF EXISTS idx_email_verification_tokens_user_id;
DROP INDEX IF EXISTS idx_email_verification_tokens_token;
DROP INDEX IF EXISTS idx_password_reset_tokens_user_id;
DROP INDEX IF EXISTS idx_password_reset_tokens_token;
DROP TABLE IF EXISTS email_verification_tokens;
DROP TABLE IF EXISTS password_reset_tokens;
