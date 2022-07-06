<template>
  <div class="notifications_wrapper">
    <div v-show="visible" :class="bodyClass" class="notifications_toast">
      <div>
        {{ title }}
      </div>
      <div>
        {{ desription }}
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
    this.$root.on("notify", (title, description, bodyClass) => {
      this.title = title;
      this.description = description;
      this.bodyClass = bodyClass;
      this.visible = true;
      setTimeout(() => {
        this.visible = false;
      }, 7000);
    });
  },
};
</script>
