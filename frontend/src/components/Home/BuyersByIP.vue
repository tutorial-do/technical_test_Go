<template>
  <v-card
    elevation="5"
  >
    <v-card-text>
      <h1>Buyers by IP: {{ currentIP }}</h1>
    </v-card-text>
    <template>
      <v-data-table
        :headers="headers"
        :items="buyersIP"
        :items-per-page="5"
        class="elevation-1"
      ></v-data-table>
    </template>
  </v-card>
</template>
<script>

import { mapState } from 'vuex';

export default {
  name: 'BuyersByIP',
  props: {
    data: {
      type: [Array, String],
      default: '',
      required: false,
    },
  },
  data() {
    return {
      headers: [
        {
          text: 'Name',
          align: 'start',
          sortable: false,
          value: 'name',
        },
        { text: 'ID', value: 'id' },
        { text: 'Age', value: 'age' },
      ],
    };
  },
  computed: mapState({
    buyersIP: (state) => state.buyersIP,
    currentIP: (state) => state.currentIP,
  }),
  methods: {
  },
};
</script>
<style scoped lang="scss">
</style>
