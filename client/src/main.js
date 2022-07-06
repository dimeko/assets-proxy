import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './css/main.scss'

Vue.config.productionTip = false
Vue.prototype.$apiServerUrl = `${window.location.protocol}//${window.location.hostname}:${process.env.API_SERVER_PORT}`

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')