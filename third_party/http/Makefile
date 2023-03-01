all:httpd.exe client.exe

httpd.exe : httpd.c
    gcc -g -Wno-unused-result -Wno-int-to-pointer-cast -Wno-pointer-to-int-cast -Wall -O2 -o $@ $^ -lpthread

client.exe : client.c
    gcc -g -Wall -o $@ $^

clean:
    rm -rf *.exe