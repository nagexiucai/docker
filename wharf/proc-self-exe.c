#include <unistd.h>
#include <stdio.h>

int main(int argc , char* argv[])
{
    char buf[1024] = {0};
    int n = 0;

    n = readlink("/proc/self/exe" , buf , sizeof(buf));
    if( n > 0 && n < sizeof(buf))
    {
        printf("%s\n" , buf);
    }
}
