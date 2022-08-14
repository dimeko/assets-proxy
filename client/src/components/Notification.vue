<template>
  <div class="notifications_wrapper">
    <div v-show="visible" :class="bodyClass" class="notifications_toast">
      <div class="notifications_toast_title">
        {{ title }}
      </div>
      <div>
        {{ description }}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      title: "",
      description: "",
      visible: false,
    };
  },
  mounted() {
    this.$root.$on("notify", ({title, description, bodyClass}) => {
      console.log("Emited")
      this.title = title;
      this.description = description;
      this.bodyClass = bodyClass;
      this.visible = true;
      setTimeout(() => {
        this.visible = false;
      }, 5000);
    });
  },
};
</script>
