def jogar():
    print("******************************")
    print("**Bem vindo ao Jogo Da Forca**")
    print("******************************")

    palavra_secreta = "banana"

    enforcou = False
    acertou = False
    letras_acertadas = ["_","_","_","_","_","_"]
    
    print(letras_acertadas)

    while(not enforcou and not acertou):

        index = 0
        chute = input("Qual letra?")
        chute = chute.strip()


        for letra in palavra_secreta:
            
            if (chute.upper() == letra.upper()):
                letras_acertadas[index] = chute
            index = index + 1
        print(letras_acertadas)
        if (letras_acertadas.count("_") == 0):
            break
    print("Fim De Jogo!")

if (__name__ == "__main__"):
    jogar()