<template>
  <div class="login_container">
    <div class="login_modal">
      <form @submit.prevent="submit" class="login_form">
        <div class="login_title">Master assets proxy</div>
        <div class="form_field">
          <label for="username">Username</label>
          <input v-model="fInput.username" />
        </div>
        <div class="form_field">
          <label for="password">Password</label>
          <input type="password" v-model="fInput.password" />
        </div>
        <div class="form_button">
          <button class="form_button" type="submit">Submit</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "Login",
  data() {
    return {
      fInput: {
        username: "",
        password: "",
      },
      baseApiUrl: "/api/user/login",
    };
  },
  methods: {
    submit() {
      axios
        .post(this.baseApiUrl, {
          username: this.fInput.username,
          password: this.fInput.password,
        })
        .then(() => {
          this.$router.push({ name: "dashboard" });
        })
        .catch(() => {
          this.$router.push({ name: "dashboard" });
        });
    },
  },
};
</script>
