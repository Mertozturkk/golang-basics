


Test yazma surecinde dikkat edilmesi gerekenler temelde bir fonksiyon yazmakdan farksiz.

- `xxx_test.go` gibi bir go dosyasi icinde olmali.
- Test fonksiyonlari isimlendirmeleri `Test` keywordu ile baslamali.
- Test fonksiyonlari sadece bir parametre alir oda `t *testing.T` dir
- `t *testing.T` nin kullanilabilmesi icin "testing" paketinin import edilmesi gerekir. `import "testing"`

### TDD Disiplini

- Test yazmadan kod yazma
- Yazilan test ile birlikte kodun compile edilebilecek duruma gelmesi lazim
- Testi calistir, basarisiz oldugunu gor ve hata mesajinin anlamli olup olmadigini kontrol et
- Testi gececek kadar kod yaz
- Refactor et


### Golang Bilgileri

- Bir degisken tanimlanirken kucuk harfle basliyorsa private acces modifier a sahiptir. Tanimlandigi paket disindan erisilemez.
- Go metotlara veya fonksiyonlara ilettiğiniz değerlerin bir kopyasını(pass by value) gönderir. Bu nedenle degiskenin gercek degerinin degismesi gerektigi durumlarda pointer kullanilmali
- Pointerlar nil olabilirler. Bir fonksiyon bir seye point eden bir deger dondugu zaman nil olup olmadigini kontrol etmek gerekir. Aksi halde runtime panic olur.
- 