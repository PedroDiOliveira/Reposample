//Variaveis da Bolinha
let xBolinha = 300;
let yBolinha = 200;
let diametro = 20;
let raio = diametro / 2;
//Variaveis da Raquete
let xRaquete = 5;
let yRaquete = 150;
let larguraRaquete = 10;
let alturaRaquete = 90;
//Variaveis do Oponente
let xOponente = 585
let yOponente = 150
//Velocidade da Bolinha
let velocidadeXBolinha = 6;
let velocidadeYBolinha = 2;
function setup() {
  createCanvas(600, 400);
}
function draw() {
  background(0);
//Bolinha
  mostraBolinha();
  movimentaBolinha();
  verificaColisaoBorda();  
  verificaColisaoRaquete();
  //Raquete
  mostraRaquete();
  movimentaMinhaRaquete();
//Oponente
  mostraOponente();  
}
function mostraBolinha(){
circle (xBolinha,yBolinha,diametro)
}
function movimentaBolinha(){
xBolinha += velocidadeXBolinha
yBolinha += velocidadeYBolinha 
}
function verificaColisaoBorda(){
  if(xBolinha + raio > width || xBolinha -raio < 0){
velocidadeXBolinha *= -1
}   
  if(yBolinha + raio > height ||
     yBolinha - raio < 0){
velocidadeYBolinha *= -1  
  }
}
function mostraRaquete(){
  rect(xRaquete, yRaquete, larguraRaquete, alturaRaquete)
}
function mostraOponente(){
  rect(xOponente, yOponente, larguraRaquete, alturaRaquete)
}
function movimentaMinhaRaquete() {
    if (keyIsDown(UP_ARROW)) {
        yRaquete -= 10;
    }
    if (keyIsDown(DOWN_ARROW)) {
        yRaquete += 10;
    }
}
function verificaColisaoRaquete() {
    if (xBolinha - raio < xRaquete + larguraRaquete && yBolinha - raio < yRaquete + alturaRaquete && yBolinha + raio > yRaquete) {
        velocidadeXBolinha *= -1;
    }
}
