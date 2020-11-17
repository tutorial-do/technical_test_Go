<template>
  <v-app>
    <v-main>
    <router-view name="header" />
    <router-view />
    <router-view name="footer" />
    </v-main>
  </v-app>
</template>
<script>
export default {};
</script>
