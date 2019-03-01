#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <time.h>
#include <string.h>

int main() {
/* declarar as variáveis */
 struct sockaddr_in exemplo;
 int socket_entrada, socket_conexao;
 time_t data_e_hora;
 int tamanho_estrutura_socket;
 char *data_e_hora_por_extenso;

/* inicializar o socket */
 socket_entrada = socket(AF_INET, SOCK_STREAM, 0);
 exemplo.sin_family = AF_INET;
 exemplo.sin_port = htons(8000);
 exemplo.sin_addr.s_addr = INADDR_ANY;
 tamanho_estrutura_socket = sizeof(exemplo);

 if(bind(socket_entrada, (struct sockaddr *)&exemplo, tamanho_estrutura_socket)==0) {
  listen(socket_entrada, 0);
  while(1) {
/* receber conexões */
   socket_conexao=accept(socket_entrada, (struct sockaddr *)&exemplo, &tamanho_estrutura_socket);
   data_e_hora = time(NULL);

/* pegar a hora do sistema e mandar */
   data_e_hora_por_extenso = ctime(&data_e_hora);
   printf("Cliente conectado em: %s\n", data_e_hora_por_extenso);
   write(socket_conexao, data_e_hora_por_extenso, strlen(data_e_hora_por_extenso));

/* desconectar */
   close(socket_conexao);
  }
  close(socket_entrada);
 } else {
  printf("Vixe! Nao deu para abrir a porta!\n");
 }
 return(0);
}









