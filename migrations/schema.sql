CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    category TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    variables TEXT,
    dlc TEXT,
    weight INTEGER DEFAULT 1,
    impact TEXT CHECK(impact IN ('Low', 'Medium', 'High')) DEFAULT 'Low',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_events_category ON events(category);
CREATE INDEX idx_events_weight ON events(weight);
CREATE INDEX idx_events_dlc ON events(dlc);

CREATE TABLE IF NOT EXISTS history (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    event_id INTEGER,
    user_id TEXT,
    rolled_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    chosen_variables TEXT,
    FOREIGN KEY (event_id) REFERENCES events(id)
);
