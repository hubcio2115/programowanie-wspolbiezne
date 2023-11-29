**Zadanie 4.** (czas: na 15.11.2023)

1. **(8 punktów)**. Wykorzystując mechanizm kolejek FIFO, napisać program serwera prostej wielodostępnej "bazy danych" oraz program klienta tej usługi.

   Baza danych zawierająca struktury postaci (`ID,nazwisko`) gdzie identyfikator `ID` to int a `nazwisko` to napis (można przyjąć ograniczenie długości napisu do 20 znaków, wystarczy kilka zapisów w testowej "bazie danych")

   Baza danych jest tworzona i wypełniana statycznie przez program serwera (np. lista). Serwer i klienci komunikują się za pomocą **kolejek FIFO**. Serwer odbiera ze swojej kolejki wejściowej (**wspólnej dla wszystkich klientów**) zapytanie zawierające ID i ścieżkę do kolejki klienta (**każdy klient ma swoją kolejkę**), do której ma wstawić odpowiedź. Serwer odsyła do kolejki klienta nazwisko odpowiadające przysłanemu identyfikatorowi lub komunikat "Nie ma", jeżeli nie ma w bazie takiego ID. (Szkic komunikacji)

   Serwer działa w pętli nieskończonej, klient tylko jednorazowo (pytanie - odpowiedź).

2. **(2 punkty)**. Dodatkowo, zapewnić w programie serwera **przechwytywanie sygnałów** SIGHUP oraz SIGTERM (powinno to umożliwić działanie serwera nawet po zamknięciu okna, w którym go uruchomiono). Natomiast sygnał SIGUSR1 wysłany do serwera powinien kończyć jego działanie.

**UWAGA 1**. zapewnić niepodzielność wysyłania komunikatów (czyli użyć pojedynczego wywołania funkcji write() do wysłania całego komunikatu).

Mogą to być komunikaty o stałej długości, ale wtedy trzeba rozpoznawać koniec komunikatu krótszego niż o maksymalnej długości.

Można też w treści komunikatu umieścić informacje o jego długośc. Przykładowa struktura komunikatu od klienta do serwera:

|                 int                  | int |  napis (zmienna długość)   |
| :----------------------------------: | :-: | :------------------------: |
| długość pozostałej części komunikatu | ID  | ścieżka do kolejki klienta |

\
\
Przykładowa struktura komunikatu od serwera do klienta:

|                 int                  | napis (zmienna długość) |
| :----------------------------------: | :---------------------: |
| długość pozostałej części komunikatu |        odpowiedź        |

**UWAGA 2.** Testowanie powinno sprawdzić działanie programu również wtedy, gdy w kolejce serwera będzie więcej niż jedno zapytanie - w tym celu trzeba wprowadzić spowolnienie działania serwera przed udzieleniem odpowiedzi, żeby zdążyć wysłać dodatkowe zapytania od jeszcze dwóch klientów

**UWAGA 3.** Ponieważ używamy kolejek FIFO z blokującym odczytem w przypadku pustej kolejki, nieoptymalną praktyką jest stosowanie aktywnego czekania na pojawienie się czegoś w kolejce (przez sprawdzanie w pętli czy jest niepusta)

**UWAGA 4.** Należy również przetestować obsługę sygnałów

**Wskazówki** - materiały z wykładu:

- [Potoki (kolejki FIFO)](https://inf.ug.edu.pl/~pmp/Z/Wspolb21P/wspolb5p.pdf). Przykłady: [fifoS1.py](https://inf.ug.edu.pl/~pmp/Z/Wspolb21P/fifoS1.py) , [fifoK1.py](https://inf.ug.edu.pl/~pmp/Z/Wspolb21P/fifoK1.py)
- [Obsługa sygnałów](https://inf.ug.edu.pl/~pmp/Z/Wspolb21P/wspolbSygnPnotatki21p.pdf). Przykład: [petlaS1.py](https://inf.ug.edu.pl/~pmp/Z/Wspolb21P/petlaS1.py)
