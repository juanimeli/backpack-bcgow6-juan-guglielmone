package main

/*

CLASE 1

TM Ejercicio 2 - Hola {nombre}
Crea dentro de la carpeta go-web un archivo llamado main.go
Crea un servidor web con Gin que te responda un JSON que tenga una clave “message” y diga Hola seguido por tu nombre.
Pegale al endpoint para corroborar que la respuesta sea la correcta.

TM Ejercicio 3 - Listar Entidad
Ya habiendo creado y probado nuestra API que nos saluda, generamos una ruta que devuelve un listado de la temática elegida.
Dentro del “main.go”, crea una estructura según la temática con los campos correspondientes.
Genera un endpoint cuya ruta sea /temática (en plural). Ejemplo: “/productos”
Genera un handler para el endpoint llamado “GetAll”.
Crea una slice de la estructura, luego devuelvela a través de nuestro endpoint.

TT Ejercicio 1 - Filtremos nuestro endpoint
Según la temática elegida, necesitamos agregarles filtros a nuestro endpoint, el mismo se tiene que poder filtrar por todos los campos.
Dentro del handler del endpoint, recibí del contexto los valores a filtrar.
Luego genera la lógica de filtrado de nuestro array.
Devolver por el endpoint el array filtrado.


TT Ejercicio 2 - Get one endpoint
Generar un nuevo endpoint que nos permita traer un solo resultado del array de la temática. Utilizando path parameters el endpoint
debería ser /temática/:id (recuerda que siempre tiene que ser en plural la temática). Una vez recibido el id devuelve la posición correspondiente.
Genera una nueva ruta.
Genera un handler para la ruta creada.
Dentro del handler busca el item que necesitas.
Devuelve el item según el id.
Si no encontraste ningún elemento con ese id devolver como código de respuesta 404.

Clase 2

TM Ejercicio 1 - Crear Entidad
Se debe implementar la funcionalidad para crear la entidad. pasa eso se deben seguir los
siguientes pasos:
1. Crea un endpoint mediante POST el cual reciba la entidad.
2. Se debe tener un array de la entidad en memoria (a nivel global), en el cual se
deberán ir guardando todas las peticiones que se vayan realizando.
3. Al momento de realizar la petición se debe generar el ID. Para generar el ID se debe
buscar el ID del último registro generado, incrementarlo en 1 y asignarlo a nuestro
nuevo registro (sin tener una variable de último ID a nivel global).

TM Ejercicio 2 - Validación de campos
Se debe implementar las validaciones de los campos al momento de enviar la petición, para
eso se deben seguir los siguientes pasos:
1. Se debe validar todos los campos enviados en la petición, todos los campos son
requeridos
2. En caso que algún campo no esté completo se debe retornar un código de error 400
con el mensaje “el campo %s es requerido”.
(En %s debe ir el nombre del campo que no está completo).

TM Ejercicio 3 - Validar Token
Para agregar seguridad a la aplicación se debe enviar la petición con un token, para eso se
deben seguir los siguientes pasos::

1. Al momento de enviar la petición se debe validar que un token sea enviado
2. Se debe validar ese token en nuestro código (el token puede estar hardcodeado).
3. En caso que el token enviado no sea correcto debemos retornar un error 401 y un
mensaje que “no tiene permisos para realizar la petición solicitada”.


TT Ejercicio 2 - Generar paquete server

Se debe separar la estructura del proyecto, como segundo paso se debe generar
el paquete server donde se agregaran las funcionalidades del proyecto que dependan
de paquetes externos y el main del programa.

Dentro del paquete deben estar:
El main del programa.
Se debe importar e inyectar el repositorio, servicio y handler
Se debe implementar el router para los diferentes endpoints

Clase 3
TM - Ejercicio 1 - Generar método PUT

Se solicita implementar una funcionalidad que modifique completamente una entidad. Para
lograrlo, es necesario, seguir los siguientes pasos:
1. Generar un método PUT para modificar la entidad completa
2. Desde el Path enviar el ID de la entidad que se modificará
3. En caso de no existir, retornar un error 404
4. Realizar todas las validaciones (todos los campos son requeridos)

TM - Ejercicio 2 - Generar método DELETE
Es necesario implementar una funcionalidad para eliminar una entidad.
Para lograrlo, es necesario, seguir los siguientes pasos:
1. Generar un método DELETE para eliminar la entidad en base al ID
2. En caso de no existir, retornar un error 404

TM - Ejercicio 3 - Generar método PATCH

Se requiere implementar una funcionalidad que modifique la entidad parcialmente, solo se
deben modificar 2 campos:
- Si se seleccionó Productos, los campos nombre y precio.
- Si se seleccionó Usuarios, los campos apellido y edad.
- Si se seleccionó Transacciones, los campos código de transacción y monto.
.Para lograrlo, es necesario, seguir los siguientes pasos:
1. Generar un método PATCH para modificar la entidad parcialmente, modificando solo 2
campo (a elección)
2. Desde el Path enviar el ID de la entidad que se modificara
3. En caso de no existir, retornar un error 404
4. Realizar las validaciones de los 2 campos a enviar

*/

import (
	"github.com/gin-gonic/gin"
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/goWeb/cmd/server/handler"
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/goWeb/internal/transactions"
)

func main() {

	//fmt.Println(ReadJson(filePath))

	repo := transactions.NewRepository()
	service := transactions.NewService(repo)

	t := handler.NewTransaction(service)

	router := gin.Default()

	/*router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hola Juani!",
		})
	})


	transactionsaux, err := ReadJson(filePath)
	if err != nil {
		fmt.Println("error")
		os.Exit(1)
	}
	transactions = transactionsaux
	*/

	transactionsR := router.Group("/transactions")
	{
		transactionsR.GET("/", t.GetAll())
		transactionsR.POST("/", t.AddTransaction())
		transactionsR.PUT("/:ID", t.Update())
		transactionsR.DELETE("/:ID", t.Delete())
		transactionsR.PATCH("/:ID", t.UpdateCodnAmount())
		//transactionsR.GET("?Moneda=USD", FilterTransactions)
		//transactionsR.GET("/:ID", FindTransaction)
	}

	router.Run()

}

/*
	func FindTransaction(ctx *gin.Context) {

	transactions, err := ReadJson(filePath)
	if err != nil {
		ctx.JSON(404, err)
		return
	}
	idParam, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	for _, transaction := range transactions {
		if transaction.ID == idParam {
			ctx.JSON(200, transaction)
			return
		}
	}
	ctx.JSON(404, gin.H{"error": "id not found"})

}

func FilterTransactions(ctx *gin.Context) {

	transactions, err := ReadJson(filePath)
	var filtrados []*Request
	if err != nil {
		ctx.JSON(404, err)
	}

	for _, t := range transactions {
		if ctx.Query("USD") == t.Moneda {
			filtrados = append(filtrados, &t)
		}
	}
	ctx.JSON(200, filtrados)
}

func ReadJson(filePath string) ([]Request, error) {

	var transactions []Request

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("error reading the file")
	}
	err = json.Unmarshal(data, &transactions)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}
	return transactions, nil

}
*/
