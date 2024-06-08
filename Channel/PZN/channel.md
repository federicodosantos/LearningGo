# Channel
Channel adalah tempat komunikasi secara synchronous(satu arah) yang dilakukan oleh goroutine. Di channel terdapat pengirim dan penerima, biasanya pengirim dan penerima adalah goroutine yang berbeda. Saat melakukan pengiriman data ke channel, goroutine akan ter-block, sampai ada yang menerima data tersebut. Maka dari itu, channel disebut sebagai alat komunikasi synchronous(blocking).

## Karakteristik Channel
Secara default channel hanya bisa menampung satu data, jika kita ingin menambahkan data lagi maka harus menunggu data yang ada di channel itu diambil. Lalu, channel hanya bisa menerima satu jenis data. Channel bisa diambil dari lebih dari satu goroutine. Jadi channel itu dapat dipakai oleh banyak goroutine akan tetapi harus gantian dengan goroutine lainnya. 

## Channel sebagai parameter
Dala, kenyataan pembuatan aplikasi, seringnya kita akan mengirim channel ke function lain via parameter. Sebelumnya kita tahu bahkan di Go-Lang by default, parameter adalah pass by value, artinya value akan diduplikasi lalu dikirim ke function parameter, sehingga jika kita ingin mengirim data asli, kita biasa gunakan pointer (agar pass by reference). Berbeda dengan channel, kita tidak perlu melakukan hal tersebut.