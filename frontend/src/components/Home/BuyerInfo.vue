<template>
  <section class="section">
    <div class="container">
      <div class="row">
        <div class="col-md-8">
          <v-card
          elevation="5"
          >
          <h1>Buyer Information</h1>
          <!-- <h3>Name: {{ activeBuyer.name }}</h3>
          <h3>Age: {{ activeBuyer.age }}</h3>
          <h3>Buyer Id: {{ activeBuyer.id }}</h3> -->
          <template
          padding-bottom=5px
          >
            <v-card
              class="mx-auto"
              max-width="344"
              outlined
            >
              <v-list-item three-line>
                <v-list-item-content>
                  <v-list-item-title class="headline mb-1">
                    {{ activeBuyer.name }}
                  </v-list-item-title>
                  <v-list-item-subtitle>Age: {{ activeBuyer.age }}</v-list-item-subtitle>
                  <v-list-item-subtitle>Buyer Id: {{ activeBuyer.id }}</v-list-item-subtitle>
                </v-list-item-content>
              </v-list-item>
            </v-card>
          </template>
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
