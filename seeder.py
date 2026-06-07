import psycopg2
import csv
import os

# 1. Konfigurasi Koneksi Database (Docker Port: 28711)
DB_CONFIG = {
    "dbname": "sims4_db",
    "user": "postgres",
    "password": "supersecretpassword",
    "host": "localhost",
    "port": "28712"
}

def load_csv(filename):
    filepath = os.path.join(os.path.dirname(__file__), 'data', filename)
    with open(filepath, 'r', encoding='utf-8') as f:
        reader = csv.reader(f)
        next(reader) # skip header
        return [tuple(row) for row in reader]

def seed_database():
    conn = None
    cursor = None
    try:
        conn = psycopg2.connect(**DB_CONFIG)
        cursor = conn.cursor()
        print("Successfully connected to the database!")

        # 🏗️ LANGKAH 1: BUAT TABEL OTOMATIS JIKA BELUM ADA
        print("🏗️ Creating tables if they don't exist...")
        
        cursor.execute("""
            CREATE TABLE IF NOT EXISTS traits (
                id SERIAL PRIMARY KEY,
                name VARCHAR(100) NOT NULL,
                category VARCHAR(100),
                generated_mood VARCHAR(100)
            );
        """)

        cursor.execute("""
            CREATE TABLE IF NOT EXISTS careers (
                id SERIAL PRIMARY KEY,
                name VARCHAR(100) NOT NULL,
                branch VARCHAR(100),
                base_salary INT,
                ideal_mood VARCHAR(100)
            );
        """)

        cursor.execute("""
            CREATE TABLE IF NOT EXISTS career_recommendations (
                career_id INT REFERENCES careers(id) ON DELETE CASCADE,
                trait_id INT REFERENCES traits(id) ON DELETE CASCADE,
                compatibility_score INT CHECK (compatibility_score BETWEEN 1 AND 5),
                reason TEXT,
                PRIMARY KEY (career_id, trait_id)
            );
        """)
        print("✅ Tables created/verified successfully!")

        # 🔄 LANGKAH 2: BERSIHKAN DATA LAMA AGAR TIDAK DUPLIKAT
        cursor.execute("TRUNCATE TABLE career_recommendations, careers, traits RESTART IDENTITY CASCADE;")

        # 🌱 LANGKAH 3: SUNTIK DATA MENTAH DARI FILE CSV (ETL)
        traits_data = load_csv('traits.csv')
        for trait in traits_data:
            cursor.execute("INSERT INTO traits (name, category, generated_mood) VALUES (%s, %s, %s);", trait)
        print("🌱 Traits data seeded successfully!")

        careers_data = load_csv('careers.csv')
        for career in careers_data:
            cursor.execute("INSERT INTO careers (name, branch, base_salary, ideal_mood) VALUES (%s, %s, %s, %s);", career)
        print("🌱 Careers data seeded successfully!")

        recommendations_data = load_csv('recommendations.csv')
        for rec in recommendations_data:
            cursor.execute("""
                INSERT INTO career_recommendations (career_id, trait_id, compatibility_score, reason) 
                VALUES (%s, %s, %s, %s);
            """, rec)
        print("🌱 Compatibility recommendations data seeded successfully!")

        conn.commit()
        print("🎉 All tables and data have been successfully committed to Docker PostgreSQL!")

    except Exception as e:
        print(f"❌ Error occurred: {e}")
        if conn: conn.rollback()
    finally:
        if cursor: cursor.close()
        if conn: conn.close()

if __name__ == "__main__":
    seed_database()