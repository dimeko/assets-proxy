<template>
  <div class="dashboard_container">
    <div class="dashboard">
      <div class="dashboard_header">
        <router-link :to="{ name: 'editor' }"><div>Editor</div></router-link>
        <div class="logout_button"><button @click="logout">Logout</button></div>
      </div>
      <div class="dashboard_editor">
        <div style="width: 50%; margin: 0 4px">
          <vue-json-editor
            v-model="jsonDir"
            :show-btns="true"
            :expandedOnStart="true"
          ></vue-json-editor>
        </div>
        <div style="width: 50%; margin: 0 4px">
          <vue-json-editor
            v-model="jsonImgDir"
            :show-btns="true"
            :expandedOnStart="true"
          ></vue-json-editor>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import vueJsonEditor from "vue-json-editor";
import axios from "axios";
const BASE_API_URL = "/api/server/";
const BASE_USER_URL = "/api/user/";
export default {
  name: "Dashboard",
  components: {
    vueJsonEditor,
  },
  data() {
    return {
      jsonDir: {},
      jsonImgDir: {},
    };
  },
  mounted() {
    this.mapDbDirectory();
    this.mapImgDirectory();
  },
  methods: {
    async mapDbDirectory() {
      const url = `${BASE_API_URL}map-db-directory`;
      const response = await axios.get(url);

      this.jsonDir = response.data;
    },
    async mapImgDirectory() {
      const url = `${BASE_API_URL}map-img-files-directory`;
      const response = await axios.get(url);

      this.jsonImgDir = response.data;
    },
    async logout() {
      const url = `${BASE_USER_URL}logout`;
      await axios.post(url).then(() => {
        this.$router.push({ name: "login" });
      }).catch(() => {
         this.$root.$emit("notify", {
            title: "Error",
            description: "Error occured while trying to logout!",
            bodyClass: "error_toast",
          });
      });
    }
  },
};
</script>
