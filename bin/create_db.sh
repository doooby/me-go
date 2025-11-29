#!/usr/bin/env bash

db_path=var/db.sqlite

if [ -f $db_path ]; then
  echo "file $db_path already exists"
  exit 0
fi

read -r -d '' sql << EOF
CREATE TABLE tasks (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  task TEXT NOT NULL,
  message TEXT,
  start_at TEXT NOT NULL,
  end_at TEXT,
  created_at TEXT NOT NULL,
  updated_at TEXT NOT NULL
);
EOF

echo "$sql"
sqlite3 $db_path <<< "$sql" 
echo -e "\e[32mdone\e[0m"
