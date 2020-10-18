-- bulletins
ALTER TABLE bulletin DROP CONSTRAINT bulletin_user_id_fkey;
ALTER TABLE bulletin ADD CONSTRAINT bulletin_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.app_user(id) ON DELETE CASCADE;

-- contact methods
ALTER TABLE contact_method DROP CONSTRAINT contact_method_user_id_fkey;
ALTER TABLE contact_method ADD CONSTRAINT contact_method_user_id_fkey FOREIGN KEY (user_id) REFERENCES app_user(id) ON DELETE CASCADE;

-- match groups
ALTER TABLE match_group_app_user DROP CONSTRAINT match_group_app_user_user_id_fkey;
ALTER TABLE match_group_app_user ADD CONSTRAINT match_group_app_user_user_id_fkey FOREIGN KEY (user_id) REFERENCES app_user(id) ON DELETE CASCADE;