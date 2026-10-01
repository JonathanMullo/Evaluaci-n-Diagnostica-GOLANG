package main

import "fmt"

var productosVendidos []string
var subtotales []float64

func RegistrarVenta(nombre string, precio float64, cantidad int) {
	subtotal := precio * float64(cantidad)

	productosVendidos = append(productosVendidos, nombre)
	subtotales = append(subtotales, subtotal)

	fmt.Println("Venta registrada correctamente.")
	fmt.Println("Producto:", nombre)
	fmt.Println("Subtotal:", subtotal)
}

func MostrarEstadisticas() {
	if len(subtotales) == 0 {
		fmt.Println("No existen ventas registradas")
		return
	}

	total := 0.0

	for _, subtotal := range subtotales {
		total = total + subtotal
	}

	fmt.Println("Total Calculado:", total)
}

func main() {
	for {
		fmt.Println("-- SELECCIONE UNA OPCION --")
		fmt.Println("1._ Registrar una nueva venta")
		fmt.Println("2._ Mostrar Estadisticas")
		fmt.Println("3._ Salir")

		var opcion int
		fmt.Scanf("%d", &opcion)

		switch opcion {

		case 1:
			fmt.Println("PRODUCTOS")
			fmt.Println("1._ Arroz $1.25")
			fmt.Println("2._ Leche $0.95")
			fmt.Println("3._ Pan $0.50")

			var producto int
			fmt.Println("Seleccione el producto:")
			fmt.Scanf("%d", &producto)

			var nombre string
			var precio float64

			switch producto {
			case 1:
				nombre = "Arroz"
				precio = 1.25

			case 2:
				nombre = "Leche"
				precio = 0.95

			case 3:
				nombre = "Pan"
				precio = 0.50

			default:
				fmt.Println("Ese producto no existe")
				continue
			}

			var cantidad int
			fmt.Println("Ingrese la cantidad vendida:")
			fmt.Scanf("%d", &cantidad)

			if cantidad <= 0 {
				fmt.Println("La cantidad debe ser mayor que cero.")
				continue
			}

			RegistrarVenta(nombre, precio, cantidad)

		case 2:
			fmt.Println("ESTADISTICAS")
			MostrarEstadisticas()

		case 3:
			fmt.Println("Programa finalizado.")
			return

		default:
			fmt.Println("Opción no válida.")
		}
	}
}
```























