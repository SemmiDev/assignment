-- TODO: answer here

-- Kalian diminta untuk membuat query untuk menampilkan data murid yang tidak lulus. Murid yang tidak lulus adalah murid yang memiliki nilai akhir dibawah 70 atau jumlah ketidak hadirannya lebih dari 5 kali.

SELECT id, CONCAT(first_name, ' ', last_name) AS student_name, student_class, final_score, absent
FROM reports WHERE final_score < 70 OR absent > 5;
