ALTER TABLE auth ADD CONSTRAINT unique_api_key UNIQUE (`api-key`);
ALTER TABLE user_profile ADD CONSTRAINT unique_user_PHONE UNIQUE (`phone`);
