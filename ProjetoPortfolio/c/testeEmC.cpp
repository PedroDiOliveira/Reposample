//////////////////////////////////////
//programa de Preenchimento de vetor//
//////////////////////////////////////

//Importaçâo de pacotes
#import <stdio.h>

//Declaraçâo de Variaveis
int vetor[10];
int i = 0;

//////////////////////////////////////////
//funçâo principal de execuçâo de codigo//
//////////////////////////////////////////
int main(){
	//limpeza da memoria
	for (i = 0;i < 10;i ++){
		vetor[i] = 0;
	}
	
	//Preenchimento do vetor
	for (i = 0;i < 10; i ++){
		printf("Digite o valor da posicao %d do vetor:", i);
		scanf("%d", &vetor[i]);
	}
	//Impressâo do vetor
	for (i = 0;i < 10;i ++){
		printf("vetor [%d] = %d\n", i, vetor[i]);
	}
}
 
