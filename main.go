package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//SLICE PARA EL REGISTRO DEL TIEMPO
var horasDeIngreso []string

//----------------------------------- FUNCIONES PARA GENERAR BICICLETAS ALEATORIAMENTE ------------------------------------------------
func numAleatorio(maximo int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) //Semilla inicial
	maximo = maximo + 1
	valor := r.Intn(maximo)
	return valor
}

func newPropietario(nombresHombre [79]string, nombresMujer [74]string) string {
	var nombre string
	sexo := numAleatorio(1)
	if sexo == 0 {
		nombre = nombresHombre[numAleatorio(78)]
	} else if sexo == 1 {
		nombre = nombresMujer[numAleatorio(73)]
	}
	return nombre
}

func newEmail(nombre string) string {
	email := nombre + strconv.Itoa(numAleatorio(99)) + "@gmail.com"
	return email
}

func newDireccion(ciudades [15]string) string {
	return ciudades[numAleatorio(14)] + " CALLE " + strconv.Itoa(numAleatorio(99)) + " No. " + strconv.Itoa(numAleatorio(99)) + "-" + strconv.Itoa(numAleatorio(99))
}

func newColor(Ncolores [10]string) string {
	return Ncolores[numAleatorio(9)]
}

func newMarca(marcas [10]string) string {
	return marcas[numAleatorio(9)]
}

func nuevaAleatoria(nombresHombre [79]string, nombresMujer [74]string, ciudades [15]string, Ncolores [10]string, marcas [10]string) *Bicicleta {
	propietario := newPropietario(nombresHombre, nombresMujer)
	email := newEmail(propietario)
	direccion := newDireccion(ciudades)
	edad := numAleatorio(99)
	telefono := numAleatorio(999999999)
	color := newColor(Ncolores)
	marca := newMarca(marcas)
	serial := strconv.Itoa(numAleatorio(1000000000))

	var dueno *Propietario = New_Propietario(propietario, email, direccion, edad, telefono)
	var bici *Bicicleta = New_Bici(*dueno, color, marca, serial)

	return bici
}

//---------------------------------------------------- PROPIETARIO --------------------------------------------------------------------
type Propietario struct {
	Nombre, Email, Direccion string
	Edad, Telefono           int
}

func New_Propietario(Nombre string, Email string, Direccion string, Edad int, Telefono int) *Propietario {
	return &Propietario{Nombre: Nombre, Email: Email, Direccion: Direccion, Edad: Edad, Telefono: Telefono}
}

//---------------------------------------------------- BICICLETA ----------------------------------------------------------------------
type Bicicleta struct {
	Propietario
	Color, Marca, Serial string
}

func New_Bici(propietario Propietario, Color string, Marca string, Serial string) *Bicicleta {
	return &Bicicleta{Propietario: propietario, Color: Color, Marca: Marca, Serial: Serial}
}

func CrearBici() *Bicicleta {
	var propietario string
	var email string
	var direccion string
	var color string
	var marca string
	var serial string
	var edad int
	var telefono int

	fmt.Print("\n\n\nREGISTRO ENTRADA DE BICICLETAS\n\n")
	fmt.Println("DATOS PROPIETARIO:")
	fmt.Print("Nombre del propietario: ")
	fmt.Scanf("%s", &propietario)
	fmt.Print("Correo electrónico: ")
	fmt.Scanf("%s", &email)
	fmt.Print("Dirección de residencia: ")
	fmt.Scanf("%s", &direccion)
	fmt.Print("Edad: ")
	fmt.Scanf("%d", &edad)
	fmt.Print("Número de celular: ")
	fmt.Scanf("%d", &telefono)

	fmt.Print("\nDATOS BICICLETA:\n")
	fmt.Print("Color: ")
	fmt.Scanf("%s", &color)
	fmt.Print("Marca: ")
	fmt.Scanf("%s", &marca)
	fmt.Print("Serial: ")
	fmt.Scanf("%s", &serial)

	var dueno *Propietario = New_Propietario(propietario, email, direccion, edad, telefono)
	var bici *Bicicleta = New_Bici(*dueno, color, marca, serial)

	return bici
}

//-------------------------------------------------- LISTA ENLAZADA -------------------------------------------------------------------
type Nodo struct {
	siguiente *Nodo
	bici      *Bicicleta
}

