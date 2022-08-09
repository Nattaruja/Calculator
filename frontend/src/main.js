import Vue from 'vue'
import VueRouter from 'vue-router';
import BootstrapVue from "bootstrap-vue"
import App from './App.vue'
import "bootstrap/dist/css/bootstrap.min.css"
import "bootstrap-vue/dist/bootstrap-vue.css"

Vue.config.productionTip = false
Vue.use(BootstrapVue)
Vue.use(VueRouter)

const routes = [
  {
    name: "Customers",
    path: "/customer",
    component: Customers
  },
  {
    name: "Items",
    path: "/items",
    component: Items
  }
];

const router = new VueRouter({ mode: "history", routes: routes });

new Vue({
  render: h => h(App),
  router
}).$mount('#app')
