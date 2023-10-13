


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

