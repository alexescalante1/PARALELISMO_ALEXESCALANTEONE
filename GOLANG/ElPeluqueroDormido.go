package main
import (
	"fmt"
	"sync"
	"time"
)

type cliente struct {
	cabello int
	sync.Mutex
}

type barbero struct {
	durmiendo bool
}

func (p cliente) atender(c chan *cliente, wg *sync.WaitGroup, ncliente int) {//
	
	time.Sleep(1*time.Second)

	p.Lock()
		
		fmt.Println("Cortando al cliente: ",ncliente)
		
		for i := p.cabello ; i >= 0 ; i--{
			fmt.Print(" ", i)
			p.cabello = i
		}
					
	p.Unlock()
	c <- &p
	
	wg.Done()
}

func host(c chan *cliente, wg *sync.WaitGroup) {
	for {
		if len(c) == 1 {
			<-c
			time.Sleep(20 * time.Millisecond)
		}
	}
}

func main() {
	
	AFOROLOCAL := 5
	NUMEROCLIENTES := 8

	c := make(chan *cliente,1)
	var wg sync.WaitGroup
	
	Clientes := make([]*cliente, NUMEROCLIENTES)

	for i := 0; i < NUMEROCLIENTES; i++ {
		Clientes[i] = new(cliente)
		Clientes[i].cabello = 5 + i
	}

	for i := 0; i < AFOROLOCAL; i++ {
		fmt.Println("Cliente: ", i ," Entra a la peluqueria, cantidad de cabello: ", Clientes[i].cabello)
	}

	for i := AFOROLOCAL; i < NUMEROCLIENTES; i++ {
		fmt.Println("Cliente: ", i ," Se va porque ya esta lleno")
	}

	fmt.Println()
	wg.Add(AFOROLOCAL) //CANTIDAD DE SILLAS DE ESPERA
		
		fmt.Println("Despierta el Barbero\n")

		go host(c, &wg)
		for i := 0; i < AFOROLOCAL &&  i < NUMEROCLIENTES; i++ {

			start := time.Now()
			go Clientes[i].atender(c, &wg, i)
			ATENDIDO := <- c
			t := time.Now()
			TiempoDeEjecucion := t.Sub(start)
			if ATENDIDO.cabello == 0{
				
				fmt.Println("\nTermino de cortar!\nTiempo: ",TiempoDeEjecucion)
				fmt.Println("-SIGUIENTE...\n")
				
			}else {
				fmt.Println("ERROR NO SE TERMINO DE CORTAR...\n")
				i--
			}
			
		}
		fmt.Println("-NO HAY MAS CLIENTES...\n")

		fmt.Println("El Barbero se duerme")

	wg.Wait()

}
