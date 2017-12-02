Bahasa

STEP 1
Install mariadb atau mysql bagi yang belum ada
1. Download untuk windows  https://downloads.mariadb.org/mariadb/10.2.10/#os_group=windows
   (OS lain silahkan menyesuaikan pilihannya pada halamna web tersebut)
2. Install mariaDB
   Untuk windows,(Next and next)
   Silahkan ikuti langkah di https://mariadb.com/kb/en/library/installing-mariadb-msi-packages-on-windows/
   Ingat user dan password nya ya :)

STEP 2
Install HeidiSQL untuk akses database
1. Download untuk windows https://www.heidisql.com/download.php
2. Install HeidiSQL (next and next)
3. Konfigurasi akses ke database mariadb yang sudah diinstall pada step 1
Langkah 2 dan 3, bisa lihat di link https://www.youtube.com/watch?v=qo0XTfq52cU

STEP 3
Import Tabel
1. Cretae database (mis: cruddb)
2. import tabel dan data awal sesuai dengan SQL berikut
   --------------------
   CREATE TABLE IF NOT EXISTS `user` (
	  `ID` int(11) NOT NULL,
	  `Firstname` varchar(50) DEFAULT NULL,
	  `Lastname` varchar(50) DEFAULT NULL,
	  `City` varchar(50) DEFAULT NULL,
	  `Country` varchar(50) DEFAULT NULL,
	  PRIMARY KEY (`ID`)
	) ENGINE=InnoDB DEFAULT CHARSET=latin1;


	INSERT INTO `user` (`ID`, `Firstname`, `Lastname`, `City`, `Country`) VALUES
		(1, 'Leo', 'Raja', 'Bandung', 'Indonesia'),
		(2, 'Dedi', 'Alan', 'Jakarta', 'Indonesia');
	-----------------------
langkah ini bisa dilihat di link  https://www.youtube.com/watch?v=xBbqDLXrZGY
untuk bagian import ada di meni 2:05 (kurang lebih)

STEP 4
get library  from
1. buka command prompt
2. mysql : jalankan script ini "go get -u github.com/go-sql-driver/mysql"
3. mux   : go get -u github.com/gorilla/mux

STEP 5
Jalankan code dengan perintah "go run crud-db.go"

STEP 6
Silahkan coba lakukan testing menggunakan tool yang disukai,
saya sarankan gunakan REST CLIENT yang ada di bowser
Sesuaikan dengan Selera guys

Moziella : https://addons.mozilla.org/id/firefox/addon/restclient/
Chrome   : https://chrome.google.com/webstore/detail/advanced-rest-client/hgmloofddffdnphfgcellkdfbfbjeloo

Note :
Jika ketemu error, anggaplah sebagai tantangan, hehehe
Enjoy ya guys

