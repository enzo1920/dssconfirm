# Programm for dss confirmation

## Последовательность вызовов
 
 1. Аутентификация пользователя
 2. Отправка запроса
 3. Ожидание ответа




## 1.Аутентификация CriptoAuth
Из ответа сервера необходимо  извлечь значение параметра access_token (JWT токен).
Он будет использован для аутентификации пользователя в последующих запросах.
В запросе используются параметры:

```
 - client_id
 - Resource - Идентификатор ресурса
 - phone
 - url
```
## 2.Отправка запроса StartReq

В запросе передаются следующие параметры:

```
 - в заголовке Authorization передаётся access_token, полученный на предыдущем шаге.
 - Resource - Идентификатор ресурса
 - ConfirmationScope - Идентификатор операции
 - CallbackUri - (опциональный) адрес Callback для асинхронного ожидания// не передается
```
В функции передается token, url, client_id
Из ответа сервера необходимо извлечь значение параметра
Challenge -> ContextData -> RefID
Параметр определяет Идентификатор запроса. 
Он используется в последующих запросах.

Так же необходимо проверить значение параметра
```
 IsError
```
Если он равен true, то отправка запроса  была не успешной.
Ответ сервера будет содержать поля
```
 - Error  				- код ошибки
 - ErrorDescription     - текстовое описание ошибки
```
## 3. Проверка статуса запроса
В запросе передаются следующие параметры:
```
 - в заголовке Authorization передаётся access_token, полученный на предыдущем шаге.
 - Resource - Идентификатор ресурса
 - Идентификатор операции, полученный на предыдущем шаге (Challenge -> ContextData -> RefID)
```
ResponseCheck(authtoken string, url string, refid string)
При обработке ответа вызывающее приложение должно смотреть на значение двух флагов: IsFinal и IsError.
Если получен ответ с IsError - true, то дальнейшая обработка не возможна.
Если получен ответ с IsFinal - false, то подтверждение транзакции ещё не завершено.

То есть в случае получения успешного ответа  севрер вернёт 
```
IsError = false, IsFianl = true.
```

## Compile
Используется Makefile
```
make windows
or
make linux 
```

## Запуск
```
./dssconfirm -m MSISDN
```

## Создание статической библиотеки для C
http://blog.ralch.com/tutorial/golang-sharing-libraries/
import "C"в пакете main
```
go install -buildmode=shared -linkshared std
cd c_lib/
go build -buildmode=c-shared -o kryptoesim.a kryptoesim.go

```


## Log 



## Authors

***Sergei*** - - https://github.com/enzo1920/
