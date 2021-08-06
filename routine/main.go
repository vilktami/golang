func main() {
	begin := time.Now()
	for i:= 0;i < 10;i++ {
		go lazy(i)
	}
	fmt.Println(time.Since(begin))
}

func lazy(i int) {
	time.Sleep(100*time.Millisecond)
	fmt.Println(i)
	
}