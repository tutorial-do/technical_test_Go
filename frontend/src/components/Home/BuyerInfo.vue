<template>
  <v-card
  class="mx-auto"
  >
    <v-card-text>
      <h2
      class="h1"
      >Buyer Information</h2>
      <p class="display-2 text--primary mt-5">
        {{ activeBuyer.name }}
      </p>
      <p
      class="text--primary"
      >Age: {{ activeBuyer.age }}</p>
      <p
      class="text--primary">Buyer Id: {{ activeBuyer.id }}</p>
    </v-card-text>
  </v-card>
</template>
<script>

import { mapState } from 'vuex';

export default {
  name: 'BuyerInfo',
  props: {
    data: {},
  },
  computed: mapState({
    activeBuyer: (state) => state.activeBuyer,
  }),
};
</script>
<style scoped lang="scss">
</style>
