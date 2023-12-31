# Устные вопросы

### Вопрос 1

Какой самый эффективный способ конкатенации строк?

```
strings.Builder минимизирует количество аллокаций памяти, что делает его
предпочтительным выбором для высокопроизводительных операций.
```

### Вопрос 2.

Что такое интерфейсы, как они применяются в Go?

```
Интерфейс - это абстракция, которая позволяет определить общий набор методов,
которые должны быть реализованы для разных типов данных. Это позволяет 
разработчикам обрабатывать разные типы данных одним и с использованием
одного и того же кода (полиморфизм)
```

### Вопрос 3.

Чем отличаются RWMutex от Mutex?

```
Mutex полностью блокирует ресурс для чтения и записи, а RWMutex позволяет
одновременный доступ для чтения нескольким потокам выполнения, однако, 
только один поток может иметь доступ на запись в любой момент времени.

Это полезно, когда у вас есть ресурс, к которому вы часто обращаетесь 
для чтения, но редко записываете, и вы хотите оптимизировать производительность.
```

### Вопрос 4.

Чем отличаются буферизированные и не буферизированные каналы?

```
Буферизированный канал записывает данные в буффер и берет из буффера (stack, FIFO).
Это позволяет отправлять данные в канал не дожидаясь их получения.

При отправке данных в небуферизированный канал мы получим deadlock,
если другая горутина не будет читать из этого канала
```

### Вопрос 5.

Какой размер у структуры `struct{}{}?`

```
Размер 0 байт, т.к. она не имеет полей. 

Под капотом структура все равно будет занимать размер равный
размеру указателя на эту структуру, чтобы хранить адрес этой структуры
в памяти (32/64 бита)
```

### Вопрос 6.

Есть ли в Go перегрузка методов или операторов?

```
Нет, в Go нет перегрузки методов или операторов. В этом языке используется 
строгая статическая типизация, поэтому каждый метод и оператор может быть 
вызван или использован только с определёнными типами данных.
```

### Вопрос 7.

В какой последовательности будут выведены элементы `map[int]int`?

Пример:

```
m[0]=1
m[1]=124
m[2]=281
```

```
Элементы будут в случайной последовательности. map это хэш функция и нее
нет индексации элементов
```

### Вопрос 8.

В чем разница make и new?

```
- make возвращает инициализированный тип, готовый к использованию
- new возвращает указатель на тип с его нулевым значением

- make используется как конструктор для slice, map, chan
- new для всех типов данных
```

### Вопрос 9.

Сколько существует способов задать переменную типа slice или map?

```
Есть 3 способа

1. Использование встроенной функции make:
    films := make([]string, 8)  
    person := make(map[string]string)

2. Создание экземпляра с полем нужного типа:
	type films struct {
		parts []int
		info  map[string]string
	}
	harryPotter := films{
		parts: []int{1, 2, 3, 4, 5, 6, 7, 8},
		info:  map[string]string{"trilogy": "Harry Potter"},
	}

3. Использование литерала нужного типа:
    films := []int{1, 2, 3, 4, 5, 6, 7, 8}
```

### Вопрос 10.

Что выведет данная программа и почему?

```
func update(p *int) {
	b := 2
	p = &b
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p)
	fmt.Println(*p)
}
```

```
Функция выведет 1 и 1

- a = 1, p = &2
- разыменовываем p и выводим 1
- передаем в update ссылку на p
- внутри update создаем b и присваиваем 2
- p = &b
- p внутри области видимости функции update p = &b
- выходя из функции update p снова становится равен &a (1)

если бы мы хотели изменить p на 2, нам нужно было бы изменить
значение по указателю *p = b
```

### Вопрос 11.

Что выведет данная программа и почему?

```
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
```

```
Мы получим дедлок, потому что внутри горутины создаем копию wg
Для решения проблемы нужно передать &wg и wg *sync.WaitGroup
либо вообще убрать wg и wg *sync.WaitGroup из параметров
```

### Вопрос 12.

Что выведет данная программа и почему?

```
func main() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n)
}
```

```
Мы получим 0, потому что внутри блока if объявляется (:=) копия n
которая удаляется после завершения блока
```

### Вопрос 13.

Что выведет данная программа и почему?

```
func someAction(v []int8, b int8) {
	v[0] = 100
	v = append(v, b)
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}
```

```
результат будет - [100 2 3 4 5]
v[0] = 100 изменяет исходный объект,
а v = append(...) создает копию
```

### Вопрос 14.

Что выведет данная программа и почему?

```
func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}
```

```
мы получим - [b b a][a a]
потому что сначала распечатаем слайс внутри блока func,
затем распечатаем слайс из функции main

append поместил в переменную слайс новое значение, а поэтому
исходный слайс не был модифицирован при обращении к его данным
по индексам элементов

если же убрать строку с append, то вывод будет [b b][b b]
т.к мы будем ссылаться на одни и те же данные в обоих вызовах принта
```

