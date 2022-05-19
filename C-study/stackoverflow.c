#include <stdio.h>

int main(int argc, char* argv[]){
    int i = 0;
    printf("%p\n",&i);
    int arr[3] = {0};
    for(; i<=3; i++){
        printf("%p\n",&(arr[i]));
        arr[i] = 0;
        printf("hello world\n");
    }
    return 0;
}