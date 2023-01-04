-- Sequence and defined type

-- Table Definition
CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);

-- INSERT INTO "expenses" ("id", "title", "amount", "note", "tags") VALUES (1, 'Expense 1', 100, 'Note for expense 1', ARRAY ['tag1','tag2']);