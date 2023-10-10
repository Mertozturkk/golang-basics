


Test yazma surecinde dikkat edilmesi gerekenler temelde bir fonksiyon yazmakdan farksiz.

- `xxx_test.go` gibi bir go dosyasi icinde olmali.
- Test fonksiyonlari isimlendirmeleri `Test` keywordu ile baslamali.
- Test fonksiyonlari sadece bir parametre alir oda `t *testing.T` dir
- `t *testing.T` nin kullanilabilmesi icin "testing" paketinin import edilmesi gerekir. `import "testing"`

