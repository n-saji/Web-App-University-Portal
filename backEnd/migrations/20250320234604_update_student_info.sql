-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF EXISTS student_infos
DROP COLUMN IF EXISTS course_id,
DROP COLUMN IF EXISTS marks_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS student_infos
ADD COLUMN IF NOT EXISTS course_id uuid REFERENCES course_infos(id),
ADD COLUMN IF NOT EXISTS marks_id uuid REFERENCES student_marks(id);
-- +goose StatementEnd