type Lista struct {
	primero  *Nodo
	ultimo   *Nodo
	Contador int
}

func New_Nodo(bici *Bicicleta) *Nodo {
	return &Nodo{siguiente: nil, bici: bici}
}

func New_Lista() *Lista {
	return &Lista{nil, nil, 0}
}

func Insertar(bici *Bicicleta, lista *Lista) {
	var nuevo *Nodo = New_Nodo(bici)
	if lista.primero == nil {
		lista.primero = nuevo
		lista.ultimo = nuevo
		lista.Contador += 1
	} else {
		lista.ultimo.siguiente = nuevo
		lista.ultimo = lista.ultimo.siguiente
		lista.Contador += 1
	}

	hora := time.Now()
	ingreso := hora.Format("02/01/2006") + " - " + hora.Format("15:04")
	horasDeIngreso = append(horasDeIngreso, ingreso)

	fmt.Println("\nBicicleta de " + bici.Propietario.Nombre + " ingresada correctamente al parqueadero el " + ingreso)
}

func eliminarBici(lista *Lista, name string) {
	var aux_borrar *Nodo
	aux_borrar = lista.primero
	var anterior *Nodo
	anterior = nil

	for (aux_borrar != nil) && (aux_borrar.bici.Propietario.Nombre != name) {
		anterior = aux_borrar
		aux_borrar = aux_borrar.siguiente
	}

	if aux_borrar == nil { //Si se recorrió la lista y no se encontró el elemento
		fmt.Println("No hay bicicletas en el parqueadero a nombre de " + name + "\n")
	} else if anterior == nil { //Si se encuentra en la primera posición
		lista.primero = lista.primero.siguiente
		lista.Contador -= 1
		fmt.Println("Bicicleta de " + name + " retirada correctamente del parqueadero\n")
	} else { //Si está en cualquier otra posición
		anterior.siguiente = aux_borrar.siguiente
		lista.Contador -= 1
		fmt.Println("Bicicleta de " + name + " retirada correctamente del parqueadero\n")
	}
	horasDeIngreso = append(horasDeIngreso[:len(horasDeIngreso)-1])

}

func espaciar(tamanio int, valor int) string {
	espacio := 0
	texto := ""
	espacio = valor - tamanio
	for i := 0; i < espacio; i++ {
		texto = texto + " "
	}
	return texto
}

func Mostrar(lista *Lista) {

	if lista.Contador > 0 {
		actual := lista.primero

		fmt.Println("\n\n\nREGISTRO DE BICICLETAS")

		posHora := -1
		for actual != nil {
			auxColor := actual.bici.Color
			auxMarca := actual.bici.Marca
			auxSerial := actual.bici.Serial
			auxNombre := actual.bici.Propietario.Nombre
			auxEmail := actual.bici.Propietario.Email
			auxDireccion := actual.bici.Propietario.Direccion
			auxEdad := actual.bici.Propietario.Edad
			auxTelefono := actual.bici.Propietario.Telefono

			posHora++

			fmt.Println("\n\nBICICLETA                               PROPIETARIO" + espaciar(11, 45) + "FECHA DE INGRESO:")
			fmt.Println("Color: " + auxColor + espaciar(len(auxColor)+7, 40) + "Nombre: " + auxNombre + espaciar(len(auxNombre)+8, 45) + horasDeIngreso[posHora])
			fmt.Println("Marca: "+auxMarca+espaciar(len(auxMarca)+7, 40)+"Edad:", auxEdad)
			fmt.Println("Serial: "+auxSerial+espaciar(len(auxSerial)+8, 40)+"Teléfono:", auxTelefono)
			fmt.Println(espaciar(0, 40) + "Email: " + auxEmail)
			fmt.Println(espaciar(0, 40) + "Dirección: " + auxDireccion)

			actual = actual.siguiente
		}
	} else {
		fmt.Println("\nPARQUEADERO VACÍO")
	}
}

func contar(lista *Lista) {
	fmt.Println("\nNúmero de bicicletas en el parqueadero:", lista.Contador)
}

