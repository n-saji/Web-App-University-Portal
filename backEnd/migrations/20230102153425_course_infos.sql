-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS course_infos (
	id uuid NOT NULL,
	course_name text NULL,
	CONSTRAINT course_infos_pkey PRIMARY KEY (id)
);

INSERT INTO public.course_infos
(id, course_name)
VALUES('285aa383-53b5-4e26-bf3a-c76b3245c617', 'No Course');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS course_infos;
-- +goose StatementEnd
