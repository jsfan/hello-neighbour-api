BEGIN;

DROP TABLE IF EXISTS bulletin;

DROP TABLE IF EXISTS match_group_question;

DROP TABLE IF EXISTS match_group_app_user;

DROP TABLE IF EXISTS match_group;

DROP TABLE IF EXISTS question;

DROP TABLE IF EXISTS contact_method;

DROP TABLE IF EXISTS app_user;

DROP TYPE IF EXISTS ROLE;
DROP TYPE IF EXISTS GENDER;

DROP TABLE IF EXISTS church;

DROP TYPE IF EXISTS GROUP_SIZE;

DROP EXTENSION IF EXISTS "uuid-ossp";

COMMIT;