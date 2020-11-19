<template>
  <v-app>
    <v-main>
      <router-view name="header" />
      <div class="router-view">
        <router-view />
      </div>
      <router-view name="footer" />
    </v-main>
  </v-app>
</template>
<script>
export default {};
</script>
<style scoped lang="scss">
.router-view {
  min-height: calc(100vh - 179px);
}
</style>
