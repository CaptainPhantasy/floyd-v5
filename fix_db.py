#!/usr/bin/env python3
import sqlite3
import os

db_path = os.path.expanduser("~/.floyd/floyd.db")
print(f"Checking database: {db_path}")

conn = sqlite3.connect(db_path)
cursor = conn.cursor()

# Check if cache_read_tokens column exists
cursor.execute("PRAGMA table_info(sessions)")
columns = [row[1] for row in cursor.fetchall()]
print("Current columns:", columns)

if 'cache_read_tokens' not in columns:
    print("Adding cache_read_tokens column...")
    try:
        cursor.execute("""
            ALTER TABLE sessions 
            ADD COLUMN cache_read_tokens INTEGER NOT NULL DEFAULT 0
        """)
        conn.commit()
        print("Column added successfully.")
    except sqlite3.Error as e:
        print(f"Error adding column: {e}")
        # Check if column already exists (maybe case sensitivity)
        cursor.execute("""
            SELECT COUNT(*) FROM pragma_table_info('sessions') 
            WHERE LOWER(name) = 'cache_read_tokens'
        """)
        if cursor.fetchone()[0] > 0:
            print("Column exists with different case.")
else:
    print("cache_read_tokens column already exists.")

# Verify
cursor.execute("PRAGMA table_info(sessions)")
for row in cursor.fetchall():
    print(f"{row[0]}: {row[1]} ({row[2]})")

conn.close()
print("Done.")