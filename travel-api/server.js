const express = require('express');
const cors = require('cors');

const app = express();
app.use(cors());

// Mock Data untuk Paket Liburan The Sims 4 (Akurat sesuai Expansion/Game Packs)
const travelPackages = [
    {
        destination: "Selvadorada (Jungle Adventure)",
        description: "Eksplorasi kuil kuno Omiscan dan kumpulkan artefak langka. Awas jebakan racun!",
        price: 750, // Harga sewa per malam rata-rata
        ideal_mood: "Focused",
        recommended_for_salary: 300 // Cocok untuk kelas menengah ke atas
    },
    {
        destination: "Granite Falls (Outdoor Retreat)",
        description: "Kemping di hutan pinus, memanggang marshmallow, dan mencari serangga langka.",
        price: 250,
        ideal_mood: "Inspired",
        recommended_for_salary: 100 // Cocok untuk semua kalangan
    },
    {
        destination: "Mt. Komorebi (Snowy Escape)",
        description: "Bermain ski, snowboarding, memanjat tebing es, dan relaksasi di Onsen Bathhouse.",
        price: 1200,
        ideal_mood: "Energized",
        recommended_for_salary: 400 // Butuh dana besar
    },
    {
        destination: "Sulani (Island Living)",
        description: "Liburan tropis berjemur di pantai pasir putih, menyelam, dan bertemu putri duyung.",
        price: 900,
        ideal_mood: "Confident",
        recommended_for_salary: 350
    },
    {
        destination: "Tartosa (My Wedding Stories)",
        description: "Destinasi romantis ala Mediterania, sempurna untuk merayakan cinta dan bulan madu.",
        price: 1500,
        ideal_mood: "Flirty",
        recommended_for_salary: 450
    }
];

app.get('/api/travel-packages', (req, res) => {
    res.json({
        success: true,
        data: travelPackages
    });
});

// Endpoint untuk kesehatan service
app.get('/health', (req, res) => {
    res.json({ status: 'OK' });
});

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`Travel API is running on port ${PORT}`);
});
