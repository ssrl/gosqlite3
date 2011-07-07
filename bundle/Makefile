all:
ifeq ($(GOOS),windows)
	gcc -shared -O2 -o ../sqlite3.dll sqlite3.c -Wl,--output-def,sqlite3.def,--out-implib,../libsqlite3.a
else ifeq ($(GOOS),darwin)
	gcc -c sqlite3.c
	ar cr ../libsqlite3.la sqlite3.o
	ranlib ../libsqlite3.la
	gcc -shared -Wl -ldl -lpthread -o ../libsqlite3.dylib sqlite3.o
else
	gcc -c sqlite3.c
	ar cr ../libsqlite3.a sqlite3.o
	ranlib ../libsqlite3.a
	gcc -shared -Wl,-soname,libsqlite3.so -ldl -lpthread -o ../libsqlite3.so sqlite3.o
endif
