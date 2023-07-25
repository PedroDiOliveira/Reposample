//////////////////////////////////////
//programa de Preenchimento de vetor//
//////////////////////////////////////

//Importaçâo de pacotes
#import <stdio.h>

//Declaraçâo de Variaveis
int matrizOriginal[2][3];
int matrizInvertida[3][2];
int i = 0;
int j = 0;

//////////////////////////////////////////
//funçâo principal de execuçâo de codigo//
//////////////////////////////////////////
int main(){
	//limpando matriz
	for(i = 0;i < 2;i ++){
		for(j = 0; i < 3; i ++){
			matrizOriginal[i][j] = 0;
			matrizInvertida[j][i] = 0;
		}
	}
	//preencher matriz original
	for(i = 0;i < 2;i ++){
		for(j = 0; j < 3; j ++){
			printf("Digite o valor da matriz[%d][%d] = ", i, j);
			scanf("%d", &matrizOriginal[i][j]);
		}
	}
	//transposicao propiamente dita
	for(i = 0; i < 2; i++){
		for(j = 0; j < 3; j ++){
			matrizInvertida[j][i] = matrizOriginal[i][j];
		}
	}
	
	//impressao da matriz
	for(i = 0;i < 2;i ++){
		for(j = 0; j < 3; j ++){
			printf("%d\t", matrizOriginal[i][j]);	
		}
		printf("\n");		
	}
	printf("\n");
	for(i = 0;i < 3;i ++){
		for(j = 0; j < 2; j ++){
			printf("%d\t", matrizInvertida[i][j]);	
		}
		printf("\n");
	}
}



