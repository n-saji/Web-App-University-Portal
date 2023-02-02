-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF EXISTS public.instructor_logins
ADD COLUMN  IF NOT EXISTS instructor_id uuid;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS public.instructor_logins
DROP COLUMN IF EXISTS instructor_id;
-- +goose StatementEnd
