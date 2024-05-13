-- Insert demo data into user table
-- Password: demo#123
INSERT INTO
    "users" (
        is_admin,
        username,
        email,
        password,
        first_name,
        last_name
    )
values
    (
        TRUE,
        'amir_admin',
        'amir@admin.com',
        '$2a$10$AWuR/dHlOdwYY0Vwtez28.thr67ir8LoB964QQr8QS2tX/eYKh8yS',
        'Amir',
        'Rostami'
    ),
    (
        FALSE,
        'amir_demo',
        'amir@demo.com',
        '$2a$10$AWuR/dHlOdwYY0Vwtez28.thr67ir8LoB964QQr8QS2tX/eYKh8yS',
        'Amir',
        'Rostami Demo'
    )