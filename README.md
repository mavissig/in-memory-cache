# In-Memory кэш на Go

**In-Memory** – это технология, обеспечивающая загрузку данных из источников в оперативную память с последующим 
обращением к ним «на лету» без необходимости обращения к диску или СУБД, за счёт чего значительно увеличивается 
скорость доступа к данным и их обработки.

В нашем случае для реализации используется `map[string]any`

## Реализованные методы

- `New()` - создание нового кэша
    > cache := New()
- `Set()` - добавление элемента в кэш 
    > cache.Set("key_name", "value")
- `Get()` - получение элемента из кэша 
    > val, ok := cache.Get("key_name")
- `Remove()` - удаление элемента из кэша
    > cache.Remove("key_name")