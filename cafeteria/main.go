package main

import (
	"fmt"
)

type Cliente struct {
	ID      int
	Nombre  string
	Carrera string // "TI" o "SO"
	Saldo   float64
}

// Producto representa algo que se vende en la cafetería: bebidas, snacks, almuerzos.
// Tiene un stock — cuando alguien compra, se descuenta del stock disponible.
type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string // "bebida", "snack", "almuerzo"
}

// Pedido representa una transacción: un cliente compra una cantidad de un producto.
// No contiene el cliente ni el producto completos — solo guarda sus IDs.
type Pedido struct {
	ID         int
	ClienteID  int
	ProductoID int
	Cantidad   int
	Total      float64
	Fecha      string // formato libre, ej: "2026-04-16"
}

// =============================================================================
// SECCIÓN 2 — FUNCIONES DE VISUALIZACIÓN
// =============================================================================

// ListarClientes imprime todos los clientes registrados en formato tabla.
func ListarClientes(clientes []Cliente) {
	fmt.Println("\n=== CLIENTES REGISTRADOS ===")
	if len(clientes) == 0 {
		fmt.Println("(no hay clientes registrados)")
		return
	}
	fmt.Println("ID | Nombre               | Carrera | Saldo")
	fmt.Println("---------------------------------------------")
	for _, c := range clientes {
		fmt.Printf("%-2d | %-20s | %-7s | $%.2f\n",
			c.ID, c.Nombre, c.Carrera, c.Saldo)
	}
}

// =============================================================================
// SECCIÓN 3 — MAIN CON DATOS INICIALES
// =============================================================================

func main() {
	// =========================================================================
	// DATOS INICIALES (simulan lo que en el futuro vendrá de una base de datos)
	// =========================================================================

	// Slice de clientes (mínimo 3)
	clientes := []Cliente{
		{ID: 1, Nombre: "Ana López", Carrera: "TI", Saldo: 25.50},
		{ID: 2, Nombre: "Carlos Méndez", Carrera: "SO", Saldo: 15.00},
		{ID: 3, Nombre: "Diana Ruiz", Carrera: "TI", Saldo: 30.00},
		{ID: 4, Nombre: "Esteban Gómez", Carrera: "SO", Saldo: 10.50},
	}

	// Slice de productos (mínimo 4)
	productos := []Producto{
		{ID: 1, Nombre: "Café Americano", Precio: 1.50, Stock: 20, Categoria: "bebida"},
		{ID: 2, Nombre: "Sandwich de Pollo", Precio: 3.50, Stock: 10, Categoria: "snack"},
		{ID: 3, Nombre: "Almuerzo Ejecutivo", Precio: 5.00, Stock: 8, Categoria: "almuerzo"},
		{ID: 4, Nombre: "Jugo Natural", Precio: 1.00, Stock: 15, Categoria: "bebida"},
		{ID: 5, Nombre: "Empanada de Queso", Precio: 1.00, Stock: 25, Categoria: "snack"},
	}

	// Slice de pedidos (vacío inicialmente)
	pedidos := []Pedido{}

	// =========================================================================
	// PRUEBA DEL CHECKPOINT 1
	// =========================================================================
	fmt.Println("══════════════════════════════════════════")
	fmt.Println("  MINI-CAFETERÍA DOÑA ROSA - Checkpoint 1")
	fmt.Println("══════════════════════════════════════════")

	// Probamos la función ListarClientes
	ListarClientes(clientes)

	// También podemos mostrar que los slices existen
	fmt.Printf("\n✓ Datos cargados: %d clientes, %d productos, %d pedidos\n",
		len(clientes), len(productos), len(pedidos))
}
