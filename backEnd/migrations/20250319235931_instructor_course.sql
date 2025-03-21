-- +goose Up
-- +goose StatementBegin
CREATE TABLE instructor_courses(
    instructor_id uuid NOT NULL,
    course_id uuid NOT NULL,
    students_limit INT NOT NULL,
    students_enrolled INT NULL,
    course_rating INT NULL,
    is_deleted BOOLEAN DEFAULT FALSE,
    PRIMARY KEY (instructor_id, course_id),
    FOREIGN KEY (instructor_id) REFERENCES instructor_details (id),
    FOREIGN KEY (course_id) REFERENCES course_infos (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS instructor_courses ;
-- +goose StatementEnd
