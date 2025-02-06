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

INSERT INTO course_infos (id, course_name) VALUES
('550e8400-e29b-41d4-a716-446655440000', 'Introduction to Computer Science'),
('123e4567-e89b-12d3-a456-426614174001', 'Data Structures and Algorithms'),
('987f6543-b21c-34d2-c789-123456789abc', 'Database Management Systems'),
('456e7890-c12d-45e6-b789-987654321def', 'Operating Systems'),
('321f6547-d34e-56f7-e890-234567890ghi', 'Artificial Intelligence'),
('789a1234-e45f-67g8-f901-345678901jkl', 'Machine Learning'),
('654b3210-f56g-78h9-g012-456789012mno', 'Software Engineering'),
('234c5678-g78h-90i1-h123-567890123pqr', 'Computer Networks'),
('890d1234-h90i-12j3-i234-678901234stu', 'Web Development'),
('567e8901-i12j-34k5-j345-789012345vwx', 'Cloud Computing');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS course_infos;
-- +goose StatementEnd
