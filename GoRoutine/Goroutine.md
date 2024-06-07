# Concurency
Menjalankan beberapa pekerjaan secara bergatian. Dalam concurreny biasanya membutuhkan hanya sedikit thread.
## Contoh : 
Saat kita makan lalapan di warung, kita bisa makan lalu minum es jeruk, makan lagi, dan seterusnya. Akan tetapi, hal2 tersebut tidak dapat dilakukan secara bersamaan, namun bisa berganti kapanpun kita mau.

Golang cocok untuk pemrograman yang menggunakan IO Bound. 
## IO Bound
IO Bound adalah algoritma yang sangat tergantung dengan kecepatan input output devices yang digunakan. Contohnya aplikasi seperti membaca data dari file, database dan lain-lain.

## Goroutine
Goroutine adalah sebuah thread ringan yang dikelola oleh Go Runtime. Ukuran goroutine sangat kecil sekitar 2kb, sedangkan thread bisa sampai 1mb atau 1000kb. Goroutine berjalan secara konkuren.


