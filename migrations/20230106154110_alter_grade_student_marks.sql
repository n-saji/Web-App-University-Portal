-- +goose Up
-- +goose StatementBegin
UPDATE student_marks
SET grade = 'not graded'
WHERE grade = '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
