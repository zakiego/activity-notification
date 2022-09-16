# Telegram Bot to Monitor Turning On and Off The Computer

## Instalasi

1. Buat bot Telegram melalui [@BotFather](http://t.me/BotFather)
2. Buat grup baru di Telegram dan masukkan bot yang sudah dibuat ke grup tersebut
3. Salin token bot yang diberikan oleh [@BotFather](http://t.me/BotFather) dan ganti di file [`utils.go`](./utils/utils.go)
4. Salin ID grup yang dibuat dan ganti di file [`utils.go`](./utils/utils.go)
5. Ubah pesan yang akan dikirim saat menyalakan di file [`on/main.go`](./on/main.go)
6. Ubah pesan yang akan dikirim saat mematikan di file [`off/main.go`](./off/main.go)
7. Build dengan menjalankan `./build.sh`

## Menjalankan Notifikasi Saat Menyalakan Komputer

1. Ketik `crontab -e` di terminal
2. Lalu tambahkan di baris terakhir

   ```bash
   @reboot cd ~/Code/Golang/activity-notification && ./notif-on
   ```

   Catatan: Ubah `~/Code/Golang/activity-notification` sesuai dengan lokasi folder ini

## Menjalankan Notifikasi Saat Mematikan Komputer

1. Buat alias di file `~/.bashrc` atau `~/.zshrc` dengan menambahkan baris berikut

   Catatan: Ubah `~/Code/Golang/notif-poweroff` sesuai dengan lokasi folder ini

   ```bash
   # Poweroff notification
   alias po="cd ~/Code/Golang/notif-poweroff && sudo ./notif-off && sudo poweroff"
   ```

2. Ketik `source ~/.bashrc` atau `source ~/.zshrc` di terminal
3. Ketik `po` di terminal setiap kali untuk mematikan komputer
