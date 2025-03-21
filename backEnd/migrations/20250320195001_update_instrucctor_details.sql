-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF EXISTS instructor_details DROP CONSTRAINT IF EXISTS instructor_details_course_id_fkey;
ALTER TABLE if exists instructor_details
DROP COLUMN IF EXISTS course_id,
DROP COLUMN IF EXISTS info;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE if exists instructor_details
ADD COLUMN if not exists course_id uuid REFERENCES course_infos(id), 
add column if not exists info jsonb default '{}';
-- +goose StatementEnd
