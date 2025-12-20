## Pemahaman Problem

1. Ada kecepatan angin saat ini w
2. Ada n jalan, masing-masing punya nama dan batas kecepatan angin maksimum yang aman
3. Jika kecepatan angin saat ini (w) ≤ batas aman jalan → jalan "opin" (buka)
4. Jika kecepatan angin saat ini (w) > batas aman jalan → jalan "lokud" (tutup)


## Solusi

```bash
1. Baca w (kecepatan angin saat ini)
2. Baca n (jumlah jalan)

3. Untuk setiap jalan (i dari 1 sampai n):
   - Baca nama jalan dan batas maksimum (max_safe)
   - Jika w <= max_safe:
     - Print "[nama] opin"
   - Else:
     - Print "[nama] lokud"
```
