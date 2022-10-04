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

*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Request struct {
	ID       int     `json:"ID"`
	Codigo   string  `json:"cod" binding:"required"`
	Moneda   string  `json:"currency" binding:"required"`
	Monto    float64 `json:"amount" binding:"required"`
	Emisor   string  `json:"sender" binding:"required"`
	Receptor string  `json:"receiver" binding:"required"`
	Fecha    string  `json:"date" binding:"required"`
}

const (
	filePath = "./transactions.json"
	Token    = "12345"
)

var transactions []Request

func main() {

	fmt.Println(ReadJson(filePath))

	router := gin.Default()

	router.GET("/hello", func(ctx *gin.Context) {
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

	transactionsR := router.Group("/transactions")
	{
		transactionsR.GET("", GetAll)
		transactionsR.GET("?Moneda=USD", FilterTransactions)
		transactionsR.GET("/:ID", FindTransaction)
		transactionsR.POST("/", AddTransaction())
	}

	router.Run()

}

func AddTransaction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != Token {
			ctx.JSON(401, gin.H{"error": "invalid token"})
			return
		}

		var r Request
		if err := ctx.ShouldBindJSON(&r); err != nil {
			ctx.JSON(404, gin.H{"error": fmt.Errorf("field %s is required", err.Error())})
			return
		}
		r.ID = len(transactions) + 1
		transactions = append(transactions, r)
		ctx.JSON(200, gin.H{"transaction added": r})
	}

}

func GetAll(ctx *gin.Context) {

	/*transactions, err := ReadJson(filePath)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}*/
	ctx.JSON(200, gin.H{"content": transactions})

}

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
