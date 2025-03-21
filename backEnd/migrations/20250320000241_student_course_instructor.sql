-- +goose Up
-- +goose StatementBegin
CREATE TABLE student_course_instructors(
    student_id uuid NOT NULL,
    course_id uuid NOT NULL,
    instructor_id uuid NOT NULL,
    marks INT NULL,
    is_deleted BOOLEAN DEFAULT FALSE,

    PRIMARY KEY (student_id, course_id, instructor_id),
    FOREIGN KEY (student_id) REFERENCES student_infos (id),
    FOREIGN KEY (course_id) REFERENCES course_infos (id),
    FOREIGN KEY (instructor_id) REFERENCES instructor_details (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS student_course_instructors ;
-- +goose StatementEnd
