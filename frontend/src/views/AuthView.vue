<script setup>
import { reactive} from "vue";
import axios from "axios";

let users = reactive({
    email:'',
    password:''
})

let submit=()=>{
    axios.post('http://127.0.0.1:8080/user/auth', users, {
        headers:{
            "Content-Type":"application/json"
        }
    })
        .then(response => {
            console.log(response)
        })
        .catch(error => {
            if(error.status == 200) {
                console.log('Good')
            }else{
                console.log('Not good')
            }
        });
}
</script>

<template>
  <div class="container flex mt-10 mb-10 mx-auto bg-gray-200">
    <form class="w-1/2 h-full relative" @submit.prevent.stop="submit">
        <h1 class="text-center text-5xl font-bold text-red mt-20">Авторизация</h1>
        <div class="ms-20 mt-10 me-20">
            <p class="text-3xl pb-2 font-bold">Введите почту или логин</p>
            <input type="text" v-model="users.loginOrEmail" class="border-2 p-2 border-gray-400 w-full"/>
        </div>

        <div class="ms-20 mt-10 me-20">
            <p class="text-3xl pb-2 font-bold">Введите пароль</p>
            <input type="password" v-model="users.password" class="border-2 p-2 border-gray-400 w-full"/>
        </div>
        <div class="ms-20 mt-10 me-20 absolute bottom-[10vh] left-0 right-0">
          <button type="submit" class="p-2 w-full bg-emerald-400 text-white font-bold text-2xl">Авторизация</button>
        </div>
    </form>
    <img
        class="w-1/2 h-full"
        src="https://sun9-48.userapi.com/impg/tOX_17XaHfP8hBkC1y7C3WG1D0UzkpEUXgDp2g/TxoyM5xbD1Q.jpg?size=564x564&quality=95&sign=1abfa9d4ab21347c1bf6d81ec8da88be&type=album"/>
  </div>
</template>

<style scoped>

</style>