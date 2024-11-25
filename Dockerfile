# Use uma imagem base com Node.js
FROM node:18-alpine

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie os arquivos de configuração para o contêiner
COPY package.json package-lock.json ./

# Instale as dependências
RUN npm install

# Copie o restante do código do projeto
COPY . .

# Exponha a porta usada pelo Vite
EXPOSE 5173

# Comando padrão para rodar o servidor de desenvolvimento
CMD ["npm", "run", "dev"]