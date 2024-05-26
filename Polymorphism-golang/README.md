# Polymorphism Kavramı

Polymorphism, bir programlama dilinde nesnelerin farklı türlerde olmalarına rağmen aynı arayüzü kullanarak işlevleri çağırabilmesini sağlayan bir özelliktir. Bu kavram, özellikle nesne yönelimli programlamada (OOP) önemli bir yere sahiptir. Polymorphism sayesinde farklı nesneler aynı metodları farklı şekillerde uygulayabilir, bu da kodun esnekliğini ve genişletilebilirliğini artırır.

## Go Dilinde Polymorphism

Go dili, nesne yönelimli programlama özellikleri içermese de, arayüzler (interfaces) aracılığıyla polymorphism sağlar. Arayüzler, belirli metodlara sahip olması gereken bir türün (type) belirlenmesine yardımcı olur. Bu sayede, farklı türlerdeki nesneler aynı arayüzü uygulayarak polymorphic davranış sergileyebilirler.

## Örnek Senaryo: Sepette Toplam Tutar Hesaplama

Bu örnekte, polymorphism kavramını kullanarak üç farklı ürün türü için vergilendirme yaparak bir sepetin toplam tutarını hesaplayacağız. Her ürün türü için farklı bir vergi hesaplama yöntemi uygulanacak, ancak hepsi aynı _**purchasable**_ arayüzünü uygulayacak. Bu arayüz, ürünlerin vergi hesaplama metodunu içerir. Sepete eklenen ürünler, arayüzü uygulayarak kendi vergilendirme metodlarını kullanır ve toplam tutar hesaplanır.

Polymorphism’in Avantajları

	1.	Esneklik: Farklı türdeki nesneler aynı arayüzü kullanarak farklı şekillerde davranabilirler.
	2.	Kodun Genişletilebilirliği: Yeni türler eklemek kolaydır; sadece arayüzü uygulamak yeterlidir.
	3.	Bakım Kolaylığı: Kod daha modüler ve anlaşılır hale gelir, bu da bakım ve güncellemeleri kolaylaştırır.