<script setup>
  import Header from './components/header.vue'
</script>

<template>
  <main>
      <Header />
      <RouterView />
  </main>
</template>

<style scoped>
  main{
      margin: 0 15vw 15vw 15vw;
  }
</style>
