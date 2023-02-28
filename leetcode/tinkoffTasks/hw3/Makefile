OUT	= candles_10m.csv candles_1m.csv  candles_2m.csv

EXEC = main

all:	clean
		go build main.go
		./main

clean:
		rm -f $(OUT)

fclean:	clean
		rm -f ./main

.PHONY:	clean, fclean, all