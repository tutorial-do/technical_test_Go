<template>
  <section class="section">
    <div class="container">
      <div class="row">
        <div class="col-md-12">
          <v-card
          elevation="5"
          >
          <h1>Buyer Information</h1>
          <h3>name: {{ activeBuyer.name }}</h3>
          <h3>Age: {{ activeBuyer.age }}</h3>
          <h3>Buyer Id: {{ activeBuyer.id }}</h3>
          </v-card>
        </div>
      </div>
    </div>
  </section>
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
