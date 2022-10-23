package main

/*
Clase 2 - TM
Ejercicio 1 - Test Unitario GetAll()
Generar un Stub del Store cuya función “Read” retorne dos productos con las especificaciones que deseen.
Comprobar que GetAll() retorne la información exactamente igual a la esperada. Para esto:
Dentro de la carpeta /internal/(producto/usuario/transacción), crear un archivo repository_test.go
con el test diseñado.

Ejercicio 2 - Test Unitario UpdateName()
Diseñar Test de UpdateName, donde se valide que la respuesta retornada sea correcta para la actualización
del nombre de un producto/usuario/transacción específico. Y además se compruebe que efectivamente se usa
el método “Read” del Storage para buscar el producto. Para esto:
Crear un mock de Storage, dicho mock debe contener en su data un producto/usuario/transacción específico
cuyo nombre puede ser “Before Update”.
El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado.
Puede ser a través de un boolean como se observó en la clase.
Para dar el test como OK debe validarse que al invocar el método del Repository UpdateName, con el
id del producto/usuario/transacción mockeado y con el nuevo nombre “After Update”, efectivamente haga
la actualización. También debe validarse que el método Read haya sido ejecutado durante el test.

Clase 2 - TT
Ejercicio 1 - Service/Repo/Db Update()
Diseñar un test que pruebe en la capa service, el método o función Update(). Para lograrlo se deberá:
1. Crear un mock de Storage, dicho mock debe contener en su data un producto con las especificaciones que desee.
2. El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado.
3. Para dar el test como OK debe validarse que al invocar el método del Service Update(),  retorne el producto con
mismo Id y los datos actualizados. Validar también que  Read() del Store haya sido ejecutado durante el test.

Ejercicio 2 - Service/Repo/Db Delete()
Diseñar un test que pruebe en la capa service, el método o función Delete(). Se debe probar la correcta eliminación
de un producto, y el error cuando el producto no existe. Para lograrlo puede:
1. Crear un mock de Storage, dicho mock debe contener en su data un producto con las especificaciones que desee.
2. Ejecutar el test con dos id’s de producto distintos, siendo uno de ellos un id inexistente en el Mock de Storage.
3. Para dar el test como OK debe validarse que efectivamente el producto borrado ya no exista en Storage luego del Delete().
También que cuando se intenta borrar un producto  inexistente, se debe obtener el error correspondiente.

*/

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/goTesting/goWebTesting/cmd/server/handler"
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/goTesting/goWebTesting/docs"
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/goTesting/goWebTesting/internal/transactions"
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/goTesting/goWebTesting/pkg/store"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	dbFilePath = "./transactions.json"
)

// @title MELI Bootcamp Go W6 - API
// @version 4.2
// @description This API handle transactions
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error: file .env is broken or does not exist")
	}

	db := store.New(store.FileType, dbFilePath)
	repo := transactions.NewRepository(db) //
	service := transactions.NewService(repo)

	t := handler.NewTransaction(service)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	transactionsR := router.Group("/transactions")
	{
		transactionsR.GET("/", t.GetAll())
		transactionsR.POST("/", t.Store())
		transactionsR.PUT("/:ID", t.Update())
		transactionsR.DELETE("/:ID", t.Delete())
		transactionsR.PATCH("/:ID", t.UpdateCodnAmount())
		//transactionsR.GET("?Moneda=USD", FilterTransactions)
		//transactionsR.GET("/:ID", FindTransaction)

	}

	if err := router.Run(); err != nil {
		log.Panic()
	}

}
