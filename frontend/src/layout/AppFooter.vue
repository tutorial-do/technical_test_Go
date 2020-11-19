<template>
  <v-footer
    dark
    padless
  >
    <v-card
      flat
      tile
      class="light-blue lighten-1 white--text text-center flex"
    >
      <v-card-text>
        <!-- <v-btn
          v-for="icon in icons"
          :key="icon"
          class="mx-4 white--text"
          icon
        >
          <v-icon size="24px">
            {{ icon }}
          </v-icon>
        </v-btn> -->
        <v-btn
        class="mx-4 white--text"
        href="https://www.linkedin.com/in/julianfrancor"
        icon
        >
          <v-icon size="24px">mdi-linkedin</v-icon>
        </v-btn>
        <v-btn
        class="mx-4 white--text"
        href="https://twitter.com/JulianFrancoR"
        icon
        >
          <v-icon size="24px">mdi-twitter</v-icon>
        </v-btn>
        <v-btn
        class="mx-4 white--text"
        href="https://github.com/julianfrancor"
        icon
        >
          <v-icon size="24px">mdi-github</v-icon>
        </v-btn>
      </v-card-text>
      <v-divider></v-divider>
      <v-card-text class="white--text pt-2">
        {{ new Date().getFullYear() }} - Julian Franco
      </v-card-text>
    </v-card>
  </v-footer>
</template>
<script>
export default {
  name: 'AppFooter',
  data() {
    return {
      icons: [
        'mdi-twitter',
        'mdi-linkedin',
        'mdi-github',
      ],
      year: new Date().getFullYear(),
    };
  },
};
</script>
<style>
</style>
