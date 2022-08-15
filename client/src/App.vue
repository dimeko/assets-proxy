<template>
  <div id="app">
    <notification-toast></notification-toast>
    <router-view></router-view>
  </div>
</template>

<script>
import notifications_toast from "./components/Notification.vue"
export default {
  name: 'App',
  components: {
    NotificationToast: notifications_toast
  }
}
</script>

<style>
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
