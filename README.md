# jetstyle-rest-api
REST API for jetstyle test  

Документация: http://localhost:8080/swagger/index.html

### Run
 * make build ; make up - Windows
 * make build & make up - Unix

При создание пользоваетля не забыть указать -db:host т.к конфиг по умолчанию подразумевает исполнение в контейнере.  
Пример: go run .\cmd\app\main.go create -l login -p password -db:host localhost
  
Всего существует две сабкоманды.
 * serve - старт сервера
 ```
 go run .\cmd\app\main.go serve --help 
Usage of serve:
  -bind-addr string
        server address (default ":8080")
  -db:SSLMode string
        database SSLMode (default "disable")
  -db:host string
        database host (default "db")
  -db:name string
        database name (default "jetstyle")
  -db:password string
        database password (default "123")
  -db:port string
        database port (default "5432")
  -db:user string
        database user (default "postgres")
```
 * create - создание юзера
```
go run .\cmd\app\main.go create --help 
Usage of create:
  -db:SSLMode string
        database SSLMode (default "disable")
  -db:host string
        database host (default "db")
  -db:name string
        database name (default "jetstyle")
  -db:password string
        database password (default "123")
  -db:port string
        database port (default "5432")
  -db:user string
        database user (default "postgres")
  -l string
        user login
  -p string
        user password
 ```
Все значения по умолчанию берутся из конфига (configs\configs.toml)  
Задать конфиг по умолчанию
```
go run .\cmd\app\main.go --help
  -config-path string
        path to config file (default "configs/config.toml")
```
### Тестирование  
Тесты написаны для хендлеров и репозиториев.  
Для репозиториев используется тестовой инстанс бд c конфигом по умолчанию. (tests\psql\configs\config.toml)  
Для хендлеров просто key-value хранилище.
 * make store ; make test - Windows
 * make store && make test - Unix
  
### Валидация 
Есть ограничение на длину пароля и поля name для тасков.  
К read-only полям пользователь просто не имеет доступа, все они изменяются с помощью тригеров.  
Поле update_date изменяется только тогда, когда произошли реальные изменения, конкретно поля owner или name.  
Изменение статуса задачи так же не влияет на update_date т.к эту дату хранит complete_date.  
  
### Реализация
Старался использовать сторонние пакеты по минимуму. По сути можно выделить только пакет github.com/gorilla/mux.   
Нет никакой системы мониторинга.  
Реализовано несколько простых middleware'ор для логирования и авторизации.  
