table users
- uuid
- name
- occupation
- email
- password_hash
- avatar_filename
- role
- created_at

CREATE TABLE local_crowdfunding.m_users (
	uuid varchar(100) NULL,
	name varchar(100) NULL,
	occupation varchar(100) NULL,
	email varchar(100) NULL,
	password_hash varchar(100) NULL,
	avatar_filename varchar(100) NULL,
	`role` varchar(100) NULL,
	created_at DATETIME NULL,
	created_by varchar(100) NULL,
	updated_at DATETIME NULL,
	updated_by varchar(100) NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;
