CREATE TABLE gocrud.chat (
	id BIGINT auto_increment NOT NULL,
	users_ids varchar(100) NOT NULL,
	CONSTRAINT chat_pk PRIMARY KEY (id),
	CONSTRAINT chat_unique UNIQUE KEY (users_ids)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;



-- gocrud.users definition

CREATE TABLE `users` (
  `name` varchar(100) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `Id` bigint NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- gocrud.message definition

CREATE TABLE gocrud.message (
	id BIGINT auto_increment NULL,
	chat_id BIGINT NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;