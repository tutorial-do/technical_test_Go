<template>
  <section class="section">
    <div class="container">
      <div class="row">
        <div class="col-md-12">
          <v-card
          elevation="5"
          >
          <h1>Products By Transaction</h1>
          <v-data-table
            :headers="headers"
            :items="productsByTransaction"
            :items-per-page="5"
            class="elevation-1"
          ></v-data-table>
          </v-card>
        </div>
      </div>
    </div>
  </section>
</template>
<script>

import { mapState } from 'vuex';

export default {
  name: 'ProductsByTransaction',
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
          text: 'Product Name',
          align: 'start',
          sortable: false,
          value: 'name',
        },
        { text: 'Product ID', value: 'id' },
        { text: 'Product Price $', value: 'price' },
      ],
    };
  },
  computed: mapState({
    productsByTransaction: (state) => state.productsByTransaction,
  }),
  mounted() {},
  methods: {},
};
</script>
<style scoped lang="scss">
</style>
