**Zadanie 5.** (czas: na 22.11.2023)

Utworzyć program serwera udostępniającego dwie **kolejki komunikatów IPC**: wejściową i wyjściową, oraz program klienta przesyłającego pojedyncze zapytanie z PID-em swojego procesu jako _typem komunikatu_, a następnie odbierającego odpowiedź z drugiej kolejki. (Obie kolejki wspólne dla serwera i wszystkich klientów)

Serwer realizuje funkcję słownika polsko-angielskiego czyli otrzymuje napis zawierający słowo polskie i odsyła odpowiadające mu słowo angielskie lub komunikat "Nie znam takiego słowa". Do testowania wystarczy kilka słów w słowniku.

UWAGI

- w Pythonie trzeba zainstalować moduł z kolejkami komunikatów IPC. Są dwie możliwości
  - https://semanchuk.com/philip/sysv_ipc/ (wersja preferowana, bo będziemy tego pakietu jeszcze używać. Instalacja: pip sysv_ipc)
  - https://pypi.org/project/ipcqueue/ (wersja System V a nie POSIX)
- przyjąć, że numery (klucze) obu kolejek serwera są ustalone i znane klientowi.
- ponieważ są osobne kolejki wejściowa i wyjściowa, serwer może odbierać komunikaty dowolnego typu (parametr typ = 0), a odpowiedzi do klientów mogę iść do wspólnej kolejki bo rozróżni je typ będący numerem procesu klienta
- w czasie testowania uwzględnić sytuację, w której dwóch klientów umieściło zapytania w kolejce - trzeba wprowadzić opóźnienie w serwerze, żeby uzyskać taką sytuację

Wskazówka:

- materiały z wykładu: [Kolejki komunikatów IPC](https://inf.ug.edu.pl/~pmp/Z/Wspolb21P/wspolb6pNotatki2022p.pdf). Przykład: [ipcS.py](https://inf.ug.edu.pl/~pmp/Z/Wspolb21P/ipcS.py) , [ipcK1.py](https://inf.ug.edu.pl/~pmp/Z/Wspolb21P/ipcK1.py), [ipcK2y.py](https://inf.ug.edu.pl/~pmp/Z/Wspolb21P/ipcK2.py), [ipcS1.py](https://inf.ug.edu.pl/~pmp/Z/Wspolb21P/ipcS1.py)
- podobne przykłady w C: [kolejki IPC - przykłady w C](https://inf.ug.edu.pl/~pmp/Z/Wspolb22/kolejkiIPC.html)
