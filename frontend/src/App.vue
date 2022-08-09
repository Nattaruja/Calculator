<template>
  <div id="app">
    <MainSidebar/>
  <div class="page-content">
    <Calculator msg="Welcome to Your Vue.js App"/>
  </div>
  </div>
</template>

<script>
import Calculator from './components/Calculator.vue'
import './App.css';
import MainSidebar from './components/MainSidebar.vue'


export default {
  name: 'app',
  components: {
    Calculator,
    MainSidebar,
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
