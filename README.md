# crud-user-desafio

Essa é uma api simples de cadastro de úsuario feita durante meus estudos de Golang e Angular.


**Crud Go**

com go instalado e configurado em sua maquina adicione os modúlos

go get -u github.com/gorilla/mux

go get go.mongodb.org/mongo-driver/mongo



**MongoDb**

Inicie o mongodb em sua máquina na porta padrão :27017

Crie uma database chamada *users* com uma collection chamada *user*;

Insira  os campos na collection criada:

**_id**  do tipo ObjectId 

**nome**  do tipo string

**email** do tipo string

**senha** do tipo string



**Inicie a Api**

abra o terminal na pasta *api-user* do projeto e inicie com o comando 

$ go run api-user.go


A api está configurada na porta :7000 e roteada em /usuarios

http://localhost:7000/usuarios

obs. eu utilizei o POSTMAN para ler(GET) e inserir(POST) os dados.



**Front-end**
Utilizei o Angular com material

Abra a pasta *frontend* do projeto e inicie 

$ npm init

Baixe as dependências node 

$ npm install

De o start

$ npm start

Vizualize em seu navegador

http://localhost:4200


**:)**

Desde já quero pontuar: estou estudando sobre as tecnologias utilizadas nesse projeto; sei que preciso aprender muitas coisas, desde de padrões,
segurança e até boas práticas. Se você tiver dicas e até projetos que possa compartilhar, ficaria agradecido. :)
