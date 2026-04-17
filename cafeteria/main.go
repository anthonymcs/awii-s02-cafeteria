package main

//Declaración de entidades

type Cliente struct {
	ID      int
	nombre  string
	Carrera string
	Saldo   float64
}

type Producto struct {
	ID        int
	nombre    string
	Precio    float64
	Stock     int
	Categoria string
}

type Pedido struct {
	ClienteID  int
	ProductoID int
	Cantidad   int
	Total      float64
	Fecha      string
}

func main() {
	// Aqui se crean las listas
	var clientes []Cliente
	var productos []Producto
	var pedidos []Pedido

	// Agregamos 3 cliente y 4 productos al slice
	clientes = append(clientes, Cliente{ID: 1, nombre: "Juan Perez", Carrera: "Ingeniería", Saldo: 100.0})
	clientes = append(clientes, Cliente{ID: 2, nombre: "Maria Gomez", Carrera: "Medicina", Saldo: 150.0})
	clientes = append(clientes, Cliente{ID: 3, nombre: "Carlos Ruiz", Carrera: "Arquitectura", Saldo: 200.0})
	productos = append(productos, Producto{ID: 1, nombre: "Pan", Precio: 1.5, Stock: 10, Categoria: "Alimentos"})
	productos = append(productos, Producto{ID: 2, nombre: "Leche", Precio: 2.0, Stock: 5, Categoria: "Alimentos"})
	productos = append(productos, Producto{ID: 3, nombre: "Huevos", Precio: 3.0, Stock: 8, Categoria: "Alimentos"})
	productos = append(productos, Producto{ID: 4, nombre: "Pan Dulce", Precio: 2.5, Stock: 12, Categoria: "Alimentos"})
}
