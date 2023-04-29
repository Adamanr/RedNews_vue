<script setup>
import {reactive} from "vue";
import axios from "axios";

let statuss = reactive({
    text:'None'
})

let user = reactive({
    login:'',
    password:'',
    email:''
})

let submit = ()=>{
    axios.post('http://127.0.0.1:8080/user/reg', user, {
        headers:{
            "Content-Type":"application/json"
        }
    })
        .then(response => {
            statuss.status=response.status
        })
        .catch(error => {
            statuss.text = error.status
            if(error.status == 200) {
                console.log('Good')
            }else{
                console.log('Not good')
            }
        });
}
</script>

<template>
    <div class="flex w-full">
        <div class="flex w-1/2 bg-[#f0f1ec]">
          <form @submit.prevent.stop="submit" class="container m-24">
            <h1 class="text-center text-5xl pt-4 font-bold text-red">Регистрация</h1>
              <h1>Статус: {{ statuss.text }}</h1>
              <div class="field m-10">
                  <p class="text-3xl text-red">Логин</p>
                  <input type="text" v-model="user.login" placeholder="Введите ваш логин" class="border-2 w-full border-gray-300 p-2"/>
              </div>
              <div class="field m-10">
                  <p class="text-3xl text-red">Почта</p>
                  <input type="text" v-model="user.email" placeholder="Введите вашу почту" class="border-2 w-full border-gray-300 p-2"/>
              </div>

              <div class="field m-10">
                  <p class="text-3xl text-red">Пароль</p>
                  <input type="text" v-model="user.password" placeholder="Введите ваш пароль" class="border-2 w-full border-gray-300 p-2"/>
              </div>

              <div class="field m-10 mt-20">
                  <button
                      @submit.prevent.stop="submit"
                      class="w-full rounded-lg text-white text-3xl bg-emerald-400 p-2 font-bold">Зарегистрироваться</button>
              </div>
          </form>
        </div>
        <img
            class="flex w-1/2"
            src="https://sun3-23.userapi.com/impg/qo3bslmYuEFQXTGNScRQ1jvkHEbrzfcZcu5aew/-77FTfgDpp0.jpg?size=1000x1000&quality=95&sign=60810a3975ae4aec02fef0893ec19fde&type=album"/>
    </div>
</template>

<style scoped>
    body {
        background-color: #f0f1ec;
    }
</style>