func buscar(lista *Lista, name string) {
	bandera := false
	posicion := 0
	var actual *Nodo
	actual = lista.primero
	aux := actual.bici.Propietario.Nombre

	if aux == name {
		bandera = true
		posicion = 1
	}

	for (actual != nil) && (aux != name) {

		if !bandera {
			posicion++
		}

		if actual.bici.Propietario.Nombre == name {
			bandera = true
			aux = name
		}

		actual = actual.siguiente
	}

	if bandera {
		fmt.Println("\nLa bicicleta de "+name+" ha sido encontrada en la posición", posicion, "del parqueadero")
	} else {
		fmt.Println("\nNo hay bicicletas en el parqueadero a nombre de " + name)
	}

}

func vaciar(lista *Lista) {
	lista.primero = lista.ultimo.siguiente
	horasDeIngreso = nil
	lista.Contador = 0
	fmt.Println("\nParqueadero vaciado correctamente")
}

//-------------------------------------------------- FUNCIÓN PRINCIPAL -----------------------------------------------------------------
func main() {

	//79 nombres de hombre
	nombresHombre := [79]string{"ANTONIO",
		"MANUEL",
		"JOSE",
		"FRANCISCO",
		"DAVID",
		"JUAN",
		"JAVIER",
		"DANIEL",
		"CARLOS",
		"JESUS",
		"ALEJANDRO",
		"MIGUEL",
		"RAFAEL",
		"PABLO",
		"PEDRO",
		"ANGEL",
		"SERGIO",
		"FERNANDO",
		"JORGE",
		"LUIS",
		"ALBERTO",
		"ALVARO",
		"ADRIAN",
		"DIEGO",
		"RAUL",
		"IVAN",
		"RUBEN",
		"ENRIQUE",
		"OSCAR",
		"RAMON",
		"ANDRES",
		"VICENTE",
		"SANTIAGO",
		"JOAQUIN",
		"VICTOR",
		"MARIO",
		"EDUARDO",
		"ROBERTO",
		"JAIME",
		"MARCOS",
		"IGNACIO",
		"HUGO",
		"ALFONSO",
		"JORDI",
		"RICARDO",
		"SALVADOR",
		"GUILLERMO",
		"GABRIEL",
		"MARC",
		"EMILIO",
		"MOHAMED",
		"GONZALO",
		"JULIO",
		"JULIAN",
		"MARTIN",
		"TOMAS",
		"AGUSTIN",
		"NICOLAS",
		"SAMUEL",
		"ISMAEL",
		"JOAN",
		"CRISTIAN",
		"FELIX",
		"LUCAS",
		"AITOR",
		"HECTOR",
		"IKER",
		"ALEX",
		"JOSEP",
		"SEBASTIAN",
		"MARIANO",
		"CESAR",
		"ALFREDO",
		"DOMINGO",
		"FELIPE",
		"RODRIGO",
		"MATEO",
		"XAVIER",
		"ALBERT"}

	//74 nombres de mujer
	nombresMujer := [74]string{"MARIA",
		"CARMEN",
		"JOSEFA",
		"ISABEL",
		"LAURA",
		"ANA",
		"CRISTINA",
		"MARTA",
		"LUCIA",
		"FRANCISCA",
		"ANTONIA",
		"DOLORES",
		"SARA",
		"PAULA",
		"ELENA",
		"RAQUEL",
		"PILAR",
		"MANUELA",
		"CONCEPCION",
		"MERCEDES",
		"JULIA",
		"BEATRIZ",
		"NURIA",
		"SILVIA",
		"ALBA",
		"IRENE",
		"ROSARIO",
		"JUANA",
		"TERESA",
		"PATRICIA",
		"ENCARNACION",
		"MONTSERRAT",
		"ANDREA",
		"ROCIO",
		"MONICA",
		"ALICIA",
		"ROSA",
		"SONIA",
		"SANDRA",
		"MARINA",
		"ANGELA",
		"SUSANA",
		"NATALIA",
		"YOLANDA",
		"MARGARITA",
		"CLAUDIA",
		"SOFIA",
		"EVA",
		"CARLA",
		"INMACULADA",
		"ESTHER",
		"NOELIA",
		"VERONICA",
		"NEREA",
		"CAROLINA",
		"ANGELES",
		"DANIELA",
		"INES",
		"MIRIAM",
		"LORENA",
		"VICTORIA",
		"AMPARO",
		"MARTINA",
		"ALEJANDRA",
		"LIDIA",
		"CATALINA",
		"CELIA",
		"CONSUELO",
		"FATIMA",
		"OLGA",
		"AINHOA",
		"GLORIA",
		"CLARA",
		"EMILIA"}

	//15 ciudades
	ciudades := [15]string{"MEDELLÍN",
		"BOGOTÁ",
		"CARTAGENA",
		"BARRANQUILLA",
		"SANTA MARTA",
		"MANIZALES",
		"MONTERIA",
		"RIOHACHA",
		"CALI",
		"PEREIRA",
		"SINCELEJO",
		"ARMENIA",
		"BUCARAMANGA",
		"TUNJA",
		"ARAUCA"}

	//10 colores
	Ncolores := [10]string{"AMARILLA",
		"ROJA",
		"NEGRA",
		"BLANCA",
		"VERDE",
		"AZUL",
		"ROSADA",
		"GRIS",
		"MORADA",
		"CAFE"}

	//10 marcas
	marcas := [10]string{"GW",
		"SHIMANU",
		"SPECIALIZED",
		"TREK",
		"SCOTT",
		"CUBE",
		"CERVELO",
		"VITTORIA",
		"CHAOYANG",
		"SUNTOUR"}

	var lista *Lista = New_Lista()

	opcion := 1
	var name string

	for opcion != 7 {
		fmt.Println("\nBIENVENIDO AL PARQUEADERO DE BICICLETAS MR.JULIAN")
		fmt.Println("¿Qué deseas hacer?:")
		fmt.Println("1. Registrar ingreso de bicileta")
		fmt.Println("2. Registrar salida de bicicleta")
		fmt.Println("3. Ver registro de biciletas")
		fmt.Println("4. Número de bicicletas en el parqueadero")
		fmt.Println("5. Buscar bicicleta")
		fmt.Println("6. Vaciar el parqueadero")
		fmt.Println("7. Salir del sistema")
		fmt.Print("Opción: ")
		fmt.Scanf("%d", &opcion)

		switch opcion {

		case 1:
			var opcion2 int
			fmt.Print("\n¿Cómo deseas realizar el registro?:\n1. Aleatoriamente\n2. Manualmente\nOpción: ")
			fmt.Scanf("%d", &opcion2)

			var ingresos int
			fmt.Print("\n¿Cuantos bicicletas deseas ingresar?: ")
			fmt.Scanf("%d", &ingresos)

			for i := 0; i < ingresos; i++ {
				if opcion2 == 1 { //Aleatoriamente
					Insertar(nuevaAleatoria(nombresHombre, nombresMujer, ciudades, Ncolores, marcas), lista)
				} else if opcion2 == 2 { //Manualmente
					Insertar(CrearBici(), lista)
				}
			}
			fmt.Println("")

			break

		case 2:
			if lista.Contador > 0 {
				fmt.Print("\n\nNombre del propietario: ")
				//name, _ = reader.ReadString('\n')
				fmt.Scanf("%s", &name)
				eliminarBici(lista, name)
				fmt.Println("")
			} else {
				fmt.Println("\n\nPARQUEADERO VACÍO")
			}
			break

		case 3:
			Mostrar(lista)
			fmt.Println("")
			break

		case 4:
			contar(lista)
			fmt.Println("")
			break

		case 5:
			if lista.Contador > 0 {
				fmt.Print("\n\nNombre del propietario: ")
				//name, _ = reader.ReadString('\n')
				fmt.Scanf("%s", &name)
				buscar(lista, name)
				fmt.Println("")
			} else {
				fmt.Println("\n\nPARQUEADERO VACÍO")
			}
			break

		case 6:
			if lista.Contador > 0 {
				var opc int
				fmt.Print("\n¿Estás seguro de dejar el parqueadero vacío?\n1. Sí, continuar\n2. No, regresar al menú\nOpción: ")
				fmt.Scanf("%d", &opc)
				if opc == 1 {
					vaciar(lista)
					fmt.Println("")
				}
			} else {
				fmt.Println("\n\nEL PARQUEADERO YA ESTÁ VACÍO")
			}
			break

		case 7:
			break
		}
	}

}
