-- +goose Up
-- +goose StatementBegin
Alter table if exists instructor_details 
add column if not exists students_list jsonb default '{}';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Alter table if exists instructor_details 
drop column if exists students_list;
-- +goose StatementEnd
