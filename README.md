# vite_grpc_problem

**frontend**
Шаблон:
```
npx degit sveltejs/template .
```

Необходимо сделать аналогичное для 
```
npm create vite@latest . -- --template svelte
```
<br/>
<br/>

В src/App.svelte заменить 
```
const clt = new LotsServiceClient("<host>:<port>");
```
<br/>

```
make install
make run
```
<br/>

**backend**

```
make install
make run
```
