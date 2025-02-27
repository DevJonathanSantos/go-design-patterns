# **Abstract Factory em Go**

## 📌 Sobre o Abstract Factory

O **Abstract Factory** é um padrão de projeto **criacional** que resolve o problema de criar **famílias inteiras de produtos** sem precisar especificar suas classes concretas.

Ele funciona definindo uma **interface** para criar todos os produtos relacionados, mas delega a implementação para **fábricas concretas**, que produzem diferentes variações dos produtos.

## 🛠️ **Como Funciona?**

✅ O código cliente chama os métodos de uma **fábrica abstrata** em vez de instanciar objetos diretamente.  
✅ Cada fábrica concreta é responsável por criar uma **variante específica dos produtos**.  
✅ O código cliente trabalha apenas com **interfaces abstratas**, garantindo que seja compatível com qualquer variante de produtos.  
✅ Para adicionar uma nova variação de produtos, basta criar uma **nova fábrica concreta** sem alterar o código existente.

## 📌 **Benefícios do Abstract Factory**

🔹 **Desacoplamento** – O cliente não precisa conhecer as classes concretas dos produtos.  
🔹 **Facilidade de expansão** – Adicionar novos produtos ou fábricas não exige modificar o código existente.  
🔹 **Código mais limpo e organizado** – Mantém a responsabilidade da criação dos objetos isolada.

## 📂 **Exemplo de Estrutura**
