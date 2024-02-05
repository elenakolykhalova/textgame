Текстовая игра
==============

Мы пишем простую игру, которая реагирует на команды игрока. 
Игровой мир обычно состоит из локаций, где может происходить какое-то действие.
Так же у нас есть игрок. Как у игрока,
так и у локации есть состояние.

initGame делает нового игрока и задаёт ему начальное состояние. 
В данной версии можно обойтись глобальными переменными для игрока и мира(локаций). 
Команда в handleCommand парсится как имя команды и список необязательных параметров. 
`команда параметр1 параметр2 параметр3 ...`

Задача - сделать правильно. Под правильным понимается универсально, 
чтобы можно было без проблем что-то добавить или убрать. 
Т.е. бесконечный набор захардкоженных if'ов для всего мира не подойдёт. 
Конкретные условия могут быть только внутри конкретной локации. 
Надо думать в сторону объектов, вызова функций, структур, которые описывают состояние локации 
и игрока, функций которые описывают какой-то интерактив в локации. 
Глобальная мапа с полной командой от юзера - это тоже считается за хардкод.

В идеале ваша архитектура должна без проблем пережить добавление дополнительных локаций, предметов и команд.

Команды игрока:
- осмотреться
- идти <имя локации>
- взять <имя предмета>

Локации:
- кухня
- коридор
- комната
- улица

Предметы:
- рюкзак
- конспекты
- ключи

Тест-кейсы:
```
______________________________________________________
> осмотреться
ты находишься на кухне, надо собрать рюкзак и идти в универ. можно пройти - коридор

> завтракать
неизвестная команда

> идти комната
нет пути в комната

> идти коридор
ничего интересного. можно пройти - кухня, комната, улица

> идти комната
ты в своей комнате. можно пройти - коридор

> осмотреться
на столе: ключи, конспекты, рюкзак. можно пройти - коридор

> взять рюкзак
предмет добавлен в инвентарь: рюкзак

> осмотреться
на столе: ключи, конспекты. можно пройти - коридор

> взять ключи
предмет добавлен в инвентарь: ключи

> взять телефон
нет такого

> взять ключи
нет такого

> осмотреться
на столе: конспекты. можно пройти - коридор

> взять конспекты
предмет добавлен в инвентарь: конспекты

> осмотреться
пустая комната. можно пройти - коридор

> идти коридор
ничего интересного. можно пройти - кухня, комната, улица

> идти кухня
кухня, ничего интересного. можно пройти - коридор

> осмотреться
ты находишься на кухне, надо собрать рюкзак и идти в универ. можно пройти - коридор

> идти коридор
ничего интересного. можно пройти - кухня, комната, улица

> идти улица
на улице весна. можно пройти - домой
```

Результат

Консольное приложение с игрой.

Требования

Допускается решение на одном из языков программирования: C#, Java, Python, JS, C/C++.
Для реализации игры разрешается использовать только стандартные библиотеки, пакеты для вашего языка программирования.
Тесты, комментарии и пояснения к коду приветствуются.
*Для решения на go
Версия 1.21+
