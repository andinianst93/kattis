1. Baca n (jumlah hotel)
2. Baca array d (jarak ke tiap hotel)

3. Buat fungsi ways(s) yang return jumlah cara dapat sum s:
   - Jika s < 2 atau s > 12: return 0
   - Else: return 6 - abs(s - 7)

4. total_ways = 0
5. Untuk setiap jarak di dalam d:
   - total_ways += ways(jarak)

6. probability = total_ways / 36.0

7. Print probability
