
Основные виды очередей продемонстрированы в папке services/broker/init:
- Limited queue (services/broker/init/limited_queue)
- Queue with dead_letter_queue (services/broker/init/dead_letter)

Чтобы запустить нужный вам очередь, сначала поменяйте конструктор очереди на файле
```
services/broker/init/cmd/main.go
```
затем запустите команду:
```
go run services/broker/init/cmd/main.go
```