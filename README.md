# edit_product

Bu proje, Go ile basit bir REST API kullanarak ürünleri listeleyen ve yeni ürün ekleyen bir uygulamayı içerir.

Microsoft Visual Studio Code için gerekli talimatları aşağıda veriyorum.

Projeyi ilk kez çalıştırırken;

1) npm install -g json-server
2) json-server --watch db.json

sırayla bu 2 komutu terminalden çalıştırın ve json server ı kurun. json-server --watch db.json yazdığınızda 

""""" cannot be loaded because running scripts is disabled on this system. For more information, see about_Exec 
ution_Policies at https:/go.microsoft.com/fwlink/?LinkID=135170. """""

hatasını alırsanız Set-ExecutionPolicy RemoteSigned -Scope CurrentUser komutunu çalıştırın gerekli sonucu alacaksınız.

go run main.go komutuyla çalışır.

ÖNEMLİ! Json server ı kurmadan önce bilgisayarınızda NodeJs in son sürümünün bulunduğundan emin olun. 
(NodeJs İndirme Linki) https://nodejs.org/en/download/package-manager
