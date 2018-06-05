# Магия, не трогать!

## Объекты
```
type Station struct {
	id	int64
        name    []string
        color   []int
        x       float32
        y       float32
        active  []bool
}

type Event struct {
        id      int             
        name    string          
        time    int
	stations []int
}
```

## Методы

`GET /events/` -- Возвращает все события и список станций, доступных на момент события
`GET /station/{id}` -- Возвращает подробную информацию о станции с id = {id}



