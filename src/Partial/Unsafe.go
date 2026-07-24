func _UnsafePartial(f func(any) any) any {
	return f(nil)
}
