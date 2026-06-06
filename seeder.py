import psycopg2

# 1. Konfigurasi Koneksi Database
DB_CONFIG = {
    "dbname": "sims4_db",
    "user": "postgres",
    "password": "warayufa28",  # ⚠️ GANTI dengan password PostgreSQL laptopmu
    "host": "localhost",
    "port": "28711"
}

# 2. Data Mentah (Mock Data) The Sims 4
TRAITS_DATA = [
    ("Genius", "Mental", "Focused"),
    ("Geek", "Hobby", "Focused"),
    ("Ambitious", "Emotional", "Confident"),
    ("Creative", "Emotional", "Inspired")
]

CAREERS_DATA = [
    ("Tech Guru", "Startup Entrepreneur", 51, "Focused"),
    ("Tech Guru", "Pro Gamer", 43, "Focused"),
    ("Astronaut", "Space Ranger", 114, "Energized"),
    ("Writer", "Author", 55, "Inspired")
]

RECOMMENDATIONS_DATA = [
    (1, 1, 5, "Sifat Genius sangat membantu mempertahankan mood Focused saat coding di Tech Guru."),
    (1, 2, 4, "Sifat Geek memberikan keuntungan tersendiri saat berinteraksi dengan komputer."),
    (2, 2, 5, "Seorang Pro Gamer wajib memiliki sifat Geek untuk performa game maksimal."),
    (4, 4, 5, "Sifat Creative membuat Sim lebih cepat menghasilkan buku berkualitas mahakarya.")
]

def seed_database():
    try:
        conn = psycopg2.connect(**DB_CONFIG)
        cursor = conn.cursor()
        print("Successfully connected to the database!")

        for trait in TRAITS_DATA:
            cursor.execute("INSERT INTO traits (name, category, generated_mood) VALUES (%s, %s, %s);", trait)
        print("🌱 Traits data seeded successfully!")

        for career in CAREERS_DATA:
            cursor.execute("INSERT INTO careers (name, branch, base_salary, ideal_mood) VALUES (%s, %s, %s, %s);", career)
        print("🌱 Careers data seeded successfully!")

        for rec in RECOMMENDATIONS_DATA:
            cursor.execute("""
                INSERT INTO career_trait_recommendations (career_id, trait_id, compatibility_score, reason) 
                VALUES (%s, %s, %s, %s);
            """, rec)
        print("🌱 Compatibility recommendations data seeded successfully!")

        conn.commit()
        print("🎉 All data has been successfully committed to sims4_db!")

    except Exception as e:
        print(f"❌ Error occurred: {e}")
        if conn: conn.rollback()
    finally:
        if cursor: cursor.close()
        if conn: conn.close()

if __name__ == "__main__":
    seed_database()