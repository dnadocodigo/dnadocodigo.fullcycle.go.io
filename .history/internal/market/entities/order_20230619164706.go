package entity

type Order struct{
	ID: type SortBy []Type
	
	func (a SortBy) Len() int           { return len(a) }
	func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
	func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }
	Investor: *Investor

}