//////////////////////////////////////
//programa de Preenchimento de vetor//
//////////////////////////////////////

//Importaçâo de pacotes
#import <stdio.h>

//Declaraçâo de Variaveis
int vetor[10];
int vetor2[10];
int i = 0;

//////////////////////////////////////////
//funçâo principal de execuçâo de codigo//
//////////////////////////////////////////
int main(){
	//limpeza da memoria
	for (i = 0;i < 10;i ++){
		vetor[i] = 0;
		vetor2[i]= 0;
	}
	
	//Preenchimento do vetor
	for (i = 0;i < 10; i ++){
		printf("Digite o valor da posicao %d do vetor:", i);
		scanf("%d", &vetor[i]);
	}
	
	//invercao do vetor
	for (i = 9;i >= 0; i --){
		vetor2[9 - i] = vetor[i]; 
	}
	
	//impressao do vertor original
	for (i = 0;i < 10; i ++){
		printf("vetor[%d] = %d\n", i, vetor[i]);
	}
	//impressao do vetor invertido
	for (i = 0;i < 10; i ++){
		printf("vetor2[%d] = %d\n", i, vetor2[i]);
	}
}
