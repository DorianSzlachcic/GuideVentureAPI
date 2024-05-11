package queries

const Create = `CREATE TABLE IF NOT EXISTS games (
					id INTEGER NOT NULL PRIMARY KEY,
					title TEXT NOT NULL,
					introduction TEXT
					);
				CREATE TABLE IF NOT EXISTS steps (
					id INTEGER NOT NULL PRIMARY KEY,
					game_id INTEGER,
					step_index INTEGER NOT NULL,
					step_type TEXT CHECK(step_type IN ('navigate', 'puzzle', 'quiz', 'photo')) NOT NULL DEFAULT 'navigate',
					points INTEGER NOT NULL DEFAULT 100,
					description TEXT,
					latitude REAL,
					longitude REAL,
					image_source TEXT,
					FOREIGN KEY(game_id) REFERENCES games(id)
					);
				CREATE TABLE IF NOT EXISTS questions (
					id INTEGER NOT NULL PRIMARY KEY,
					step_id INTEGER,
					question_type TEXT CHECK(question_type IN ('multiple_choice', 'single_choice', 'text_answer')) NOT NULL DEFAULT 'single_choice',
					question TEXT NOT NULL,
					answers TEXT NOT NULL,  -- divided by ';', first considered as correct answer
					num_of_correct_ans INTEGER CHECK(num_of_correct_ans >= 1) NOT NULL DEFAULT 1,
					FOREIGN KEY(step_id) REFERENCES steps(id)
				);`

const DummyData = `INSERT INTO games VALUES(1, 'Przykładowa gra', 'Gra przykładowa stworzona na potrzebny developmentu');
					INSERT INTO steps VALUES(1, 1, 1, 'navigate', 50, 'Przejdź do punktu startowego', 51.109730, 17.030655, NULL);
					INSERT INTO steps VALUES(2, 1, 2, 'quiz', 300, 'Odpowiedz na poniższe pytania', NULL, NULL, NULL);
					INSERT INTO questions VALUES(1, 2, 'single_choice', 'Gdzie się znajdujesz?', 'Wrocław;Gdańsk;Poznań;Inne', 1);
					INSERT INTO questions VALUES(2, 2, 'text_answer', 'Jaki jest najlepszy uniwersytet we Wrocławiu?', 'UWR', 1);
					INSERT INTO questions VALUES(3, 2, 'multiple_choice', 'Jakie są największe atrakcje Wrocławia?', 'Rynek;WFiA;Wyspa Słodowa;Nadodrze', 2);
					INSERT INTO steps VALUES(3, 1, 3, 'navigate', 50, 'Przejdź do następnego punktu', 51.114730, 17.042187, NULL);
					INSERT INTO steps VALUES(4, 1, 4, 'photo', 100, 'Wykonaj sobie zdjęcie', NULL, NULL, NULL);
					INSERT INTO steps VALUES(5, 1, 5, 'navigate', 50, 'Przejdź do następnego punktu', 51.114046, 17.031246, NULL);
					INSERT INTO steps VALUES(6, 1, 6, 'puzzle', 300, 'Ułóżcie puzzle', NULL, NULL, 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/b6/Wroclaw_-_Uniwersytet_Wroclawski_o_poranku.jpg/1200px-Wroclaw_-_Uniwersytet_Wroclawski_o_poranku.jpg');
					`

const SelectGames = `SELECT * FROM games;`
const SelectGameById = `SELECT * FROM games WHERE id=?;`
