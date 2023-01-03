-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS course_infos (
	id uuid NOT NULL,
	course_name text NULL,
	CONSTRAINT course_infos_pkey PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF NOT EXISTS course_infos;
-- +goose StatementEnd
