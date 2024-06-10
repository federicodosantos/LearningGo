# Channel
Channel adalah tempat komunikasi secara synchronous(satu arah) yang dilakukan oleh goroutine. Di channel terdapat pengirim dan penerima, biasanya pengirim dan penerima adalah goroutine yang berbeda. Saat melakukan pengiriman data ke channel, goroutine akan ter-block, sampai ada yang menerima data tersebut. Maka dari itu, channel disebut sebagai alat komunikasi synchronous(blocking).

## Karakteristik Channel
Secara default channel hanya bisa menampung satu data, jika kita ingin menambahkan data lagi maka harus menunggu data yang ada di channel itu diambil. Lalu, channel hanya bisa menerima satu jenis data. Channel bisa diambil dari lebih dari satu goroutine. Jadi channel itu dapat dipakai oleh banyak goroutine akan tetapi harus gantian dengan goroutine lainnya. 

## Channel sebagai parameter
Dala, kenyataan pembuatan aplikasi, seringnya kita akan mengirim channel ke function lain via parameter. Sebelumnya kita tahu bahkan di Go-Lang by default, parameter adalah pass by value, artinya value akan diduplikasi lalu dikirim ke function parameter, sehingga jika kita ingin mengirim data asli, kita biasa gunakan pointer (agar pass by reference). Berbeda dengan channel, kita tidak perlu melakukan hal tersebut.

## Buffered Channel
Secara default channel hanya bisa menerima 1 data. Kadang-kadang ada kasus dimana pengirim lebih cepat dibanding penerima, dalam hal ini jika kita menggunakan channel, maka otomatis pengirim ikut lambat juga. Untuk mengatasi hal tersebut, kita gunakan buffered channel, yaitu buffer yang bisa digunakan untuk menampung data antrian di channel.

## Range Channel
Ada kasus dimana sebuah channel mengirim data secara terus-terusan oleh pengirim. Kadang tidak jelas kapan channel tersebut akan berhenti menerima data dari si pengirim. Untuk mengatasi hal itu kita dapat menggunakan perulangan range ketika menerima data dari channel. Ketika sebuah channel di close() maka perualangan range tersebut akan berhenti.

## Select Channel
Ada kasus dimana kita membuat beberapa channel dan menjalankan beberapa goroutine. Lalu kita ingin mendapat data dari semua channel tersebut. Pada kasus tersebut kita dapat menggunakan perintah 'select' channel. Dengan menggunakan select, kita bisa memilih channel tercepat, jika ada beberapa channel yang sama2 tercepat maka akan dipilih salah satu secara random. 