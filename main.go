package main

func main() {
	todos := Todos{}
	todos.add("Buy milk")
	todos.add("Buy book")
	todos.toggle(0)
	todos.print()
}
